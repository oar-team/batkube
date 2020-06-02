// Code generated by go-swagger; DO NOT EDIT.

package apps_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/models"
)

// CreateAppsV1NamespacedDeploymentOKCode is the HTTP code returned for type CreateAppsV1NamespacedDeploymentOK
const CreateAppsV1NamespacedDeploymentOKCode int = 200

/*CreateAppsV1NamespacedDeploymentOK OK

swagger:response createAppsV1NamespacedDeploymentOK
*/
type CreateAppsV1NamespacedDeploymentOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIAppsV1Deployment `json:"body,omitempty"`
}

// NewCreateAppsV1NamespacedDeploymentOK creates CreateAppsV1NamespacedDeploymentOK with default headers values
func NewCreateAppsV1NamespacedDeploymentOK() *CreateAppsV1NamespacedDeploymentOK {

	return &CreateAppsV1NamespacedDeploymentOK{}
}

// WithPayload adds the payload to the create apps v1 namespaced deployment o k response
func (o *CreateAppsV1NamespacedDeploymentOK) WithPayload(payload *models.IoK8sAPIAppsV1Deployment) *CreateAppsV1NamespacedDeploymentOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create apps v1 namespaced deployment o k response
func (o *CreateAppsV1NamespacedDeploymentOK) SetPayload(payload *models.IoK8sAPIAppsV1Deployment) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateAppsV1NamespacedDeploymentOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateAppsV1NamespacedDeploymentCreatedCode is the HTTP code returned for type CreateAppsV1NamespacedDeploymentCreated
const CreateAppsV1NamespacedDeploymentCreatedCode int = 201

/*CreateAppsV1NamespacedDeploymentCreated Created

swagger:response createAppsV1NamespacedDeploymentCreated
*/
type CreateAppsV1NamespacedDeploymentCreated struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIAppsV1Deployment `json:"body,omitempty"`
}

// NewCreateAppsV1NamespacedDeploymentCreated creates CreateAppsV1NamespacedDeploymentCreated with default headers values
func NewCreateAppsV1NamespacedDeploymentCreated() *CreateAppsV1NamespacedDeploymentCreated {

	return &CreateAppsV1NamespacedDeploymentCreated{}
}

// WithPayload adds the payload to the create apps v1 namespaced deployment created response
func (o *CreateAppsV1NamespacedDeploymentCreated) WithPayload(payload *models.IoK8sAPIAppsV1Deployment) *CreateAppsV1NamespacedDeploymentCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create apps v1 namespaced deployment created response
func (o *CreateAppsV1NamespacedDeploymentCreated) SetPayload(payload *models.IoK8sAPIAppsV1Deployment) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateAppsV1NamespacedDeploymentCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateAppsV1NamespacedDeploymentAcceptedCode is the HTTP code returned for type CreateAppsV1NamespacedDeploymentAccepted
const CreateAppsV1NamespacedDeploymentAcceptedCode int = 202

/*CreateAppsV1NamespacedDeploymentAccepted Accepted

swagger:response createAppsV1NamespacedDeploymentAccepted
*/
type CreateAppsV1NamespacedDeploymentAccepted struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIAppsV1Deployment `json:"body,omitempty"`
}

// NewCreateAppsV1NamespacedDeploymentAccepted creates CreateAppsV1NamespacedDeploymentAccepted with default headers values
func NewCreateAppsV1NamespacedDeploymentAccepted() *CreateAppsV1NamespacedDeploymentAccepted {

	return &CreateAppsV1NamespacedDeploymentAccepted{}
}

// WithPayload adds the payload to the create apps v1 namespaced deployment accepted response
func (o *CreateAppsV1NamespacedDeploymentAccepted) WithPayload(payload *models.IoK8sAPIAppsV1Deployment) *CreateAppsV1NamespacedDeploymentAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create apps v1 namespaced deployment accepted response
func (o *CreateAppsV1NamespacedDeploymentAccepted) SetPayload(payload *models.IoK8sAPIAppsV1Deployment) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateAppsV1NamespacedDeploymentAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(202)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateAppsV1NamespacedDeploymentUnauthorizedCode is the HTTP code returned for type CreateAppsV1NamespacedDeploymentUnauthorized
const CreateAppsV1NamespacedDeploymentUnauthorizedCode int = 401

/*CreateAppsV1NamespacedDeploymentUnauthorized Unauthorized

swagger:response createAppsV1NamespacedDeploymentUnauthorized
*/
type CreateAppsV1NamespacedDeploymentUnauthorized struct {
}

// NewCreateAppsV1NamespacedDeploymentUnauthorized creates CreateAppsV1NamespacedDeploymentUnauthorized with default headers values
func NewCreateAppsV1NamespacedDeploymentUnauthorized() *CreateAppsV1NamespacedDeploymentUnauthorized {

	return &CreateAppsV1NamespacedDeploymentUnauthorized{}
}

// WriteResponse to the client
func (o *CreateAppsV1NamespacedDeploymentUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
