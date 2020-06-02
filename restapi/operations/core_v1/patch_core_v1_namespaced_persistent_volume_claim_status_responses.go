// Code generated by go-swagger; DO NOT EDIT.

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/models"
)

// PatchCoreV1NamespacedPersistentVolumeClaimStatusOKCode is the HTTP code returned for type PatchCoreV1NamespacedPersistentVolumeClaimStatusOK
const PatchCoreV1NamespacedPersistentVolumeClaimStatusOKCode int = 200

/*PatchCoreV1NamespacedPersistentVolumeClaimStatusOK OK

swagger:response patchCoreV1NamespacedPersistentVolumeClaimStatusOK
*/
type PatchCoreV1NamespacedPersistentVolumeClaimStatusOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPICoreV1PersistentVolumeClaim `json:"body,omitempty"`
}

// NewPatchCoreV1NamespacedPersistentVolumeClaimStatusOK creates PatchCoreV1NamespacedPersistentVolumeClaimStatusOK with default headers values
func NewPatchCoreV1NamespacedPersistentVolumeClaimStatusOK() *PatchCoreV1NamespacedPersistentVolumeClaimStatusOK {

	return &PatchCoreV1NamespacedPersistentVolumeClaimStatusOK{}
}

// WithPayload adds the payload to the patch core v1 namespaced persistent volume claim status o k response
func (o *PatchCoreV1NamespacedPersistentVolumeClaimStatusOK) WithPayload(payload *models.IoK8sAPICoreV1PersistentVolumeClaim) *PatchCoreV1NamespacedPersistentVolumeClaimStatusOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch core v1 namespaced persistent volume claim status o k response
func (o *PatchCoreV1NamespacedPersistentVolumeClaimStatusOK) SetPayload(payload *models.IoK8sAPICoreV1PersistentVolumeClaim) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchCoreV1NamespacedPersistentVolumeClaimStatusOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchCoreV1NamespacedPersistentVolumeClaimStatusUnauthorizedCode is the HTTP code returned for type PatchCoreV1NamespacedPersistentVolumeClaimStatusUnauthorized
const PatchCoreV1NamespacedPersistentVolumeClaimStatusUnauthorizedCode int = 401

/*PatchCoreV1NamespacedPersistentVolumeClaimStatusUnauthorized Unauthorized

swagger:response patchCoreV1NamespacedPersistentVolumeClaimStatusUnauthorized
*/
type PatchCoreV1NamespacedPersistentVolumeClaimStatusUnauthorized struct {
}

// NewPatchCoreV1NamespacedPersistentVolumeClaimStatusUnauthorized creates PatchCoreV1NamespacedPersistentVolumeClaimStatusUnauthorized with default headers values
func NewPatchCoreV1NamespacedPersistentVolumeClaimStatusUnauthorized() *PatchCoreV1NamespacedPersistentVolumeClaimStatusUnauthorized {

	return &PatchCoreV1NamespacedPersistentVolumeClaimStatusUnauthorized{}
}

// WriteResponse to the client
func (o *PatchCoreV1NamespacedPersistentVolumeClaimStatusUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
