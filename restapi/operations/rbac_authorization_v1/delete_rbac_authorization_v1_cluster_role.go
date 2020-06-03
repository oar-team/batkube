// Code generated by go-swagger; DO NOT EDIT.

package rbac_authorization_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteRbacAuthorizationV1ClusterRoleHandlerFunc turns a function with the right signature into a delete rbac authorization v1 cluster role handler
type DeleteRbacAuthorizationV1ClusterRoleHandlerFunc func(DeleteRbacAuthorizationV1ClusterRoleParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteRbacAuthorizationV1ClusterRoleHandlerFunc) Handle(params DeleteRbacAuthorizationV1ClusterRoleParams) middleware.Responder {
	return fn(params)
}

// DeleteRbacAuthorizationV1ClusterRoleHandler interface for that can handle valid delete rbac authorization v1 cluster role params
type DeleteRbacAuthorizationV1ClusterRoleHandler interface {
	Handle(DeleteRbacAuthorizationV1ClusterRoleParams) middleware.Responder
}

// NewDeleteRbacAuthorizationV1ClusterRole creates a new http.Handler for the delete rbac authorization v1 cluster role operation
func NewDeleteRbacAuthorizationV1ClusterRole(ctx *middleware.Context, handler DeleteRbacAuthorizationV1ClusterRoleHandler) *DeleteRbacAuthorizationV1ClusterRole {
	return &DeleteRbacAuthorizationV1ClusterRole{Context: ctx, Handler: handler}
}

/*DeleteRbacAuthorizationV1ClusterRole swagger:route DELETE /apis/rbac.authorization.k8s.io/v1/clusterroles/{name} rbacAuthorization_v1 deleteRbacAuthorizationV1ClusterRole

delete a ClusterRole

*/
type DeleteRbacAuthorizationV1ClusterRole struct {
	Context *middleware.Context
	Handler DeleteRbacAuthorizationV1ClusterRoleHandler
}

func (o *DeleteRbacAuthorizationV1ClusterRole) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteRbacAuthorizationV1ClusterRoleParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}