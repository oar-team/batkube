// Code generated by go-swagger; DO NOT EDIT.

package apiextensions_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/models"
)

// ReadApiextensionsV1beta1CustomResourceDefinitionStatusOKCode is the HTTP code returned for type ReadApiextensionsV1beta1CustomResourceDefinitionStatusOK
const ReadApiextensionsV1beta1CustomResourceDefinitionStatusOKCode int = 200

/*ReadApiextensionsV1beta1CustomResourceDefinitionStatusOK OK

swagger:response readApiextensionsV1beta1CustomResourceDefinitionStatusOK
*/
type ReadApiextensionsV1beta1CustomResourceDefinitionStatusOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApiextensionsApiserverPkgApisApiextensionsV1beta1CustomResourceDefinition `json:"body,omitempty"`
}

// NewReadApiextensionsV1beta1CustomResourceDefinitionStatusOK creates ReadApiextensionsV1beta1CustomResourceDefinitionStatusOK with default headers values
func NewReadApiextensionsV1beta1CustomResourceDefinitionStatusOK() *ReadApiextensionsV1beta1CustomResourceDefinitionStatusOK {

	return &ReadApiextensionsV1beta1CustomResourceDefinitionStatusOK{}
}

// WithPayload adds the payload to the read apiextensions v1beta1 custom resource definition status o k response
func (o *ReadApiextensionsV1beta1CustomResourceDefinitionStatusOK) WithPayload(payload *models.IoK8sApiextensionsApiserverPkgApisApiextensionsV1beta1CustomResourceDefinition) *ReadApiextensionsV1beta1CustomResourceDefinitionStatusOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the read apiextensions v1beta1 custom resource definition status o k response
func (o *ReadApiextensionsV1beta1CustomResourceDefinitionStatusOK) SetPayload(payload *models.IoK8sApiextensionsApiserverPkgApisApiextensionsV1beta1CustomResourceDefinition) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReadApiextensionsV1beta1CustomResourceDefinitionStatusOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReadApiextensionsV1beta1CustomResourceDefinitionStatusUnauthorizedCode is the HTTP code returned for type ReadApiextensionsV1beta1CustomResourceDefinitionStatusUnauthorized
const ReadApiextensionsV1beta1CustomResourceDefinitionStatusUnauthorizedCode int = 401

/*ReadApiextensionsV1beta1CustomResourceDefinitionStatusUnauthorized Unauthorized

swagger:response readApiextensionsV1beta1CustomResourceDefinitionStatusUnauthorized
*/
type ReadApiextensionsV1beta1CustomResourceDefinitionStatusUnauthorized struct {
}

// NewReadApiextensionsV1beta1CustomResourceDefinitionStatusUnauthorized creates ReadApiextensionsV1beta1CustomResourceDefinitionStatusUnauthorized with default headers values
func NewReadApiextensionsV1beta1CustomResourceDefinitionStatusUnauthorized() *ReadApiextensionsV1beta1CustomResourceDefinitionStatusUnauthorized {

	return &ReadApiextensionsV1beta1CustomResourceDefinitionStatusUnauthorized{}
}

// WriteResponse to the client
func (o *ReadApiextensionsV1beta1CustomResourceDefinitionStatusUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
