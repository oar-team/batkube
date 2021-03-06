// Code generated by go-swagger; DO NOT EDIT.

package apiregistration_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// ReplaceApiregistrationV1beta1APIServiceOKCode is the HTTP code returned for type ReplaceApiregistrationV1beta1APIServiceOK
const ReplaceApiregistrationV1beta1APIServiceOKCode int = 200

/*ReplaceApiregistrationV1beta1APIServiceOK OK

swagger:response replaceApiregistrationV1beta1ApiServiceOK
*/
type ReplaceApiregistrationV1beta1APIServiceOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sKubeAggregatorPkgApisApiregistrationV1beta1APIService `json:"body,omitempty"`
}

// NewReplaceApiregistrationV1beta1APIServiceOK creates ReplaceApiregistrationV1beta1APIServiceOK with default headers values
func NewReplaceApiregistrationV1beta1APIServiceOK() *ReplaceApiregistrationV1beta1APIServiceOK {

	return &ReplaceApiregistrationV1beta1APIServiceOK{}
}

// WithPayload adds the payload to the replace apiregistration v1beta1 Api service o k response
func (o *ReplaceApiregistrationV1beta1APIServiceOK) WithPayload(payload *models.IoK8sKubeAggregatorPkgApisApiregistrationV1beta1APIService) *ReplaceApiregistrationV1beta1APIServiceOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace apiregistration v1beta1 Api service o k response
func (o *ReplaceApiregistrationV1beta1APIServiceOK) SetPayload(payload *models.IoK8sKubeAggregatorPkgApisApiregistrationV1beta1APIService) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceApiregistrationV1beta1APIServiceOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceApiregistrationV1beta1APIServiceCreatedCode is the HTTP code returned for type ReplaceApiregistrationV1beta1APIServiceCreated
const ReplaceApiregistrationV1beta1APIServiceCreatedCode int = 201

/*ReplaceApiregistrationV1beta1APIServiceCreated Created

swagger:response replaceApiregistrationV1beta1ApiServiceCreated
*/
type ReplaceApiregistrationV1beta1APIServiceCreated struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sKubeAggregatorPkgApisApiregistrationV1beta1APIService `json:"body,omitempty"`
}

// NewReplaceApiregistrationV1beta1APIServiceCreated creates ReplaceApiregistrationV1beta1APIServiceCreated with default headers values
func NewReplaceApiregistrationV1beta1APIServiceCreated() *ReplaceApiregistrationV1beta1APIServiceCreated {

	return &ReplaceApiregistrationV1beta1APIServiceCreated{}
}

// WithPayload adds the payload to the replace apiregistration v1beta1 Api service created response
func (o *ReplaceApiregistrationV1beta1APIServiceCreated) WithPayload(payload *models.IoK8sKubeAggregatorPkgApisApiregistrationV1beta1APIService) *ReplaceApiregistrationV1beta1APIServiceCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace apiregistration v1beta1 Api service created response
func (o *ReplaceApiregistrationV1beta1APIServiceCreated) SetPayload(payload *models.IoK8sKubeAggregatorPkgApisApiregistrationV1beta1APIService) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceApiregistrationV1beta1APIServiceCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceApiregistrationV1beta1APIServiceUnauthorizedCode is the HTTP code returned for type ReplaceApiregistrationV1beta1APIServiceUnauthorized
const ReplaceApiregistrationV1beta1APIServiceUnauthorizedCode int = 401

/*ReplaceApiregistrationV1beta1APIServiceUnauthorized Unauthorized

swagger:response replaceApiregistrationV1beta1ApiServiceUnauthorized
*/
type ReplaceApiregistrationV1beta1APIServiceUnauthorized struct {
}

// NewReplaceApiregistrationV1beta1APIServiceUnauthorized creates ReplaceApiregistrationV1beta1APIServiceUnauthorized with default headers values
func NewReplaceApiregistrationV1beta1APIServiceUnauthorized() *ReplaceApiregistrationV1beta1APIServiceUnauthorized {

	return &ReplaceApiregistrationV1beta1APIServiceUnauthorized{}
}

// WriteResponse to the client
func (o *ReplaceApiregistrationV1beta1APIServiceUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
