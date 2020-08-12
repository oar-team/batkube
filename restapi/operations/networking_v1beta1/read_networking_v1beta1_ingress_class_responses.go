// Code generated by go-swagger; DO NOT EDIT.

package networking_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// ReadNetworkingV1beta1IngressClassOKCode is the HTTP code returned for type ReadNetworkingV1beta1IngressClassOK
const ReadNetworkingV1beta1IngressClassOKCode int = 200

/*ReadNetworkingV1beta1IngressClassOK OK

swagger:response readNetworkingV1beta1IngressClassOK
*/
type ReadNetworkingV1beta1IngressClassOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPINetworkingV1beta1IngressClass `json:"body,omitempty"`
}

// NewReadNetworkingV1beta1IngressClassOK creates ReadNetworkingV1beta1IngressClassOK with default headers values
func NewReadNetworkingV1beta1IngressClassOK() *ReadNetworkingV1beta1IngressClassOK {

	return &ReadNetworkingV1beta1IngressClassOK{}
}

// WithPayload adds the payload to the read networking v1beta1 ingress class o k response
func (o *ReadNetworkingV1beta1IngressClassOK) WithPayload(payload *models.IoK8sAPINetworkingV1beta1IngressClass) *ReadNetworkingV1beta1IngressClassOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the read networking v1beta1 ingress class o k response
func (o *ReadNetworkingV1beta1IngressClassOK) SetPayload(payload *models.IoK8sAPINetworkingV1beta1IngressClass) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReadNetworkingV1beta1IngressClassOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReadNetworkingV1beta1IngressClassUnauthorizedCode is the HTTP code returned for type ReadNetworkingV1beta1IngressClassUnauthorized
const ReadNetworkingV1beta1IngressClassUnauthorizedCode int = 401

/*ReadNetworkingV1beta1IngressClassUnauthorized Unauthorized

swagger:response readNetworkingV1beta1IngressClassUnauthorized
*/
type ReadNetworkingV1beta1IngressClassUnauthorized struct {
}

// NewReadNetworkingV1beta1IngressClassUnauthorized creates ReadNetworkingV1beta1IngressClassUnauthorized with default headers values
func NewReadNetworkingV1beta1IngressClassUnauthorized() *ReadNetworkingV1beta1IngressClassUnauthorized {

	return &ReadNetworkingV1beta1IngressClassUnauthorized{}
}

// WriteResponse to the client
func (o *ReadNetworkingV1beta1IngressClassUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
