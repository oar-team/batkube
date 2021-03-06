// Code generated by go-swagger; DO NOT EDIT.

package rbac_authorization_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// PatchRbacAuthorizationV1ClusterRoleOKCode is the HTTP code returned for type PatchRbacAuthorizationV1ClusterRoleOK
const PatchRbacAuthorizationV1ClusterRoleOKCode int = 200

/*PatchRbacAuthorizationV1ClusterRoleOK OK

swagger:response patchRbacAuthorizationV1ClusterRoleOK
*/
type PatchRbacAuthorizationV1ClusterRoleOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIRbacV1ClusterRole `json:"body,omitempty"`
}

// NewPatchRbacAuthorizationV1ClusterRoleOK creates PatchRbacAuthorizationV1ClusterRoleOK with default headers values
func NewPatchRbacAuthorizationV1ClusterRoleOK() *PatchRbacAuthorizationV1ClusterRoleOK {

	return &PatchRbacAuthorizationV1ClusterRoleOK{}
}

// WithPayload adds the payload to the patch rbac authorization v1 cluster role o k response
func (o *PatchRbacAuthorizationV1ClusterRoleOK) WithPayload(payload *models.IoK8sAPIRbacV1ClusterRole) *PatchRbacAuthorizationV1ClusterRoleOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch rbac authorization v1 cluster role o k response
func (o *PatchRbacAuthorizationV1ClusterRoleOK) SetPayload(payload *models.IoK8sAPIRbacV1ClusterRole) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchRbacAuthorizationV1ClusterRoleOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchRbacAuthorizationV1ClusterRoleUnauthorizedCode is the HTTP code returned for type PatchRbacAuthorizationV1ClusterRoleUnauthorized
const PatchRbacAuthorizationV1ClusterRoleUnauthorizedCode int = 401

/*PatchRbacAuthorizationV1ClusterRoleUnauthorized Unauthorized

swagger:response patchRbacAuthorizationV1ClusterRoleUnauthorized
*/
type PatchRbacAuthorizationV1ClusterRoleUnauthorized struct {
}

// NewPatchRbacAuthorizationV1ClusterRoleUnauthorized creates PatchRbacAuthorizationV1ClusterRoleUnauthorized with default headers values
func NewPatchRbacAuthorizationV1ClusterRoleUnauthorized() *PatchRbacAuthorizationV1ClusterRoleUnauthorized {

	return &PatchRbacAuthorizationV1ClusterRoleUnauthorized{}
}

// WriteResponse to the client
func (o *PatchRbacAuthorizationV1ClusterRoleUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
