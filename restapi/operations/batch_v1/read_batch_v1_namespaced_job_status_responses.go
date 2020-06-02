// Code generated by go-swagger; DO NOT EDIT.

package batch_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/models"
)

// ReadBatchV1NamespacedJobStatusOKCode is the HTTP code returned for type ReadBatchV1NamespacedJobStatusOK
const ReadBatchV1NamespacedJobStatusOKCode int = 200

/*ReadBatchV1NamespacedJobStatusOK OK

swagger:response readBatchV1NamespacedJobStatusOK
*/
type ReadBatchV1NamespacedJobStatusOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIBatchV1Job `json:"body,omitempty"`
}

// NewReadBatchV1NamespacedJobStatusOK creates ReadBatchV1NamespacedJobStatusOK with default headers values
func NewReadBatchV1NamespacedJobStatusOK() *ReadBatchV1NamespacedJobStatusOK {

	return &ReadBatchV1NamespacedJobStatusOK{}
}

// WithPayload adds the payload to the read batch v1 namespaced job status o k response
func (o *ReadBatchV1NamespacedJobStatusOK) WithPayload(payload *models.IoK8sAPIBatchV1Job) *ReadBatchV1NamespacedJobStatusOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the read batch v1 namespaced job status o k response
func (o *ReadBatchV1NamespacedJobStatusOK) SetPayload(payload *models.IoK8sAPIBatchV1Job) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReadBatchV1NamespacedJobStatusOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReadBatchV1NamespacedJobStatusUnauthorizedCode is the HTTP code returned for type ReadBatchV1NamespacedJobStatusUnauthorized
const ReadBatchV1NamespacedJobStatusUnauthorizedCode int = 401

/*ReadBatchV1NamespacedJobStatusUnauthorized Unauthorized

swagger:response readBatchV1NamespacedJobStatusUnauthorized
*/
type ReadBatchV1NamespacedJobStatusUnauthorized struct {
}

// NewReadBatchV1NamespacedJobStatusUnauthorized creates ReadBatchV1NamespacedJobStatusUnauthorized with default headers values
func NewReadBatchV1NamespacedJobStatusUnauthorized() *ReadBatchV1NamespacedJobStatusUnauthorized {

	return &ReadBatchV1NamespacedJobStatusUnauthorized{}
}

// WriteResponse to the client
func (o *ReadBatchV1NamespacedJobStatusUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
