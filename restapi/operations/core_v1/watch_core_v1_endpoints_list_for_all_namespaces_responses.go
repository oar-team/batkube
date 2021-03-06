// Code generated by go-swagger; DO NOT EDIT.

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// WatchCoreV1EndpointsListForAllNamespacesOKCode is the HTTP code returned for type WatchCoreV1EndpointsListForAllNamespacesOK
const WatchCoreV1EndpointsListForAllNamespacesOKCode int = 200

/*WatchCoreV1EndpointsListForAllNamespacesOK OK

swagger:response watchCoreV1EndpointsListForAllNamespacesOK
*/
type WatchCoreV1EndpointsListForAllNamespacesOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent `json:"body,omitempty"`
}

// NewWatchCoreV1EndpointsListForAllNamespacesOK creates WatchCoreV1EndpointsListForAllNamespacesOK with default headers values
func NewWatchCoreV1EndpointsListForAllNamespacesOK() *WatchCoreV1EndpointsListForAllNamespacesOK {

	return &WatchCoreV1EndpointsListForAllNamespacesOK{}
}

// WithPayload adds the payload to the watch core v1 endpoints list for all namespaces o k response
func (o *WatchCoreV1EndpointsListForAllNamespacesOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) *WatchCoreV1EndpointsListForAllNamespacesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the watch core v1 endpoints list for all namespaces o k response
func (o *WatchCoreV1EndpointsListForAllNamespacesOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WatchCoreV1EndpointsListForAllNamespacesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WatchCoreV1EndpointsListForAllNamespacesUnauthorizedCode is the HTTP code returned for type WatchCoreV1EndpointsListForAllNamespacesUnauthorized
const WatchCoreV1EndpointsListForAllNamespacesUnauthorizedCode int = 401

/*WatchCoreV1EndpointsListForAllNamespacesUnauthorized Unauthorized

swagger:response watchCoreV1EndpointsListForAllNamespacesUnauthorized
*/
type WatchCoreV1EndpointsListForAllNamespacesUnauthorized struct {
}

// NewWatchCoreV1EndpointsListForAllNamespacesUnauthorized creates WatchCoreV1EndpointsListForAllNamespacesUnauthorized with default headers values
func NewWatchCoreV1EndpointsListForAllNamespacesUnauthorized() *WatchCoreV1EndpointsListForAllNamespacesUnauthorized {

	return &WatchCoreV1EndpointsListForAllNamespacesUnauthorized{}
}

// WriteResponse to the client
func (o *WatchCoreV1EndpointsListForAllNamespacesUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
