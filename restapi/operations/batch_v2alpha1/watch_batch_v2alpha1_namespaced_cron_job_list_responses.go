// Code generated by go-swagger; DO NOT EDIT.

package batch_v2alpha1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// WatchBatchV2alpha1NamespacedCronJobListOKCode is the HTTP code returned for type WatchBatchV2alpha1NamespacedCronJobListOK
const WatchBatchV2alpha1NamespacedCronJobListOKCode int = 200

/*WatchBatchV2alpha1NamespacedCronJobListOK OK

swagger:response watchBatchV2alpha1NamespacedCronJobListOK
*/
type WatchBatchV2alpha1NamespacedCronJobListOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent `json:"body,omitempty"`
}

// NewWatchBatchV2alpha1NamespacedCronJobListOK creates WatchBatchV2alpha1NamespacedCronJobListOK with default headers values
func NewWatchBatchV2alpha1NamespacedCronJobListOK() *WatchBatchV2alpha1NamespacedCronJobListOK {

	return &WatchBatchV2alpha1NamespacedCronJobListOK{}
}

// WithPayload adds the payload to the watch batch v2alpha1 namespaced cron job list o k response
func (o *WatchBatchV2alpha1NamespacedCronJobListOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) *WatchBatchV2alpha1NamespacedCronJobListOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the watch batch v2alpha1 namespaced cron job list o k response
func (o *WatchBatchV2alpha1NamespacedCronJobListOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WatchBatchV2alpha1NamespacedCronJobListOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WatchBatchV2alpha1NamespacedCronJobListUnauthorizedCode is the HTTP code returned for type WatchBatchV2alpha1NamespacedCronJobListUnauthorized
const WatchBatchV2alpha1NamespacedCronJobListUnauthorizedCode int = 401

/*WatchBatchV2alpha1NamespacedCronJobListUnauthorized Unauthorized

swagger:response watchBatchV2alpha1NamespacedCronJobListUnauthorized
*/
type WatchBatchV2alpha1NamespacedCronJobListUnauthorized struct {
}

// NewWatchBatchV2alpha1NamespacedCronJobListUnauthorized creates WatchBatchV2alpha1NamespacedCronJobListUnauthorized with default headers values
func NewWatchBatchV2alpha1NamespacedCronJobListUnauthorized() *WatchBatchV2alpha1NamespacedCronJobListUnauthorized {

	return &WatchBatchV2alpha1NamespacedCronJobListUnauthorized{}
}

// WriteResponse to the client
func (o *WatchBatchV2alpha1NamespacedCronJobListUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
