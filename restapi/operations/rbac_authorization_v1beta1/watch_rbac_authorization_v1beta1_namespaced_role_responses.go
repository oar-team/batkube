// Code generated by go-swagger; DO NOT EDIT.

package rbac_authorization_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// WatchRbacAuthorizationV1beta1NamespacedRoleOKCode is the HTTP code returned for type WatchRbacAuthorizationV1beta1NamespacedRoleOK
const WatchRbacAuthorizationV1beta1NamespacedRoleOKCode int = 200

/*WatchRbacAuthorizationV1beta1NamespacedRoleOK OK

swagger:response watchRbacAuthorizationV1beta1NamespacedRoleOK
*/
type WatchRbacAuthorizationV1beta1NamespacedRoleOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent `json:"body,omitempty"`
}

// NewWatchRbacAuthorizationV1beta1NamespacedRoleOK creates WatchRbacAuthorizationV1beta1NamespacedRoleOK with default headers values
func NewWatchRbacAuthorizationV1beta1NamespacedRoleOK() *WatchRbacAuthorizationV1beta1NamespacedRoleOK {

	return &WatchRbacAuthorizationV1beta1NamespacedRoleOK{}
}

// WithPayload adds the payload to the watch rbac authorization v1beta1 namespaced role o k response
func (o *WatchRbacAuthorizationV1beta1NamespacedRoleOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) *WatchRbacAuthorizationV1beta1NamespacedRoleOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the watch rbac authorization v1beta1 namespaced role o k response
func (o *WatchRbacAuthorizationV1beta1NamespacedRoleOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WatchRbacAuthorizationV1beta1NamespacedRoleOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WatchRbacAuthorizationV1beta1NamespacedRoleUnauthorizedCode is the HTTP code returned for type WatchRbacAuthorizationV1beta1NamespacedRoleUnauthorized
const WatchRbacAuthorizationV1beta1NamespacedRoleUnauthorizedCode int = 401

/*WatchRbacAuthorizationV1beta1NamespacedRoleUnauthorized Unauthorized

swagger:response watchRbacAuthorizationV1beta1NamespacedRoleUnauthorized
*/
type WatchRbacAuthorizationV1beta1NamespacedRoleUnauthorized struct {
}

// NewWatchRbacAuthorizationV1beta1NamespacedRoleUnauthorized creates WatchRbacAuthorizationV1beta1NamespacedRoleUnauthorized with default headers values
func NewWatchRbacAuthorizationV1beta1NamespacedRoleUnauthorized() *WatchRbacAuthorizationV1beta1NamespacedRoleUnauthorized {

	return &WatchRbacAuthorizationV1beta1NamespacedRoleUnauthorized{}
}

// WriteResponse to the client
func (o *WatchRbacAuthorizationV1beta1NamespacedRoleUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
