// Code generated by go-swagger; DO NOT EDIT.

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/models"
)

// DeleteCoreV1CollectionNamespacedServiceAccountOKCode is the HTTP code returned for type DeleteCoreV1CollectionNamespacedServiceAccountOK
const DeleteCoreV1CollectionNamespacedServiceAccountOKCode int = 200

/*DeleteCoreV1CollectionNamespacedServiceAccountOK OK

swagger:response deleteCoreV1CollectionNamespacedServiceAccountOK
*/
type DeleteCoreV1CollectionNamespacedServiceAccountOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1Status `json:"body,omitempty"`
}

// NewDeleteCoreV1CollectionNamespacedServiceAccountOK creates DeleteCoreV1CollectionNamespacedServiceAccountOK with default headers values
func NewDeleteCoreV1CollectionNamespacedServiceAccountOK() *DeleteCoreV1CollectionNamespacedServiceAccountOK {

	return &DeleteCoreV1CollectionNamespacedServiceAccountOK{}
}

// WithPayload adds the payload to the delete core v1 collection namespaced service account o k response
func (o *DeleteCoreV1CollectionNamespacedServiceAccountOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) *DeleteCoreV1CollectionNamespacedServiceAccountOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete core v1 collection namespaced service account o k response
func (o *DeleteCoreV1CollectionNamespacedServiceAccountOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteCoreV1CollectionNamespacedServiceAccountOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteCoreV1CollectionNamespacedServiceAccountUnauthorizedCode is the HTTP code returned for type DeleteCoreV1CollectionNamespacedServiceAccountUnauthorized
const DeleteCoreV1CollectionNamespacedServiceAccountUnauthorizedCode int = 401

/*DeleteCoreV1CollectionNamespacedServiceAccountUnauthorized Unauthorized

swagger:response deleteCoreV1CollectionNamespacedServiceAccountUnauthorized
*/
type DeleteCoreV1CollectionNamespacedServiceAccountUnauthorized struct {
}

// NewDeleteCoreV1CollectionNamespacedServiceAccountUnauthorized creates DeleteCoreV1CollectionNamespacedServiceAccountUnauthorized with default headers values
func NewDeleteCoreV1CollectionNamespacedServiceAccountUnauthorized() *DeleteCoreV1CollectionNamespacedServiceAccountUnauthorized {

	return &DeleteCoreV1CollectionNamespacedServiceAccountUnauthorized{}
}

// WriteResponse to the client
func (o *DeleteCoreV1CollectionNamespacedServiceAccountUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}