// Code generated by go-swagger; DO NOT EDIT.

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// PatchCoreV1NamespacedConfigMapOKCode is the HTTP code returned for type PatchCoreV1NamespacedConfigMapOK
const PatchCoreV1NamespacedConfigMapOKCode int = 200

/*PatchCoreV1NamespacedConfigMapOK OK

swagger:response patchCoreV1NamespacedConfigMapOK
*/
type PatchCoreV1NamespacedConfigMapOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPICoreV1ConfigMap `json:"body,omitempty"`
}

// NewPatchCoreV1NamespacedConfigMapOK creates PatchCoreV1NamespacedConfigMapOK with default headers values
func NewPatchCoreV1NamespacedConfigMapOK() *PatchCoreV1NamespacedConfigMapOK {

	return &PatchCoreV1NamespacedConfigMapOK{}
}

// WithPayload adds the payload to the patch core v1 namespaced config map o k response
func (o *PatchCoreV1NamespacedConfigMapOK) WithPayload(payload *models.IoK8sAPICoreV1ConfigMap) *PatchCoreV1NamespacedConfigMapOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch core v1 namespaced config map o k response
func (o *PatchCoreV1NamespacedConfigMapOK) SetPayload(payload *models.IoK8sAPICoreV1ConfigMap) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchCoreV1NamespacedConfigMapOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchCoreV1NamespacedConfigMapUnauthorizedCode is the HTTP code returned for type PatchCoreV1NamespacedConfigMapUnauthorized
const PatchCoreV1NamespacedConfigMapUnauthorizedCode int = 401

/*PatchCoreV1NamespacedConfigMapUnauthorized Unauthorized

swagger:response patchCoreV1NamespacedConfigMapUnauthorized
*/
type PatchCoreV1NamespacedConfigMapUnauthorized struct {
}

// NewPatchCoreV1NamespacedConfigMapUnauthorized creates PatchCoreV1NamespacedConfigMapUnauthorized with default headers values
func NewPatchCoreV1NamespacedConfigMapUnauthorized() *PatchCoreV1NamespacedConfigMapUnauthorized {

	return &PatchCoreV1NamespacedConfigMapUnauthorized{}
}

// WriteResponse to the client
func (o *PatchCoreV1NamespacedConfigMapUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
