// Code generated by go-swagger; DO NOT EDIT.

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// WatchCoreV1PodTemplateListForAllNamespacesOKCode is the HTTP code returned for type WatchCoreV1PodTemplateListForAllNamespacesOK
const WatchCoreV1PodTemplateListForAllNamespacesOKCode int = 200

/*WatchCoreV1PodTemplateListForAllNamespacesOK OK

swagger:response watchCoreV1PodTemplateListForAllNamespacesOK
*/
type WatchCoreV1PodTemplateListForAllNamespacesOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent `json:"body,omitempty"`
}

// NewWatchCoreV1PodTemplateListForAllNamespacesOK creates WatchCoreV1PodTemplateListForAllNamespacesOK with default headers values
func NewWatchCoreV1PodTemplateListForAllNamespacesOK() *WatchCoreV1PodTemplateListForAllNamespacesOK {

	return &WatchCoreV1PodTemplateListForAllNamespacesOK{}
}

// WithPayload adds the payload to the watch core v1 pod template list for all namespaces o k response
func (o *WatchCoreV1PodTemplateListForAllNamespacesOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) *WatchCoreV1PodTemplateListForAllNamespacesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the watch core v1 pod template list for all namespaces o k response
func (o *WatchCoreV1PodTemplateListForAllNamespacesOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WatchCoreV1PodTemplateListForAllNamespacesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WatchCoreV1PodTemplateListForAllNamespacesUnauthorizedCode is the HTTP code returned for type WatchCoreV1PodTemplateListForAllNamespacesUnauthorized
const WatchCoreV1PodTemplateListForAllNamespacesUnauthorizedCode int = 401

/*WatchCoreV1PodTemplateListForAllNamespacesUnauthorized Unauthorized

swagger:response watchCoreV1PodTemplateListForAllNamespacesUnauthorized
*/
type WatchCoreV1PodTemplateListForAllNamespacesUnauthorized struct {
}

// NewWatchCoreV1PodTemplateListForAllNamespacesUnauthorized creates WatchCoreV1PodTemplateListForAllNamespacesUnauthorized with default headers values
func NewWatchCoreV1PodTemplateListForAllNamespacesUnauthorized() *WatchCoreV1PodTemplateListForAllNamespacesUnauthorized {

	return &WatchCoreV1PodTemplateListForAllNamespacesUnauthorized{}
}

// WriteResponse to the client
func (o *WatchCoreV1PodTemplateListForAllNamespacesUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
