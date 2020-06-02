// Code generated by go-swagger; DO NOT EDIT.

package apps_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/models"
)

// WatchAppsV1ReplicaSetListForAllNamespacesOKCode is the HTTP code returned for type WatchAppsV1ReplicaSetListForAllNamespacesOK
const WatchAppsV1ReplicaSetListForAllNamespacesOKCode int = 200

/*WatchAppsV1ReplicaSetListForAllNamespacesOK OK

swagger:response watchAppsV1ReplicaSetListForAllNamespacesOK
*/
type WatchAppsV1ReplicaSetListForAllNamespacesOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent `json:"body,omitempty"`
}

// NewWatchAppsV1ReplicaSetListForAllNamespacesOK creates WatchAppsV1ReplicaSetListForAllNamespacesOK with default headers values
func NewWatchAppsV1ReplicaSetListForAllNamespacesOK() *WatchAppsV1ReplicaSetListForAllNamespacesOK {

	return &WatchAppsV1ReplicaSetListForAllNamespacesOK{}
}

// WithPayload adds the payload to the watch apps v1 replica set list for all namespaces o k response
func (o *WatchAppsV1ReplicaSetListForAllNamespacesOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) *WatchAppsV1ReplicaSetListForAllNamespacesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the watch apps v1 replica set list for all namespaces o k response
func (o *WatchAppsV1ReplicaSetListForAllNamespacesOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WatchAppsV1ReplicaSetListForAllNamespacesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WatchAppsV1ReplicaSetListForAllNamespacesUnauthorizedCode is the HTTP code returned for type WatchAppsV1ReplicaSetListForAllNamespacesUnauthorized
const WatchAppsV1ReplicaSetListForAllNamespacesUnauthorizedCode int = 401

/*WatchAppsV1ReplicaSetListForAllNamespacesUnauthorized Unauthorized

swagger:response watchAppsV1ReplicaSetListForAllNamespacesUnauthorized
*/
type WatchAppsV1ReplicaSetListForAllNamespacesUnauthorized struct {
}

// NewWatchAppsV1ReplicaSetListForAllNamespacesUnauthorized creates WatchAppsV1ReplicaSetListForAllNamespacesUnauthorized with default headers values
func NewWatchAppsV1ReplicaSetListForAllNamespacesUnauthorized() *WatchAppsV1ReplicaSetListForAllNamespacesUnauthorized {

	return &WatchAppsV1ReplicaSetListForAllNamespacesUnauthorized{}
}

// WriteResponse to the client
func (o *WatchAppsV1ReplicaSetListForAllNamespacesUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
