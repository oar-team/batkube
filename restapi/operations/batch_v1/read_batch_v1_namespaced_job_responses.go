// Code generated by go-swagger; DO NOT EDIT.

package batch_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/models"
)

// ReadBatchV1NamespacedJobOKCode is the HTTP code returned for type ReadBatchV1NamespacedJobOK
const ReadBatchV1NamespacedJobOKCode int = 200

/*ReadBatchV1NamespacedJobOK OK

swagger:response readBatchV1NamespacedJobOK
*/
type ReadBatchV1NamespacedJobOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIBatchV1Job `json:"body,omitempty"`
}

// NewReadBatchV1NamespacedJobOK creates ReadBatchV1NamespacedJobOK with default headers values
func NewReadBatchV1NamespacedJobOK() *ReadBatchV1NamespacedJobOK {

	return &ReadBatchV1NamespacedJobOK{}
}

// WithPayload adds the payload to the read batch v1 namespaced job o k response
func (o *ReadBatchV1NamespacedJobOK) WithPayload(payload *models.IoK8sAPIBatchV1Job) *ReadBatchV1NamespacedJobOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the read batch v1 namespaced job o k response
func (o *ReadBatchV1NamespacedJobOK) SetPayload(payload *models.IoK8sAPIBatchV1Job) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReadBatchV1NamespacedJobOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReadBatchV1NamespacedJobUnauthorizedCode is the HTTP code returned for type ReadBatchV1NamespacedJobUnauthorized
const ReadBatchV1NamespacedJobUnauthorizedCode int = 401

/*ReadBatchV1NamespacedJobUnauthorized Unauthorized

swagger:response readBatchV1NamespacedJobUnauthorized
*/
type ReadBatchV1NamespacedJobUnauthorized struct {
}

// NewReadBatchV1NamespacedJobUnauthorized creates ReadBatchV1NamespacedJobUnauthorized with default headers values
func NewReadBatchV1NamespacedJobUnauthorized() *ReadBatchV1NamespacedJobUnauthorized {

	return &ReadBatchV1NamespacedJobUnauthorized{}
}

// WriteResponse to the client
func (o *ReadBatchV1NamespacedJobUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}