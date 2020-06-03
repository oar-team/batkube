// Code generated by go-swagger; DO NOT EDIT.

package coordination_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/models"
)

// PatchCoordinationV1beta1NamespacedLeaseOKCode is the HTTP code returned for type PatchCoordinationV1beta1NamespacedLeaseOK
const PatchCoordinationV1beta1NamespacedLeaseOKCode int = 200

/*PatchCoordinationV1beta1NamespacedLeaseOK OK

swagger:response patchCoordinationV1beta1NamespacedLeaseOK
*/
type PatchCoordinationV1beta1NamespacedLeaseOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPICoordinationV1beta1Lease `json:"body,omitempty"`
}

// NewPatchCoordinationV1beta1NamespacedLeaseOK creates PatchCoordinationV1beta1NamespacedLeaseOK with default headers values
func NewPatchCoordinationV1beta1NamespacedLeaseOK() *PatchCoordinationV1beta1NamespacedLeaseOK {

	return &PatchCoordinationV1beta1NamespacedLeaseOK{}
}

// WithPayload adds the payload to the patch coordination v1beta1 namespaced lease o k response
func (o *PatchCoordinationV1beta1NamespacedLeaseOK) WithPayload(payload *models.IoK8sAPICoordinationV1beta1Lease) *PatchCoordinationV1beta1NamespacedLeaseOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch coordination v1beta1 namespaced lease o k response
func (o *PatchCoordinationV1beta1NamespacedLeaseOK) SetPayload(payload *models.IoK8sAPICoordinationV1beta1Lease) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchCoordinationV1beta1NamespacedLeaseOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchCoordinationV1beta1NamespacedLeaseUnauthorizedCode is the HTTP code returned for type PatchCoordinationV1beta1NamespacedLeaseUnauthorized
const PatchCoordinationV1beta1NamespacedLeaseUnauthorizedCode int = 401

/*PatchCoordinationV1beta1NamespacedLeaseUnauthorized Unauthorized

swagger:response patchCoordinationV1beta1NamespacedLeaseUnauthorized
*/
type PatchCoordinationV1beta1NamespacedLeaseUnauthorized struct {
}

// NewPatchCoordinationV1beta1NamespacedLeaseUnauthorized creates PatchCoordinationV1beta1NamespacedLeaseUnauthorized with default headers values
func NewPatchCoordinationV1beta1NamespacedLeaseUnauthorized() *PatchCoordinationV1beta1NamespacedLeaseUnauthorized {

	return &PatchCoordinationV1beta1NamespacedLeaseUnauthorized{}
}

// WriteResponse to the client
func (o *PatchCoordinationV1beta1NamespacedLeaseUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}