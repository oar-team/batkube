// Code generated by go-swagger; DO NOT EDIT.

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// ListCoreV1PersistentVolumeOKCode is the HTTP code returned for type ListCoreV1PersistentVolumeOK
const ListCoreV1PersistentVolumeOKCode int = 200

/*ListCoreV1PersistentVolumeOK OK

swagger:response listCoreV1PersistentVolumeOK
*/
type ListCoreV1PersistentVolumeOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPICoreV1PersistentVolumeList `json:"body,omitempty"`
}

// NewListCoreV1PersistentVolumeOK creates ListCoreV1PersistentVolumeOK with default headers values
func NewListCoreV1PersistentVolumeOK() *ListCoreV1PersistentVolumeOK {

	return &ListCoreV1PersistentVolumeOK{}
}

// WithPayload adds the payload to the list core v1 persistent volume o k response
func (o *ListCoreV1PersistentVolumeOK) WithPayload(payload *models.IoK8sAPICoreV1PersistentVolumeList) *ListCoreV1PersistentVolumeOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list core v1 persistent volume o k response
func (o *ListCoreV1PersistentVolumeOK) SetPayload(payload *models.IoK8sAPICoreV1PersistentVolumeList) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListCoreV1PersistentVolumeOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListCoreV1PersistentVolumeUnauthorizedCode is the HTTP code returned for type ListCoreV1PersistentVolumeUnauthorized
const ListCoreV1PersistentVolumeUnauthorizedCode int = 401

/*ListCoreV1PersistentVolumeUnauthorized Unauthorized

swagger:response listCoreV1PersistentVolumeUnauthorized
*/
type ListCoreV1PersistentVolumeUnauthorized struct {
}

// NewListCoreV1PersistentVolumeUnauthorized creates ListCoreV1PersistentVolumeUnauthorized with default headers values
func NewListCoreV1PersistentVolumeUnauthorized() *ListCoreV1PersistentVolumeUnauthorized {

	return &ListCoreV1PersistentVolumeUnauthorized{}
}

// WriteResponse to the client
func (o *ListCoreV1PersistentVolumeUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
