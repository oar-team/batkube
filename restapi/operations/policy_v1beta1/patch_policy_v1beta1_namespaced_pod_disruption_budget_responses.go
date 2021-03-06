// Code generated by go-swagger; DO NOT EDIT.

package policy_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// PatchPolicyV1beta1NamespacedPodDisruptionBudgetOKCode is the HTTP code returned for type PatchPolicyV1beta1NamespacedPodDisruptionBudgetOK
const PatchPolicyV1beta1NamespacedPodDisruptionBudgetOKCode int = 200

/*PatchPolicyV1beta1NamespacedPodDisruptionBudgetOK OK

swagger:response patchPolicyV1beta1NamespacedPodDisruptionBudgetOK
*/
type PatchPolicyV1beta1NamespacedPodDisruptionBudgetOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIPolicyV1beta1PodDisruptionBudget `json:"body,omitempty"`
}

// NewPatchPolicyV1beta1NamespacedPodDisruptionBudgetOK creates PatchPolicyV1beta1NamespacedPodDisruptionBudgetOK with default headers values
func NewPatchPolicyV1beta1NamespacedPodDisruptionBudgetOK() *PatchPolicyV1beta1NamespacedPodDisruptionBudgetOK {

	return &PatchPolicyV1beta1NamespacedPodDisruptionBudgetOK{}
}

// WithPayload adds the payload to the patch policy v1beta1 namespaced pod disruption budget o k response
func (o *PatchPolicyV1beta1NamespacedPodDisruptionBudgetOK) WithPayload(payload *models.IoK8sAPIPolicyV1beta1PodDisruptionBudget) *PatchPolicyV1beta1NamespacedPodDisruptionBudgetOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch policy v1beta1 namespaced pod disruption budget o k response
func (o *PatchPolicyV1beta1NamespacedPodDisruptionBudgetOK) SetPayload(payload *models.IoK8sAPIPolicyV1beta1PodDisruptionBudget) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchPolicyV1beta1NamespacedPodDisruptionBudgetOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchPolicyV1beta1NamespacedPodDisruptionBudgetUnauthorizedCode is the HTTP code returned for type PatchPolicyV1beta1NamespacedPodDisruptionBudgetUnauthorized
const PatchPolicyV1beta1NamespacedPodDisruptionBudgetUnauthorizedCode int = 401

/*PatchPolicyV1beta1NamespacedPodDisruptionBudgetUnauthorized Unauthorized

swagger:response patchPolicyV1beta1NamespacedPodDisruptionBudgetUnauthorized
*/
type PatchPolicyV1beta1NamespacedPodDisruptionBudgetUnauthorized struct {
}

// NewPatchPolicyV1beta1NamespacedPodDisruptionBudgetUnauthorized creates PatchPolicyV1beta1NamespacedPodDisruptionBudgetUnauthorized with default headers values
func NewPatchPolicyV1beta1NamespacedPodDisruptionBudgetUnauthorized() *PatchPolicyV1beta1NamespacedPodDisruptionBudgetUnauthorized {

	return &PatchPolicyV1beta1NamespacedPodDisruptionBudgetUnauthorized{}
}

// WriteResponse to the client
func (o *PatchPolicyV1beta1NamespacedPodDisruptionBudgetUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
