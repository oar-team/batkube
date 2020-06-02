// Code generated by go-swagger; DO NOT EDIT.

package apiregistration_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/models"
)

// GetApiregistrationV1beta1APIResourcesOKCode is the HTTP code returned for type GetApiregistrationV1beta1APIResourcesOK
const GetApiregistrationV1beta1APIResourcesOKCode int = 200

/*GetApiregistrationV1beta1APIResourcesOK OK

swagger:response getApiregistrationV1beta1ApiResourcesOK
*/
type GetApiregistrationV1beta1APIResourcesOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1APIResourceList `json:"body,omitempty"`
}

// NewGetApiregistrationV1beta1APIResourcesOK creates GetApiregistrationV1beta1APIResourcesOK with default headers values
func NewGetApiregistrationV1beta1APIResourcesOK() *GetApiregistrationV1beta1APIResourcesOK {

	return &GetApiregistrationV1beta1APIResourcesOK{}
}

// WithPayload adds the payload to the get apiregistration v1beta1 Api resources o k response
func (o *GetApiregistrationV1beta1APIResourcesOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1APIResourceList) *GetApiregistrationV1beta1APIResourcesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get apiregistration v1beta1 Api resources o k response
func (o *GetApiregistrationV1beta1APIResourcesOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1APIResourceList) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetApiregistrationV1beta1APIResourcesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetApiregistrationV1beta1APIResourcesUnauthorizedCode is the HTTP code returned for type GetApiregistrationV1beta1APIResourcesUnauthorized
const GetApiregistrationV1beta1APIResourcesUnauthorizedCode int = 401

/*GetApiregistrationV1beta1APIResourcesUnauthorized Unauthorized

swagger:response getApiregistrationV1beta1ApiResourcesUnauthorized
*/
type GetApiregistrationV1beta1APIResourcesUnauthorized struct {
}

// NewGetApiregistrationV1beta1APIResourcesUnauthorized creates GetApiregistrationV1beta1APIResourcesUnauthorized with default headers values
func NewGetApiregistrationV1beta1APIResourcesUnauthorized() *GetApiregistrationV1beta1APIResourcesUnauthorized {

	return &GetApiregistrationV1beta1APIResourcesUnauthorized{}
}

// WriteResponse to the client
func (o *GetApiregistrationV1beta1APIResourcesUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
