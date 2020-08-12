// Code generated by go-swagger; DO NOT EDIT.

package apps_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// DeleteAppsV1CollectionNamespacedStatefulSetOKCode is the HTTP code returned for type DeleteAppsV1CollectionNamespacedStatefulSetOK
const DeleteAppsV1CollectionNamespacedStatefulSetOKCode int = 200

/*DeleteAppsV1CollectionNamespacedStatefulSetOK OK

swagger:response deleteAppsV1CollectionNamespacedStatefulSetOK
*/
type DeleteAppsV1CollectionNamespacedStatefulSetOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1Status `json:"body,omitempty"`
}

// NewDeleteAppsV1CollectionNamespacedStatefulSetOK creates DeleteAppsV1CollectionNamespacedStatefulSetOK with default headers values
func NewDeleteAppsV1CollectionNamespacedStatefulSetOK() *DeleteAppsV1CollectionNamespacedStatefulSetOK {

	return &DeleteAppsV1CollectionNamespacedStatefulSetOK{}
}

// WithPayload adds the payload to the delete apps v1 collection namespaced stateful set o k response
func (o *DeleteAppsV1CollectionNamespacedStatefulSetOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) *DeleteAppsV1CollectionNamespacedStatefulSetOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete apps v1 collection namespaced stateful set o k response
func (o *DeleteAppsV1CollectionNamespacedStatefulSetOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteAppsV1CollectionNamespacedStatefulSetOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteAppsV1CollectionNamespacedStatefulSetUnauthorizedCode is the HTTP code returned for type DeleteAppsV1CollectionNamespacedStatefulSetUnauthorized
const DeleteAppsV1CollectionNamespacedStatefulSetUnauthorizedCode int = 401

/*DeleteAppsV1CollectionNamespacedStatefulSetUnauthorized Unauthorized

swagger:response deleteAppsV1CollectionNamespacedStatefulSetUnauthorized
*/
type DeleteAppsV1CollectionNamespacedStatefulSetUnauthorized struct {
}

// NewDeleteAppsV1CollectionNamespacedStatefulSetUnauthorized creates DeleteAppsV1CollectionNamespacedStatefulSetUnauthorized with default headers values
func NewDeleteAppsV1CollectionNamespacedStatefulSetUnauthorized() *DeleteAppsV1CollectionNamespacedStatefulSetUnauthorized {

	return &DeleteAppsV1CollectionNamespacedStatefulSetUnauthorized{}
}

// WriteResponse to the client
func (o *DeleteAppsV1CollectionNamespacedStatefulSetUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
