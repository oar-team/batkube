// Code generated by go-swagger; DO NOT EDIT.

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// ReadCoreV1NamespacedConfigMapOKCode is the HTTP code returned for type ReadCoreV1NamespacedConfigMapOK
const ReadCoreV1NamespacedConfigMapOKCode int = 200

/*ReadCoreV1NamespacedConfigMapOK OK

swagger:response readCoreV1NamespacedConfigMapOK
*/
type ReadCoreV1NamespacedConfigMapOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPICoreV1ConfigMap `json:"body,omitempty"`
}

// NewReadCoreV1NamespacedConfigMapOK creates ReadCoreV1NamespacedConfigMapOK with default headers values
func NewReadCoreV1NamespacedConfigMapOK() *ReadCoreV1NamespacedConfigMapOK {

	return &ReadCoreV1NamespacedConfigMapOK{}
}

// WithPayload adds the payload to the read core v1 namespaced config map o k response
func (o *ReadCoreV1NamespacedConfigMapOK) WithPayload(payload *models.IoK8sAPICoreV1ConfigMap) *ReadCoreV1NamespacedConfigMapOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the read core v1 namespaced config map o k response
func (o *ReadCoreV1NamespacedConfigMapOK) SetPayload(payload *models.IoK8sAPICoreV1ConfigMap) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReadCoreV1NamespacedConfigMapOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReadCoreV1NamespacedConfigMapUnauthorizedCode is the HTTP code returned for type ReadCoreV1NamespacedConfigMapUnauthorized
const ReadCoreV1NamespacedConfigMapUnauthorizedCode int = 401

/*ReadCoreV1NamespacedConfigMapUnauthorized Unauthorized

swagger:response readCoreV1NamespacedConfigMapUnauthorized
*/
type ReadCoreV1NamespacedConfigMapUnauthorized struct {
}

// NewReadCoreV1NamespacedConfigMapUnauthorized creates ReadCoreV1NamespacedConfigMapUnauthorized with default headers values
func NewReadCoreV1NamespacedConfigMapUnauthorized() *ReadCoreV1NamespacedConfigMapUnauthorized {

	return &ReadCoreV1NamespacedConfigMapUnauthorized{}
}

// WriteResponse to the client
func (o *ReadCoreV1NamespacedConfigMapUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
