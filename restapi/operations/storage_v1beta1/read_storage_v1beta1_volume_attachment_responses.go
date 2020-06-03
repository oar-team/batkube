// Code generated by go-swagger; DO NOT EDIT.

package storage_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/models"
)

// ReadStorageV1beta1VolumeAttachmentOKCode is the HTTP code returned for type ReadStorageV1beta1VolumeAttachmentOK
const ReadStorageV1beta1VolumeAttachmentOKCode int = 200

/*ReadStorageV1beta1VolumeAttachmentOK OK

swagger:response readStorageV1beta1VolumeAttachmentOK
*/
type ReadStorageV1beta1VolumeAttachmentOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIStorageV1beta1VolumeAttachment `json:"body,omitempty"`
}

// NewReadStorageV1beta1VolumeAttachmentOK creates ReadStorageV1beta1VolumeAttachmentOK with default headers values
func NewReadStorageV1beta1VolumeAttachmentOK() *ReadStorageV1beta1VolumeAttachmentOK {

	return &ReadStorageV1beta1VolumeAttachmentOK{}
}

// WithPayload adds the payload to the read storage v1beta1 volume attachment o k response
func (o *ReadStorageV1beta1VolumeAttachmentOK) WithPayload(payload *models.IoK8sAPIStorageV1beta1VolumeAttachment) *ReadStorageV1beta1VolumeAttachmentOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the read storage v1beta1 volume attachment o k response
func (o *ReadStorageV1beta1VolumeAttachmentOK) SetPayload(payload *models.IoK8sAPIStorageV1beta1VolumeAttachment) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReadStorageV1beta1VolumeAttachmentOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReadStorageV1beta1VolumeAttachmentUnauthorizedCode is the HTTP code returned for type ReadStorageV1beta1VolumeAttachmentUnauthorized
const ReadStorageV1beta1VolumeAttachmentUnauthorizedCode int = 401

/*ReadStorageV1beta1VolumeAttachmentUnauthorized Unauthorized

swagger:response readStorageV1beta1VolumeAttachmentUnauthorized
*/
type ReadStorageV1beta1VolumeAttachmentUnauthorized struct {
}

// NewReadStorageV1beta1VolumeAttachmentUnauthorized creates ReadStorageV1beta1VolumeAttachmentUnauthorized with default headers values
func NewReadStorageV1beta1VolumeAttachmentUnauthorized() *ReadStorageV1beta1VolumeAttachmentUnauthorized {

	return &ReadStorageV1beta1VolumeAttachmentUnauthorized{}
}

// WriteResponse to the client
func (o *ReadStorageV1beta1VolumeAttachmentUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}