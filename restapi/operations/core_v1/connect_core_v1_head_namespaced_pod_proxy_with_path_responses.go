// Code generated by go-swagger; DO NOT EDIT.

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// ConnectCoreV1HeadNamespacedPodProxyWithPathOKCode is the HTTP code returned for type ConnectCoreV1HeadNamespacedPodProxyWithPathOK
const ConnectCoreV1HeadNamespacedPodProxyWithPathOKCode int = 200

/*ConnectCoreV1HeadNamespacedPodProxyWithPathOK OK

swagger:response connectCoreV1HeadNamespacedPodProxyWithPathOK
*/
type ConnectCoreV1HeadNamespacedPodProxyWithPathOK struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewConnectCoreV1HeadNamespacedPodProxyWithPathOK creates ConnectCoreV1HeadNamespacedPodProxyWithPathOK with default headers values
func NewConnectCoreV1HeadNamespacedPodProxyWithPathOK() *ConnectCoreV1HeadNamespacedPodProxyWithPathOK {

	return &ConnectCoreV1HeadNamespacedPodProxyWithPathOK{}
}

// WithPayload adds the payload to the connect core v1 head namespaced pod proxy with path o k response
func (o *ConnectCoreV1HeadNamespacedPodProxyWithPathOK) WithPayload(payload string) *ConnectCoreV1HeadNamespacedPodProxyWithPathOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the connect core v1 head namespaced pod proxy with path o k response
func (o *ConnectCoreV1HeadNamespacedPodProxyWithPathOK) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ConnectCoreV1HeadNamespacedPodProxyWithPathOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// ConnectCoreV1HeadNamespacedPodProxyWithPathUnauthorizedCode is the HTTP code returned for type ConnectCoreV1HeadNamespacedPodProxyWithPathUnauthorized
const ConnectCoreV1HeadNamespacedPodProxyWithPathUnauthorizedCode int = 401

/*ConnectCoreV1HeadNamespacedPodProxyWithPathUnauthorized Unauthorized

swagger:response connectCoreV1HeadNamespacedPodProxyWithPathUnauthorized
*/
type ConnectCoreV1HeadNamespacedPodProxyWithPathUnauthorized struct {
}

// NewConnectCoreV1HeadNamespacedPodProxyWithPathUnauthorized creates ConnectCoreV1HeadNamespacedPodProxyWithPathUnauthorized with default headers values
func NewConnectCoreV1HeadNamespacedPodProxyWithPathUnauthorized() *ConnectCoreV1HeadNamespacedPodProxyWithPathUnauthorized {

	return &ConnectCoreV1HeadNamespacedPodProxyWithPathUnauthorized{}
}

// WriteResponse to the client
func (o *ConnectCoreV1HeadNamespacedPodProxyWithPathUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
