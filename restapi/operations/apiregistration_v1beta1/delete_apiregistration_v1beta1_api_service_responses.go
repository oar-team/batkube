// Code generated by go-swagger; DO NOT EDIT.

package apiregistration_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/models"
)

// DeleteApiregistrationV1beta1APIServiceOKCode is the HTTP code returned for type DeleteApiregistrationV1beta1APIServiceOK
const DeleteApiregistrationV1beta1APIServiceOKCode int = 200

/*DeleteApiregistrationV1beta1APIServiceOK OK

swagger:response deleteApiregistrationV1beta1ApiServiceOK
*/
type DeleteApiregistrationV1beta1APIServiceOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1Status `json:"body,omitempty"`
}

// NewDeleteApiregistrationV1beta1APIServiceOK creates DeleteApiregistrationV1beta1APIServiceOK with default headers values
func NewDeleteApiregistrationV1beta1APIServiceOK() *DeleteApiregistrationV1beta1APIServiceOK {

	return &DeleteApiregistrationV1beta1APIServiceOK{}
}

// WithPayload adds the payload to the delete apiregistration v1beta1 Api service o k response
func (o *DeleteApiregistrationV1beta1APIServiceOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) *DeleteApiregistrationV1beta1APIServiceOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete apiregistration v1beta1 Api service o k response
func (o *DeleteApiregistrationV1beta1APIServiceOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteApiregistrationV1beta1APIServiceOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteApiregistrationV1beta1APIServiceAcceptedCode is the HTTP code returned for type DeleteApiregistrationV1beta1APIServiceAccepted
const DeleteApiregistrationV1beta1APIServiceAcceptedCode int = 202

/*DeleteApiregistrationV1beta1APIServiceAccepted Accepted

swagger:response deleteApiregistrationV1beta1ApiServiceAccepted
*/
type DeleteApiregistrationV1beta1APIServiceAccepted struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1Status `json:"body,omitempty"`
}

// NewDeleteApiregistrationV1beta1APIServiceAccepted creates DeleteApiregistrationV1beta1APIServiceAccepted with default headers values
func NewDeleteApiregistrationV1beta1APIServiceAccepted() *DeleteApiregistrationV1beta1APIServiceAccepted {

	return &DeleteApiregistrationV1beta1APIServiceAccepted{}
}

// WithPayload adds the payload to the delete apiregistration v1beta1 Api service accepted response
func (o *DeleteApiregistrationV1beta1APIServiceAccepted) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) *DeleteApiregistrationV1beta1APIServiceAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete apiregistration v1beta1 Api service accepted response
func (o *DeleteApiregistrationV1beta1APIServiceAccepted) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteApiregistrationV1beta1APIServiceAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(202)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteApiregistrationV1beta1APIServiceUnauthorizedCode is the HTTP code returned for type DeleteApiregistrationV1beta1APIServiceUnauthorized
const DeleteApiregistrationV1beta1APIServiceUnauthorizedCode int = 401

/*DeleteApiregistrationV1beta1APIServiceUnauthorized Unauthorized

swagger:response deleteApiregistrationV1beta1ApiServiceUnauthorized
*/
type DeleteApiregistrationV1beta1APIServiceUnauthorized struct {
}

// NewDeleteApiregistrationV1beta1APIServiceUnauthorized creates DeleteApiregistrationV1beta1APIServiceUnauthorized with default headers values
func NewDeleteApiregistrationV1beta1APIServiceUnauthorized() *DeleteApiregistrationV1beta1APIServiceUnauthorized {

	return &DeleteApiregistrationV1beta1APIServiceUnauthorized{}
}

// WriteResponse to the client
func (o *DeleteApiregistrationV1beta1APIServiceUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
