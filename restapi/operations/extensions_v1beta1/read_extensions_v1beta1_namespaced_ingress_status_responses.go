// Code generated by go-swagger; DO NOT EDIT.

package extensions_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// ReadExtensionsV1beta1NamespacedIngressStatusOKCode is the HTTP code returned for type ReadExtensionsV1beta1NamespacedIngressStatusOK
const ReadExtensionsV1beta1NamespacedIngressStatusOKCode int = 200

/*ReadExtensionsV1beta1NamespacedIngressStatusOK OK

swagger:response readExtensionsV1beta1NamespacedIngressStatusOK
*/
type ReadExtensionsV1beta1NamespacedIngressStatusOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIExtensionsV1beta1Ingress `json:"body,omitempty"`
}

// NewReadExtensionsV1beta1NamespacedIngressStatusOK creates ReadExtensionsV1beta1NamespacedIngressStatusOK with default headers values
func NewReadExtensionsV1beta1NamespacedIngressStatusOK() *ReadExtensionsV1beta1NamespacedIngressStatusOK {

	return &ReadExtensionsV1beta1NamespacedIngressStatusOK{}
}

// WithPayload adds the payload to the read extensions v1beta1 namespaced ingress status o k response
func (o *ReadExtensionsV1beta1NamespacedIngressStatusOK) WithPayload(payload *models.IoK8sAPIExtensionsV1beta1Ingress) *ReadExtensionsV1beta1NamespacedIngressStatusOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the read extensions v1beta1 namespaced ingress status o k response
func (o *ReadExtensionsV1beta1NamespacedIngressStatusOK) SetPayload(payload *models.IoK8sAPIExtensionsV1beta1Ingress) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReadExtensionsV1beta1NamespacedIngressStatusOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReadExtensionsV1beta1NamespacedIngressStatusUnauthorizedCode is the HTTP code returned for type ReadExtensionsV1beta1NamespacedIngressStatusUnauthorized
const ReadExtensionsV1beta1NamespacedIngressStatusUnauthorizedCode int = 401

/*ReadExtensionsV1beta1NamespacedIngressStatusUnauthorized Unauthorized

swagger:response readExtensionsV1beta1NamespacedIngressStatusUnauthorized
*/
type ReadExtensionsV1beta1NamespacedIngressStatusUnauthorized struct {
}

// NewReadExtensionsV1beta1NamespacedIngressStatusUnauthorized creates ReadExtensionsV1beta1NamespacedIngressStatusUnauthorized with default headers values
func NewReadExtensionsV1beta1NamespacedIngressStatusUnauthorized() *ReadExtensionsV1beta1NamespacedIngressStatusUnauthorized {

	return &ReadExtensionsV1beta1NamespacedIngressStatusUnauthorized{}
}

// WriteResponse to the client
func (o *ReadExtensionsV1beta1NamespacedIngressStatusUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
