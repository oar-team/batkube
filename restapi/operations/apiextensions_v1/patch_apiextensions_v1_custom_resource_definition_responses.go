// Code generated by go-swagger; DO NOT EDIT.

package apiextensions_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// PatchApiextensionsV1CustomResourceDefinitionOKCode is the HTTP code returned for type PatchApiextensionsV1CustomResourceDefinitionOK
const PatchApiextensionsV1CustomResourceDefinitionOKCode int = 200

/*PatchApiextensionsV1CustomResourceDefinitionOK OK

swagger:response patchApiextensionsV1CustomResourceDefinitionOK
*/
type PatchApiextensionsV1CustomResourceDefinitionOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinition `json:"body,omitempty"`
}

// NewPatchApiextensionsV1CustomResourceDefinitionOK creates PatchApiextensionsV1CustomResourceDefinitionOK with default headers values
func NewPatchApiextensionsV1CustomResourceDefinitionOK() *PatchApiextensionsV1CustomResourceDefinitionOK {

	return &PatchApiextensionsV1CustomResourceDefinitionOK{}
}

// WithPayload adds the payload to the patch apiextensions v1 custom resource definition o k response
func (o *PatchApiextensionsV1CustomResourceDefinitionOK) WithPayload(payload *models.IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinition) *PatchApiextensionsV1CustomResourceDefinitionOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch apiextensions v1 custom resource definition o k response
func (o *PatchApiextensionsV1CustomResourceDefinitionOK) SetPayload(payload *models.IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinition) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchApiextensionsV1CustomResourceDefinitionOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchApiextensionsV1CustomResourceDefinitionUnauthorizedCode is the HTTP code returned for type PatchApiextensionsV1CustomResourceDefinitionUnauthorized
const PatchApiextensionsV1CustomResourceDefinitionUnauthorizedCode int = 401

/*PatchApiextensionsV1CustomResourceDefinitionUnauthorized Unauthorized

swagger:response patchApiextensionsV1CustomResourceDefinitionUnauthorized
*/
type PatchApiextensionsV1CustomResourceDefinitionUnauthorized struct {
}

// NewPatchApiextensionsV1CustomResourceDefinitionUnauthorized creates PatchApiextensionsV1CustomResourceDefinitionUnauthorized with default headers values
func NewPatchApiextensionsV1CustomResourceDefinitionUnauthorized() *PatchApiextensionsV1CustomResourceDefinitionUnauthorized {

	return &PatchApiextensionsV1CustomResourceDefinitionUnauthorized{}
}

// WriteResponse to the client
func (o *PatchApiextensionsV1CustomResourceDefinitionUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
