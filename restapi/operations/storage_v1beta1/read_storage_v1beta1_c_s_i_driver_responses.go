// Code generated by go-swagger; DO NOT EDIT.

package storage_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// ReadStorageV1beta1CSIDriverOKCode is the HTTP code returned for type ReadStorageV1beta1CSIDriverOK
const ReadStorageV1beta1CSIDriverOKCode int = 200

/*ReadStorageV1beta1CSIDriverOK OK

swagger:response readStorageV1beta1CSIDriverOK
*/
type ReadStorageV1beta1CSIDriverOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIStorageV1beta1CSIDriver `json:"body,omitempty"`
}

// NewReadStorageV1beta1CSIDriverOK creates ReadStorageV1beta1CSIDriverOK with default headers values
func NewReadStorageV1beta1CSIDriverOK() *ReadStorageV1beta1CSIDriverOK {

	return &ReadStorageV1beta1CSIDriverOK{}
}

// WithPayload adds the payload to the read storage v1beta1 c s i driver o k response
func (o *ReadStorageV1beta1CSIDriverOK) WithPayload(payload *models.IoK8sAPIStorageV1beta1CSIDriver) *ReadStorageV1beta1CSIDriverOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the read storage v1beta1 c s i driver o k response
func (o *ReadStorageV1beta1CSIDriverOK) SetPayload(payload *models.IoK8sAPIStorageV1beta1CSIDriver) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReadStorageV1beta1CSIDriverOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReadStorageV1beta1CSIDriverUnauthorizedCode is the HTTP code returned for type ReadStorageV1beta1CSIDriverUnauthorized
const ReadStorageV1beta1CSIDriverUnauthorizedCode int = 401

/*ReadStorageV1beta1CSIDriverUnauthorized Unauthorized

swagger:response readStorageV1beta1CSIDriverUnauthorized
*/
type ReadStorageV1beta1CSIDriverUnauthorized struct {
}

// NewReadStorageV1beta1CSIDriverUnauthorized creates ReadStorageV1beta1CSIDriverUnauthorized with default headers values
func NewReadStorageV1beta1CSIDriverUnauthorized() *ReadStorageV1beta1CSIDriverUnauthorized {

	return &ReadStorageV1beta1CSIDriverUnauthorized{}
}

// WriteResponse to the client
func (o *ReadStorageV1beta1CSIDriverUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
