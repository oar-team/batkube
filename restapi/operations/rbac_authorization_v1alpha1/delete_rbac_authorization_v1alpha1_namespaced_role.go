// Code generated by go-swagger; DO NOT EDIT.

package rbac_authorization_v1alpha1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteRbacAuthorizationV1alpha1NamespacedRoleHandlerFunc turns a function with the right signature into a delete rbac authorization v1alpha1 namespaced role handler
type DeleteRbacAuthorizationV1alpha1NamespacedRoleHandlerFunc func(DeleteRbacAuthorizationV1alpha1NamespacedRoleParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteRbacAuthorizationV1alpha1NamespacedRoleHandlerFunc) Handle(params DeleteRbacAuthorizationV1alpha1NamespacedRoleParams) middleware.Responder {
	return fn(params)
}

// DeleteRbacAuthorizationV1alpha1NamespacedRoleHandler interface for that can handle valid delete rbac authorization v1alpha1 namespaced role params
type DeleteRbacAuthorizationV1alpha1NamespacedRoleHandler interface {
	Handle(DeleteRbacAuthorizationV1alpha1NamespacedRoleParams) middleware.Responder
}

// NewDeleteRbacAuthorizationV1alpha1NamespacedRole creates a new http.Handler for the delete rbac authorization v1alpha1 namespaced role operation
func NewDeleteRbacAuthorizationV1alpha1NamespacedRole(ctx *middleware.Context, handler DeleteRbacAuthorizationV1alpha1NamespacedRoleHandler) *DeleteRbacAuthorizationV1alpha1NamespacedRole {
	return &DeleteRbacAuthorizationV1alpha1NamespacedRole{Context: ctx, Handler: handler}
}

/*DeleteRbacAuthorizationV1alpha1NamespacedRole swagger:route DELETE /apis/rbac.authorization.k8s.io/v1alpha1/namespaces/{namespace}/roles/{name} rbacAuthorization_v1alpha1 deleteRbacAuthorizationV1alpha1NamespacedRole

delete a Role

*/
type DeleteRbacAuthorizationV1alpha1NamespacedRole struct {
	Context *middleware.Context
	Handler DeleteRbacAuthorizationV1alpha1NamespacedRoleHandler
}

func (o *DeleteRbacAuthorizationV1alpha1NamespacedRole) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteRbacAuthorizationV1alpha1NamespacedRoleParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}