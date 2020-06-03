// Code generated by go-swagger; DO NOT EDIT.

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/models"
)

// PatchCoreV1NamespacedPodTemplateOKCode is the HTTP code returned for type PatchCoreV1NamespacedPodTemplateOK
const PatchCoreV1NamespacedPodTemplateOKCode int = 200

/*PatchCoreV1NamespacedPodTemplateOK OK

swagger:response patchCoreV1NamespacedPodTemplateOK
*/
type PatchCoreV1NamespacedPodTemplateOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPICoreV1PodTemplate `json:"body,omitempty"`
}

// NewPatchCoreV1NamespacedPodTemplateOK creates PatchCoreV1NamespacedPodTemplateOK with default headers values
func NewPatchCoreV1NamespacedPodTemplateOK() *PatchCoreV1NamespacedPodTemplateOK {

	return &PatchCoreV1NamespacedPodTemplateOK{}
}

// WithPayload adds the payload to the patch core v1 namespaced pod template o k response
func (o *PatchCoreV1NamespacedPodTemplateOK) WithPayload(payload *models.IoK8sAPICoreV1PodTemplate) *PatchCoreV1NamespacedPodTemplateOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch core v1 namespaced pod template o k response
func (o *PatchCoreV1NamespacedPodTemplateOK) SetPayload(payload *models.IoK8sAPICoreV1PodTemplate) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchCoreV1NamespacedPodTemplateOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchCoreV1NamespacedPodTemplateUnauthorizedCode is the HTTP code returned for type PatchCoreV1NamespacedPodTemplateUnauthorized
const PatchCoreV1NamespacedPodTemplateUnauthorizedCode int = 401

/*PatchCoreV1NamespacedPodTemplateUnauthorized Unauthorized

swagger:response patchCoreV1NamespacedPodTemplateUnauthorized
*/
type PatchCoreV1NamespacedPodTemplateUnauthorized struct {
}

// NewPatchCoreV1NamespacedPodTemplateUnauthorized creates PatchCoreV1NamespacedPodTemplateUnauthorized with default headers values
func NewPatchCoreV1NamespacedPodTemplateUnauthorized() *PatchCoreV1NamespacedPodTemplateUnauthorized {

	return &PatchCoreV1NamespacedPodTemplateUnauthorized{}
}

// WriteResponse to the client
func (o *PatchCoreV1NamespacedPodTemplateUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}