// Code generated by go-swagger; DO NOT EDIT.

package rbac_authorization_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// WatchRbacAuthorizationV1beta1ClusterRoleBindingListOKCode is the HTTP code returned for type WatchRbacAuthorizationV1beta1ClusterRoleBindingListOK
const WatchRbacAuthorizationV1beta1ClusterRoleBindingListOKCode int = 200

/*WatchRbacAuthorizationV1beta1ClusterRoleBindingListOK OK

swagger:response watchRbacAuthorizationV1beta1ClusterRoleBindingListOK
*/
type WatchRbacAuthorizationV1beta1ClusterRoleBindingListOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent `json:"body,omitempty"`
}

// NewWatchRbacAuthorizationV1beta1ClusterRoleBindingListOK creates WatchRbacAuthorizationV1beta1ClusterRoleBindingListOK with default headers values
func NewWatchRbacAuthorizationV1beta1ClusterRoleBindingListOK() *WatchRbacAuthorizationV1beta1ClusterRoleBindingListOK {

	return &WatchRbacAuthorizationV1beta1ClusterRoleBindingListOK{}
}

// WithPayload adds the payload to the watch rbac authorization v1beta1 cluster role binding list o k response
func (o *WatchRbacAuthorizationV1beta1ClusterRoleBindingListOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) *WatchRbacAuthorizationV1beta1ClusterRoleBindingListOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the watch rbac authorization v1beta1 cluster role binding list o k response
func (o *WatchRbacAuthorizationV1beta1ClusterRoleBindingListOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WatchRbacAuthorizationV1beta1ClusterRoleBindingListOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WatchRbacAuthorizationV1beta1ClusterRoleBindingListUnauthorizedCode is the HTTP code returned for type WatchRbacAuthorizationV1beta1ClusterRoleBindingListUnauthorized
const WatchRbacAuthorizationV1beta1ClusterRoleBindingListUnauthorizedCode int = 401

/*WatchRbacAuthorizationV1beta1ClusterRoleBindingListUnauthorized Unauthorized

swagger:response watchRbacAuthorizationV1beta1ClusterRoleBindingListUnauthorized
*/
type WatchRbacAuthorizationV1beta1ClusterRoleBindingListUnauthorized struct {
}

// NewWatchRbacAuthorizationV1beta1ClusterRoleBindingListUnauthorized creates WatchRbacAuthorizationV1beta1ClusterRoleBindingListUnauthorized with default headers values
func NewWatchRbacAuthorizationV1beta1ClusterRoleBindingListUnauthorized() *WatchRbacAuthorizationV1beta1ClusterRoleBindingListUnauthorized {

	return &WatchRbacAuthorizationV1beta1ClusterRoleBindingListUnauthorized{}
}

// WriteResponse to the client
func (o *WatchRbacAuthorizationV1beta1ClusterRoleBindingListUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
