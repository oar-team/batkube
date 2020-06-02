// Code generated by go-swagger; DO NOT EDIT.

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// ConnectCoreV1HeadNodeProxyOKCode is the HTTP code returned for type ConnectCoreV1HeadNodeProxyOK
const ConnectCoreV1HeadNodeProxyOKCode int = 200

/*ConnectCoreV1HeadNodeProxyOK OK

swagger:response connectCoreV1HeadNodeProxyOK
*/
type ConnectCoreV1HeadNodeProxyOK struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewConnectCoreV1HeadNodeProxyOK creates ConnectCoreV1HeadNodeProxyOK with default headers values
func NewConnectCoreV1HeadNodeProxyOK() *ConnectCoreV1HeadNodeProxyOK {

	return &ConnectCoreV1HeadNodeProxyOK{}
}

// WithPayload adds the payload to the connect core v1 head node proxy o k response
func (o *ConnectCoreV1HeadNodeProxyOK) WithPayload(payload string) *ConnectCoreV1HeadNodeProxyOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the connect core v1 head node proxy o k response
func (o *ConnectCoreV1HeadNodeProxyOK) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ConnectCoreV1HeadNodeProxyOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// ConnectCoreV1HeadNodeProxyUnauthorizedCode is the HTTP code returned for type ConnectCoreV1HeadNodeProxyUnauthorized
const ConnectCoreV1HeadNodeProxyUnauthorizedCode int = 401

/*ConnectCoreV1HeadNodeProxyUnauthorized Unauthorized

swagger:response connectCoreV1HeadNodeProxyUnauthorized
*/
type ConnectCoreV1HeadNodeProxyUnauthorized struct {
}

// NewConnectCoreV1HeadNodeProxyUnauthorized creates ConnectCoreV1HeadNodeProxyUnauthorized with default headers values
func NewConnectCoreV1HeadNodeProxyUnauthorized() *ConnectCoreV1HeadNodeProxyUnauthorized {

	return &ConnectCoreV1HeadNodeProxyUnauthorized{}
}

// WriteResponse to the client
func (o *ConnectCoreV1HeadNodeProxyUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
