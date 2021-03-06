// Code generated by go-swagger; DO NOT EDIT.

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// PatchCoreV1NamespacedReplicationControllerScaleOKCode is the HTTP code returned for type PatchCoreV1NamespacedReplicationControllerScaleOK
const PatchCoreV1NamespacedReplicationControllerScaleOKCode int = 200

/*PatchCoreV1NamespacedReplicationControllerScaleOK OK

swagger:response patchCoreV1NamespacedReplicationControllerScaleOK
*/
type PatchCoreV1NamespacedReplicationControllerScaleOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIAutoscalingV1Scale `json:"body,omitempty"`
}

// NewPatchCoreV1NamespacedReplicationControllerScaleOK creates PatchCoreV1NamespacedReplicationControllerScaleOK with default headers values
func NewPatchCoreV1NamespacedReplicationControllerScaleOK() *PatchCoreV1NamespacedReplicationControllerScaleOK {

	return &PatchCoreV1NamespacedReplicationControllerScaleOK{}
}

// WithPayload adds the payload to the patch core v1 namespaced replication controller scale o k response
func (o *PatchCoreV1NamespacedReplicationControllerScaleOK) WithPayload(payload *models.IoK8sAPIAutoscalingV1Scale) *PatchCoreV1NamespacedReplicationControllerScaleOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch core v1 namespaced replication controller scale o k response
func (o *PatchCoreV1NamespacedReplicationControllerScaleOK) SetPayload(payload *models.IoK8sAPIAutoscalingV1Scale) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchCoreV1NamespacedReplicationControllerScaleOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchCoreV1NamespacedReplicationControllerScaleUnauthorizedCode is the HTTP code returned for type PatchCoreV1NamespacedReplicationControllerScaleUnauthorized
const PatchCoreV1NamespacedReplicationControllerScaleUnauthorizedCode int = 401

/*PatchCoreV1NamespacedReplicationControllerScaleUnauthorized Unauthorized

swagger:response patchCoreV1NamespacedReplicationControllerScaleUnauthorized
*/
type PatchCoreV1NamespacedReplicationControllerScaleUnauthorized struct {
}

// NewPatchCoreV1NamespacedReplicationControllerScaleUnauthorized creates PatchCoreV1NamespacedReplicationControllerScaleUnauthorized with default headers values
func NewPatchCoreV1NamespacedReplicationControllerScaleUnauthorized() *PatchCoreV1NamespacedReplicationControllerScaleUnauthorized {

	return &PatchCoreV1NamespacedReplicationControllerScaleUnauthorized{}
}

// WriteResponse to the client
func (o *PatchCoreV1NamespacedReplicationControllerScaleUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
