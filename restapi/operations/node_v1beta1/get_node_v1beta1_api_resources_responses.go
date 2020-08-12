// Code generated by go-swagger; DO NOT EDIT.

package node_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// GetNodeV1beta1APIResourcesOKCode is the HTTP code returned for type GetNodeV1beta1APIResourcesOK
const GetNodeV1beta1APIResourcesOKCode int = 200

/*GetNodeV1beta1APIResourcesOK OK

swagger:response getNodeV1beta1ApiResourcesOK
*/
type GetNodeV1beta1APIResourcesOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1APIResourceList `json:"body,omitempty"`
}

// NewGetNodeV1beta1APIResourcesOK creates GetNodeV1beta1APIResourcesOK with default headers values
func NewGetNodeV1beta1APIResourcesOK() *GetNodeV1beta1APIResourcesOK {

	return &GetNodeV1beta1APIResourcesOK{}
}

// WithPayload adds the payload to the get node v1beta1 Api resources o k response
func (o *GetNodeV1beta1APIResourcesOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1APIResourceList) *GetNodeV1beta1APIResourcesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get node v1beta1 Api resources o k response
func (o *GetNodeV1beta1APIResourcesOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1APIResourceList) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetNodeV1beta1APIResourcesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetNodeV1beta1APIResourcesUnauthorizedCode is the HTTP code returned for type GetNodeV1beta1APIResourcesUnauthorized
const GetNodeV1beta1APIResourcesUnauthorizedCode int = 401

/*GetNodeV1beta1APIResourcesUnauthorized Unauthorized

swagger:response getNodeV1beta1ApiResourcesUnauthorized
*/
type GetNodeV1beta1APIResourcesUnauthorized struct {
}

// NewGetNodeV1beta1APIResourcesUnauthorized creates GetNodeV1beta1APIResourcesUnauthorized with default headers values
func NewGetNodeV1beta1APIResourcesUnauthorized() *GetNodeV1beta1APIResourcesUnauthorized {

	return &GetNodeV1beta1APIResourcesUnauthorized{}
}

// WriteResponse to the client
func (o *GetNodeV1beta1APIResourcesUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
