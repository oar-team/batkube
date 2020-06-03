// Code generated by go-swagger; DO NOT EDIT.

package apps_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/models"
)

// ReadAppsV1NamespacedDeploymentStatusOKCode is the HTTP code returned for type ReadAppsV1NamespacedDeploymentStatusOK
const ReadAppsV1NamespacedDeploymentStatusOKCode int = 200

/*ReadAppsV1NamespacedDeploymentStatusOK OK

swagger:response readAppsV1NamespacedDeploymentStatusOK
*/
type ReadAppsV1NamespacedDeploymentStatusOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIAppsV1Deployment `json:"body,omitempty"`
}

// NewReadAppsV1NamespacedDeploymentStatusOK creates ReadAppsV1NamespacedDeploymentStatusOK with default headers values
func NewReadAppsV1NamespacedDeploymentStatusOK() *ReadAppsV1NamespacedDeploymentStatusOK {

	return &ReadAppsV1NamespacedDeploymentStatusOK{}
}

// WithPayload adds the payload to the read apps v1 namespaced deployment status o k response
func (o *ReadAppsV1NamespacedDeploymentStatusOK) WithPayload(payload *models.IoK8sAPIAppsV1Deployment) *ReadAppsV1NamespacedDeploymentStatusOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the read apps v1 namespaced deployment status o k response
func (o *ReadAppsV1NamespacedDeploymentStatusOK) SetPayload(payload *models.IoK8sAPIAppsV1Deployment) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReadAppsV1NamespacedDeploymentStatusOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReadAppsV1NamespacedDeploymentStatusUnauthorizedCode is the HTTP code returned for type ReadAppsV1NamespacedDeploymentStatusUnauthorized
const ReadAppsV1NamespacedDeploymentStatusUnauthorizedCode int = 401

/*ReadAppsV1NamespacedDeploymentStatusUnauthorized Unauthorized

swagger:response readAppsV1NamespacedDeploymentStatusUnauthorized
*/
type ReadAppsV1NamespacedDeploymentStatusUnauthorized struct {
}

// NewReadAppsV1NamespacedDeploymentStatusUnauthorized creates ReadAppsV1NamespacedDeploymentStatusUnauthorized with default headers values
func NewReadAppsV1NamespacedDeploymentStatusUnauthorized() *ReadAppsV1NamespacedDeploymentStatusUnauthorized {

	return &ReadAppsV1NamespacedDeploymentStatusUnauthorized{}
}

// WriteResponse to the client
func (o *ReadAppsV1NamespacedDeploymentStatusUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}