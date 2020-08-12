// Code generated by go-swagger; DO NOT EDIT.

package admissionregistration_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// ListAdmissionregistrationV1MutatingWebhookConfigurationOKCode is the HTTP code returned for type ListAdmissionregistrationV1MutatingWebhookConfigurationOK
const ListAdmissionregistrationV1MutatingWebhookConfigurationOKCode int = 200

/*ListAdmissionregistrationV1MutatingWebhookConfigurationOK OK

swagger:response listAdmissionregistrationV1MutatingWebhookConfigurationOK
*/
type ListAdmissionregistrationV1MutatingWebhookConfigurationOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIAdmissionregistrationV1MutatingWebhookConfigurationList `json:"body,omitempty"`
}

// NewListAdmissionregistrationV1MutatingWebhookConfigurationOK creates ListAdmissionregistrationV1MutatingWebhookConfigurationOK with default headers values
func NewListAdmissionregistrationV1MutatingWebhookConfigurationOK() *ListAdmissionregistrationV1MutatingWebhookConfigurationOK {

	return &ListAdmissionregistrationV1MutatingWebhookConfigurationOK{}
}

// WithPayload adds the payload to the list admissionregistration v1 mutating webhook configuration o k response
func (o *ListAdmissionregistrationV1MutatingWebhookConfigurationOK) WithPayload(payload *models.IoK8sAPIAdmissionregistrationV1MutatingWebhookConfigurationList) *ListAdmissionregistrationV1MutatingWebhookConfigurationOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list admissionregistration v1 mutating webhook configuration o k response
func (o *ListAdmissionregistrationV1MutatingWebhookConfigurationOK) SetPayload(payload *models.IoK8sAPIAdmissionregistrationV1MutatingWebhookConfigurationList) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListAdmissionregistrationV1MutatingWebhookConfigurationOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListAdmissionregistrationV1MutatingWebhookConfigurationUnauthorizedCode is the HTTP code returned for type ListAdmissionregistrationV1MutatingWebhookConfigurationUnauthorized
const ListAdmissionregistrationV1MutatingWebhookConfigurationUnauthorizedCode int = 401

/*ListAdmissionregistrationV1MutatingWebhookConfigurationUnauthorized Unauthorized

swagger:response listAdmissionregistrationV1MutatingWebhookConfigurationUnauthorized
*/
type ListAdmissionregistrationV1MutatingWebhookConfigurationUnauthorized struct {
}

// NewListAdmissionregistrationV1MutatingWebhookConfigurationUnauthorized creates ListAdmissionregistrationV1MutatingWebhookConfigurationUnauthorized with default headers values
func NewListAdmissionregistrationV1MutatingWebhookConfigurationUnauthorized() *ListAdmissionregistrationV1MutatingWebhookConfigurationUnauthorized {

	return &ListAdmissionregistrationV1MutatingWebhookConfigurationUnauthorized{}
}

// WriteResponse to the client
func (o *ListAdmissionregistrationV1MutatingWebhookConfigurationUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
