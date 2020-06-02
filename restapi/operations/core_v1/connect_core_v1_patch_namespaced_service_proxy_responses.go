// Code generated by go-swagger; DO NOT EDIT.

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// ConnectCoreV1PatchNamespacedServiceProxyOKCode is the HTTP code returned for type ConnectCoreV1PatchNamespacedServiceProxyOK
const ConnectCoreV1PatchNamespacedServiceProxyOKCode int = 200

/*ConnectCoreV1PatchNamespacedServiceProxyOK OK

swagger:response connectCoreV1PatchNamespacedServiceProxyOK
*/
type ConnectCoreV1PatchNamespacedServiceProxyOK struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewConnectCoreV1PatchNamespacedServiceProxyOK creates ConnectCoreV1PatchNamespacedServiceProxyOK with default headers values
func NewConnectCoreV1PatchNamespacedServiceProxyOK() *ConnectCoreV1PatchNamespacedServiceProxyOK {

	return &ConnectCoreV1PatchNamespacedServiceProxyOK{}
}

// WithPayload adds the payload to the connect core v1 patch namespaced service proxy o k response
func (o *ConnectCoreV1PatchNamespacedServiceProxyOK) WithPayload(payload string) *ConnectCoreV1PatchNamespacedServiceProxyOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the connect core v1 patch namespaced service proxy o k response
func (o *ConnectCoreV1PatchNamespacedServiceProxyOK) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ConnectCoreV1PatchNamespacedServiceProxyOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// ConnectCoreV1PatchNamespacedServiceProxyUnauthorizedCode is the HTTP code returned for type ConnectCoreV1PatchNamespacedServiceProxyUnauthorized
const ConnectCoreV1PatchNamespacedServiceProxyUnauthorizedCode int = 401

/*ConnectCoreV1PatchNamespacedServiceProxyUnauthorized Unauthorized

swagger:response connectCoreV1PatchNamespacedServiceProxyUnauthorized
*/
type ConnectCoreV1PatchNamespacedServiceProxyUnauthorized struct {
}

// NewConnectCoreV1PatchNamespacedServiceProxyUnauthorized creates ConnectCoreV1PatchNamespacedServiceProxyUnauthorized with default headers values
func NewConnectCoreV1PatchNamespacedServiceProxyUnauthorized() *ConnectCoreV1PatchNamespacedServiceProxyUnauthorized {

	return &ConnectCoreV1PatchNamespacedServiceProxyUnauthorized{}
}

// WriteResponse to the client
func (o *ConnectCoreV1PatchNamespacedServiceProxyUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
