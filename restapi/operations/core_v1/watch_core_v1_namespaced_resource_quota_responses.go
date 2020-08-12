// Code generated by go-swagger; DO NOT EDIT.

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// WatchCoreV1NamespacedResourceQuotaOKCode is the HTTP code returned for type WatchCoreV1NamespacedResourceQuotaOK
const WatchCoreV1NamespacedResourceQuotaOKCode int = 200

/*WatchCoreV1NamespacedResourceQuotaOK OK

swagger:response watchCoreV1NamespacedResourceQuotaOK
*/
type WatchCoreV1NamespacedResourceQuotaOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent `json:"body,omitempty"`
}

// NewWatchCoreV1NamespacedResourceQuotaOK creates WatchCoreV1NamespacedResourceQuotaOK with default headers values
func NewWatchCoreV1NamespacedResourceQuotaOK() *WatchCoreV1NamespacedResourceQuotaOK {

	return &WatchCoreV1NamespacedResourceQuotaOK{}
}

// WithPayload adds the payload to the watch core v1 namespaced resource quota o k response
func (o *WatchCoreV1NamespacedResourceQuotaOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) *WatchCoreV1NamespacedResourceQuotaOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the watch core v1 namespaced resource quota o k response
func (o *WatchCoreV1NamespacedResourceQuotaOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WatchCoreV1NamespacedResourceQuotaOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WatchCoreV1NamespacedResourceQuotaUnauthorizedCode is the HTTP code returned for type WatchCoreV1NamespacedResourceQuotaUnauthorized
const WatchCoreV1NamespacedResourceQuotaUnauthorizedCode int = 401

/*WatchCoreV1NamespacedResourceQuotaUnauthorized Unauthorized

swagger:response watchCoreV1NamespacedResourceQuotaUnauthorized
*/
type WatchCoreV1NamespacedResourceQuotaUnauthorized struct {
}

// NewWatchCoreV1NamespacedResourceQuotaUnauthorized creates WatchCoreV1NamespacedResourceQuotaUnauthorized with default headers values
func NewWatchCoreV1NamespacedResourceQuotaUnauthorized() *WatchCoreV1NamespacedResourceQuotaUnauthorized {

	return &WatchCoreV1NamespacedResourceQuotaUnauthorized{}
}

// WriteResponse to the client
func (o *WatchCoreV1NamespacedResourceQuotaUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
