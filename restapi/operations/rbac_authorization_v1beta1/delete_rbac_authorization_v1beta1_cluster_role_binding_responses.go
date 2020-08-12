// Code generated by go-swagger; DO NOT EDIT.

package rbac_authorization_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// DeleteRbacAuthorizationV1beta1ClusterRoleBindingOKCode is the HTTP code returned for type DeleteRbacAuthorizationV1beta1ClusterRoleBindingOK
const DeleteRbacAuthorizationV1beta1ClusterRoleBindingOKCode int = 200

/*DeleteRbacAuthorizationV1beta1ClusterRoleBindingOK OK

swagger:response deleteRbacAuthorizationV1beta1ClusterRoleBindingOK
*/
type DeleteRbacAuthorizationV1beta1ClusterRoleBindingOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1Status `json:"body,omitempty"`
}

// NewDeleteRbacAuthorizationV1beta1ClusterRoleBindingOK creates DeleteRbacAuthorizationV1beta1ClusterRoleBindingOK with default headers values
func NewDeleteRbacAuthorizationV1beta1ClusterRoleBindingOK() *DeleteRbacAuthorizationV1beta1ClusterRoleBindingOK {

	return &DeleteRbacAuthorizationV1beta1ClusterRoleBindingOK{}
}

// WithPayload adds the payload to the delete rbac authorization v1beta1 cluster role binding o k response
func (o *DeleteRbacAuthorizationV1beta1ClusterRoleBindingOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) *DeleteRbacAuthorizationV1beta1ClusterRoleBindingOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete rbac authorization v1beta1 cluster role binding o k response
func (o *DeleteRbacAuthorizationV1beta1ClusterRoleBindingOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteRbacAuthorizationV1beta1ClusterRoleBindingOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteRbacAuthorizationV1beta1ClusterRoleBindingAcceptedCode is the HTTP code returned for type DeleteRbacAuthorizationV1beta1ClusterRoleBindingAccepted
const DeleteRbacAuthorizationV1beta1ClusterRoleBindingAcceptedCode int = 202

/*DeleteRbacAuthorizationV1beta1ClusterRoleBindingAccepted Accepted

swagger:response deleteRbacAuthorizationV1beta1ClusterRoleBindingAccepted
*/
type DeleteRbacAuthorizationV1beta1ClusterRoleBindingAccepted struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1Status `json:"body,omitempty"`
}

// NewDeleteRbacAuthorizationV1beta1ClusterRoleBindingAccepted creates DeleteRbacAuthorizationV1beta1ClusterRoleBindingAccepted with default headers values
func NewDeleteRbacAuthorizationV1beta1ClusterRoleBindingAccepted() *DeleteRbacAuthorizationV1beta1ClusterRoleBindingAccepted {

	return &DeleteRbacAuthorizationV1beta1ClusterRoleBindingAccepted{}
}

// WithPayload adds the payload to the delete rbac authorization v1beta1 cluster role binding accepted response
func (o *DeleteRbacAuthorizationV1beta1ClusterRoleBindingAccepted) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) *DeleteRbacAuthorizationV1beta1ClusterRoleBindingAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete rbac authorization v1beta1 cluster role binding accepted response
func (o *DeleteRbacAuthorizationV1beta1ClusterRoleBindingAccepted) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteRbacAuthorizationV1beta1ClusterRoleBindingAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(202)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteRbacAuthorizationV1beta1ClusterRoleBindingUnauthorizedCode is the HTTP code returned for type DeleteRbacAuthorizationV1beta1ClusterRoleBindingUnauthorized
const DeleteRbacAuthorizationV1beta1ClusterRoleBindingUnauthorizedCode int = 401

/*DeleteRbacAuthorizationV1beta1ClusterRoleBindingUnauthorized Unauthorized

swagger:response deleteRbacAuthorizationV1beta1ClusterRoleBindingUnauthorized
*/
type DeleteRbacAuthorizationV1beta1ClusterRoleBindingUnauthorized struct {
}

// NewDeleteRbacAuthorizationV1beta1ClusterRoleBindingUnauthorized creates DeleteRbacAuthorizationV1beta1ClusterRoleBindingUnauthorized with default headers values
func NewDeleteRbacAuthorizationV1beta1ClusterRoleBindingUnauthorized() *DeleteRbacAuthorizationV1beta1ClusterRoleBindingUnauthorized {

	return &DeleteRbacAuthorizationV1beta1ClusterRoleBindingUnauthorized{}
}

// WriteResponse to the client
func (o *DeleteRbacAuthorizationV1beta1ClusterRoleBindingUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
