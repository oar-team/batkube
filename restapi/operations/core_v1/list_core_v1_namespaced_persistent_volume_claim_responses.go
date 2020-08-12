// Code generated by go-swagger; DO NOT EDIT.

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// ListCoreV1NamespacedPersistentVolumeClaimOKCode is the HTTP code returned for type ListCoreV1NamespacedPersistentVolumeClaimOK
const ListCoreV1NamespacedPersistentVolumeClaimOKCode int = 200

/*ListCoreV1NamespacedPersistentVolumeClaimOK OK

swagger:response listCoreV1NamespacedPersistentVolumeClaimOK
*/
type ListCoreV1NamespacedPersistentVolumeClaimOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPICoreV1PersistentVolumeClaimList `json:"body,omitempty"`
}

// NewListCoreV1NamespacedPersistentVolumeClaimOK creates ListCoreV1NamespacedPersistentVolumeClaimOK with default headers values
func NewListCoreV1NamespacedPersistentVolumeClaimOK() *ListCoreV1NamespacedPersistentVolumeClaimOK {

	return &ListCoreV1NamespacedPersistentVolumeClaimOK{}
}

// WithPayload adds the payload to the list core v1 namespaced persistent volume claim o k response
func (o *ListCoreV1NamespacedPersistentVolumeClaimOK) WithPayload(payload *models.IoK8sAPICoreV1PersistentVolumeClaimList) *ListCoreV1NamespacedPersistentVolumeClaimOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list core v1 namespaced persistent volume claim o k response
func (o *ListCoreV1NamespacedPersistentVolumeClaimOK) SetPayload(payload *models.IoK8sAPICoreV1PersistentVolumeClaimList) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListCoreV1NamespacedPersistentVolumeClaimOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListCoreV1NamespacedPersistentVolumeClaimUnauthorizedCode is the HTTP code returned for type ListCoreV1NamespacedPersistentVolumeClaimUnauthorized
const ListCoreV1NamespacedPersistentVolumeClaimUnauthorizedCode int = 401

/*ListCoreV1NamespacedPersistentVolumeClaimUnauthorized Unauthorized

swagger:response listCoreV1NamespacedPersistentVolumeClaimUnauthorized
*/
type ListCoreV1NamespacedPersistentVolumeClaimUnauthorized struct {
}

// NewListCoreV1NamespacedPersistentVolumeClaimUnauthorized creates ListCoreV1NamespacedPersistentVolumeClaimUnauthorized with default headers values
func NewListCoreV1NamespacedPersistentVolumeClaimUnauthorized() *ListCoreV1NamespacedPersistentVolumeClaimUnauthorized {

	return &ListCoreV1NamespacedPersistentVolumeClaimUnauthorized{}
}

// WriteResponse to the client
func (o *ListCoreV1NamespacedPersistentVolumeClaimUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
