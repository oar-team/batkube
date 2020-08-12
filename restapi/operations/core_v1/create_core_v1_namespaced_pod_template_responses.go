// Code generated by go-swagger; DO NOT EDIT.

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// CreateCoreV1NamespacedPodTemplateOKCode is the HTTP code returned for type CreateCoreV1NamespacedPodTemplateOK
const CreateCoreV1NamespacedPodTemplateOKCode int = 200

/*CreateCoreV1NamespacedPodTemplateOK OK

swagger:response createCoreV1NamespacedPodTemplateOK
*/
type CreateCoreV1NamespacedPodTemplateOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPICoreV1PodTemplate `json:"body,omitempty"`
}

// NewCreateCoreV1NamespacedPodTemplateOK creates CreateCoreV1NamespacedPodTemplateOK with default headers values
func NewCreateCoreV1NamespacedPodTemplateOK() *CreateCoreV1NamespacedPodTemplateOK {

	return &CreateCoreV1NamespacedPodTemplateOK{}
}

// WithPayload adds the payload to the create core v1 namespaced pod template o k response
func (o *CreateCoreV1NamespacedPodTemplateOK) WithPayload(payload *models.IoK8sAPICoreV1PodTemplate) *CreateCoreV1NamespacedPodTemplateOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create core v1 namespaced pod template o k response
func (o *CreateCoreV1NamespacedPodTemplateOK) SetPayload(payload *models.IoK8sAPICoreV1PodTemplate) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateCoreV1NamespacedPodTemplateOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateCoreV1NamespacedPodTemplateCreatedCode is the HTTP code returned for type CreateCoreV1NamespacedPodTemplateCreated
const CreateCoreV1NamespacedPodTemplateCreatedCode int = 201

/*CreateCoreV1NamespacedPodTemplateCreated Created

swagger:response createCoreV1NamespacedPodTemplateCreated
*/
type CreateCoreV1NamespacedPodTemplateCreated struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPICoreV1PodTemplate `json:"body,omitempty"`
}

// NewCreateCoreV1NamespacedPodTemplateCreated creates CreateCoreV1NamespacedPodTemplateCreated with default headers values
func NewCreateCoreV1NamespacedPodTemplateCreated() *CreateCoreV1NamespacedPodTemplateCreated {

	return &CreateCoreV1NamespacedPodTemplateCreated{}
}

// WithPayload adds the payload to the create core v1 namespaced pod template created response
func (o *CreateCoreV1NamespacedPodTemplateCreated) WithPayload(payload *models.IoK8sAPICoreV1PodTemplate) *CreateCoreV1NamespacedPodTemplateCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create core v1 namespaced pod template created response
func (o *CreateCoreV1NamespacedPodTemplateCreated) SetPayload(payload *models.IoK8sAPICoreV1PodTemplate) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateCoreV1NamespacedPodTemplateCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateCoreV1NamespacedPodTemplateAcceptedCode is the HTTP code returned for type CreateCoreV1NamespacedPodTemplateAccepted
const CreateCoreV1NamespacedPodTemplateAcceptedCode int = 202

/*CreateCoreV1NamespacedPodTemplateAccepted Accepted

swagger:response createCoreV1NamespacedPodTemplateAccepted
*/
type CreateCoreV1NamespacedPodTemplateAccepted struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPICoreV1PodTemplate `json:"body,omitempty"`
}

// NewCreateCoreV1NamespacedPodTemplateAccepted creates CreateCoreV1NamespacedPodTemplateAccepted with default headers values
func NewCreateCoreV1NamespacedPodTemplateAccepted() *CreateCoreV1NamespacedPodTemplateAccepted {

	return &CreateCoreV1NamespacedPodTemplateAccepted{}
}

// WithPayload adds the payload to the create core v1 namespaced pod template accepted response
func (o *CreateCoreV1NamespacedPodTemplateAccepted) WithPayload(payload *models.IoK8sAPICoreV1PodTemplate) *CreateCoreV1NamespacedPodTemplateAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create core v1 namespaced pod template accepted response
func (o *CreateCoreV1NamespacedPodTemplateAccepted) SetPayload(payload *models.IoK8sAPICoreV1PodTemplate) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateCoreV1NamespacedPodTemplateAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(202)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateCoreV1NamespacedPodTemplateUnauthorizedCode is the HTTP code returned for type CreateCoreV1NamespacedPodTemplateUnauthorized
const CreateCoreV1NamespacedPodTemplateUnauthorizedCode int = 401

/*CreateCoreV1NamespacedPodTemplateUnauthorized Unauthorized

swagger:response createCoreV1NamespacedPodTemplateUnauthorized
*/
type CreateCoreV1NamespacedPodTemplateUnauthorized struct {
}

// NewCreateCoreV1NamespacedPodTemplateUnauthorized creates CreateCoreV1NamespacedPodTemplateUnauthorized with default headers values
func NewCreateCoreV1NamespacedPodTemplateUnauthorized() *CreateCoreV1NamespacedPodTemplateUnauthorized {

	return &CreateCoreV1NamespacedPodTemplateUnauthorized{}
}

// WriteResponse to the client
func (o *CreateCoreV1NamespacedPodTemplateUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
