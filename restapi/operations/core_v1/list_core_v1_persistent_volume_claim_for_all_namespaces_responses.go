// Code generated by go-swagger; DO NOT EDIT.

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/models"
)

// ListCoreV1PersistentVolumeClaimForAllNamespacesOKCode is the HTTP code returned for type ListCoreV1PersistentVolumeClaimForAllNamespacesOK
const ListCoreV1PersistentVolumeClaimForAllNamespacesOKCode int = 200

/*ListCoreV1PersistentVolumeClaimForAllNamespacesOK OK

swagger:response listCoreV1PersistentVolumeClaimForAllNamespacesOK
*/
type ListCoreV1PersistentVolumeClaimForAllNamespacesOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPICoreV1PersistentVolumeClaimList `json:"body,omitempty"`
}

// NewListCoreV1PersistentVolumeClaimForAllNamespacesOK creates ListCoreV1PersistentVolumeClaimForAllNamespacesOK with default headers values
func NewListCoreV1PersistentVolumeClaimForAllNamespacesOK() *ListCoreV1PersistentVolumeClaimForAllNamespacesOK {

	return &ListCoreV1PersistentVolumeClaimForAllNamespacesOK{}
}

// WithPayload adds the payload to the list core v1 persistent volume claim for all namespaces o k response
func (o *ListCoreV1PersistentVolumeClaimForAllNamespacesOK) WithPayload(payload *models.IoK8sAPICoreV1PersistentVolumeClaimList) *ListCoreV1PersistentVolumeClaimForAllNamespacesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list core v1 persistent volume claim for all namespaces o k response
func (o *ListCoreV1PersistentVolumeClaimForAllNamespacesOK) SetPayload(payload *models.IoK8sAPICoreV1PersistentVolumeClaimList) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListCoreV1PersistentVolumeClaimForAllNamespacesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListCoreV1PersistentVolumeClaimForAllNamespacesUnauthorizedCode is the HTTP code returned for type ListCoreV1PersistentVolumeClaimForAllNamespacesUnauthorized
const ListCoreV1PersistentVolumeClaimForAllNamespacesUnauthorizedCode int = 401

/*ListCoreV1PersistentVolumeClaimForAllNamespacesUnauthorized Unauthorized

swagger:response listCoreV1PersistentVolumeClaimForAllNamespacesUnauthorized
*/
type ListCoreV1PersistentVolumeClaimForAllNamespacesUnauthorized struct {
}

// NewListCoreV1PersistentVolumeClaimForAllNamespacesUnauthorized creates ListCoreV1PersistentVolumeClaimForAllNamespacesUnauthorized with default headers values
func NewListCoreV1PersistentVolumeClaimForAllNamespacesUnauthorized() *ListCoreV1PersistentVolumeClaimForAllNamespacesUnauthorized {

	return &ListCoreV1PersistentVolumeClaimForAllNamespacesUnauthorized{}
}

// WriteResponse to the client
func (o *ListCoreV1PersistentVolumeClaimForAllNamespacesUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}