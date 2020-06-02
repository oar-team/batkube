// Code generated by go-swagger; DO NOT EDIT.

package networking_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/models"
)

// GetNetworkingV1APIResourcesOKCode is the HTTP code returned for type GetNetworkingV1APIResourcesOK
const GetNetworkingV1APIResourcesOKCode int = 200

/*GetNetworkingV1APIResourcesOK OK

swagger:response getNetworkingV1ApiResourcesOK
*/
type GetNetworkingV1APIResourcesOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1APIResourceList `json:"body,omitempty"`
}

// NewGetNetworkingV1APIResourcesOK creates GetNetworkingV1APIResourcesOK with default headers values
func NewGetNetworkingV1APIResourcesOK() *GetNetworkingV1APIResourcesOK {

	return &GetNetworkingV1APIResourcesOK{}
}

// WithPayload adds the payload to the get networking v1 Api resources o k response
func (o *GetNetworkingV1APIResourcesOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1APIResourceList) *GetNetworkingV1APIResourcesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get networking v1 Api resources o k response
func (o *GetNetworkingV1APIResourcesOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1APIResourceList) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetNetworkingV1APIResourcesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetNetworkingV1APIResourcesUnauthorizedCode is the HTTP code returned for type GetNetworkingV1APIResourcesUnauthorized
const GetNetworkingV1APIResourcesUnauthorizedCode int = 401

/*GetNetworkingV1APIResourcesUnauthorized Unauthorized

swagger:response getNetworkingV1ApiResourcesUnauthorized
*/
type GetNetworkingV1APIResourcesUnauthorized struct {
}

// NewGetNetworkingV1APIResourcesUnauthorized creates GetNetworkingV1APIResourcesUnauthorized with default headers values
func NewGetNetworkingV1APIResourcesUnauthorized() *GetNetworkingV1APIResourcesUnauthorized {

	return &GetNetworkingV1APIResourcesUnauthorized{}
}

// WriteResponse to the client
func (o *GetNetworkingV1APIResourcesUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
