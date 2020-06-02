// Code generated by go-swagger; DO NOT EDIT.

package admissionregistration_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/models"
)

// WatchAdmissionregistrationV1beta1MutatingWebhookConfigurationOKCode is the HTTP code returned for type WatchAdmissionregistrationV1beta1MutatingWebhookConfigurationOK
const WatchAdmissionregistrationV1beta1MutatingWebhookConfigurationOKCode int = 200

/*WatchAdmissionregistrationV1beta1MutatingWebhookConfigurationOK OK

swagger:response watchAdmissionregistrationV1beta1MutatingWebhookConfigurationOK
*/
type WatchAdmissionregistrationV1beta1MutatingWebhookConfigurationOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent `json:"body,omitempty"`
}

// NewWatchAdmissionregistrationV1beta1MutatingWebhookConfigurationOK creates WatchAdmissionregistrationV1beta1MutatingWebhookConfigurationOK with default headers values
func NewWatchAdmissionregistrationV1beta1MutatingWebhookConfigurationOK() *WatchAdmissionregistrationV1beta1MutatingWebhookConfigurationOK {

	return &WatchAdmissionregistrationV1beta1MutatingWebhookConfigurationOK{}
}

// WithPayload adds the payload to the watch admissionregistration v1beta1 mutating webhook configuration o k response
func (o *WatchAdmissionregistrationV1beta1MutatingWebhookConfigurationOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) *WatchAdmissionregistrationV1beta1MutatingWebhookConfigurationOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the watch admissionregistration v1beta1 mutating webhook configuration o k response
func (o *WatchAdmissionregistrationV1beta1MutatingWebhookConfigurationOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WatchAdmissionregistrationV1beta1MutatingWebhookConfigurationOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WatchAdmissionregistrationV1beta1MutatingWebhookConfigurationUnauthorizedCode is the HTTP code returned for type WatchAdmissionregistrationV1beta1MutatingWebhookConfigurationUnauthorized
const WatchAdmissionregistrationV1beta1MutatingWebhookConfigurationUnauthorizedCode int = 401

/*WatchAdmissionregistrationV1beta1MutatingWebhookConfigurationUnauthorized Unauthorized

swagger:response watchAdmissionregistrationV1beta1MutatingWebhookConfigurationUnauthorized
*/
type WatchAdmissionregistrationV1beta1MutatingWebhookConfigurationUnauthorized struct {
}

// NewWatchAdmissionregistrationV1beta1MutatingWebhookConfigurationUnauthorized creates WatchAdmissionregistrationV1beta1MutatingWebhookConfigurationUnauthorized with default headers values
func NewWatchAdmissionregistrationV1beta1MutatingWebhookConfigurationUnauthorized() *WatchAdmissionregistrationV1beta1MutatingWebhookConfigurationUnauthorized {

	return &WatchAdmissionregistrationV1beta1MutatingWebhookConfigurationUnauthorized{}
}

// WriteResponse to the client
func (o *WatchAdmissionregistrationV1beta1MutatingWebhookConfigurationUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
