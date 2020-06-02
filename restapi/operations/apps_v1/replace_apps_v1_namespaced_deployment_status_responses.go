// Code generated by go-swagger; DO NOT EDIT.

package apps_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/models"
)

// ReplaceAppsV1NamespacedDeploymentStatusOKCode is the HTTP code returned for type ReplaceAppsV1NamespacedDeploymentStatusOK
const ReplaceAppsV1NamespacedDeploymentStatusOKCode int = 200

/*ReplaceAppsV1NamespacedDeploymentStatusOK OK

swagger:response replaceAppsV1NamespacedDeploymentStatusOK
*/
type ReplaceAppsV1NamespacedDeploymentStatusOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIAppsV1Deployment `json:"body,omitempty"`
}

// NewReplaceAppsV1NamespacedDeploymentStatusOK creates ReplaceAppsV1NamespacedDeploymentStatusOK with default headers values
func NewReplaceAppsV1NamespacedDeploymentStatusOK() *ReplaceAppsV1NamespacedDeploymentStatusOK {

	return &ReplaceAppsV1NamespacedDeploymentStatusOK{}
}

// WithPayload adds the payload to the replace apps v1 namespaced deployment status o k response
func (o *ReplaceAppsV1NamespacedDeploymentStatusOK) WithPayload(payload *models.IoK8sAPIAppsV1Deployment) *ReplaceAppsV1NamespacedDeploymentStatusOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace apps v1 namespaced deployment status o k response
func (o *ReplaceAppsV1NamespacedDeploymentStatusOK) SetPayload(payload *models.IoK8sAPIAppsV1Deployment) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceAppsV1NamespacedDeploymentStatusOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceAppsV1NamespacedDeploymentStatusCreatedCode is the HTTP code returned for type ReplaceAppsV1NamespacedDeploymentStatusCreated
const ReplaceAppsV1NamespacedDeploymentStatusCreatedCode int = 201

/*ReplaceAppsV1NamespacedDeploymentStatusCreated Created

swagger:response replaceAppsV1NamespacedDeploymentStatusCreated
*/
type ReplaceAppsV1NamespacedDeploymentStatusCreated struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIAppsV1Deployment `json:"body,omitempty"`
}

// NewReplaceAppsV1NamespacedDeploymentStatusCreated creates ReplaceAppsV1NamespacedDeploymentStatusCreated with default headers values
func NewReplaceAppsV1NamespacedDeploymentStatusCreated() *ReplaceAppsV1NamespacedDeploymentStatusCreated {

	return &ReplaceAppsV1NamespacedDeploymentStatusCreated{}
}

// WithPayload adds the payload to the replace apps v1 namespaced deployment status created response
func (o *ReplaceAppsV1NamespacedDeploymentStatusCreated) WithPayload(payload *models.IoK8sAPIAppsV1Deployment) *ReplaceAppsV1NamespacedDeploymentStatusCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace apps v1 namespaced deployment status created response
func (o *ReplaceAppsV1NamespacedDeploymentStatusCreated) SetPayload(payload *models.IoK8sAPIAppsV1Deployment) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceAppsV1NamespacedDeploymentStatusCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceAppsV1NamespacedDeploymentStatusUnauthorizedCode is the HTTP code returned for type ReplaceAppsV1NamespacedDeploymentStatusUnauthorized
const ReplaceAppsV1NamespacedDeploymentStatusUnauthorizedCode int = 401

/*ReplaceAppsV1NamespacedDeploymentStatusUnauthorized Unauthorized

swagger:response replaceAppsV1NamespacedDeploymentStatusUnauthorized
*/
type ReplaceAppsV1NamespacedDeploymentStatusUnauthorized struct {
}

// NewReplaceAppsV1NamespacedDeploymentStatusUnauthorized creates ReplaceAppsV1NamespacedDeploymentStatusUnauthorized with default headers values
func NewReplaceAppsV1NamespacedDeploymentStatusUnauthorized() *ReplaceAppsV1NamespacedDeploymentStatusUnauthorized {

	return &ReplaceAppsV1NamespacedDeploymentStatusUnauthorized{}
}

// WriteResponse to the client
func (o *ReplaceAppsV1NamespacedDeploymentStatusUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
