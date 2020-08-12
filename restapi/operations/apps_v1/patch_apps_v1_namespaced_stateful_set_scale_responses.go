// Code generated by go-swagger; DO NOT EDIT.

package apps_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// PatchAppsV1NamespacedStatefulSetScaleOKCode is the HTTP code returned for type PatchAppsV1NamespacedStatefulSetScaleOK
const PatchAppsV1NamespacedStatefulSetScaleOKCode int = 200

/*PatchAppsV1NamespacedStatefulSetScaleOK OK

swagger:response patchAppsV1NamespacedStatefulSetScaleOK
*/
type PatchAppsV1NamespacedStatefulSetScaleOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIAutoscalingV1Scale `json:"body,omitempty"`
}

// NewPatchAppsV1NamespacedStatefulSetScaleOK creates PatchAppsV1NamespacedStatefulSetScaleOK with default headers values
func NewPatchAppsV1NamespacedStatefulSetScaleOK() *PatchAppsV1NamespacedStatefulSetScaleOK {

	return &PatchAppsV1NamespacedStatefulSetScaleOK{}
}

// WithPayload adds the payload to the patch apps v1 namespaced stateful set scale o k response
func (o *PatchAppsV1NamespacedStatefulSetScaleOK) WithPayload(payload *models.IoK8sAPIAutoscalingV1Scale) *PatchAppsV1NamespacedStatefulSetScaleOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch apps v1 namespaced stateful set scale o k response
func (o *PatchAppsV1NamespacedStatefulSetScaleOK) SetPayload(payload *models.IoK8sAPIAutoscalingV1Scale) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchAppsV1NamespacedStatefulSetScaleOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchAppsV1NamespacedStatefulSetScaleUnauthorizedCode is the HTTP code returned for type PatchAppsV1NamespacedStatefulSetScaleUnauthorized
const PatchAppsV1NamespacedStatefulSetScaleUnauthorizedCode int = 401

/*PatchAppsV1NamespacedStatefulSetScaleUnauthorized Unauthorized

swagger:response patchAppsV1NamespacedStatefulSetScaleUnauthorized
*/
type PatchAppsV1NamespacedStatefulSetScaleUnauthorized struct {
}

// NewPatchAppsV1NamespacedStatefulSetScaleUnauthorized creates PatchAppsV1NamespacedStatefulSetScaleUnauthorized with default headers values
func NewPatchAppsV1NamespacedStatefulSetScaleUnauthorized() *PatchAppsV1NamespacedStatefulSetScaleUnauthorized {

	return &PatchAppsV1NamespacedStatefulSetScaleUnauthorized{}
}

// WriteResponse to the client
func (o *PatchAppsV1NamespacedStatefulSetScaleUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
