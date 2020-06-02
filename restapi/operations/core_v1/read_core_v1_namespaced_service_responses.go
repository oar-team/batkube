// Code generated by go-swagger; DO NOT EDIT.

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/models"
)

// ReadCoreV1NamespacedServiceOKCode is the HTTP code returned for type ReadCoreV1NamespacedServiceOK
const ReadCoreV1NamespacedServiceOKCode int = 200

/*ReadCoreV1NamespacedServiceOK OK

swagger:response readCoreV1NamespacedServiceOK
*/
type ReadCoreV1NamespacedServiceOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPICoreV1Service `json:"body,omitempty"`
}

// NewReadCoreV1NamespacedServiceOK creates ReadCoreV1NamespacedServiceOK with default headers values
func NewReadCoreV1NamespacedServiceOK() *ReadCoreV1NamespacedServiceOK {

	return &ReadCoreV1NamespacedServiceOK{}
}

// WithPayload adds the payload to the read core v1 namespaced service o k response
func (o *ReadCoreV1NamespacedServiceOK) WithPayload(payload *models.IoK8sAPICoreV1Service) *ReadCoreV1NamespacedServiceOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the read core v1 namespaced service o k response
func (o *ReadCoreV1NamespacedServiceOK) SetPayload(payload *models.IoK8sAPICoreV1Service) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReadCoreV1NamespacedServiceOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReadCoreV1NamespacedServiceUnauthorizedCode is the HTTP code returned for type ReadCoreV1NamespacedServiceUnauthorized
const ReadCoreV1NamespacedServiceUnauthorizedCode int = 401

/*ReadCoreV1NamespacedServiceUnauthorized Unauthorized

swagger:response readCoreV1NamespacedServiceUnauthorized
*/
type ReadCoreV1NamespacedServiceUnauthorized struct {
}

// NewReadCoreV1NamespacedServiceUnauthorized creates ReadCoreV1NamespacedServiceUnauthorized with default headers values
func NewReadCoreV1NamespacedServiceUnauthorized() *ReadCoreV1NamespacedServiceUnauthorized {

	return &ReadCoreV1NamespacedServiceUnauthorized{}
}

// WriteResponse to the client
func (o *ReadCoreV1NamespacedServiceUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
