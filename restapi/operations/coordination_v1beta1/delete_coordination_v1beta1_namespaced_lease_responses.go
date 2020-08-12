// Code generated by go-swagger; DO NOT EDIT.

package coordination_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// DeleteCoordinationV1beta1NamespacedLeaseOKCode is the HTTP code returned for type DeleteCoordinationV1beta1NamespacedLeaseOK
const DeleteCoordinationV1beta1NamespacedLeaseOKCode int = 200

/*DeleteCoordinationV1beta1NamespacedLeaseOK OK

swagger:response deleteCoordinationV1beta1NamespacedLeaseOK
*/
type DeleteCoordinationV1beta1NamespacedLeaseOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1Status `json:"body,omitempty"`
}

// NewDeleteCoordinationV1beta1NamespacedLeaseOK creates DeleteCoordinationV1beta1NamespacedLeaseOK with default headers values
func NewDeleteCoordinationV1beta1NamespacedLeaseOK() *DeleteCoordinationV1beta1NamespacedLeaseOK {

	return &DeleteCoordinationV1beta1NamespacedLeaseOK{}
}

// WithPayload adds the payload to the delete coordination v1beta1 namespaced lease o k response
func (o *DeleteCoordinationV1beta1NamespacedLeaseOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) *DeleteCoordinationV1beta1NamespacedLeaseOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete coordination v1beta1 namespaced lease o k response
func (o *DeleteCoordinationV1beta1NamespacedLeaseOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteCoordinationV1beta1NamespacedLeaseOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteCoordinationV1beta1NamespacedLeaseAcceptedCode is the HTTP code returned for type DeleteCoordinationV1beta1NamespacedLeaseAccepted
const DeleteCoordinationV1beta1NamespacedLeaseAcceptedCode int = 202

/*DeleteCoordinationV1beta1NamespacedLeaseAccepted Accepted

swagger:response deleteCoordinationV1beta1NamespacedLeaseAccepted
*/
type DeleteCoordinationV1beta1NamespacedLeaseAccepted struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1Status `json:"body,omitempty"`
}

// NewDeleteCoordinationV1beta1NamespacedLeaseAccepted creates DeleteCoordinationV1beta1NamespacedLeaseAccepted with default headers values
func NewDeleteCoordinationV1beta1NamespacedLeaseAccepted() *DeleteCoordinationV1beta1NamespacedLeaseAccepted {

	return &DeleteCoordinationV1beta1NamespacedLeaseAccepted{}
}

// WithPayload adds the payload to the delete coordination v1beta1 namespaced lease accepted response
func (o *DeleteCoordinationV1beta1NamespacedLeaseAccepted) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) *DeleteCoordinationV1beta1NamespacedLeaseAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete coordination v1beta1 namespaced lease accepted response
func (o *DeleteCoordinationV1beta1NamespacedLeaseAccepted) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteCoordinationV1beta1NamespacedLeaseAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(202)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteCoordinationV1beta1NamespacedLeaseUnauthorizedCode is the HTTP code returned for type DeleteCoordinationV1beta1NamespacedLeaseUnauthorized
const DeleteCoordinationV1beta1NamespacedLeaseUnauthorizedCode int = 401

/*DeleteCoordinationV1beta1NamespacedLeaseUnauthorized Unauthorized

swagger:response deleteCoordinationV1beta1NamespacedLeaseUnauthorized
*/
type DeleteCoordinationV1beta1NamespacedLeaseUnauthorized struct {
}

// NewDeleteCoordinationV1beta1NamespacedLeaseUnauthorized creates DeleteCoordinationV1beta1NamespacedLeaseUnauthorized with default headers values
func NewDeleteCoordinationV1beta1NamespacedLeaseUnauthorized() *DeleteCoordinationV1beta1NamespacedLeaseUnauthorized {

	return &DeleteCoordinationV1beta1NamespacedLeaseUnauthorized{}
}

// WriteResponse to the client
func (o *DeleteCoordinationV1beta1NamespacedLeaseUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
