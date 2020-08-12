// Code generated by go-swagger; DO NOT EDIT.

package extensions_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// PatchExtensionsV1beta1NamespacedIngressOKCode is the HTTP code returned for type PatchExtensionsV1beta1NamespacedIngressOK
const PatchExtensionsV1beta1NamespacedIngressOKCode int = 200

/*PatchExtensionsV1beta1NamespacedIngressOK OK

swagger:response patchExtensionsV1beta1NamespacedIngressOK
*/
type PatchExtensionsV1beta1NamespacedIngressOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIExtensionsV1beta1Ingress `json:"body,omitempty"`
}

// NewPatchExtensionsV1beta1NamespacedIngressOK creates PatchExtensionsV1beta1NamespacedIngressOK with default headers values
func NewPatchExtensionsV1beta1NamespacedIngressOK() *PatchExtensionsV1beta1NamespacedIngressOK {

	return &PatchExtensionsV1beta1NamespacedIngressOK{}
}

// WithPayload adds the payload to the patch extensions v1beta1 namespaced ingress o k response
func (o *PatchExtensionsV1beta1NamespacedIngressOK) WithPayload(payload *models.IoK8sAPIExtensionsV1beta1Ingress) *PatchExtensionsV1beta1NamespacedIngressOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch extensions v1beta1 namespaced ingress o k response
func (o *PatchExtensionsV1beta1NamespacedIngressOK) SetPayload(payload *models.IoK8sAPIExtensionsV1beta1Ingress) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchExtensionsV1beta1NamespacedIngressOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchExtensionsV1beta1NamespacedIngressUnauthorizedCode is the HTTP code returned for type PatchExtensionsV1beta1NamespacedIngressUnauthorized
const PatchExtensionsV1beta1NamespacedIngressUnauthorizedCode int = 401

/*PatchExtensionsV1beta1NamespacedIngressUnauthorized Unauthorized

swagger:response patchExtensionsV1beta1NamespacedIngressUnauthorized
*/
type PatchExtensionsV1beta1NamespacedIngressUnauthorized struct {
}

// NewPatchExtensionsV1beta1NamespacedIngressUnauthorized creates PatchExtensionsV1beta1NamespacedIngressUnauthorized with default headers values
func NewPatchExtensionsV1beta1NamespacedIngressUnauthorized() *PatchExtensionsV1beta1NamespacedIngressUnauthorized {

	return &PatchExtensionsV1beta1NamespacedIngressUnauthorized{}
}

// WriteResponse to the client
func (o *PatchExtensionsV1beta1NamespacedIngressUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
