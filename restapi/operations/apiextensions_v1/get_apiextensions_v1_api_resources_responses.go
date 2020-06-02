// Code generated by go-swagger; DO NOT EDIT.

package apiextensions_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/models"
)

// GetApiextensionsV1APIResourcesOKCode is the HTTP code returned for type GetApiextensionsV1APIResourcesOK
const GetApiextensionsV1APIResourcesOKCode int = 200

/*GetApiextensionsV1APIResourcesOK OK

swagger:response getApiextensionsV1ApiResourcesOK
*/
type GetApiextensionsV1APIResourcesOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1APIResourceList `json:"body,omitempty"`
}

// NewGetApiextensionsV1APIResourcesOK creates GetApiextensionsV1APIResourcesOK with default headers values
func NewGetApiextensionsV1APIResourcesOK() *GetApiextensionsV1APIResourcesOK {

	return &GetApiextensionsV1APIResourcesOK{}
}

// WithPayload adds the payload to the get apiextensions v1 Api resources o k response
func (o *GetApiextensionsV1APIResourcesOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1APIResourceList) *GetApiextensionsV1APIResourcesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get apiextensions v1 Api resources o k response
func (o *GetApiextensionsV1APIResourcesOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1APIResourceList) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetApiextensionsV1APIResourcesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetApiextensionsV1APIResourcesUnauthorizedCode is the HTTP code returned for type GetApiextensionsV1APIResourcesUnauthorized
const GetApiextensionsV1APIResourcesUnauthorizedCode int = 401

/*GetApiextensionsV1APIResourcesUnauthorized Unauthorized

swagger:response getApiextensionsV1ApiResourcesUnauthorized
*/
type GetApiextensionsV1APIResourcesUnauthorized struct {
}

// NewGetApiextensionsV1APIResourcesUnauthorized creates GetApiextensionsV1APIResourcesUnauthorized with default headers values
func NewGetApiextensionsV1APIResourcesUnauthorized() *GetApiextensionsV1APIResourcesUnauthorized {

	return &GetApiextensionsV1APIResourcesUnauthorized{}
}

// WriteResponse to the client
func (o *GetApiextensionsV1APIResourcesUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
