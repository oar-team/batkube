// Code generated by go-swagger; DO NOT EDIT.

package rbac_authorization_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ReplaceRbacAuthorizationV1NamespacedRoleHandlerFunc turns a function with the right signature into a replace rbac authorization v1 namespaced role handler
type ReplaceRbacAuthorizationV1NamespacedRoleHandlerFunc func(ReplaceRbacAuthorizationV1NamespacedRoleParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn ReplaceRbacAuthorizationV1NamespacedRoleHandlerFunc) Handle(params ReplaceRbacAuthorizationV1NamespacedRoleParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// ReplaceRbacAuthorizationV1NamespacedRoleHandler interface for that can handle valid replace rbac authorization v1 namespaced role params
type ReplaceRbacAuthorizationV1NamespacedRoleHandler interface {
	Handle(ReplaceRbacAuthorizationV1NamespacedRoleParams, interface{}) middleware.Responder
}

// NewReplaceRbacAuthorizationV1NamespacedRole creates a new http.Handler for the replace rbac authorization v1 namespaced role operation
func NewReplaceRbacAuthorizationV1NamespacedRole(ctx *middleware.Context, handler ReplaceRbacAuthorizationV1NamespacedRoleHandler) *ReplaceRbacAuthorizationV1NamespacedRole {
	return &ReplaceRbacAuthorizationV1NamespacedRole{Context: ctx, Handler: handler}
}

/*ReplaceRbacAuthorizationV1NamespacedRole swagger:route PUT /apis/rbac.authorization.k8s.io/v1/namespaces/{namespace}/roles/{name} rbacAuthorization_v1 replaceRbacAuthorizationV1NamespacedRole

replace the specified Role

*/
type ReplaceRbacAuthorizationV1NamespacedRole struct {
	Context *middleware.Context
	Handler ReplaceRbacAuthorizationV1NamespacedRoleHandler
}

func (o *ReplaceRbacAuthorizationV1NamespacedRole) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewReplaceRbacAuthorizationV1NamespacedRoleParams()

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
