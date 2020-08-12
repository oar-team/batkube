// Code generated by go-swagger; DO NOT EDIT.

package rbac_authorization_v1alpha1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oar-team/batkube/models"
)

// ReplaceRbacAuthorizationV1alpha1ClusterRoleBindingOKCode is the HTTP code returned for type ReplaceRbacAuthorizationV1alpha1ClusterRoleBindingOK
const ReplaceRbacAuthorizationV1alpha1ClusterRoleBindingOKCode int = 200

/*ReplaceRbacAuthorizationV1alpha1ClusterRoleBindingOK OK

swagger:response replaceRbacAuthorizationV1alpha1ClusterRoleBindingOK
*/
type ReplaceRbacAuthorizationV1alpha1ClusterRoleBindingOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIRbacV1alpha1ClusterRoleBinding `json:"body,omitempty"`
}

// NewReplaceRbacAuthorizationV1alpha1ClusterRoleBindingOK creates ReplaceRbacAuthorizationV1alpha1ClusterRoleBindingOK with default headers values
func NewReplaceRbacAuthorizationV1alpha1ClusterRoleBindingOK() *ReplaceRbacAuthorizationV1alpha1ClusterRoleBindingOK {

	return &ReplaceRbacAuthorizationV1alpha1ClusterRoleBindingOK{}
}

// WithPayload adds the payload to the replace rbac authorization v1alpha1 cluster role binding o k response
func (o *ReplaceRbacAuthorizationV1alpha1ClusterRoleBindingOK) WithPayload(payload *models.IoK8sAPIRbacV1alpha1ClusterRoleBinding) *ReplaceRbacAuthorizationV1alpha1ClusterRoleBindingOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace rbac authorization v1alpha1 cluster role binding o k response
func (o *ReplaceRbacAuthorizationV1alpha1ClusterRoleBindingOK) SetPayload(payload *models.IoK8sAPIRbacV1alpha1ClusterRoleBinding) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceRbacAuthorizationV1alpha1ClusterRoleBindingOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceRbacAuthorizationV1alpha1ClusterRoleBindingCreatedCode is the HTTP code returned for type ReplaceRbacAuthorizationV1alpha1ClusterRoleBindingCreated
const ReplaceRbacAuthorizationV1alpha1ClusterRoleBindingCreatedCode int = 201

/*ReplaceRbacAuthorizationV1alpha1ClusterRoleBindingCreated Created

swagger:response replaceRbacAuthorizationV1alpha1ClusterRoleBindingCreated
*/
type ReplaceRbacAuthorizationV1alpha1ClusterRoleBindingCreated struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIRbacV1alpha1ClusterRoleBinding `json:"body,omitempty"`
}

// NewReplaceRbacAuthorizationV1alpha1ClusterRoleBindingCreated creates ReplaceRbacAuthorizationV1alpha1ClusterRoleBindingCreated with default headers values
func NewReplaceRbacAuthorizationV1alpha1ClusterRoleBindingCreated() *ReplaceRbacAuthorizationV1alpha1ClusterRoleBindingCreated {

	return &ReplaceRbacAuthorizationV1alpha1ClusterRoleBindingCreated{}
}

// WithPayload adds the payload to the replace rbac authorization v1alpha1 cluster role binding created response
func (o *ReplaceRbacAuthorizationV1alpha1ClusterRoleBindingCreated) WithPayload(payload *models.IoK8sAPIRbacV1alpha1ClusterRoleBinding) *ReplaceRbacAuthorizationV1alpha1ClusterRoleBindingCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace rbac authorization v1alpha1 cluster role binding created response
func (o *ReplaceRbacAuthorizationV1alpha1ClusterRoleBindingCreated) SetPayload(payload *models.IoK8sAPIRbacV1alpha1ClusterRoleBinding) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceRbacAuthorizationV1alpha1ClusterRoleBindingCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceRbacAuthorizationV1alpha1ClusterRoleBindingUnauthorizedCode is the HTTP code returned for type ReplaceRbacAuthorizationV1alpha1ClusterRoleBindingUnauthorized
const ReplaceRbacAuthorizationV1alpha1ClusterRoleBindingUnauthorizedCode int = 401

/*ReplaceRbacAuthorizationV1alpha1ClusterRoleBindingUnauthorized Unauthorized

swagger:response replaceRbacAuthorizationV1alpha1ClusterRoleBindingUnauthorized
*/
type ReplaceRbacAuthorizationV1alpha1ClusterRoleBindingUnauthorized struct {
}

// NewReplaceRbacAuthorizationV1alpha1ClusterRoleBindingUnauthorized creates ReplaceRbacAuthorizationV1alpha1ClusterRoleBindingUnauthorized with default headers values
func NewReplaceRbacAuthorizationV1alpha1ClusterRoleBindingUnauthorized() *ReplaceRbacAuthorizationV1alpha1ClusterRoleBindingUnauthorized {

	return &ReplaceRbacAuthorizationV1alpha1ClusterRoleBindingUnauthorized{}
}

// WriteResponse to the client
func (o *ReplaceRbacAuthorizationV1alpha1ClusterRoleBindingUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
