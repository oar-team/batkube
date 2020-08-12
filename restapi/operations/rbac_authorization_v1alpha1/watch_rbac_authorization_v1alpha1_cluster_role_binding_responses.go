// Code generated by go-swagger; DO NOT EDIT.

package rbac_authorization_v1alpha1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// WatchRbacAuthorizationV1alpha1ClusterRoleBindingOKCode is the HTTP code returned for type WatchRbacAuthorizationV1alpha1ClusterRoleBindingOK
const WatchRbacAuthorizationV1alpha1ClusterRoleBindingOKCode int = 200

/*WatchRbacAuthorizationV1alpha1ClusterRoleBindingOK OK

swagger:response watchRbacAuthorizationV1alpha1ClusterRoleBindingOK
*/
type WatchRbacAuthorizationV1alpha1ClusterRoleBindingOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent `json:"body,omitempty"`
}

// NewWatchRbacAuthorizationV1alpha1ClusterRoleBindingOK creates WatchRbacAuthorizationV1alpha1ClusterRoleBindingOK with default headers values
func NewWatchRbacAuthorizationV1alpha1ClusterRoleBindingOK() *WatchRbacAuthorizationV1alpha1ClusterRoleBindingOK {

	return &WatchRbacAuthorizationV1alpha1ClusterRoleBindingOK{}
}

// WithPayload adds the payload to the watch rbac authorization v1alpha1 cluster role binding o k response
func (o *WatchRbacAuthorizationV1alpha1ClusterRoleBindingOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) *WatchRbacAuthorizationV1alpha1ClusterRoleBindingOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the watch rbac authorization v1alpha1 cluster role binding o k response
func (o *WatchRbacAuthorizationV1alpha1ClusterRoleBindingOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WatchRbacAuthorizationV1alpha1ClusterRoleBindingOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WatchRbacAuthorizationV1alpha1ClusterRoleBindingUnauthorizedCode is the HTTP code returned for type WatchRbacAuthorizationV1alpha1ClusterRoleBindingUnauthorized
const WatchRbacAuthorizationV1alpha1ClusterRoleBindingUnauthorizedCode int = 401

/*WatchRbacAuthorizationV1alpha1ClusterRoleBindingUnauthorized Unauthorized

swagger:response watchRbacAuthorizationV1alpha1ClusterRoleBindingUnauthorized
*/
type WatchRbacAuthorizationV1alpha1ClusterRoleBindingUnauthorized struct {
}

// NewWatchRbacAuthorizationV1alpha1ClusterRoleBindingUnauthorized creates WatchRbacAuthorizationV1alpha1ClusterRoleBindingUnauthorized with default headers values
func NewWatchRbacAuthorizationV1alpha1ClusterRoleBindingUnauthorized() *WatchRbacAuthorizationV1alpha1ClusterRoleBindingUnauthorized {

	return &WatchRbacAuthorizationV1alpha1ClusterRoleBindingUnauthorized{}
}

// WriteResponse to the client
func (o *WatchRbacAuthorizationV1alpha1ClusterRoleBindingUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
