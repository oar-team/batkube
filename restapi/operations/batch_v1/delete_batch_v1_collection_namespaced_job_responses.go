// Code generated by go-swagger; DO NOT EDIT.

package batch_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// DeleteBatchV1CollectionNamespacedJobOKCode is the HTTP code returned for type DeleteBatchV1CollectionNamespacedJobOK
const DeleteBatchV1CollectionNamespacedJobOKCode int = 200

/*DeleteBatchV1CollectionNamespacedJobOK OK

swagger:response deleteBatchV1CollectionNamespacedJobOK
*/
type DeleteBatchV1CollectionNamespacedJobOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1Status `json:"body,omitempty"`
}

// NewDeleteBatchV1CollectionNamespacedJobOK creates DeleteBatchV1CollectionNamespacedJobOK with default headers values
func NewDeleteBatchV1CollectionNamespacedJobOK() *DeleteBatchV1CollectionNamespacedJobOK {

	return &DeleteBatchV1CollectionNamespacedJobOK{}
}

// WithPayload adds the payload to the delete batch v1 collection namespaced job o k response
func (o *DeleteBatchV1CollectionNamespacedJobOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) *DeleteBatchV1CollectionNamespacedJobOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete batch v1 collection namespaced job o k response
func (o *DeleteBatchV1CollectionNamespacedJobOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteBatchV1CollectionNamespacedJobOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteBatchV1CollectionNamespacedJobUnauthorizedCode is the HTTP code returned for type DeleteBatchV1CollectionNamespacedJobUnauthorized
const DeleteBatchV1CollectionNamespacedJobUnauthorizedCode int = 401

/*DeleteBatchV1CollectionNamespacedJobUnauthorized Unauthorized

swagger:response deleteBatchV1CollectionNamespacedJobUnauthorized
*/
type DeleteBatchV1CollectionNamespacedJobUnauthorized struct {
}

// NewDeleteBatchV1CollectionNamespacedJobUnauthorized creates DeleteBatchV1CollectionNamespacedJobUnauthorized with default headers values
func NewDeleteBatchV1CollectionNamespacedJobUnauthorized() *DeleteBatchV1CollectionNamespacedJobUnauthorized {

	return &DeleteBatchV1CollectionNamespacedJobUnauthorized{}
}

// WriteResponse to the client
func (o *DeleteBatchV1CollectionNamespacedJobUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
