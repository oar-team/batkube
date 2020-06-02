// Code generated by go-swagger; DO NOT EDIT.

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// ConnectCoreV1GetNamespacedPodExecOKCode is the HTTP code returned for type ConnectCoreV1GetNamespacedPodExecOK
const ConnectCoreV1GetNamespacedPodExecOKCode int = 200

/*ConnectCoreV1GetNamespacedPodExecOK OK

swagger:response connectCoreV1GetNamespacedPodExecOK
*/
type ConnectCoreV1GetNamespacedPodExecOK struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewConnectCoreV1GetNamespacedPodExecOK creates ConnectCoreV1GetNamespacedPodExecOK with default headers values
func NewConnectCoreV1GetNamespacedPodExecOK() *ConnectCoreV1GetNamespacedPodExecOK {

	return &ConnectCoreV1GetNamespacedPodExecOK{}
}

// WithPayload adds the payload to the connect core v1 get namespaced pod exec o k response
func (o *ConnectCoreV1GetNamespacedPodExecOK) WithPayload(payload string) *ConnectCoreV1GetNamespacedPodExecOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the connect core v1 get namespaced pod exec o k response
func (o *ConnectCoreV1GetNamespacedPodExecOK) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ConnectCoreV1GetNamespacedPodExecOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// ConnectCoreV1GetNamespacedPodExecUnauthorizedCode is the HTTP code returned for type ConnectCoreV1GetNamespacedPodExecUnauthorized
const ConnectCoreV1GetNamespacedPodExecUnauthorizedCode int = 401

/*ConnectCoreV1GetNamespacedPodExecUnauthorized Unauthorized

swagger:response connectCoreV1GetNamespacedPodExecUnauthorized
*/
type ConnectCoreV1GetNamespacedPodExecUnauthorized struct {
}

// NewConnectCoreV1GetNamespacedPodExecUnauthorized creates ConnectCoreV1GetNamespacedPodExecUnauthorized with default headers values
func NewConnectCoreV1GetNamespacedPodExecUnauthorized() *ConnectCoreV1GetNamespacedPodExecUnauthorized {

	return &ConnectCoreV1GetNamespacedPodExecUnauthorized{}
}

// WriteResponse to the client
func (o *ConnectCoreV1GetNamespacedPodExecUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
