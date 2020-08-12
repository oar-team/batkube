// Code generated by go-swagger; DO NOT EDIT.

package node_v1alpha1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// PatchNodeV1alpha1RuntimeClassOKCode is the HTTP code returned for type PatchNodeV1alpha1RuntimeClassOK
const PatchNodeV1alpha1RuntimeClassOKCode int = 200

/*PatchNodeV1alpha1RuntimeClassOK OK

swagger:response patchNodeV1alpha1RuntimeClassOK
*/
type PatchNodeV1alpha1RuntimeClassOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPINodeV1alpha1RuntimeClass `json:"body,omitempty"`
}

// NewPatchNodeV1alpha1RuntimeClassOK creates PatchNodeV1alpha1RuntimeClassOK with default headers values
func NewPatchNodeV1alpha1RuntimeClassOK() *PatchNodeV1alpha1RuntimeClassOK {

	return &PatchNodeV1alpha1RuntimeClassOK{}
}

// WithPayload adds the payload to the patch node v1alpha1 runtime class o k response
func (o *PatchNodeV1alpha1RuntimeClassOK) WithPayload(payload *models.IoK8sAPINodeV1alpha1RuntimeClass) *PatchNodeV1alpha1RuntimeClassOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch node v1alpha1 runtime class o k response
func (o *PatchNodeV1alpha1RuntimeClassOK) SetPayload(payload *models.IoK8sAPINodeV1alpha1RuntimeClass) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchNodeV1alpha1RuntimeClassOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchNodeV1alpha1RuntimeClassUnauthorizedCode is the HTTP code returned for type PatchNodeV1alpha1RuntimeClassUnauthorized
const PatchNodeV1alpha1RuntimeClassUnauthorizedCode int = 401

/*PatchNodeV1alpha1RuntimeClassUnauthorized Unauthorized

swagger:response patchNodeV1alpha1RuntimeClassUnauthorized
*/
type PatchNodeV1alpha1RuntimeClassUnauthorized struct {
}

// NewPatchNodeV1alpha1RuntimeClassUnauthorized creates PatchNodeV1alpha1RuntimeClassUnauthorized with default headers values
func NewPatchNodeV1alpha1RuntimeClassUnauthorized() *PatchNodeV1alpha1RuntimeClassUnauthorized {

	return &PatchNodeV1alpha1RuntimeClassUnauthorized{}
}

// WriteResponse to the client
func (o *PatchNodeV1alpha1RuntimeClassUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
