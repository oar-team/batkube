// Code generated by go-swagger; DO NOT EDIT.

package authentication_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// GetAuthenticationV1APIResourcesOKCode is the HTTP code returned for type GetAuthenticationV1APIResourcesOK
const GetAuthenticationV1APIResourcesOKCode int = 200

/*GetAuthenticationV1APIResourcesOK OK

swagger:response getAuthenticationV1ApiResourcesOK
*/
type GetAuthenticationV1APIResourcesOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1APIResourceList `json:"body,omitempty"`
}

// NewGetAuthenticationV1APIResourcesOK creates GetAuthenticationV1APIResourcesOK with default headers values
func NewGetAuthenticationV1APIResourcesOK() *GetAuthenticationV1APIResourcesOK {

	return &GetAuthenticationV1APIResourcesOK{}
}

// WithPayload adds the payload to the get authentication v1 Api resources o k response
func (o *GetAuthenticationV1APIResourcesOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1APIResourceList) *GetAuthenticationV1APIResourcesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get authentication v1 Api resources o k response
func (o *GetAuthenticationV1APIResourcesOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1APIResourceList) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAuthenticationV1APIResourcesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetAuthenticationV1APIResourcesUnauthorizedCode is the HTTP code returned for type GetAuthenticationV1APIResourcesUnauthorized
const GetAuthenticationV1APIResourcesUnauthorizedCode int = 401

/*GetAuthenticationV1APIResourcesUnauthorized Unauthorized

swagger:response getAuthenticationV1ApiResourcesUnauthorized
*/
type GetAuthenticationV1APIResourcesUnauthorized struct {
}

// NewGetAuthenticationV1APIResourcesUnauthorized creates GetAuthenticationV1APIResourcesUnauthorized with default headers values
func NewGetAuthenticationV1APIResourcesUnauthorized() *GetAuthenticationV1APIResourcesUnauthorized {

	return &GetAuthenticationV1APIResourcesUnauthorized{}
}

// WriteResponse to the client
func (o *GetAuthenticationV1APIResourcesUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
