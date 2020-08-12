// Code generated by go-swagger; DO NOT EDIT.

package batch_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// WatchBatchV1JobListForAllNamespacesOKCode is the HTTP code returned for type WatchBatchV1JobListForAllNamespacesOK
const WatchBatchV1JobListForAllNamespacesOKCode int = 200

/*WatchBatchV1JobListForAllNamespacesOK OK

swagger:response watchBatchV1JobListForAllNamespacesOK
*/
type WatchBatchV1JobListForAllNamespacesOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent `json:"body,omitempty"`
}

// NewWatchBatchV1JobListForAllNamespacesOK creates WatchBatchV1JobListForAllNamespacesOK with default headers values
func NewWatchBatchV1JobListForAllNamespacesOK() *WatchBatchV1JobListForAllNamespacesOK {

	return &WatchBatchV1JobListForAllNamespacesOK{}
}

// WithPayload adds the payload to the watch batch v1 job list for all namespaces o k response
func (o *WatchBatchV1JobListForAllNamespacesOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) *WatchBatchV1JobListForAllNamespacesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the watch batch v1 job list for all namespaces o k response
func (o *WatchBatchV1JobListForAllNamespacesOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WatchBatchV1JobListForAllNamespacesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WatchBatchV1JobListForAllNamespacesUnauthorizedCode is the HTTP code returned for type WatchBatchV1JobListForAllNamespacesUnauthorized
const WatchBatchV1JobListForAllNamespacesUnauthorizedCode int = 401

/*WatchBatchV1JobListForAllNamespacesUnauthorized Unauthorized

swagger:response watchBatchV1JobListForAllNamespacesUnauthorized
*/
type WatchBatchV1JobListForAllNamespacesUnauthorized struct {
}

// NewWatchBatchV1JobListForAllNamespacesUnauthorized creates WatchBatchV1JobListForAllNamespacesUnauthorized with default headers values
func NewWatchBatchV1JobListForAllNamespacesUnauthorized() *WatchBatchV1JobListForAllNamespacesUnauthorized {

	return &WatchBatchV1JobListForAllNamespacesUnauthorized{}
}

// WriteResponse to the client
func (o *WatchBatchV1JobListForAllNamespacesUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
