// Code generated by go-swagger; DO NOT EDIT.

package node_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// ReadNodeV1beta1RuntimeClassOKCode is the HTTP code returned for type ReadNodeV1beta1RuntimeClassOK
const ReadNodeV1beta1RuntimeClassOKCode int = 200

/*ReadNodeV1beta1RuntimeClassOK OK

swagger:response readNodeV1beta1RuntimeClassOK
*/
type ReadNodeV1beta1RuntimeClassOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPINodeV1beta1RuntimeClass `json:"body,omitempty"`
}

// NewReadNodeV1beta1RuntimeClassOK creates ReadNodeV1beta1RuntimeClassOK with default headers values
func NewReadNodeV1beta1RuntimeClassOK() *ReadNodeV1beta1RuntimeClassOK {

	return &ReadNodeV1beta1RuntimeClassOK{}
}

// WithPayload adds the payload to the read node v1beta1 runtime class o k response
func (o *ReadNodeV1beta1RuntimeClassOK) WithPayload(payload *models.IoK8sAPINodeV1beta1RuntimeClass) *ReadNodeV1beta1RuntimeClassOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the read node v1beta1 runtime class o k response
func (o *ReadNodeV1beta1RuntimeClassOK) SetPayload(payload *models.IoK8sAPINodeV1beta1RuntimeClass) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReadNodeV1beta1RuntimeClassOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReadNodeV1beta1RuntimeClassUnauthorizedCode is the HTTP code returned for type ReadNodeV1beta1RuntimeClassUnauthorized
const ReadNodeV1beta1RuntimeClassUnauthorizedCode int = 401

/*ReadNodeV1beta1RuntimeClassUnauthorized Unauthorized

swagger:response readNodeV1beta1RuntimeClassUnauthorized
*/
type ReadNodeV1beta1RuntimeClassUnauthorized struct {
}

// NewReadNodeV1beta1RuntimeClassUnauthorized creates ReadNodeV1beta1RuntimeClassUnauthorized with default headers values
func NewReadNodeV1beta1RuntimeClassUnauthorized() *ReadNodeV1beta1RuntimeClassUnauthorized {

	return &ReadNodeV1beta1RuntimeClassUnauthorized{}
}

// WriteResponse to the client
func (o *ReadNodeV1beta1RuntimeClassUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
