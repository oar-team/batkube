package broker

import (
	"encoding/binary"
	"encoding/json"
	"time"

	zmq "github.com/pebbe/zmq4"
	log "github.com/sirupsen/logrus"
	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/pkg/translate"
)

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
	nowValue := <-now // has to be initialized

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

			// conversion from nanoseconds to seconds.milliseconds
			df := float64(d) / 1e9
			requested := nowValue + df
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

		// Answer the messages
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
	log.Infoln("[broker] Launching the Broker")

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

		log.Infoln("[broker] Batsim -> Broker:\n", string(batMsgBytes))
		for _, event := range batMsg.Events {
			if event.Type == "SIMULATION_ENDS" {
				thisIsTheEnd = true
				end <- true
			}
		}

		// Handle the message, take back the response
		batMsg = handleBatMessage(batMsg)

		// Get pending events to send to Batsim
		elapsedSinceLastMessage := time.Duration(0)
		lastMessageTime := time.Now()
		timeout := 500 * time.Millisecond
		stopReceivingEvents := false
		stopCondition := "non-empty"
		for !stopReceivingEvents {
			updateNow(now, batMsg.Now)
			select {
			case event := <-timeEvents:
				// Call me laters from time requests
				batMsg.Now += float64(elapsedSinceLastMessage) / 1e9
				lastMessageTime = time.Now()
				batMsg.Events = append(batMsg.Events, event)
			case pod := <-ToExecute:
				// Jobs sent over by the api
				batMsg.Now += float64(elapsedSinceLastMessage) / 1e9
				lastMessageTime = time.Now()
				err, executeJob := translate.MakeEvent(batMsg.Now, "EXECUTE_JOB", translate.PodToExecuteJobData(pod))
				if err != nil {
					log.Panic("Failed to create event:", err)
				}
				batMsg.Events = append(batMsg.Events, executeJob)
				log.Infof("[broker:bathandler] pod %s was scheduled on node %s", pod.Metadata.Name, pod.Spec.NodeName)
			default:
				// TODO : use masks to parameterize stop conditions.
				switch stopCondition {
				case "non-empty":
					if len(batMsg.Events) > 0 {
						stopReceivingEvents = true
					}
				case "timeout":
					elapsedSinceLastMessage = time.Now().Sub(lastMessageTime)
					if elapsedSinceLastMessage >= timeout {
						stopReceivingEvents = true
					}
				}
			}
		}

		batMsgBytes, err = json.Marshal(batMsg)
		if err != nil {
			log.Panicln("Error in message merging: " + err.Error())
		}
		_, err = batSock.SendBytes(batMsgBytes, 0)
		if err != nil {
			log.Panicln("Error while sending message to batsim: " + err.Error())
		}
		log.Infoln("[broker] Broker -> Batsim:\n", string(batMsgBytes))
	}
	log.Infoln("[broker] Simulation finished successfully!")
}

// Sync with handleTimeRequests
func updateNow(now chan float64, nowValue float64) {
	// This is a non blocking send because now is buffered
	if len(now) == cap(now) {
		select {
		case <-now:
		default:
		}
	}
	now <- nowValue
}

func isIn(i int64, s []int64) bool {
	for _, e := range s {
		if i == e {
			return true
		}
	}
	return false
}