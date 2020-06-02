// Code generated by go-swagger; DO NOT EDIT.

package storage_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/models"
)

// DeleteStorageV1beta1StorageClassOKCode is the HTTP code returned for type DeleteStorageV1beta1StorageClassOK
const DeleteStorageV1beta1StorageClassOKCode int = 200

/*DeleteStorageV1beta1StorageClassOK OK

swagger:response deleteStorageV1beta1StorageClassOK
*/
type DeleteStorageV1beta1StorageClassOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIStorageV1beta1StorageClass `json:"body,omitempty"`
}

// NewDeleteStorageV1beta1StorageClassOK creates DeleteStorageV1beta1StorageClassOK with default headers values
func NewDeleteStorageV1beta1StorageClassOK() *DeleteStorageV1beta1StorageClassOK {

	return &DeleteStorageV1beta1StorageClassOK{}
}

// WithPayload adds the payload to the delete storage v1beta1 storage class o k response
func (o *DeleteStorageV1beta1StorageClassOK) WithPayload(payload *models.IoK8sAPIStorageV1beta1StorageClass) *DeleteStorageV1beta1StorageClassOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete storage v1beta1 storage class o k response
func (o *DeleteStorageV1beta1StorageClassOK) SetPayload(payload *models.IoK8sAPIStorageV1beta1StorageClass) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteStorageV1beta1StorageClassOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteStorageV1beta1StorageClassAcceptedCode is the HTTP code returned for type DeleteStorageV1beta1StorageClassAccepted
const DeleteStorageV1beta1StorageClassAcceptedCode int = 202

/*DeleteStorageV1beta1StorageClassAccepted Accepted

swagger:response deleteStorageV1beta1StorageClassAccepted
*/
type DeleteStorageV1beta1StorageClassAccepted struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIStorageV1beta1StorageClass `json:"body,omitempty"`
}

// NewDeleteStorageV1beta1StorageClassAccepted creates DeleteStorageV1beta1StorageClassAccepted with default headers values
func NewDeleteStorageV1beta1StorageClassAccepted() *DeleteStorageV1beta1StorageClassAccepted {

	return &DeleteStorageV1beta1StorageClassAccepted{}
}

// WithPayload adds the payload to the delete storage v1beta1 storage class accepted response
func (o *DeleteStorageV1beta1StorageClassAccepted) WithPayload(payload *models.IoK8sAPIStorageV1beta1StorageClass) *DeleteStorageV1beta1StorageClassAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete storage v1beta1 storage class accepted response
func (o *DeleteStorageV1beta1StorageClassAccepted) SetPayload(payload *models.IoK8sAPIStorageV1beta1StorageClass) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteStorageV1beta1StorageClassAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(202)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteStorageV1beta1StorageClassUnauthorizedCode is the HTTP code returned for type DeleteStorageV1beta1StorageClassUnauthorized
const DeleteStorageV1beta1StorageClassUnauthorizedCode int = 401

/*DeleteStorageV1beta1StorageClassUnauthorized Unauthorized

swagger:response deleteStorageV1beta1StorageClassUnauthorized
*/
type DeleteStorageV1beta1StorageClassUnauthorized struct {
}

// NewDeleteStorageV1beta1StorageClassUnauthorized creates DeleteStorageV1beta1StorageClassUnauthorized with default headers values
func NewDeleteStorageV1beta1StorageClassUnauthorized() *DeleteStorageV1beta1StorageClassUnauthorized {

	return &DeleteStorageV1beta1StorageClassUnauthorized{}
}

// WriteResponse to the client
func (o *DeleteStorageV1beta1StorageClassUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
