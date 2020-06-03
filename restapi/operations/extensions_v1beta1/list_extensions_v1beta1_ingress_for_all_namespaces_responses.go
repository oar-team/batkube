// Code generated by go-swagger; DO NOT EDIT.

package extensions_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/models"
)

// ListExtensionsV1beta1IngressForAllNamespacesOKCode is the HTTP code returned for type ListExtensionsV1beta1IngressForAllNamespacesOK
const ListExtensionsV1beta1IngressForAllNamespacesOKCode int = 200

/*ListExtensionsV1beta1IngressForAllNamespacesOK OK

swagger:response listExtensionsV1beta1IngressForAllNamespacesOK
*/
type ListExtensionsV1beta1IngressForAllNamespacesOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIExtensionsV1beta1IngressList `json:"body,omitempty"`
}

// NewListExtensionsV1beta1IngressForAllNamespacesOK creates ListExtensionsV1beta1IngressForAllNamespacesOK with default headers values
func NewListExtensionsV1beta1IngressForAllNamespacesOK() *ListExtensionsV1beta1IngressForAllNamespacesOK {

	return &ListExtensionsV1beta1IngressForAllNamespacesOK{}
}

// WithPayload adds the payload to the list extensions v1beta1 ingress for all namespaces o k response
func (o *ListExtensionsV1beta1IngressForAllNamespacesOK) WithPayload(payload *models.IoK8sAPIExtensionsV1beta1IngressList) *ListExtensionsV1beta1IngressForAllNamespacesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list extensions v1beta1 ingress for all namespaces o k response
func (o *ListExtensionsV1beta1IngressForAllNamespacesOK) SetPayload(payload *models.IoK8sAPIExtensionsV1beta1IngressList) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListExtensionsV1beta1IngressForAllNamespacesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListExtensionsV1beta1IngressForAllNamespacesUnauthorizedCode is the HTTP code returned for type ListExtensionsV1beta1IngressForAllNamespacesUnauthorized
const ListExtensionsV1beta1IngressForAllNamespacesUnauthorizedCode int = 401

/*ListExtensionsV1beta1IngressForAllNamespacesUnauthorized Unauthorized

swagger:response listExtensionsV1beta1IngressForAllNamespacesUnauthorized
*/
type ListExtensionsV1beta1IngressForAllNamespacesUnauthorized struct {
}

// NewListExtensionsV1beta1IngressForAllNamespacesUnauthorized creates ListExtensionsV1beta1IngressForAllNamespacesUnauthorized with default headers values
func NewListExtensionsV1beta1IngressForAllNamespacesUnauthorized() *ListExtensionsV1beta1IngressForAllNamespacesUnauthorized {

	return &ListExtensionsV1beta1IngressForAllNamespacesUnauthorized{}
}

// WriteResponse to the client
func (o *ListExtensionsV1beta1IngressForAllNamespacesUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}