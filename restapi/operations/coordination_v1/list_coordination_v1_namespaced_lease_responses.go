// Code generated by go-swagger; DO NOT EDIT.

package coordination_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// ListCoordinationV1NamespacedLeaseOKCode is the HTTP code returned for type ListCoordinationV1NamespacedLeaseOK
const ListCoordinationV1NamespacedLeaseOKCode int = 200

/*ListCoordinationV1NamespacedLeaseOK OK

swagger:response listCoordinationV1NamespacedLeaseOK
*/
type ListCoordinationV1NamespacedLeaseOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPICoordinationV1LeaseList `json:"body,omitempty"`
}

// NewListCoordinationV1NamespacedLeaseOK creates ListCoordinationV1NamespacedLeaseOK with default headers values
func NewListCoordinationV1NamespacedLeaseOK() *ListCoordinationV1NamespacedLeaseOK {

	return &ListCoordinationV1NamespacedLeaseOK{}
}

// WithPayload adds the payload to the list coordination v1 namespaced lease o k response
func (o *ListCoordinationV1NamespacedLeaseOK) WithPayload(payload *models.IoK8sAPICoordinationV1LeaseList) *ListCoordinationV1NamespacedLeaseOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list coordination v1 namespaced lease o k response
func (o *ListCoordinationV1NamespacedLeaseOK) SetPayload(payload *models.IoK8sAPICoordinationV1LeaseList) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListCoordinationV1NamespacedLeaseOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListCoordinationV1NamespacedLeaseUnauthorizedCode is the HTTP code returned for type ListCoordinationV1NamespacedLeaseUnauthorized
const ListCoordinationV1NamespacedLeaseUnauthorizedCode int = 401

/*ListCoordinationV1NamespacedLeaseUnauthorized Unauthorized

swagger:response listCoordinationV1NamespacedLeaseUnauthorized
*/
type ListCoordinationV1NamespacedLeaseUnauthorized struct {
}

// NewListCoordinationV1NamespacedLeaseUnauthorized creates ListCoordinationV1NamespacedLeaseUnauthorized with default headers values
func NewListCoordinationV1NamespacedLeaseUnauthorized() *ListCoordinationV1NamespacedLeaseUnauthorized {

	return &ListCoordinationV1NamespacedLeaseUnauthorized{}
}

// WriteResponse to the client
func (o *ListCoordinationV1NamespacedLeaseUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
