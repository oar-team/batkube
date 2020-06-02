// Code generated by go-swagger; DO NOT EDIT.

package rbac_authorization_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteRbacAuthorizationV1beta1CollectionNamespacedRoleHandlerFunc turns a function with the right signature into a delete rbac authorization v1beta1 collection namespaced role handler
type DeleteRbacAuthorizationV1beta1CollectionNamespacedRoleHandlerFunc func(DeleteRbacAuthorizationV1beta1CollectionNamespacedRoleParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteRbacAuthorizationV1beta1CollectionNamespacedRoleHandlerFunc) Handle(params DeleteRbacAuthorizationV1beta1CollectionNamespacedRoleParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// DeleteRbacAuthorizationV1beta1CollectionNamespacedRoleHandler interface for that can handle valid delete rbac authorization v1beta1 collection namespaced role params
type DeleteRbacAuthorizationV1beta1CollectionNamespacedRoleHandler interface {
	Handle(DeleteRbacAuthorizationV1beta1CollectionNamespacedRoleParams, interface{}) middleware.Responder
}

// NewDeleteRbacAuthorizationV1beta1CollectionNamespacedRole creates a new http.Handler for the delete rbac authorization v1beta1 collection namespaced role operation
func NewDeleteRbacAuthorizationV1beta1CollectionNamespacedRole(ctx *middleware.Context, handler DeleteRbacAuthorizationV1beta1CollectionNamespacedRoleHandler) *DeleteRbacAuthorizationV1beta1CollectionNamespacedRole {
	return &DeleteRbacAuthorizationV1beta1CollectionNamespacedRole{Context: ctx, Handler: handler}
}

/*DeleteRbacAuthorizationV1beta1CollectionNamespacedRole swagger:route DELETE /apis/rbac.authorization.k8s.io/v1beta1/namespaces/{namespace}/roles rbacAuthorization_v1beta1 deleteRbacAuthorizationV1beta1CollectionNamespacedRole

delete collection of Role

*/
type DeleteRbacAuthorizationV1beta1CollectionNamespacedRole struct {
	Context *middleware.Context
	Handler DeleteRbacAuthorizationV1beta1CollectionNamespacedRoleHandler
}

func (o *DeleteRbacAuthorizationV1beta1CollectionNamespacedRole) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteRbacAuthorizationV1beta1CollectionNamespacedRoleParams()

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
