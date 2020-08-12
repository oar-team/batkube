// Code generated by go-swagger; DO NOT EDIT.

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// DeleteCoreV1NodeOKCode is the HTTP code returned for type DeleteCoreV1NodeOK
const DeleteCoreV1NodeOKCode int = 200

/*DeleteCoreV1NodeOK OK

swagger:response deleteCoreV1NodeOK
*/
type DeleteCoreV1NodeOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1Status `json:"body,omitempty"`
}

// NewDeleteCoreV1NodeOK creates DeleteCoreV1NodeOK with default headers values
func NewDeleteCoreV1NodeOK() *DeleteCoreV1NodeOK {

	return &DeleteCoreV1NodeOK{}
}

// WithPayload adds the payload to the delete core v1 node o k response
func (o *DeleteCoreV1NodeOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) *DeleteCoreV1NodeOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete core v1 node o k response
func (o *DeleteCoreV1NodeOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteCoreV1NodeOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteCoreV1NodeAcceptedCode is the HTTP code returned for type DeleteCoreV1NodeAccepted
const DeleteCoreV1NodeAcceptedCode int = 202

/*DeleteCoreV1NodeAccepted Accepted

swagger:response deleteCoreV1NodeAccepted
*/
type DeleteCoreV1NodeAccepted struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1Status `json:"body,omitempty"`
}

// NewDeleteCoreV1NodeAccepted creates DeleteCoreV1NodeAccepted with default headers values
func NewDeleteCoreV1NodeAccepted() *DeleteCoreV1NodeAccepted {

	return &DeleteCoreV1NodeAccepted{}
}

// WithPayload adds the payload to the delete core v1 node accepted response
func (o *DeleteCoreV1NodeAccepted) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) *DeleteCoreV1NodeAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete core v1 node accepted response
func (o *DeleteCoreV1NodeAccepted) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteCoreV1NodeAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(202)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteCoreV1NodeUnauthorizedCode is the HTTP code returned for type DeleteCoreV1NodeUnauthorized
const DeleteCoreV1NodeUnauthorizedCode int = 401

/*DeleteCoreV1NodeUnauthorized Unauthorized

swagger:response deleteCoreV1NodeUnauthorized
*/
type DeleteCoreV1NodeUnauthorized struct {
}

// NewDeleteCoreV1NodeUnauthorized creates DeleteCoreV1NodeUnauthorized with default headers values
func NewDeleteCoreV1NodeUnauthorized() *DeleteCoreV1NodeUnauthorized {

	return &DeleteCoreV1NodeUnauthorized{}
}

// WriteResponse to the client
func (o *DeleteCoreV1NodeUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
