// Code generated by go-swagger; DO NOT EDIT.

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// PatchCoreV1PersistentVolumeStatusOKCode is the HTTP code returned for type PatchCoreV1PersistentVolumeStatusOK
const PatchCoreV1PersistentVolumeStatusOKCode int = 200

/*PatchCoreV1PersistentVolumeStatusOK OK

swagger:response patchCoreV1PersistentVolumeStatusOK
*/
type PatchCoreV1PersistentVolumeStatusOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPICoreV1PersistentVolume `json:"body,omitempty"`
}

// NewPatchCoreV1PersistentVolumeStatusOK creates PatchCoreV1PersistentVolumeStatusOK with default headers values
func NewPatchCoreV1PersistentVolumeStatusOK() *PatchCoreV1PersistentVolumeStatusOK {

	return &PatchCoreV1PersistentVolumeStatusOK{}
}

// WithPayload adds the payload to the patch core v1 persistent volume status o k response
func (o *PatchCoreV1PersistentVolumeStatusOK) WithPayload(payload *models.IoK8sAPICoreV1PersistentVolume) *PatchCoreV1PersistentVolumeStatusOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch core v1 persistent volume status o k response
func (o *PatchCoreV1PersistentVolumeStatusOK) SetPayload(payload *models.IoK8sAPICoreV1PersistentVolume) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchCoreV1PersistentVolumeStatusOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchCoreV1PersistentVolumeStatusUnauthorizedCode is the HTTP code returned for type PatchCoreV1PersistentVolumeStatusUnauthorized
const PatchCoreV1PersistentVolumeStatusUnauthorizedCode int = 401

/*PatchCoreV1PersistentVolumeStatusUnauthorized Unauthorized

swagger:response patchCoreV1PersistentVolumeStatusUnauthorized
*/
type PatchCoreV1PersistentVolumeStatusUnauthorized struct {
}

// NewPatchCoreV1PersistentVolumeStatusUnauthorized creates PatchCoreV1PersistentVolumeStatusUnauthorized with default headers values
func NewPatchCoreV1PersistentVolumeStatusUnauthorized() *PatchCoreV1PersistentVolumeStatusUnauthorized {

	return &PatchCoreV1PersistentVolumeStatusUnauthorized{}
}

// WriteResponse to the client
func (o *PatchCoreV1PersistentVolumeStatusUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
