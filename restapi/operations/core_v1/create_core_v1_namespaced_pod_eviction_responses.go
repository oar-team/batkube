// Code generated by go-swagger; DO NOT EDIT.

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// CreateCoreV1NamespacedPodEvictionOKCode is the HTTP code returned for type CreateCoreV1NamespacedPodEvictionOK
const CreateCoreV1NamespacedPodEvictionOKCode int = 200

/*CreateCoreV1NamespacedPodEvictionOK OK

swagger:response createCoreV1NamespacedPodEvictionOK
*/
type CreateCoreV1NamespacedPodEvictionOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIPolicyV1beta1Eviction `json:"body,omitempty"`
}

// NewCreateCoreV1NamespacedPodEvictionOK creates CreateCoreV1NamespacedPodEvictionOK with default headers values
func NewCreateCoreV1NamespacedPodEvictionOK() *CreateCoreV1NamespacedPodEvictionOK {

	return &CreateCoreV1NamespacedPodEvictionOK{}
}

// WithPayload adds the payload to the create core v1 namespaced pod eviction o k response
func (o *CreateCoreV1NamespacedPodEvictionOK) WithPayload(payload *models.IoK8sAPIPolicyV1beta1Eviction) *CreateCoreV1NamespacedPodEvictionOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create core v1 namespaced pod eviction o k response
func (o *CreateCoreV1NamespacedPodEvictionOK) SetPayload(payload *models.IoK8sAPIPolicyV1beta1Eviction) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateCoreV1NamespacedPodEvictionOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateCoreV1NamespacedPodEvictionCreatedCode is the HTTP code returned for type CreateCoreV1NamespacedPodEvictionCreated
const CreateCoreV1NamespacedPodEvictionCreatedCode int = 201

/*CreateCoreV1NamespacedPodEvictionCreated Created

swagger:response createCoreV1NamespacedPodEvictionCreated
*/
type CreateCoreV1NamespacedPodEvictionCreated struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIPolicyV1beta1Eviction `json:"body,omitempty"`
}

// NewCreateCoreV1NamespacedPodEvictionCreated creates CreateCoreV1NamespacedPodEvictionCreated with default headers values
func NewCreateCoreV1NamespacedPodEvictionCreated() *CreateCoreV1NamespacedPodEvictionCreated {

	return &CreateCoreV1NamespacedPodEvictionCreated{}
}

// WithPayload adds the payload to the create core v1 namespaced pod eviction created response
func (o *CreateCoreV1NamespacedPodEvictionCreated) WithPayload(payload *models.IoK8sAPIPolicyV1beta1Eviction) *CreateCoreV1NamespacedPodEvictionCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create core v1 namespaced pod eviction created response
func (o *CreateCoreV1NamespacedPodEvictionCreated) SetPayload(payload *models.IoK8sAPIPolicyV1beta1Eviction) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateCoreV1NamespacedPodEvictionCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateCoreV1NamespacedPodEvictionAcceptedCode is the HTTP code returned for type CreateCoreV1NamespacedPodEvictionAccepted
const CreateCoreV1NamespacedPodEvictionAcceptedCode int = 202

/*CreateCoreV1NamespacedPodEvictionAccepted Accepted

swagger:response createCoreV1NamespacedPodEvictionAccepted
*/
type CreateCoreV1NamespacedPodEvictionAccepted struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIPolicyV1beta1Eviction `json:"body,omitempty"`
}

// NewCreateCoreV1NamespacedPodEvictionAccepted creates CreateCoreV1NamespacedPodEvictionAccepted with default headers values
func NewCreateCoreV1NamespacedPodEvictionAccepted() *CreateCoreV1NamespacedPodEvictionAccepted {

	return &CreateCoreV1NamespacedPodEvictionAccepted{}
}

// WithPayload adds the payload to the create core v1 namespaced pod eviction accepted response
func (o *CreateCoreV1NamespacedPodEvictionAccepted) WithPayload(payload *models.IoK8sAPIPolicyV1beta1Eviction) *CreateCoreV1NamespacedPodEvictionAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create core v1 namespaced pod eviction accepted response
func (o *CreateCoreV1NamespacedPodEvictionAccepted) SetPayload(payload *models.IoK8sAPIPolicyV1beta1Eviction) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateCoreV1NamespacedPodEvictionAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(202)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateCoreV1NamespacedPodEvictionUnauthorizedCode is the HTTP code returned for type CreateCoreV1NamespacedPodEvictionUnauthorized
const CreateCoreV1NamespacedPodEvictionUnauthorizedCode int = 401

/*CreateCoreV1NamespacedPodEvictionUnauthorized Unauthorized

swagger:response createCoreV1NamespacedPodEvictionUnauthorized
*/
type CreateCoreV1NamespacedPodEvictionUnauthorized struct {
}

// NewCreateCoreV1NamespacedPodEvictionUnauthorized creates CreateCoreV1NamespacedPodEvictionUnauthorized with default headers values
func NewCreateCoreV1NamespacedPodEvictionUnauthorized() *CreateCoreV1NamespacedPodEvictionUnauthorized {

	return &CreateCoreV1NamespacedPodEvictionUnauthorized{}
}

// WriteResponse to the client
func (o *CreateCoreV1NamespacedPodEvictionUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
