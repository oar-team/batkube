// Code generated by go-swagger; DO NOT EDIT.

package events_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// ListEventsV1beta1NamespacedEventOKCode is the HTTP code returned for type ListEventsV1beta1NamespacedEventOK
const ListEventsV1beta1NamespacedEventOKCode int = 200

/*ListEventsV1beta1NamespacedEventOK OK

swagger:response listEventsV1beta1NamespacedEventOK
*/
type ListEventsV1beta1NamespacedEventOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIEventsV1beta1EventList `json:"body,omitempty"`
}

// NewListEventsV1beta1NamespacedEventOK creates ListEventsV1beta1NamespacedEventOK with default headers values
func NewListEventsV1beta1NamespacedEventOK() *ListEventsV1beta1NamespacedEventOK {

	return &ListEventsV1beta1NamespacedEventOK{}
}

// WithPayload adds the payload to the list events v1beta1 namespaced event o k response
func (o *ListEventsV1beta1NamespacedEventOK) WithPayload(payload *models.IoK8sAPIEventsV1beta1EventList) *ListEventsV1beta1NamespacedEventOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list events v1beta1 namespaced event o k response
func (o *ListEventsV1beta1NamespacedEventOK) SetPayload(payload *models.IoK8sAPIEventsV1beta1EventList) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListEventsV1beta1NamespacedEventOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListEventsV1beta1NamespacedEventUnauthorizedCode is the HTTP code returned for type ListEventsV1beta1NamespacedEventUnauthorized
const ListEventsV1beta1NamespacedEventUnauthorizedCode int = 401

/*ListEventsV1beta1NamespacedEventUnauthorized Unauthorized

swagger:response listEventsV1beta1NamespacedEventUnauthorized
*/
type ListEventsV1beta1NamespacedEventUnauthorized struct {
}

// NewListEventsV1beta1NamespacedEventUnauthorized creates ListEventsV1beta1NamespacedEventUnauthorized with default headers values
func NewListEventsV1beta1NamespacedEventUnauthorized() *ListEventsV1beta1NamespacedEventUnauthorized {

	return &ListEventsV1beta1NamespacedEventUnauthorized{}
}

// WriteResponse to the client
func (o *ListEventsV1beta1NamespacedEventUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
