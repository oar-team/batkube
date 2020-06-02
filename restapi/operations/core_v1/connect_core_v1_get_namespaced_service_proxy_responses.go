// Code generated by go-swagger; DO NOT EDIT.

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// ConnectCoreV1GetNamespacedServiceProxyOKCode is the HTTP code returned for type ConnectCoreV1GetNamespacedServiceProxyOK
const ConnectCoreV1GetNamespacedServiceProxyOKCode int = 200

/*ConnectCoreV1GetNamespacedServiceProxyOK OK

swagger:response connectCoreV1GetNamespacedServiceProxyOK
*/
type ConnectCoreV1GetNamespacedServiceProxyOK struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewConnectCoreV1GetNamespacedServiceProxyOK creates ConnectCoreV1GetNamespacedServiceProxyOK with default headers values
func NewConnectCoreV1GetNamespacedServiceProxyOK() *ConnectCoreV1GetNamespacedServiceProxyOK {

	return &ConnectCoreV1GetNamespacedServiceProxyOK{}
}

// WithPayload adds the payload to the connect core v1 get namespaced service proxy o k response
func (o *ConnectCoreV1GetNamespacedServiceProxyOK) WithPayload(payload string) *ConnectCoreV1GetNamespacedServiceProxyOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the connect core v1 get namespaced service proxy o k response
func (o *ConnectCoreV1GetNamespacedServiceProxyOK) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ConnectCoreV1GetNamespacedServiceProxyOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// ConnectCoreV1GetNamespacedServiceProxyUnauthorizedCode is the HTTP code returned for type ConnectCoreV1GetNamespacedServiceProxyUnauthorized
const ConnectCoreV1GetNamespacedServiceProxyUnauthorizedCode int = 401

/*ConnectCoreV1GetNamespacedServiceProxyUnauthorized Unauthorized

swagger:response connectCoreV1GetNamespacedServiceProxyUnauthorized
*/
type ConnectCoreV1GetNamespacedServiceProxyUnauthorized struct {
}

// NewConnectCoreV1GetNamespacedServiceProxyUnauthorized creates ConnectCoreV1GetNamespacedServiceProxyUnauthorized with default headers values
func NewConnectCoreV1GetNamespacedServiceProxyUnauthorized() *ConnectCoreV1GetNamespacedServiceProxyUnauthorized {

	return &ConnectCoreV1GetNamespacedServiceProxyUnauthorized{}
}

// WriteResponse to the client
func (o *ConnectCoreV1GetNamespacedServiceProxyUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
