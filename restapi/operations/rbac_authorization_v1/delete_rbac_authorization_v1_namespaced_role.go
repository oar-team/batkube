// Code generated by go-swagger; DO NOT EDIT.

package rbac_authorization_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteRbacAuthorizationV1NamespacedRoleHandlerFunc turns a function with the right signature into a delete rbac authorization v1 namespaced role handler
type DeleteRbacAuthorizationV1NamespacedRoleHandlerFunc func(DeleteRbacAuthorizationV1NamespacedRoleParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteRbacAuthorizationV1NamespacedRoleHandlerFunc) Handle(params DeleteRbacAuthorizationV1NamespacedRoleParams) middleware.Responder {
	return fn(params)
}

// DeleteRbacAuthorizationV1NamespacedRoleHandler interface for that can handle valid delete rbac authorization v1 namespaced role params
type DeleteRbacAuthorizationV1NamespacedRoleHandler interface {
	Handle(DeleteRbacAuthorizationV1NamespacedRoleParams) middleware.Responder
}

// NewDeleteRbacAuthorizationV1NamespacedRole creates a new http.Handler for the delete rbac authorization v1 namespaced role operation
func NewDeleteRbacAuthorizationV1NamespacedRole(ctx *middleware.Context, handler DeleteRbacAuthorizationV1NamespacedRoleHandler) *DeleteRbacAuthorizationV1NamespacedRole {
	return &DeleteRbacAuthorizationV1NamespacedRole{Context: ctx, Handler: handler}
}

/*DeleteRbacAuthorizationV1NamespacedRole swagger:route DELETE /apis/rbac.authorization.k8s.io/v1/namespaces/{namespace}/roles/{name} rbacAuthorization_v1 deleteRbacAuthorizationV1NamespacedRole

delete a Role

*/
type DeleteRbacAuthorizationV1NamespacedRole struct {
	Context *middleware.Context
	Handler DeleteRbacAuthorizationV1NamespacedRoleHandler
}

func (o *DeleteRbacAuthorizationV1NamespacedRole) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteRbacAuthorizationV1NamespacedRoleParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
