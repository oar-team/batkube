// Code generated by go-swagger; DO NOT EDIT.

package extensions_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// ReplaceExtensionsV1beta1NamespacedIngressStatusOKCode is the HTTP code returned for type ReplaceExtensionsV1beta1NamespacedIngressStatusOK
const ReplaceExtensionsV1beta1NamespacedIngressStatusOKCode int = 200

/*ReplaceExtensionsV1beta1NamespacedIngressStatusOK OK

swagger:response replaceExtensionsV1beta1NamespacedIngressStatusOK
*/
type ReplaceExtensionsV1beta1NamespacedIngressStatusOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIExtensionsV1beta1Ingress `json:"body,omitempty"`
}

// NewReplaceExtensionsV1beta1NamespacedIngressStatusOK creates ReplaceExtensionsV1beta1NamespacedIngressStatusOK with default headers values
func NewReplaceExtensionsV1beta1NamespacedIngressStatusOK() *ReplaceExtensionsV1beta1NamespacedIngressStatusOK {

	return &ReplaceExtensionsV1beta1NamespacedIngressStatusOK{}
}

// WithPayload adds the payload to the replace extensions v1beta1 namespaced ingress status o k response
func (o *ReplaceExtensionsV1beta1NamespacedIngressStatusOK) WithPayload(payload *models.IoK8sAPIExtensionsV1beta1Ingress) *ReplaceExtensionsV1beta1NamespacedIngressStatusOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace extensions v1beta1 namespaced ingress status o k response
func (o *ReplaceExtensionsV1beta1NamespacedIngressStatusOK) SetPayload(payload *models.IoK8sAPIExtensionsV1beta1Ingress) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceExtensionsV1beta1NamespacedIngressStatusOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceExtensionsV1beta1NamespacedIngressStatusCreatedCode is the HTTP code returned for type ReplaceExtensionsV1beta1NamespacedIngressStatusCreated
const ReplaceExtensionsV1beta1NamespacedIngressStatusCreatedCode int = 201

/*ReplaceExtensionsV1beta1NamespacedIngressStatusCreated Created

swagger:response replaceExtensionsV1beta1NamespacedIngressStatusCreated
*/
type ReplaceExtensionsV1beta1NamespacedIngressStatusCreated struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIExtensionsV1beta1Ingress `json:"body,omitempty"`
}

// NewReplaceExtensionsV1beta1NamespacedIngressStatusCreated creates ReplaceExtensionsV1beta1NamespacedIngressStatusCreated with default headers values
func NewReplaceExtensionsV1beta1NamespacedIngressStatusCreated() *ReplaceExtensionsV1beta1NamespacedIngressStatusCreated {

	return &ReplaceExtensionsV1beta1NamespacedIngressStatusCreated{}
}

// WithPayload adds the payload to the replace extensions v1beta1 namespaced ingress status created response
func (o *ReplaceExtensionsV1beta1NamespacedIngressStatusCreated) WithPayload(payload *models.IoK8sAPIExtensionsV1beta1Ingress) *ReplaceExtensionsV1beta1NamespacedIngressStatusCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace extensions v1beta1 namespaced ingress status created response
func (o *ReplaceExtensionsV1beta1NamespacedIngressStatusCreated) SetPayload(payload *models.IoK8sAPIExtensionsV1beta1Ingress) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceExtensionsV1beta1NamespacedIngressStatusCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceExtensionsV1beta1NamespacedIngressStatusUnauthorizedCode is the HTTP code returned for type ReplaceExtensionsV1beta1NamespacedIngressStatusUnauthorized
const ReplaceExtensionsV1beta1NamespacedIngressStatusUnauthorizedCode int = 401

/*ReplaceExtensionsV1beta1NamespacedIngressStatusUnauthorized Unauthorized

swagger:response replaceExtensionsV1beta1NamespacedIngressStatusUnauthorized
*/
type ReplaceExtensionsV1beta1NamespacedIngressStatusUnauthorized struct {
}

// NewReplaceExtensionsV1beta1NamespacedIngressStatusUnauthorized creates ReplaceExtensionsV1beta1NamespacedIngressStatusUnauthorized with default headers values
func NewReplaceExtensionsV1beta1NamespacedIngressStatusUnauthorized() *ReplaceExtensionsV1beta1NamespacedIngressStatusUnauthorized {

	return &ReplaceExtensionsV1beta1NamespacedIngressStatusUnauthorized{}
}

// WriteResponse to the client
func (o *ReplaceExtensionsV1beta1NamespacedIngressStatusUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
