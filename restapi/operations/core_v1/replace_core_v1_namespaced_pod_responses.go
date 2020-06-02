// Code generated by go-swagger; DO NOT EDIT.

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/models"
)

// ReplaceCoreV1NamespacedPodOKCode is the HTTP code returned for type ReplaceCoreV1NamespacedPodOK
const ReplaceCoreV1NamespacedPodOKCode int = 200

/*ReplaceCoreV1NamespacedPodOK OK

swagger:response replaceCoreV1NamespacedPodOK
*/
type ReplaceCoreV1NamespacedPodOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPICoreV1Pod `json:"body,omitempty"`
}

// NewReplaceCoreV1NamespacedPodOK creates ReplaceCoreV1NamespacedPodOK with default headers values
func NewReplaceCoreV1NamespacedPodOK() *ReplaceCoreV1NamespacedPodOK {

	return &ReplaceCoreV1NamespacedPodOK{}
}

// WithPayload adds the payload to the replace core v1 namespaced pod o k response
func (o *ReplaceCoreV1NamespacedPodOK) WithPayload(payload *models.IoK8sAPICoreV1Pod) *ReplaceCoreV1NamespacedPodOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace core v1 namespaced pod o k response
func (o *ReplaceCoreV1NamespacedPodOK) SetPayload(payload *models.IoK8sAPICoreV1Pod) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceCoreV1NamespacedPodOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceCoreV1NamespacedPodCreatedCode is the HTTP code returned for type ReplaceCoreV1NamespacedPodCreated
const ReplaceCoreV1NamespacedPodCreatedCode int = 201

/*ReplaceCoreV1NamespacedPodCreated Created

swagger:response replaceCoreV1NamespacedPodCreated
*/
type ReplaceCoreV1NamespacedPodCreated struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPICoreV1Pod `json:"body,omitempty"`
}

// NewReplaceCoreV1NamespacedPodCreated creates ReplaceCoreV1NamespacedPodCreated with default headers values
func NewReplaceCoreV1NamespacedPodCreated() *ReplaceCoreV1NamespacedPodCreated {

	return &ReplaceCoreV1NamespacedPodCreated{}
}

// WithPayload adds the payload to the replace core v1 namespaced pod created response
func (o *ReplaceCoreV1NamespacedPodCreated) WithPayload(payload *models.IoK8sAPICoreV1Pod) *ReplaceCoreV1NamespacedPodCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace core v1 namespaced pod created response
func (o *ReplaceCoreV1NamespacedPodCreated) SetPayload(payload *models.IoK8sAPICoreV1Pod) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceCoreV1NamespacedPodCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceCoreV1NamespacedPodUnauthorizedCode is the HTTP code returned for type ReplaceCoreV1NamespacedPodUnauthorized
const ReplaceCoreV1NamespacedPodUnauthorizedCode int = 401

/*ReplaceCoreV1NamespacedPodUnauthorized Unauthorized

swagger:response replaceCoreV1NamespacedPodUnauthorized
*/
type ReplaceCoreV1NamespacedPodUnauthorized struct {
}

// NewReplaceCoreV1NamespacedPodUnauthorized creates ReplaceCoreV1NamespacedPodUnauthorized with default headers values
func NewReplaceCoreV1NamespacedPodUnauthorized() *ReplaceCoreV1NamespacedPodUnauthorized {

	return &ReplaceCoreV1NamespacedPodUnauthorized{}
}

// WriteResponse to the client
func (o *ReplaceCoreV1NamespacedPodUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
