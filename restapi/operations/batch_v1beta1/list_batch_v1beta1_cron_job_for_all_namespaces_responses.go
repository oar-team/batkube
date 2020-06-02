// Code generated by go-swagger; DO NOT EDIT.

package batch_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/models"
)

// ListBatchV1beta1CronJobForAllNamespacesOKCode is the HTTP code returned for type ListBatchV1beta1CronJobForAllNamespacesOK
const ListBatchV1beta1CronJobForAllNamespacesOKCode int = 200

/*ListBatchV1beta1CronJobForAllNamespacesOK OK

swagger:response listBatchV1beta1CronJobForAllNamespacesOK
*/
type ListBatchV1beta1CronJobForAllNamespacesOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIBatchV1beta1CronJobList `json:"body,omitempty"`
}

// NewListBatchV1beta1CronJobForAllNamespacesOK creates ListBatchV1beta1CronJobForAllNamespacesOK with default headers values
func NewListBatchV1beta1CronJobForAllNamespacesOK() *ListBatchV1beta1CronJobForAllNamespacesOK {

	return &ListBatchV1beta1CronJobForAllNamespacesOK{}
}

// WithPayload adds the payload to the list batch v1beta1 cron job for all namespaces o k response
func (o *ListBatchV1beta1CronJobForAllNamespacesOK) WithPayload(payload *models.IoK8sAPIBatchV1beta1CronJobList) *ListBatchV1beta1CronJobForAllNamespacesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list batch v1beta1 cron job for all namespaces o k response
func (o *ListBatchV1beta1CronJobForAllNamespacesOK) SetPayload(payload *models.IoK8sAPIBatchV1beta1CronJobList) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListBatchV1beta1CronJobForAllNamespacesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListBatchV1beta1CronJobForAllNamespacesUnauthorizedCode is the HTTP code returned for type ListBatchV1beta1CronJobForAllNamespacesUnauthorized
const ListBatchV1beta1CronJobForAllNamespacesUnauthorizedCode int = 401

/*ListBatchV1beta1CronJobForAllNamespacesUnauthorized Unauthorized

swagger:response listBatchV1beta1CronJobForAllNamespacesUnauthorized
*/
type ListBatchV1beta1CronJobForAllNamespacesUnauthorized struct {
}

// NewListBatchV1beta1CronJobForAllNamespacesUnauthorized creates ListBatchV1beta1CronJobForAllNamespacesUnauthorized with default headers values
func NewListBatchV1beta1CronJobForAllNamespacesUnauthorized() *ListBatchV1beta1CronJobForAllNamespacesUnauthorized {

	return &ListBatchV1beta1CronJobForAllNamespacesUnauthorized{}
}

// WriteResponse to the client
func (o *ListBatchV1beta1CronJobForAllNamespacesUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
