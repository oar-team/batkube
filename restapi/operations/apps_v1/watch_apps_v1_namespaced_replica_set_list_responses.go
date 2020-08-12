// Code generated by go-swagger; DO NOT EDIT.

package apps_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// WatchAppsV1NamespacedReplicaSetListOKCode is the HTTP code returned for type WatchAppsV1NamespacedReplicaSetListOK
const WatchAppsV1NamespacedReplicaSetListOKCode int = 200

/*WatchAppsV1NamespacedReplicaSetListOK OK

swagger:response watchAppsV1NamespacedReplicaSetListOK
*/
type WatchAppsV1NamespacedReplicaSetListOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent `json:"body,omitempty"`
}

// NewWatchAppsV1NamespacedReplicaSetListOK creates WatchAppsV1NamespacedReplicaSetListOK with default headers values
func NewWatchAppsV1NamespacedReplicaSetListOK() *WatchAppsV1NamespacedReplicaSetListOK {

	return &WatchAppsV1NamespacedReplicaSetListOK{}
}

// WithPayload adds the payload to the watch apps v1 namespaced replica set list o k response
func (o *WatchAppsV1NamespacedReplicaSetListOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) *WatchAppsV1NamespacedReplicaSetListOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the watch apps v1 namespaced replica set list o k response
func (o *WatchAppsV1NamespacedReplicaSetListOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WatchAppsV1NamespacedReplicaSetListOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WatchAppsV1NamespacedReplicaSetListUnauthorizedCode is the HTTP code returned for type WatchAppsV1NamespacedReplicaSetListUnauthorized
const WatchAppsV1NamespacedReplicaSetListUnauthorizedCode int = 401

/*WatchAppsV1NamespacedReplicaSetListUnauthorized Unauthorized

swagger:response watchAppsV1NamespacedReplicaSetListUnauthorized
*/
type WatchAppsV1NamespacedReplicaSetListUnauthorized struct {
}

// NewWatchAppsV1NamespacedReplicaSetListUnauthorized creates WatchAppsV1NamespacedReplicaSetListUnauthorized with default headers values
func NewWatchAppsV1NamespacedReplicaSetListUnauthorized() *WatchAppsV1NamespacedReplicaSetListUnauthorized {

	return &WatchAppsV1NamespacedReplicaSetListUnauthorized{}
}

// WriteResponse to the client
func (o *WatchAppsV1NamespacedReplicaSetListUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
