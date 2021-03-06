// Code generated by go-swagger; DO NOT EDIT.

package apiregistration_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// PatchApiregistrationV1APIServiceStatusOKCode is the HTTP code returned for type PatchApiregistrationV1APIServiceStatusOK
const PatchApiregistrationV1APIServiceStatusOKCode int = 200

/*PatchApiregistrationV1APIServiceStatusOK OK

swagger:response patchApiregistrationV1ApiServiceStatusOK
*/
type PatchApiregistrationV1APIServiceStatusOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sKubeAggregatorPkgApisApiregistrationV1APIService `json:"body,omitempty"`
}

// NewPatchApiregistrationV1APIServiceStatusOK creates PatchApiregistrationV1APIServiceStatusOK with default headers values
func NewPatchApiregistrationV1APIServiceStatusOK() *PatchApiregistrationV1APIServiceStatusOK {

	return &PatchApiregistrationV1APIServiceStatusOK{}
}

// WithPayload adds the payload to the patch apiregistration v1 Api service status o k response
func (o *PatchApiregistrationV1APIServiceStatusOK) WithPayload(payload *models.IoK8sKubeAggregatorPkgApisApiregistrationV1APIService) *PatchApiregistrationV1APIServiceStatusOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch apiregistration v1 Api service status o k response
func (o *PatchApiregistrationV1APIServiceStatusOK) SetPayload(payload *models.IoK8sKubeAggregatorPkgApisApiregistrationV1APIService) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchApiregistrationV1APIServiceStatusOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchApiregistrationV1APIServiceStatusUnauthorizedCode is the HTTP code returned for type PatchApiregistrationV1APIServiceStatusUnauthorized
const PatchApiregistrationV1APIServiceStatusUnauthorizedCode int = 401

/*PatchApiregistrationV1APIServiceStatusUnauthorized Unauthorized

swagger:response patchApiregistrationV1ApiServiceStatusUnauthorized
*/
type PatchApiregistrationV1APIServiceStatusUnauthorized struct {
}

// NewPatchApiregistrationV1APIServiceStatusUnauthorized creates PatchApiregistrationV1APIServiceStatusUnauthorized with default headers values
func NewPatchApiregistrationV1APIServiceStatusUnauthorized() *PatchApiregistrationV1APIServiceStatusUnauthorized {

	return &PatchApiregistrationV1APIServiceStatusUnauthorized{}
}

// WriteResponse to the client
func (o *PatchApiregistrationV1APIServiceStatusUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
