// Code generated by go-swagger; DO NOT EDIT.

package extensions_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/models"
)

// WatchExtensionsV1beta1NamespacedIngressListOKCode is the HTTP code returned for type WatchExtensionsV1beta1NamespacedIngressListOK
const WatchExtensionsV1beta1NamespacedIngressListOKCode int = 200

/*WatchExtensionsV1beta1NamespacedIngressListOK OK

swagger:response watchExtensionsV1beta1NamespacedIngressListOK
*/
type WatchExtensionsV1beta1NamespacedIngressListOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent `json:"body,omitempty"`
}

// NewWatchExtensionsV1beta1NamespacedIngressListOK creates WatchExtensionsV1beta1NamespacedIngressListOK with default headers values
func NewWatchExtensionsV1beta1NamespacedIngressListOK() *WatchExtensionsV1beta1NamespacedIngressListOK {

	return &WatchExtensionsV1beta1NamespacedIngressListOK{}
}

// WithPayload adds the payload to the watch extensions v1beta1 namespaced ingress list o k response
func (o *WatchExtensionsV1beta1NamespacedIngressListOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) *WatchExtensionsV1beta1NamespacedIngressListOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the watch extensions v1beta1 namespaced ingress list o k response
func (o *WatchExtensionsV1beta1NamespacedIngressListOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WatchExtensionsV1beta1NamespacedIngressListOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WatchExtensionsV1beta1NamespacedIngressListUnauthorizedCode is the HTTP code returned for type WatchExtensionsV1beta1NamespacedIngressListUnauthorized
const WatchExtensionsV1beta1NamespacedIngressListUnauthorizedCode int = 401

/*WatchExtensionsV1beta1NamespacedIngressListUnauthorized Unauthorized

swagger:response watchExtensionsV1beta1NamespacedIngressListUnauthorized
*/
type WatchExtensionsV1beta1NamespacedIngressListUnauthorized struct {
}

// NewWatchExtensionsV1beta1NamespacedIngressListUnauthorized creates WatchExtensionsV1beta1NamespacedIngressListUnauthorized with default headers values
func NewWatchExtensionsV1beta1NamespacedIngressListUnauthorized() *WatchExtensionsV1beta1NamespacedIngressListUnauthorized {

	return &WatchExtensionsV1beta1NamespacedIngressListUnauthorized{}
}

// WriteResponse to the client
func (o *WatchExtensionsV1beta1NamespacedIngressListUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
