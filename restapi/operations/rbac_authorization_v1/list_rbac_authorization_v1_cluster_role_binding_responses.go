// Code generated by go-swagger; DO NOT EDIT.

package rbac_authorization_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/models"
)

// ListRbacAuthorizationV1ClusterRoleBindingOKCode is the HTTP code returned for type ListRbacAuthorizationV1ClusterRoleBindingOK
const ListRbacAuthorizationV1ClusterRoleBindingOKCode int = 200

/*ListRbacAuthorizationV1ClusterRoleBindingOK OK

swagger:response listRbacAuthorizationV1ClusterRoleBindingOK
*/
type ListRbacAuthorizationV1ClusterRoleBindingOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIRbacV1ClusterRoleBindingList `json:"body,omitempty"`
}

// NewListRbacAuthorizationV1ClusterRoleBindingOK creates ListRbacAuthorizationV1ClusterRoleBindingOK with default headers values
func NewListRbacAuthorizationV1ClusterRoleBindingOK() *ListRbacAuthorizationV1ClusterRoleBindingOK {

	return &ListRbacAuthorizationV1ClusterRoleBindingOK{}
}

// WithPayload adds the payload to the list rbac authorization v1 cluster role binding o k response
func (o *ListRbacAuthorizationV1ClusterRoleBindingOK) WithPayload(payload *models.IoK8sAPIRbacV1ClusterRoleBindingList) *ListRbacAuthorizationV1ClusterRoleBindingOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list rbac authorization v1 cluster role binding o k response
func (o *ListRbacAuthorizationV1ClusterRoleBindingOK) SetPayload(payload *models.IoK8sAPIRbacV1ClusterRoleBindingList) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListRbacAuthorizationV1ClusterRoleBindingOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListRbacAuthorizationV1ClusterRoleBindingUnauthorizedCode is the HTTP code returned for type ListRbacAuthorizationV1ClusterRoleBindingUnauthorized
const ListRbacAuthorizationV1ClusterRoleBindingUnauthorizedCode int = 401

/*ListRbacAuthorizationV1ClusterRoleBindingUnauthorized Unauthorized

swagger:response listRbacAuthorizationV1ClusterRoleBindingUnauthorized
*/
type ListRbacAuthorizationV1ClusterRoleBindingUnauthorized struct {
}

// NewListRbacAuthorizationV1ClusterRoleBindingUnauthorized creates ListRbacAuthorizationV1ClusterRoleBindingUnauthorized with default headers values
func NewListRbacAuthorizationV1ClusterRoleBindingUnauthorized() *ListRbacAuthorizationV1ClusterRoleBindingUnauthorized {

	return &ListRbacAuthorizationV1ClusterRoleBindingUnauthorized{}
}

// WriteResponse to the client
func (o *ListRbacAuthorizationV1ClusterRoleBindingUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
