// Code generated by go-swagger; DO NOT EDIT.

package admissionregistration_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/models"
)

// ListAdmissionregistrationV1beta1MutatingWebhookConfigurationOKCode is the HTTP code returned for type ListAdmissionregistrationV1beta1MutatingWebhookConfigurationOK
const ListAdmissionregistrationV1beta1MutatingWebhookConfigurationOKCode int = 200

/*ListAdmissionregistrationV1beta1MutatingWebhookConfigurationOK OK

swagger:response listAdmissionregistrationV1beta1MutatingWebhookConfigurationOK
*/
type ListAdmissionregistrationV1beta1MutatingWebhookConfigurationOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIAdmissionregistrationV1beta1MutatingWebhookConfigurationList `json:"body,omitempty"`
}

// NewListAdmissionregistrationV1beta1MutatingWebhookConfigurationOK creates ListAdmissionregistrationV1beta1MutatingWebhookConfigurationOK with default headers values
func NewListAdmissionregistrationV1beta1MutatingWebhookConfigurationOK() *ListAdmissionregistrationV1beta1MutatingWebhookConfigurationOK {

	return &ListAdmissionregistrationV1beta1MutatingWebhookConfigurationOK{}
}

// WithPayload adds the payload to the list admissionregistration v1beta1 mutating webhook configuration o k response
func (o *ListAdmissionregistrationV1beta1MutatingWebhookConfigurationOK) WithPayload(payload *models.IoK8sAPIAdmissionregistrationV1beta1MutatingWebhookConfigurationList) *ListAdmissionregistrationV1beta1MutatingWebhookConfigurationOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list admissionregistration v1beta1 mutating webhook configuration o k response
func (o *ListAdmissionregistrationV1beta1MutatingWebhookConfigurationOK) SetPayload(payload *models.IoK8sAPIAdmissionregistrationV1beta1MutatingWebhookConfigurationList) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListAdmissionregistrationV1beta1MutatingWebhookConfigurationOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListAdmissionregistrationV1beta1MutatingWebhookConfigurationUnauthorizedCode is the HTTP code returned for type ListAdmissionregistrationV1beta1MutatingWebhookConfigurationUnauthorized
const ListAdmissionregistrationV1beta1MutatingWebhookConfigurationUnauthorizedCode int = 401

/*ListAdmissionregistrationV1beta1MutatingWebhookConfigurationUnauthorized Unauthorized

swagger:response listAdmissionregistrationV1beta1MutatingWebhookConfigurationUnauthorized
*/
type ListAdmissionregistrationV1beta1MutatingWebhookConfigurationUnauthorized struct {
}

// NewListAdmissionregistrationV1beta1MutatingWebhookConfigurationUnauthorized creates ListAdmissionregistrationV1beta1MutatingWebhookConfigurationUnauthorized with default headers values
func NewListAdmissionregistrationV1beta1MutatingWebhookConfigurationUnauthorized() *ListAdmissionregistrationV1beta1MutatingWebhookConfigurationUnauthorized {

	return &ListAdmissionregistrationV1beta1MutatingWebhookConfigurationUnauthorized{}
}

// WriteResponse to the client
func (o *ListAdmissionregistrationV1beta1MutatingWebhookConfigurationUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}