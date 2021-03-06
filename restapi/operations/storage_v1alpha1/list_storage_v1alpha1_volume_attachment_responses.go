// Code generated by go-swagger; DO NOT EDIT.

package storage_v1alpha1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// ListStorageV1alpha1VolumeAttachmentOKCode is the HTTP code returned for type ListStorageV1alpha1VolumeAttachmentOK
const ListStorageV1alpha1VolumeAttachmentOKCode int = 200

/*ListStorageV1alpha1VolumeAttachmentOK OK

swagger:response listStorageV1alpha1VolumeAttachmentOK
*/
type ListStorageV1alpha1VolumeAttachmentOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIStorageV1alpha1VolumeAttachmentList `json:"body,omitempty"`
}

// NewListStorageV1alpha1VolumeAttachmentOK creates ListStorageV1alpha1VolumeAttachmentOK with default headers values
func NewListStorageV1alpha1VolumeAttachmentOK() *ListStorageV1alpha1VolumeAttachmentOK {

	return &ListStorageV1alpha1VolumeAttachmentOK{}
}

// WithPayload adds the payload to the list storage v1alpha1 volume attachment o k response
func (o *ListStorageV1alpha1VolumeAttachmentOK) WithPayload(payload *models.IoK8sAPIStorageV1alpha1VolumeAttachmentList) *ListStorageV1alpha1VolumeAttachmentOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list storage v1alpha1 volume attachment o k response
func (o *ListStorageV1alpha1VolumeAttachmentOK) SetPayload(payload *models.IoK8sAPIStorageV1alpha1VolumeAttachmentList) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListStorageV1alpha1VolumeAttachmentOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListStorageV1alpha1VolumeAttachmentUnauthorizedCode is the HTTP code returned for type ListStorageV1alpha1VolumeAttachmentUnauthorized
const ListStorageV1alpha1VolumeAttachmentUnauthorizedCode int = 401

/*ListStorageV1alpha1VolumeAttachmentUnauthorized Unauthorized

swagger:response listStorageV1alpha1VolumeAttachmentUnauthorized
*/
type ListStorageV1alpha1VolumeAttachmentUnauthorized struct {
}

// NewListStorageV1alpha1VolumeAttachmentUnauthorized creates ListStorageV1alpha1VolumeAttachmentUnauthorized with default headers values
func NewListStorageV1alpha1VolumeAttachmentUnauthorized() *ListStorageV1alpha1VolumeAttachmentUnauthorized {

	return &ListStorageV1alpha1VolumeAttachmentUnauthorized{}
}

// WriteResponse to the client
func (o *ListStorageV1alpha1VolumeAttachmentUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
