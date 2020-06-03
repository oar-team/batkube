// Code generated by go-swagger; DO NOT EDIT.

package rbac_authorization_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/models"
)

// PatchRbacAuthorizationV1beta1ClusterRoleBindingOKCode is the HTTP code returned for type PatchRbacAuthorizationV1beta1ClusterRoleBindingOK
const PatchRbacAuthorizationV1beta1ClusterRoleBindingOKCode int = 200

/*PatchRbacAuthorizationV1beta1ClusterRoleBindingOK OK

swagger:response patchRbacAuthorizationV1beta1ClusterRoleBindingOK
*/
type PatchRbacAuthorizationV1beta1ClusterRoleBindingOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIRbacV1beta1ClusterRoleBinding `json:"body,omitempty"`
}

// NewPatchRbacAuthorizationV1beta1ClusterRoleBindingOK creates PatchRbacAuthorizationV1beta1ClusterRoleBindingOK with default headers values
func NewPatchRbacAuthorizationV1beta1ClusterRoleBindingOK() *PatchRbacAuthorizationV1beta1ClusterRoleBindingOK {

	return &PatchRbacAuthorizationV1beta1ClusterRoleBindingOK{}
}

// WithPayload adds the payload to the patch rbac authorization v1beta1 cluster role binding o k response
func (o *PatchRbacAuthorizationV1beta1ClusterRoleBindingOK) WithPayload(payload *models.IoK8sAPIRbacV1beta1ClusterRoleBinding) *PatchRbacAuthorizationV1beta1ClusterRoleBindingOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch rbac authorization v1beta1 cluster role binding o k response
func (o *PatchRbacAuthorizationV1beta1ClusterRoleBindingOK) SetPayload(payload *models.IoK8sAPIRbacV1beta1ClusterRoleBinding) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchRbacAuthorizationV1beta1ClusterRoleBindingOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchRbacAuthorizationV1beta1ClusterRoleBindingUnauthorizedCode is the HTTP code returned for type PatchRbacAuthorizationV1beta1ClusterRoleBindingUnauthorized
const PatchRbacAuthorizationV1beta1ClusterRoleBindingUnauthorizedCode int = 401

/*PatchRbacAuthorizationV1beta1ClusterRoleBindingUnauthorized Unauthorized

swagger:response patchRbacAuthorizationV1beta1ClusterRoleBindingUnauthorized
*/
type PatchRbacAuthorizationV1beta1ClusterRoleBindingUnauthorized struct {
}

// NewPatchRbacAuthorizationV1beta1ClusterRoleBindingUnauthorized creates PatchRbacAuthorizationV1beta1ClusterRoleBindingUnauthorized with default headers values
func NewPatchRbacAuthorizationV1beta1ClusterRoleBindingUnauthorized() *PatchRbacAuthorizationV1beta1ClusterRoleBindingUnauthorized {

	return &PatchRbacAuthorizationV1beta1ClusterRoleBindingUnauthorized{}
}

// WriteResponse to the client
func (o *PatchRbacAuthorizationV1beta1ClusterRoleBindingUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}