// Code generated by go-swagger; DO NOT EDIT.

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// ReplaceCoreV1NamespacedServiceStatusOKCode is the HTTP code returned for type ReplaceCoreV1NamespacedServiceStatusOK
const ReplaceCoreV1NamespacedServiceStatusOKCode int = 200

/*ReplaceCoreV1NamespacedServiceStatusOK OK

swagger:response replaceCoreV1NamespacedServiceStatusOK
*/
type ReplaceCoreV1NamespacedServiceStatusOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPICoreV1Service `json:"body,omitempty"`
}

// NewReplaceCoreV1NamespacedServiceStatusOK creates ReplaceCoreV1NamespacedServiceStatusOK with default headers values
func NewReplaceCoreV1NamespacedServiceStatusOK() *ReplaceCoreV1NamespacedServiceStatusOK {

	return &ReplaceCoreV1NamespacedServiceStatusOK{}
}

// WithPayload adds the payload to the replace core v1 namespaced service status o k response
func (o *ReplaceCoreV1NamespacedServiceStatusOK) WithPayload(payload *models.IoK8sAPICoreV1Service) *ReplaceCoreV1NamespacedServiceStatusOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace core v1 namespaced service status o k response
func (o *ReplaceCoreV1NamespacedServiceStatusOK) SetPayload(payload *models.IoK8sAPICoreV1Service) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceCoreV1NamespacedServiceStatusOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceCoreV1NamespacedServiceStatusCreatedCode is the HTTP code returned for type ReplaceCoreV1NamespacedServiceStatusCreated
const ReplaceCoreV1NamespacedServiceStatusCreatedCode int = 201

/*ReplaceCoreV1NamespacedServiceStatusCreated Created

swagger:response replaceCoreV1NamespacedServiceStatusCreated
*/
type ReplaceCoreV1NamespacedServiceStatusCreated struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPICoreV1Service `json:"body,omitempty"`
}

// NewReplaceCoreV1NamespacedServiceStatusCreated creates ReplaceCoreV1NamespacedServiceStatusCreated with default headers values
func NewReplaceCoreV1NamespacedServiceStatusCreated() *ReplaceCoreV1NamespacedServiceStatusCreated {

	return &ReplaceCoreV1NamespacedServiceStatusCreated{}
}

// WithPayload adds the payload to the replace core v1 namespaced service status created response
func (o *ReplaceCoreV1NamespacedServiceStatusCreated) WithPayload(payload *models.IoK8sAPICoreV1Service) *ReplaceCoreV1NamespacedServiceStatusCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace core v1 namespaced service status created response
func (o *ReplaceCoreV1NamespacedServiceStatusCreated) SetPayload(payload *models.IoK8sAPICoreV1Service) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceCoreV1NamespacedServiceStatusCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceCoreV1NamespacedServiceStatusUnauthorizedCode is the HTTP code returned for type ReplaceCoreV1NamespacedServiceStatusUnauthorized
const ReplaceCoreV1NamespacedServiceStatusUnauthorizedCode int = 401

/*ReplaceCoreV1NamespacedServiceStatusUnauthorized Unauthorized

swagger:response replaceCoreV1NamespacedServiceStatusUnauthorized
*/
type ReplaceCoreV1NamespacedServiceStatusUnauthorized struct {
}

// NewReplaceCoreV1NamespacedServiceStatusUnauthorized creates ReplaceCoreV1NamespacedServiceStatusUnauthorized with default headers values
func NewReplaceCoreV1NamespacedServiceStatusUnauthorized() *ReplaceCoreV1NamespacedServiceStatusUnauthorized {

	return &ReplaceCoreV1NamespacedServiceStatusUnauthorized{}
}

// WriteResponse to the client
func (o *ReplaceCoreV1NamespacedServiceStatusUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
