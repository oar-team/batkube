// Code generated by go-swagger; DO NOT EDIT.

package networking_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/models"
)

// ReplaceNetworkingV1beta1IngressClassOKCode is the HTTP code returned for type ReplaceNetworkingV1beta1IngressClassOK
const ReplaceNetworkingV1beta1IngressClassOKCode int = 200

/*ReplaceNetworkingV1beta1IngressClassOK OK

swagger:response replaceNetworkingV1beta1IngressClassOK
*/
type ReplaceNetworkingV1beta1IngressClassOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPINetworkingV1beta1IngressClass `json:"body,omitempty"`
}

// NewReplaceNetworkingV1beta1IngressClassOK creates ReplaceNetworkingV1beta1IngressClassOK with default headers values
func NewReplaceNetworkingV1beta1IngressClassOK() *ReplaceNetworkingV1beta1IngressClassOK {

	return &ReplaceNetworkingV1beta1IngressClassOK{}
}

// WithPayload adds the payload to the replace networking v1beta1 ingress class o k response
func (o *ReplaceNetworkingV1beta1IngressClassOK) WithPayload(payload *models.IoK8sAPINetworkingV1beta1IngressClass) *ReplaceNetworkingV1beta1IngressClassOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace networking v1beta1 ingress class o k response
func (o *ReplaceNetworkingV1beta1IngressClassOK) SetPayload(payload *models.IoK8sAPINetworkingV1beta1IngressClass) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceNetworkingV1beta1IngressClassOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceNetworkingV1beta1IngressClassCreatedCode is the HTTP code returned for type ReplaceNetworkingV1beta1IngressClassCreated
const ReplaceNetworkingV1beta1IngressClassCreatedCode int = 201

/*ReplaceNetworkingV1beta1IngressClassCreated Created

swagger:response replaceNetworkingV1beta1IngressClassCreated
*/
type ReplaceNetworkingV1beta1IngressClassCreated struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPINetworkingV1beta1IngressClass `json:"body,omitempty"`
}

// NewReplaceNetworkingV1beta1IngressClassCreated creates ReplaceNetworkingV1beta1IngressClassCreated with default headers values
func NewReplaceNetworkingV1beta1IngressClassCreated() *ReplaceNetworkingV1beta1IngressClassCreated {

	return &ReplaceNetworkingV1beta1IngressClassCreated{}
}

// WithPayload adds the payload to the replace networking v1beta1 ingress class created response
func (o *ReplaceNetworkingV1beta1IngressClassCreated) WithPayload(payload *models.IoK8sAPINetworkingV1beta1IngressClass) *ReplaceNetworkingV1beta1IngressClassCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace networking v1beta1 ingress class created response
func (o *ReplaceNetworkingV1beta1IngressClassCreated) SetPayload(payload *models.IoK8sAPINetworkingV1beta1IngressClass) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceNetworkingV1beta1IngressClassCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceNetworkingV1beta1IngressClassUnauthorizedCode is the HTTP code returned for type ReplaceNetworkingV1beta1IngressClassUnauthorized
const ReplaceNetworkingV1beta1IngressClassUnauthorizedCode int = 401

/*ReplaceNetworkingV1beta1IngressClassUnauthorized Unauthorized

swagger:response replaceNetworkingV1beta1IngressClassUnauthorized
*/
type ReplaceNetworkingV1beta1IngressClassUnauthorized struct {
}

// NewReplaceNetworkingV1beta1IngressClassUnauthorized creates ReplaceNetworkingV1beta1IngressClassUnauthorized with default headers values
func NewReplaceNetworkingV1beta1IngressClassUnauthorized() *ReplaceNetworkingV1beta1IngressClassUnauthorized {

	return &ReplaceNetworkingV1beta1IngressClassUnauthorized{}
}

// WriteResponse to the client
func (o *ReplaceNetworkingV1beta1IngressClassUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
