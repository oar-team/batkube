// Code generated by go-swagger; DO NOT EDIT.

package apiextensions_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// ListApiextensionsV1CustomResourceDefinitionOKCode is the HTTP code returned for type ListApiextensionsV1CustomResourceDefinitionOK
const ListApiextensionsV1CustomResourceDefinitionOKCode int = 200

/*ListApiextensionsV1CustomResourceDefinitionOK OK

swagger:response listApiextensionsV1CustomResourceDefinitionOK
*/
type ListApiextensionsV1CustomResourceDefinitionOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionList `json:"body,omitempty"`
}

// NewListApiextensionsV1CustomResourceDefinitionOK creates ListApiextensionsV1CustomResourceDefinitionOK with default headers values
func NewListApiextensionsV1CustomResourceDefinitionOK() *ListApiextensionsV1CustomResourceDefinitionOK {

	return &ListApiextensionsV1CustomResourceDefinitionOK{}
}

// WithPayload adds the payload to the list apiextensions v1 custom resource definition o k response
func (o *ListApiextensionsV1CustomResourceDefinitionOK) WithPayload(payload *models.IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionList) *ListApiextensionsV1CustomResourceDefinitionOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list apiextensions v1 custom resource definition o k response
func (o *ListApiextensionsV1CustomResourceDefinitionOK) SetPayload(payload *models.IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionList) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListApiextensionsV1CustomResourceDefinitionOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListApiextensionsV1CustomResourceDefinitionUnauthorizedCode is the HTTP code returned for type ListApiextensionsV1CustomResourceDefinitionUnauthorized
const ListApiextensionsV1CustomResourceDefinitionUnauthorizedCode int = 401

/*ListApiextensionsV1CustomResourceDefinitionUnauthorized Unauthorized

swagger:response listApiextensionsV1CustomResourceDefinitionUnauthorized
*/
type ListApiextensionsV1CustomResourceDefinitionUnauthorized struct {
}

// NewListApiextensionsV1CustomResourceDefinitionUnauthorized creates ListApiextensionsV1CustomResourceDefinitionUnauthorized with default headers values
func NewListApiextensionsV1CustomResourceDefinitionUnauthorized() *ListApiextensionsV1CustomResourceDefinitionUnauthorized {

	return &ListApiextensionsV1CustomResourceDefinitionUnauthorized{}
}

// WriteResponse to the client
func (o *ListApiextensionsV1CustomResourceDefinitionUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
