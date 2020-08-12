// Code generated by go-swagger; DO NOT EDIT.

package storage_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// ReadStorageV1StorageClassOKCode is the HTTP code returned for type ReadStorageV1StorageClassOK
const ReadStorageV1StorageClassOKCode int = 200

/*ReadStorageV1StorageClassOK OK

swagger:response readStorageV1StorageClassOK
*/
type ReadStorageV1StorageClassOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIStorageV1StorageClass `json:"body,omitempty"`
}

// NewReadStorageV1StorageClassOK creates ReadStorageV1StorageClassOK with default headers values
func NewReadStorageV1StorageClassOK() *ReadStorageV1StorageClassOK {

	return &ReadStorageV1StorageClassOK{}
}

// WithPayload adds the payload to the read storage v1 storage class o k response
func (o *ReadStorageV1StorageClassOK) WithPayload(payload *models.IoK8sAPIStorageV1StorageClass) *ReadStorageV1StorageClassOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the read storage v1 storage class o k response
func (o *ReadStorageV1StorageClassOK) SetPayload(payload *models.IoK8sAPIStorageV1StorageClass) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReadStorageV1StorageClassOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReadStorageV1StorageClassUnauthorizedCode is the HTTP code returned for type ReadStorageV1StorageClassUnauthorized
const ReadStorageV1StorageClassUnauthorizedCode int = 401

/*ReadStorageV1StorageClassUnauthorized Unauthorized

swagger:response readStorageV1StorageClassUnauthorized
*/
type ReadStorageV1StorageClassUnauthorized struct {
}

// NewReadStorageV1StorageClassUnauthorized creates ReadStorageV1StorageClassUnauthorized with default headers values
func NewReadStorageV1StorageClassUnauthorized() *ReadStorageV1StorageClassUnauthorized {

	return &ReadStorageV1StorageClassUnauthorized{}
}

// WriteResponse to the client
func (o *ReadStorageV1StorageClassUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
