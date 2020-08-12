// Code generated by go-swagger; DO NOT EDIT.

package rbac_authorization_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// WatchRbacAuthorizationV1beta1RoleBindingListForAllNamespacesOKCode is the HTTP code returned for type WatchRbacAuthorizationV1beta1RoleBindingListForAllNamespacesOK
const WatchRbacAuthorizationV1beta1RoleBindingListForAllNamespacesOKCode int = 200

/*WatchRbacAuthorizationV1beta1RoleBindingListForAllNamespacesOK OK

swagger:response watchRbacAuthorizationV1beta1RoleBindingListForAllNamespacesOK
*/
type WatchRbacAuthorizationV1beta1RoleBindingListForAllNamespacesOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent `json:"body,omitempty"`
}

// NewWatchRbacAuthorizationV1beta1RoleBindingListForAllNamespacesOK creates WatchRbacAuthorizationV1beta1RoleBindingListForAllNamespacesOK with default headers values
func NewWatchRbacAuthorizationV1beta1RoleBindingListForAllNamespacesOK() *WatchRbacAuthorizationV1beta1RoleBindingListForAllNamespacesOK {

	return &WatchRbacAuthorizationV1beta1RoleBindingListForAllNamespacesOK{}
}

// WithPayload adds the payload to the watch rbac authorization v1beta1 role binding list for all namespaces o k response
func (o *WatchRbacAuthorizationV1beta1RoleBindingListForAllNamespacesOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) *WatchRbacAuthorizationV1beta1RoleBindingListForAllNamespacesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the watch rbac authorization v1beta1 role binding list for all namespaces o k response
func (o *WatchRbacAuthorizationV1beta1RoleBindingListForAllNamespacesOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WatchRbacAuthorizationV1beta1RoleBindingListForAllNamespacesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WatchRbacAuthorizationV1beta1RoleBindingListForAllNamespacesUnauthorizedCode is the HTTP code returned for type WatchRbacAuthorizationV1beta1RoleBindingListForAllNamespacesUnauthorized
const WatchRbacAuthorizationV1beta1RoleBindingListForAllNamespacesUnauthorizedCode int = 401

/*WatchRbacAuthorizationV1beta1RoleBindingListForAllNamespacesUnauthorized Unauthorized

swagger:response watchRbacAuthorizationV1beta1RoleBindingListForAllNamespacesUnauthorized
*/
type WatchRbacAuthorizationV1beta1RoleBindingListForAllNamespacesUnauthorized struct {
}

// NewWatchRbacAuthorizationV1beta1RoleBindingListForAllNamespacesUnauthorized creates WatchRbacAuthorizationV1beta1RoleBindingListForAllNamespacesUnauthorized with default headers values
func NewWatchRbacAuthorizationV1beta1RoleBindingListForAllNamespacesUnauthorized() *WatchRbacAuthorizationV1beta1RoleBindingListForAllNamespacesUnauthorized {

	return &WatchRbacAuthorizationV1beta1RoleBindingListForAllNamespacesUnauthorized{}
}

// WriteResponse to the client
func (o *WatchRbacAuthorizationV1beta1RoleBindingListForAllNamespacesUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
