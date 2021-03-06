// Code generated by go-swagger; DO NOT EDIT.

package rbac_authorization_v1alpha1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// DeleteRbacAuthorizationV1alpha1CollectionNamespacedRoleBindingOKCode is the HTTP code returned for type DeleteRbacAuthorizationV1alpha1CollectionNamespacedRoleBindingOK
const DeleteRbacAuthorizationV1alpha1CollectionNamespacedRoleBindingOKCode int = 200

/*DeleteRbacAuthorizationV1alpha1CollectionNamespacedRoleBindingOK OK

swagger:response deleteRbacAuthorizationV1alpha1CollectionNamespacedRoleBindingOK
*/
type DeleteRbacAuthorizationV1alpha1CollectionNamespacedRoleBindingOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1Status `json:"body,omitempty"`
}

// NewDeleteRbacAuthorizationV1alpha1CollectionNamespacedRoleBindingOK creates DeleteRbacAuthorizationV1alpha1CollectionNamespacedRoleBindingOK with default headers values
func NewDeleteRbacAuthorizationV1alpha1CollectionNamespacedRoleBindingOK() *DeleteRbacAuthorizationV1alpha1CollectionNamespacedRoleBindingOK {

	return &DeleteRbacAuthorizationV1alpha1CollectionNamespacedRoleBindingOK{}
}

// WithPayload adds the payload to the delete rbac authorization v1alpha1 collection namespaced role binding o k response
func (o *DeleteRbacAuthorizationV1alpha1CollectionNamespacedRoleBindingOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) *DeleteRbacAuthorizationV1alpha1CollectionNamespacedRoleBindingOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete rbac authorization v1alpha1 collection namespaced role binding o k response
func (o *DeleteRbacAuthorizationV1alpha1CollectionNamespacedRoleBindingOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteRbacAuthorizationV1alpha1CollectionNamespacedRoleBindingOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteRbacAuthorizationV1alpha1CollectionNamespacedRoleBindingUnauthorizedCode is the HTTP code returned for type DeleteRbacAuthorizationV1alpha1CollectionNamespacedRoleBindingUnauthorized
const DeleteRbacAuthorizationV1alpha1CollectionNamespacedRoleBindingUnauthorizedCode int = 401

/*DeleteRbacAuthorizationV1alpha1CollectionNamespacedRoleBindingUnauthorized Unauthorized

swagger:response deleteRbacAuthorizationV1alpha1CollectionNamespacedRoleBindingUnauthorized
*/
type DeleteRbacAuthorizationV1alpha1CollectionNamespacedRoleBindingUnauthorized struct {
}

// NewDeleteRbacAuthorizationV1alpha1CollectionNamespacedRoleBindingUnauthorized creates DeleteRbacAuthorizationV1alpha1CollectionNamespacedRoleBindingUnauthorized with default headers values
func NewDeleteRbacAuthorizationV1alpha1CollectionNamespacedRoleBindingUnauthorized() *DeleteRbacAuthorizationV1alpha1CollectionNamespacedRoleBindingUnauthorized {

	return &DeleteRbacAuthorizationV1alpha1CollectionNamespacedRoleBindingUnauthorized{}
}

// WriteResponse to the client
func (o *DeleteRbacAuthorizationV1alpha1CollectionNamespacedRoleBindingUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
