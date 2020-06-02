// Code generated by go-swagger; DO NOT EDIT.

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// ConnectCoreV1PutNamespacedServiceProxyWithPathOKCode is the HTTP code returned for type ConnectCoreV1PutNamespacedServiceProxyWithPathOK
const ConnectCoreV1PutNamespacedServiceProxyWithPathOKCode int = 200

/*ConnectCoreV1PutNamespacedServiceProxyWithPathOK OK

swagger:response connectCoreV1PutNamespacedServiceProxyWithPathOK
*/
type ConnectCoreV1PutNamespacedServiceProxyWithPathOK struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewConnectCoreV1PutNamespacedServiceProxyWithPathOK creates ConnectCoreV1PutNamespacedServiceProxyWithPathOK with default headers values
func NewConnectCoreV1PutNamespacedServiceProxyWithPathOK() *ConnectCoreV1PutNamespacedServiceProxyWithPathOK {

	return &ConnectCoreV1PutNamespacedServiceProxyWithPathOK{}
}

// WithPayload adds the payload to the connect core v1 put namespaced service proxy with path o k response
func (o *ConnectCoreV1PutNamespacedServiceProxyWithPathOK) WithPayload(payload string) *ConnectCoreV1PutNamespacedServiceProxyWithPathOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the connect core v1 put namespaced service proxy with path o k response
func (o *ConnectCoreV1PutNamespacedServiceProxyWithPathOK) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ConnectCoreV1PutNamespacedServiceProxyWithPathOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// ConnectCoreV1PutNamespacedServiceProxyWithPathUnauthorizedCode is the HTTP code returned for type ConnectCoreV1PutNamespacedServiceProxyWithPathUnauthorized
const ConnectCoreV1PutNamespacedServiceProxyWithPathUnauthorizedCode int = 401

/*ConnectCoreV1PutNamespacedServiceProxyWithPathUnauthorized Unauthorized

swagger:response connectCoreV1PutNamespacedServiceProxyWithPathUnauthorized
*/
type ConnectCoreV1PutNamespacedServiceProxyWithPathUnauthorized struct {
}

// NewConnectCoreV1PutNamespacedServiceProxyWithPathUnauthorized creates ConnectCoreV1PutNamespacedServiceProxyWithPathUnauthorized with default headers values
func NewConnectCoreV1PutNamespacedServiceProxyWithPathUnauthorized() *ConnectCoreV1PutNamespacedServiceProxyWithPathUnauthorized {

	return &ConnectCoreV1PutNamespacedServiceProxyWithPathUnauthorized{}
}

// WriteResponse to the client
func (o *ConnectCoreV1PutNamespacedServiceProxyWithPathUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
