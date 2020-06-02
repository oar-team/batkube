// Code generated by go-swagger; DO NOT EDIT.

package scheduling_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/models"
)

// DeleteSchedulingV1CollectionPriorityClassOKCode is the HTTP code returned for type DeleteSchedulingV1CollectionPriorityClassOK
const DeleteSchedulingV1CollectionPriorityClassOKCode int = 200

/*DeleteSchedulingV1CollectionPriorityClassOK OK

swagger:response deleteSchedulingV1CollectionPriorityClassOK
*/
type DeleteSchedulingV1CollectionPriorityClassOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1Status `json:"body,omitempty"`
}

// NewDeleteSchedulingV1CollectionPriorityClassOK creates DeleteSchedulingV1CollectionPriorityClassOK with default headers values
func NewDeleteSchedulingV1CollectionPriorityClassOK() *DeleteSchedulingV1CollectionPriorityClassOK {

	return &DeleteSchedulingV1CollectionPriorityClassOK{}
}

// WithPayload adds the payload to the delete scheduling v1 collection priority class o k response
func (o *DeleteSchedulingV1CollectionPriorityClassOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) *DeleteSchedulingV1CollectionPriorityClassOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete scheduling v1 collection priority class o k response
func (o *DeleteSchedulingV1CollectionPriorityClassOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteSchedulingV1CollectionPriorityClassOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteSchedulingV1CollectionPriorityClassUnauthorizedCode is the HTTP code returned for type DeleteSchedulingV1CollectionPriorityClassUnauthorized
const DeleteSchedulingV1CollectionPriorityClassUnauthorizedCode int = 401

/*DeleteSchedulingV1CollectionPriorityClassUnauthorized Unauthorized

swagger:response deleteSchedulingV1CollectionPriorityClassUnauthorized
*/
type DeleteSchedulingV1CollectionPriorityClassUnauthorized struct {
}

// NewDeleteSchedulingV1CollectionPriorityClassUnauthorized creates DeleteSchedulingV1CollectionPriorityClassUnauthorized with default headers values
func NewDeleteSchedulingV1CollectionPriorityClassUnauthorized() *DeleteSchedulingV1CollectionPriorityClassUnauthorized {

	return &DeleteSchedulingV1CollectionPriorityClassUnauthorized{}
}

// WriteResponse to the client
func (o *DeleteSchedulingV1CollectionPriorityClassUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
