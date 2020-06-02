// Code generated by go-swagger; DO NOT EDIT.

package storage_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/models"
)

// PatchStorageV1VolumeAttachmentStatusOKCode is the HTTP code returned for type PatchStorageV1VolumeAttachmentStatusOK
const PatchStorageV1VolumeAttachmentStatusOKCode int = 200

/*PatchStorageV1VolumeAttachmentStatusOK OK

swagger:response patchStorageV1VolumeAttachmentStatusOK
*/
type PatchStorageV1VolumeAttachmentStatusOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIStorageV1VolumeAttachment `json:"body,omitempty"`
}

// NewPatchStorageV1VolumeAttachmentStatusOK creates PatchStorageV1VolumeAttachmentStatusOK with default headers values
func NewPatchStorageV1VolumeAttachmentStatusOK() *PatchStorageV1VolumeAttachmentStatusOK {

	return &PatchStorageV1VolumeAttachmentStatusOK{}
}

// WithPayload adds the payload to the patch storage v1 volume attachment status o k response
func (o *PatchStorageV1VolumeAttachmentStatusOK) WithPayload(payload *models.IoK8sAPIStorageV1VolumeAttachment) *PatchStorageV1VolumeAttachmentStatusOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch storage v1 volume attachment status o k response
func (o *PatchStorageV1VolumeAttachmentStatusOK) SetPayload(payload *models.IoK8sAPIStorageV1VolumeAttachment) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchStorageV1VolumeAttachmentStatusOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchStorageV1VolumeAttachmentStatusUnauthorizedCode is the HTTP code returned for type PatchStorageV1VolumeAttachmentStatusUnauthorized
const PatchStorageV1VolumeAttachmentStatusUnauthorizedCode int = 401

/*PatchStorageV1VolumeAttachmentStatusUnauthorized Unauthorized

swagger:response patchStorageV1VolumeAttachmentStatusUnauthorized
*/
type PatchStorageV1VolumeAttachmentStatusUnauthorized struct {
}

// NewPatchStorageV1VolumeAttachmentStatusUnauthorized creates PatchStorageV1VolumeAttachmentStatusUnauthorized with default headers values
func NewPatchStorageV1VolumeAttachmentStatusUnauthorized() *PatchStorageV1VolumeAttachmentStatusUnauthorized {

	return &PatchStorageV1VolumeAttachmentStatusUnauthorized{}
}

// WriteResponse to the client
func (o *PatchStorageV1VolumeAttachmentStatusUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
