// Code generated by go-swagger; DO NOT EDIT.

package autoscaling_v2beta2

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// PatchAutoscalingV2beta2NamespacedHorizontalPodAutoscalerOKCode is the HTTP code returned for type PatchAutoscalingV2beta2NamespacedHorizontalPodAutoscalerOK
const PatchAutoscalingV2beta2NamespacedHorizontalPodAutoscalerOKCode int = 200

/*PatchAutoscalingV2beta2NamespacedHorizontalPodAutoscalerOK OK

swagger:response patchAutoscalingV2beta2NamespacedHorizontalPodAutoscalerOK
*/
type PatchAutoscalingV2beta2NamespacedHorizontalPodAutoscalerOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIAutoscalingV2beta2HorizontalPodAutoscaler `json:"body,omitempty"`
}

// NewPatchAutoscalingV2beta2NamespacedHorizontalPodAutoscalerOK creates PatchAutoscalingV2beta2NamespacedHorizontalPodAutoscalerOK with default headers values
func NewPatchAutoscalingV2beta2NamespacedHorizontalPodAutoscalerOK() *PatchAutoscalingV2beta2NamespacedHorizontalPodAutoscalerOK {

	return &PatchAutoscalingV2beta2NamespacedHorizontalPodAutoscalerOK{}
}

// WithPayload adds the payload to the patch autoscaling v2beta2 namespaced horizontal pod autoscaler o k response
func (o *PatchAutoscalingV2beta2NamespacedHorizontalPodAutoscalerOK) WithPayload(payload *models.IoK8sAPIAutoscalingV2beta2HorizontalPodAutoscaler) *PatchAutoscalingV2beta2NamespacedHorizontalPodAutoscalerOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch autoscaling v2beta2 namespaced horizontal pod autoscaler o k response
func (o *PatchAutoscalingV2beta2NamespacedHorizontalPodAutoscalerOK) SetPayload(payload *models.IoK8sAPIAutoscalingV2beta2HorizontalPodAutoscaler) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchAutoscalingV2beta2NamespacedHorizontalPodAutoscalerOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchAutoscalingV2beta2NamespacedHorizontalPodAutoscalerUnauthorizedCode is the HTTP code returned for type PatchAutoscalingV2beta2NamespacedHorizontalPodAutoscalerUnauthorized
const PatchAutoscalingV2beta2NamespacedHorizontalPodAutoscalerUnauthorizedCode int = 401

/*PatchAutoscalingV2beta2NamespacedHorizontalPodAutoscalerUnauthorized Unauthorized

swagger:response patchAutoscalingV2beta2NamespacedHorizontalPodAutoscalerUnauthorized
*/
type PatchAutoscalingV2beta2NamespacedHorizontalPodAutoscalerUnauthorized struct {
}

// NewPatchAutoscalingV2beta2NamespacedHorizontalPodAutoscalerUnauthorized creates PatchAutoscalingV2beta2NamespacedHorizontalPodAutoscalerUnauthorized with default headers values
func NewPatchAutoscalingV2beta2NamespacedHorizontalPodAutoscalerUnauthorized() *PatchAutoscalingV2beta2NamespacedHorizontalPodAutoscalerUnauthorized {

	return &PatchAutoscalingV2beta2NamespacedHorizontalPodAutoscalerUnauthorized{}
}

// WriteResponse to the client
func (o *PatchAutoscalingV2beta2NamespacedHorizontalPodAutoscalerUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
