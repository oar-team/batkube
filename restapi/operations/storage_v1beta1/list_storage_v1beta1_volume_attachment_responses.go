// Code generated by go-swagger; DO NOT EDIT.

package storage_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// ListStorageV1beta1VolumeAttachmentOKCode is the HTTP code returned for type ListStorageV1beta1VolumeAttachmentOK
const ListStorageV1beta1VolumeAttachmentOKCode int = 200

/*ListStorageV1beta1VolumeAttachmentOK OK

swagger:response listStorageV1beta1VolumeAttachmentOK
*/
type ListStorageV1beta1VolumeAttachmentOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIStorageV1beta1VolumeAttachmentList `json:"body,omitempty"`
}

// NewListStorageV1beta1VolumeAttachmentOK creates ListStorageV1beta1VolumeAttachmentOK with default headers values
func NewListStorageV1beta1VolumeAttachmentOK() *ListStorageV1beta1VolumeAttachmentOK {

	return &ListStorageV1beta1VolumeAttachmentOK{}
}

// WithPayload adds the payload to the list storage v1beta1 volume attachment o k response
func (o *ListStorageV1beta1VolumeAttachmentOK) WithPayload(payload *models.IoK8sAPIStorageV1beta1VolumeAttachmentList) *ListStorageV1beta1VolumeAttachmentOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list storage v1beta1 volume attachment o k response
func (o *ListStorageV1beta1VolumeAttachmentOK) SetPayload(payload *models.IoK8sAPIStorageV1beta1VolumeAttachmentList) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListStorageV1beta1VolumeAttachmentOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListStorageV1beta1VolumeAttachmentUnauthorizedCode is the HTTP code returned for type ListStorageV1beta1VolumeAttachmentUnauthorized
const ListStorageV1beta1VolumeAttachmentUnauthorizedCode int = 401

/*ListStorageV1beta1VolumeAttachmentUnauthorized Unauthorized

swagger:response listStorageV1beta1VolumeAttachmentUnauthorized
*/
type ListStorageV1beta1VolumeAttachmentUnauthorized struct {
}

// NewListStorageV1beta1VolumeAttachmentUnauthorized creates ListStorageV1beta1VolumeAttachmentUnauthorized with default headers values
func NewListStorageV1beta1VolumeAttachmentUnauthorized() *ListStorageV1beta1VolumeAttachmentUnauthorized {

	return &ListStorageV1beta1VolumeAttachmentUnauthorized{}
}

// WriteResponse to the client
func (o *ListStorageV1beta1VolumeAttachmentUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
