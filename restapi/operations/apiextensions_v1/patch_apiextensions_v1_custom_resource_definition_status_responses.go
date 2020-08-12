// Code generated by go-swagger; DO NOT EDIT.

package apiextensions_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// PatchApiextensionsV1CustomResourceDefinitionStatusOKCode is the HTTP code returned for type PatchApiextensionsV1CustomResourceDefinitionStatusOK
const PatchApiextensionsV1CustomResourceDefinitionStatusOKCode int = 200

/*PatchApiextensionsV1CustomResourceDefinitionStatusOK OK

swagger:response patchApiextensionsV1CustomResourceDefinitionStatusOK
*/
type PatchApiextensionsV1CustomResourceDefinitionStatusOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinition `json:"body,omitempty"`
}

// NewPatchApiextensionsV1CustomResourceDefinitionStatusOK creates PatchApiextensionsV1CustomResourceDefinitionStatusOK with default headers values
func NewPatchApiextensionsV1CustomResourceDefinitionStatusOK() *PatchApiextensionsV1CustomResourceDefinitionStatusOK {

	return &PatchApiextensionsV1CustomResourceDefinitionStatusOK{}
}

// WithPayload adds the payload to the patch apiextensions v1 custom resource definition status o k response
func (o *PatchApiextensionsV1CustomResourceDefinitionStatusOK) WithPayload(payload *models.IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinition) *PatchApiextensionsV1CustomResourceDefinitionStatusOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch apiextensions v1 custom resource definition status o k response
func (o *PatchApiextensionsV1CustomResourceDefinitionStatusOK) SetPayload(payload *models.IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinition) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchApiextensionsV1CustomResourceDefinitionStatusOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchApiextensionsV1CustomResourceDefinitionStatusUnauthorizedCode is the HTTP code returned for type PatchApiextensionsV1CustomResourceDefinitionStatusUnauthorized
const PatchApiextensionsV1CustomResourceDefinitionStatusUnauthorizedCode int = 401

/*PatchApiextensionsV1CustomResourceDefinitionStatusUnauthorized Unauthorized

swagger:response patchApiextensionsV1CustomResourceDefinitionStatusUnauthorized
*/
type PatchApiextensionsV1CustomResourceDefinitionStatusUnauthorized struct {
}

// NewPatchApiextensionsV1CustomResourceDefinitionStatusUnauthorized creates PatchApiextensionsV1CustomResourceDefinitionStatusUnauthorized with default headers values
func NewPatchApiextensionsV1CustomResourceDefinitionStatusUnauthorized() *PatchApiextensionsV1CustomResourceDefinitionStatusUnauthorized {

	return &PatchApiextensionsV1CustomResourceDefinitionStatusUnauthorized{}
}

// WriteResponse to the client
func (o *PatchApiextensionsV1CustomResourceDefinitionStatusUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
