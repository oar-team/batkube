// Code generated by go-swagger; DO NOT EDIT.

package batch_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/models"
)

// ListBatchV1NamespacedJobOKCode is the HTTP code returned for type ListBatchV1NamespacedJobOK
const ListBatchV1NamespacedJobOKCode int = 200

/*ListBatchV1NamespacedJobOK OK

swagger:response listBatchV1NamespacedJobOK
*/
type ListBatchV1NamespacedJobOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIBatchV1JobList `json:"body,omitempty"`
}

// NewListBatchV1NamespacedJobOK creates ListBatchV1NamespacedJobOK with default headers values
func NewListBatchV1NamespacedJobOK() *ListBatchV1NamespacedJobOK {

	return &ListBatchV1NamespacedJobOK{}
}

// WithPayload adds the payload to the list batch v1 namespaced job o k response
func (o *ListBatchV1NamespacedJobOK) WithPayload(payload *models.IoK8sAPIBatchV1JobList) *ListBatchV1NamespacedJobOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list batch v1 namespaced job o k response
func (o *ListBatchV1NamespacedJobOK) SetPayload(payload *models.IoK8sAPIBatchV1JobList) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListBatchV1NamespacedJobOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListBatchV1NamespacedJobUnauthorizedCode is the HTTP code returned for type ListBatchV1NamespacedJobUnauthorized
const ListBatchV1NamespacedJobUnauthorizedCode int = 401

/*ListBatchV1NamespacedJobUnauthorized Unauthorized

swagger:response listBatchV1NamespacedJobUnauthorized
*/
type ListBatchV1NamespacedJobUnauthorized struct {
}

// NewListBatchV1NamespacedJobUnauthorized creates ListBatchV1NamespacedJobUnauthorized with default headers values
func NewListBatchV1NamespacedJobUnauthorized() *ListBatchV1NamespacedJobUnauthorized {

	return &ListBatchV1NamespacedJobUnauthorized{}
}

// WriteResponse to the client
func (o *ListBatchV1NamespacedJobUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}