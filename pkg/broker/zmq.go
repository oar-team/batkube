package broker

import (
	zmq "github.com/pebbe/zmq4"

	"fmt"
)

func newSocket(endpoint string, zmq_socket_type zmq.Type) *zmq.Socket {
	// Create Request socket
	fmt.Println("Connecting to " + endpoint)
	socket, err := zmq.NewSocket(zmq_socket_type)
	if err != nil {
		panic("Unable to create the socket" + endpoint + " with error: " + err.Error())
	}
	return socket
}

// Public API

func NewRequestSocket(endpoint string) *zmq.Socket {
	requester := newSocket(endpoint, zmq.REQ)
	// Connect to Reply socket
	err := requester.Connect(endpoint)
	if err != nil {
		panic("Error while connecting to socket: " + err.Error())
	}
	time, _ := requester.GetLinger()
	fmt.Println("ZMQ_LINGER for socket " + endpoint + " is " + time.String())
	return requester
}

func NewReplySocket(endpoint string) *zmq.Socket {
	responder := newSocket(endpoint, zmq.REP)
	err := responder.Bind(endpoint)
	if err != nil {
		panic("Error while binding socket to \"" + endpoint + "\" with error: " + err.Error())
	}
	time, _ := responder.GetLinger()
	fmt.Println("ZMQ_LINGER for socket " + endpoint + " is " + time.String())
	return responder
}
