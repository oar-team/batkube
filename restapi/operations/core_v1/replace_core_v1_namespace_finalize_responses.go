// Code generated by go-swagger; DO NOT EDIT.

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/models"
)

// ReplaceCoreV1NamespaceFinalizeOKCode is the HTTP code returned for type ReplaceCoreV1NamespaceFinalizeOK
const ReplaceCoreV1NamespaceFinalizeOKCode int = 200

/*ReplaceCoreV1NamespaceFinalizeOK OK

swagger:response replaceCoreV1NamespaceFinalizeOK
*/
type ReplaceCoreV1NamespaceFinalizeOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPICoreV1Namespace `json:"body,omitempty"`
}

// NewReplaceCoreV1NamespaceFinalizeOK creates ReplaceCoreV1NamespaceFinalizeOK with default headers values
func NewReplaceCoreV1NamespaceFinalizeOK() *ReplaceCoreV1NamespaceFinalizeOK {

	return &ReplaceCoreV1NamespaceFinalizeOK{}
}

// WithPayload adds the payload to the replace core v1 namespace finalize o k response
func (o *ReplaceCoreV1NamespaceFinalizeOK) WithPayload(payload *models.IoK8sAPICoreV1Namespace) *ReplaceCoreV1NamespaceFinalizeOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace core v1 namespace finalize o k response
func (o *ReplaceCoreV1NamespaceFinalizeOK) SetPayload(payload *models.IoK8sAPICoreV1Namespace) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceCoreV1NamespaceFinalizeOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceCoreV1NamespaceFinalizeCreatedCode is the HTTP code returned for type ReplaceCoreV1NamespaceFinalizeCreated
const ReplaceCoreV1NamespaceFinalizeCreatedCode int = 201

/*ReplaceCoreV1NamespaceFinalizeCreated Created

swagger:response replaceCoreV1NamespaceFinalizeCreated
*/
type ReplaceCoreV1NamespaceFinalizeCreated struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPICoreV1Namespace `json:"body,omitempty"`
}

// NewReplaceCoreV1NamespaceFinalizeCreated creates ReplaceCoreV1NamespaceFinalizeCreated with default headers values
func NewReplaceCoreV1NamespaceFinalizeCreated() *ReplaceCoreV1NamespaceFinalizeCreated {

	return &ReplaceCoreV1NamespaceFinalizeCreated{}
}

// WithPayload adds the payload to the replace core v1 namespace finalize created response
func (o *ReplaceCoreV1NamespaceFinalizeCreated) WithPayload(payload *models.IoK8sAPICoreV1Namespace) *ReplaceCoreV1NamespaceFinalizeCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace core v1 namespace finalize created response
func (o *ReplaceCoreV1NamespaceFinalizeCreated) SetPayload(payload *models.IoK8sAPICoreV1Namespace) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceCoreV1NamespaceFinalizeCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceCoreV1NamespaceFinalizeUnauthorizedCode is the HTTP code returned for type ReplaceCoreV1NamespaceFinalizeUnauthorized
const ReplaceCoreV1NamespaceFinalizeUnauthorizedCode int = 401

/*ReplaceCoreV1NamespaceFinalizeUnauthorized Unauthorized

swagger:response replaceCoreV1NamespaceFinalizeUnauthorized
*/
type ReplaceCoreV1NamespaceFinalizeUnauthorized struct {
}

// NewReplaceCoreV1NamespaceFinalizeUnauthorized creates ReplaceCoreV1NamespaceFinalizeUnauthorized with default headers values
func NewReplaceCoreV1NamespaceFinalizeUnauthorized() *ReplaceCoreV1NamespaceFinalizeUnauthorized {

	return &ReplaceCoreV1NamespaceFinalizeUnauthorized{}
}

// WriteResponse to the client
func (o *ReplaceCoreV1NamespaceFinalizeUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
