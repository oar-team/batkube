package broker

import (
	"encoding/binary"
	"encoding/json"
	"math"
	"os"
	"time"

	zmq "github.com/pebbe/zmq4"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/pkg/translate"
)

const timeout = 1
const nonEmpty = 1 << 1

// Stop condition when waiting for new messages to send to Batsim. Ex :
// stopWaitingForMessages = nonEmpty | timeout will stop waiting for messages if
// 	- the waiting time since the last message from the scheduler is higher
// 	than timeoutValue
// 	- OR if the Events slice contained in the response is not empty.
var sendMessageCondition = nonEmpty | timeout
var timeoutValue = 200 * time.Millisecond

// Minimal amount of time to wait for messages from the scheduler.  Not having
// a minimal amount of waiting time leads to incorrect behavior from the
// scheduler. Most of the time, it gets stuck, though I am not sure why.
var minimalWaitDelay = 10 * time.Millisecond

// Set to true when a no_more_static_job_to_submit NOTIFY is received.
var noMoreJobs bool

// Number of pods in running or pending status.
var unfinishedJobs int

// Number of unanswered call_me_later events. If that number reached zero and
// an empty response is not excepted, batkube should not reply with an empty
// message.
var callMeLaters int

// It is ok to send empty messages (that is to say, with an empty event list)
// upon reception of certain events from batsim as Batsim will not error out on those.
// Those events are :
// - SIMULATION_BEGINS
// - NOTIFY
// - JOB_COMPLETED
// - SIMULATION_ENDS
//
// Note that empty messages may still be sent for the other types if the
// callMeLaters count is above 0, as Batsim will not error out on this case.
var expectedEmptyResponse bool

// Batkube received a SIMULATION_ENDS event.
var receivedSimulationEnded bool

/*
Handles time requests asynchronously. All sends and receives are non blocking

end : send a boolean on this channel to end the loop.
now : whenever now is updated, send it to this channel.
events : CALL_ME_LATER events are sent to this channel. They are stacked and sent whenever the receiver is ready.
*/
func handleTimeRequests(timeSock *zmq.Socket, end chan bool, now chan float64, events chan translate.Event) {
	if cap(now) != 1 {
		panic("now should be a buffered channel with capacity 1")
	}
	nowValue := <-now // Blocking receive : the value has to be initialized

	thisIsTheEnd := false
	for !thisIsTheEnd {
		// Get latest now value, if it is there
		// Note : to make message passing between time, batkube and
		// batsim synchronous again, make this a blocking receive.
		select {
		case nowValue = <-now:
		default:
		}

		_, err := timeSock.SendBytes([]byte("ready"), 0) // Tell time we're ready to receive
		timeMsgBytes, err := timeSock.RecvBytes(0)       //  Wait for next request from client
		durations := make([]int64, 0)
		if err != nil {
			panic("Error receiving message:" + err.Error())
		}
		if err = json.Unmarshal(timeMsgBytes, &durations); err != nil {
			panic("Could not unmarshal data:" + err.Error())
		}

		seen := make([]int64, 0)
		for _, d := range durations {
			if d <= 0 {
				panic("Got a negative duration")
			}

			// remove duplicates
			if isIn(d, seen) {
				continue
			}
			seen = append(seen, d)

			requested := addAndRound(nowValue, time.Duration(d))
			// TODO : can the contrary ever happen? What to do if it happens?
			if requested > nowValue {
				err, callMeLater := translate.MakeEvent(nowValue, "CALL_ME_LATER", translate.CallMeLaterData{Timestamp: requested})
				if err != nil {
					log.Panic("Failed to create event:", err)
				}
				// Non blocking send. Select isn't used to make sure it is actually sent.
				go func() {
					events <- callMeLater
				}()
			}
		}

		// Answer the time requests
		nowNano := uint64(nowValue * 1e9)
		b := make([]byte, 8)
		binary.LittleEndian.PutUint64(b, nowNano)
		_, err = timeSock.SendBytes(b, 0)
		if err != nil {
			panic("Error sending message: " + err.Error())
		}
		timeMsgBytes, err = timeSock.RecvBytes(0)
		if err != nil {
			panic(err)
		}
		done := string(timeMsgBytes)
		if done != "done" {
			log.Panicf("Failed exchange: Expected %s, got %s", "done", done)
		}

		select {
		case <-end:
			thisIsTheEnd = true
		default:
		}
	}
}

func Run(batEndpoint string) {
	var err error

	logrus.SetLevel(logrus.InfoLevel)
	if level, err := logrus.ParseLevel(os.Getenv("LOGLEVEL")); err == nil {
		logrus.SetLevel(level)
	}

	log.Infoln("[broker] Launching the Broker")
	InitResources()

	log.Infoln("[broker] Listening to batsim on", batEndpoint)
	batSock := NewReplySocket(batEndpoint)

	timeSock := NewRequestSocket("tcp://127.0.0.1:27000")

	defer func() {
		log.Infoln("[broker] Closing Batsim socket...")
		err = batSock.Close()
		if err != nil {
			log.Errorln("[broker] Error while closing Batsim socket: ", err)
		}
		log.Infoln("[broker] Batsim socket closed.") //TODO : is it?

		if err = timeSock.Close(); err != nil {
			log.Errorln("[broker] Error while closing time socket: ", err)
		}

		// WARNING: This is required to avoid stalling of the connection: the
		// last message is never sent if not called
		zmq.Term()
	}()

	// Loop responsible for time requests
	var now = make(chan float64, 1)
	var end = make(chan bool)
	var timeEvents = make(chan translate.Event)
	go handleTimeRequests(timeSock, end, now, timeEvents)

	// Main loop
	var batMsg translate.BatMessage
	var batMsgBytes []byte
	thisIsTheEnd := false
	for !thisIsTheEnd {
		batMsg = translate.BatMessage{}

		// Batsim message : receive it, decode it, handle it
		batMsgBytes, err = batSock.RecvBytes(0)
		if err != nil {
			log.Panicln("Error while receiving Batsim message: " + err.Error())
		}
		if err := json.Unmarshal(batMsgBytes, &batMsg); err != nil {
			log.Panicln(err)
		}
		batMsg.Now = round(batMsg.Now) // There is a rounding issue with some timestamps
		updateNow(now, batMsg)

		// Sending an mepty message as a response to an empty message
		// makes Batsim erorr out. We want to avoid this.
		lastMessageWasEmpty := false
		if len(batMsg.Events) == 0 {
			lastMessageWasEmpty = true
		}

		//log.Infoln("[broker] Batsim -> Broker:\n", string(batMsgBytes))
		for _, event := range batMsg.Events {
			if event.Type == "SIMULATION_ENDS" {
				thisIsTheEnd = true
				receivedSimulationEnded = true
			}
		}
		if noMoreJobs && unfinishedJobs == 0 {
			thisIsTheEnd = true
		}

		// Handle the message
		handleBatMessage(batMsg)
		batMsg.Events = make([]translate.Event, 0)

		// Get pending events to send to Batsim
		elapsedSinceLastMessage := time.Duration(0)
		lastMessageTime := time.Now()
		stopReceivingEvents := false
		loopStartTime := time.Now()
		for !stopReceivingEvents {
			updateNow(now, batMsg)
			select {
			case event := <-timeEvents: // Call me later events from time requests
				// Note : rounding and adding to now value each
				// time is not very precise. Errors due to
				// rounded values add up over the many
				// iterations of the loop. At the same time, it
				// filters noise and only takes into account
				// significant scheduler delays.
				batMsg.Now = addAndRound(batMsg.Now, elapsedSinceLastMessage)
				batMsg.Events = append(batMsg.Events, event)
				lastMessageTime = time.Now()
			case pod := <-ToExecute: // Jobs sent over by the api
				if pod.Status.Phase == "Running" {
					// This is an error. It means that the pod was binded twice.
					continue
				}
				batMsg.Now = addAndRound(batMsg.Now, elapsedSinceLastMessage)
				err, executeJob := translate.MakeEvent(batMsg.Now, "EXECUTE_JOB", translate.PodToExecuteJobData(pod))
				pod.Status.Phase = "Running"
				if err != nil {
					log.Panic("Failed to create event:", err)
				}
				batMsg.Events = append(batMsg.Events, executeJob)
				log.Infof("[broker:bathandler] pod %s was scheduled on node %s", pod.Metadata.Name, pod.Spec.NodeName)
				lastMessageTime = time.Now()
			default:
				// Wait at least minimalWaitDelay
				if time.Now().Sub(loopStartTime) < minimalWaitDelay {
					continue
				}
				// If Batsim has no pending requested calls and a response is expected,
				// do not send an empty message
				if len(batMsg.Events) == 0 && callMeLaters == 0 && !expectedEmptyResponse {
					continue
				}
				if sendMessageCondition&nonEmpty != 0 {
					if len(batMsg.Events) > 0 && !lastMessageWasEmpty {
						removeOutdatedEvents(&batMsg)
					}
					if len(batMsg.Events) > 0 {
						stopReceivingEvents = true
					}
				}
				if sendMessageCondition&timeout != 0 {
					elapsedSinceLastMessage = time.Now().Sub(lastMessageTime)
					if elapsedSinceLastMessage >= timeoutValue {
						removeOutdatedEvents(&batMsg)
						stopReceivingEvents = true
					}
				}
			}
		}

		if thisIsTheEnd {
			// This channel send may shut down the scheduler
			// prematurely. Maybe we can find another solution to
			// let the scheduler run for a while without
			// terminating zmq right now.
			end <- true
			log.Debug("It seems like the simulation ended.")
			emptyEvents(&batMsg)
			if !receivedSimulationEnded {
				emptyRequestedCallStack(&batMsg, batSock)
			}
		}

		countCallMeLaters(batMsg)

		batMsgBytes, err = json.Marshal(batMsg)
		if err != nil {
			log.Panicln("Error in message merging: " + err.Error())
		}
		_, err = batSock.SendBytes(batMsgBytes, 0)
		if err != nil {
			log.Panicln("Error while sending message to batsim: " + err.Error())
		}
		//log.Infoln("[broker] Broker -> Batsim:\n", string(batMsgBytes))
	}
	log.Infoln("[broker] Simulation finished successfully!")
}

// Counting the amount of unanswered call_me_laters can be done on the go,
// without a separate function, but it is done this was to simplify the code.
// This counter is decremented upon reception of REQUESTED_CALLs (in bathandler)
func countCallMeLaters(batMsg translate.BatMessage) {
	for _, event := range batMsg.Events {
		if event.Type == "CALL_ME_LATER" {
			callMeLaters++
		}
	}
}

// Sync with handleTimeRequests
func updateNow(now chan float64, batMsg translate.BatMessage) {
	// This is a non blocking send because now is buffered
	if len(now) == cap(now) {
		select {
		case <-now:
		default:
		}
	}
	now <- batMsg.Now
}

func isIn(i int64, s []int64) bool {
	for _, e := range s {
		if i == e {
			return true
		}
	}
	return false
}

/*
Remove any call_me_later that would be outdated

Returns if events have been removed
*/
func removeOutdatedEvents(batMsg *translate.BatMessage) {
	toRemove := make([]int, 0)
	for i, event := range batMsg.Events {
		if event.Type == "CALL_ME_LATER" && event.Data["timestamp"].(float64) <= batMsg.Now {
			toRemove = append(toRemove, i)
		}
	}

	// Do a reverse range to avoid index error
	last := len(toRemove) - 1
	for i := range toRemove {
		reverse_i := toRemove[last-i]
		batMsg.Events = append(
			batMsg.Events[:reverse_i],
			batMsg.Events[reverse_i+1:]...,
		)
	}
}

func addAndRound(now float64, d time.Duration) float64 {
	df := float64(d) / 1e9 // d is in nanoseconds
	return round(now + df)
}

func round(now float64) float64 {
	return math.Round(now*1000) / 1000 // we want to round to the closest millisecond
}

func emptyEvents(batMsg *translate.BatMessage) {
	for _, event := range batMsg.Events {
		if event.Type != "CALL_ME_LATER" {
			log.Panicf("Unexpected message type from the scheduler : %s. Did the simulation really end?", event.Type)
		}
	}
	batMsg.Events = make([]translate.Event, 0)
}

func emptyRequestedCallStack(batMsg *translate.BatMessage, batSock *zmq.Socket) {
	if len(batMsg.Events) != 0 {
		log.Panicf("emptyRequestedCallStack called with non-empty event list")
	}
	for !receivedSimulationEnded {
		batMsgBytes, err := json.Marshal(batMsg)
		if err != nil {
			log.Panicln("Error in message merging: " + err.Error())
		}
		_, err = batSock.SendBytes(batMsgBytes, 0)
		if err != nil {
			log.Panicln("Error while sending message to batsim: " + err.Error())
		}
		batMsgBytes, err = batSock.RecvBytes(0)
		if err != nil {
			log.Panicln("Error while receiving Batsim message: " + err.Error())
		}
		if err := json.Unmarshal(batMsgBytes, &batMsg); err != nil {
			log.Panicln(err)
		}
		for _, event := range batMsg.Events {
			if event.Type == "SIMULATION_ENDS" {
				receivedSimulationEnded = true
			} else if event.Type != "REQUESTED_CALL" {
				log.Panicf("Unexpected message type from Batsim : %s. Did the simulation really end?", event.Type)
			}
		}
		batMsg.Events = make([]translate.Event, 0)
	}
}
