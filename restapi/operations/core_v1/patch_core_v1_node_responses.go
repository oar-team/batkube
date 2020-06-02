// Code generated by go-swagger; DO NOT EDIT.

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/models"
)

// PatchCoreV1NodeOKCode is the HTTP code returned for type PatchCoreV1NodeOK
const PatchCoreV1NodeOKCode int = 200

/*PatchCoreV1NodeOK OK

swagger:response patchCoreV1NodeOK
*/
type PatchCoreV1NodeOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPICoreV1Node `json:"body,omitempty"`
}

// NewPatchCoreV1NodeOK creates PatchCoreV1NodeOK with default headers values
func NewPatchCoreV1NodeOK() *PatchCoreV1NodeOK {

	return &PatchCoreV1NodeOK{}
}

// WithPayload adds the payload to the patch core v1 node o k response
func (o *PatchCoreV1NodeOK) WithPayload(payload *models.IoK8sAPICoreV1Node) *PatchCoreV1NodeOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch core v1 node o k response
func (o *PatchCoreV1NodeOK) SetPayload(payload *models.IoK8sAPICoreV1Node) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchCoreV1NodeOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchCoreV1NodeUnauthorizedCode is the HTTP code returned for type PatchCoreV1NodeUnauthorized
const PatchCoreV1NodeUnauthorizedCode int = 401

/*PatchCoreV1NodeUnauthorized Unauthorized

swagger:response patchCoreV1NodeUnauthorized
*/
type PatchCoreV1NodeUnauthorized struct {
}

// NewPatchCoreV1NodeUnauthorized creates PatchCoreV1NodeUnauthorized with default headers values
func NewPatchCoreV1NodeUnauthorized() *PatchCoreV1NodeUnauthorized {

	return &PatchCoreV1NodeUnauthorized{}
}

// WriteResponse to the client
func (o *PatchCoreV1NodeUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
