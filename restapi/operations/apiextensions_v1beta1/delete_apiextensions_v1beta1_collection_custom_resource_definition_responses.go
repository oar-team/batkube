// Code generated by go-swagger; DO NOT EDIT.

package apiextensions_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// DeleteApiextensionsV1beta1CollectionCustomResourceDefinitionOKCode is the HTTP code returned for type DeleteApiextensionsV1beta1CollectionCustomResourceDefinitionOK
const DeleteApiextensionsV1beta1CollectionCustomResourceDefinitionOKCode int = 200

/*DeleteApiextensionsV1beta1CollectionCustomResourceDefinitionOK OK

swagger:response deleteApiextensionsV1beta1CollectionCustomResourceDefinitionOK
*/
type DeleteApiextensionsV1beta1CollectionCustomResourceDefinitionOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1Status `json:"body,omitempty"`
}

// NewDeleteApiextensionsV1beta1CollectionCustomResourceDefinitionOK creates DeleteApiextensionsV1beta1CollectionCustomResourceDefinitionOK with default headers values
func NewDeleteApiextensionsV1beta1CollectionCustomResourceDefinitionOK() *DeleteApiextensionsV1beta1CollectionCustomResourceDefinitionOK {

	return &DeleteApiextensionsV1beta1CollectionCustomResourceDefinitionOK{}
}

// WithPayload adds the payload to the delete apiextensions v1beta1 collection custom resource definition o k response
func (o *DeleteApiextensionsV1beta1CollectionCustomResourceDefinitionOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) *DeleteApiextensionsV1beta1CollectionCustomResourceDefinitionOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete apiextensions v1beta1 collection custom resource definition o k response
func (o *DeleteApiextensionsV1beta1CollectionCustomResourceDefinitionOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteApiextensionsV1beta1CollectionCustomResourceDefinitionOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteApiextensionsV1beta1CollectionCustomResourceDefinitionUnauthorizedCode is the HTTP code returned for type DeleteApiextensionsV1beta1CollectionCustomResourceDefinitionUnauthorized
const DeleteApiextensionsV1beta1CollectionCustomResourceDefinitionUnauthorizedCode int = 401

/*DeleteApiextensionsV1beta1CollectionCustomResourceDefinitionUnauthorized Unauthorized

swagger:response deleteApiextensionsV1beta1CollectionCustomResourceDefinitionUnauthorized
*/
type DeleteApiextensionsV1beta1CollectionCustomResourceDefinitionUnauthorized struct {
}

// NewDeleteApiextensionsV1beta1CollectionCustomResourceDefinitionUnauthorized creates DeleteApiextensionsV1beta1CollectionCustomResourceDefinitionUnauthorized with default headers values
func NewDeleteApiextensionsV1beta1CollectionCustomResourceDefinitionUnauthorized() *DeleteApiextensionsV1beta1CollectionCustomResourceDefinitionUnauthorized {

	return &DeleteApiextensionsV1beta1CollectionCustomResourceDefinitionUnauthorized{}
}

// WriteResponse to the client
func (o *DeleteApiextensionsV1beta1CollectionCustomResourceDefinitionUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
