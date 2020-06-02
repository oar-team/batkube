// Code generated by go-swagger; DO NOT EDIT.

package batch_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/models"
)

// ReplaceBatchV1NamespacedJobOKCode is the HTTP code returned for type ReplaceBatchV1NamespacedJobOK
const ReplaceBatchV1NamespacedJobOKCode int = 200

/*ReplaceBatchV1NamespacedJobOK OK

swagger:response replaceBatchV1NamespacedJobOK
*/
type ReplaceBatchV1NamespacedJobOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIBatchV1Job `json:"body,omitempty"`
}

// NewReplaceBatchV1NamespacedJobOK creates ReplaceBatchV1NamespacedJobOK with default headers values
func NewReplaceBatchV1NamespacedJobOK() *ReplaceBatchV1NamespacedJobOK {

	return &ReplaceBatchV1NamespacedJobOK{}
}

// WithPayload adds the payload to the replace batch v1 namespaced job o k response
func (o *ReplaceBatchV1NamespacedJobOK) WithPayload(payload *models.IoK8sAPIBatchV1Job) *ReplaceBatchV1NamespacedJobOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace batch v1 namespaced job o k response
func (o *ReplaceBatchV1NamespacedJobOK) SetPayload(payload *models.IoK8sAPIBatchV1Job) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceBatchV1NamespacedJobOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceBatchV1NamespacedJobCreatedCode is the HTTP code returned for type ReplaceBatchV1NamespacedJobCreated
const ReplaceBatchV1NamespacedJobCreatedCode int = 201

/*ReplaceBatchV1NamespacedJobCreated Created

swagger:response replaceBatchV1NamespacedJobCreated
*/
type ReplaceBatchV1NamespacedJobCreated struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIBatchV1Job `json:"body,omitempty"`
}

// NewReplaceBatchV1NamespacedJobCreated creates ReplaceBatchV1NamespacedJobCreated with default headers values
func NewReplaceBatchV1NamespacedJobCreated() *ReplaceBatchV1NamespacedJobCreated {

	return &ReplaceBatchV1NamespacedJobCreated{}
}

// WithPayload adds the payload to the replace batch v1 namespaced job created response
func (o *ReplaceBatchV1NamespacedJobCreated) WithPayload(payload *models.IoK8sAPIBatchV1Job) *ReplaceBatchV1NamespacedJobCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace batch v1 namespaced job created response
func (o *ReplaceBatchV1NamespacedJobCreated) SetPayload(payload *models.IoK8sAPIBatchV1Job) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceBatchV1NamespacedJobCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceBatchV1NamespacedJobUnauthorizedCode is the HTTP code returned for type ReplaceBatchV1NamespacedJobUnauthorized
const ReplaceBatchV1NamespacedJobUnauthorizedCode int = 401

/*ReplaceBatchV1NamespacedJobUnauthorized Unauthorized

swagger:response replaceBatchV1NamespacedJobUnauthorized
*/
type ReplaceBatchV1NamespacedJobUnauthorized struct {
}

// NewReplaceBatchV1NamespacedJobUnauthorized creates ReplaceBatchV1NamespacedJobUnauthorized with default headers values
func NewReplaceBatchV1NamespacedJobUnauthorized() *ReplaceBatchV1NamespacedJobUnauthorized {

	return &ReplaceBatchV1NamespacedJobUnauthorized{}
}

// WriteResponse to the client
func (o *ReplaceBatchV1NamespacedJobUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
