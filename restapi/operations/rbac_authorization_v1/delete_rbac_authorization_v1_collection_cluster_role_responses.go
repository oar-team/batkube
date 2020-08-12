// Code generated by go-swagger; DO NOT EDIT.

package rbac_authorization_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// DeleteRbacAuthorizationV1CollectionClusterRoleOKCode is the HTTP code returned for type DeleteRbacAuthorizationV1CollectionClusterRoleOK
const DeleteRbacAuthorizationV1CollectionClusterRoleOKCode int = 200

/*DeleteRbacAuthorizationV1CollectionClusterRoleOK OK

swagger:response deleteRbacAuthorizationV1CollectionClusterRoleOK
*/
type DeleteRbacAuthorizationV1CollectionClusterRoleOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1Status `json:"body,omitempty"`
}

// NewDeleteRbacAuthorizationV1CollectionClusterRoleOK creates DeleteRbacAuthorizationV1CollectionClusterRoleOK with default headers values
func NewDeleteRbacAuthorizationV1CollectionClusterRoleOK() *DeleteRbacAuthorizationV1CollectionClusterRoleOK {

	return &DeleteRbacAuthorizationV1CollectionClusterRoleOK{}
}

// WithPayload adds the payload to the delete rbac authorization v1 collection cluster role o k response
func (o *DeleteRbacAuthorizationV1CollectionClusterRoleOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) *DeleteRbacAuthorizationV1CollectionClusterRoleOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete rbac authorization v1 collection cluster role o k response
func (o *DeleteRbacAuthorizationV1CollectionClusterRoleOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteRbacAuthorizationV1CollectionClusterRoleOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteRbacAuthorizationV1CollectionClusterRoleUnauthorizedCode is the HTTP code returned for type DeleteRbacAuthorizationV1CollectionClusterRoleUnauthorized
const DeleteRbacAuthorizationV1CollectionClusterRoleUnauthorizedCode int = 401

/*DeleteRbacAuthorizationV1CollectionClusterRoleUnauthorized Unauthorized

swagger:response deleteRbacAuthorizationV1CollectionClusterRoleUnauthorized
*/
type DeleteRbacAuthorizationV1CollectionClusterRoleUnauthorized struct {
}

// NewDeleteRbacAuthorizationV1CollectionClusterRoleUnauthorized creates DeleteRbacAuthorizationV1CollectionClusterRoleUnauthorized with default headers values
func NewDeleteRbacAuthorizationV1CollectionClusterRoleUnauthorized() *DeleteRbacAuthorizationV1CollectionClusterRoleUnauthorized {

	return &DeleteRbacAuthorizationV1CollectionClusterRoleUnauthorized{}
}

// WriteResponse to the client
func (o *DeleteRbacAuthorizationV1CollectionClusterRoleUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
