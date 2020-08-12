// Code generated by go-swagger; DO NOT EDIT.

package apps_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// DeleteAppsV1CollectionNamespacedReplicaSetOKCode is the HTTP code returned for type DeleteAppsV1CollectionNamespacedReplicaSetOK
const DeleteAppsV1CollectionNamespacedReplicaSetOKCode int = 200

/*DeleteAppsV1CollectionNamespacedReplicaSetOK OK

swagger:response deleteAppsV1CollectionNamespacedReplicaSetOK
*/
type DeleteAppsV1CollectionNamespacedReplicaSetOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1Status `json:"body,omitempty"`
}

// NewDeleteAppsV1CollectionNamespacedReplicaSetOK creates DeleteAppsV1CollectionNamespacedReplicaSetOK with default headers values
func NewDeleteAppsV1CollectionNamespacedReplicaSetOK() *DeleteAppsV1CollectionNamespacedReplicaSetOK {

	return &DeleteAppsV1CollectionNamespacedReplicaSetOK{}
}

// WithPayload adds the payload to the delete apps v1 collection namespaced replica set o k response
func (o *DeleteAppsV1CollectionNamespacedReplicaSetOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) *DeleteAppsV1CollectionNamespacedReplicaSetOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete apps v1 collection namespaced replica set o k response
func (o *DeleteAppsV1CollectionNamespacedReplicaSetOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteAppsV1CollectionNamespacedReplicaSetOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteAppsV1CollectionNamespacedReplicaSetUnauthorizedCode is the HTTP code returned for type DeleteAppsV1CollectionNamespacedReplicaSetUnauthorized
const DeleteAppsV1CollectionNamespacedReplicaSetUnauthorizedCode int = 401

/*DeleteAppsV1CollectionNamespacedReplicaSetUnauthorized Unauthorized

swagger:response deleteAppsV1CollectionNamespacedReplicaSetUnauthorized
*/
type DeleteAppsV1CollectionNamespacedReplicaSetUnauthorized struct {
}

// NewDeleteAppsV1CollectionNamespacedReplicaSetUnauthorized creates DeleteAppsV1CollectionNamespacedReplicaSetUnauthorized with default headers values
func NewDeleteAppsV1CollectionNamespacedReplicaSetUnauthorized() *DeleteAppsV1CollectionNamespacedReplicaSetUnauthorized {

	return &DeleteAppsV1CollectionNamespacedReplicaSetUnauthorized{}
}

// WriteResponse to the client
func (o *DeleteAppsV1CollectionNamespacedReplicaSetUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
