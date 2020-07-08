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

// stop conditions for the loop responsible of transfering scheduler decisions.
const timeout = 1                // Real time timeout
const nonEmpty = 1 << 1          // A non-empty message is defined as a message containing at least one event that is not a CALL_ME_LATER
const simulationTimeout = 1 << 2 // Just like the real time timeout, but for simulation time.

// Stop condition when waiting for new messages to send to Batsim. Ex :
// stopWaitingForMessages = nonEmpty | timeout will stop waiting for messages if
// 	- the waiting time since the last message from the scheduler is higher
// 	than timeoutValue
// 	- OR if the Events slice contained in the response is not empty.
var sendMessageCondition = timeout | simulationTimeout

// If the duration since the last message received by the scheduler exceeds
// timeoutValue, Batkube will consider it as a timeout and will proceed with
// sending a message to Batsim
var timeoutValue = 20 * time.Millisecond

// Minimal amount of time to wait for messages from the scheduler.  Not having
// a minimal waiting time or having an insufficient minimal time leads to
// incorrect behavior from the scheduler, which does not have enough time to
// react.
// TODO : remove this variable, as it does not affect the simulation in any
// way. timeout is enough.
var minimalWaitDelay = 0 * time.Millisecond

// Enable CALL_ME__LATER events. These events originate from timer requests
// from the scheduler, allowing to fast forward in time to the next timestamp
// where a timer is supposed to fire.
var enableCallMeLaters bool = true

// Maximum time step to advance in the simulation. This prevents Batsim to jump
// too much forward in time, preventing the scheduler to properly take
// decisions.
var simulationTimestep = 500 * time.Millisecond

// Enable time incremental increases. When waiting for a response from the
// scheduler, increase the simulated time with a real time period by
// incrementValue with a real time period of incrementTimeStep.
// It is more efficient than waking up Batsim constantly with call_me_laters,
// but also makes the simulation slightly less accurate as events could happen
// during this time, that Batsim would have to wait for the next exchange to
// send.
var enableIncrementalTime bool = false
var incrementUpperLimit = 1 * time.Second    // Setting a fair upper limit moderates the effect described above.
var incrementTimeStep = 1 * time.Millisecond // Basically, slows down the simulation to give more time for the scheduler.
var incrementValue = 10 * time.Millisecond   // Having a low increment may (or may not) result in a more accurate simulation.

// Reset the time at wich we start counting to determine a timeout at each increment
// TODO : remove this
var resetLoopTimeOnIncrement bool = false

// Set to true when a no_more_static_job_to_submit NOTIFY is received.
var noMoreJobs bool

// Number of pods in running or pending status.
var unfinishedJobs int

// Number of jobs that are still running
var runningJobs int

// Allows us to keep track of pending requested calls
var requestedCalls []float64 = make([]float64, 0)

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
func handleTimeRequests(timeSock *zmq.Socket, end chan bool, now chan float64, timers chan float64) {
	if cap(now) != 1 {
		panic("now should be a buffered channel with capacity 1")
	}
	nowValue := <-now // Blocking receive : the value has to be initialized

	thisIsTheEnd := false
	for !thisIsTheEnd {
		// Get latest now value (if it is there).
		// Note : to make message passing between time, batkube and
		// batsim completely synchronous, make this a blocking receive.
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

		if enableCallMeLaters {
			processTimerRequests(nowValue, durations, timers)
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

/*
durations is an integer slice representing the durations of the timers
requested by the scheduler. For each of these timers, this function emits a
CALL_ME_LATER event to the events channel.
*/
func processTimerRequests(nowValue float64, durations []int64, timers chan float64) {
	for _, d := range durations {
		if d <= 0 {
			panic("Got a negative duration")
		}

		requested := addAndRound(nowValue, time.Duration(d))
		if requested < nowValue {
			panic("Requested a timer prior of current time")
		}

		// Non blocking send. Select isn't used to make sure it is actually sent.
		go func() {
			timers <- requested
		}()
	}
}

/*
Processes the messages coming from the scheduler and populates batMsg.Events
with them.

It handles all the options defined above regarding timing and other message
sending conditions.
*/
func processMessagesToSend(batMsg *translate.BatMessage, now chan float64, timers chan float64) {
	batMsg.Events = make([]translate.Event, 0)

	// Get pending events to send to Batsim
	var elapsedSinceLastMessage time.Duration
	var incremented time.Duration
	lastMessageTime := time.Now()
	loopStartTime := time.Now()
	lastIncrement := time.Now()
	var stopReceivingEvents bool // Go compiler does not like infinite loops
	var messageIsNotEmpty bool
	for !stopReceivingEvents {
		updateNow(now, *batMsg)
		select {
		case requested := <-timers:
			// Duplicates cause issues when keeping track of requested calls
			if !isAlreadyRequested(requested) {
				batMsg.Events = append(batMsg.Events, newCallMeLater(batMsg.Now, requested))
			}
		case pod := <-ToExecute: // Jobs sent over by the api
			messageIsNotEmpty = true
			// TODO : better study how to account for scheduling time
			batMsg.Now = addAndRound(batMsg.Now, elapsedSinceLastMessage)

			err, executeJob := translate.MakeEvent(batMsg.Now, "EXECUTE_JOB", translate.PodToExecuteJobData(pod))

			translate.UpdatePodStatusForScheduling(pod, translate.BatsimNowToMetaV1Time(batMsg.Now))
			IncrementResourceVersion(pod.Metadata)

			runningJobs++
			AddEvent(&translate.Modified, pod)
			if err != nil {
				log.Panic("Failed to create event:", err)
			}
			batMsg.Events = append(batMsg.Events, executeJob)
			log.Infof("[broker:bathandler] pod %s was scheduled on node %s", pod.Metadata.Name, pod.Spec.NodeName)
			lastMessageTime = time.Now()
		default:
			//fmt.Printf("[%f] len(events) %d; callMeLaters %d; unfinishedJobs %d; runningJobs %d\n", batMsg.Now, len(batMsg.Events), callMeLaters, unfinishedJobs, runningJobs)
			elapsedSinceLastMessage = time.Now().Sub(lastMessageTime)

			// Wait at least minimalWaitDelay
			if time.Now().Sub(loopStartTime) < minimalWaitDelay {
				continue
			}

			// increment simulation time if nothing happened
			if enableIncrementalTime &&
				incremented < incrementUpperLimit &&
				time.Now().Sub(lastIncrement) > incrementTimeStep {

				batMsg.Now = addAndRound(batMsg.Now, incrementValue)
				updateNow(now, *batMsg)
				lastIncrement = time.Now()
				incremented += incrementValue
				if resetLoopTimeOnIncrement {
					// Reset the loop start time value,
					// resulting in a minimal wait time
					// once again.
					loopStartTime = time.Now()
				}
			}

			if sendMessageCondition&nonEmpty != 0 && messageIsNotEmpty ||
				sendMessageCondition&timeout != 0 && elapsedSinceLastMessage > timeoutValue ||
				sendMessageCondition&simulationTimeout != 0 && incremented >= incrementUpperLimit {
				stopReceivingEvents = true
			}
		}
	}

	removeOutdatedEvents(batMsg)

	// We don't want to jump too much forward in time, so the scheduler can
	// make its decisions in time
	nextStep := addAndRound(batMsg.Now, simulationTimestep)
	if len(requestedCalls) == 0 || requestedCalls[0] > nextStep {
		batMsg.Events = append(batMsg.Events, newCallMeLater(batMsg.Now, nextStep))
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
	var timers = make(chan float64)
	go handleTimeRequests(timeSock, end, now, timers)

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

		//UpdateProbeAndHeartbeatTimes(batMsg.Now)
		//IncrementAllResourceVersions()

		// Handle the message
		handleBatMessage(batMsg)
		batMsg.Events = make([]translate.Event, 0)

		//log.Infoln("[broker] Batsim -> Broker:\n", string(batMsgBytes))
		if receivedSimulationEnded || noMoreJobs && unfinishedJobs == 0 {
			thisIsTheEnd = true
		}

		if thisIsTheEnd {
			// This channel send may shut down the scheduler
			// prematurely. Maybe we can find another solution to
			// let the scheduler run for a while without
			// terminating zmq right now.
			log.Info("It seems like the simulation ended.")
			go func() {
				end <- true
			}()

			if !receivedSimulationEnded {
				emptyRequestedCallStack(&batMsg, batSock)
			}
		} else {
			processMessagesToSend(&batMsg, now, timers)
		}

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

/*
Remove any call_me_later that would be outdated

Returns if events have been removed
*/
func removeOutdatedEvents(batMsg *translate.BatMessage) {
	toRemove := make([]int, 0)
	for i, event := range batMsg.Events {
		if event.Type == "CALL_ME_LATER" && event.Data["timestamp"].(float64) <= batMsg.Now {
			toRemove = append(toRemove, i)
			deleteRequestedCall(event.Data["timestamp"].(float64))
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

/*
Warning : not thread safe.
*/
func newCallMeLater(now, timestamp float64) translate.Event {
	err, callMeLater := translate.MakeEvent(now, "CALL_ME_LATER", translate.CallMeLaterData{Timestamp: timestamp})
	if err != nil {
		log.Panic("Failed to create event:", err)
	}

	// Add a new requested call to the stack
	added := false
	for i, call := range requestedCalls {
		if call > timestamp {
			requestedCalls = append(requestedCalls, 0)
			copy(requestedCalls[i+1:], requestedCalls[i:])
			requestedCalls[i] = timestamp
			added = true
			break
		}
	}
	// If nothing was done, it means that either requestedCalls is empty or
	// timestamp exceeds every other value in the slice
	if !added {
		requestedCalls = append(requestedCalls, timestamp)
	}

	return callMeLater
}

/*
Warning : not thread safe.
*/
func deleteRequestedCall(timestamp float64) {
	for i, requested := range requestedCalls {
		if requested == timestamp {
			n := len(requestedCalls)
			requestedCalls[n-1], requestedCalls[i] = requestedCalls[i], requestedCalls[n-1]
			requestedCalls = requestedCalls[:n-1]
			break
		}
	}
}

func isAlreadyRequested(requested float64) bool {
	for _, timestamp := range requestedCalls {
		if requested == timestamp {
			return true
		}
	}
	return false
}
