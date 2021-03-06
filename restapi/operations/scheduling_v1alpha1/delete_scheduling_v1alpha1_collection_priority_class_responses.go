// Code generated by go-swagger; DO NOT EDIT.

package scheduling_v1alpha1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// DeleteSchedulingV1alpha1CollectionPriorityClassOKCode is the HTTP code returned for type DeleteSchedulingV1alpha1CollectionPriorityClassOK
const DeleteSchedulingV1alpha1CollectionPriorityClassOKCode int = 200

/*DeleteSchedulingV1alpha1CollectionPriorityClassOK OK

swagger:response deleteSchedulingV1alpha1CollectionPriorityClassOK
*/
type DeleteSchedulingV1alpha1CollectionPriorityClassOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1Status `json:"body,omitempty"`
}

// NewDeleteSchedulingV1alpha1CollectionPriorityClassOK creates DeleteSchedulingV1alpha1CollectionPriorityClassOK with default headers values
func NewDeleteSchedulingV1alpha1CollectionPriorityClassOK() *DeleteSchedulingV1alpha1CollectionPriorityClassOK {

	return &DeleteSchedulingV1alpha1CollectionPriorityClassOK{}
}

// WithPayload adds the payload to the delete scheduling v1alpha1 collection priority class o k response
func (o *DeleteSchedulingV1alpha1CollectionPriorityClassOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) *DeleteSchedulingV1alpha1CollectionPriorityClassOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete scheduling v1alpha1 collection priority class o k response
func (o *DeleteSchedulingV1alpha1CollectionPriorityClassOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteSchedulingV1alpha1CollectionPriorityClassOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteSchedulingV1alpha1CollectionPriorityClassUnauthorizedCode is the HTTP code returned for type DeleteSchedulingV1alpha1CollectionPriorityClassUnauthorized
const DeleteSchedulingV1alpha1CollectionPriorityClassUnauthorizedCode int = 401

/*DeleteSchedulingV1alpha1CollectionPriorityClassUnauthorized Unauthorized

swagger:response deleteSchedulingV1alpha1CollectionPriorityClassUnauthorized
*/
type DeleteSchedulingV1alpha1CollectionPriorityClassUnauthorized struct {
}

// NewDeleteSchedulingV1alpha1CollectionPriorityClassUnauthorized creates DeleteSchedulingV1alpha1CollectionPriorityClassUnauthorized with default headers values
func NewDeleteSchedulingV1alpha1CollectionPriorityClassUnauthorized() *DeleteSchedulingV1alpha1CollectionPriorityClassUnauthorized {

	return &DeleteSchedulingV1alpha1CollectionPriorityClassUnauthorized{}
}

// WriteResponse to the client
func (o *DeleteSchedulingV1alpha1CollectionPriorityClassUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
