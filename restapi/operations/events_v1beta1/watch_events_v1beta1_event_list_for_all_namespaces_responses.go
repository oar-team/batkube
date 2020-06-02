// Code generated by go-swagger; DO NOT EDIT.

package events_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/models"
)

// WatchEventsV1beta1EventListForAllNamespacesOKCode is the HTTP code returned for type WatchEventsV1beta1EventListForAllNamespacesOK
const WatchEventsV1beta1EventListForAllNamespacesOKCode int = 200

/*WatchEventsV1beta1EventListForAllNamespacesOK OK

swagger:response watchEventsV1beta1EventListForAllNamespacesOK
*/
type WatchEventsV1beta1EventListForAllNamespacesOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent `json:"body,omitempty"`
}

// NewWatchEventsV1beta1EventListForAllNamespacesOK creates WatchEventsV1beta1EventListForAllNamespacesOK with default headers values
func NewWatchEventsV1beta1EventListForAllNamespacesOK() *WatchEventsV1beta1EventListForAllNamespacesOK {

	return &WatchEventsV1beta1EventListForAllNamespacesOK{}
}

// WithPayload adds the payload to the watch events v1beta1 event list for all namespaces o k response
func (o *WatchEventsV1beta1EventListForAllNamespacesOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) *WatchEventsV1beta1EventListForAllNamespacesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the watch events v1beta1 event list for all namespaces o k response
func (o *WatchEventsV1beta1EventListForAllNamespacesOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WatchEventsV1beta1EventListForAllNamespacesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WatchEventsV1beta1EventListForAllNamespacesUnauthorizedCode is the HTTP code returned for type WatchEventsV1beta1EventListForAllNamespacesUnauthorized
const WatchEventsV1beta1EventListForAllNamespacesUnauthorizedCode int = 401

/*WatchEventsV1beta1EventListForAllNamespacesUnauthorized Unauthorized

swagger:response watchEventsV1beta1EventListForAllNamespacesUnauthorized
*/
type WatchEventsV1beta1EventListForAllNamespacesUnauthorized struct {
}

// NewWatchEventsV1beta1EventListForAllNamespacesUnauthorized creates WatchEventsV1beta1EventListForAllNamespacesUnauthorized with default headers values
func NewWatchEventsV1beta1EventListForAllNamespacesUnauthorized() *WatchEventsV1beta1EventListForAllNamespacesUnauthorized {

	return &WatchEventsV1beta1EventListForAllNamespacesUnauthorized{}
}

// WriteResponse to the client
func (o *WatchEventsV1beta1EventListForAllNamespacesUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
