// Code generated by go-swagger; DO NOT EDIT.

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// CreateCoreV1NamespacedServiceOKCode is the HTTP code returned for type CreateCoreV1NamespacedServiceOK
const CreateCoreV1NamespacedServiceOKCode int = 200

/*CreateCoreV1NamespacedServiceOK OK

swagger:response createCoreV1NamespacedServiceOK
*/
type CreateCoreV1NamespacedServiceOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPICoreV1Service `json:"body,omitempty"`
}

// NewCreateCoreV1NamespacedServiceOK creates CreateCoreV1NamespacedServiceOK with default headers values
func NewCreateCoreV1NamespacedServiceOK() *CreateCoreV1NamespacedServiceOK {

	return &CreateCoreV1NamespacedServiceOK{}
}

// WithPayload adds the payload to the create core v1 namespaced service o k response
func (o *CreateCoreV1NamespacedServiceOK) WithPayload(payload *models.IoK8sAPICoreV1Service) *CreateCoreV1NamespacedServiceOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create core v1 namespaced service o k response
func (o *CreateCoreV1NamespacedServiceOK) SetPayload(payload *models.IoK8sAPICoreV1Service) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateCoreV1NamespacedServiceOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateCoreV1NamespacedServiceCreatedCode is the HTTP code returned for type CreateCoreV1NamespacedServiceCreated
const CreateCoreV1NamespacedServiceCreatedCode int = 201

/*CreateCoreV1NamespacedServiceCreated Created

swagger:response createCoreV1NamespacedServiceCreated
*/
type CreateCoreV1NamespacedServiceCreated struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPICoreV1Service `json:"body,omitempty"`
}

// NewCreateCoreV1NamespacedServiceCreated creates CreateCoreV1NamespacedServiceCreated with default headers values
func NewCreateCoreV1NamespacedServiceCreated() *CreateCoreV1NamespacedServiceCreated {

	return &CreateCoreV1NamespacedServiceCreated{}
}

// WithPayload adds the payload to the create core v1 namespaced service created response
func (o *CreateCoreV1NamespacedServiceCreated) WithPayload(payload *models.IoK8sAPICoreV1Service) *CreateCoreV1NamespacedServiceCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create core v1 namespaced service created response
func (o *CreateCoreV1NamespacedServiceCreated) SetPayload(payload *models.IoK8sAPICoreV1Service) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateCoreV1NamespacedServiceCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateCoreV1NamespacedServiceAcceptedCode is the HTTP code returned for type CreateCoreV1NamespacedServiceAccepted
const CreateCoreV1NamespacedServiceAcceptedCode int = 202

/*CreateCoreV1NamespacedServiceAccepted Accepted

swagger:response createCoreV1NamespacedServiceAccepted
*/
type CreateCoreV1NamespacedServiceAccepted struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPICoreV1Service `json:"body,omitempty"`
}

// NewCreateCoreV1NamespacedServiceAccepted creates CreateCoreV1NamespacedServiceAccepted with default headers values
func NewCreateCoreV1NamespacedServiceAccepted() *CreateCoreV1NamespacedServiceAccepted {

	return &CreateCoreV1NamespacedServiceAccepted{}
}

// WithPayload adds the payload to the create core v1 namespaced service accepted response
func (o *CreateCoreV1NamespacedServiceAccepted) WithPayload(payload *models.IoK8sAPICoreV1Service) *CreateCoreV1NamespacedServiceAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create core v1 namespaced service accepted response
func (o *CreateCoreV1NamespacedServiceAccepted) SetPayload(payload *models.IoK8sAPICoreV1Service) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateCoreV1NamespacedServiceAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(202)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateCoreV1NamespacedServiceUnauthorizedCode is the HTTP code returned for type CreateCoreV1NamespacedServiceUnauthorized
const CreateCoreV1NamespacedServiceUnauthorizedCode int = 401

/*CreateCoreV1NamespacedServiceUnauthorized Unauthorized

swagger:response createCoreV1NamespacedServiceUnauthorized
*/
type CreateCoreV1NamespacedServiceUnauthorized struct {
}

// NewCreateCoreV1NamespacedServiceUnauthorized creates CreateCoreV1NamespacedServiceUnauthorized with default headers values
func NewCreateCoreV1NamespacedServiceUnauthorized() *CreateCoreV1NamespacedServiceUnauthorized {

	return &CreateCoreV1NamespacedServiceUnauthorized{}
}

// WriteResponse to the client
func (o *CreateCoreV1NamespacedServiceUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
