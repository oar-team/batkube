// Code generated by go-swagger; DO NOT EDIT.

package scheduling_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// DeleteSchedulingV1beta1PriorityClassOKCode is the HTTP code returned for type DeleteSchedulingV1beta1PriorityClassOK
const DeleteSchedulingV1beta1PriorityClassOKCode int = 200

/*DeleteSchedulingV1beta1PriorityClassOK OK

swagger:response deleteSchedulingV1beta1PriorityClassOK
*/
type DeleteSchedulingV1beta1PriorityClassOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1Status `json:"body,omitempty"`
}

// NewDeleteSchedulingV1beta1PriorityClassOK creates DeleteSchedulingV1beta1PriorityClassOK with default headers values
func NewDeleteSchedulingV1beta1PriorityClassOK() *DeleteSchedulingV1beta1PriorityClassOK {

	return &DeleteSchedulingV1beta1PriorityClassOK{}
}

// WithPayload adds the payload to the delete scheduling v1beta1 priority class o k response
func (o *DeleteSchedulingV1beta1PriorityClassOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) *DeleteSchedulingV1beta1PriorityClassOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete scheduling v1beta1 priority class o k response
func (o *DeleteSchedulingV1beta1PriorityClassOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteSchedulingV1beta1PriorityClassOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteSchedulingV1beta1PriorityClassAcceptedCode is the HTTP code returned for type DeleteSchedulingV1beta1PriorityClassAccepted
const DeleteSchedulingV1beta1PriorityClassAcceptedCode int = 202

/*DeleteSchedulingV1beta1PriorityClassAccepted Accepted

swagger:response deleteSchedulingV1beta1PriorityClassAccepted
*/
type DeleteSchedulingV1beta1PriorityClassAccepted struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1Status `json:"body,omitempty"`
}

// NewDeleteSchedulingV1beta1PriorityClassAccepted creates DeleteSchedulingV1beta1PriorityClassAccepted with default headers values
func NewDeleteSchedulingV1beta1PriorityClassAccepted() *DeleteSchedulingV1beta1PriorityClassAccepted {

	return &DeleteSchedulingV1beta1PriorityClassAccepted{}
}

// WithPayload adds the payload to the delete scheduling v1beta1 priority class accepted response
func (o *DeleteSchedulingV1beta1PriorityClassAccepted) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) *DeleteSchedulingV1beta1PriorityClassAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete scheduling v1beta1 priority class accepted response
func (o *DeleteSchedulingV1beta1PriorityClassAccepted) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteSchedulingV1beta1PriorityClassAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(202)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteSchedulingV1beta1PriorityClassUnauthorizedCode is the HTTP code returned for type DeleteSchedulingV1beta1PriorityClassUnauthorized
const DeleteSchedulingV1beta1PriorityClassUnauthorizedCode int = 401

/*DeleteSchedulingV1beta1PriorityClassUnauthorized Unauthorized

swagger:response deleteSchedulingV1beta1PriorityClassUnauthorized
*/
type DeleteSchedulingV1beta1PriorityClassUnauthorized struct {
}

// NewDeleteSchedulingV1beta1PriorityClassUnauthorized creates DeleteSchedulingV1beta1PriorityClassUnauthorized with default headers values
func NewDeleteSchedulingV1beta1PriorityClassUnauthorized() *DeleteSchedulingV1beta1PriorityClassUnauthorized {

	return &DeleteSchedulingV1beta1PriorityClassUnauthorized{}
}

// WriteResponse to the client
func (o *DeleteSchedulingV1beta1PriorityClassUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
