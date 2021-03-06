// Code generated by go-swagger; DO NOT EDIT.

package admissionregistration_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// PatchAdmissionregistrationV1ValidatingWebhookConfigurationOKCode is the HTTP code returned for type PatchAdmissionregistrationV1ValidatingWebhookConfigurationOK
const PatchAdmissionregistrationV1ValidatingWebhookConfigurationOKCode int = 200

/*PatchAdmissionregistrationV1ValidatingWebhookConfigurationOK OK

swagger:response patchAdmissionregistrationV1ValidatingWebhookConfigurationOK
*/
type PatchAdmissionregistrationV1ValidatingWebhookConfigurationOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIAdmissionregistrationV1ValidatingWebhookConfiguration `json:"body,omitempty"`
}

// NewPatchAdmissionregistrationV1ValidatingWebhookConfigurationOK creates PatchAdmissionregistrationV1ValidatingWebhookConfigurationOK with default headers values
func NewPatchAdmissionregistrationV1ValidatingWebhookConfigurationOK() *PatchAdmissionregistrationV1ValidatingWebhookConfigurationOK {

	return &PatchAdmissionregistrationV1ValidatingWebhookConfigurationOK{}
}

// WithPayload adds the payload to the patch admissionregistration v1 validating webhook configuration o k response
func (o *PatchAdmissionregistrationV1ValidatingWebhookConfigurationOK) WithPayload(payload *models.IoK8sAPIAdmissionregistrationV1ValidatingWebhookConfiguration) *PatchAdmissionregistrationV1ValidatingWebhookConfigurationOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch admissionregistration v1 validating webhook configuration o k response
func (o *PatchAdmissionregistrationV1ValidatingWebhookConfigurationOK) SetPayload(payload *models.IoK8sAPIAdmissionregistrationV1ValidatingWebhookConfiguration) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchAdmissionregistrationV1ValidatingWebhookConfigurationOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchAdmissionregistrationV1ValidatingWebhookConfigurationUnauthorizedCode is the HTTP code returned for type PatchAdmissionregistrationV1ValidatingWebhookConfigurationUnauthorized
const PatchAdmissionregistrationV1ValidatingWebhookConfigurationUnauthorizedCode int = 401

/*PatchAdmissionregistrationV1ValidatingWebhookConfigurationUnauthorized Unauthorized

swagger:response patchAdmissionregistrationV1ValidatingWebhookConfigurationUnauthorized
*/
type PatchAdmissionregistrationV1ValidatingWebhookConfigurationUnauthorized struct {
}

// NewPatchAdmissionregistrationV1ValidatingWebhookConfigurationUnauthorized creates PatchAdmissionregistrationV1ValidatingWebhookConfigurationUnauthorized with default headers values
func NewPatchAdmissionregistrationV1ValidatingWebhookConfigurationUnauthorized() *PatchAdmissionregistrationV1ValidatingWebhookConfigurationUnauthorized {

	return &PatchAdmissionregistrationV1ValidatingWebhookConfigurationUnauthorized{}
}

// WriteResponse to the client
func (o *PatchAdmissionregistrationV1ValidatingWebhookConfigurationUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
