// Code generated by go-swagger; DO NOT EDIT.

package admissionregistration_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/models"
)

// CreateAdmissionregistrationV1MutatingWebhookConfigurationOKCode is the HTTP code returned for type CreateAdmissionregistrationV1MutatingWebhookConfigurationOK
const CreateAdmissionregistrationV1MutatingWebhookConfigurationOKCode int = 200

/*CreateAdmissionregistrationV1MutatingWebhookConfigurationOK OK

swagger:response createAdmissionregistrationV1MutatingWebhookConfigurationOK
*/
type CreateAdmissionregistrationV1MutatingWebhookConfigurationOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIAdmissionregistrationV1MutatingWebhookConfiguration `json:"body,omitempty"`
}

// NewCreateAdmissionregistrationV1MutatingWebhookConfigurationOK creates CreateAdmissionregistrationV1MutatingWebhookConfigurationOK with default headers values
func NewCreateAdmissionregistrationV1MutatingWebhookConfigurationOK() *CreateAdmissionregistrationV1MutatingWebhookConfigurationOK {

	return &CreateAdmissionregistrationV1MutatingWebhookConfigurationOK{}
}

// WithPayload adds the payload to the create admissionregistration v1 mutating webhook configuration o k response
func (o *CreateAdmissionregistrationV1MutatingWebhookConfigurationOK) WithPayload(payload *models.IoK8sAPIAdmissionregistrationV1MutatingWebhookConfiguration) *CreateAdmissionregistrationV1MutatingWebhookConfigurationOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create admissionregistration v1 mutating webhook configuration o k response
func (o *CreateAdmissionregistrationV1MutatingWebhookConfigurationOK) SetPayload(payload *models.IoK8sAPIAdmissionregistrationV1MutatingWebhookConfiguration) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateAdmissionregistrationV1MutatingWebhookConfigurationOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateAdmissionregistrationV1MutatingWebhookConfigurationCreatedCode is the HTTP code returned for type CreateAdmissionregistrationV1MutatingWebhookConfigurationCreated
const CreateAdmissionregistrationV1MutatingWebhookConfigurationCreatedCode int = 201

/*CreateAdmissionregistrationV1MutatingWebhookConfigurationCreated Created

swagger:response createAdmissionregistrationV1MutatingWebhookConfigurationCreated
*/
type CreateAdmissionregistrationV1MutatingWebhookConfigurationCreated struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIAdmissionregistrationV1MutatingWebhookConfiguration `json:"body,omitempty"`
}

// NewCreateAdmissionregistrationV1MutatingWebhookConfigurationCreated creates CreateAdmissionregistrationV1MutatingWebhookConfigurationCreated with default headers values
func NewCreateAdmissionregistrationV1MutatingWebhookConfigurationCreated() *CreateAdmissionregistrationV1MutatingWebhookConfigurationCreated {

	return &CreateAdmissionregistrationV1MutatingWebhookConfigurationCreated{}
}

// WithPayload adds the payload to the create admissionregistration v1 mutating webhook configuration created response
func (o *CreateAdmissionregistrationV1MutatingWebhookConfigurationCreated) WithPayload(payload *models.IoK8sAPIAdmissionregistrationV1MutatingWebhookConfiguration) *CreateAdmissionregistrationV1MutatingWebhookConfigurationCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create admissionregistration v1 mutating webhook configuration created response
func (o *CreateAdmissionregistrationV1MutatingWebhookConfigurationCreated) SetPayload(payload *models.IoK8sAPIAdmissionregistrationV1MutatingWebhookConfiguration) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateAdmissionregistrationV1MutatingWebhookConfigurationCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateAdmissionregistrationV1MutatingWebhookConfigurationAcceptedCode is the HTTP code returned for type CreateAdmissionregistrationV1MutatingWebhookConfigurationAccepted
const CreateAdmissionregistrationV1MutatingWebhookConfigurationAcceptedCode int = 202

/*CreateAdmissionregistrationV1MutatingWebhookConfigurationAccepted Accepted

swagger:response createAdmissionregistrationV1MutatingWebhookConfigurationAccepted
*/
type CreateAdmissionregistrationV1MutatingWebhookConfigurationAccepted struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIAdmissionregistrationV1MutatingWebhookConfiguration `json:"body,omitempty"`
}

// NewCreateAdmissionregistrationV1MutatingWebhookConfigurationAccepted creates CreateAdmissionregistrationV1MutatingWebhookConfigurationAccepted with default headers values
func NewCreateAdmissionregistrationV1MutatingWebhookConfigurationAccepted() *CreateAdmissionregistrationV1MutatingWebhookConfigurationAccepted {

	return &CreateAdmissionregistrationV1MutatingWebhookConfigurationAccepted{}
}

// WithPayload adds the payload to the create admissionregistration v1 mutating webhook configuration accepted response
func (o *CreateAdmissionregistrationV1MutatingWebhookConfigurationAccepted) WithPayload(payload *models.IoK8sAPIAdmissionregistrationV1MutatingWebhookConfiguration) *CreateAdmissionregistrationV1MutatingWebhookConfigurationAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create admissionregistration v1 mutating webhook configuration accepted response
func (o *CreateAdmissionregistrationV1MutatingWebhookConfigurationAccepted) SetPayload(payload *models.IoK8sAPIAdmissionregistrationV1MutatingWebhookConfiguration) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateAdmissionregistrationV1MutatingWebhookConfigurationAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(202)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateAdmissionregistrationV1MutatingWebhookConfigurationUnauthorizedCode is the HTTP code returned for type CreateAdmissionregistrationV1MutatingWebhookConfigurationUnauthorized
const CreateAdmissionregistrationV1MutatingWebhookConfigurationUnauthorizedCode int = 401

/*CreateAdmissionregistrationV1MutatingWebhookConfigurationUnauthorized Unauthorized

swagger:response createAdmissionregistrationV1MutatingWebhookConfigurationUnauthorized
*/
type CreateAdmissionregistrationV1MutatingWebhookConfigurationUnauthorized struct {
}

// NewCreateAdmissionregistrationV1MutatingWebhookConfigurationUnauthorized creates CreateAdmissionregistrationV1MutatingWebhookConfigurationUnauthorized with default headers values
func NewCreateAdmissionregistrationV1MutatingWebhookConfigurationUnauthorized() *CreateAdmissionregistrationV1MutatingWebhookConfigurationUnauthorized {

	return &CreateAdmissionregistrationV1MutatingWebhookConfigurationUnauthorized{}
}

// WriteResponse to the client
func (o *CreateAdmissionregistrationV1MutatingWebhookConfigurationUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
