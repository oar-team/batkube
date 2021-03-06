// Code generated by go-swagger; DO NOT EDIT.

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// ReplaceCoreV1NamespacedResourceQuotaStatusOKCode is the HTTP code returned for type ReplaceCoreV1NamespacedResourceQuotaStatusOK
const ReplaceCoreV1NamespacedResourceQuotaStatusOKCode int = 200

/*ReplaceCoreV1NamespacedResourceQuotaStatusOK OK

swagger:response replaceCoreV1NamespacedResourceQuotaStatusOK
*/
type ReplaceCoreV1NamespacedResourceQuotaStatusOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPICoreV1ResourceQuota `json:"body,omitempty"`
}

// NewReplaceCoreV1NamespacedResourceQuotaStatusOK creates ReplaceCoreV1NamespacedResourceQuotaStatusOK with default headers values
func NewReplaceCoreV1NamespacedResourceQuotaStatusOK() *ReplaceCoreV1NamespacedResourceQuotaStatusOK {

	return &ReplaceCoreV1NamespacedResourceQuotaStatusOK{}
}

// WithPayload adds the payload to the replace core v1 namespaced resource quota status o k response
func (o *ReplaceCoreV1NamespacedResourceQuotaStatusOK) WithPayload(payload *models.IoK8sAPICoreV1ResourceQuota) *ReplaceCoreV1NamespacedResourceQuotaStatusOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace core v1 namespaced resource quota status o k response
func (o *ReplaceCoreV1NamespacedResourceQuotaStatusOK) SetPayload(payload *models.IoK8sAPICoreV1ResourceQuota) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceCoreV1NamespacedResourceQuotaStatusOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceCoreV1NamespacedResourceQuotaStatusCreatedCode is the HTTP code returned for type ReplaceCoreV1NamespacedResourceQuotaStatusCreated
const ReplaceCoreV1NamespacedResourceQuotaStatusCreatedCode int = 201

/*ReplaceCoreV1NamespacedResourceQuotaStatusCreated Created

swagger:response replaceCoreV1NamespacedResourceQuotaStatusCreated
*/
type ReplaceCoreV1NamespacedResourceQuotaStatusCreated struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPICoreV1ResourceQuota `json:"body,omitempty"`
}

// NewReplaceCoreV1NamespacedResourceQuotaStatusCreated creates ReplaceCoreV1NamespacedResourceQuotaStatusCreated with default headers values
func NewReplaceCoreV1NamespacedResourceQuotaStatusCreated() *ReplaceCoreV1NamespacedResourceQuotaStatusCreated {

	return &ReplaceCoreV1NamespacedResourceQuotaStatusCreated{}
}

// WithPayload adds the payload to the replace core v1 namespaced resource quota status created response
func (o *ReplaceCoreV1NamespacedResourceQuotaStatusCreated) WithPayload(payload *models.IoK8sAPICoreV1ResourceQuota) *ReplaceCoreV1NamespacedResourceQuotaStatusCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace core v1 namespaced resource quota status created response
func (o *ReplaceCoreV1NamespacedResourceQuotaStatusCreated) SetPayload(payload *models.IoK8sAPICoreV1ResourceQuota) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceCoreV1NamespacedResourceQuotaStatusCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceCoreV1NamespacedResourceQuotaStatusUnauthorizedCode is the HTTP code returned for type ReplaceCoreV1NamespacedResourceQuotaStatusUnauthorized
const ReplaceCoreV1NamespacedResourceQuotaStatusUnauthorizedCode int = 401

/*ReplaceCoreV1NamespacedResourceQuotaStatusUnauthorized Unauthorized

swagger:response replaceCoreV1NamespacedResourceQuotaStatusUnauthorized
*/
type ReplaceCoreV1NamespacedResourceQuotaStatusUnauthorized struct {
}

// NewReplaceCoreV1NamespacedResourceQuotaStatusUnauthorized creates ReplaceCoreV1NamespacedResourceQuotaStatusUnauthorized with default headers values
func NewReplaceCoreV1NamespacedResourceQuotaStatusUnauthorized() *ReplaceCoreV1NamespacedResourceQuotaStatusUnauthorized {

	return &ReplaceCoreV1NamespacedResourceQuotaStatusUnauthorized{}
}

// WriteResponse to the client
func (o *ReplaceCoreV1NamespacedResourceQuotaStatusUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
