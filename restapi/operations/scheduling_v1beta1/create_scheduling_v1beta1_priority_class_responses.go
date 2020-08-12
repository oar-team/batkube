// Code generated by go-swagger; DO NOT EDIT.

package scheduling_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// CreateSchedulingV1beta1PriorityClassOKCode is the HTTP code returned for type CreateSchedulingV1beta1PriorityClassOK
const CreateSchedulingV1beta1PriorityClassOKCode int = 200

/*CreateSchedulingV1beta1PriorityClassOK OK

swagger:response createSchedulingV1beta1PriorityClassOK
*/
type CreateSchedulingV1beta1PriorityClassOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPISchedulingV1beta1PriorityClass `json:"body,omitempty"`
}

// NewCreateSchedulingV1beta1PriorityClassOK creates CreateSchedulingV1beta1PriorityClassOK with default headers values
func NewCreateSchedulingV1beta1PriorityClassOK() *CreateSchedulingV1beta1PriorityClassOK {

	return &CreateSchedulingV1beta1PriorityClassOK{}
}

// WithPayload adds the payload to the create scheduling v1beta1 priority class o k response
func (o *CreateSchedulingV1beta1PriorityClassOK) WithPayload(payload *models.IoK8sAPISchedulingV1beta1PriorityClass) *CreateSchedulingV1beta1PriorityClassOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create scheduling v1beta1 priority class o k response
func (o *CreateSchedulingV1beta1PriorityClassOK) SetPayload(payload *models.IoK8sAPISchedulingV1beta1PriorityClass) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateSchedulingV1beta1PriorityClassOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateSchedulingV1beta1PriorityClassCreatedCode is the HTTP code returned for type CreateSchedulingV1beta1PriorityClassCreated
const CreateSchedulingV1beta1PriorityClassCreatedCode int = 201

/*CreateSchedulingV1beta1PriorityClassCreated Created

swagger:response createSchedulingV1beta1PriorityClassCreated
*/
type CreateSchedulingV1beta1PriorityClassCreated struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPISchedulingV1beta1PriorityClass `json:"body,omitempty"`
}

// NewCreateSchedulingV1beta1PriorityClassCreated creates CreateSchedulingV1beta1PriorityClassCreated with default headers values
func NewCreateSchedulingV1beta1PriorityClassCreated() *CreateSchedulingV1beta1PriorityClassCreated {

	return &CreateSchedulingV1beta1PriorityClassCreated{}
}

// WithPayload adds the payload to the create scheduling v1beta1 priority class created response
func (o *CreateSchedulingV1beta1PriorityClassCreated) WithPayload(payload *models.IoK8sAPISchedulingV1beta1PriorityClass) *CreateSchedulingV1beta1PriorityClassCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create scheduling v1beta1 priority class created response
func (o *CreateSchedulingV1beta1PriorityClassCreated) SetPayload(payload *models.IoK8sAPISchedulingV1beta1PriorityClass) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateSchedulingV1beta1PriorityClassCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateSchedulingV1beta1PriorityClassAcceptedCode is the HTTP code returned for type CreateSchedulingV1beta1PriorityClassAccepted
const CreateSchedulingV1beta1PriorityClassAcceptedCode int = 202

/*CreateSchedulingV1beta1PriorityClassAccepted Accepted

swagger:response createSchedulingV1beta1PriorityClassAccepted
*/
type CreateSchedulingV1beta1PriorityClassAccepted struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPISchedulingV1beta1PriorityClass `json:"body,omitempty"`
}

// NewCreateSchedulingV1beta1PriorityClassAccepted creates CreateSchedulingV1beta1PriorityClassAccepted with default headers values
func NewCreateSchedulingV1beta1PriorityClassAccepted() *CreateSchedulingV1beta1PriorityClassAccepted {

	return &CreateSchedulingV1beta1PriorityClassAccepted{}
}

// WithPayload adds the payload to the create scheduling v1beta1 priority class accepted response
func (o *CreateSchedulingV1beta1PriorityClassAccepted) WithPayload(payload *models.IoK8sAPISchedulingV1beta1PriorityClass) *CreateSchedulingV1beta1PriorityClassAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create scheduling v1beta1 priority class accepted response
func (o *CreateSchedulingV1beta1PriorityClassAccepted) SetPayload(payload *models.IoK8sAPISchedulingV1beta1PriorityClass) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateSchedulingV1beta1PriorityClassAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(202)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateSchedulingV1beta1PriorityClassUnauthorizedCode is the HTTP code returned for type CreateSchedulingV1beta1PriorityClassUnauthorized
const CreateSchedulingV1beta1PriorityClassUnauthorizedCode int = 401

/*CreateSchedulingV1beta1PriorityClassUnauthorized Unauthorized

swagger:response createSchedulingV1beta1PriorityClassUnauthorized
*/
type CreateSchedulingV1beta1PriorityClassUnauthorized struct {
}

// NewCreateSchedulingV1beta1PriorityClassUnauthorized creates CreateSchedulingV1beta1PriorityClassUnauthorized with default headers values
func NewCreateSchedulingV1beta1PriorityClassUnauthorized() *CreateSchedulingV1beta1PriorityClassUnauthorized {

	return &CreateSchedulingV1beta1PriorityClassUnauthorized{}
}

// WriteResponse to the client
func (o *CreateSchedulingV1beta1PriorityClassUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
