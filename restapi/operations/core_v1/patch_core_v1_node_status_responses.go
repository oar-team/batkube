// Code generated by go-swagger; DO NOT EDIT.

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// PatchCoreV1NodeStatusOKCode is the HTTP code returned for type PatchCoreV1NodeStatusOK
const PatchCoreV1NodeStatusOKCode int = 200

/*PatchCoreV1NodeStatusOK OK

swagger:response patchCoreV1NodeStatusOK
*/
type PatchCoreV1NodeStatusOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPICoreV1Node `json:"body,omitempty"`
}

// NewPatchCoreV1NodeStatusOK creates PatchCoreV1NodeStatusOK with default headers values
func NewPatchCoreV1NodeStatusOK() *PatchCoreV1NodeStatusOK {

	return &PatchCoreV1NodeStatusOK{}
}

// WithPayload adds the payload to the patch core v1 node status o k response
func (o *PatchCoreV1NodeStatusOK) WithPayload(payload *models.IoK8sAPICoreV1Node) *PatchCoreV1NodeStatusOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch core v1 node status o k response
func (o *PatchCoreV1NodeStatusOK) SetPayload(payload *models.IoK8sAPICoreV1Node) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchCoreV1NodeStatusOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchCoreV1NodeStatusUnauthorizedCode is the HTTP code returned for type PatchCoreV1NodeStatusUnauthorized
const PatchCoreV1NodeStatusUnauthorizedCode int = 401

/*PatchCoreV1NodeStatusUnauthorized Unauthorized

swagger:response patchCoreV1NodeStatusUnauthorized
*/
type PatchCoreV1NodeStatusUnauthorized struct {
}

// NewPatchCoreV1NodeStatusUnauthorized creates PatchCoreV1NodeStatusUnauthorized with default headers values
func NewPatchCoreV1NodeStatusUnauthorized() *PatchCoreV1NodeStatusUnauthorized {

	return &PatchCoreV1NodeStatusUnauthorized{}
}

// WriteResponse to the client
func (o *PatchCoreV1NodeStatusUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
