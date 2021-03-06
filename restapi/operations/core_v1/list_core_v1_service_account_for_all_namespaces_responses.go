// Code generated by go-swagger; DO NOT EDIT.

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// ListCoreV1ServiceAccountForAllNamespacesOKCode is the HTTP code returned for type ListCoreV1ServiceAccountForAllNamespacesOK
const ListCoreV1ServiceAccountForAllNamespacesOKCode int = 200

/*ListCoreV1ServiceAccountForAllNamespacesOK OK

swagger:response listCoreV1ServiceAccountForAllNamespacesOK
*/
type ListCoreV1ServiceAccountForAllNamespacesOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPICoreV1ServiceAccountList `json:"body,omitempty"`
}

// NewListCoreV1ServiceAccountForAllNamespacesOK creates ListCoreV1ServiceAccountForAllNamespacesOK with default headers values
func NewListCoreV1ServiceAccountForAllNamespacesOK() *ListCoreV1ServiceAccountForAllNamespacesOK {

	return &ListCoreV1ServiceAccountForAllNamespacesOK{}
}

// WithPayload adds the payload to the list core v1 service account for all namespaces o k response
func (o *ListCoreV1ServiceAccountForAllNamespacesOK) WithPayload(payload *models.IoK8sAPICoreV1ServiceAccountList) *ListCoreV1ServiceAccountForAllNamespacesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list core v1 service account for all namespaces o k response
func (o *ListCoreV1ServiceAccountForAllNamespacesOK) SetPayload(payload *models.IoK8sAPICoreV1ServiceAccountList) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListCoreV1ServiceAccountForAllNamespacesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListCoreV1ServiceAccountForAllNamespacesUnauthorizedCode is the HTTP code returned for type ListCoreV1ServiceAccountForAllNamespacesUnauthorized
const ListCoreV1ServiceAccountForAllNamespacesUnauthorizedCode int = 401

/*ListCoreV1ServiceAccountForAllNamespacesUnauthorized Unauthorized

swagger:response listCoreV1ServiceAccountForAllNamespacesUnauthorized
*/
type ListCoreV1ServiceAccountForAllNamespacesUnauthorized struct {
}

// NewListCoreV1ServiceAccountForAllNamespacesUnauthorized creates ListCoreV1ServiceAccountForAllNamespacesUnauthorized with default headers values
func NewListCoreV1ServiceAccountForAllNamespacesUnauthorized() *ListCoreV1ServiceAccountForAllNamespacesUnauthorized {

	return &ListCoreV1ServiceAccountForAllNamespacesUnauthorized{}
}

// WriteResponse to the client
func (o *ListCoreV1ServiceAccountForAllNamespacesUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
