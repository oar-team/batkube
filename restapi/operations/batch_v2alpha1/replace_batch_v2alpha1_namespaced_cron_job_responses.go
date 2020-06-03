// Code generated by go-swagger; DO NOT EDIT.

package batch_v2alpha1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/models"
)

// ReplaceBatchV2alpha1NamespacedCronJobOKCode is the HTTP code returned for type ReplaceBatchV2alpha1NamespacedCronJobOK
const ReplaceBatchV2alpha1NamespacedCronJobOKCode int = 200

/*ReplaceBatchV2alpha1NamespacedCronJobOK OK

swagger:response replaceBatchV2alpha1NamespacedCronJobOK
*/
type ReplaceBatchV2alpha1NamespacedCronJobOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIBatchV2alpha1CronJob `json:"body,omitempty"`
}

// NewReplaceBatchV2alpha1NamespacedCronJobOK creates ReplaceBatchV2alpha1NamespacedCronJobOK with default headers values
func NewReplaceBatchV2alpha1NamespacedCronJobOK() *ReplaceBatchV2alpha1NamespacedCronJobOK {

	return &ReplaceBatchV2alpha1NamespacedCronJobOK{}
}

// WithPayload adds the payload to the replace batch v2alpha1 namespaced cron job o k response
func (o *ReplaceBatchV2alpha1NamespacedCronJobOK) WithPayload(payload *models.IoK8sAPIBatchV2alpha1CronJob) *ReplaceBatchV2alpha1NamespacedCronJobOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace batch v2alpha1 namespaced cron job o k response
func (o *ReplaceBatchV2alpha1NamespacedCronJobOK) SetPayload(payload *models.IoK8sAPIBatchV2alpha1CronJob) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceBatchV2alpha1NamespacedCronJobOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceBatchV2alpha1NamespacedCronJobCreatedCode is the HTTP code returned for type ReplaceBatchV2alpha1NamespacedCronJobCreated
const ReplaceBatchV2alpha1NamespacedCronJobCreatedCode int = 201

/*ReplaceBatchV2alpha1NamespacedCronJobCreated Created

swagger:response replaceBatchV2alpha1NamespacedCronJobCreated
*/
type ReplaceBatchV2alpha1NamespacedCronJobCreated struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIBatchV2alpha1CronJob `json:"body,omitempty"`
}

// NewReplaceBatchV2alpha1NamespacedCronJobCreated creates ReplaceBatchV2alpha1NamespacedCronJobCreated with default headers values
func NewReplaceBatchV2alpha1NamespacedCronJobCreated() *ReplaceBatchV2alpha1NamespacedCronJobCreated {

	return &ReplaceBatchV2alpha1NamespacedCronJobCreated{}
}

// WithPayload adds the payload to the replace batch v2alpha1 namespaced cron job created response
func (o *ReplaceBatchV2alpha1NamespacedCronJobCreated) WithPayload(payload *models.IoK8sAPIBatchV2alpha1CronJob) *ReplaceBatchV2alpha1NamespacedCronJobCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace batch v2alpha1 namespaced cron job created response
func (o *ReplaceBatchV2alpha1NamespacedCronJobCreated) SetPayload(payload *models.IoK8sAPIBatchV2alpha1CronJob) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceBatchV2alpha1NamespacedCronJobCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceBatchV2alpha1NamespacedCronJobUnauthorizedCode is the HTTP code returned for type ReplaceBatchV2alpha1NamespacedCronJobUnauthorized
const ReplaceBatchV2alpha1NamespacedCronJobUnauthorizedCode int = 401

/*ReplaceBatchV2alpha1NamespacedCronJobUnauthorized Unauthorized

swagger:response replaceBatchV2alpha1NamespacedCronJobUnauthorized
*/
type ReplaceBatchV2alpha1NamespacedCronJobUnauthorized struct {
}

// NewReplaceBatchV2alpha1NamespacedCronJobUnauthorized creates ReplaceBatchV2alpha1NamespacedCronJobUnauthorized with default headers values
func NewReplaceBatchV2alpha1NamespacedCronJobUnauthorized() *ReplaceBatchV2alpha1NamespacedCronJobUnauthorized {

	return &ReplaceBatchV2alpha1NamespacedCronJobUnauthorized{}
}

// WriteResponse to the client
func (o *ReplaceBatchV2alpha1NamespacedCronJobUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}