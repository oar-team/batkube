// Code generated by go-swagger; DO NOT EDIT.

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/models"
)

// ListCoreV1NamespacedConfigMapOKCode is the HTTP code returned for type ListCoreV1NamespacedConfigMapOK
const ListCoreV1NamespacedConfigMapOKCode int = 200

/*ListCoreV1NamespacedConfigMapOK OK

swagger:response listCoreV1NamespacedConfigMapOK
*/
type ListCoreV1NamespacedConfigMapOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPICoreV1ConfigMapList `json:"body,omitempty"`
}

// NewListCoreV1NamespacedConfigMapOK creates ListCoreV1NamespacedConfigMapOK with default headers values
func NewListCoreV1NamespacedConfigMapOK() *ListCoreV1NamespacedConfigMapOK {

	return &ListCoreV1NamespacedConfigMapOK{}
}

// WithPayload adds the payload to the list core v1 namespaced config map o k response
func (o *ListCoreV1NamespacedConfigMapOK) WithPayload(payload *models.IoK8sAPICoreV1ConfigMapList) *ListCoreV1NamespacedConfigMapOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list core v1 namespaced config map o k response
func (o *ListCoreV1NamespacedConfigMapOK) SetPayload(payload *models.IoK8sAPICoreV1ConfigMapList) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListCoreV1NamespacedConfigMapOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListCoreV1NamespacedConfigMapUnauthorizedCode is the HTTP code returned for type ListCoreV1NamespacedConfigMapUnauthorized
const ListCoreV1NamespacedConfigMapUnauthorizedCode int = 401

/*ListCoreV1NamespacedConfigMapUnauthorized Unauthorized

swagger:response listCoreV1NamespacedConfigMapUnauthorized
*/
type ListCoreV1NamespacedConfigMapUnauthorized struct {
}

// NewListCoreV1NamespacedConfigMapUnauthorized creates ListCoreV1NamespacedConfigMapUnauthorized with default headers values
func NewListCoreV1NamespacedConfigMapUnauthorized() *ListCoreV1NamespacedConfigMapUnauthorized {

	return &ListCoreV1NamespacedConfigMapUnauthorized{}
}

// WriteResponse to the client
func (o *ListCoreV1NamespacedConfigMapUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}