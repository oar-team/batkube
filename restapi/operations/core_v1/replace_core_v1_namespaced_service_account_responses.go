// Code generated by go-swagger; DO NOT EDIT.

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// ReplaceCoreV1NamespacedServiceAccountOKCode is the HTTP code returned for type ReplaceCoreV1NamespacedServiceAccountOK
const ReplaceCoreV1NamespacedServiceAccountOKCode int = 200

/*ReplaceCoreV1NamespacedServiceAccountOK OK

swagger:response replaceCoreV1NamespacedServiceAccountOK
*/
type ReplaceCoreV1NamespacedServiceAccountOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPICoreV1ServiceAccount `json:"body,omitempty"`
}

// NewReplaceCoreV1NamespacedServiceAccountOK creates ReplaceCoreV1NamespacedServiceAccountOK with default headers values
func NewReplaceCoreV1NamespacedServiceAccountOK() *ReplaceCoreV1NamespacedServiceAccountOK {

	return &ReplaceCoreV1NamespacedServiceAccountOK{}
}

// WithPayload adds the payload to the replace core v1 namespaced service account o k response
func (o *ReplaceCoreV1NamespacedServiceAccountOK) WithPayload(payload *models.IoK8sAPICoreV1ServiceAccount) *ReplaceCoreV1NamespacedServiceAccountOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace core v1 namespaced service account o k response
func (o *ReplaceCoreV1NamespacedServiceAccountOK) SetPayload(payload *models.IoK8sAPICoreV1ServiceAccount) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceCoreV1NamespacedServiceAccountOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceCoreV1NamespacedServiceAccountCreatedCode is the HTTP code returned for type ReplaceCoreV1NamespacedServiceAccountCreated
const ReplaceCoreV1NamespacedServiceAccountCreatedCode int = 201

/*ReplaceCoreV1NamespacedServiceAccountCreated Created

swagger:response replaceCoreV1NamespacedServiceAccountCreated
*/
type ReplaceCoreV1NamespacedServiceAccountCreated struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPICoreV1ServiceAccount `json:"body,omitempty"`
}

// NewReplaceCoreV1NamespacedServiceAccountCreated creates ReplaceCoreV1NamespacedServiceAccountCreated with default headers values
func NewReplaceCoreV1NamespacedServiceAccountCreated() *ReplaceCoreV1NamespacedServiceAccountCreated {

	return &ReplaceCoreV1NamespacedServiceAccountCreated{}
}

// WithPayload adds the payload to the replace core v1 namespaced service account created response
func (o *ReplaceCoreV1NamespacedServiceAccountCreated) WithPayload(payload *models.IoK8sAPICoreV1ServiceAccount) *ReplaceCoreV1NamespacedServiceAccountCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace core v1 namespaced service account created response
func (o *ReplaceCoreV1NamespacedServiceAccountCreated) SetPayload(payload *models.IoK8sAPICoreV1ServiceAccount) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceCoreV1NamespacedServiceAccountCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceCoreV1NamespacedServiceAccountUnauthorizedCode is the HTTP code returned for type ReplaceCoreV1NamespacedServiceAccountUnauthorized
const ReplaceCoreV1NamespacedServiceAccountUnauthorizedCode int = 401

/*ReplaceCoreV1NamespacedServiceAccountUnauthorized Unauthorized

swagger:response replaceCoreV1NamespacedServiceAccountUnauthorized
*/
type ReplaceCoreV1NamespacedServiceAccountUnauthorized struct {
}

// NewReplaceCoreV1NamespacedServiceAccountUnauthorized creates ReplaceCoreV1NamespacedServiceAccountUnauthorized with default headers values
func NewReplaceCoreV1NamespacedServiceAccountUnauthorized() *ReplaceCoreV1NamespacedServiceAccountUnauthorized {

	return &ReplaceCoreV1NamespacedServiceAccountUnauthorized{}
}

// WriteResponse to the client
func (o *ReplaceCoreV1NamespacedServiceAccountUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
