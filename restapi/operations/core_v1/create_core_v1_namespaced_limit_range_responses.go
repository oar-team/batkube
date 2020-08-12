// Code generated by go-swagger; DO NOT EDIT.

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// CreateCoreV1NamespacedLimitRangeOKCode is the HTTP code returned for type CreateCoreV1NamespacedLimitRangeOK
const CreateCoreV1NamespacedLimitRangeOKCode int = 200

/*CreateCoreV1NamespacedLimitRangeOK OK

swagger:response createCoreV1NamespacedLimitRangeOK
*/
type CreateCoreV1NamespacedLimitRangeOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPICoreV1LimitRange `json:"body,omitempty"`
}

// NewCreateCoreV1NamespacedLimitRangeOK creates CreateCoreV1NamespacedLimitRangeOK with default headers values
func NewCreateCoreV1NamespacedLimitRangeOK() *CreateCoreV1NamespacedLimitRangeOK {

	return &CreateCoreV1NamespacedLimitRangeOK{}
}

// WithPayload adds the payload to the create core v1 namespaced limit range o k response
func (o *CreateCoreV1NamespacedLimitRangeOK) WithPayload(payload *models.IoK8sAPICoreV1LimitRange) *CreateCoreV1NamespacedLimitRangeOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create core v1 namespaced limit range o k response
func (o *CreateCoreV1NamespacedLimitRangeOK) SetPayload(payload *models.IoK8sAPICoreV1LimitRange) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateCoreV1NamespacedLimitRangeOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateCoreV1NamespacedLimitRangeCreatedCode is the HTTP code returned for type CreateCoreV1NamespacedLimitRangeCreated
const CreateCoreV1NamespacedLimitRangeCreatedCode int = 201

/*CreateCoreV1NamespacedLimitRangeCreated Created

swagger:response createCoreV1NamespacedLimitRangeCreated
*/
type CreateCoreV1NamespacedLimitRangeCreated struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPICoreV1LimitRange `json:"body,omitempty"`
}

// NewCreateCoreV1NamespacedLimitRangeCreated creates CreateCoreV1NamespacedLimitRangeCreated with default headers values
func NewCreateCoreV1NamespacedLimitRangeCreated() *CreateCoreV1NamespacedLimitRangeCreated {

	return &CreateCoreV1NamespacedLimitRangeCreated{}
}

// WithPayload adds the payload to the create core v1 namespaced limit range created response
func (o *CreateCoreV1NamespacedLimitRangeCreated) WithPayload(payload *models.IoK8sAPICoreV1LimitRange) *CreateCoreV1NamespacedLimitRangeCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create core v1 namespaced limit range created response
func (o *CreateCoreV1NamespacedLimitRangeCreated) SetPayload(payload *models.IoK8sAPICoreV1LimitRange) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateCoreV1NamespacedLimitRangeCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateCoreV1NamespacedLimitRangeAcceptedCode is the HTTP code returned for type CreateCoreV1NamespacedLimitRangeAccepted
const CreateCoreV1NamespacedLimitRangeAcceptedCode int = 202

/*CreateCoreV1NamespacedLimitRangeAccepted Accepted

swagger:response createCoreV1NamespacedLimitRangeAccepted
*/
type CreateCoreV1NamespacedLimitRangeAccepted struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPICoreV1LimitRange `json:"body,omitempty"`
}

// NewCreateCoreV1NamespacedLimitRangeAccepted creates CreateCoreV1NamespacedLimitRangeAccepted with default headers values
func NewCreateCoreV1NamespacedLimitRangeAccepted() *CreateCoreV1NamespacedLimitRangeAccepted {

	return &CreateCoreV1NamespacedLimitRangeAccepted{}
}

// WithPayload adds the payload to the create core v1 namespaced limit range accepted response
func (o *CreateCoreV1NamespacedLimitRangeAccepted) WithPayload(payload *models.IoK8sAPICoreV1LimitRange) *CreateCoreV1NamespacedLimitRangeAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create core v1 namespaced limit range accepted response
func (o *CreateCoreV1NamespacedLimitRangeAccepted) SetPayload(payload *models.IoK8sAPICoreV1LimitRange) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateCoreV1NamespacedLimitRangeAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(202)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateCoreV1NamespacedLimitRangeUnauthorizedCode is the HTTP code returned for type CreateCoreV1NamespacedLimitRangeUnauthorized
const CreateCoreV1NamespacedLimitRangeUnauthorizedCode int = 401

/*CreateCoreV1NamespacedLimitRangeUnauthorized Unauthorized

swagger:response createCoreV1NamespacedLimitRangeUnauthorized
*/
type CreateCoreV1NamespacedLimitRangeUnauthorized struct {
}

// NewCreateCoreV1NamespacedLimitRangeUnauthorized creates CreateCoreV1NamespacedLimitRangeUnauthorized with default headers values
func NewCreateCoreV1NamespacedLimitRangeUnauthorized() *CreateCoreV1NamespacedLimitRangeUnauthorized {

	return &CreateCoreV1NamespacedLimitRangeUnauthorized{}
}

// WriteResponse to the client
func (o *CreateCoreV1NamespacedLimitRangeUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
