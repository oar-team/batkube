// Code generated by go-swagger; DO NOT EDIT.

package extensions_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// WatchExtensionsV1beta1NamespacedIngressOKCode is the HTTP code returned for type WatchExtensionsV1beta1NamespacedIngressOK
const WatchExtensionsV1beta1NamespacedIngressOKCode int = 200

/*WatchExtensionsV1beta1NamespacedIngressOK OK

swagger:response watchExtensionsV1beta1NamespacedIngressOK
*/
type WatchExtensionsV1beta1NamespacedIngressOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent `json:"body,omitempty"`
}

// NewWatchExtensionsV1beta1NamespacedIngressOK creates WatchExtensionsV1beta1NamespacedIngressOK with default headers values
func NewWatchExtensionsV1beta1NamespacedIngressOK() *WatchExtensionsV1beta1NamespacedIngressOK {

	return &WatchExtensionsV1beta1NamespacedIngressOK{}
}

// WithPayload adds the payload to the watch extensions v1beta1 namespaced ingress o k response
func (o *WatchExtensionsV1beta1NamespacedIngressOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) *WatchExtensionsV1beta1NamespacedIngressOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the watch extensions v1beta1 namespaced ingress o k response
func (o *WatchExtensionsV1beta1NamespacedIngressOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WatchExtensionsV1beta1NamespacedIngressOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WatchExtensionsV1beta1NamespacedIngressUnauthorizedCode is the HTTP code returned for type WatchExtensionsV1beta1NamespacedIngressUnauthorized
const WatchExtensionsV1beta1NamespacedIngressUnauthorizedCode int = 401

/*WatchExtensionsV1beta1NamespacedIngressUnauthorized Unauthorized

swagger:response watchExtensionsV1beta1NamespacedIngressUnauthorized
*/
type WatchExtensionsV1beta1NamespacedIngressUnauthorized struct {
}

// NewWatchExtensionsV1beta1NamespacedIngressUnauthorized creates WatchExtensionsV1beta1NamespacedIngressUnauthorized with default headers values
func NewWatchExtensionsV1beta1NamespacedIngressUnauthorized() *WatchExtensionsV1beta1NamespacedIngressUnauthorized {

	return &WatchExtensionsV1beta1NamespacedIngressUnauthorized{}
}

// WriteResponse to the client
func (o *WatchExtensionsV1beta1NamespacedIngressUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
