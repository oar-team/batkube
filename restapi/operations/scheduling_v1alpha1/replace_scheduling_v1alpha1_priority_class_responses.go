// Code generated by go-swagger; DO NOT EDIT.

package scheduling_v1alpha1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// ReplaceSchedulingV1alpha1PriorityClassOKCode is the HTTP code returned for type ReplaceSchedulingV1alpha1PriorityClassOK
const ReplaceSchedulingV1alpha1PriorityClassOKCode int = 200

/*ReplaceSchedulingV1alpha1PriorityClassOK OK

swagger:response replaceSchedulingV1alpha1PriorityClassOK
*/
type ReplaceSchedulingV1alpha1PriorityClassOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPISchedulingV1alpha1PriorityClass `json:"body,omitempty"`
}

// NewReplaceSchedulingV1alpha1PriorityClassOK creates ReplaceSchedulingV1alpha1PriorityClassOK with default headers values
func NewReplaceSchedulingV1alpha1PriorityClassOK() *ReplaceSchedulingV1alpha1PriorityClassOK {

	return &ReplaceSchedulingV1alpha1PriorityClassOK{}
}

// WithPayload adds the payload to the replace scheduling v1alpha1 priority class o k response
func (o *ReplaceSchedulingV1alpha1PriorityClassOK) WithPayload(payload *models.IoK8sAPISchedulingV1alpha1PriorityClass) *ReplaceSchedulingV1alpha1PriorityClassOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace scheduling v1alpha1 priority class o k response
func (o *ReplaceSchedulingV1alpha1PriorityClassOK) SetPayload(payload *models.IoK8sAPISchedulingV1alpha1PriorityClass) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceSchedulingV1alpha1PriorityClassOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceSchedulingV1alpha1PriorityClassCreatedCode is the HTTP code returned for type ReplaceSchedulingV1alpha1PriorityClassCreated
const ReplaceSchedulingV1alpha1PriorityClassCreatedCode int = 201

/*ReplaceSchedulingV1alpha1PriorityClassCreated Created

swagger:response replaceSchedulingV1alpha1PriorityClassCreated
*/
type ReplaceSchedulingV1alpha1PriorityClassCreated struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPISchedulingV1alpha1PriorityClass `json:"body,omitempty"`
}

// NewReplaceSchedulingV1alpha1PriorityClassCreated creates ReplaceSchedulingV1alpha1PriorityClassCreated with default headers values
func NewReplaceSchedulingV1alpha1PriorityClassCreated() *ReplaceSchedulingV1alpha1PriorityClassCreated {

	return &ReplaceSchedulingV1alpha1PriorityClassCreated{}
}

// WithPayload adds the payload to the replace scheduling v1alpha1 priority class created response
func (o *ReplaceSchedulingV1alpha1PriorityClassCreated) WithPayload(payload *models.IoK8sAPISchedulingV1alpha1PriorityClass) *ReplaceSchedulingV1alpha1PriorityClassCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace scheduling v1alpha1 priority class created response
func (o *ReplaceSchedulingV1alpha1PriorityClassCreated) SetPayload(payload *models.IoK8sAPISchedulingV1alpha1PriorityClass) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceSchedulingV1alpha1PriorityClassCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceSchedulingV1alpha1PriorityClassUnauthorizedCode is the HTTP code returned for type ReplaceSchedulingV1alpha1PriorityClassUnauthorized
const ReplaceSchedulingV1alpha1PriorityClassUnauthorizedCode int = 401

/*ReplaceSchedulingV1alpha1PriorityClassUnauthorized Unauthorized

swagger:response replaceSchedulingV1alpha1PriorityClassUnauthorized
*/
type ReplaceSchedulingV1alpha1PriorityClassUnauthorized struct {
}

// NewReplaceSchedulingV1alpha1PriorityClassUnauthorized creates ReplaceSchedulingV1alpha1PriorityClassUnauthorized with default headers values
func NewReplaceSchedulingV1alpha1PriorityClassUnauthorized() *ReplaceSchedulingV1alpha1PriorityClassUnauthorized {

	return &ReplaceSchedulingV1alpha1PriorityClassUnauthorized{}
}

// WriteResponse to the client
func (o *ReplaceSchedulingV1alpha1PriorityClassUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
