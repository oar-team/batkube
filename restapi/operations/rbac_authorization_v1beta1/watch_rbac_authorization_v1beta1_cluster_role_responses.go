// Code generated by go-swagger; DO NOT EDIT.

package rbac_authorization_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// WatchRbacAuthorizationV1beta1ClusterRoleOKCode is the HTTP code returned for type WatchRbacAuthorizationV1beta1ClusterRoleOK
const WatchRbacAuthorizationV1beta1ClusterRoleOKCode int = 200

/*WatchRbacAuthorizationV1beta1ClusterRoleOK OK

swagger:response watchRbacAuthorizationV1beta1ClusterRoleOK
*/
type WatchRbacAuthorizationV1beta1ClusterRoleOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent `json:"body,omitempty"`
}

// NewWatchRbacAuthorizationV1beta1ClusterRoleOK creates WatchRbacAuthorizationV1beta1ClusterRoleOK with default headers values
func NewWatchRbacAuthorizationV1beta1ClusterRoleOK() *WatchRbacAuthorizationV1beta1ClusterRoleOK {

	return &WatchRbacAuthorizationV1beta1ClusterRoleOK{}
}

// WithPayload adds the payload to the watch rbac authorization v1beta1 cluster role o k response
func (o *WatchRbacAuthorizationV1beta1ClusterRoleOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) *WatchRbacAuthorizationV1beta1ClusterRoleOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the watch rbac authorization v1beta1 cluster role o k response
func (o *WatchRbacAuthorizationV1beta1ClusterRoleOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WatchRbacAuthorizationV1beta1ClusterRoleOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WatchRbacAuthorizationV1beta1ClusterRoleUnauthorizedCode is the HTTP code returned for type WatchRbacAuthorizationV1beta1ClusterRoleUnauthorized
const WatchRbacAuthorizationV1beta1ClusterRoleUnauthorizedCode int = 401

/*WatchRbacAuthorizationV1beta1ClusterRoleUnauthorized Unauthorized

swagger:response watchRbacAuthorizationV1beta1ClusterRoleUnauthorized
*/
type WatchRbacAuthorizationV1beta1ClusterRoleUnauthorized struct {
}

// NewWatchRbacAuthorizationV1beta1ClusterRoleUnauthorized creates WatchRbacAuthorizationV1beta1ClusterRoleUnauthorized with default headers values
func NewWatchRbacAuthorizationV1beta1ClusterRoleUnauthorized() *WatchRbacAuthorizationV1beta1ClusterRoleUnauthorized {

	return &WatchRbacAuthorizationV1beta1ClusterRoleUnauthorized{}
}

// WriteResponse to the client
func (o *WatchRbacAuthorizationV1beta1ClusterRoleUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
