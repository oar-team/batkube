// Code generated by go-swagger; DO NOT EDIT.

package settings_v1alpha1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// GetSettingsV1alpha1APIResourcesOKCode is the HTTP code returned for type GetSettingsV1alpha1APIResourcesOK
const GetSettingsV1alpha1APIResourcesOKCode int = 200

/*GetSettingsV1alpha1APIResourcesOK OK

swagger:response getSettingsV1alpha1ApiResourcesOK
*/
type GetSettingsV1alpha1APIResourcesOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1APIResourceList `json:"body,omitempty"`
}

// NewGetSettingsV1alpha1APIResourcesOK creates GetSettingsV1alpha1APIResourcesOK with default headers values
func NewGetSettingsV1alpha1APIResourcesOK() *GetSettingsV1alpha1APIResourcesOK {

	return &GetSettingsV1alpha1APIResourcesOK{}
}

// WithPayload adds the payload to the get settings v1alpha1 Api resources o k response
func (o *GetSettingsV1alpha1APIResourcesOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1APIResourceList) *GetSettingsV1alpha1APIResourcesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get settings v1alpha1 Api resources o k response
func (o *GetSettingsV1alpha1APIResourcesOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1APIResourceList) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSettingsV1alpha1APIResourcesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetSettingsV1alpha1APIResourcesUnauthorizedCode is the HTTP code returned for type GetSettingsV1alpha1APIResourcesUnauthorized
const GetSettingsV1alpha1APIResourcesUnauthorizedCode int = 401

/*GetSettingsV1alpha1APIResourcesUnauthorized Unauthorized

swagger:response getSettingsV1alpha1ApiResourcesUnauthorized
*/
type GetSettingsV1alpha1APIResourcesUnauthorized struct {
}

// NewGetSettingsV1alpha1APIResourcesUnauthorized creates GetSettingsV1alpha1APIResourcesUnauthorized with default headers values
func NewGetSettingsV1alpha1APIResourcesUnauthorized() *GetSettingsV1alpha1APIResourcesUnauthorized {

	return &GetSettingsV1alpha1APIResourcesUnauthorized{}
}

// WriteResponse to the client
func (o *GetSettingsV1alpha1APIResourcesUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
