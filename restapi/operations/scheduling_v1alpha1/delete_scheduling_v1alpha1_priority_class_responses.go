// Code generated by go-swagger; DO NOT EDIT.

package scheduling_v1alpha1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// DeleteSchedulingV1alpha1PriorityClassOKCode is the HTTP code returned for type DeleteSchedulingV1alpha1PriorityClassOK
const DeleteSchedulingV1alpha1PriorityClassOKCode int = 200

/*DeleteSchedulingV1alpha1PriorityClassOK OK

swagger:response deleteSchedulingV1alpha1PriorityClassOK
*/
type DeleteSchedulingV1alpha1PriorityClassOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1Status `json:"body,omitempty"`
}

// NewDeleteSchedulingV1alpha1PriorityClassOK creates DeleteSchedulingV1alpha1PriorityClassOK with default headers values
func NewDeleteSchedulingV1alpha1PriorityClassOK() *DeleteSchedulingV1alpha1PriorityClassOK {

	return &DeleteSchedulingV1alpha1PriorityClassOK{}
}

// WithPayload adds the payload to the delete scheduling v1alpha1 priority class o k response
func (o *DeleteSchedulingV1alpha1PriorityClassOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) *DeleteSchedulingV1alpha1PriorityClassOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete scheduling v1alpha1 priority class o k response
func (o *DeleteSchedulingV1alpha1PriorityClassOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteSchedulingV1alpha1PriorityClassOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteSchedulingV1alpha1PriorityClassAcceptedCode is the HTTP code returned for type DeleteSchedulingV1alpha1PriorityClassAccepted
const DeleteSchedulingV1alpha1PriorityClassAcceptedCode int = 202

/*DeleteSchedulingV1alpha1PriorityClassAccepted Accepted

swagger:response deleteSchedulingV1alpha1PriorityClassAccepted
*/
type DeleteSchedulingV1alpha1PriorityClassAccepted struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1Status `json:"body,omitempty"`
}

// NewDeleteSchedulingV1alpha1PriorityClassAccepted creates DeleteSchedulingV1alpha1PriorityClassAccepted with default headers values
func NewDeleteSchedulingV1alpha1PriorityClassAccepted() *DeleteSchedulingV1alpha1PriorityClassAccepted {

	return &DeleteSchedulingV1alpha1PriorityClassAccepted{}
}

// WithPayload adds the payload to the delete scheduling v1alpha1 priority class accepted response
func (o *DeleteSchedulingV1alpha1PriorityClassAccepted) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) *DeleteSchedulingV1alpha1PriorityClassAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete scheduling v1alpha1 priority class accepted response
func (o *DeleteSchedulingV1alpha1PriorityClassAccepted) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteSchedulingV1alpha1PriorityClassAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(202)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteSchedulingV1alpha1PriorityClassUnauthorizedCode is the HTTP code returned for type DeleteSchedulingV1alpha1PriorityClassUnauthorized
const DeleteSchedulingV1alpha1PriorityClassUnauthorizedCode int = 401

/*DeleteSchedulingV1alpha1PriorityClassUnauthorized Unauthorized

swagger:response deleteSchedulingV1alpha1PriorityClassUnauthorized
*/
type DeleteSchedulingV1alpha1PriorityClassUnauthorized struct {
}

// NewDeleteSchedulingV1alpha1PriorityClassUnauthorized creates DeleteSchedulingV1alpha1PriorityClassUnauthorized with default headers values
func NewDeleteSchedulingV1alpha1PriorityClassUnauthorized() *DeleteSchedulingV1alpha1PriorityClassUnauthorized {

	return &DeleteSchedulingV1alpha1PriorityClassUnauthorized{}
}

// WriteResponse to the client
func (o *DeleteSchedulingV1alpha1PriorityClassUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
