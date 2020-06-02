// Code generated by go-swagger; DO NOT EDIT.

package apps_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/models"
)

// ReadAppsV1NamespacedDeploymentOKCode is the HTTP code returned for type ReadAppsV1NamespacedDeploymentOK
const ReadAppsV1NamespacedDeploymentOKCode int = 200

/*ReadAppsV1NamespacedDeploymentOK OK

swagger:response readAppsV1NamespacedDeploymentOK
*/
type ReadAppsV1NamespacedDeploymentOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIAppsV1Deployment `json:"body,omitempty"`
}

// NewReadAppsV1NamespacedDeploymentOK creates ReadAppsV1NamespacedDeploymentOK with default headers values
func NewReadAppsV1NamespacedDeploymentOK() *ReadAppsV1NamespacedDeploymentOK {

	return &ReadAppsV1NamespacedDeploymentOK{}
}

// WithPayload adds the payload to the read apps v1 namespaced deployment o k response
func (o *ReadAppsV1NamespacedDeploymentOK) WithPayload(payload *models.IoK8sAPIAppsV1Deployment) *ReadAppsV1NamespacedDeploymentOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the read apps v1 namespaced deployment o k response
func (o *ReadAppsV1NamespacedDeploymentOK) SetPayload(payload *models.IoK8sAPIAppsV1Deployment) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReadAppsV1NamespacedDeploymentOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReadAppsV1NamespacedDeploymentUnauthorizedCode is the HTTP code returned for type ReadAppsV1NamespacedDeploymentUnauthorized
const ReadAppsV1NamespacedDeploymentUnauthorizedCode int = 401

/*ReadAppsV1NamespacedDeploymentUnauthorized Unauthorized

swagger:response readAppsV1NamespacedDeploymentUnauthorized
*/
type ReadAppsV1NamespacedDeploymentUnauthorized struct {
}

// NewReadAppsV1NamespacedDeploymentUnauthorized creates ReadAppsV1NamespacedDeploymentUnauthorized with default headers values
func NewReadAppsV1NamespacedDeploymentUnauthorized() *ReadAppsV1NamespacedDeploymentUnauthorized {

	return &ReadAppsV1NamespacedDeploymentUnauthorized{}
}

// WriteResponse to the client
func (o *ReadAppsV1NamespacedDeploymentUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
