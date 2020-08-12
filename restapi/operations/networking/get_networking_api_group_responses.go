// Code generated by go-swagger; DO NOT EDIT.

package networking

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// GetNetworkingAPIGroupOKCode is the HTTP code returned for type GetNetworkingAPIGroupOK
const GetNetworkingAPIGroupOKCode int = 200

/*GetNetworkingAPIGroupOK OK

swagger:response getNetworkingApiGroupOK
*/
type GetNetworkingAPIGroupOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1APIGroup `json:"body,omitempty"`
}

// NewGetNetworkingAPIGroupOK creates GetNetworkingAPIGroupOK with default headers values
func NewGetNetworkingAPIGroupOK() *GetNetworkingAPIGroupOK {

	return &GetNetworkingAPIGroupOK{}
}

// WithPayload adds the payload to the get networking Api group o k response
func (o *GetNetworkingAPIGroupOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1APIGroup) *GetNetworkingAPIGroupOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get networking Api group o k response
func (o *GetNetworkingAPIGroupOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1APIGroup) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetNetworkingAPIGroupOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetNetworkingAPIGroupUnauthorizedCode is the HTTP code returned for type GetNetworkingAPIGroupUnauthorized
const GetNetworkingAPIGroupUnauthorizedCode int = 401

/*GetNetworkingAPIGroupUnauthorized Unauthorized

swagger:response getNetworkingApiGroupUnauthorized
*/
type GetNetworkingAPIGroupUnauthorized struct {
}

// NewGetNetworkingAPIGroupUnauthorized creates GetNetworkingAPIGroupUnauthorized with default headers values
func NewGetNetworkingAPIGroupUnauthorized() *GetNetworkingAPIGroupUnauthorized {

	return &GetNetworkingAPIGroupUnauthorized{}
}

// WriteResponse to the client
func (o *GetNetworkingAPIGroupUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
