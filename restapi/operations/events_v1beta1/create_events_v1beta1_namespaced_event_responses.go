// Code generated by go-swagger; DO NOT EDIT.

package events_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/models"
)

// CreateEventsV1beta1NamespacedEventOKCode is the HTTP code returned for type CreateEventsV1beta1NamespacedEventOK
const CreateEventsV1beta1NamespacedEventOKCode int = 200

/*CreateEventsV1beta1NamespacedEventOK OK

swagger:response createEventsV1beta1NamespacedEventOK
*/
type CreateEventsV1beta1NamespacedEventOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIEventsV1beta1Event `json:"body,omitempty"`
}

// NewCreateEventsV1beta1NamespacedEventOK creates CreateEventsV1beta1NamespacedEventOK with default headers values
func NewCreateEventsV1beta1NamespacedEventOK() *CreateEventsV1beta1NamespacedEventOK {

	return &CreateEventsV1beta1NamespacedEventOK{}
}

// WithPayload adds the payload to the create events v1beta1 namespaced event o k response
func (o *CreateEventsV1beta1NamespacedEventOK) WithPayload(payload *models.IoK8sAPIEventsV1beta1Event) *CreateEventsV1beta1NamespacedEventOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create events v1beta1 namespaced event o k response
func (o *CreateEventsV1beta1NamespacedEventOK) SetPayload(payload *models.IoK8sAPIEventsV1beta1Event) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateEventsV1beta1NamespacedEventOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateEventsV1beta1NamespacedEventCreatedCode is the HTTP code returned for type CreateEventsV1beta1NamespacedEventCreated
const CreateEventsV1beta1NamespacedEventCreatedCode int = 201

/*CreateEventsV1beta1NamespacedEventCreated Created

swagger:response createEventsV1beta1NamespacedEventCreated
*/
type CreateEventsV1beta1NamespacedEventCreated struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIEventsV1beta1Event `json:"body,omitempty"`
}

// NewCreateEventsV1beta1NamespacedEventCreated creates CreateEventsV1beta1NamespacedEventCreated with default headers values
func NewCreateEventsV1beta1NamespacedEventCreated() *CreateEventsV1beta1NamespacedEventCreated {

	return &CreateEventsV1beta1NamespacedEventCreated{}
}

// WithPayload adds the payload to the create events v1beta1 namespaced event created response
func (o *CreateEventsV1beta1NamespacedEventCreated) WithPayload(payload *models.IoK8sAPIEventsV1beta1Event) *CreateEventsV1beta1NamespacedEventCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create events v1beta1 namespaced event created response
func (o *CreateEventsV1beta1NamespacedEventCreated) SetPayload(payload *models.IoK8sAPIEventsV1beta1Event) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateEventsV1beta1NamespacedEventCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateEventsV1beta1NamespacedEventAcceptedCode is the HTTP code returned for type CreateEventsV1beta1NamespacedEventAccepted
const CreateEventsV1beta1NamespacedEventAcceptedCode int = 202

/*CreateEventsV1beta1NamespacedEventAccepted Accepted

swagger:response createEventsV1beta1NamespacedEventAccepted
*/
type CreateEventsV1beta1NamespacedEventAccepted struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIEventsV1beta1Event `json:"body,omitempty"`
}

// NewCreateEventsV1beta1NamespacedEventAccepted creates CreateEventsV1beta1NamespacedEventAccepted with default headers values
func NewCreateEventsV1beta1NamespacedEventAccepted() *CreateEventsV1beta1NamespacedEventAccepted {

	return &CreateEventsV1beta1NamespacedEventAccepted{}
}

// WithPayload adds the payload to the create events v1beta1 namespaced event accepted response
func (o *CreateEventsV1beta1NamespacedEventAccepted) WithPayload(payload *models.IoK8sAPIEventsV1beta1Event) *CreateEventsV1beta1NamespacedEventAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create events v1beta1 namespaced event accepted response
func (o *CreateEventsV1beta1NamespacedEventAccepted) SetPayload(payload *models.IoK8sAPIEventsV1beta1Event) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateEventsV1beta1NamespacedEventAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(202)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateEventsV1beta1NamespacedEventUnauthorizedCode is the HTTP code returned for type CreateEventsV1beta1NamespacedEventUnauthorized
const CreateEventsV1beta1NamespacedEventUnauthorizedCode int = 401

/*CreateEventsV1beta1NamespacedEventUnauthorized Unauthorized

swagger:response createEventsV1beta1NamespacedEventUnauthorized
*/
type CreateEventsV1beta1NamespacedEventUnauthorized struct {
}

// NewCreateEventsV1beta1NamespacedEventUnauthorized creates CreateEventsV1beta1NamespacedEventUnauthorized with default headers values
func NewCreateEventsV1beta1NamespacedEventUnauthorized() *CreateEventsV1beta1NamespacedEventUnauthorized {

	return &CreateEventsV1beta1NamespacedEventUnauthorized{}
}

// WriteResponse to the client
func (o *CreateEventsV1beta1NamespacedEventUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}