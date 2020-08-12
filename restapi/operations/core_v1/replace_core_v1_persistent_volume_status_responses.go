// Code generated by go-swagger; DO NOT EDIT.

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// ReplaceCoreV1PersistentVolumeStatusOKCode is the HTTP code returned for type ReplaceCoreV1PersistentVolumeStatusOK
const ReplaceCoreV1PersistentVolumeStatusOKCode int = 200

/*ReplaceCoreV1PersistentVolumeStatusOK OK

swagger:response replaceCoreV1PersistentVolumeStatusOK
*/
type ReplaceCoreV1PersistentVolumeStatusOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPICoreV1PersistentVolume `json:"body,omitempty"`
}

// NewReplaceCoreV1PersistentVolumeStatusOK creates ReplaceCoreV1PersistentVolumeStatusOK with default headers values
func NewReplaceCoreV1PersistentVolumeStatusOK() *ReplaceCoreV1PersistentVolumeStatusOK {

	return &ReplaceCoreV1PersistentVolumeStatusOK{}
}

// WithPayload adds the payload to the replace core v1 persistent volume status o k response
func (o *ReplaceCoreV1PersistentVolumeStatusOK) WithPayload(payload *models.IoK8sAPICoreV1PersistentVolume) *ReplaceCoreV1PersistentVolumeStatusOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace core v1 persistent volume status o k response
func (o *ReplaceCoreV1PersistentVolumeStatusOK) SetPayload(payload *models.IoK8sAPICoreV1PersistentVolume) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceCoreV1PersistentVolumeStatusOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceCoreV1PersistentVolumeStatusCreatedCode is the HTTP code returned for type ReplaceCoreV1PersistentVolumeStatusCreated
const ReplaceCoreV1PersistentVolumeStatusCreatedCode int = 201

/*ReplaceCoreV1PersistentVolumeStatusCreated Created

swagger:response replaceCoreV1PersistentVolumeStatusCreated
*/
type ReplaceCoreV1PersistentVolumeStatusCreated struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPICoreV1PersistentVolume `json:"body,omitempty"`
}

// NewReplaceCoreV1PersistentVolumeStatusCreated creates ReplaceCoreV1PersistentVolumeStatusCreated with default headers values
func NewReplaceCoreV1PersistentVolumeStatusCreated() *ReplaceCoreV1PersistentVolumeStatusCreated {

	return &ReplaceCoreV1PersistentVolumeStatusCreated{}
}

// WithPayload adds the payload to the replace core v1 persistent volume status created response
func (o *ReplaceCoreV1PersistentVolumeStatusCreated) WithPayload(payload *models.IoK8sAPICoreV1PersistentVolume) *ReplaceCoreV1PersistentVolumeStatusCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace core v1 persistent volume status created response
func (o *ReplaceCoreV1PersistentVolumeStatusCreated) SetPayload(payload *models.IoK8sAPICoreV1PersistentVolume) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceCoreV1PersistentVolumeStatusCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceCoreV1PersistentVolumeStatusUnauthorizedCode is the HTTP code returned for type ReplaceCoreV1PersistentVolumeStatusUnauthorized
const ReplaceCoreV1PersistentVolumeStatusUnauthorizedCode int = 401

/*ReplaceCoreV1PersistentVolumeStatusUnauthorized Unauthorized

swagger:response replaceCoreV1PersistentVolumeStatusUnauthorized
*/
type ReplaceCoreV1PersistentVolumeStatusUnauthorized struct {
}

// NewReplaceCoreV1PersistentVolumeStatusUnauthorized creates ReplaceCoreV1PersistentVolumeStatusUnauthorized with default headers values
func NewReplaceCoreV1PersistentVolumeStatusUnauthorized() *ReplaceCoreV1PersistentVolumeStatusUnauthorized {

	return &ReplaceCoreV1PersistentVolumeStatusUnauthorized{}
}

// WriteResponse to the client
func (o *ReplaceCoreV1PersistentVolumeStatusUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
