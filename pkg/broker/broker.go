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

type broker struct {
	// If the duration since the last message received by the scheduler exceeds
	// timeoutValue, Batkube will consider it as a timeout and will proceed with
	// sending a message to Batsim
	timeoutValue time.Duration

	// Deprecated. Might be uselful in some cases, might not.
	//
	// Enable CALL_ME_LATER events triggered by scheduler timers. That way, the
	// scheduler gets control back whenever a timer is supposed to  fire.  This
	// results in scheduler code executing at the exact time it is supposed to be
	// executed, at the expense of very slow simulations.
	enableCallMeLaters bool

	// Each cycle, a call me later event is emitted in order to advance in the
	// simulation, fast forwarding into the future with a certain timestep value.
	// This value is a starting point, that will increase or decrease depending on
	// a backoff policy and other logic based on the context. See getNextWakeUp for
	// implementation details
	baseSimulationTimestep time.Duration

	// Maximum amount of time Batsim is allowed to jump forward in time
	maxSimulationTimestep time.Duration
	backoffMultiplier     float64

	// Schedulers may take a while to startup. This factors slows down the
	// simulation until Batkube gets it first decision from the scheduler.
	startupSlowDownFactor float64

	currentSimulationTimestep time.Duration

	// Set this to true if the scheduler is not excepted to make any decisions if
	// no jobs are in a pending state. Set it to false otherwise (for exemple, if
	// the scheduler has policies like preemption).
	fastForwardOnNoPendingJobs bool

	// Set to true when a no_more_static_job_to_submit NOTIFY is received.
	noMoreJobs bool

	// Number of pods in running or pending status.
	unfinishedJobs int

	// Number of jobs that are still running
	runningJobs int

	// Allows us to keep track of pending requested calls
	requestedCalls []float64

	// Batkube received a SIMULATION_ENDS event.
	receivedSimulationEnded bool

	// The scheduler is a bit slow at startup, so the simulation shouldn't be too
	// fast before it reacts for the first time
	firstJobWasScheduled bool
}

type BatkubeOptions struct {
	TimeoutValue               int64   `long:"timeout-value" description:"maximum amount of time spent waiting for messages from the scheduler, in milliseconds" default:"20"`
	BaseSimulationTimestep     int64   `long:"base-simulation-timestep" description:"maximum amount of time Batsim is allowed to jump forward in time, in milliseconds. This value increases according to a backoff policy, up to a maximum amount" default:"100"`
	MaxSimulationTimestep      int64   `long:"max-simulation-timestep" description:"maximum value authorized for simulationTimestep, in seconds" default:"50"`
	BackoffMultiplier          float64 `long:"backoff-multiplier" description:"each time the scheduler did not react, simulationTimestep is multiplied by this amount" default:"2"`
	FastForwardOnNoPendingJobs bool    `long:"fast-forward-on-no-pending-jobs" description:"if there are no pending jobs the simulation may fast forwards to the next Batsim event, potentially skipping some scheduler decisions"`
}

func NewBroker(options *BatkubeOptions) *broker {
	return &broker{
		startupSlowDownFactor:      5,
		requestedCalls:             make([]float64, 0),
		timeoutValue:               time.Duration(options.TimeoutValue) * time.Millisecond,
		baseSimulationTimestep:     time.Duration(options.BaseSimulationTimestep) * time.Millisecond,
		maxSimulationTimestep:      time.Duration(options.MaxSimulationTimestep) * time.Second,
		backoffMultiplier:          options.BackoffMultiplier,
		fastForwardOnNoPendingJobs: options.FastForwardOnNoPendingJobs,
	}
}

func (b *broker) Run(batEndpoint string) {
	var err error

	//TODO: set currentSimulationTimestep

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
		log.Infoln("[broker] Batsim socket closed.") // TODO : is it?

		if err = timeSock.Close(); err != nil {
			log.Errorln("[broker] Error while closing time socket: ", err)
		}

		// WARNING: This is required to avoid stalling of the connection: the
		// last message is never sent if not called
		zmq.Term()
		os.Exit(1)
	}()

	// Loop responsible for time requests
	var now = make(chan float64, 1)
	var end = make(chan bool)
	var timers = make(chan float64)
	go handleTimeRequests(b, timeSock, end, now, timers)

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
		handleBatMessage(b, batMsg)
		batMsg.Events = make([]translate.Event, 0)

		//log.Infoln("[broker] Batsim -> Broker:\n", string(batMsgBytes))
		if b.receivedSimulationEnded || b.noMoreJobs && b.unfinishedJobs == 0 {
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

			if !b.receivedSimulationEnded {
				emptyRequestedCallStack(b, &batMsg, batSock)
			}
		} else {
			processMessagesToSend(b, &batMsg, now, timers)
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

/*
Handles time requests asynchronously. All sends and receives are non blocking

end : send a boolean on this channel to end the loop.
now : whenever now is updated, send it to this channel.
events : CALL_ME_LATER events are sent to this channel. They are stacked and sent whenever the receiver is ready.
*/
func handleTimeRequests(b *broker, timeSock *zmq.Socket, end chan bool, now chan float64, timers chan float64) {
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

		if b.enableCallMeLaters {
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

		// Non blocking send. Select isn't used to make sure it is
		// actually sent.  The send has to be non blocking so the
		// scheduler gets the time as quickly as possible.
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
func processMessagesToSend(b *broker, batMsg *translate.BatMessage, now chan float64, timers chan float64) {
	batMsg.Events = make([]translate.Event, 0)

	// Get pending events to send to Batsim
	loopStartTime := time.Now()
	loopSimulatedStartTime := batMsg.Now
	var elapsedSinceLoopStart time.Duration
	var stopReceivingEvents bool // Go compiler does not like infinite loops
	var messageIsNotEmpty bool
	for !stopReceivingEvents {
		select {
		case requested := <-timers:
			// Duplicates cause issues when keeping track of
			// requested calls A check for duplicates is already
			// made when receiving timer requests but it does not
			// account for previous requests.
			if !isAlreadyRequested(b, requested) {
				batMsg.Events = append(batMsg.Events, newCallMeLater(b, batMsg.Now, requested))
			}
		case pod := <-ToExecute: // Jobs sent over by the api
			messageIsNotEmpty = true
			b.firstJobWasScheduled = true
			b.currentSimulationTimestep = b.baseSimulationTimestep
			b.runningJobs++

			// Resource housekeeping
			err, executeJob := translate.MakeEvent(batMsg.Now, "EXECUTE_JOB", translate.PodToExecuteJobData(pod))
			translate.UpdatePodStatusForScheduling(pod, translate.BatsimNowToMetaV1Time(batMsg.Now))
			IncrementResourceVersion(pod.Metadata)
			AddEvent(&translate.Modified, pod)
			if err != nil {
				log.Panic("Failed to create event:", err)
			}
			batMsg.Events = append(batMsg.Events, executeJob)
			log.Infof("[broker:bathandler] pod %s was scheduled on node %s", pod.Metadata.Name, pod.Spec.NodeName)
		default:
		}

		elapsedSinceLoopStart = time.Now().Sub(loopStartTime)
		batMsg.Now = addAndRound(loopSimulatedStartTime, elapsedSinceLoopStart)
		updateNow(now, *batMsg)

		if messageIsNotEmpty || elapsedSinceLoopStart > b.timeoutValue {
			stopReceivingEvents = true
		}
	}

	// If CALL_ME_LATER on scheduler timers are enables, some may point to
	// timestamps in the past now
	removeOutdatedCML(b, batMsg)

	// We don't want to jump too much forward in time, so the scheduler can
	// make its decisions in time. This "not too much forward in time" is
	// given by getNextWakeUp and depends on the context given to the
	// simulator at startup.
	nextStep := getNextWakeUp(b, batMsg.Now)
	if nextStep > batMsg.Now && (len(b.requestedCalls) == 0 || b.requestedCalls[0] > nextStep) {
		batMsg.Events = append(batMsg.Events, newCallMeLater(b, batMsg.Now, nextStep))
	}
}

// Sync with handleTimeRequests
func updateNow(now chan float64, batMsg translate.BatMessage) {
	// Empty any outdated value (there is one at most)
	select {
	case <-now:
	default:
	}
	// This is a non blocking send because now is buffered
	now <- batMsg.Now
}

/*
Remove any call_me_later that would be outdated

Returns if events have been removed
*/
func removeOutdatedCML(b *broker, batMsg *translate.BatMessage) {
	toRemove := make([]int, 0)
	for i, event := range batMsg.Events {
		if event.Type == "CALL_ME_LATER" && event.Data["timestamp"].(float64) <= batMsg.Now {
			toRemove = append(toRemove, i)
			deleteRequestedCall(b, event.Data["timestamp"].(float64))
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

func emptyRequestedCallStack(b *broker, batMsg *translate.BatMessage, batSock *zmq.Socket) {
	if len(batMsg.Events) != 0 {
		log.Panicf("emptyRequestedCallStack called with non-empty event list")
	}
	for !b.receivedSimulationEnded {
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
				b.receivedSimulationEnded = true
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
func newCallMeLater(b *broker, now, timestamp float64) translate.Event {
	err, callMeLater := translate.MakeEvent(now, "CALL_ME_LATER", translate.CallMeLaterData{Timestamp: timestamp})
	if err != nil {
		log.Panic("Failed to create event:", err)
	}

	// Add a new requested call to the stack
	added := false
	for i, call := range b.requestedCalls {
		if call > timestamp {
			b.requestedCalls = append(b.requestedCalls, 0)
			copy(b.requestedCalls[i+1:], b.requestedCalls[i:])
			b.requestedCalls[i] = timestamp
			added = true
			break
		}
	}
	// If nothing was done, it means that either requestedCalls is empty or
	// timestamp exceeds every other value in the slice
	if !added {
		b.requestedCalls = append(b.requestedCalls, timestamp)
	}

	return callMeLater
}

/*
Warning : not thread safe.
*/
func deleteRequestedCall(b *broker, timestamp float64) {
	for i, requested := range b.requestedCalls {
		if requested == timestamp {
			n := len(b.requestedCalls)
			b.requestedCalls[n-1], b.requestedCalls[i] = b.requestedCalls[i], b.requestedCalls[n-1]
			b.requestedCalls = b.requestedCalls[:n-1]
			break
		}
	}
}

func isAlreadyRequested(b *broker, requested float64) bool {
	for _, timestamp := range b.requestedCalls {
		if requested == timestamp {
			return true
		}
	}
	return false
}

/*
Returns the next step Batsim is allowed to fast forward to
*/
func getNextWakeUp(b *broker, now float64) float64 {
	var timestep time.Duration

	if !b.firstJobWasScheduled {
		// It does make sense to to this : Durations are nothing more than int64.
		timestep = b.baseSimulationTimestep / time.Duration(b.startupSlowDownFactor)
	} else if b.fastForwardOnNoPendingJobs && b.unfinishedJobs-b.runningJobs == 0 {
		timestep = 0
	} else if b.currentSimulationTimestep < b.maxSimulationTimestep {
		b.currentSimulationTimestep = time.Duration(float64(b.currentSimulationTimestep) * b.backoffMultiplier)
		if b.currentSimulationTimestep > b.maxSimulationTimestep {
			b.currentSimulationTimestep = b.maxSimulationTimestep
		}
		timestep = b.currentSimulationTimestep
	} else {
		timestep = b.maxSimulationTimestep
	}
	return addAndRound(now, timestep)
}
