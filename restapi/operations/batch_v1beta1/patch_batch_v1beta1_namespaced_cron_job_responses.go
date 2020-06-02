// Code generated by go-swagger; DO NOT EDIT.

package batch_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/models"
)

// PatchBatchV1beta1NamespacedCronJobOKCode is the HTTP code returned for type PatchBatchV1beta1NamespacedCronJobOK
const PatchBatchV1beta1NamespacedCronJobOKCode int = 200

/*PatchBatchV1beta1NamespacedCronJobOK OK

swagger:response patchBatchV1beta1NamespacedCronJobOK
*/
type PatchBatchV1beta1NamespacedCronJobOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIBatchV1beta1CronJob `json:"body,omitempty"`
}

// NewPatchBatchV1beta1NamespacedCronJobOK creates PatchBatchV1beta1NamespacedCronJobOK with default headers values
func NewPatchBatchV1beta1NamespacedCronJobOK() *PatchBatchV1beta1NamespacedCronJobOK {

	return &PatchBatchV1beta1NamespacedCronJobOK{}
}

// WithPayload adds the payload to the patch batch v1beta1 namespaced cron job o k response
func (o *PatchBatchV1beta1NamespacedCronJobOK) WithPayload(payload *models.IoK8sAPIBatchV1beta1CronJob) *PatchBatchV1beta1NamespacedCronJobOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch batch v1beta1 namespaced cron job o k response
func (o *PatchBatchV1beta1NamespacedCronJobOK) SetPayload(payload *models.IoK8sAPIBatchV1beta1CronJob) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchBatchV1beta1NamespacedCronJobOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchBatchV1beta1NamespacedCronJobUnauthorizedCode is the HTTP code returned for type PatchBatchV1beta1NamespacedCronJobUnauthorized
const PatchBatchV1beta1NamespacedCronJobUnauthorizedCode int = 401

/*PatchBatchV1beta1NamespacedCronJobUnauthorized Unauthorized

swagger:response patchBatchV1beta1NamespacedCronJobUnauthorized
*/
type PatchBatchV1beta1NamespacedCronJobUnauthorized struct {
}

// NewPatchBatchV1beta1NamespacedCronJobUnauthorized creates PatchBatchV1beta1NamespacedCronJobUnauthorized with default headers values
func NewPatchBatchV1beta1NamespacedCronJobUnauthorized() *PatchBatchV1beta1NamespacedCronJobUnauthorized {

	return &PatchBatchV1beta1NamespacedCronJobUnauthorized{}
}

// WriteResponse to the client
func (o *PatchBatchV1beta1NamespacedCronJobUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
