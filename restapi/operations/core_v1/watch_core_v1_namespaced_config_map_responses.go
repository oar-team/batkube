// Code generated by go-swagger; DO NOT EDIT.

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// WatchCoreV1NamespacedConfigMapOKCode is the HTTP code returned for type WatchCoreV1NamespacedConfigMapOK
const WatchCoreV1NamespacedConfigMapOKCode int = 200

/*WatchCoreV1NamespacedConfigMapOK OK

swagger:response watchCoreV1NamespacedConfigMapOK
*/
type WatchCoreV1NamespacedConfigMapOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent `json:"body,omitempty"`
}

// NewWatchCoreV1NamespacedConfigMapOK creates WatchCoreV1NamespacedConfigMapOK with default headers values
func NewWatchCoreV1NamespacedConfigMapOK() *WatchCoreV1NamespacedConfigMapOK {

	return &WatchCoreV1NamespacedConfigMapOK{}
}

// WithPayload adds the payload to the watch core v1 namespaced config map o k response
func (o *WatchCoreV1NamespacedConfigMapOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) *WatchCoreV1NamespacedConfigMapOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the watch core v1 namespaced config map o k response
func (o *WatchCoreV1NamespacedConfigMapOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WatchCoreV1NamespacedConfigMapOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WatchCoreV1NamespacedConfigMapUnauthorizedCode is the HTTP code returned for type WatchCoreV1NamespacedConfigMapUnauthorized
const WatchCoreV1NamespacedConfigMapUnauthorizedCode int = 401

/*WatchCoreV1NamespacedConfigMapUnauthorized Unauthorized

swagger:response watchCoreV1NamespacedConfigMapUnauthorized
*/
type WatchCoreV1NamespacedConfigMapUnauthorized struct {
}

// NewWatchCoreV1NamespacedConfigMapUnauthorized creates WatchCoreV1NamespacedConfigMapUnauthorized with default headers values
func NewWatchCoreV1NamespacedConfigMapUnauthorized() *WatchCoreV1NamespacedConfigMapUnauthorized {

	return &WatchCoreV1NamespacedConfigMapUnauthorized{}
}

// WriteResponse to the client
func (o *WatchCoreV1NamespacedConfigMapUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
