// Code generated by go-swagger; DO NOT EDIT.

package storage_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// CreateStorageV1CSIDriverOKCode is the HTTP code returned for type CreateStorageV1CSIDriverOK
const CreateStorageV1CSIDriverOKCode int = 200

/*CreateStorageV1CSIDriverOK OK

swagger:response createStorageV1CSIDriverOK
*/
type CreateStorageV1CSIDriverOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIStorageV1CSIDriver `json:"body,omitempty"`
}

// NewCreateStorageV1CSIDriverOK creates CreateStorageV1CSIDriverOK with default headers values
func NewCreateStorageV1CSIDriverOK() *CreateStorageV1CSIDriverOK {

	return &CreateStorageV1CSIDriverOK{}
}

// WithPayload adds the payload to the create storage v1 c s i driver o k response
func (o *CreateStorageV1CSIDriverOK) WithPayload(payload *models.IoK8sAPIStorageV1CSIDriver) *CreateStorageV1CSIDriverOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create storage v1 c s i driver o k response
func (o *CreateStorageV1CSIDriverOK) SetPayload(payload *models.IoK8sAPIStorageV1CSIDriver) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateStorageV1CSIDriverOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateStorageV1CSIDriverCreatedCode is the HTTP code returned for type CreateStorageV1CSIDriverCreated
const CreateStorageV1CSIDriverCreatedCode int = 201

/*CreateStorageV1CSIDriverCreated Created

swagger:response createStorageV1CSIDriverCreated
*/
type CreateStorageV1CSIDriverCreated struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIStorageV1CSIDriver `json:"body,omitempty"`
}

// NewCreateStorageV1CSIDriverCreated creates CreateStorageV1CSIDriverCreated with default headers values
func NewCreateStorageV1CSIDriverCreated() *CreateStorageV1CSIDriverCreated {

	return &CreateStorageV1CSIDriverCreated{}
}

// WithPayload adds the payload to the create storage v1 c s i driver created response
func (o *CreateStorageV1CSIDriverCreated) WithPayload(payload *models.IoK8sAPIStorageV1CSIDriver) *CreateStorageV1CSIDriverCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create storage v1 c s i driver created response
func (o *CreateStorageV1CSIDriverCreated) SetPayload(payload *models.IoK8sAPIStorageV1CSIDriver) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateStorageV1CSIDriverCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateStorageV1CSIDriverAcceptedCode is the HTTP code returned for type CreateStorageV1CSIDriverAccepted
const CreateStorageV1CSIDriverAcceptedCode int = 202

/*CreateStorageV1CSIDriverAccepted Accepted

swagger:response createStorageV1CSIDriverAccepted
*/
type CreateStorageV1CSIDriverAccepted struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIStorageV1CSIDriver `json:"body,omitempty"`
}

// NewCreateStorageV1CSIDriverAccepted creates CreateStorageV1CSIDriverAccepted with default headers values
func NewCreateStorageV1CSIDriverAccepted() *CreateStorageV1CSIDriverAccepted {

	return &CreateStorageV1CSIDriverAccepted{}
}

// WithPayload adds the payload to the create storage v1 c s i driver accepted response
func (o *CreateStorageV1CSIDriverAccepted) WithPayload(payload *models.IoK8sAPIStorageV1CSIDriver) *CreateStorageV1CSIDriverAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create storage v1 c s i driver accepted response
func (o *CreateStorageV1CSIDriverAccepted) SetPayload(payload *models.IoK8sAPIStorageV1CSIDriver) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateStorageV1CSIDriverAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(202)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateStorageV1CSIDriverUnauthorizedCode is the HTTP code returned for type CreateStorageV1CSIDriverUnauthorized
const CreateStorageV1CSIDriverUnauthorizedCode int = 401

/*CreateStorageV1CSIDriverUnauthorized Unauthorized

swagger:response createStorageV1CSIDriverUnauthorized
*/
type CreateStorageV1CSIDriverUnauthorized struct {
}

// NewCreateStorageV1CSIDriverUnauthorized creates CreateStorageV1CSIDriverUnauthorized with default headers values
func NewCreateStorageV1CSIDriverUnauthorized() *CreateStorageV1CSIDriverUnauthorized {

	return &CreateStorageV1CSIDriverUnauthorized{}
}

// WriteResponse to the client
func (o *CreateStorageV1CSIDriverUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
