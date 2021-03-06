// Code generated by go-swagger; DO NOT EDIT.

package rbac_authorization_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// DeleteRbacAuthorizationV1NamespacedRoleOKCode is the HTTP code returned for type DeleteRbacAuthorizationV1NamespacedRoleOK
const DeleteRbacAuthorizationV1NamespacedRoleOKCode int = 200

/*DeleteRbacAuthorizationV1NamespacedRoleOK OK

swagger:response deleteRbacAuthorizationV1NamespacedRoleOK
*/
type DeleteRbacAuthorizationV1NamespacedRoleOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1Status `json:"body,omitempty"`
}

// NewDeleteRbacAuthorizationV1NamespacedRoleOK creates DeleteRbacAuthorizationV1NamespacedRoleOK with default headers values
func NewDeleteRbacAuthorizationV1NamespacedRoleOK() *DeleteRbacAuthorizationV1NamespacedRoleOK {

	return &DeleteRbacAuthorizationV1NamespacedRoleOK{}
}

// WithPayload adds the payload to the delete rbac authorization v1 namespaced role o k response
func (o *DeleteRbacAuthorizationV1NamespacedRoleOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) *DeleteRbacAuthorizationV1NamespacedRoleOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete rbac authorization v1 namespaced role o k response
func (o *DeleteRbacAuthorizationV1NamespacedRoleOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteRbacAuthorizationV1NamespacedRoleOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteRbacAuthorizationV1NamespacedRoleAcceptedCode is the HTTP code returned for type DeleteRbacAuthorizationV1NamespacedRoleAccepted
const DeleteRbacAuthorizationV1NamespacedRoleAcceptedCode int = 202

/*DeleteRbacAuthorizationV1NamespacedRoleAccepted Accepted

swagger:response deleteRbacAuthorizationV1NamespacedRoleAccepted
*/
type DeleteRbacAuthorizationV1NamespacedRoleAccepted struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1Status `json:"body,omitempty"`
}

// NewDeleteRbacAuthorizationV1NamespacedRoleAccepted creates DeleteRbacAuthorizationV1NamespacedRoleAccepted with default headers values
func NewDeleteRbacAuthorizationV1NamespacedRoleAccepted() *DeleteRbacAuthorizationV1NamespacedRoleAccepted {

	return &DeleteRbacAuthorizationV1NamespacedRoleAccepted{}
}

// WithPayload adds the payload to the delete rbac authorization v1 namespaced role accepted response
func (o *DeleteRbacAuthorizationV1NamespacedRoleAccepted) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) *DeleteRbacAuthorizationV1NamespacedRoleAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete rbac authorization v1 namespaced role accepted response
func (o *DeleteRbacAuthorizationV1NamespacedRoleAccepted) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteRbacAuthorizationV1NamespacedRoleAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(202)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteRbacAuthorizationV1NamespacedRoleUnauthorizedCode is the HTTP code returned for type DeleteRbacAuthorizationV1NamespacedRoleUnauthorized
const DeleteRbacAuthorizationV1NamespacedRoleUnauthorizedCode int = 401

/*DeleteRbacAuthorizationV1NamespacedRoleUnauthorized Unauthorized

swagger:response deleteRbacAuthorizationV1NamespacedRoleUnauthorized
*/
type DeleteRbacAuthorizationV1NamespacedRoleUnauthorized struct {
}

// NewDeleteRbacAuthorizationV1NamespacedRoleUnauthorized creates DeleteRbacAuthorizationV1NamespacedRoleUnauthorized with default headers values
func NewDeleteRbacAuthorizationV1NamespacedRoleUnauthorized() *DeleteRbacAuthorizationV1NamespacedRoleUnauthorized {

	return &DeleteRbacAuthorizationV1NamespacedRoleUnauthorized{}
}

// WriteResponse to the client
func (o *DeleteRbacAuthorizationV1NamespacedRoleUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
