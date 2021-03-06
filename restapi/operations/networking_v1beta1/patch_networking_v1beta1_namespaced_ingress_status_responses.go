// Code generated by go-swagger; DO NOT EDIT.

package networking_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// PatchNetworkingV1beta1NamespacedIngressStatusOKCode is the HTTP code returned for type PatchNetworkingV1beta1NamespacedIngressStatusOK
const PatchNetworkingV1beta1NamespacedIngressStatusOKCode int = 200

/*PatchNetworkingV1beta1NamespacedIngressStatusOK OK

swagger:response patchNetworkingV1beta1NamespacedIngressStatusOK
*/
type PatchNetworkingV1beta1NamespacedIngressStatusOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPINetworkingV1beta1Ingress `json:"body,omitempty"`
}

// NewPatchNetworkingV1beta1NamespacedIngressStatusOK creates PatchNetworkingV1beta1NamespacedIngressStatusOK with default headers values
func NewPatchNetworkingV1beta1NamespacedIngressStatusOK() *PatchNetworkingV1beta1NamespacedIngressStatusOK {

	return &PatchNetworkingV1beta1NamespacedIngressStatusOK{}
}

// WithPayload adds the payload to the patch networking v1beta1 namespaced ingress status o k response
func (o *PatchNetworkingV1beta1NamespacedIngressStatusOK) WithPayload(payload *models.IoK8sAPINetworkingV1beta1Ingress) *PatchNetworkingV1beta1NamespacedIngressStatusOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch networking v1beta1 namespaced ingress status o k response
func (o *PatchNetworkingV1beta1NamespacedIngressStatusOK) SetPayload(payload *models.IoK8sAPINetworkingV1beta1Ingress) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchNetworkingV1beta1NamespacedIngressStatusOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchNetworkingV1beta1NamespacedIngressStatusUnauthorizedCode is the HTTP code returned for type PatchNetworkingV1beta1NamespacedIngressStatusUnauthorized
const PatchNetworkingV1beta1NamespacedIngressStatusUnauthorizedCode int = 401

/*PatchNetworkingV1beta1NamespacedIngressStatusUnauthorized Unauthorized

swagger:response patchNetworkingV1beta1NamespacedIngressStatusUnauthorized
*/
type PatchNetworkingV1beta1NamespacedIngressStatusUnauthorized struct {
}

// NewPatchNetworkingV1beta1NamespacedIngressStatusUnauthorized creates PatchNetworkingV1beta1NamespacedIngressStatusUnauthorized with default headers values
func NewPatchNetworkingV1beta1NamespacedIngressStatusUnauthorized() *PatchNetworkingV1beta1NamespacedIngressStatusUnauthorized {

	return &PatchNetworkingV1beta1NamespacedIngressStatusUnauthorized{}
}

// WriteResponse to the client
func (o *PatchNetworkingV1beta1NamespacedIngressStatusUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
