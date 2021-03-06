// Code generated by go-swagger; DO NOT EDIT.

package extensions_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// WatchExtensionsV1beta1IngressListForAllNamespacesOKCode is the HTTP code returned for type WatchExtensionsV1beta1IngressListForAllNamespacesOK
const WatchExtensionsV1beta1IngressListForAllNamespacesOKCode int = 200

/*WatchExtensionsV1beta1IngressListForAllNamespacesOK OK

swagger:response watchExtensionsV1beta1IngressListForAllNamespacesOK
*/
type WatchExtensionsV1beta1IngressListForAllNamespacesOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent `json:"body,omitempty"`
}

// NewWatchExtensionsV1beta1IngressListForAllNamespacesOK creates WatchExtensionsV1beta1IngressListForAllNamespacesOK with default headers values
func NewWatchExtensionsV1beta1IngressListForAllNamespacesOK() *WatchExtensionsV1beta1IngressListForAllNamespacesOK {

	return &WatchExtensionsV1beta1IngressListForAllNamespacesOK{}
}

// WithPayload adds the payload to the watch extensions v1beta1 ingress list for all namespaces o k response
func (o *WatchExtensionsV1beta1IngressListForAllNamespacesOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) *WatchExtensionsV1beta1IngressListForAllNamespacesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the watch extensions v1beta1 ingress list for all namespaces o k response
func (o *WatchExtensionsV1beta1IngressListForAllNamespacesOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WatchExtensionsV1beta1IngressListForAllNamespacesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WatchExtensionsV1beta1IngressListForAllNamespacesUnauthorizedCode is the HTTP code returned for type WatchExtensionsV1beta1IngressListForAllNamespacesUnauthorized
const WatchExtensionsV1beta1IngressListForAllNamespacesUnauthorizedCode int = 401

/*WatchExtensionsV1beta1IngressListForAllNamespacesUnauthorized Unauthorized

swagger:response watchExtensionsV1beta1IngressListForAllNamespacesUnauthorized
*/
type WatchExtensionsV1beta1IngressListForAllNamespacesUnauthorized struct {
}

// NewWatchExtensionsV1beta1IngressListForAllNamespacesUnauthorized creates WatchExtensionsV1beta1IngressListForAllNamespacesUnauthorized with default headers values
func NewWatchExtensionsV1beta1IngressListForAllNamespacesUnauthorized() *WatchExtensionsV1beta1IngressListForAllNamespacesUnauthorized {

	return &WatchExtensionsV1beta1IngressListForAllNamespacesUnauthorized{}
}

// WriteResponse to the client
func (o *WatchExtensionsV1beta1IngressListForAllNamespacesUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
