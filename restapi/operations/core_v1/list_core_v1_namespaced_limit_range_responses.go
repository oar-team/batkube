// Code generated by go-swagger; DO NOT EDIT.

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// ListCoreV1NamespacedLimitRangeOKCode is the HTTP code returned for type ListCoreV1NamespacedLimitRangeOK
const ListCoreV1NamespacedLimitRangeOKCode int = 200

/*ListCoreV1NamespacedLimitRangeOK OK

swagger:response listCoreV1NamespacedLimitRangeOK
*/
type ListCoreV1NamespacedLimitRangeOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPICoreV1LimitRangeList `json:"body,omitempty"`
}

// NewListCoreV1NamespacedLimitRangeOK creates ListCoreV1NamespacedLimitRangeOK with default headers values
func NewListCoreV1NamespacedLimitRangeOK() *ListCoreV1NamespacedLimitRangeOK {

	return &ListCoreV1NamespacedLimitRangeOK{}
}

// WithPayload adds the payload to the list core v1 namespaced limit range o k response
func (o *ListCoreV1NamespacedLimitRangeOK) WithPayload(payload *models.IoK8sAPICoreV1LimitRangeList) *ListCoreV1NamespacedLimitRangeOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list core v1 namespaced limit range o k response
func (o *ListCoreV1NamespacedLimitRangeOK) SetPayload(payload *models.IoK8sAPICoreV1LimitRangeList) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListCoreV1NamespacedLimitRangeOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListCoreV1NamespacedLimitRangeUnauthorizedCode is the HTTP code returned for type ListCoreV1NamespacedLimitRangeUnauthorized
const ListCoreV1NamespacedLimitRangeUnauthorizedCode int = 401

/*ListCoreV1NamespacedLimitRangeUnauthorized Unauthorized

swagger:response listCoreV1NamespacedLimitRangeUnauthorized
*/
type ListCoreV1NamespacedLimitRangeUnauthorized struct {
}

// NewListCoreV1NamespacedLimitRangeUnauthorized creates ListCoreV1NamespacedLimitRangeUnauthorized with default headers values
func NewListCoreV1NamespacedLimitRangeUnauthorized() *ListCoreV1NamespacedLimitRangeUnauthorized {

	return &ListCoreV1NamespacedLimitRangeUnauthorized{}
}

// WriteResponse to the client
func (o *ListCoreV1NamespacedLimitRangeUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
