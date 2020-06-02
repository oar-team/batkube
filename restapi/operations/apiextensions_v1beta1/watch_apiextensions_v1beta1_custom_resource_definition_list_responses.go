// Code generated by go-swagger; DO NOT EDIT.

package apiextensions_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/models"
)

// WatchApiextensionsV1beta1CustomResourceDefinitionListOKCode is the HTTP code returned for type WatchApiextensionsV1beta1CustomResourceDefinitionListOK
const WatchApiextensionsV1beta1CustomResourceDefinitionListOKCode int = 200

/*WatchApiextensionsV1beta1CustomResourceDefinitionListOK OK

swagger:response watchApiextensionsV1beta1CustomResourceDefinitionListOK
*/
type WatchApiextensionsV1beta1CustomResourceDefinitionListOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent `json:"body,omitempty"`
}

// NewWatchApiextensionsV1beta1CustomResourceDefinitionListOK creates WatchApiextensionsV1beta1CustomResourceDefinitionListOK with default headers values
func NewWatchApiextensionsV1beta1CustomResourceDefinitionListOK() *WatchApiextensionsV1beta1CustomResourceDefinitionListOK {

	return &WatchApiextensionsV1beta1CustomResourceDefinitionListOK{}
}

// WithPayload adds the payload to the watch apiextensions v1beta1 custom resource definition list o k response
func (o *WatchApiextensionsV1beta1CustomResourceDefinitionListOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) *WatchApiextensionsV1beta1CustomResourceDefinitionListOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the watch apiextensions v1beta1 custom resource definition list o k response
func (o *WatchApiextensionsV1beta1CustomResourceDefinitionListOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WatchApiextensionsV1beta1CustomResourceDefinitionListOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WatchApiextensionsV1beta1CustomResourceDefinitionListUnauthorizedCode is the HTTP code returned for type WatchApiextensionsV1beta1CustomResourceDefinitionListUnauthorized
const WatchApiextensionsV1beta1CustomResourceDefinitionListUnauthorizedCode int = 401

/*WatchApiextensionsV1beta1CustomResourceDefinitionListUnauthorized Unauthorized

swagger:response watchApiextensionsV1beta1CustomResourceDefinitionListUnauthorized
*/
type WatchApiextensionsV1beta1CustomResourceDefinitionListUnauthorized struct {
}

// NewWatchApiextensionsV1beta1CustomResourceDefinitionListUnauthorized creates WatchApiextensionsV1beta1CustomResourceDefinitionListUnauthorized with default headers values
func NewWatchApiextensionsV1beta1CustomResourceDefinitionListUnauthorized() *WatchApiextensionsV1beta1CustomResourceDefinitionListUnauthorized {

	return &WatchApiextensionsV1beta1CustomResourceDefinitionListUnauthorized{}
}

// WriteResponse to the client
func (o *WatchApiextensionsV1beta1CustomResourceDefinitionListUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
