// Code generated by go-swagger; DO NOT EDIT.

package rbac_authorization_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// WatchRbacAuthorizationV1beta1NamespacedRoleListOKCode is the HTTP code returned for type WatchRbacAuthorizationV1beta1NamespacedRoleListOK
const WatchRbacAuthorizationV1beta1NamespacedRoleListOKCode int = 200

/*WatchRbacAuthorizationV1beta1NamespacedRoleListOK OK

swagger:response watchRbacAuthorizationV1beta1NamespacedRoleListOK
*/
type WatchRbacAuthorizationV1beta1NamespacedRoleListOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent `json:"body,omitempty"`
}

// NewWatchRbacAuthorizationV1beta1NamespacedRoleListOK creates WatchRbacAuthorizationV1beta1NamespacedRoleListOK with default headers values
func NewWatchRbacAuthorizationV1beta1NamespacedRoleListOK() *WatchRbacAuthorizationV1beta1NamespacedRoleListOK {

	return &WatchRbacAuthorizationV1beta1NamespacedRoleListOK{}
}

// WithPayload adds the payload to the watch rbac authorization v1beta1 namespaced role list o k response
func (o *WatchRbacAuthorizationV1beta1NamespacedRoleListOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) *WatchRbacAuthorizationV1beta1NamespacedRoleListOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the watch rbac authorization v1beta1 namespaced role list o k response
func (o *WatchRbacAuthorizationV1beta1NamespacedRoleListOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WatchRbacAuthorizationV1beta1NamespacedRoleListOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WatchRbacAuthorizationV1beta1NamespacedRoleListUnauthorizedCode is the HTTP code returned for type WatchRbacAuthorizationV1beta1NamespacedRoleListUnauthorized
const WatchRbacAuthorizationV1beta1NamespacedRoleListUnauthorizedCode int = 401

/*WatchRbacAuthorizationV1beta1NamespacedRoleListUnauthorized Unauthorized

swagger:response watchRbacAuthorizationV1beta1NamespacedRoleListUnauthorized
*/
type WatchRbacAuthorizationV1beta1NamespacedRoleListUnauthorized struct {
}

// NewWatchRbacAuthorizationV1beta1NamespacedRoleListUnauthorized creates WatchRbacAuthorizationV1beta1NamespacedRoleListUnauthorized with default headers values
func NewWatchRbacAuthorizationV1beta1NamespacedRoleListUnauthorized() *WatchRbacAuthorizationV1beta1NamespacedRoleListUnauthorized {

	return &WatchRbacAuthorizationV1beta1NamespacedRoleListUnauthorized{}
}

// WriteResponse to the client
func (o *WatchRbacAuthorizationV1beta1NamespacedRoleListUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
