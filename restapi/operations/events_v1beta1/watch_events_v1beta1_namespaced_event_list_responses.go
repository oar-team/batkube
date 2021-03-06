// Code generated by go-swagger; DO NOT EDIT.

package events_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// WatchEventsV1beta1NamespacedEventListOKCode is the HTTP code returned for type WatchEventsV1beta1NamespacedEventListOK
const WatchEventsV1beta1NamespacedEventListOKCode int = 200

/*WatchEventsV1beta1NamespacedEventListOK OK

swagger:response watchEventsV1beta1NamespacedEventListOK
*/
type WatchEventsV1beta1NamespacedEventListOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent `json:"body,omitempty"`
}

// NewWatchEventsV1beta1NamespacedEventListOK creates WatchEventsV1beta1NamespacedEventListOK with default headers values
func NewWatchEventsV1beta1NamespacedEventListOK() *WatchEventsV1beta1NamespacedEventListOK {

	return &WatchEventsV1beta1NamespacedEventListOK{}
}

// WithPayload adds the payload to the watch events v1beta1 namespaced event list o k response
func (o *WatchEventsV1beta1NamespacedEventListOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) *WatchEventsV1beta1NamespacedEventListOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the watch events v1beta1 namespaced event list o k response
func (o *WatchEventsV1beta1NamespacedEventListOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WatchEventsV1beta1NamespacedEventListOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WatchEventsV1beta1NamespacedEventListUnauthorizedCode is the HTTP code returned for type WatchEventsV1beta1NamespacedEventListUnauthorized
const WatchEventsV1beta1NamespacedEventListUnauthorizedCode int = 401

/*WatchEventsV1beta1NamespacedEventListUnauthorized Unauthorized

swagger:response watchEventsV1beta1NamespacedEventListUnauthorized
*/
type WatchEventsV1beta1NamespacedEventListUnauthorized struct {
}

// NewWatchEventsV1beta1NamespacedEventListUnauthorized creates WatchEventsV1beta1NamespacedEventListUnauthorized with default headers values
func NewWatchEventsV1beta1NamespacedEventListUnauthorized() *WatchEventsV1beta1NamespacedEventListUnauthorized {

	return &WatchEventsV1beta1NamespacedEventListUnauthorized{}
}

// WriteResponse to the client
func (o *WatchEventsV1beta1NamespacedEventListUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
