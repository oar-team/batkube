// Code generated by go-swagger; DO NOT EDIT.

package coordination_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// DeleteCoordinationV1CollectionNamespacedLeaseOKCode is the HTTP code returned for type DeleteCoordinationV1CollectionNamespacedLeaseOK
const DeleteCoordinationV1CollectionNamespacedLeaseOKCode int = 200

/*DeleteCoordinationV1CollectionNamespacedLeaseOK OK

swagger:response deleteCoordinationV1CollectionNamespacedLeaseOK
*/
type DeleteCoordinationV1CollectionNamespacedLeaseOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1Status `json:"body,omitempty"`
}

// NewDeleteCoordinationV1CollectionNamespacedLeaseOK creates DeleteCoordinationV1CollectionNamespacedLeaseOK with default headers values
func NewDeleteCoordinationV1CollectionNamespacedLeaseOK() *DeleteCoordinationV1CollectionNamespacedLeaseOK {

	return &DeleteCoordinationV1CollectionNamespacedLeaseOK{}
}

// WithPayload adds the payload to the delete coordination v1 collection namespaced lease o k response
func (o *DeleteCoordinationV1CollectionNamespacedLeaseOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) *DeleteCoordinationV1CollectionNamespacedLeaseOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete coordination v1 collection namespaced lease o k response
func (o *DeleteCoordinationV1CollectionNamespacedLeaseOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteCoordinationV1CollectionNamespacedLeaseOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteCoordinationV1CollectionNamespacedLeaseUnauthorizedCode is the HTTP code returned for type DeleteCoordinationV1CollectionNamespacedLeaseUnauthorized
const DeleteCoordinationV1CollectionNamespacedLeaseUnauthorizedCode int = 401

/*DeleteCoordinationV1CollectionNamespacedLeaseUnauthorized Unauthorized

swagger:response deleteCoordinationV1CollectionNamespacedLeaseUnauthorized
*/
type DeleteCoordinationV1CollectionNamespacedLeaseUnauthorized struct {
}

// NewDeleteCoordinationV1CollectionNamespacedLeaseUnauthorized creates DeleteCoordinationV1CollectionNamespacedLeaseUnauthorized with default headers values
func NewDeleteCoordinationV1CollectionNamespacedLeaseUnauthorized() *DeleteCoordinationV1CollectionNamespacedLeaseUnauthorized {

	return &DeleteCoordinationV1CollectionNamespacedLeaseUnauthorized{}
}

// WriteResponse to the client
func (o *DeleteCoordinationV1CollectionNamespacedLeaseUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
