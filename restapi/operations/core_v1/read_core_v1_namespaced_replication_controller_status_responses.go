// Code generated by go-swagger; DO NOT EDIT.

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// ReadCoreV1NamespacedReplicationControllerStatusOKCode is the HTTP code returned for type ReadCoreV1NamespacedReplicationControllerStatusOK
const ReadCoreV1NamespacedReplicationControllerStatusOKCode int = 200

/*ReadCoreV1NamespacedReplicationControllerStatusOK OK

swagger:response readCoreV1NamespacedReplicationControllerStatusOK
*/
type ReadCoreV1NamespacedReplicationControllerStatusOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPICoreV1ReplicationController `json:"body,omitempty"`
}

// NewReadCoreV1NamespacedReplicationControllerStatusOK creates ReadCoreV1NamespacedReplicationControllerStatusOK with default headers values
func NewReadCoreV1NamespacedReplicationControllerStatusOK() *ReadCoreV1NamespacedReplicationControllerStatusOK {

	return &ReadCoreV1NamespacedReplicationControllerStatusOK{}
}

// WithPayload adds the payload to the read core v1 namespaced replication controller status o k response
func (o *ReadCoreV1NamespacedReplicationControllerStatusOK) WithPayload(payload *models.IoK8sAPICoreV1ReplicationController) *ReadCoreV1NamespacedReplicationControllerStatusOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the read core v1 namespaced replication controller status o k response
func (o *ReadCoreV1NamespacedReplicationControllerStatusOK) SetPayload(payload *models.IoK8sAPICoreV1ReplicationController) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReadCoreV1NamespacedReplicationControllerStatusOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReadCoreV1NamespacedReplicationControllerStatusUnauthorizedCode is the HTTP code returned for type ReadCoreV1NamespacedReplicationControllerStatusUnauthorized
const ReadCoreV1NamespacedReplicationControllerStatusUnauthorizedCode int = 401

/*ReadCoreV1NamespacedReplicationControllerStatusUnauthorized Unauthorized

swagger:response readCoreV1NamespacedReplicationControllerStatusUnauthorized
*/
type ReadCoreV1NamespacedReplicationControllerStatusUnauthorized struct {
}

// NewReadCoreV1NamespacedReplicationControllerStatusUnauthorized creates ReadCoreV1NamespacedReplicationControllerStatusUnauthorized with default headers values
func NewReadCoreV1NamespacedReplicationControllerStatusUnauthorized() *ReadCoreV1NamespacedReplicationControllerStatusUnauthorized {

	return &ReadCoreV1NamespacedReplicationControllerStatusUnauthorized{}
}

// WriteResponse to the client
func (o *ReadCoreV1NamespacedReplicationControllerStatusUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
