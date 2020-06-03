// Code generated by go-swagger; DO NOT EDIT.

package storage_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/models"
)

// CreateStorageV1StorageClassOKCode is the HTTP code returned for type CreateStorageV1StorageClassOK
const CreateStorageV1StorageClassOKCode int = 200

/*CreateStorageV1StorageClassOK OK

swagger:response createStorageV1StorageClassOK
*/
type CreateStorageV1StorageClassOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIStorageV1StorageClass `json:"body,omitempty"`
}

// NewCreateStorageV1StorageClassOK creates CreateStorageV1StorageClassOK with default headers values
func NewCreateStorageV1StorageClassOK() *CreateStorageV1StorageClassOK {

	return &CreateStorageV1StorageClassOK{}
}

// WithPayload adds the payload to the create storage v1 storage class o k response
func (o *CreateStorageV1StorageClassOK) WithPayload(payload *models.IoK8sAPIStorageV1StorageClass) *CreateStorageV1StorageClassOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create storage v1 storage class o k response
func (o *CreateStorageV1StorageClassOK) SetPayload(payload *models.IoK8sAPIStorageV1StorageClass) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateStorageV1StorageClassOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateStorageV1StorageClassCreatedCode is the HTTP code returned for type CreateStorageV1StorageClassCreated
const CreateStorageV1StorageClassCreatedCode int = 201

/*CreateStorageV1StorageClassCreated Created

swagger:response createStorageV1StorageClassCreated
*/
type CreateStorageV1StorageClassCreated struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIStorageV1StorageClass `json:"body,omitempty"`
}

// NewCreateStorageV1StorageClassCreated creates CreateStorageV1StorageClassCreated with default headers values
func NewCreateStorageV1StorageClassCreated() *CreateStorageV1StorageClassCreated {

	return &CreateStorageV1StorageClassCreated{}
}

// WithPayload adds the payload to the create storage v1 storage class created response
func (o *CreateStorageV1StorageClassCreated) WithPayload(payload *models.IoK8sAPIStorageV1StorageClass) *CreateStorageV1StorageClassCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create storage v1 storage class created response
func (o *CreateStorageV1StorageClassCreated) SetPayload(payload *models.IoK8sAPIStorageV1StorageClass) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateStorageV1StorageClassCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateStorageV1StorageClassAcceptedCode is the HTTP code returned for type CreateStorageV1StorageClassAccepted
const CreateStorageV1StorageClassAcceptedCode int = 202

/*CreateStorageV1StorageClassAccepted Accepted

swagger:response createStorageV1StorageClassAccepted
*/
type CreateStorageV1StorageClassAccepted struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIStorageV1StorageClass `json:"body,omitempty"`
}

// NewCreateStorageV1StorageClassAccepted creates CreateStorageV1StorageClassAccepted with default headers values
func NewCreateStorageV1StorageClassAccepted() *CreateStorageV1StorageClassAccepted {

	return &CreateStorageV1StorageClassAccepted{}
}

// WithPayload adds the payload to the create storage v1 storage class accepted response
func (o *CreateStorageV1StorageClassAccepted) WithPayload(payload *models.IoK8sAPIStorageV1StorageClass) *CreateStorageV1StorageClassAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create storage v1 storage class accepted response
func (o *CreateStorageV1StorageClassAccepted) SetPayload(payload *models.IoK8sAPIStorageV1StorageClass) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateStorageV1StorageClassAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(202)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateStorageV1StorageClassUnauthorizedCode is the HTTP code returned for type CreateStorageV1StorageClassUnauthorized
const CreateStorageV1StorageClassUnauthorizedCode int = 401

/*CreateStorageV1StorageClassUnauthorized Unauthorized

swagger:response createStorageV1StorageClassUnauthorized
*/
type CreateStorageV1StorageClassUnauthorized struct {
}

// NewCreateStorageV1StorageClassUnauthorized creates CreateStorageV1StorageClassUnauthorized with default headers values
func NewCreateStorageV1StorageClassUnauthorized() *CreateStorageV1StorageClassUnauthorized {

	return &CreateStorageV1StorageClassUnauthorized{}
}

// WriteResponse to the client
func (o *CreateStorageV1StorageClassUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}