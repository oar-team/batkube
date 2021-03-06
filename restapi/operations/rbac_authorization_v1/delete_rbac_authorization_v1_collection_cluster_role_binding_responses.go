// Code generated by go-swagger; DO NOT EDIT.

package rbac_authorization_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// DeleteRbacAuthorizationV1CollectionClusterRoleBindingOKCode is the HTTP code returned for type DeleteRbacAuthorizationV1CollectionClusterRoleBindingOK
const DeleteRbacAuthorizationV1CollectionClusterRoleBindingOKCode int = 200

/*DeleteRbacAuthorizationV1CollectionClusterRoleBindingOK OK

swagger:response deleteRbacAuthorizationV1CollectionClusterRoleBindingOK
*/
type DeleteRbacAuthorizationV1CollectionClusterRoleBindingOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1Status `json:"body,omitempty"`
}

// NewDeleteRbacAuthorizationV1CollectionClusterRoleBindingOK creates DeleteRbacAuthorizationV1CollectionClusterRoleBindingOK with default headers values
func NewDeleteRbacAuthorizationV1CollectionClusterRoleBindingOK() *DeleteRbacAuthorizationV1CollectionClusterRoleBindingOK {

	return &DeleteRbacAuthorizationV1CollectionClusterRoleBindingOK{}
}

// WithPayload adds the payload to the delete rbac authorization v1 collection cluster role binding o k response
func (o *DeleteRbacAuthorizationV1CollectionClusterRoleBindingOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) *DeleteRbacAuthorizationV1CollectionClusterRoleBindingOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete rbac authorization v1 collection cluster role binding o k response
func (o *DeleteRbacAuthorizationV1CollectionClusterRoleBindingOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteRbacAuthorizationV1CollectionClusterRoleBindingOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteRbacAuthorizationV1CollectionClusterRoleBindingUnauthorizedCode is the HTTP code returned for type DeleteRbacAuthorizationV1CollectionClusterRoleBindingUnauthorized
const DeleteRbacAuthorizationV1CollectionClusterRoleBindingUnauthorizedCode int = 401

/*DeleteRbacAuthorizationV1CollectionClusterRoleBindingUnauthorized Unauthorized

swagger:response deleteRbacAuthorizationV1CollectionClusterRoleBindingUnauthorized
*/
type DeleteRbacAuthorizationV1CollectionClusterRoleBindingUnauthorized struct {
}

// NewDeleteRbacAuthorizationV1CollectionClusterRoleBindingUnauthorized creates DeleteRbacAuthorizationV1CollectionClusterRoleBindingUnauthorized with default headers values
func NewDeleteRbacAuthorizationV1CollectionClusterRoleBindingUnauthorized() *DeleteRbacAuthorizationV1CollectionClusterRoleBindingUnauthorized {

	return &DeleteRbacAuthorizationV1CollectionClusterRoleBindingUnauthorized{}
}

// WriteResponse to the client
func (o *DeleteRbacAuthorizationV1CollectionClusterRoleBindingUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
