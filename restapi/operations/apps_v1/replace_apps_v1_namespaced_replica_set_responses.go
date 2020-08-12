// Code generated by go-swagger; DO NOT EDIT.

package apps_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// ReplaceAppsV1NamespacedReplicaSetOKCode is the HTTP code returned for type ReplaceAppsV1NamespacedReplicaSetOK
const ReplaceAppsV1NamespacedReplicaSetOKCode int = 200

/*ReplaceAppsV1NamespacedReplicaSetOK OK

swagger:response replaceAppsV1NamespacedReplicaSetOK
*/
type ReplaceAppsV1NamespacedReplicaSetOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIAppsV1ReplicaSet `json:"body,omitempty"`
}

// NewReplaceAppsV1NamespacedReplicaSetOK creates ReplaceAppsV1NamespacedReplicaSetOK with default headers values
func NewReplaceAppsV1NamespacedReplicaSetOK() *ReplaceAppsV1NamespacedReplicaSetOK {

	return &ReplaceAppsV1NamespacedReplicaSetOK{}
}

// WithPayload adds the payload to the replace apps v1 namespaced replica set o k response
func (o *ReplaceAppsV1NamespacedReplicaSetOK) WithPayload(payload *models.IoK8sAPIAppsV1ReplicaSet) *ReplaceAppsV1NamespacedReplicaSetOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace apps v1 namespaced replica set o k response
func (o *ReplaceAppsV1NamespacedReplicaSetOK) SetPayload(payload *models.IoK8sAPIAppsV1ReplicaSet) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceAppsV1NamespacedReplicaSetOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceAppsV1NamespacedReplicaSetCreatedCode is the HTTP code returned for type ReplaceAppsV1NamespacedReplicaSetCreated
const ReplaceAppsV1NamespacedReplicaSetCreatedCode int = 201

/*ReplaceAppsV1NamespacedReplicaSetCreated Created

swagger:response replaceAppsV1NamespacedReplicaSetCreated
*/
type ReplaceAppsV1NamespacedReplicaSetCreated struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIAppsV1ReplicaSet `json:"body,omitempty"`
}

// NewReplaceAppsV1NamespacedReplicaSetCreated creates ReplaceAppsV1NamespacedReplicaSetCreated with default headers values
func NewReplaceAppsV1NamespacedReplicaSetCreated() *ReplaceAppsV1NamespacedReplicaSetCreated {

	return &ReplaceAppsV1NamespacedReplicaSetCreated{}
}

// WithPayload adds the payload to the replace apps v1 namespaced replica set created response
func (o *ReplaceAppsV1NamespacedReplicaSetCreated) WithPayload(payload *models.IoK8sAPIAppsV1ReplicaSet) *ReplaceAppsV1NamespacedReplicaSetCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace apps v1 namespaced replica set created response
func (o *ReplaceAppsV1NamespacedReplicaSetCreated) SetPayload(payload *models.IoK8sAPIAppsV1ReplicaSet) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceAppsV1NamespacedReplicaSetCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceAppsV1NamespacedReplicaSetUnauthorizedCode is the HTTP code returned for type ReplaceAppsV1NamespacedReplicaSetUnauthorized
const ReplaceAppsV1NamespacedReplicaSetUnauthorizedCode int = 401

/*ReplaceAppsV1NamespacedReplicaSetUnauthorized Unauthorized

swagger:response replaceAppsV1NamespacedReplicaSetUnauthorized
*/
type ReplaceAppsV1NamespacedReplicaSetUnauthorized struct {
}

// NewReplaceAppsV1NamespacedReplicaSetUnauthorized creates ReplaceAppsV1NamespacedReplicaSetUnauthorized with default headers values
func NewReplaceAppsV1NamespacedReplicaSetUnauthorized() *ReplaceAppsV1NamespacedReplicaSetUnauthorized {

	return &ReplaceAppsV1NamespacedReplicaSetUnauthorized{}
}

// WriteResponse to the client
func (o *ReplaceAppsV1NamespacedReplicaSetUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
