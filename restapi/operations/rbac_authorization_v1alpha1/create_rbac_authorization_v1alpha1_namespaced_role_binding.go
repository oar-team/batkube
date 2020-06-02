// Code generated by go-swagger; DO NOT EDIT.

package rbac_authorization_v1alpha1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// CreateRbacAuthorizationV1alpha1NamespacedRoleBindingHandlerFunc turns a function with the right signature into a create rbac authorization v1alpha1 namespaced role binding handler
type CreateRbacAuthorizationV1alpha1NamespacedRoleBindingHandlerFunc func(CreateRbacAuthorizationV1alpha1NamespacedRoleBindingParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateRbacAuthorizationV1alpha1NamespacedRoleBindingHandlerFunc) Handle(params CreateRbacAuthorizationV1alpha1NamespacedRoleBindingParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// CreateRbacAuthorizationV1alpha1NamespacedRoleBindingHandler interface for that can handle valid create rbac authorization v1alpha1 namespaced role binding params
type CreateRbacAuthorizationV1alpha1NamespacedRoleBindingHandler interface {
	Handle(CreateRbacAuthorizationV1alpha1NamespacedRoleBindingParams, interface{}) middleware.Responder
}

// NewCreateRbacAuthorizationV1alpha1NamespacedRoleBinding creates a new http.Handler for the create rbac authorization v1alpha1 namespaced role binding operation
func NewCreateRbacAuthorizationV1alpha1NamespacedRoleBinding(ctx *middleware.Context, handler CreateRbacAuthorizationV1alpha1NamespacedRoleBindingHandler) *CreateRbacAuthorizationV1alpha1NamespacedRoleBinding {
	return &CreateRbacAuthorizationV1alpha1NamespacedRoleBinding{Context: ctx, Handler: handler}
}

/*CreateRbacAuthorizationV1alpha1NamespacedRoleBinding swagger:route POST /apis/rbac.authorization.k8s.io/v1alpha1/namespaces/{namespace}/rolebindings rbacAuthorization_v1alpha1 createRbacAuthorizationV1alpha1NamespacedRoleBinding

create a RoleBinding

*/
type CreateRbacAuthorizationV1alpha1NamespacedRoleBinding struct {
	Context *middleware.Context
	Handler CreateRbacAuthorizationV1alpha1NamespacedRoleBindingHandler
}

func (o *CreateRbacAuthorizationV1alpha1NamespacedRoleBinding) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCreateRbacAuthorizationV1alpha1NamespacedRoleBindingParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
