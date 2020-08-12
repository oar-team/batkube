// Code generated by go-swagger; DO NOT EDIT.

package apps_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// WatchAppsV1NamespacedDeploymentListOKCode is the HTTP code returned for type WatchAppsV1NamespacedDeploymentListOK
const WatchAppsV1NamespacedDeploymentListOKCode int = 200

/*WatchAppsV1NamespacedDeploymentListOK OK

swagger:response watchAppsV1NamespacedDeploymentListOK
*/
type WatchAppsV1NamespacedDeploymentListOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent `json:"body,omitempty"`
}

// NewWatchAppsV1NamespacedDeploymentListOK creates WatchAppsV1NamespacedDeploymentListOK with default headers values
func NewWatchAppsV1NamespacedDeploymentListOK() *WatchAppsV1NamespacedDeploymentListOK {

	return &WatchAppsV1NamespacedDeploymentListOK{}
}

// WithPayload adds the payload to the watch apps v1 namespaced deployment list o k response
func (o *WatchAppsV1NamespacedDeploymentListOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) *WatchAppsV1NamespacedDeploymentListOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the watch apps v1 namespaced deployment list o k response
func (o *WatchAppsV1NamespacedDeploymentListOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WatchAppsV1NamespacedDeploymentListOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WatchAppsV1NamespacedDeploymentListUnauthorizedCode is the HTTP code returned for type WatchAppsV1NamespacedDeploymentListUnauthorized
const WatchAppsV1NamespacedDeploymentListUnauthorizedCode int = 401

/*WatchAppsV1NamespacedDeploymentListUnauthorized Unauthorized

swagger:response watchAppsV1NamespacedDeploymentListUnauthorized
*/
type WatchAppsV1NamespacedDeploymentListUnauthorized struct {
}

// NewWatchAppsV1NamespacedDeploymentListUnauthorized creates WatchAppsV1NamespacedDeploymentListUnauthorized with default headers values
func NewWatchAppsV1NamespacedDeploymentListUnauthorized() *WatchAppsV1NamespacedDeploymentListUnauthorized {

	return &WatchAppsV1NamespacedDeploymentListUnauthorized{}
}

// WriteResponse to the client
func (o *WatchAppsV1NamespacedDeploymentListUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
