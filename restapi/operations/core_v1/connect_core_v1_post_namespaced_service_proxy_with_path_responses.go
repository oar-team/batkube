// Code generated by go-swagger; DO NOT EDIT.

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// ConnectCoreV1PostNamespacedServiceProxyWithPathOKCode is the HTTP code returned for type ConnectCoreV1PostNamespacedServiceProxyWithPathOK
const ConnectCoreV1PostNamespacedServiceProxyWithPathOKCode int = 200

/*ConnectCoreV1PostNamespacedServiceProxyWithPathOK OK

swagger:response connectCoreV1PostNamespacedServiceProxyWithPathOK
*/
type ConnectCoreV1PostNamespacedServiceProxyWithPathOK struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewConnectCoreV1PostNamespacedServiceProxyWithPathOK creates ConnectCoreV1PostNamespacedServiceProxyWithPathOK with default headers values
func NewConnectCoreV1PostNamespacedServiceProxyWithPathOK() *ConnectCoreV1PostNamespacedServiceProxyWithPathOK {

	return &ConnectCoreV1PostNamespacedServiceProxyWithPathOK{}
}

// WithPayload adds the payload to the connect core v1 post namespaced service proxy with path o k response
func (o *ConnectCoreV1PostNamespacedServiceProxyWithPathOK) WithPayload(payload string) *ConnectCoreV1PostNamespacedServiceProxyWithPathOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the connect core v1 post namespaced service proxy with path o k response
func (o *ConnectCoreV1PostNamespacedServiceProxyWithPathOK) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ConnectCoreV1PostNamespacedServiceProxyWithPathOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// ConnectCoreV1PostNamespacedServiceProxyWithPathUnauthorizedCode is the HTTP code returned for type ConnectCoreV1PostNamespacedServiceProxyWithPathUnauthorized
const ConnectCoreV1PostNamespacedServiceProxyWithPathUnauthorizedCode int = 401

/*ConnectCoreV1PostNamespacedServiceProxyWithPathUnauthorized Unauthorized

swagger:response connectCoreV1PostNamespacedServiceProxyWithPathUnauthorized
*/
type ConnectCoreV1PostNamespacedServiceProxyWithPathUnauthorized struct {
}

// NewConnectCoreV1PostNamespacedServiceProxyWithPathUnauthorized creates ConnectCoreV1PostNamespacedServiceProxyWithPathUnauthorized with default headers values
func NewConnectCoreV1PostNamespacedServiceProxyWithPathUnauthorized() *ConnectCoreV1PostNamespacedServiceProxyWithPathUnauthorized {

	return &ConnectCoreV1PostNamespacedServiceProxyWithPathUnauthorized{}
}

// WriteResponse to the client
func (o *ConnectCoreV1PostNamespacedServiceProxyWithPathUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
