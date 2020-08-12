// Code generated by go-swagger; DO NOT EDIT.

package node_v1alpha1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// WatchNodeV1alpha1RuntimeClassOKCode is the HTTP code returned for type WatchNodeV1alpha1RuntimeClassOK
const WatchNodeV1alpha1RuntimeClassOKCode int = 200

/*WatchNodeV1alpha1RuntimeClassOK OK

swagger:response watchNodeV1alpha1RuntimeClassOK
*/
type WatchNodeV1alpha1RuntimeClassOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent `json:"body,omitempty"`
}

// NewWatchNodeV1alpha1RuntimeClassOK creates WatchNodeV1alpha1RuntimeClassOK with default headers values
func NewWatchNodeV1alpha1RuntimeClassOK() *WatchNodeV1alpha1RuntimeClassOK {

	return &WatchNodeV1alpha1RuntimeClassOK{}
}

// WithPayload adds the payload to the watch node v1alpha1 runtime class o k response
func (o *WatchNodeV1alpha1RuntimeClassOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) *WatchNodeV1alpha1RuntimeClassOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the watch node v1alpha1 runtime class o k response
func (o *WatchNodeV1alpha1RuntimeClassOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WatchNodeV1alpha1RuntimeClassOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WatchNodeV1alpha1RuntimeClassUnauthorizedCode is the HTTP code returned for type WatchNodeV1alpha1RuntimeClassUnauthorized
const WatchNodeV1alpha1RuntimeClassUnauthorizedCode int = 401

/*WatchNodeV1alpha1RuntimeClassUnauthorized Unauthorized

swagger:response watchNodeV1alpha1RuntimeClassUnauthorized
*/
type WatchNodeV1alpha1RuntimeClassUnauthorized struct {
}

// NewWatchNodeV1alpha1RuntimeClassUnauthorized creates WatchNodeV1alpha1RuntimeClassUnauthorized with default headers values
func NewWatchNodeV1alpha1RuntimeClassUnauthorized() *WatchNodeV1alpha1RuntimeClassUnauthorized {

	return &WatchNodeV1alpha1RuntimeClassUnauthorized{}
}

// WriteResponse to the client
func (o *WatchNodeV1alpha1RuntimeClassUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
