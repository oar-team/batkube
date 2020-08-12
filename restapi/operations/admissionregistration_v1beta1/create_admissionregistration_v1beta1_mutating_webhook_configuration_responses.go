// Code generated by go-swagger; DO NOT EDIT.

package admissionregistration_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// CreateAdmissionregistrationV1beta1MutatingWebhookConfigurationOKCode is the HTTP code returned for type CreateAdmissionregistrationV1beta1MutatingWebhookConfigurationOK
const CreateAdmissionregistrationV1beta1MutatingWebhookConfigurationOKCode int = 200

/*CreateAdmissionregistrationV1beta1MutatingWebhookConfigurationOK OK

swagger:response createAdmissionregistrationV1beta1MutatingWebhookConfigurationOK
*/
type CreateAdmissionregistrationV1beta1MutatingWebhookConfigurationOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIAdmissionregistrationV1beta1MutatingWebhookConfiguration `json:"body,omitempty"`
}

// NewCreateAdmissionregistrationV1beta1MutatingWebhookConfigurationOK creates CreateAdmissionregistrationV1beta1MutatingWebhookConfigurationOK with default headers values
func NewCreateAdmissionregistrationV1beta1MutatingWebhookConfigurationOK() *CreateAdmissionregistrationV1beta1MutatingWebhookConfigurationOK {

	return &CreateAdmissionregistrationV1beta1MutatingWebhookConfigurationOK{}
}

// WithPayload adds the payload to the create admissionregistration v1beta1 mutating webhook configuration o k response
func (o *CreateAdmissionregistrationV1beta1MutatingWebhookConfigurationOK) WithPayload(payload *models.IoK8sAPIAdmissionregistrationV1beta1MutatingWebhookConfiguration) *CreateAdmissionregistrationV1beta1MutatingWebhookConfigurationOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create admissionregistration v1beta1 mutating webhook configuration o k response
func (o *CreateAdmissionregistrationV1beta1MutatingWebhookConfigurationOK) SetPayload(payload *models.IoK8sAPIAdmissionregistrationV1beta1MutatingWebhookConfiguration) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateAdmissionregistrationV1beta1MutatingWebhookConfigurationOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateAdmissionregistrationV1beta1MutatingWebhookConfigurationCreatedCode is the HTTP code returned for type CreateAdmissionregistrationV1beta1MutatingWebhookConfigurationCreated
const CreateAdmissionregistrationV1beta1MutatingWebhookConfigurationCreatedCode int = 201

/*CreateAdmissionregistrationV1beta1MutatingWebhookConfigurationCreated Created

swagger:response createAdmissionregistrationV1beta1MutatingWebhookConfigurationCreated
*/
type CreateAdmissionregistrationV1beta1MutatingWebhookConfigurationCreated struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIAdmissionregistrationV1beta1MutatingWebhookConfiguration `json:"body,omitempty"`
}

// NewCreateAdmissionregistrationV1beta1MutatingWebhookConfigurationCreated creates CreateAdmissionregistrationV1beta1MutatingWebhookConfigurationCreated with default headers values
func NewCreateAdmissionregistrationV1beta1MutatingWebhookConfigurationCreated() *CreateAdmissionregistrationV1beta1MutatingWebhookConfigurationCreated {

	return &CreateAdmissionregistrationV1beta1MutatingWebhookConfigurationCreated{}
}

// WithPayload adds the payload to the create admissionregistration v1beta1 mutating webhook configuration created response
func (o *CreateAdmissionregistrationV1beta1MutatingWebhookConfigurationCreated) WithPayload(payload *models.IoK8sAPIAdmissionregistrationV1beta1MutatingWebhookConfiguration) *CreateAdmissionregistrationV1beta1MutatingWebhookConfigurationCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create admissionregistration v1beta1 mutating webhook configuration created response
func (o *CreateAdmissionregistrationV1beta1MutatingWebhookConfigurationCreated) SetPayload(payload *models.IoK8sAPIAdmissionregistrationV1beta1MutatingWebhookConfiguration) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateAdmissionregistrationV1beta1MutatingWebhookConfigurationCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateAdmissionregistrationV1beta1MutatingWebhookConfigurationAcceptedCode is the HTTP code returned for type CreateAdmissionregistrationV1beta1MutatingWebhookConfigurationAccepted
const CreateAdmissionregistrationV1beta1MutatingWebhookConfigurationAcceptedCode int = 202

/*CreateAdmissionregistrationV1beta1MutatingWebhookConfigurationAccepted Accepted

swagger:response createAdmissionregistrationV1beta1MutatingWebhookConfigurationAccepted
*/
type CreateAdmissionregistrationV1beta1MutatingWebhookConfigurationAccepted struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIAdmissionregistrationV1beta1MutatingWebhookConfiguration `json:"body,omitempty"`
}

// NewCreateAdmissionregistrationV1beta1MutatingWebhookConfigurationAccepted creates CreateAdmissionregistrationV1beta1MutatingWebhookConfigurationAccepted with default headers values
func NewCreateAdmissionregistrationV1beta1MutatingWebhookConfigurationAccepted() *CreateAdmissionregistrationV1beta1MutatingWebhookConfigurationAccepted {

	return &CreateAdmissionregistrationV1beta1MutatingWebhookConfigurationAccepted{}
}

// WithPayload adds the payload to the create admissionregistration v1beta1 mutating webhook configuration accepted response
func (o *CreateAdmissionregistrationV1beta1MutatingWebhookConfigurationAccepted) WithPayload(payload *models.IoK8sAPIAdmissionregistrationV1beta1MutatingWebhookConfiguration) *CreateAdmissionregistrationV1beta1MutatingWebhookConfigurationAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create admissionregistration v1beta1 mutating webhook configuration accepted response
func (o *CreateAdmissionregistrationV1beta1MutatingWebhookConfigurationAccepted) SetPayload(payload *models.IoK8sAPIAdmissionregistrationV1beta1MutatingWebhookConfiguration) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateAdmissionregistrationV1beta1MutatingWebhookConfigurationAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(202)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateAdmissionregistrationV1beta1MutatingWebhookConfigurationUnauthorizedCode is the HTTP code returned for type CreateAdmissionregistrationV1beta1MutatingWebhookConfigurationUnauthorized
const CreateAdmissionregistrationV1beta1MutatingWebhookConfigurationUnauthorizedCode int = 401

/*CreateAdmissionregistrationV1beta1MutatingWebhookConfigurationUnauthorized Unauthorized

swagger:response createAdmissionregistrationV1beta1MutatingWebhookConfigurationUnauthorized
*/
type CreateAdmissionregistrationV1beta1MutatingWebhookConfigurationUnauthorized struct {
}

// NewCreateAdmissionregistrationV1beta1MutatingWebhookConfigurationUnauthorized creates CreateAdmissionregistrationV1beta1MutatingWebhookConfigurationUnauthorized with default headers values
func NewCreateAdmissionregistrationV1beta1MutatingWebhookConfigurationUnauthorized() *CreateAdmissionregistrationV1beta1MutatingWebhookConfigurationUnauthorized {

	return &CreateAdmissionregistrationV1beta1MutatingWebhookConfigurationUnauthorized{}
}

// WriteResponse to the client
func (o *CreateAdmissionregistrationV1beta1MutatingWebhookConfigurationUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
