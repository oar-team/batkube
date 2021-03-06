// Code generated by go-swagger; DO NOT EDIT.

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// WatchCoreV1NamespacedPersistentVolumeClaimListOKCode is the HTTP code returned for type WatchCoreV1NamespacedPersistentVolumeClaimListOK
const WatchCoreV1NamespacedPersistentVolumeClaimListOKCode int = 200

/*WatchCoreV1NamespacedPersistentVolumeClaimListOK OK

swagger:response watchCoreV1NamespacedPersistentVolumeClaimListOK
*/
type WatchCoreV1NamespacedPersistentVolumeClaimListOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent `json:"body,omitempty"`
}

// NewWatchCoreV1NamespacedPersistentVolumeClaimListOK creates WatchCoreV1NamespacedPersistentVolumeClaimListOK with default headers values
func NewWatchCoreV1NamespacedPersistentVolumeClaimListOK() *WatchCoreV1NamespacedPersistentVolumeClaimListOK {

	return &WatchCoreV1NamespacedPersistentVolumeClaimListOK{}
}

// WithPayload adds the payload to the watch core v1 namespaced persistent volume claim list o k response
func (o *WatchCoreV1NamespacedPersistentVolumeClaimListOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) *WatchCoreV1NamespacedPersistentVolumeClaimListOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the watch core v1 namespaced persistent volume claim list o k response
func (o *WatchCoreV1NamespacedPersistentVolumeClaimListOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WatchCoreV1NamespacedPersistentVolumeClaimListOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WatchCoreV1NamespacedPersistentVolumeClaimListUnauthorizedCode is the HTTP code returned for type WatchCoreV1NamespacedPersistentVolumeClaimListUnauthorized
const WatchCoreV1NamespacedPersistentVolumeClaimListUnauthorizedCode int = 401

/*WatchCoreV1NamespacedPersistentVolumeClaimListUnauthorized Unauthorized

swagger:response watchCoreV1NamespacedPersistentVolumeClaimListUnauthorized
*/
type WatchCoreV1NamespacedPersistentVolumeClaimListUnauthorized struct {
}

// NewWatchCoreV1NamespacedPersistentVolumeClaimListUnauthorized creates WatchCoreV1NamespacedPersistentVolumeClaimListUnauthorized with default headers values
func NewWatchCoreV1NamespacedPersistentVolumeClaimListUnauthorized() *WatchCoreV1NamespacedPersistentVolumeClaimListUnauthorized {

	return &WatchCoreV1NamespacedPersistentVolumeClaimListUnauthorized{}
}

// WriteResponse to the client
func (o *WatchCoreV1NamespacedPersistentVolumeClaimListUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
