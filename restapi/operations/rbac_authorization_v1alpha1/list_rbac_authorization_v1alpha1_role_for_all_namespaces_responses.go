// Code generated by go-swagger; DO NOT EDIT.

package rbac_authorization_v1alpha1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// ListRbacAuthorizationV1alpha1RoleForAllNamespacesOKCode is the HTTP code returned for type ListRbacAuthorizationV1alpha1RoleForAllNamespacesOK
const ListRbacAuthorizationV1alpha1RoleForAllNamespacesOKCode int = 200

/*ListRbacAuthorizationV1alpha1RoleForAllNamespacesOK OK

swagger:response listRbacAuthorizationV1alpha1RoleForAllNamespacesOK
*/
type ListRbacAuthorizationV1alpha1RoleForAllNamespacesOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIRbacV1alpha1RoleList `json:"body,omitempty"`
}

// NewListRbacAuthorizationV1alpha1RoleForAllNamespacesOK creates ListRbacAuthorizationV1alpha1RoleForAllNamespacesOK with default headers values
func NewListRbacAuthorizationV1alpha1RoleForAllNamespacesOK() *ListRbacAuthorizationV1alpha1RoleForAllNamespacesOK {

	return &ListRbacAuthorizationV1alpha1RoleForAllNamespacesOK{}
}

// WithPayload adds the payload to the list rbac authorization v1alpha1 role for all namespaces o k response
func (o *ListRbacAuthorizationV1alpha1RoleForAllNamespacesOK) WithPayload(payload *models.IoK8sAPIRbacV1alpha1RoleList) *ListRbacAuthorizationV1alpha1RoleForAllNamespacesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list rbac authorization v1alpha1 role for all namespaces o k response
func (o *ListRbacAuthorizationV1alpha1RoleForAllNamespacesOK) SetPayload(payload *models.IoK8sAPIRbacV1alpha1RoleList) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListRbacAuthorizationV1alpha1RoleForAllNamespacesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListRbacAuthorizationV1alpha1RoleForAllNamespacesUnauthorizedCode is the HTTP code returned for type ListRbacAuthorizationV1alpha1RoleForAllNamespacesUnauthorized
const ListRbacAuthorizationV1alpha1RoleForAllNamespacesUnauthorizedCode int = 401

/*ListRbacAuthorizationV1alpha1RoleForAllNamespacesUnauthorized Unauthorized

swagger:response listRbacAuthorizationV1alpha1RoleForAllNamespacesUnauthorized
*/
type ListRbacAuthorizationV1alpha1RoleForAllNamespacesUnauthorized struct {
}

// NewListRbacAuthorizationV1alpha1RoleForAllNamespacesUnauthorized creates ListRbacAuthorizationV1alpha1RoleForAllNamespacesUnauthorized with default headers values
func NewListRbacAuthorizationV1alpha1RoleForAllNamespacesUnauthorized() *ListRbacAuthorizationV1alpha1RoleForAllNamespacesUnauthorized {

	return &ListRbacAuthorizationV1alpha1RoleForAllNamespacesUnauthorized{}
}

// WriteResponse to the client
func (o *ListRbacAuthorizationV1alpha1RoleForAllNamespacesUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
