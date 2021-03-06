// Code generated by go-swagger; DO NOT EDIT.

package rbac_authorization_v1alpha1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// ReadRbacAuthorizationV1alpha1NamespacedRoleOKCode is the HTTP code returned for type ReadRbacAuthorizationV1alpha1NamespacedRoleOK
const ReadRbacAuthorizationV1alpha1NamespacedRoleOKCode int = 200

/*ReadRbacAuthorizationV1alpha1NamespacedRoleOK OK

swagger:response readRbacAuthorizationV1alpha1NamespacedRoleOK
*/
type ReadRbacAuthorizationV1alpha1NamespacedRoleOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIRbacV1alpha1Role `json:"body,omitempty"`
}

// NewReadRbacAuthorizationV1alpha1NamespacedRoleOK creates ReadRbacAuthorizationV1alpha1NamespacedRoleOK with default headers values
func NewReadRbacAuthorizationV1alpha1NamespacedRoleOK() *ReadRbacAuthorizationV1alpha1NamespacedRoleOK {

	return &ReadRbacAuthorizationV1alpha1NamespacedRoleOK{}
}

// WithPayload adds the payload to the read rbac authorization v1alpha1 namespaced role o k response
func (o *ReadRbacAuthorizationV1alpha1NamespacedRoleOK) WithPayload(payload *models.IoK8sAPIRbacV1alpha1Role) *ReadRbacAuthorizationV1alpha1NamespacedRoleOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the read rbac authorization v1alpha1 namespaced role o k response
func (o *ReadRbacAuthorizationV1alpha1NamespacedRoleOK) SetPayload(payload *models.IoK8sAPIRbacV1alpha1Role) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReadRbacAuthorizationV1alpha1NamespacedRoleOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReadRbacAuthorizationV1alpha1NamespacedRoleUnauthorizedCode is the HTTP code returned for type ReadRbacAuthorizationV1alpha1NamespacedRoleUnauthorized
const ReadRbacAuthorizationV1alpha1NamespacedRoleUnauthorizedCode int = 401

/*ReadRbacAuthorizationV1alpha1NamespacedRoleUnauthorized Unauthorized

swagger:response readRbacAuthorizationV1alpha1NamespacedRoleUnauthorized
*/
type ReadRbacAuthorizationV1alpha1NamespacedRoleUnauthorized struct {
}

// NewReadRbacAuthorizationV1alpha1NamespacedRoleUnauthorized creates ReadRbacAuthorizationV1alpha1NamespacedRoleUnauthorized with default headers values
func NewReadRbacAuthorizationV1alpha1NamespacedRoleUnauthorized() *ReadRbacAuthorizationV1alpha1NamespacedRoleUnauthorized {

	return &ReadRbacAuthorizationV1alpha1NamespacedRoleUnauthorized{}
}

// WriteResponse to the client
func (o *ReadRbacAuthorizationV1alpha1NamespacedRoleUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
