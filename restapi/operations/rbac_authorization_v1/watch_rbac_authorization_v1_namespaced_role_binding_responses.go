// Code generated by go-swagger; DO NOT EDIT.

package rbac_authorization_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/models"
)

// WatchRbacAuthorizationV1NamespacedRoleBindingOKCode is the HTTP code returned for type WatchRbacAuthorizationV1NamespacedRoleBindingOK
const WatchRbacAuthorizationV1NamespacedRoleBindingOKCode int = 200

/*WatchRbacAuthorizationV1NamespacedRoleBindingOK OK

swagger:response watchRbacAuthorizationV1NamespacedRoleBindingOK
*/
type WatchRbacAuthorizationV1NamespacedRoleBindingOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent `json:"body,omitempty"`
}

// NewWatchRbacAuthorizationV1NamespacedRoleBindingOK creates WatchRbacAuthorizationV1NamespacedRoleBindingOK with default headers values
func NewWatchRbacAuthorizationV1NamespacedRoleBindingOK() *WatchRbacAuthorizationV1NamespacedRoleBindingOK {

	return &WatchRbacAuthorizationV1NamespacedRoleBindingOK{}
}

// WithPayload adds the payload to the watch rbac authorization v1 namespaced role binding o k response
func (o *WatchRbacAuthorizationV1NamespacedRoleBindingOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) *WatchRbacAuthorizationV1NamespacedRoleBindingOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the watch rbac authorization v1 namespaced role binding o k response
func (o *WatchRbacAuthorizationV1NamespacedRoleBindingOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WatchRbacAuthorizationV1NamespacedRoleBindingOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WatchRbacAuthorizationV1NamespacedRoleBindingUnauthorizedCode is the HTTP code returned for type WatchRbacAuthorizationV1NamespacedRoleBindingUnauthorized
const WatchRbacAuthorizationV1NamespacedRoleBindingUnauthorizedCode int = 401

/*WatchRbacAuthorizationV1NamespacedRoleBindingUnauthorized Unauthorized

swagger:response watchRbacAuthorizationV1NamespacedRoleBindingUnauthorized
*/
type WatchRbacAuthorizationV1NamespacedRoleBindingUnauthorized struct {
}

// NewWatchRbacAuthorizationV1NamespacedRoleBindingUnauthorized creates WatchRbacAuthorizationV1NamespacedRoleBindingUnauthorized with default headers values
func NewWatchRbacAuthorizationV1NamespacedRoleBindingUnauthorized() *WatchRbacAuthorizationV1NamespacedRoleBindingUnauthorized {

	return &WatchRbacAuthorizationV1NamespacedRoleBindingUnauthorized{}
}

// WriteResponse to the client
func (o *WatchRbacAuthorizationV1NamespacedRoleBindingUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
