// Code generated by go-swagger; DO NOT EDIT.

package storage_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/models"
)

// DeleteStorageV1beta1CSIDriverOKCode is the HTTP code returned for type DeleteStorageV1beta1CSIDriverOK
const DeleteStorageV1beta1CSIDriverOKCode int = 200

/*DeleteStorageV1beta1CSIDriverOK OK

swagger:response deleteStorageV1beta1CSIDriverOK
*/
type DeleteStorageV1beta1CSIDriverOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIStorageV1beta1CSIDriver `json:"body,omitempty"`
}

// NewDeleteStorageV1beta1CSIDriverOK creates DeleteStorageV1beta1CSIDriverOK with default headers values
func NewDeleteStorageV1beta1CSIDriverOK() *DeleteStorageV1beta1CSIDriverOK {

	return &DeleteStorageV1beta1CSIDriverOK{}
}

// WithPayload adds the payload to the delete storage v1beta1 c s i driver o k response
func (o *DeleteStorageV1beta1CSIDriverOK) WithPayload(payload *models.IoK8sAPIStorageV1beta1CSIDriver) *DeleteStorageV1beta1CSIDriverOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete storage v1beta1 c s i driver o k response
func (o *DeleteStorageV1beta1CSIDriverOK) SetPayload(payload *models.IoK8sAPIStorageV1beta1CSIDriver) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteStorageV1beta1CSIDriverOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteStorageV1beta1CSIDriverAcceptedCode is the HTTP code returned for type DeleteStorageV1beta1CSIDriverAccepted
const DeleteStorageV1beta1CSIDriverAcceptedCode int = 202

/*DeleteStorageV1beta1CSIDriverAccepted Accepted

swagger:response deleteStorageV1beta1CSIDriverAccepted
*/
type DeleteStorageV1beta1CSIDriverAccepted struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIStorageV1beta1CSIDriver `json:"body,omitempty"`
}

// NewDeleteStorageV1beta1CSIDriverAccepted creates DeleteStorageV1beta1CSIDriverAccepted with default headers values
func NewDeleteStorageV1beta1CSIDriverAccepted() *DeleteStorageV1beta1CSIDriverAccepted {

	return &DeleteStorageV1beta1CSIDriverAccepted{}
}

// WithPayload adds the payload to the delete storage v1beta1 c s i driver accepted response
func (o *DeleteStorageV1beta1CSIDriverAccepted) WithPayload(payload *models.IoK8sAPIStorageV1beta1CSIDriver) *DeleteStorageV1beta1CSIDriverAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete storage v1beta1 c s i driver accepted response
func (o *DeleteStorageV1beta1CSIDriverAccepted) SetPayload(payload *models.IoK8sAPIStorageV1beta1CSIDriver) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteStorageV1beta1CSIDriverAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(202)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteStorageV1beta1CSIDriverUnauthorizedCode is the HTTP code returned for type DeleteStorageV1beta1CSIDriverUnauthorized
const DeleteStorageV1beta1CSIDriverUnauthorizedCode int = 401

/*DeleteStorageV1beta1CSIDriverUnauthorized Unauthorized

swagger:response deleteStorageV1beta1CSIDriverUnauthorized
*/
type DeleteStorageV1beta1CSIDriverUnauthorized struct {
}

// NewDeleteStorageV1beta1CSIDriverUnauthorized creates DeleteStorageV1beta1CSIDriverUnauthorized with default headers values
func NewDeleteStorageV1beta1CSIDriverUnauthorized() *DeleteStorageV1beta1CSIDriverUnauthorized {

	return &DeleteStorageV1beta1CSIDriverUnauthorized{}
}

// WriteResponse to the client
func (o *DeleteStorageV1beta1CSIDriverUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}