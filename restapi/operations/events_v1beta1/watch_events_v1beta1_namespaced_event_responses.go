// Code generated by go-swagger; DO NOT EDIT.

package events_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/models"
)

// WatchEventsV1beta1NamespacedEventOKCode is the HTTP code returned for type WatchEventsV1beta1NamespacedEventOK
const WatchEventsV1beta1NamespacedEventOKCode int = 200

/*WatchEventsV1beta1NamespacedEventOK OK

swagger:response watchEventsV1beta1NamespacedEventOK
*/
type WatchEventsV1beta1NamespacedEventOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent `json:"body,omitempty"`
}

// NewWatchEventsV1beta1NamespacedEventOK creates WatchEventsV1beta1NamespacedEventOK with default headers values
func NewWatchEventsV1beta1NamespacedEventOK() *WatchEventsV1beta1NamespacedEventOK {

	return &WatchEventsV1beta1NamespacedEventOK{}
}

// WithPayload adds the payload to the watch events v1beta1 namespaced event o k response
func (o *WatchEventsV1beta1NamespacedEventOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) *WatchEventsV1beta1NamespacedEventOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the watch events v1beta1 namespaced event o k response
func (o *WatchEventsV1beta1NamespacedEventOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WatchEventsV1beta1NamespacedEventOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WatchEventsV1beta1NamespacedEventUnauthorizedCode is the HTTP code returned for type WatchEventsV1beta1NamespacedEventUnauthorized
const WatchEventsV1beta1NamespacedEventUnauthorizedCode int = 401

/*WatchEventsV1beta1NamespacedEventUnauthorized Unauthorized

swagger:response watchEventsV1beta1NamespacedEventUnauthorized
*/
type WatchEventsV1beta1NamespacedEventUnauthorized struct {
}

// NewWatchEventsV1beta1NamespacedEventUnauthorized creates WatchEventsV1beta1NamespacedEventUnauthorized with default headers values
func NewWatchEventsV1beta1NamespacedEventUnauthorized() *WatchEventsV1beta1NamespacedEventUnauthorized {

	return &WatchEventsV1beta1NamespacedEventUnauthorized{}
}

// WriteResponse to the client
func (o *WatchEventsV1beta1NamespacedEventUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
