// Code generated by go-swagger; DO NOT EDIT.

package storage_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/models"
)

// ListStorageV1VolumeAttachmentOKCode is the HTTP code returned for type ListStorageV1VolumeAttachmentOK
const ListStorageV1VolumeAttachmentOKCode int = 200

/*ListStorageV1VolumeAttachmentOK OK

swagger:response listStorageV1VolumeAttachmentOK
*/
type ListStorageV1VolumeAttachmentOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIStorageV1VolumeAttachmentList `json:"body,omitempty"`
}

// NewListStorageV1VolumeAttachmentOK creates ListStorageV1VolumeAttachmentOK with default headers values
func NewListStorageV1VolumeAttachmentOK() *ListStorageV1VolumeAttachmentOK {

	return &ListStorageV1VolumeAttachmentOK{}
}

// WithPayload adds the payload to the list storage v1 volume attachment o k response
func (o *ListStorageV1VolumeAttachmentOK) WithPayload(payload *models.IoK8sAPIStorageV1VolumeAttachmentList) *ListStorageV1VolumeAttachmentOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list storage v1 volume attachment o k response
func (o *ListStorageV1VolumeAttachmentOK) SetPayload(payload *models.IoK8sAPIStorageV1VolumeAttachmentList) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListStorageV1VolumeAttachmentOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListStorageV1VolumeAttachmentUnauthorizedCode is the HTTP code returned for type ListStorageV1VolumeAttachmentUnauthorized
const ListStorageV1VolumeAttachmentUnauthorizedCode int = 401

/*ListStorageV1VolumeAttachmentUnauthorized Unauthorized

swagger:response listStorageV1VolumeAttachmentUnauthorized
*/
type ListStorageV1VolumeAttachmentUnauthorized struct {
}

// NewListStorageV1VolumeAttachmentUnauthorized creates ListStorageV1VolumeAttachmentUnauthorized with default headers values
func NewListStorageV1VolumeAttachmentUnauthorized() *ListStorageV1VolumeAttachmentUnauthorized {

	return &ListStorageV1VolumeAttachmentUnauthorized{}
}

// WriteResponse to the client
func (o *ListStorageV1VolumeAttachmentUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
