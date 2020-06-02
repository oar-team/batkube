// Code generated by go-swagger; DO NOT EDIT.

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/models"
)

// ReplaceCoreV1NamespacedReplicationControllerScaleOKCode is the HTTP code returned for type ReplaceCoreV1NamespacedReplicationControllerScaleOK
const ReplaceCoreV1NamespacedReplicationControllerScaleOKCode int = 200

/*ReplaceCoreV1NamespacedReplicationControllerScaleOK OK

swagger:response replaceCoreV1NamespacedReplicationControllerScaleOK
*/
type ReplaceCoreV1NamespacedReplicationControllerScaleOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIAutoscalingV1Scale `json:"body,omitempty"`
}

// NewReplaceCoreV1NamespacedReplicationControllerScaleOK creates ReplaceCoreV1NamespacedReplicationControllerScaleOK with default headers values
func NewReplaceCoreV1NamespacedReplicationControllerScaleOK() *ReplaceCoreV1NamespacedReplicationControllerScaleOK {

	return &ReplaceCoreV1NamespacedReplicationControllerScaleOK{}
}

// WithPayload adds the payload to the replace core v1 namespaced replication controller scale o k response
func (o *ReplaceCoreV1NamespacedReplicationControllerScaleOK) WithPayload(payload *models.IoK8sAPIAutoscalingV1Scale) *ReplaceCoreV1NamespacedReplicationControllerScaleOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace core v1 namespaced replication controller scale o k response
func (o *ReplaceCoreV1NamespacedReplicationControllerScaleOK) SetPayload(payload *models.IoK8sAPIAutoscalingV1Scale) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceCoreV1NamespacedReplicationControllerScaleOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceCoreV1NamespacedReplicationControllerScaleCreatedCode is the HTTP code returned for type ReplaceCoreV1NamespacedReplicationControllerScaleCreated
const ReplaceCoreV1NamespacedReplicationControllerScaleCreatedCode int = 201

/*ReplaceCoreV1NamespacedReplicationControllerScaleCreated Created

swagger:response replaceCoreV1NamespacedReplicationControllerScaleCreated
*/
type ReplaceCoreV1NamespacedReplicationControllerScaleCreated struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIAutoscalingV1Scale `json:"body,omitempty"`
}

// NewReplaceCoreV1NamespacedReplicationControllerScaleCreated creates ReplaceCoreV1NamespacedReplicationControllerScaleCreated with default headers values
func NewReplaceCoreV1NamespacedReplicationControllerScaleCreated() *ReplaceCoreV1NamespacedReplicationControllerScaleCreated {

	return &ReplaceCoreV1NamespacedReplicationControllerScaleCreated{}
}

// WithPayload adds the payload to the replace core v1 namespaced replication controller scale created response
func (o *ReplaceCoreV1NamespacedReplicationControllerScaleCreated) WithPayload(payload *models.IoK8sAPIAutoscalingV1Scale) *ReplaceCoreV1NamespacedReplicationControllerScaleCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace core v1 namespaced replication controller scale created response
func (o *ReplaceCoreV1NamespacedReplicationControllerScaleCreated) SetPayload(payload *models.IoK8sAPIAutoscalingV1Scale) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceCoreV1NamespacedReplicationControllerScaleCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceCoreV1NamespacedReplicationControllerScaleUnauthorizedCode is the HTTP code returned for type ReplaceCoreV1NamespacedReplicationControllerScaleUnauthorized
const ReplaceCoreV1NamespacedReplicationControllerScaleUnauthorizedCode int = 401

/*ReplaceCoreV1NamespacedReplicationControllerScaleUnauthorized Unauthorized

swagger:response replaceCoreV1NamespacedReplicationControllerScaleUnauthorized
*/
type ReplaceCoreV1NamespacedReplicationControllerScaleUnauthorized struct {
}

// NewReplaceCoreV1NamespacedReplicationControllerScaleUnauthorized creates ReplaceCoreV1NamespacedReplicationControllerScaleUnauthorized with default headers values
func NewReplaceCoreV1NamespacedReplicationControllerScaleUnauthorized() *ReplaceCoreV1NamespacedReplicationControllerScaleUnauthorized {

	return &ReplaceCoreV1NamespacedReplicationControllerScaleUnauthorized{}
}

// WriteResponse to the client
func (o *ReplaceCoreV1NamespacedReplicationControllerScaleUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
