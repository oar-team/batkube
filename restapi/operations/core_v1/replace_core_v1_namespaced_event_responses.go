// Code generated by go-swagger; DO NOT EDIT.

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/models"
)

// ReplaceCoreV1NamespacedEventOKCode is the HTTP code returned for type ReplaceCoreV1NamespacedEventOK
const ReplaceCoreV1NamespacedEventOKCode int = 200

/*ReplaceCoreV1NamespacedEventOK OK

swagger:response replaceCoreV1NamespacedEventOK
*/
type ReplaceCoreV1NamespacedEventOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPICoreV1Event `json:"body,omitempty"`
}

// NewReplaceCoreV1NamespacedEventOK creates ReplaceCoreV1NamespacedEventOK with default headers values
func NewReplaceCoreV1NamespacedEventOK() *ReplaceCoreV1NamespacedEventOK {

	return &ReplaceCoreV1NamespacedEventOK{}
}

// WithPayload adds the payload to the replace core v1 namespaced event o k response
func (o *ReplaceCoreV1NamespacedEventOK) WithPayload(payload *models.IoK8sAPICoreV1Event) *ReplaceCoreV1NamespacedEventOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace core v1 namespaced event o k response
func (o *ReplaceCoreV1NamespacedEventOK) SetPayload(payload *models.IoK8sAPICoreV1Event) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceCoreV1NamespacedEventOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceCoreV1NamespacedEventCreatedCode is the HTTP code returned for type ReplaceCoreV1NamespacedEventCreated
const ReplaceCoreV1NamespacedEventCreatedCode int = 201

/*ReplaceCoreV1NamespacedEventCreated Created

swagger:response replaceCoreV1NamespacedEventCreated
*/
type ReplaceCoreV1NamespacedEventCreated struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPICoreV1Event `json:"body,omitempty"`
}

// NewReplaceCoreV1NamespacedEventCreated creates ReplaceCoreV1NamespacedEventCreated with default headers values
func NewReplaceCoreV1NamespacedEventCreated() *ReplaceCoreV1NamespacedEventCreated {

	return &ReplaceCoreV1NamespacedEventCreated{}
}

// WithPayload adds the payload to the replace core v1 namespaced event created response
func (o *ReplaceCoreV1NamespacedEventCreated) WithPayload(payload *models.IoK8sAPICoreV1Event) *ReplaceCoreV1NamespacedEventCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace core v1 namespaced event created response
func (o *ReplaceCoreV1NamespacedEventCreated) SetPayload(payload *models.IoK8sAPICoreV1Event) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceCoreV1NamespacedEventCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceCoreV1NamespacedEventUnauthorizedCode is the HTTP code returned for type ReplaceCoreV1NamespacedEventUnauthorized
const ReplaceCoreV1NamespacedEventUnauthorizedCode int = 401

/*ReplaceCoreV1NamespacedEventUnauthorized Unauthorized

swagger:response replaceCoreV1NamespacedEventUnauthorized
*/
type ReplaceCoreV1NamespacedEventUnauthorized struct {
}

// NewReplaceCoreV1NamespacedEventUnauthorized creates ReplaceCoreV1NamespacedEventUnauthorized with default headers values
func NewReplaceCoreV1NamespacedEventUnauthorized() *ReplaceCoreV1NamespacedEventUnauthorized {

	return &ReplaceCoreV1NamespacedEventUnauthorized{}
}

// WriteResponse to the client
func (o *ReplaceCoreV1NamespacedEventUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}