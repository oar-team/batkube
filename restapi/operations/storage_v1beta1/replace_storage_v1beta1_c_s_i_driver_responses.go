// Code generated by go-swagger; DO NOT EDIT.

package storage_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/models"
)

// ReplaceStorageV1beta1CSIDriverOKCode is the HTTP code returned for type ReplaceStorageV1beta1CSIDriverOK
const ReplaceStorageV1beta1CSIDriverOKCode int = 200

/*ReplaceStorageV1beta1CSIDriverOK OK

swagger:response replaceStorageV1beta1CSIDriverOK
*/
type ReplaceStorageV1beta1CSIDriverOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIStorageV1beta1CSIDriver `json:"body,omitempty"`
}

// NewReplaceStorageV1beta1CSIDriverOK creates ReplaceStorageV1beta1CSIDriverOK with default headers values
func NewReplaceStorageV1beta1CSIDriverOK() *ReplaceStorageV1beta1CSIDriverOK {

	return &ReplaceStorageV1beta1CSIDriverOK{}
}

// WithPayload adds the payload to the replace storage v1beta1 c s i driver o k response
func (o *ReplaceStorageV1beta1CSIDriverOK) WithPayload(payload *models.IoK8sAPIStorageV1beta1CSIDriver) *ReplaceStorageV1beta1CSIDriverOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace storage v1beta1 c s i driver o k response
func (o *ReplaceStorageV1beta1CSIDriverOK) SetPayload(payload *models.IoK8sAPIStorageV1beta1CSIDriver) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceStorageV1beta1CSIDriverOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceStorageV1beta1CSIDriverCreatedCode is the HTTP code returned for type ReplaceStorageV1beta1CSIDriverCreated
const ReplaceStorageV1beta1CSIDriverCreatedCode int = 201

/*ReplaceStorageV1beta1CSIDriverCreated Created

swagger:response replaceStorageV1beta1CSIDriverCreated
*/
type ReplaceStorageV1beta1CSIDriverCreated struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIStorageV1beta1CSIDriver `json:"body,omitempty"`
}

// NewReplaceStorageV1beta1CSIDriverCreated creates ReplaceStorageV1beta1CSIDriverCreated with default headers values
func NewReplaceStorageV1beta1CSIDriverCreated() *ReplaceStorageV1beta1CSIDriverCreated {

	return &ReplaceStorageV1beta1CSIDriverCreated{}
}

// WithPayload adds the payload to the replace storage v1beta1 c s i driver created response
func (o *ReplaceStorageV1beta1CSIDriverCreated) WithPayload(payload *models.IoK8sAPIStorageV1beta1CSIDriver) *ReplaceStorageV1beta1CSIDriverCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace storage v1beta1 c s i driver created response
func (o *ReplaceStorageV1beta1CSIDriverCreated) SetPayload(payload *models.IoK8sAPIStorageV1beta1CSIDriver) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceStorageV1beta1CSIDriverCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceStorageV1beta1CSIDriverUnauthorizedCode is the HTTP code returned for type ReplaceStorageV1beta1CSIDriverUnauthorized
const ReplaceStorageV1beta1CSIDriverUnauthorizedCode int = 401

/*ReplaceStorageV1beta1CSIDriverUnauthorized Unauthorized

swagger:response replaceStorageV1beta1CSIDriverUnauthorized
*/
type ReplaceStorageV1beta1CSIDriverUnauthorized struct {
}

// NewReplaceStorageV1beta1CSIDriverUnauthorized creates ReplaceStorageV1beta1CSIDriverUnauthorized with default headers values
func NewReplaceStorageV1beta1CSIDriverUnauthorized() *ReplaceStorageV1beta1CSIDriverUnauthorized {

	return &ReplaceStorageV1beta1CSIDriverUnauthorized{}
}

// WriteResponse to the client
func (o *ReplaceStorageV1beta1CSIDriverUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
