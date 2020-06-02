// Code generated by go-swagger; DO NOT EDIT.

package storage_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/models"
)

// ListStorageV1StorageClassOKCode is the HTTP code returned for type ListStorageV1StorageClassOK
const ListStorageV1StorageClassOKCode int = 200

/*ListStorageV1StorageClassOK OK

swagger:response listStorageV1StorageClassOK
*/
type ListStorageV1StorageClassOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIStorageV1StorageClassList `json:"body,omitempty"`
}

// NewListStorageV1StorageClassOK creates ListStorageV1StorageClassOK with default headers values
func NewListStorageV1StorageClassOK() *ListStorageV1StorageClassOK {

	return &ListStorageV1StorageClassOK{}
}

// WithPayload adds the payload to the list storage v1 storage class o k response
func (o *ListStorageV1StorageClassOK) WithPayload(payload *models.IoK8sAPIStorageV1StorageClassList) *ListStorageV1StorageClassOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list storage v1 storage class o k response
func (o *ListStorageV1StorageClassOK) SetPayload(payload *models.IoK8sAPIStorageV1StorageClassList) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListStorageV1StorageClassOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListStorageV1StorageClassUnauthorizedCode is the HTTP code returned for type ListStorageV1StorageClassUnauthorized
const ListStorageV1StorageClassUnauthorizedCode int = 401

/*ListStorageV1StorageClassUnauthorized Unauthorized

swagger:response listStorageV1StorageClassUnauthorized
*/
type ListStorageV1StorageClassUnauthorized struct {
}

// NewListStorageV1StorageClassUnauthorized creates ListStorageV1StorageClassUnauthorized with default headers values
func NewListStorageV1StorageClassUnauthorized() *ListStorageV1StorageClassUnauthorized {

	return &ListStorageV1StorageClassUnauthorized{}
}

// WriteResponse to the client
func (o *ListStorageV1StorageClassUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
