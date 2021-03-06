// Code generated by go-swagger; DO NOT EDIT.

package scheduling_v1alpha1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// PatchSchedulingV1alpha1PriorityClassOKCode is the HTTP code returned for type PatchSchedulingV1alpha1PriorityClassOK
const PatchSchedulingV1alpha1PriorityClassOKCode int = 200

/*PatchSchedulingV1alpha1PriorityClassOK OK

swagger:response patchSchedulingV1alpha1PriorityClassOK
*/
type PatchSchedulingV1alpha1PriorityClassOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPISchedulingV1alpha1PriorityClass `json:"body,omitempty"`
}

// NewPatchSchedulingV1alpha1PriorityClassOK creates PatchSchedulingV1alpha1PriorityClassOK with default headers values
func NewPatchSchedulingV1alpha1PriorityClassOK() *PatchSchedulingV1alpha1PriorityClassOK {

	return &PatchSchedulingV1alpha1PriorityClassOK{}
}

// WithPayload adds the payload to the patch scheduling v1alpha1 priority class o k response
func (o *PatchSchedulingV1alpha1PriorityClassOK) WithPayload(payload *models.IoK8sAPISchedulingV1alpha1PriorityClass) *PatchSchedulingV1alpha1PriorityClassOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch scheduling v1alpha1 priority class o k response
func (o *PatchSchedulingV1alpha1PriorityClassOK) SetPayload(payload *models.IoK8sAPISchedulingV1alpha1PriorityClass) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchSchedulingV1alpha1PriorityClassOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchSchedulingV1alpha1PriorityClassUnauthorizedCode is the HTTP code returned for type PatchSchedulingV1alpha1PriorityClassUnauthorized
const PatchSchedulingV1alpha1PriorityClassUnauthorizedCode int = 401

/*PatchSchedulingV1alpha1PriorityClassUnauthorized Unauthorized

swagger:response patchSchedulingV1alpha1PriorityClassUnauthorized
*/
type PatchSchedulingV1alpha1PriorityClassUnauthorized struct {
}

// NewPatchSchedulingV1alpha1PriorityClassUnauthorized creates PatchSchedulingV1alpha1PriorityClassUnauthorized with default headers values
func NewPatchSchedulingV1alpha1PriorityClassUnauthorized() *PatchSchedulingV1alpha1PriorityClassUnauthorized {

	return &PatchSchedulingV1alpha1PriorityClassUnauthorized{}
}

// WriteResponse to the client
func (o *PatchSchedulingV1alpha1PriorityClassUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
