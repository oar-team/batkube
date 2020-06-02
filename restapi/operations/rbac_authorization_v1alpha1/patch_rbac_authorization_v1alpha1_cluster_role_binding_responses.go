// Code generated by go-swagger; DO NOT EDIT.

package rbac_authorization_v1alpha1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/models"
)

// PatchRbacAuthorizationV1alpha1ClusterRoleBindingOKCode is the HTTP code returned for type PatchRbacAuthorizationV1alpha1ClusterRoleBindingOK
const PatchRbacAuthorizationV1alpha1ClusterRoleBindingOKCode int = 200

/*PatchRbacAuthorizationV1alpha1ClusterRoleBindingOK OK

swagger:response patchRbacAuthorizationV1alpha1ClusterRoleBindingOK
*/
type PatchRbacAuthorizationV1alpha1ClusterRoleBindingOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIRbacV1alpha1ClusterRoleBinding `json:"body,omitempty"`
}

// NewPatchRbacAuthorizationV1alpha1ClusterRoleBindingOK creates PatchRbacAuthorizationV1alpha1ClusterRoleBindingOK with default headers values
func NewPatchRbacAuthorizationV1alpha1ClusterRoleBindingOK() *PatchRbacAuthorizationV1alpha1ClusterRoleBindingOK {

	return &PatchRbacAuthorizationV1alpha1ClusterRoleBindingOK{}
}

// WithPayload adds the payload to the patch rbac authorization v1alpha1 cluster role binding o k response
func (o *PatchRbacAuthorizationV1alpha1ClusterRoleBindingOK) WithPayload(payload *models.IoK8sAPIRbacV1alpha1ClusterRoleBinding) *PatchRbacAuthorizationV1alpha1ClusterRoleBindingOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch rbac authorization v1alpha1 cluster role binding o k response
func (o *PatchRbacAuthorizationV1alpha1ClusterRoleBindingOK) SetPayload(payload *models.IoK8sAPIRbacV1alpha1ClusterRoleBinding) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchRbacAuthorizationV1alpha1ClusterRoleBindingOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchRbacAuthorizationV1alpha1ClusterRoleBindingUnauthorizedCode is the HTTP code returned for type PatchRbacAuthorizationV1alpha1ClusterRoleBindingUnauthorized
const PatchRbacAuthorizationV1alpha1ClusterRoleBindingUnauthorizedCode int = 401

/*PatchRbacAuthorizationV1alpha1ClusterRoleBindingUnauthorized Unauthorized

swagger:response patchRbacAuthorizationV1alpha1ClusterRoleBindingUnauthorized
*/
type PatchRbacAuthorizationV1alpha1ClusterRoleBindingUnauthorized struct {
}

// NewPatchRbacAuthorizationV1alpha1ClusterRoleBindingUnauthorized creates PatchRbacAuthorizationV1alpha1ClusterRoleBindingUnauthorized with default headers values
func NewPatchRbacAuthorizationV1alpha1ClusterRoleBindingUnauthorized() *PatchRbacAuthorizationV1alpha1ClusterRoleBindingUnauthorized {

	return &PatchRbacAuthorizationV1alpha1ClusterRoleBindingUnauthorized{}
}

// WriteResponse to the client
func (o *PatchRbacAuthorizationV1alpha1ClusterRoleBindingUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
