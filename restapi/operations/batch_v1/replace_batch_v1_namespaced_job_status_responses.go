// Code generated by go-swagger; DO NOT EDIT.

package batch_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/models"
)

// ReplaceBatchV1NamespacedJobStatusOKCode is the HTTP code returned for type ReplaceBatchV1NamespacedJobStatusOK
const ReplaceBatchV1NamespacedJobStatusOKCode int = 200

/*ReplaceBatchV1NamespacedJobStatusOK OK

swagger:response replaceBatchV1NamespacedJobStatusOK
*/
type ReplaceBatchV1NamespacedJobStatusOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIBatchV1Job `json:"body,omitempty"`
}

// NewReplaceBatchV1NamespacedJobStatusOK creates ReplaceBatchV1NamespacedJobStatusOK with default headers values
func NewReplaceBatchV1NamespacedJobStatusOK() *ReplaceBatchV1NamespacedJobStatusOK {

	return &ReplaceBatchV1NamespacedJobStatusOK{}
}

// WithPayload adds the payload to the replace batch v1 namespaced job status o k response
func (o *ReplaceBatchV1NamespacedJobStatusOK) WithPayload(payload *models.IoK8sAPIBatchV1Job) *ReplaceBatchV1NamespacedJobStatusOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace batch v1 namespaced job status o k response
func (o *ReplaceBatchV1NamespacedJobStatusOK) SetPayload(payload *models.IoK8sAPIBatchV1Job) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceBatchV1NamespacedJobStatusOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceBatchV1NamespacedJobStatusCreatedCode is the HTTP code returned for type ReplaceBatchV1NamespacedJobStatusCreated
const ReplaceBatchV1NamespacedJobStatusCreatedCode int = 201

/*ReplaceBatchV1NamespacedJobStatusCreated Created

swagger:response replaceBatchV1NamespacedJobStatusCreated
*/
type ReplaceBatchV1NamespacedJobStatusCreated struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIBatchV1Job `json:"body,omitempty"`
}

// NewReplaceBatchV1NamespacedJobStatusCreated creates ReplaceBatchV1NamespacedJobStatusCreated with default headers values
func NewReplaceBatchV1NamespacedJobStatusCreated() *ReplaceBatchV1NamespacedJobStatusCreated {

	return &ReplaceBatchV1NamespacedJobStatusCreated{}
}

// WithPayload adds the payload to the replace batch v1 namespaced job status created response
func (o *ReplaceBatchV1NamespacedJobStatusCreated) WithPayload(payload *models.IoK8sAPIBatchV1Job) *ReplaceBatchV1NamespacedJobStatusCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace batch v1 namespaced job status created response
func (o *ReplaceBatchV1NamespacedJobStatusCreated) SetPayload(payload *models.IoK8sAPIBatchV1Job) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceBatchV1NamespacedJobStatusCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceBatchV1NamespacedJobStatusUnauthorizedCode is the HTTP code returned for type ReplaceBatchV1NamespacedJobStatusUnauthorized
const ReplaceBatchV1NamespacedJobStatusUnauthorizedCode int = 401

/*ReplaceBatchV1NamespacedJobStatusUnauthorized Unauthorized

swagger:response replaceBatchV1NamespacedJobStatusUnauthorized
*/
type ReplaceBatchV1NamespacedJobStatusUnauthorized struct {
}

// NewReplaceBatchV1NamespacedJobStatusUnauthorized creates ReplaceBatchV1NamespacedJobStatusUnauthorized with default headers values
func NewReplaceBatchV1NamespacedJobStatusUnauthorized() *ReplaceBatchV1NamespacedJobStatusUnauthorized {

	return &ReplaceBatchV1NamespacedJobStatusUnauthorized{}
}

// WriteResponse to the client
func (o *ReplaceBatchV1NamespacedJobStatusUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
