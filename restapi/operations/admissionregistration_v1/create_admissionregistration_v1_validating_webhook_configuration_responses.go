// Code generated by go-swagger; DO NOT EDIT.

package admissionregistration_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/models"
)

// CreateAdmissionregistrationV1ValidatingWebhookConfigurationOKCode is the HTTP code returned for type CreateAdmissionregistrationV1ValidatingWebhookConfigurationOK
const CreateAdmissionregistrationV1ValidatingWebhookConfigurationOKCode int = 200

/*CreateAdmissionregistrationV1ValidatingWebhookConfigurationOK OK

swagger:response createAdmissionregistrationV1ValidatingWebhookConfigurationOK
*/
type CreateAdmissionregistrationV1ValidatingWebhookConfigurationOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIAdmissionregistrationV1ValidatingWebhookConfiguration `json:"body,omitempty"`
}

// NewCreateAdmissionregistrationV1ValidatingWebhookConfigurationOK creates CreateAdmissionregistrationV1ValidatingWebhookConfigurationOK with default headers values
func NewCreateAdmissionregistrationV1ValidatingWebhookConfigurationOK() *CreateAdmissionregistrationV1ValidatingWebhookConfigurationOK {

	return &CreateAdmissionregistrationV1ValidatingWebhookConfigurationOK{}
}

// WithPayload adds the payload to the create admissionregistration v1 validating webhook configuration o k response
func (o *CreateAdmissionregistrationV1ValidatingWebhookConfigurationOK) WithPayload(payload *models.IoK8sAPIAdmissionregistrationV1ValidatingWebhookConfiguration) *CreateAdmissionregistrationV1ValidatingWebhookConfigurationOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create admissionregistration v1 validating webhook configuration o k response
func (o *CreateAdmissionregistrationV1ValidatingWebhookConfigurationOK) SetPayload(payload *models.IoK8sAPIAdmissionregistrationV1ValidatingWebhookConfiguration) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateAdmissionregistrationV1ValidatingWebhookConfigurationOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateAdmissionregistrationV1ValidatingWebhookConfigurationCreatedCode is the HTTP code returned for type CreateAdmissionregistrationV1ValidatingWebhookConfigurationCreated
const CreateAdmissionregistrationV1ValidatingWebhookConfigurationCreatedCode int = 201

/*CreateAdmissionregistrationV1ValidatingWebhookConfigurationCreated Created

swagger:response createAdmissionregistrationV1ValidatingWebhookConfigurationCreated
*/
type CreateAdmissionregistrationV1ValidatingWebhookConfigurationCreated struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIAdmissionregistrationV1ValidatingWebhookConfiguration `json:"body,omitempty"`
}

// NewCreateAdmissionregistrationV1ValidatingWebhookConfigurationCreated creates CreateAdmissionregistrationV1ValidatingWebhookConfigurationCreated with default headers values
func NewCreateAdmissionregistrationV1ValidatingWebhookConfigurationCreated() *CreateAdmissionregistrationV1ValidatingWebhookConfigurationCreated {

	return &CreateAdmissionregistrationV1ValidatingWebhookConfigurationCreated{}
}

// WithPayload adds the payload to the create admissionregistration v1 validating webhook configuration created response
func (o *CreateAdmissionregistrationV1ValidatingWebhookConfigurationCreated) WithPayload(payload *models.IoK8sAPIAdmissionregistrationV1ValidatingWebhookConfiguration) *CreateAdmissionregistrationV1ValidatingWebhookConfigurationCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create admissionregistration v1 validating webhook configuration created response
func (o *CreateAdmissionregistrationV1ValidatingWebhookConfigurationCreated) SetPayload(payload *models.IoK8sAPIAdmissionregistrationV1ValidatingWebhookConfiguration) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateAdmissionregistrationV1ValidatingWebhookConfigurationCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateAdmissionregistrationV1ValidatingWebhookConfigurationAcceptedCode is the HTTP code returned for type CreateAdmissionregistrationV1ValidatingWebhookConfigurationAccepted
const CreateAdmissionregistrationV1ValidatingWebhookConfigurationAcceptedCode int = 202

/*CreateAdmissionregistrationV1ValidatingWebhookConfigurationAccepted Accepted

swagger:response createAdmissionregistrationV1ValidatingWebhookConfigurationAccepted
*/
type CreateAdmissionregistrationV1ValidatingWebhookConfigurationAccepted struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIAdmissionregistrationV1ValidatingWebhookConfiguration `json:"body,omitempty"`
}

// NewCreateAdmissionregistrationV1ValidatingWebhookConfigurationAccepted creates CreateAdmissionregistrationV1ValidatingWebhookConfigurationAccepted with default headers values
func NewCreateAdmissionregistrationV1ValidatingWebhookConfigurationAccepted() *CreateAdmissionregistrationV1ValidatingWebhookConfigurationAccepted {

	return &CreateAdmissionregistrationV1ValidatingWebhookConfigurationAccepted{}
}

// WithPayload adds the payload to the create admissionregistration v1 validating webhook configuration accepted response
func (o *CreateAdmissionregistrationV1ValidatingWebhookConfigurationAccepted) WithPayload(payload *models.IoK8sAPIAdmissionregistrationV1ValidatingWebhookConfiguration) *CreateAdmissionregistrationV1ValidatingWebhookConfigurationAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create admissionregistration v1 validating webhook configuration accepted response
func (o *CreateAdmissionregistrationV1ValidatingWebhookConfigurationAccepted) SetPayload(payload *models.IoK8sAPIAdmissionregistrationV1ValidatingWebhookConfiguration) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateAdmissionregistrationV1ValidatingWebhookConfigurationAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(202)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateAdmissionregistrationV1ValidatingWebhookConfigurationUnauthorizedCode is the HTTP code returned for type CreateAdmissionregistrationV1ValidatingWebhookConfigurationUnauthorized
const CreateAdmissionregistrationV1ValidatingWebhookConfigurationUnauthorizedCode int = 401

/*CreateAdmissionregistrationV1ValidatingWebhookConfigurationUnauthorized Unauthorized

swagger:response createAdmissionregistrationV1ValidatingWebhookConfigurationUnauthorized
*/
type CreateAdmissionregistrationV1ValidatingWebhookConfigurationUnauthorized struct {
}

// NewCreateAdmissionregistrationV1ValidatingWebhookConfigurationUnauthorized creates CreateAdmissionregistrationV1ValidatingWebhookConfigurationUnauthorized with default headers values
func NewCreateAdmissionregistrationV1ValidatingWebhookConfigurationUnauthorized() *CreateAdmissionregistrationV1ValidatingWebhookConfigurationUnauthorized {

	return &CreateAdmissionregistrationV1ValidatingWebhookConfigurationUnauthorized{}
}

// WriteResponse to the client
func (o *CreateAdmissionregistrationV1ValidatingWebhookConfigurationUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}