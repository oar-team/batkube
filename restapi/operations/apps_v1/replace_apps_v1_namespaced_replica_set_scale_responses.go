// Code generated by go-swagger; DO NOT EDIT.

package apps_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// ReplaceAppsV1NamespacedReplicaSetScaleOKCode is the HTTP code returned for type ReplaceAppsV1NamespacedReplicaSetScaleOK
const ReplaceAppsV1NamespacedReplicaSetScaleOKCode int = 200

/*ReplaceAppsV1NamespacedReplicaSetScaleOK OK

swagger:response replaceAppsV1NamespacedReplicaSetScaleOK
*/
type ReplaceAppsV1NamespacedReplicaSetScaleOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIAutoscalingV1Scale `json:"body,omitempty"`
}

// NewReplaceAppsV1NamespacedReplicaSetScaleOK creates ReplaceAppsV1NamespacedReplicaSetScaleOK with default headers values
func NewReplaceAppsV1NamespacedReplicaSetScaleOK() *ReplaceAppsV1NamespacedReplicaSetScaleOK {

	return &ReplaceAppsV1NamespacedReplicaSetScaleOK{}
}

// WithPayload adds the payload to the replace apps v1 namespaced replica set scale o k response
func (o *ReplaceAppsV1NamespacedReplicaSetScaleOK) WithPayload(payload *models.IoK8sAPIAutoscalingV1Scale) *ReplaceAppsV1NamespacedReplicaSetScaleOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace apps v1 namespaced replica set scale o k response
func (o *ReplaceAppsV1NamespacedReplicaSetScaleOK) SetPayload(payload *models.IoK8sAPIAutoscalingV1Scale) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceAppsV1NamespacedReplicaSetScaleOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceAppsV1NamespacedReplicaSetScaleCreatedCode is the HTTP code returned for type ReplaceAppsV1NamespacedReplicaSetScaleCreated
const ReplaceAppsV1NamespacedReplicaSetScaleCreatedCode int = 201

/*ReplaceAppsV1NamespacedReplicaSetScaleCreated Created

swagger:response replaceAppsV1NamespacedReplicaSetScaleCreated
*/
type ReplaceAppsV1NamespacedReplicaSetScaleCreated struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIAutoscalingV1Scale `json:"body,omitempty"`
}

// NewReplaceAppsV1NamespacedReplicaSetScaleCreated creates ReplaceAppsV1NamespacedReplicaSetScaleCreated with default headers values
func NewReplaceAppsV1NamespacedReplicaSetScaleCreated() *ReplaceAppsV1NamespacedReplicaSetScaleCreated {

	return &ReplaceAppsV1NamespacedReplicaSetScaleCreated{}
}

// WithPayload adds the payload to the replace apps v1 namespaced replica set scale created response
func (o *ReplaceAppsV1NamespacedReplicaSetScaleCreated) WithPayload(payload *models.IoK8sAPIAutoscalingV1Scale) *ReplaceAppsV1NamespacedReplicaSetScaleCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace apps v1 namespaced replica set scale created response
func (o *ReplaceAppsV1NamespacedReplicaSetScaleCreated) SetPayload(payload *models.IoK8sAPIAutoscalingV1Scale) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceAppsV1NamespacedReplicaSetScaleCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceAppsV1NamespacedReplicaSetScaleUnauthorizedCode is the HTTP code returned for type ReplaceAppsV1NamespacedReplicaSetScaleUnauthorized
const ReplaceAppsV1NamespacedReplicaSetScaleUnauthorizedCode int = 401

/*ReplaceAppsV1NamespacedReplicaSetScaleUnauthorized Unauthorized

swagger:response replaceAppsV1NamespacedReplicaSetScaleUnauthorized
*/
type ReplaceAppsV1NamespacedReplicaSetScaleUnauthorized struct {
}

// NewReplaceAppsV1NamespacedReplicaSetScaleUnauthorized creates ReplaceAppsV1NamespacedReplicaSetScaleUnauthorized with default headers values
func NewReplaceAppsV1NamespacedReplicaSetScaleUnauthorized() *ReplaceAppsV1NamespacedReplicaSetScaleUnauthorized {

	return &ReplaceAppsV1NamespacedReplicaSetScaleUnauthorized{}
}

// WriteResponse to the client
func (o *ReplaceAppsV1NamespacedReplicaSetScaleUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
