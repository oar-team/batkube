// Code generated by go-swagger; DO NOT EDIT.

package coordination_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/models"
)

// ReadCoordinationV1NamespacedLeaseOKCode is the HTTP code returned for type ReadCoordinationV1NamespacedLeaseOK
const ReadCoordinationV1NamespacedLeaseOKCode int = 200

/*ReadCoordinationV1NamespacedLeaseOK OK

swagger:response readCoordinationV1NamespacedLeaseOK
*/
type ReadCoordinationV1NamespacedLeaseOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPICoordinationV1Lease `json:"body,omitempty"`
}

// NewReadCoordinationV1NamespacedLeaseOK creates ReadCoordinationV1NamespacedLeaseOK with default headers values
func NewReadCoordinationV1NamespacedLeaseOK() *ReadCoordinationV1NamespacedLeaseOK {

	return &ReadCoordinationV1NamespacedLeaseOK{}
}

// WithPayload adds the payload to the read coordination v1 namespaced lease o k response
func (o *ReadCoordinationV1NamespacedLeaseOK) WithPayload(payload *models.IoK8sAPICoordinationV1Lease) *ReadCoordinationV1NamespacedLeaseOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the read coordination v1 namespaced lease o k response
func (o *ReadCoordinationV1NamespacedLeaseOK) SetPayload(payload *models.IoK8sAPICoordinationV1Lease) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReadCoordinationV1NamespacedLeaseOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReadCoordinationV1NamespacedLeaseUnauthorizedCode is the HTTP code returned for type ReadCoordinationV1NamespacedLeaseUnauthorized
const ReadCoordinationV1NamespacedLeaseUnauthorizedCode int = 401

/*ReadCoordinationV1NamespacedLeaseUnauthorized Unauthorized

swagger:response readCoordinationV1NamespacedLeaseUnauthorized
*/
type ReadCoordinationV1NamespacedLeaseUnauthorized struct {
}

// NewReadCoordinationV1NamespacedLeaseUnauthorized creates ReadCoordinationV1NamespacedLeaseUnauthorized with default headers values
func NewReadCoordinationV1NamespacedLeaseUnauthorized() *ReadCoordinationV1NamespacedLeaseUnauthorized {

	return &ReadCoordinationV1NamespacedLeaseUnauthorized{}
}

// WriteResponse to the client
func (o *ReadCoordinationV1NamespacedLeaseUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
