// Code generated by go-swagger; DO NOT EDIT.

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// CreateCoreV1NamespacedPersistentVolumeClaimOKCode is the HTTP code returned for type CreateCoreV1NamespacedPersistentVolumeClaimOK
const CreateCoreV1NamespacedPersistentVolumeClaimOKCode int = 200

/*CreateCoreV1NamespacedPersistentVolumeClaimOK OK

swagger:response createCoreV1NamespacedPersistentVolumeClaimOK
*/
type CreateCoreV1NamespacedPersistentVolumeClaimOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPICoreV1PersistentVolumeClaim `json:"body,omitempty"`
}

// NewCreateCoreV1NamespacedPersistentVolumeClaimOK creates CreateCoreV1NamespacedPersistentVolumeClaimOK with default headers values
func NewCreateCoreV1NamespacedPersistentVolumeClaimOK() *CreateCoreV1NamespacedPersistentVolumeClaimOK {

	return &CreateCoreV1NamespacedPersistentVolumeClaimOK{}
}

// WithPayload adds the payload to the create core v1 namespaced persistent volume claim o k response
func (o *CreateCoreV1NamespacedPersistentVolumeClaimOK) WithPayload(payload *models.IoK8sAPICoreV1PersistentVolumeClaim) *CreateCoreV1NamespacedPersistentVolumeClaimOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create core v1 namespaced persistent volume claim o k response
func (o *CreateCoreV1NamespacedPersistentVolumeClaimOK) SetPayload(payload *models.IoK8sAPICoreV1PersistentVolumeClaim) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateCoreV1NamespacedPersistentVolumeClaimOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateCoreV1NamespacedPersistentVolumeClaimCreatedCode is the HTTP code returned for type CreateCoreV1NamespacedPersistentVolumeClaimCreated
const CreateCoreV1NamespacedPersistentVolumeClaimCreatedCode int = 201

/*CreateCoreV1NamespacedPersistentVolumeClaimCreated Created

swagger:response createCoreV1NamespacedPersistentVolumeClaimCreated
*/
type CreateCoreV1NamespacedPersistentVolumeClaimCreated struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPICoreV1PersistentVolumeClaim `json:"body,omitempty"`
}

// NewCreateCoreV1NamespacedPersistentVolumeClaimCreated creates CreateCoreV1NamespacedPersistentVolumeClaimCreated with default headers values
func NewCreateCoreV1NamespacedPersistentVolumeClaimCreated() *CreateCoreV1NamespacedPersistentVolumeClaimCreated {

	return &CreateCoreV1NamespacedPersistentVolumeClaimCreated{}
}

// WithPayload adds the payload to the create core v1 namespaced persistent volume claim created response
func (o *CreateCoreV1NamespacedPersistentVolumeClaimCreated) WithPayload(payload *models.IoK8sAPICoreV1PersistentVolumeClaim) *CreateCoreV1NamespacedPersistentVolumeClaimCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create core v1 namespaced persistent volume claim created response
func (o *CreateCoreV1NamespacedPersistentVolumeClaimCreated) SetPayload(payload *models.IoK8sAPICoreV1PersistentVolumeClaim) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateCoreV1NamespacedPersistentVolumeClaimCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateCoreV1NamespacedPersistentVolumeClaimAcceptedCode is the HTTP code returned for type CreateCoreV1NamespacedPersistentVolumeClaimAccepted
const CreateCoreV1NamespacedPersistentVolumeClaimAcceptedCode int = 202

/*CreateCoreV1NamespacedPersistentVolumeClaimAccepted Accepted

swagger:response createCoreV1NamespacedPersistentVolumeClaimAccepted
*/
type CreateCoreV1NamespacedPersistentVolumeClaimAccepted struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPICoreV1PersistentVolumeClaim `json:"body,omitempty"`
}

// NewCreateCoreV1NamespacedPersistentVolumeClaimAccepted creates CreateCoreV1NamespacedPersistentVolumeClaimAccepted with default headers values
func NewCreateCoreV1NamespacedPersistentVolumeClaimAccepted() *CreateCoreV1NamespacedPersistentVolumeClaimAccepted {

	return &CreateCoreV1NamespacedPersistentVolumeClaimAccepted{}
}

// WithPayload adds the payload to the create core v1 namespaced persistent volume claim accepted response
func (o *CreateCoreV1NamespacedPersistentVolumeClaimAccepted) WithPayload(payload *models.IoK8sAPICoreV1PersistentVolumeClaim) *CreateCoreV1NamespacedPersistentVolumeClaimAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create core v1 namespaced persistent volume claim accepted response
func (o *CreateCoreV1NamespacedPersistentVolumeClaimAccepted) SetPayload(payload *models.IoK8sAPICoreV1PersistentVolumeClaim) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateCoreV1NamespacedPersistentVolumeClaimAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(202)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateCoreV1NamespacedPersistentVolumeClaimUnauthorizedCode is the HTTP code returned for type CreateCoreV1NamespacedPersistentVolumeClaimUnauthorized
const CreateCoreV1NamespacedPersistentVolumeClaimUnauthorizedCode int = 401

/*CreateCoreV1NamespacedPersistentVolumeClaimUnauthorized Unauthorized

swagger:response createCoreV1NamespacedPersistentVolumeClaimUnauthorized
*/
type CreateCoreV1NamespacedPersistentVolumeClaimUnauthorized struct {
}

// NewCreateCoreV1NamespacedPersistentVolumeClaimUnauthorized creates CreateCoreV1NamespacedPersistentVolumeClaimUnauthorized with default headers values
func NewCreateCoreV1NamespacedPersistentVolumeClaimUnauthorized() *CreateCoreV1NamespacedPersistentVolumeClaimUnauthorized {

	return &CreateCoreV1NamespacedPersistentVolumeClaimUnauthorized{}
}

// WriteResponse to the client
func (o *CreateCoreV1NamespacedPersistentVolumeClaimUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
