// Code generated by go-swagger; DO NOT EDIT.

package discovery

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// GetDiscoveryAPIGroupOKCode is the HTTP code returned for type GetDiscoveryAPIGroupOK
const GetDiscoveryAPIGroupOKCode int = 200

/*GetDiscoveryAPIGroupOK OK

swagger:response getDiscoveryApiGroupOK
*/
type GetDiscoveryAPIGroupOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1APIGroup `json:"body,omitempty"`
}

// NewGetDiscoveryAPIGroupOK creates GetDiscoveryAPIGroupOK with default headers values
func NewGetDiscoveryAPIGroupOK() *GetDiscoveryAPIGroupOK {

	return &GetDiscoveryAPIGroupOK{}
}

// WithPayload adds the payload to the get discovery Api group o k response
func (o *GetDiscoveryAPIGroupOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1APIGroup) *GetDiscoveryAPIGroupOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get discovery Api group o k response
func (o *GetDiscoveryAPIGroupOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1APIGroup) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetDiscoveryAPIGroupOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetDiscoveryAPIGroupUnauthorizedCode is the HTTP code returned for type GetDiscoveryAPIGroupUnauthorized
const GetDiscoveryAPIGroupUnauthorizedCode int = 401

/*GetDiscoveryAPIGroupUnauthorized Unauthorized

swagger:response getDiscoveryApiGroupUnauthorized
*/
type GetDiscoveryAPIGroupUnauthorized struct {
}

// NewGetDiscoveryAPIGroupUnauthorized creates GetDiscoveryAPIGroupUnauthorized with default headers values
func NewGetDiscoveryAPIGroupUnauthorized() *GetDiscoveryAPIGroupUnauthorized {

	return &GetDiscoveryAPIGroupUnauthorized{}
}

// WriteResponse to the client
func (o *GetDiscoveryAPIGroupUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
