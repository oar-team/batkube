// Code generated by go-swagger; DO NOT EDIT.

package batch_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// DeleteBatchV1beta1CollectionNamespacedCronJobOKCode is the HTTP code returned for type DeleteBatchV1beta1CollectionNamespacedCronJobOK
const DeleteBatchV1beta1CollectionNamespacedCronJobOKCode int = 200

/*DeleteBatchV1beta1CollectionNamespacedCronJobOK OK

swagger:response deleteBatchV1beta1CollectionNamespacedCronJobOK
*/
type DeleteBatchV1beta1CollectionNamespacedCronJobOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1Status `json:"body,omitempty"`
}

// NewDeleteBatchV1beta1CollectionNamespacedCronJobOK creates DeleteBatchV1beta1CollectionNamespacedCronJobOK with default headers values
func NewDeleteBatchV1beta1CollectionNamespacedCronJobOK() *DeleteBatchV1beta1CollectionNamespacedCronJobOK {

	return &DeleteBatchV1beta1CollectionNamespacedCronJobOK{}
}

// WithPayload adds the payload to the delete batch v1beta1 collection namespaced cron job o k response
func (o *DeleteBatchV1beta1CollectionNamespacedCronJobOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) *DeleteBatchV1beta1CollectionNamespacedCronJobOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete batch v1beta1 collection namespaced cron job o k response
func (o *DeleteBatchV1beta1CollectionNamespacedCronJobOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteBatchV1beta1CollectionNamespacedCronJobOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteBatchV1beta1CollectionNamespacedCronJobUnauthorizedCode is the HTTP code returned for type DeleteBatchV1beta1CollectionNamespacedCronJobUnauthorized
const DeleteBatchV1beta1CollectionNamespacedCronJobUnauthorizedCode int = 401

/*DeleteBatchV1beta1CollectionNamespacedCronJobUnauthorized Unauthorized

swagger:response deleteBatchV1beta1CollectionNamespacedCronJobUnauthorized
*/
type DeleteBatchV1beta1CollectionNamespacedCronJobUnauthorized struct {
}

// NewDeleteBatchV1beta1CollectionNamespacedCronJobUnauthorized creates DeleteBatchV1beta1CollectionNamespacedCronJobUnauthorized with default headers values
func NewDeleteBatchV1beta1CollectionNamespacedCronJobUnauthorized() *DeleteBatchV1beta1CollectionNamespacedCronJobUnauthorized {

	return &DeleteBatchV1beta1CollectionNamespacedCronJobUnauthorized{}
}

// WriteResponse to the client
func (o *DeleteBatchV1beta1CollectionNamespacedCronJobUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
