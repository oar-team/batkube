// Code generated by go-swagger; DO NOT EDIT.

package events_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// GetEventsV1beta1APIResourcesOKCode is the HTTP code returned for type GetEventsV1beta1APIResourcesOK
const GetEventsV1beta1APIResourcesOKCode int = 200

/*GetEventsV1beta1APIResourcesOK OK

swagger:response getEventsV1beta1ApiResourcesOK
*/
type GetEventsV1beta1APIResourcesOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1APIResourceList `json:"body,omitempty"`
}

// NewGetEventsV1beta1APIResourcesOK creates GetEventsV1beta1APIResourcesOK with default headers values
func NewGetEventsV1beta1APIResourcesOK() *GetEventsV1beta1APIResourcesOK {

	return &GetEventsV1beta1APIResourcesOK{}
}

// WithPayload adds the payload to the get events v1beta1 Api resources o k response
func (o *GetEventsV1beta1APIResourcesOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1APIResourceList) *GetEventsV1beta1APIResourcesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get events v1beta1 Api resources o k response
func (o *GetEventsV1beta1APIResourcesOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1APIResourceList) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetEventsV1beta1APIResourcesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetEventsV1beta1APIResourcesUnauthorizedCode is the HTTP code returned for type GetEventsV1beta1APIResourcesUnauthorized
const GetEventsV1beta1APIResourcesUnauthorizedCode int = 401

/*GetEventsV1beta1APIResourcesUnauthorized Unauthorized

swagger:response getEventsV1beta1ApiResourcesUnauthorized
*/
type GetEventsV1beta1APIResourcesUnauthorized struct {
}

// NewGetEventsV1beta1APIResourcesUnauthorized creates GetEventsV1beta1APIResourcesUnauthorized with default headers values
func NewGetEventsV1beta1APIResourcesUnauthorized() *GetEventsV1beta1APIResourcesUnauthorized {

	return &GetEventsV1beta1APIResourcesUnauthorized{}
}

// WriteResponse to the client
func (o *GetEventsV1beta1APIResourcesUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
