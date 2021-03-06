// Code generated by go-swagger; DO NOT EDIT.

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// ReplaceCoreV1NamespacedReplicationControllerOKCode is the HTTP code returned for type ReplaceCoreV1NamespacedReplicationControllerOK
const ReplaceCoreV1NamespacedReplicationControllerOKCode int = 200

/*ReplaceCoreV1NamespacedReplicationControllerOK OK

swagger:response replaceCoreV1NamespacedReplicationControllerOK
*/
type ReplaceCoreV1NamespacedReplicationControllerOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPICoreV1ReplicationController `json:"body,omitempty"`
}

// NewReplaceCoreV1NamespacedReplicationControllerOK creates ReplaceCoreV1NamespacedReplicationControllerOK with default headers values
func NewReplaceCoreV1NamespacedReplicationControllerOK() *ReplaceCoreV1NamespacedReplicationControllerOK {

	return &ReplaceCoreV1NamespacedReplicationControllerOK{}
}

// WithPayload adds the payload to the replace core v1 namespaced replication controller o k response
func (o *ReplaceCoreV1NamespacedReplicationControllerOK) WithPayload(payload *models.IoK8sAPICoreV1ReplicationController) *ReplaceCoreV1NamespacedReplicationControllerOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace core v1 namespaced replication controller o k response
func (o *ReplaceCoreV1NamespacedReplicationControllerOK) SetPayload(payload *models.IoK8sAPICoreV1ReplicationController) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceCoreV1NamespacedReplicationControllerOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceCoreV1NamespacedReplicationControllerCreatedCode is the HTTP code returned for type ReplaceCoreV1NamespacedReplicationControllerCreated
const ReplaceCoreV1NamespacedReplicationControllerCreatedCode int = 201

/*ReplaceCoreV1NamespacedReplicationControllerCreated Created

swagger:response replaceCoreV1NamespacedReplicationControllerCreated
*/
type ReplaceCoreV1NamespacedReplicationControllerCreated struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPICoreV1ReplicationController `json:"body,omitempty"`
}

// NewReplaceCoreV1NamespacedReplicationControllerCreated creates ReplaceCoreV1NamespacedReplicationControllerCreated with default headers values
func NewReplaceCoreV1NamespacedReplicationControllerCreated() *ReplaceCoreV1NamespacedReplicationControllerCreated {

	return &ReplaceCoreV1NamespacedReplicationControllerCreated{}
}

// WithPayload adds the payload to the replace core v1 namespaced replication controller created response
func (o *ReplaceCoreV1NamespacedReplicationControllerCreated) WithPayload(payload *models.IoK8sAPICoreV1ReplicationController) *ReplaceCoreV1NamespacedReplicationControllerCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace core v1 namespaced replication controller created response
func (o *ReplaceCoreV1NamespacedReplicationControllerCreated) SetPayload(payload *models.IoK8sAPICoreV1ReplicationController) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceCoreV1NamespacedReplicationControllerCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceCoreV1NamespacedReplicationControllerUnauthorizedCode is the HTTP code returned for type ReplaceCoreV1NamespacedReplicationControllerUnauthorized
const ReplaceCoreV1NamespacedReplicationControllerUnauthorizedCode int = 401

/*ReplaceCoreV1NamespacedReplicationControllerUnauthorized Unauthorized

swagger:response replaceCoreV1NamespacedReplicationControllerUnauthorized
*/
type ReplaceCoreV1NamespacedReplicationControllerUnauthorized struct {
}

// NewReplaceCoreV1NamespacedReplicationControllerUnauthorized creates ReplaceCoreV1NamespacedReplicationControllerUnauthorized with default headers values
func NewReplaceCoreV1NamespacedReplicationControllerUnauthorized() *ReplaceCoreV1NamespacedReplicationControllerUnauthorized {

	return &ReplaceCoreV1NamespacedReplicationControllerUnauthorized{}
}

// WriteResponse to the client
func (o *ReplaceCoreV1NamespacedReplicationControllerUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
