// Code generated by go-swagger; DO NOT EDIT.

package coordination_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/models"
)

// DeleteCoordinationV1NamespacedLeaseOKCode is the HTTP code returned for type DeleteCoordinationV1NamespacedLeaseOK
const DeleteCoordinationV1NamespacedLeaseOKCode int = 200

/*DeleteCoordinationV1NamespacedLeaseOK OK

swagger:response deleteCoordinationV1NamespacedLeaseOK
*/
type DeleteCoordinationV1NamespacedLeaseOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1Status `json:"body,omitempty"`
}

// NewDeleteCoordinationV1NamespacedLeaseOK creates DeleteCoordinationV1NamespacedLeaseOK with default headers values
func NewDeleteCoordinationV1NamespacedLeaseOK() *DeleteCoordinationV1NamespacedLeaseOK {

	return &DeleteCoordinationV1NamespacedLeaseOK{}
}

// WithPayload adds the payload to the delete coordination v1 namespaced lease o k response
func (o *DeleteCoordinationV1NamespacedLeaseOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) *DeleteCoordinationV1NamespacedLeaseOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete coordination v1 namespaced lease o k response
func (o *DeleteCoordinationV1NamespacedLeaseOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteCoordinationV1NamespacedLeaseOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteCoordinationV1NamespacedLeaseAcceptedCode is the HTTP code returned for type DeleteCoordinationV1NamespacedLeaseAccepted
const DeleteCoordinationV1NamespacedLeaseAcceptedCode int = 202

/*DeleteCoordinationV1NamespacedLeaseAccepted Accepted

swagger:response deleteCoordinationV1NamespacedLeaseAccepted
*/
type DeleteCoordinationV1NamespacedLeaseAccepted struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1Status `json:"body,omitempty"`
}

// NewDeleteCoordinationV1NamespacedLeaseAccepted creates DeleteCoordinationV1NamespacedLeaseAccepted with default headers values
func NewDeleteCoordinationV1NamespacedLeaseAccepted() *DeleteCoordinationV1NamespacedLeaseAccepted {

	return &DeleteCoordinationV1NamespacedLeaseAccepted{}
}

// WithPayload adds the payload to the delete coordination v1 namespaced lease accepted response
func (o *DeleteCoordinationV1NamespacedLeaseAccepted) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) *DeleteCoordinationV1NamespacedLeaseAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete coordination v1 namespaced lease accepted response
func (o *DeleteCoordinationV1NamespacedLeaseAccepted) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteCoordinationV1NamespacedLeaseAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(202)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteCoordinationV1NamespacedLeaseUnauthorizedCode is the HTTP code returned for type DeleteCoordinationV1NamespacedLeaseUnauthorized
const DeleteCoordinationV1NamespacedLeaseUnauthorizedCode int = 401

/*DeleteCoordinationV1NamespacedLeaseUnauthorized Unauthorized

swagger:response deleteCoordinationV1NamespacedLeaseUnauthorized
*/
type DeleteCoordinationV1NamespacedLeaseUnauthorized struct {
}

// NewDeleteCoordinationV1NamespacedLeaseUnauthorized creates DeleteCoordinationV1NamespacedLeaseUnauthorized with default headers values
func NewDeleteCoordinationV1NamespacedLeaseUnauthorized() *DeleteCoordinationV1NamespacedLeaseUnauthorized {

	return &DeleteCoordinationV1NamespacedLeaseUnauthorized{}
}

// WriteResponse to the client
func (o *DeleteCoordinationV1NamespacedLeaseUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
