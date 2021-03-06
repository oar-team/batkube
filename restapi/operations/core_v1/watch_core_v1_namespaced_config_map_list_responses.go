// Code generated by go-swagger; DO NOT EDIT.

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// WatchCoreV1NamespacedConfigMapListOKCode is the HTTP code returned for type WatchCoreV1NamespacedConfigMapListOK
const WatchCoreV1NamespacedConfigMapListOKCode int = 200

/*WatchCoreV1NamespacedConfigMapListOK OK

swagger:response watchCoreV1NamespacedConfigMapListOK
*/
type WatchCoreV1NamespacedConfigMapListOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent `json:"body,omitempty"`
}

// NewWatchCoreV1NamespacedConfigMapListOK creates WatchCoreV1NamespacedConfigMapListOK with default headers values
func NewWatchCoreV1NamespacedConfigMapListOK() *WatchCoreV1NamespacedConfigMapListOK {

	return &WatchCoreV1NamespacedConfigMapListOK{}
}

// WithPayload adds the payload to the watch core v1 namespaced config map list o k response
func (o *WatchCoreV1NamespacedConfigMapListOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) *WatchCoreV1NamespacedConfigMapListOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the watch core v1 namespaced config map list o k response
func (o *WatchCoreV1NamespacedConfigMapListOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WatchCoreV1NamespacedConfigMapListOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WatchCoreV1NamespacedConfigMapListUnauthorizedCode is the HTTP code returned for type WatchCoreV1NamespacedConfigMapListUnauthorized
const WatchCoreV1NamespacedConfigMapListUnauthorizedCode int = 401

/*WatchCoreV1NamespacedConfigMapListUnauthorized Unauthorized

swagger:response watchCoreV1NamespacedConfigMapListUnauthorized
*/
type WatchCoreV1NamespacedConfigMapListUnauthorized struct {
}

// NewWatchCoreV1NamespacedConfigMapListUnauthorized creates WatchCoreV1NamespacedConfigMapListUnauthorized with default headers values
func NewWatchCoreV1NamespacedConfigMapListUnauthorized() *WatchCoreV1NamespacedConfigMapListUnauthorized {

	return &WatchCoreV1NamespacedConfigMapListUnauthorized{}
}

// WriteResponse to the client
func (o *WatchCoreV1NamespacedConfigMapListUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
