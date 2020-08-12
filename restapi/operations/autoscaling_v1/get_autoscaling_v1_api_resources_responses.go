// Code generated by go-swagger; DO NOT EDIT.

package autoscaling_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// GetAutoscalingV1APIResourcesOKCode is the HTTP code returned for type GetAutoscalingV1APIResourcesOK
const GetAutoscalingV1APIResourcesOKCode int = 200

/*GetAutoscalingV1APIResourcesOK OK

swagger:response getAutoscalingV1ApiResourcesOK
*/
type GetAutoscalingV1APIResourcesOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1APIResourceList `json:"body,omitempty"`
}

// NewGetAutoscalingV1APIResourcesOK creates GetAutoscalingV1APIResourcesOK with default headers values
func NewGetAutoscalingV1APIResourcesOK() *GetAutoscalingV1APIResourcesOK {

	return &GetAutoscalingV1APIResourcesOK{}
}

// WithPayload adds the payload to the get autoscaling v1 Api resources o k response
func (o *GetAutoscalingV1APIResourcesOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1APIResourceList) *GetAutoscalingV1APIResourcesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get autoscaling v1 Api resources o k response
func (o *GetAutoscalingV1APIResourcesOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1APIResourceList) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAutoscalingV1APIResourcesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetAutoscalingV1APIResourcesUnauthorizedCode is the HTTP code returned for type GetAutoscalingV1APIResourcesUnauthorized
const GetAutoscalingV1APIResourcesUnauthorizedCode int = 401

/*GetAutoscalingV1APIResourcesUnauthorized Unauthorized

swagger:response getAutoscalingV1ApiResourcesUnauthorized
*/
type GetAutoscalingV1APIResourcesUnauthorized struct {
}

// NewGetAutoscalingV1APIResourcesUnauthorized creates GetAutoscalingV1APIResourcesUnauthorized with default headers values
func NewGetAutoscalingV1APIResourcesUnauthorized() *GetAutoscalingV1APIResourcesUnauthorized {

	return &GetAutoscalingV1APIResourcesUnauthorized{}
}

// WriteResponse to the client
func (o *GetAutoscalingV1APIResourcesUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
