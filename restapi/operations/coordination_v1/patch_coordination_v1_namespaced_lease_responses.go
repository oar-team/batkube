// Code generated by go-swagger; DO NOT EDIT.

package coordination_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// PatchCoordinationV1NamespacedLeaseOKCode is the HTTP code returned for type PatchCoordinationV1NamespacedLeaseOK
const PatchCoordinationV1NamespacedLeaseOKCode int = 200

/*PatchCoordinationV1NamespacedLeaseOK OK

swagger:response patchCoordinationV1NamespacedLeaseOK
*/
type PatchCoordinationV1NamespacedLeaseOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPICoordinationV1Lease `json:"body,omitempty"`
}

// NewPatchCoordinationV1NamespacedLeaseOK creates PatchCoordinationV1NamespacedLeaseOK with default headers values
func NewPatchCoordinationV1NamespacedLeaseOK() *PatchCoordinationV1NamespacedLeaseOK {

	return &PatchCoordinationV1NamespacedLeaseOK{}
}

// WithPayload adds the payload to the patch coordination v1 namespaced lease o k response
func (o *PatchCoordinationV1NamespacedLeaseOK) WithPayload(payload *models.IoK8sAPICoordinationV1Lease) *PatchCoordinationV1NamespacedLeaseOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch coordination v1 namespaced lease o k response
func (o *PatchCoordinationV1NamespacedLeaseOK) SetPayload(payload *models.IoK8sAPICoordinationV1Lease) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchCoordinationV1NamespacedLeaseOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchCoordinationV1NamespacedLeaseUnauthorizedCode is the HTTP code returned for type PatchCoordinationV1NamespacedLeaseUnauthorized
const PatchCoordinationV1NamespacedLeaseUnauthorizedCode int = 401

/*PatchCoordinationV1NamespacedLeaseUnauthorized Unauthorized

swagger:response patchCoordinationV1NamespacedLeaseUnauthorized
*/
type PatchCoordinationV1NamespacedLeaseUnauthorized struct {
}

// NewPatchCoordinationV1NamespacedLeaseUnauthorized creates PatchCoordinationV1NamespacedLeaseUnauthorized with default headers values
func NewPatchCoordinationV1NamespacedLeaseUnauthorized() *PatchCoordinationV1NamespacedLeaseUnauthorized {

	return &PatchCoordinationV1NamespacedLeaseUnauthorized{}
}

// WriteResponse to the client
func (o *PatchCoordinationV1NamespacedLeaseUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
