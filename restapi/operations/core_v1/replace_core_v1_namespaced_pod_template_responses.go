// Code generated by go-swagger; DO NOT EDIT.

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/models"
)

// ReplaceCoreV1NamespacedPodTemplateOKCode is the HTTP code returned for type ReplaceCoreV1NamespacedPodTemplateOK
const ReplaceCoreV1NamespacedPodTemplateOKCode int = 200

/*ReplaceCoreV1NamespacedPodTemplateOK OK

swagger:response replaceCoreV1NamespacedPodTemplateOK
*/
type ReplaceCoreV1NamespacedPodTemplateOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPICoreV1PodTemplate `json:"body,omitempty"`
}

// NewReplaceCoreV1NamespacedPodTemplateOK creates ReplaceCoreV1NamespacedPodTemplateOK with default headers values
func NewReplaceCoreV1NamespacedPodTemplateOK() *ReplaceCoreV1NamespacedPodTemplateOK {

	return &ReplaceCoreV1NamespacedPodTemplateOK{}
}

// WithPayload adds the payload to the replace core v1 namespaced pod template o k response
func (o *ReplaceCoreV1NamespacedPodTemplateOK) WithPayload(payload *models.IoK8sAPICoreV1PodTemplate) *ReplaceCoreV1NamespacedPodTemplateOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace core v1 namespaced pod template o k response
func (o *ReplaceCoreV1NamespacedPodTemplateOK) SetPayload(payload *models.IoK8sAPICoreV1PodTemplate) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceCoreV1NamespacedPodTemplateOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceCoreV1NamespacedPodTemplateCreatedCode is the HTTP code returned for type ReplaceCoreV1NamespacedPodTemplateCreated
const ReplaceCoreV1NamespacedPodTemplateCreatedCode int = 201

/*ReplaceCoreV1NamespacedPodTemplateCreated Created

swagger:response replaceCoreV1NamespacedPodTemplateCreated
*/
type ReplaceCoreV1NamespacedPodTemplateCreated struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPICoreV1PodTemplate `json:"body,omitempty"`
}

// NewReplaceCoreV1NamespacedPodTemplateCreated creates ReplaceCoreV1NamespacedPodTemplateCreated with default headers values
func NewReplaceCoreV1NamespacedPodTemplateCreated() *ReplaceCoreV1NamespacedPodTemplateCreated {

	return &ReplaceCoreV1NamespacedPodTemplateCreated{}
}

// WithPayload adds the payload to the replace core v1 namespaced pod template created response
func (o *ReplaceCoreV1NamespacedPodTemplateCreated) WithPayload(payload *models.IoK8sAPICoreV1PodTemplate) *ReplaceCoreV1NamespacedPodTemplateCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace core v1 namespaced pod template created response
func (o *ReplaceCoreV1NamespacedPodTemplateCreated) SetPayload(payload *models.IoK8sAPICoreV1PodTemplate) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceCoreV1NamespacedPodTemplateCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceCoreV1NamespacedPodTemplateUnauthorizedCode is the HTTP code returned for type ReplaceCoreV1NamespacedPodTemplateUnauthorized
const ReplaceCoreV1NamespacedPodTemplateUnauthorizedCode int = 401

/*ReplaceCoreV1NamespacedPodTemplateUnauthorized Unauthorized

swagger:response replaceCoreV1NamespacedPodTemplateUnauthorized
*/
type ReplaceCoreV1NamespacedPodTemplateUnauthorized struct {
}

// NewReplaceCoreV1NamespacedPodTemplateUnauthorized creates ReplaceCoreV1NamespacedPodTemplateUnauthorized with default headers values
func NewReplaceCoreV1NamespacedPodTemplateUnauthorized() *ReplaceCoreV1NamespacedPodTemplateUnauthorized {

	return &ReplaceCoreV1NamespacedPodTemplateUnauthorized{}
}

// WriteResponse to the client
func (o *ReplaceCoreV1NamespacedPodTemplateUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}