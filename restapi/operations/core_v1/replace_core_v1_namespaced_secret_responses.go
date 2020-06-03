// Code generated by go-swagger; DO NOT EDIT.

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/models"
)

// ReplaceCoreV1NamespacedSecretOKCode is the HTTP code returned for type ReplaceCoreV1NamespacedSecretOK
const ReplaceCoreV1NamespacedSecretOKCode int = 200

/*ReplaceCoreV1NamespacedSecretOK OK

swagger:response replaceCoreV1NamespacedSecretOK
*/
type ReplaceCoreV1NamespacedSecretOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPICoreV1Secret `json:"body,omitempty"`
}

// NewReplaceCoreV1NamespacedSecretOK creates ReplaceCoreV1NamespacedSecretOK with default headers values
func NewReplaceCoreV1NamespacedSecretOK() *ReplaceCoreV1NamespacedSecretOK {

	return &ReplaceCoreV1NamespacedSecretOK{}
}

// WithPayload adds the payload to the replace core v1 namespaced secret o k response
func (o *ReplaceCoreV1NamespacedSecretOK) WithPayload(payload *models.IoK8sAPICoreV1Secret) *ReplaceCoreV1NamespacedSecretOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace core v1 namespaced secret o k response
func (o *ReplaceCoreV1NamespacedSecretOK) SetPayload(payload *models.IoK8sAPICoreV1Secret) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceCoreV1NamespacedSecretOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceCoreV1NamespacedSecretCreatedCode is the HTTP code returned for type ReplaceCoreV1NamespacedSecretCreated
const ReplaceCoreV1NamespacedSecretCreatedCode int = 201

/*ReplaceCoreV1NamespacedSecretCreated Created

swagger:response replaceCoreV1NamespacedSecretCreated
*/
type ReplaceCoreV1NamespacedSecretCreated struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPICoreV1Secret `json:"body,omitempty"`
}

// NewReplaceCoreV1NamespacedSecretCreated creates ReplaceCoreV1NamespacedSecretCreated with default headers values
func NewReplaceCoreV1NamespacedSecretCreated() *ReplaceCoreV1NamespacedSecretCreated {

	return &ReplaceCoreV1NamespacedSecretCreated{}
}

// WithPayload adds the payload to the replace core v1 namespaced secret created response
func (o *ReplaceCoreV1NamespacedSecretCreated) WithPayload(payload *models.IoK8sAPICoreV1Secret) *ReplaceCoreV1NamespacedSecretCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace core v1 namespaced secret created response
func (o *ReplaceCoreV1NamespacedSecretCreated) SetPayload(payload *models.IoK8sAPICoreV1Secret) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceCoreV1NamespacedSecretCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceCoreV1NamespacedSecretUnauthorizedCode is the HTTP code returned for type ReplaceCoreV1NamespacedSecretUnauthorized
const ReplaceCoreV1NamespacedSecretUnauthorizedCode int = 401

/*ReplaceCoreV1NamespacedSecretUnauthorized Unauthorized

swagger:response replaceCoreV1NamespacedSecretUnauthorized
*/
type ReplaceCoreV1NamespacedSecretUnauthorized struct {
}

// NewReplaceCoreV1NamespacedSecretUnauthorized creates ReplaceCoreV1NamespacedSecretUnauthorized with default headers values
func NewReplaceCoreV1NamespacedSecretUnauthorized() *ReplaceCoreV1NamespacedSecretUnauthorized {

	return &ReplaceCoreV1NamespacedSecretUnauthorized{}
}

// WriteResponse to the client
func (o *ReplaceCoreV1NamespacedSecretUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}