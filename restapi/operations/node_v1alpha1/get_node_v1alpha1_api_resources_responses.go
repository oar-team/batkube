// Code generated by go-swagger; DO NOT EDIT.

package node_v1alpha1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// GetNodeV1alpha1APIResourcesOKCode is the HTTP code returned for type GetNodeV1alpha1APIResourcesOK
const GetNodeV1alpha1APIResourcesOKCode int = 200

/*GetNodeV1alpha1APIResourcesOK OK

swagger:response getNodeV1alpha1ApiResourcesOK
*/
type GetNodeV1alpha1APIResourcesOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1APIResourceList `json:"body,omitempty"`
}

// NewGetNodeV1alpha1APIResourcesOK creates GetNodeV1alpha1APIResourcesOK with default headers values
func NewGetNodeV1alpha1APIResourcesOK() *GetNodeV1alpha1APIResourcesOK {

	return &GetNodeV1alpha1APIResourcesOK{}
}

// WithPayload adds the payload to the get node v1alpha1 Api resources o k response
func (o *GetNodeV1alpha1APIResourcesOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1APIResourceList) *GetNodeV1alpha1APIResourcesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get node v1alpha1 Api resources o k response
func (o *GetNodeV1alpha1APIResourcesOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1APIResourceList) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetNodeV1alpha1APIResourcesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetNodeV1alpha1APIResourcesUnauthorizedCode is the HTTP code returned for type GetNodeV1alpha1APIResourcesUnauthorized
const GetNodeV1alpha1APIResourcesUnauthorizedCode int = 401

/*GetNodeV1alpha1APIResourcesUnauthorized Unauthorized

swagger:response getNodeV1alpha1ApiResourcesUnauthorized
*/
type GetNodeV1alpha1APIResourcesUnauthorized struct {
}

// NewGetNodeV1alpha1APIResourcesUnauthorized creates GetNodeV1alpha1APIResourcesUnauthorized with default headers values
func NewGetNodeV1alpha1APIResourcesUnauthorized() *GetNodeV1alpha1APIResourcesUnauthorized {

	return &GetNodeV1alpha1APIResourcesUnauthorized{}
}

// WriteResponse to the client
func (o *GetNodeV1alpha1APIResourcesUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
