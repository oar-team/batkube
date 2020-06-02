// Code generated by go-swagger; DO NOT EDIT.

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// ConnectCoreV1HeadNamespacedPodProxyOKCode is the HTTP code returned for type ConnectCoreV1HeadNamespacedPodProxyOK
const ConnectCoreV1HeadNamespacedPodProxyOKCode int = 200

/*ConnectCoreV1HeadNamespacedPodProxyOK OK

swagger:response connectCoreV1HeadNamespacedPodProxyOK
*/
type ConnectCoreV1HeadNamespacedPodProxyOK struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewConnectCoreV1HeadNamespacedPodProxyOK creates ConnectCoreV1HeadNamespacedPodProxyOK with default headers values
func NewConnectCoreV1HeadNamespacedPodProxyOK() *ConnectCoreV1HeadNamespacedPodProxyOK {

	return &ConnectCoreV1HeadNamespacedPodProxyOK{}
}

// WithPayload adds the payload to the connect core v1 head namespaced pod proxy o k response
func (o *ConnectCoreV1HeadNamespacedPodProxyOK) WithPayload(payload string) *ConnectCoreV1HeadNamespacedPodProxyOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the connect core v1 head namespaced pod proxy o k response
func (o *ConnectCoreV1HeadNamespacedPodProxyOK) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ConnectCoreV1HeadNamespacedPodProxyOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// ConnectCoreV1HeadNamespacedPodProxyUnauthorizedCode is the HTTP code returned for type ConnectCoreV1HeadNamespacedPodProxyUnauthorized
const ConnectCoreV1HeadNamespacedPodProxyUnauthorizedCode int = 401

/*ConnectCoreV1HeadNamespacedPodProxyUnauthorized Unauthorized

swagger:response connectCoreV1HeadNamespacedPodProxyUnauthorized
*/
type ConnectCoreV1HeadNamespacedPodProxyUnauthorized struct {
}

// NewConnectCoreV1HeadNamespacedPodProxyUnauthorized creates ConnectCoreV1HeadNamespacedPodProxyUnauthorized with default headers values
func NewConnectCoreV1HeadNamespacedPodProxyUnauthorized() *ConnectCoreV1HeadNamespacedPodProxyUnauthorized {

	return &ConnectCoreV1HeadNamespacedPodProxyUnauthorized{}
}

// WriteResponse to the client
func (o *ConnectCoreV1HeadNamespacedPodProxyUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
