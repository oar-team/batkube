// Code generated by go-swagger; DO NOT EDIT.

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// WatchCoreV1NamespacedServiceListOKCode is the HTTP code returned for type WatchCoreV1NamespacedServiceListOK
const WatchCoreV1NamespacedServiceListOKCode int = 200

/*WatchCoreV1NamespacedServiceListOK OK

swagger:response watchCoreV1NamespacedServiceListOK
*/
type WatchCoreV1NamespacedServiceListOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent `json:"body,omitempty"`
}

// NewWatchCoreV1NamespacedServiceListOK creates WatchCoreV1NamespacedServiceListOK with default headers values
func NewWatchCoreV1NamespacedServiceListOK() *WatchCoreV1NamespacedServiceListOK {

	return &WatchCoreV1NamespacedServiceListOK{}
}

// WithPayload adds the payload to the watch core v1 namespaced service list o k response
func (o *WatchCoreV1NamespacedServiceListOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) *WatchCoreV1NamespacedServiceListOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the watch core v1 namespaced service list o k response
func (o *WatchCoreV1NamespacedServiceListOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WatchCoreV1NamespacedServiceListOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WatchCoreV1NamespacedServiceListUnauthorizedCode is the HTTP code returned for type WatchCoreV1NamespacedServiceListUnauthorized
const WatchCoreV1NamespacedServiceListUnauthorizedCode int = 401

/*WatchCoreV1NamespacedServiceListUnauthorized Unauthorized

swagger:response watchCoreV1NamespacedServiceListUnauthorized
*/
type WatchCoreV1NamespacedServiceListUnauthorized struct {
}

// NewWatchCoreV1NamespacedServiceListUnauthorized creates WatchCoreV1NamespacedServiceListUnauthorized with default headers values
func NewWatchCoreV1NamespacedServiceListUnauthorized() *WatchCoreV1NamespacedServiceListUnauthorized {

	return &WatchCoreV1NamespacedServiceListUnauthorized{}
}

// WriteResponse to the client
func (o *WatchCoreV1NamespacedServiceListUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
