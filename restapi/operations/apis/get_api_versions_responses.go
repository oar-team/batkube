// Code generated by go-swagger; DO NOT EDIT.

package apis

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// GetAPIVersionsOKCode is the HTTP code returned for type GetAPIVersionsOK
const GetAPIVersionsOKCode int = 200

/*GetAPIVersionsOK OK

swagger:response getApiVersionsOK
*/
type GetAPIVersionsOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1APIGroupList `json:"body,omitempty"`
}

// NewGetAPIVersionsOK creates GetAPIVersionsOK with default headers values
func NewGetAPIVersionsOK() *GetAPIVersionsOK {

	return &GetAPIVersionsOK{}
}

// WithPayload adds the payload to the get Api versions o k response
func (o *GetAPIVersionsOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1APIGroupList) *GetAPIVersionsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get Api versions o k response
func (o *GetAPIVersionsOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1APIGroupList) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAPIVersionsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetAPIVersionsUnauthorizedCode is the HTTP code returned for type GetAPIVersionsUnauthorized
const GetAPIVersionsUnauthorizedCode int = 401

/*GetAPIVersionsUnauthorized Unauthorized

swagger:response getApiVersionsUnauthorized
*/
type GetAPIVersionsUnauthorized struct {
}

// NewGetAPIVersionsUnauthorized creates GetAPIVersionsUnauthorized with default headers values
func NewGetAPIVersionsUnauthorized() *GetAPIVersionsUnauthorized {

	return &GetAPIVersionsUnauthorized{}
}

// WriteResponse to the client
func (o *GetAPIVersionsUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
