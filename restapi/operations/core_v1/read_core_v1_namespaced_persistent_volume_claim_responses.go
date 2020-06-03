// Code generated by go-swagger; DO NOT EDIT.

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/models"
)

// ReadCoreV1NamespacedPersistentVolumeClaimOKCode is the HTTP code returned for type ReadCoreV1NamespacedPersistentVolumeClaimOK
const ReadCoreV1NamespacedPersistentVolumeClaimOKCode int = 200

/*ReadCoreV1NamespacedPersistentVolumeClaimOK OK

swagger:response readCoreV1NamespacedPersistentVolumeClaimOK
*/
type ReadCoreV1NamespacedPersistentVolumeClaimOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPICoreV1PersistentVolumeClaim `json:"body,omitempty"`
}

// NewReadCoreV1NamespacedPersistentVolumeClaimOK creates ReadCoreV1NamespacedPersistentVolumeClaimOK with default headers values
func NewReadCoreV1NamespacedPersistentVolumeClaimOK() *ReadCoreV1NamespacedPersistentVolumeClaimOK {

	return &ReadCoreV1NamespacedPersistentVolumeClaimOK{}
}

// WithPayload adds the payload to the read core v1 namespaced persistent volume claim o k response
func (o *ReadCoreV1NamespacedPersistentVolumeClaimOK) WithPayload(payload *models.IoK8sAPICoreV1PersistentVolumeClaim) *ReadCoreV1NamespacedPersistentVolumeClaimOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the read core v1 namespaced persistent volume claim o k response
func (o *ReadCoreV1NamespacedPersistentVolumeClaimOK) SetPayload(payload *models.IoK8sAPICoreV1PersistentVolumeClaim) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReadCoreV1NamespacedPersistentVolumeClaimOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReadCoreV1NamespacedPersistentVolumeClaimUnauthorizedCode is the HTTP code returned for type ReadCoreV1NamespacedPersistentVolumeClaimUnauthorized
const ReadCoreV1NamespacedPersistentVolumeClaimUnauthorizedCode int = 401

/*ReadCoreV1NamespacedPersistentVolumeClaimUnauthorized Unauthorized

swagger:response readCoreV1NamespacedPersistentVolumeClaimUnauthorized
*/
type ReadCoreV1NamespacedPersistentVolumeClaimUnauthorized struct {
}

// NewReadCoreV1NamespacedPersistentVolumeClaimUnauthorized creates ReadCoreV1NamespacedPersistentVolumeClaimUnauthorized with default headers values
func NewReadCoreV1NamespacedPersistentVolumeClaimUnauthorized() *ReadCoreV1NamespacedPersistentVolumeClaimUnauthorized {

	return &ReadCoreV1NamespacedPersistentVolumeClaimUnauthorized{}
}

// WriteResponse to the client
func (o *ReadCoreV1NamespacedPersistentVolumeClaimUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}