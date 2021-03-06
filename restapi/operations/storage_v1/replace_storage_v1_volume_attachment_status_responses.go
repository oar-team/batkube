// Code generated by go-swagger; DO NOT EDIT.

package storage_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// ReplaceStorageV1VolumeAttachmentStatusOKCode is the HTTP code returned for type ReplaceStorageV1VolumeAttachmentStatusOK
const ReplaceStorageV1VolumeAttachmentStatusOKCode int = 200

/*ReplaceStorageV1VolumeAttachmentStatusOK OK

swagger:response replaceStorageV1VolumeAttachmentStatusOK
*/
type ReplaceStorageV1VolumeAttachmentStatusOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIStorageV1VolumeAttachment `json:"body,omitempty"`
}

// NewReplaceStorageV1VolumeAttachmentStatusOK creates ReplaceStorageV1VolumeAttachmentStatusOK with default headers values
func NewReplaceStorageV1VolumeAttachmentStatusOK() *ReplaceStorageV1VolumeAttachmentStatusOK {

	return &ReplaceStorageV1VolumeAttachmentStatusOK{}
}

// WithPayload adds the payload to the replace storage v1 volume attachment status o k response
func (o *ReplaceStorageV1VolumeAttachmentStatusOK) WithPayload(payload *models.IoK8sAPIStorageV1VolumeAttachment) *ReplaceStorageV1VolumeAttachmentStatusOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace storage v1 volume attachment status o k response
func (o *ReplaceStorageV1VolumeAttachmentStatusOK) SetPayload(payload *models.IoK8sAPIStorageV1VolumeAttachment) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceStorageV1VolumeAttachmentStatusOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceStorageV1VolumeAttachmentStatusCreatedCode is the HTTP code returned for type ReplaceStorageV1VolumeAttachmentStatusCreated
const ReplaceStorageV1VolumeAttachmentStatusCreatedCode int = 201

/*ReplaceStorageV1VolumeAttachmentStatusCreated Created

swagger:response replaceStorageV1VolumeAttachmentStatusCreated
*/
type ReplaceStorageV1VolumeAttachmentStatusCreated struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIStorageV1VolumeAttachment `json:"body,omitempty"`
}

// NewReplaceStorageV1VolumeAttachmentStatusCreated creates ReplaceStorageV1VolumeAttachmentStatusCreated with default headers values
func NewReplaceStorageV1VolumeAttachmentStatusCreated() *ReplaceStorageV1VolumeAttachmentStatusCreated {

	return &ReplaceStorageV1VolumeAttachmentStatusCreated{}
}

// WithPayload adds the payload to the replace storage v1 volume attachment status created response
func (o *ReplaceStorageV1VolumeAttachmentStatusCreated) WithPayload(payload *models.IoK8sAPIStorageV1VolumeAttachment) *ReplaceStorageV1VolumeAttachmentStatusCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace storage v1 volume attachment status created response
func (o *ReplaceStorageV1VolumeAttachmentStatusCreated) SetPayload(payload *models.IoK8sAPIStorageV1VolumeAttachment) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceStorageV1VolumeAttachmentStatusCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceStorageV1VolumeAttachmentStatusUnauthorizedCode is the HTTP code returned for type ReplaceStorageV1VolumeAttachmentStatusUnauthorized
const ReplaceStorageV1VolumeAttachmentStatusUnauthorizedCode int = 401

/*ReplaceStorageV1VolumeAttachmentStatusUnauthorized Unauthorized

swagger:response replaceStorageV1VolumeAttachmentStatusUnauthorized
*/
type ReplaceStorageV1VolumeAttachmentStatusUnauthorized struct {
}

// NewReplaceStorageV1VolumeAttachmentStatusUnauthorized creates ReplaceStorageV1VolumeAttachmentStatusUnauthorized with default headers values
func NewReplaceStorageV1VolumeAttachmentStatusUnauthorized() *ReplaceStorageV1VolumeAttachmentStatusUnauthorized {

	return &ReplaceStorageV1VolumeAttachmentStatusUnauthorized{}
}

// WriteResponse to the client
func (o *ReplaceStorageV1VolumeAttachmentStatusUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
