// Code generated by go-swagger; DO NOT EDIT.

package apps_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/models"
)

// CreateAppsV1NamespacedReplicaSetOKCode is the HTTP code returned for type CreateAppsV1NamespacedReplicaSetOK
const CreateAppsV1NamespacedReplicaSetOKCode int = 200

/*CreateAppsV1NamespacedReplicaSetOK OK

swagger:response createAppsV1NamespacedReplicaSetOK
*/
type CreateAppsV1NamespacedReplicaSetOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIAppsV1ReplicaSet `json:"body,omitempty"`
}

// NewCreateAppsV1NamespacedReplicaSetOK creates CreateAppsV1NamespacedReplicaSetOK with default headers values
func NewCreateAppsV1NamespacedReplicaSetOK() *CreateAppsV1NamespacedReplicaSetOK {

	return &CreateAppsV1NamespacedReplicaSetOK{}
}

// WithPayload adds the payload to the create apps v1 namespaced replica set o k response
func (o *CreateAppsV1NamespacedReplicaSetOK) WithPayload(payload *models.IoK8sAPIAppsV1ReplicaSet) *CreateAppsV1NamespacedReplicaSetOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create apps v1 namespaced replica set o k response
func (o *CreateAppsV1NamespacedReplicaSetOK) SetPayload(payload *models.IoK8sAPIAppsV1ReplicaSet) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateAppsV1NamespacedReplicaSetOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateAppsV1NamespacedReplicaSetCreatedCode is the HTTP code returned for type CreateAppsV1NamespacedReplicaSetCreated
const CreateAppsV1NamespacedReplicaSetCreatedCode int = 201

/*CreateAppsV1NamespacedReplicaSetCreated Created

swagger:response createAppsV1NamespacedReplicaSetCreated
*/
type CreateAppsV1NamespacedReplicaSetCreated struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIAppsV1ReplicaSet `json:"body,omitempty"`
}

// NewCreateAppsV1NamespacedReplicaSetCreated creates CreateAppsV1NamespacedReplicaSetCreated with default headers values
func NewCreateAppsV1NamespacedReplicaSetCreated() *CreateAppsV1NamespacedReplicaSetCreated {

	return &CreateAppsV1NamespacedReplicaSetCreated{}
}

// WithPayload adds the payload to the create apps v1 namespaced replica set created response
func (o *CreateAppsV1NamespacedReplicaSetCreated) WithPayload(payload *models.IoK8sAPIAppsV1ReplicaSet) *CreateAppsV1NamespacedReplicaSetCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create apps v1 namespaced replica set created response
func (o *CreateAppsV1NamespacedReplicaSetCreated) SetPayload(payload *models.IoK8sAPIAppsV1ReplicaSet) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateAppsV1NamespacedReplicaSetCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateAppsV1NamespacedReplicaSetAcceptedCode is the HTTP code returned for type CreateAppsV1NamespacedReplicaSetAccepted
const CreateAppsV1NamespacedReplicaSetAcceptedCode int = 202

/*CreateAppsV1NamespacedReplicaSetAccepted Accepted

swagger:response createAppsV1NamespacedReplicaSetAccepted
*/
type CreateAppsV1NamespacedReplicaSetAccepted struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIAppsV1ReplicaSet `json:"body,omitempty"`
}

// NewCreateAppsV1NamespacedReplicaSetAccepted creates CreateAppsV1NamespacedReplicaSetAccepted with default headers values
func NewCreateAppsV1NamespacedReplicaSetAccepted() *CreateAppsV1NamespacedReplicaSetAccepted {

	return &CreateAppsV1NamespacedReplicaSetAccepted{}
}

// WithPayload adds the payload to the create apps v1 namespaced replica set accepted response
func (o *CreateAppsV1NamespacedReplicaSetAccepted) WithPayload(payload *models.IoK8sAPIAppsV1ReplicaSet) *CreateAppsV1NamespacedReplicaSetAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create apps v1 namespaced replica set accepted response
func (o *CreateAppsV1NamespacedReplicaSetAccepted) SetPayload(payload *models.IoK8sAPIAppsV1ReplicaSet) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateAppsV1NamespacedReplicaSetAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(202)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateAppsV1NamespacedReplicaSetUnauthorizedCode is the HTTP code returned for type CreateAppsV1NamespacedReplicaSetUnauthorized
const CreateAppsV1NamespacedReplicaSetUnauthorizedCode int = 401

/*CreateAppsV1NamespacedReplicaSetUnauthorized Unauthorized

swagger:response createAppsV1NamespacedReplicaSetUnauthorized
*/
type CreateAppsV1NamespacedReplicaSetUnauthorized struct {
}

// NewCreateAppsV1NamespacedReplicaSetUnauthorized creates CreateAppsV1NamespacedReplicaSetUnauthorized with default headers values
func NewCreateAppsV1NamespacedReplicaSetUnauthorized() *CreateAppsV1NamespacedReplicaSetUnauthorized {

	return &CreateAppsV1NamespacedReplicaSetUnauthorized{}
}

// WriteResponse to the client
func (o *CreateAppsV1NamespacedReplicaSetUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}