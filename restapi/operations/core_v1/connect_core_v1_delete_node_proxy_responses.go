// Code generated by go-swagger; DO NOT EDIT.

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// ConnectCoreV1DeleteNodeProxyOKCode is the HTTP code returned for type ConnectCoreV1DeleteNodeProxyOK
const ConnectCoreV1DeleteNodeProxyOKCode int = 200

/*ConnectCoreV1DeleteNodeProxyOK OK

swagger:response connectCoreV1DeleteNodeProxyOK
*/
type ConnectCoreV1DeleteNodeProxyOK struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewConnectCoreV1DeleteNodeProxyOK creates ConnectCoreV1DeleteNodeProxyOK with default headers values
func NewConnectCoreV1DeleteNodeProxyOK() *ConnectCoreV1DeleteNodeProxyOK {

	return &ConnectCoreV1DeleteNodeProxyOK{}
}

// WithPayload adds the payload to the connect core v1 delete node proxy o k response
func (o *ConnectCoreV1DeleteNodeProxyOK) WithPayload(payload string) *ConnectCoreV1DeleteNodeProxyOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the connect core v1 delete node proxy o k response
func (o *ConnectCoreV1DeleteNodeProxyOK) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ConnectCoreV1DeleteNodeProxyOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// ConnectCoreV1DeleteNodeProxyUnauthorizedCode is the HTTP code returned for type ConnectCoreV1DeleteNodeProxyUnauthorized
const ConnectCoreV1DeleteNodeProxyUnauthorizedCode int = 401

/*ConnectCoreV1DeleteNodeProxyUnauthorized Unauthorized

swagger:response connectCoreV1DeleteNodeProxyUnauthorized
*/
type ConnectCoreV1DeleteNodeProxyUnauthorized struct {
}

// NewConnectCoreV1DeleteNodeProxyUnauthorized creates ConnectCoreV1DeleteNodeProxyUnauthorized with default headers values
func NewConnectCoreV1DeleteNodeProxyUnauthorized() *ConnectCoreV1DeleteNodeProxyUnauthorized {

	return &ConnectCoreV1DeleteNodeProxyUnauthorized{}
}

// WriteResponse to the client
func (o *ConnectCoreV1DeleteNodeProxyUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}