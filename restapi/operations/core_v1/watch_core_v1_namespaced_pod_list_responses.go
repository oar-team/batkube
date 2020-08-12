// Code generated by go-swagger; DO NOT EDIT.

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// WatchCoreV1NamespacedPodListOKCode is the HTTP code returned for type WatchCoreV1NamespacedPodListOK
const WatchCoreV1NamespacedPodListOKCode int = 200

/*WatchCoreV1NamespacedPodListOK OK

swagger:response watchCoreV1NamespacedPodListOK
*/
type WatchCoreV1NamespacedPodListOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent `json:"body,omitempty"`
}

// NewWatchCoreV1NamespacedPodListOK creates WatchCoreV1NamespacedPodListOK with default headers values
func NewWatchCoreV1NamespacedPodListOK() *WatchCoreV1NamespacedPodListOK {

	return &WatchCoreV1NamespacedPodListOK{}
}

// WithPayload adds the payload to the watch core v1 namespaced pod list o k response
func (o *WatchCoreV1NamespacedPodListOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) *WatchCoreV1NamespacedPodListOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the watch core v1 namespaced pod list o k response
func (o *WatchCoreV1NamespacedPodListOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WatchCoreV1NamespacedPodListOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WatchCoreV1NamespacedPodListUnauthorizedCode is the HTTP code returned for type WatchCoreV1NamespacedPodListUnauthorized
const WatchCoreV1NamespacedPodListUnauthorizedCode int = 401

/*WatchCoreV1NamespacedPodListUnauthorized Unauthorized

swagger:response watchCoreV1NamespacedPodListUnauthorized
*/
type WatchCoreV1NamespacedPodListUnauthorized struct {
}

// NewWatchCoreV1NamespacedPodListUnauthorized creates WatchCoreV1NamespacedPodListUnauthorized with default headers values
func NewWatchCoreV1NamespacedPodListUnauthorized() *WatchCoreV1NamespacedPodListUnauthorized {

	return &WatchCoreV1NamespacedPodListUnauthorized{}
}

// WriteResponse to the client
func (o *WatchCoreV1NamespacedPodListUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
