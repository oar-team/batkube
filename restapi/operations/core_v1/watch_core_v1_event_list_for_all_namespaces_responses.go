// Code generated by go-swagger; DO NOT EDIT.

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/models"
)

// WatchCoreV1EventListForAllNamespacesOKCode is the HTTP code returned for type WatchCoreV1EventListForAllNamespacesOK
const WatchCoreV1EventListForAllNamespacesOKCode int = 200

/*WatchCoreV1EventListForAllNamespacesOK OK

swagger:response watchCoreV1EventListForAllNamespacesOK
*/
type WatchCoreV1EventListForAllNamespacesOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent `json:"body,omitempty"`
}

// NewWatchCoreV1EventListForAllNamespacesOK creates WatchCoreV1EventListForAllNamespacesOK with default headers values
func NewWatchCoreV1EventListForAllNamespacesOK() *WatchCoreV1EventListForAllNamespacesOK {

	return &WatchCoreV1EventListForAllNamespacesOK{}
}

// WithPayload adds the payload to the watch core v1 event list for all namespaces o k response
func (o *WatchCoreV1EventListForAllNamespacesOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) *WatchCoreV1EventListForAllNamespacesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the watch core v1 event list for all namespaces o k response
func (o *WatchCoreV1EventListForAllNamespacesOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WatchCoreV1EventListForAllNamespacesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WatchCoreV1EventListForAllNamespacesUnauthorizedCode is the HTTP code returned for type WatchCoreV1EventListForAllNamespacesUnauthorized
const WatchCoreV1EventListForAllNamespacesUnauthorizedCode int = 401

/*WatchCoreV1EventListForAllNamespacesUnauthorized Unauthorized

swagger:response watchCoreV1EventListForAllNamespacesUnauthorized
*/
type WatchCoreV1EventListForAllNamespacesUnauthorized struct {
}

// NewWatchCoreV1EventListForAllNamespacesUnauthorized creates WatchCoreV1EventListForAllNamespacesUnauthorized with default headers values
func NewWatchCoreV1EventListForAllNamespacesUnauthorized() *WatchCoreV1EventListForAllNamespacesUnauthorized {

	return &WatchCoreV1EventListForAllNamespacesUnauthorized{}
}

// WriteResponse to the client
func (o *WatchCoreV1EventListForAllNamespacesUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
