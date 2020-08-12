// Code generated by go-swagger; DO NOT EDIT.

package admissionregistration_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// GetAdmissionregistrationV1beta1APIResourcesOKCode is the HTTP code returned for type GetAdmissionregistrationV1beta1APIResourcesOK
const GetAdmissionregistrationV1beta1APIResourcesOKCode int = 200

/*GetAdmissionregistrationV1beta1APIResourcesOK OK

swagger:response getAdmissionregistrationV1beta1ApiResourcesOK
*/
type GetAdmissionregistrationV1beta1APIResourcesOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1APIResourceList `json:"body,omitempty"`
}

// NewGetAdmissionregistrationV1beta1APIResourcesOK creates GetAdmissionregistrationV1beta1APIResourcesOK with default headers values
func NewGetAdmissionregistrationV1beta1APIResourcesOK() *GetAdmissionregistrationV1beta1APIResourcesOK {

	return &GetAdmissionregistrationV1beta1APIResourcesOK{}
}

// WithPayload adds the payload to the get admissionregistration v1beta1 Api resources o k response
func (o *GetAdmissionregistrationV1beta1APIResourcesOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1APIResourceList) *GetAdmissionregistrationV1beta1APIResourcesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get admissionregistration v1beta1 Api resources o k response
func (o *GetAdmissionregistrationV1beta1APIResourcesOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1APIResourceList) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAdmissionregistrationV1beta1APIResourcesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetAdmissionregistrationV1beta1APIResourcesUnauthorizedCode is the HTTP code returned for type GetAdmissionregistrationV1beta1APIResourcesUnauthorized
const GetAdmissionregistrationV1beta1APIResourcesUnauthorizedCode int = 401

/*GetAdmissionregistrationV1beta1APIResourcesUnauthorized Unauthorized

swagger:response getAdmissionregistrationV1beta1ApiResourcesUnauthorized
*/
type GetAdmissionregistrationV1beta1APIResourcesUnauthorized struct {
}

// NewGetAdmissionregistrationV1beta1APIResourcesUnauthorized creates GetAdmissionregistrationV1beta1APIResourcesUnauthorized with default headers values
func NewGetAdmissionregistrationV1beta1APIResourcesUnauthorized() *GetAdmissionregistrationV1beta1APIResourcesUnauthorized {

	return &GetAdmissionregistrationV1beta1APIResourcesUnauthorized{}
}

// WriteResponse to the client
func (o *GetAdmissionregistrationV1beta1APIResourcesUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
