// Code generated by go-swagger; DO NOT EDIT.

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// PatchCoreV1NamespacedPodOKCode is the HTTP code returned for type PatchCoreV1NamespacedPodOK
const PatchCoreV1NamespacedPodOKCode int = 200

/*PatchCoreV1NamespacedPodOK OK

swagger:response patchCoreV1NamespacedPodOK
*/
type PatchCoreV1NamespacedPodOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPICoreV1Pod `json:"body,omitempty"`
}

// NewPatchCoreV1NamespacedPodOK creates PatchCoreV1NamespacedPodOK with default headers values
func NewPatchCoreV1NamespacedPodOK() *PatchCoreV1NamespacedPodOK {

	return &PatchCoreV1NamespacedPodOK{}
}

// WithPayload adds the payload to the patch core v1 namespaced pod o k response
func (o *PatchCoreV1NamespacedPodOK) WithPayload(payload *models.IoK8sAPICoreV1Pod) *PatchCoreV1NamespacedPodOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch core v1 namespaced pod o k response
func (o *PatchCoreV1NamespacedPodOK) SetPayload(payload *models.IoK8sAPICoreV1Pod) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchCoreV1NamespacedPodOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchCoreV1NamespacedPodUnauthorizedCode is the HTTP code returned for type PatchCoreV1NamespacedPodUnauthorized
const PatchCoreV1NamespacedPodUnauthorizedCode int = 401

/*PatchCoreV1NamespacedPodUnauthorized Unauthorized

swagger:response patchCoreV1NamespacedPodUnauthorized
*/
type PatchCoreV1NamespacedPodUnauthorized struct {
}

// NewPatchCoreV1NamespacedPodUnauthorized creates PatchCoreV1NamespacedPodUnauthorized with default headers values
func NewPatchCoreV1NamespacedPodUnauthorized() *PatchCoreV1NamespacedPodUnauthorized {

	return &PatchCoreV1NamespacedPodUnauthorized{}
}

// WriteResponse to the client
func (o *PatchCoreV1NamespacedPodUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
