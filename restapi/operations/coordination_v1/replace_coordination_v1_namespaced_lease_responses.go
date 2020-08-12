// Code generated by go-swagger; DO NOT EDIT.

package coordination_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// ReplaceCoordinationV1NamespacedLeaseOKCode is the HTTP code returned for type ReplaceCoordinationV1NamespacedLeaseOK
const ReplaceCoordinationV1NamespacedLeaseOKCode int = 200

/*ReplaceCoordinationV1NamespacedLeaseOK OK

swagger:response replaceCoordinationV1NamespacedLeaseOK
*/
type ReplaceCoordinationV1NamespacedLeaseOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPICoordinationV1Lease `json:"body,omitempty"`
}

// NewReplaceCoordinationV1NamespacedLeaseOK creates ReplaceCoordinationV1NamespacedLeaseOK with default headers values
func NewReplaceCoordinationV1NamespacedLeaseOK() *ReplaceCoordinationV1NamespacedLeaseOK {

	return &ReplaceCoordinationV1NamespacedLeaseOK{}
}

// WithPayload adds the payload to the replace coordination v1 namespaced lease o k response
func (o *ReplaceCoordinationV1NamespacedLeaseOK) WithPayload(payload *models.IoK8sAPICoordinationV1Lease) *ReplaceCoordinationV1NamespacedLeaseOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace coordination v1 namespaced lease o k response
func (o *ReplaceCoordinationV1NamespacedLeaseOK) SetPayload(payload *models.IoK8sAPICoordinationV1Lease) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceCoordinationV1NamespacedLeaseOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceCoordinationV1NamespacedLeaseCreatedCode is the HTTP code returned for type ReplaceCoordinationV1NamespacedLeaseCreated
const ReplaceCoordinationV1NamespacedLeaseCreatedCode int = 201

/*ReplaceCoordinationV1NamespacedLeaseCreated Created

swagger:response replaceCoordinationV1NamespacedLeaseCreated
*/
type ReplaceCoordinationV1NamespacedLeaseCreated struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPICoordinationV1Lease `json:"body,omitempty"`
}

// NewReplaceCoordinationV1NamespacedLeaseCreated creates ReplaceCoordinationV1NamespacedLeaseCreated with default headers values
func NewReplaceCoordinationV1NamespacedLeaseCreated() *ReplaceCoordinationV1NamespacedLeaseCreated {

	return &ReplaceCoordinationV1NamespacedLeaseCreated{}
}

// WithPayload adds the payload to the replace coordination v1 namespaced lease created response
func (o *ReplaceCoordinationV1NamespacedLeaseCreated) WithPayload(payload *models.IoK8sAPICoordinationV1Lease) *ReplaceCoordinationV1NamespacedLeaseCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace coordination v1 namespaced lease created response
func (o *ReplaceCoordinationV1NamespacedLeaseCreated) SetPayload(payload *models.IoK8sAPICoordinationV1Lease) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceCoordinationV1NamespacedLeaseCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceCoordinationV1NamespacedLeaseUnauthorizedCode is the HTTP code returned for type ReplaceCoordinationV1NamespacedLeaseUnauthorized
const ReplaceCoordinationV1NamespacedLeaseUnauthorizedCode int = 401

/*ReplaceCoordinationV1NamespacedLeaseUnauthorized Unauthorized

swagger:response replaceCoordinationV1NamespacedLeaseUnauthorized
*/
type ReplaceCoordinationV1NamespacedLeaseUnauthorized struct {
}

// NewReplaceCoordinationV1NamespacedLeaseUnauthorized creates ReplaceCoordinationV1NamespacedLeaseUnauthorized with default headers values
func NewReplaceCoordinationV1NamespacedLeaseUnauthorized() *ReplaceCoordinationV1NamespacedLeaseUnauthorized {

	return &ReplaceCoordinationV1NamespacedLeaseUnauthorized{}
}

// WriteResponse to the client
func (o *ReplaceCoordinationV1NamespacedLeaseUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
