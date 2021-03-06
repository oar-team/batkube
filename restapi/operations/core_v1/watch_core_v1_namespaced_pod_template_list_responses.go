// Code generated by go-swagger; DO NOT EDIT.

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// WatchCoreV1NamespacedPodTemplateListOKCode is the HTTP code returned for type WatchCoreV1NamespacedPodTemplateListOK
const WatchCoreV1NamespacedPodTemplateListOKCode int = 200

/*WatchCoreV1NamespacedPodTemplateListOK OK

swagger:response watchCoreV1NamespacedPodTemplateListOK
*/
type WatchCoreV1NamespacedPodTemplateListOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent `json:"body,omitempty"`
}

// NewWatchCoreV1NamespacedPodTemplateListOK creates WatchCoreV1NamespacedPodTemplateListOK with default headers values
func NewWatchCoreV1NamespacedPodTemplateListOK() *WatchCoreV1NamespacedPodTemplateListOK {

	return &WatchCoreV1NamespacedPodTemplateListOK{}
}

// WithPayload adds the payload to the watch core v1 namespaced pod template list o k response
func (o *WatchCoreV1NamespacedPodTemplateListOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) *WatchCoreV1NamespacedPodTemplateListOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the watch core v1 namespaced pod template list o k response
func (o *WatchCoreV1NamespacedPodTemplateListOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WatchCoreV1NamespacedPodTemplateListOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WatchCoreV1NamespacedPodTemplateListUnauthorizedCode is the HTTP code returned for type WatchCoreV1NamespacedPodTemplateListUnauthorized
const WatchCoreV1NamespacedPodTemplateListUnauthorizedCode int = 401

/*WatchCoreV1NamespacedPodTemplateListUnauthorized Unauthorized

swagger:response watchCoreV1NamespacedPodTemplateListUnauthorized
*/
type WatchCoreV1NamespacedPodTemplateListUnauthorized struct {
}

// NewWatchCoreV1NamespacedPodTemplateListUnauthorized creates WatchCoreV1NamespacedPodTemplateListUnauthorized with default headers values
func NewWatchCoreV1NamespacedPodTemplateListUnauthorized() *WatchCoreV1NamespacedPodTemplateListUnauthorized {

	return &WatchCoreV1NamespacedPodTemplateListUnauthorized{}
}

// WriteResponse to the client
func (o *WatchCoreV1NamespacedPodTemplateListUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
