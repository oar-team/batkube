// Code generated by go-swagger; DO NOT EDIT.

package events_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// PatchEventsV1beta1NamespacedEventOKCode is the HTTP code returned for type PatchEventsV1beta1NamespacedEventOK
const PatchEventsV1beta1NamespacedEventOKCode int = 200

/*PatchEventsV1beta1NamespacedEventOK OK

swagger:response patchEventsV1beta1NamespacedEventOK
*/
type PatchEventsV1beta1NamespacedEventOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIEventsV1beta1Event `json:"body,omitempty"`
}

// NewPatchEventsV1beta1NamespacedEventOK creates PatchEventsV1beta1NamespacedEventOK with default headers values
func NewPatchEventsV1beta1NamespacedEventOK() *PatchEventsV1beta1NamespacedEventOK {

	return &PatchEventsV1beta1NamespacedEventOK{}
}

// WithPayload adds the payload to the patch events v1beta1 namespaced event o k response
func (o *PatchEventsV1beta1NamespacedEventOK) WithPayload(payload *models.IoK8sAPIEventsV1beta1Event) *PatchEventsV1beta1NamespacedEventOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch events v1beta1 namespaced event o k response
func (o *PatchEventsV1beta1NamespacedEventOK) SetPayload(payload *models.IoK8sAPIEventsV1beta1Event) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchEventsV1beta1NamespacedEventOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchEventsV1beta1NamespacedEventUnauthorizedCode is the HTTP code returned for type PatchEventsV1beta1NamespacedEventUnauthorized
const PatchEventsV1beta1NamespacedEventUnauthorizedCode int = 401

/*PatchEventsV1beta1NamespacedEventUnauthorized Unauthorized

swagger:response patchEventsV1beta1NamespacedEventUnauthorized
*/
type PatchEventsV1beta1NamespacedEventUnauthorized struct {
}

// NewPatchEventsV1beta1NamespacedEventUnauthorized creates PatchEventsV1beta1NamespacedEventUnauthorized with default headers values
func NewPatchEventsV1beta1NamespacedEventUnauthorized() *PatchEventsV1beta1NamespacedEventUnauthorized {

	return &PatchEventsV1beta1NamespacedEventUnauthorized{}
}

// WriteResponse to the client
func (o *PatchEventsV1beta1NamespacedEventUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
