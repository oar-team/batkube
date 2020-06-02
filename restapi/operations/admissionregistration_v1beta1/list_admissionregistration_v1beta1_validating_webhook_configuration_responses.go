// Code generated by go-swagger; DO NOT EDIT.

package admissionregistration_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/models"
)

// ListAdmissionregistrationV1beta1ValidatingWebhookConfigurationOKCode is the HTTP code returned for type ListAdmissionregistrationV1beta1ValidatingWebhookConfigurationOK
const ListAdmissionregistrationV1beta1ValidatingWebhookConfigurationOKCode int = 200

/*ListAdmissionregistrationV1beta1ValidatingWebhookConfigurationOK OK

swagger:response listAdmissionregistrationV1beta1ValidatingWebhookConfigurationOK
*/
type ListAdmissionregistrationV1beta1ValidatingWebhookConfigurationOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIAdmissionregistrationV1beta1ValidatingWebhookConfigurationList `json:"body,omitempty"`
}

// NewListAdmissionregistrationV1beta1ValidatingWebhookConfigurationOK creates ListAdmissionregistrationV1beta1ValidatingWebhookConfigurationOK with default headers values
func NewListAdmissionregistrationV1beta1ValidatingWebhookConfigurationOK() *ListAdmissionregistrationV1beta1ValidatingWebhookConfigurationOK {

	return &ListAdmissionregistrationV1beta1ValidatingWebhookConfigurationOK{}
}

// WithPayload adds the payload to the list admissionregistration v1beta1 validating webhook configuration o k response
func (o *ListAdmissionregistrationV1beta1ValidatingWebhookConfigurationOK) WithPayload(payload *models.IoK8sAPIAdmissionregistrationV1beta1ValidatingWebhookConfigurationList) *ListAdmissionregistrationV1beta1ValidatingWebhookConfigurationOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list admissionregistration v1beta1 validating webhook configuration o k response
func (o *ListAdmissionregistrationV1beta1ValidatingWebhookConfigurationOK) SetPayload(payload *models.IoK8sAPIAdmissionregistrationV1beta1ValidatingWebhookConfigurationList) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListAdmissionregistrationV1beta1ValidatingWebhookConfigurationOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListAdmissionregistrationV1beta1ValidatingWebhookConfigurationUnauthorizedCode is the HTTP code returned for type ListAdmissionregistrationV1beta1ValidatingWebhookConfigurationUnauthorized
const ListAdmissionregistrationV1beta1ValidatingWebhookConfigurationUnauthorizedCode int = 401

/*ListAdmissionregistrationV1beta1ValidatingWebhookConfigurationUnauthorized Unauthorized

swagger:response listAdmissionregistrationV1beta1ValidatingWebhookConfigurationUnauthorized
*/
type ListAdmissionregistrationV1beta1ValidatingWebhookConfigurationUnauthorized struct {
}

// NewListAdmissionregistrationV1beta1ValidatingWebhookConfigurationUnauthorized creates ListAdmissionregistrationV1beta1ValidatingWebhookConfigurationUnauthorized with default headers values
func NewListAdmissionregistrationV1beta1ValidatingWebhookConfigurationUnauthorized() *ListAdmissionregistrationV1beta1ValidatingWebhookConfigurationUnauthorized {

	return &ListAdmissionregistrationV1beta1ValidatingWebhookConfigurationUnauthorized{}
}

// WriteResponse to the client
func (o *ListAdmissionregistrationV1beta1ValidatingWebhookConfigurationUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
