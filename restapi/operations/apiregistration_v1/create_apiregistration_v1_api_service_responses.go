// Code generated by go-swagger; DO NOT EDIT.

package apiregistration_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/models"
)

// CreateApiregistrationV1APIServiceOKCode is the HTTP code returned for type CreateApiregistrationV1APIServiceOK
const CreateApiregistrationV1APIServiceOKCode int = 200

/*CreateApiregistrationV1APIServiceOK OK

swagger:response createApiregistrationV1ApiServiceOK
*/
type CreateApiregistrationV1APIServiceOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sKubeAggregatorPkgApisApiregistrationV1APIService `json:"body,omitempty"`
}

// NewCreateApiregistrationV1APIServiceOK creates CreateApiregistrationV1APIServiceOK with default headers values
func NewCreateApiregistrationV1APIServiceOK() *CreateApiregistrationV1APIServiceOK {

	return &CreateApiregistrationV1APIServiceOK{}
}

// WithPayload adds the payload to the create apiregistration v1 Api service o k response
func (o *CreateApiregistrationV1APIServiceOK) WithPayload(payload *models.IoK8sKubeAggregatorPkgApisApiregistrationV1APIService) *CreateApiregistrationV1APIServiceOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create apiregistration v1 Api service o k response
func (o *CreateApiregistrationV1APIServiceOK) SetPayload(payload *models.IoK8sKubeAggregatorPkgApisApiregistrationV1APIService) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateApiregistrationV1APIServiceOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateApiregistrationV1APIServiceCreatedCode is the HTTP code returned for type CreateApiregistrationV1APIServiceCreated
const CreateApiregistrationV1APIServiceCreatedCode int = 201

/*CreateApiregistrationV1APIServiceCreated Created

swagger:response createApiregistrationV1ApiServiceCreated
*/
type CreateApiregistrationV1APIServiceCreated struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sKubeAggregatorPkgApisApiregistrationV1APIService `json:"body,omitempty"`
}

// NewCreateApiregistrationV1APIServiceCreated creates CreateApiregistrationV1APIServiceCreated with default headers values
func NewCreateApiregistrationV1APIServiceCreated() *CreateApiregistrationV1APIServiceCreated {

	return &CreateApiregistrationV1APIServiceCreated{}
}

// WithPayload adds the payload to the create apiregistration v1 Api service created response
func (o *CreateApiregistrationV1APIServiceCreated) WithPayload(payload *models.IoK8sKubeAggregatorPkgApisApiregistrationV1APIService) *CreateApiregistrationV1APIServiceCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create apiregistration v1 Api service created response
func (o *CreateApiregistrationV1APIServiceCreated) SetPayload(payload *models.IoK8sKubeAggregatorPkgApisApiregistrationV1APIService) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateApiregistrationV1APIServiceCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateApiregistrationV1APIServiceAcceptedCode is the HTTP code returned for type CreateApiregistrationV1APIServiceAccepted
const CreateApiregistrationV1APIServiceAcceptedCode int = 202

/*CreateApiregistrationV1APIServiceAccepted Accepted

swagger:response createApiregistrationV1ApiServiceAccepted
*/
type CreateApiregistrationV1APIServiceAccepted struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sKubeAggregatorPkgApisApiregistrationV1APIService `json:"body,omitempty"`
}

// NewCreateApiregistrationV1APIServiceAccepted creates CreateApiregistrationV1APIServiceAccepted with default headers values
func NewCreateApiregistrationV1APIServiceAccepted() *CreateApiregistrationV1APIServiceAccepted {

	return &CreateApiregistrationV1APIServiceAccepted{}
}

// WithPayload adds the payload to the create apiregistration v1 Api service accepted response
func (o *CreateApiregistrationV1APIServiceAccepted) WithPayload(payload *models.IoK8sKubeAggregatorPkgApisApiregistrationV1APIService) *CreateApiregistrationV1APIServiceAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create apiregistration v1 Api service accepted response
func (o *CreateApiregistrationV1APIServiceAccepted) SetPayload(payload *models.IoK8sKubeAggregatorPkgApisApiregistrationV1APIService) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateApiregistrationV1APIServiceAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(202)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateApiregistrationV1APIServiceUnauthorizedCode is the HTTP code returned for type CreateApiregistrationV1APIServiceUnauthorized
const CreateApiregistrationV1APIServiceUnauthorizedCode int = 401

/*CreateApiregistrationV1APIServiceUnauthorized Unauthorized

swagger:response createApiregistrationV1ApiServiceUnauthorized
*/
type CreateApiregistrationV1APIServiceUnauthorized struct {
}

// NewCreateApiregistrationV1APIServiceUnauthorized creates CreateApiregistrationV1APIServiceUnauthorized with default headers values
func NewCreateApiregistrationV1APIServiceUnauthorized() *CreateApiregistrationV1APIServiceUnauthorized {

	return &CreateApiregistrationV1APIServiceUnauthorized{}
}

// WriteResponse to the client
func (o *CreateApiregistrationV1APIServiceUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
