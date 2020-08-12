// Code generated by go-swagger; DO NOT EDIT.

package node_v1alpha1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// DeleteNodeV1alpha1CollectionRuntimeClassOKCode is the HTTP code returned for type DeleteNodeV1alpha1CollectionRuntimeClassOK
const DeleteNodeV1alpha1CollectionRuntimeClassOKCode int = 200

/*DeleteNodeV1alpha1CollectionRuntimeClassOK OK

swagger:response deleteNodeV1alpha1CollectionRuntimeClassOK
*/
type DeleteNodeV1alpha1CollectionRuntimeClassOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1Status `json:"body,omitempty"`
}

// NewDeleteNodeV1alpha1CollectionRuntimeClassOK creates DeleteNodeV1alpha1CollectionRuntimeClassOK with default headers values
func NewDeleteNodeV1alpha1CollectionRuntimeClassOK() *DeleteNodeV1alpha1CollectionRuntimeClassOK {

	return &DeleteNodeV1alpha1CollectionRuntimeClassOK{}
}

// WithPayload adds the payload to the delete node v1alpha1 collection runtime class o k response
func (o *DeleteNodeV1alpha1CollectionRuntimeClassOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) *DeleteNodeV1alpha1CollectionRuntimeClassOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete node v1alpha1 collection runtime class o k response
func (o *DeleteNodeV1alpha1CollectionRuntimeClassOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteNodeV1alpha1CollectionRuntimeClassOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteNodeV1alpha1CollectionRuntimeClassUnauthorizedCode is the HTTP code returned for type DeleteNodeV1alpha1CollectionRuntimeClassUnauthorized
const DeleteNodeV1alpha1CollectionRuntimeClassUnauthorizedCode int = 401

/*DeleteNodeV1alpha1CollectionRuntimeClassUnauthorized Unauthorized

swagger:response deleteNodeV1alpha1CollectionRuntimeClassUnauthorized
*/
type DeleteNodeV1alpha1CollectionRuntimeClassUnauthorized struct {
}

// NewDeleteNodeV1alpha1CollectionRuntimeClassUnauthorized creates DeleteNodeV1alpha1CollectionRuntimeClassUnauthorized with default headers values
func NewDeleteNodeV1alpha1CollectionRuntimeClassUnauthorized() *DeleteNodeV1alpha1CollectionRuntimeClassUnauthorized {

	return &DeleteNodeV1alpha1CollectionRuntimeClassUnauthorized{}
}

// WriteResponse to the client
func (o *DeleteNodeV1alpha1CollectionRuntimeClassUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
