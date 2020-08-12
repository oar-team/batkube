// Code generated by go-swagger; DO NOT EDIT.

package core

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// GetCoreAPIVersionsOKCode is the HTTP code returned for type GetCoreAPIVersionsOK
const GetCoreAPIVersionsOKCode int = 200

/*GetCoreAPIVersionsOK OK

swagger:response getCoreApiVersionsOK
*/
type GetCoreAPIVersionsOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1APIVersions `json:"body,omitempty"`
}

// NewGetCoreAPIVersionsOK creates GetCoreAPIVersionsOK with default headers values
func NewGetCoreAPIVersionsOK() *GetCoreAPIVersionsOK {

	return &GetCoreAPIVersionsOK{}
}

// WithPayload adds the payload to the get core Api versions o k response
func (o *GetCoreAPIVersionsOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1APIVersions) *GetCoreAPIVersionsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get core Api versions o k response
func (o *GetCoreAPIVersionsOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1APIVersions) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCoreAPIVersionsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetCoreAPIVersionsUnauthorizedCode is the HTTP code returned for type GetCoreAPIVersionsUnauthorized
const GetCoreAPIVersionsUnauthorizedCode int = 401

/*GetCoreAPIVersionsUnauthorized Unauthorized

swagger:response getCoreApiVersionsUnauthorized
*/
type GetCoreAPIVersionsUnauthorized struct {
}

// NewGetCoreAPIVersionsUnauthorized creates GetCoreAPIVersionsUnauthorized with default headers values
func NewGetCoreAPIVersionsUnauthorized() *GetCoreAPIVersionsUnauthorized {

	return &GetCoreAPIVersionsUnauthorized{}
}

// WriteResponse to the client
func (o *GetCoreAPIVersionsUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
