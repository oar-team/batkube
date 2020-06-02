// Code generated by go-swagger; DO NOT EDIT.

package storage_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/models"
)

// ReplaceStorageV1StorageClassOKCode is the HTTP code returned for type ReplaceStorageV1StorageClassOK
const ReplaceStorageV1StorageClassOKCode int = 200

/*ReplaceStorageV1StorageClassOK OK

swagger:response replaceStorageV1StorageClassOK
*/
type ReplaceStorageV1StorageClassOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIStorageV1StorageClass `json:"body,omitempty"`
}

// NewReplaceStorageV1StorageClassOK creates ReplaceStorageV1StorageClassOK with default headers values
func NewReplaceStorageV1StorageClassOK() *ReplaceStorageV1StorageClassOK {

	return &ReplaceStorageV1StorageClassOK{}
}

// WithPayload adds the payload to the replace storage v1 storage class o k response
func (o *ReplaceStorageV1StorageClassOK) WithPayload(payload *models.IoK8sAPIStorageV1StorageClass) *ReplaceStorageV1StorageClassOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace storage v1 storage class o k response
func (o *ReplaceStorageV1StorageClassOK) SetPayload(payload *models.IoK8sAPIStorageV1StorageClass) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceStorageV1StorageClassOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceStorageV1StorageClassCreatedCode is the HTTP code returned for type ReplaceStorageV1StorageClassCreated
const ReplaceStorageV1StorageClassCreatedCode int = 201

/*ReplaceStorageV1StorageClassCreated Created

swagger:response replaceStorageV1StorageClassCreated
*/
type ReplaceStorageV1StorageClassCreated struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIStorageV1StorageClass `json:"body,omitempty"`
}

// NewReplaceStorageV1StorageClassCreated creates ReplaceStorageV1StorageClassCreated with default headers values
func NewReplaceStorageV1StorageClassCreated() *ReplaceStorageV1StorageClassCreated {

	return &ReplaceStorageV1StorageClassCreated{}
}

// WithPayload adds the payload to the replace storage v1 storage class created response
func (o *ReplaceStorageV1StorageClassCreated) WithPayload(payload *models.IoK8sAPIStorageV1StorageClass) *ReplaceStorageV1StorageClassCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace storage v1 storage class created response
func (o *ReplaceStorageV1StorageClassCreated) SetPayload(payload *models.IoK8sAPIStorageV1StorageClass) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceStorageV1StorageClassCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceStorageV1StorageClassUnauthorizedCode is the HTTP code returned for type ReplaceStorageV1StorageClassUnauthorized
const ReplaceStorageV1StorageClassUnauthorizedCode int = 401

/*ReplaceStorageV1StorageClassUnauthorized Unauthorized

swagger:response replaceStorageV1StorageClassUnauthorized
*/
type ReplaceStorageV1StorageClassUnauthorized struct {
}

// NewReplaceStorageV1StorageClassUnauthorized creates ReplaceStorageV1StorageClassUnauthorized with default headers values
func NewReplaceStorageV1StorageClassUnauthorized() *ReplaceStorageV1StorageClassUnauthorized {

	return &ReplaceStorageV1StorageClassUnauthorized{}
}

// WriteResponse to the client
func (o *ReplaceStorageV1StorageClassUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
