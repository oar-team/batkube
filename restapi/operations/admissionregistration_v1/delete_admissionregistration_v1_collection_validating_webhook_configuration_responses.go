// Code generated by go-swagger; DO NOT EDIT.

package admissionregistration_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// DeleteAdmissionregistrationV1CollectionValidatingWebhookConfigurationOKCode is the HTTP code returned for type DeleteAdmissionregistrationV1CollectionValidatingWebhookConfigurationOK
const DeleteAdmissionregistrationV1CollectionValidatingWebhookConfigurationOKCode int = 200

/*DeleteAdmissionregistrationV1CollectionValidatingWebhookConfigurationOK OK

swagger:response deleteAdmissionregistrationV1CollectionValidatingWebhookConfigurationOK
*/
type DeleteAdmissionregistrationV1CollectionValidatingWebhookConfigurationOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1Status `json:"body,omitempty"`
}

// NewDeleteAdmissionregistrationV1CollectionValidatingWebhookConfigurationOK creates DeleteAdmissionregistrationV1CollectionValidatingWebhookConfigurationOK with default headers values
func NewDeleteAdmissionregistrationV1CollectionValidatingWebhookConfigurationOK() *DeleteAdmissionregistrationV1CollectionValidatingWebhookConfigurationOK {

	return &DeleteAdmissionregistrationV1CollectionValidatingWebhookConfigurationOK{}
}

// WithPayload adds the payload to the delete admissionregistration v1 collection validating webhook configuration o k response
func (o *DeleteAdmissionregistrationV1CollectionValidatingWebhookConfigurationOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) *DeleteAdmissionregistrationV1CollectionValidatingWebhookConfigurationOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete admissionregistration v1 collection validating webhook configuration o k response
func (o *DeleteAdmissionregistrationV1CollectionValidatingWebhookConfigurationOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteAdmissionregistrationV1CollectionValidatingWebhookConfigurationOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteAdmissionregistrationV1CollectionValidatingWebhookConfigurationUnauthorizedCode is the HTTP code returned for type DeleteAdmissionregistrationV1CollectionValidatingWebhookConfigurationUnauthorized
const DeleteAdmissionregistrationV1CollectionValidatingWebhookConfigurationUnauthorizedCode int = 401

/*DeleteAdmissionregistrationV1CollectionValidatingWebhookConfigurationUnauthorized Unauthorized

swagger:response deleteAdmissionregistrationV1CollectionValidatingWebhookConfigurationUnauthorized
*/
type DeleteAdmissionregistrationV1CollectionValidatingWebhookConfigurationUnauthorized struct {
}

// NewDeleteAdmissionregistrationV1CollectionValidatingWebhookConfigurationUnauthorized creates DeleteAdmissionregistrationV1CollectionValidatingWebhookConfigurationUnauthorized with default headers values
func NewDeleteAdmissionregistrationV1CollectionValidatingWebhookConfigurationUnauthorized() *DeleteAdmissionregistrationV1CollectionValidatingWebhookConfigurationUnauthorized {

	return &DeleteAdmissionregistrationV1CollectionValidatingWebhookConfigurationUnauthorized{}
}

// WriteResponse to the client
func (o *DeleteAdmissionregistrationV1CollectionValidatingWebhookConfigurationUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
