// Code generated by go-swagger; DO NOT EDIT.

package scheduling_v1alpha1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// ListSchedulingV1alpha1PriorityClassOKCode is the HTTP code returned for type ListSchedulingV1alpha1PriorityClassOK
const ListSchedulingV1alpha1PriorityClassOKCode int = 200

/*ListSchedulingV1alpha1PriorityClassOK OK

swagger:response listSchedulingV1alpha1PriorityClassOK
*/
type ListSchedulingV1alpha1PriorityClassOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPISchedulingV1alpha1PriorityClassList `json:"body,omitempty"`
}

// NewListSchedulingV1alpha1PriorityClassOK creates ListSchedulingV1alpha1PriorityClassOK with default headers values
func NewListSchedulingV1alpha1PriorityClassOK() *ListSchedulingV1alpha1PriorityClassOK {

	return &ListSchedulingV1alpha1PriorityClassOK{}
}

// WithPayload adds the payload to the list scheduling v1alpha1 priority class o k response
func (o *ListSchedulingV1alpha1PriorityClassOK) WithPayload(payload *models.IoK8sAPISchedulingV1alpha1PriorityClassList) *ListSchedulingV1alpha1PriorityClassOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list scheduling v1alpha1 priority class o k response
func (o *ListSchedulingV1alpha1PriorityClassOK) SetPayload(payload *models.IoK8sAPISchedulingV1alpha1PriorityClassList) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListSchedulingV1alpha1PriorityClassOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListSchedulingV1alpha1PriorityClassUnauthorizedCode is the HTTP code returned for type ListSchedulingV1alpha1PriorityClassUnauthorized
const ListSchedulingV1alpha1PriorityClassUnauthorizedCode int = 401

/*ListSchedulingV1alpha1PriorityClassUnauthorized Unauthorized

swagger:response listSchedulingV1alpha1PriorityClassUnauthorized
*/
type ListSchedulingV1alpha1PriorityClassUnauthorized struct {
}

// NewListSchedulingV1alpha1PriorityClassUnauthorized creates ListSchedulingV1alpha1PriorityClassUnauthorized with default headers values
func NewListSchedulingV1alpha1PriorityClassUnauthorized() *ListSchedulingV1alpha1PriorityClassUnauthorized {

	return &ListSchedulingV1alpha1PriorityClassUnauthorized{}
}

// WriteResponse to the client
func (o *ListSchedulingV1alpha1PriorityClassUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
