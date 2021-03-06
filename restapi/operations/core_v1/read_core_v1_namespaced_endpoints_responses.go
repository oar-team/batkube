// Code generated by go-swagger; DO NOT EDIT.

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// ReadCoreV1NamespacedEndpointsOKCode is the HTTP code returned for type ReadCoreV1NamespacedEndpointsOK
const ReadCoreV1NamespacedEndpointsOKCode int = 200

/*ReadCoreV1NamespacedEndpointsOK OK

swagger:response readCoreV1NamespacedEndpointsOK
*/
type ReadCoreV1NamespacedEndpointsOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPICoreV1Endpoints `json:"body,omitempty"`
}

// NewReadCoreV1NamespacedEndpointsOK creates ReadCoreV1NamespacedEndpointsOK with default headers values
func NewReadCoreV1NamespacedEndpointsOK() *ReadCoreV1NamespacedEndpointsOK {

	return &ReadCoreV1NamespacedEndpointsOK{}
}

// WithPayload adds the payload to the read core v1 namespaced endpoints o k response
func (o *ReadCoreV1NamespacedEndpointsOK) WithPayload(payload *models.IoK8sAPICoreV1Endpoints) *ReadCoreV1NamespacedEndpointsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the read core v1 namespaced endpoints o k response
func (o *ReadCoreV1NamespacedEndpointsOK) SetPayload(payload *models.IoK8sAPICoreV1Endpoints) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReadCoreV1NamespacedEndpointsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReadCoreV1NamespacedEndpointsUnauthorizedCode is the HTTP code returned for type ReadCoreV1NamespacedEndpointsUnauthorized
const ReadCoreV1NamespacedEndpointsUnauthorizedCode int = 401

/*ReadCoreV1NamespacedEndpointsUnauthorized Unauthorized

swagger:response readCoreV1NamespacedEndpointsUnauthorized
*/
type ReadCoreV1NamespacedEndpointsUnauthorized struct {
}

// NewReadCoreV1NamespacedEndpointsUnauthorized creates ReadCoreV1NamespacedEndpointsUnauthorized with default headers values
func NewReadCoreV1NamespacedEndpointsUnauthorized() *ReadCoreV1NamespacedEndpointsUnauthorized {

	return &ReadCoreV1NamespacedEndpointsUnauthorized{}
}

// WriteResponse to the client
func (o *ReadCoreV1NamespacedEndpointsUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
