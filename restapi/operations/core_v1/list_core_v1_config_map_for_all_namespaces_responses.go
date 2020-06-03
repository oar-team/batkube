// Code generated by go-swagger; DO NOT EDIT.

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/models"
)

// ListCoreV1ConfigMapForAllNamespacesOKCode is the HTTP code returned for type ListCoreV1ConfigMapForAllNamespacesOK
const ListCoreV1ConfigMapForAllNamespacesOKCode int = 200

/*ListCoreV1ConfigMapForAllNamespacesOK OK

swagger:response listCoreV1ConfigMapForAllNamespacesOK
*/
type ListCoreV1ConfigMapForAllNamespacesOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPICoreV1ConfigMapList `json:"body,omitempty"`
}

// NewListCoreV1ConfigMapForAllNamespacesOK creates ListCoreV1ConfigMapForAllNamespacesOK with default headers values
func NewListCoreV1ConfigMapForAllNamespacesOK() *ListCoreV1ConfigMapForAllNamespacesOK {

	return &ListCoreV1ConfigMapForAllNamespacesOK{}
}

// WithPayload adds the payload to the list core v1 config map for all namespaces o k response
func (o *ListCoreV1ConfigMapForAllNamespacesOK) WithPayload(payload *models.IoK8sAPICoreV1ConfigMapList) *ListCoreV1ConfigMapForAllNamespacesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list core v1 config map for all namespaces o k response
func (o *ListCoreV1ConfigMapForAllNamespacesOK) SetPayload(payload *models.IoK8sAPICoreV1ConfigMapList) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListCoreV1ConfigMapForAllNamespacesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListCoreV1ConfigMapForAllNamespacesUnauthorizedCode is the HTTP code returned for type ListCoreV1ConfigMapForAllNamespacesUnauthorized
const ListCoreV1ConfigMapForAllNamespacesUnauthorizedCode int = 401

/*ListCoreV1ConfigMapForAllNamespacesUnauthorized Unauthorized

swagger:response listCoreV1ConfigMapForAllNamespacesUnauthorized
*/
type ListCoreV1ConfigMapForAllNamespacesUnauthorized struct {
}

// NewListCoreV1ConfigMapForAllNamespacesUnauthorized creates ListCoreV1ConfigMapForAllNamespacesUnauthorized with default headers values
func NewListCoreV1ConfigMapForAllNamespacesUnauthorized() *ListCoreV1ConfigMapForAllNamespacesUnauthorized {

	return &ListCoreV1ConfigMapForAllNamespacesUnauthorized{}
}

// WriteResponse to the client
func (o *ListCoreV1ConfigMapForAllNamespacesUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}