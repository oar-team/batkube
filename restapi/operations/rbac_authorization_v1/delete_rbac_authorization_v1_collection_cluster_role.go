// Code generated by go-swagger; DO NOT EDIT.

package rbac_authorization_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteRbacAuthorizationV1CollectionClusterRoleHandlerFunc turns a function with the right signature into a delete rbac authorization v1 collection cluster role handler
type DeleteRbacAuthorizationV1CollectionClusterRoleHandlerFunc func(DeleteRbacAuthorizationV1CollectionClusterRoleParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteRbacAuthorizationV1CollectionClusterRoleHandlerFunc) Handle(params DeleteRbacAuthorizationV1CollectionClusterRoleParams) middleware.Responder {
	return fn(params)
}

// DeleteRbacAuthorizationV1CollectionClusterRoleHandler interface for that can handle valid delete rbac authorization v1 collection cluster role params
type DeleteRbacAuthorizationV1CollectionClusterRoleHandler interface {
	Handle(DeleteRbacAuthorizationV1CollectionClusterRoleParams) middleware.Responder
}

// NewDeleteRbacAuthorizationV1CollectionClusterRole creates a new http.Handler for the delete rbac authorization v1 collection cluster role operation
func NewDeleteRbacAuthorizationV1CollectionClusterRole(ctx *middleware.Context, handler DeleteRbacAuthorizationV1CollectionClusterRoleHandler) *DeleteRbacAuthorizationV1CollectionClusterRole {
	return &DeleteRbacAuthorizationV1CollectionClusterRole{Context: ctx, Handler: handler}
}

/*DeleteRbacAuthorizationV1CollectionClusterRole swagger:route DELETE /apis/rbac.authorization.k8s.io/v1/clusterroles rbacAuthorization_v1 deleteRbacAuthorizationV1CollectionClusterRole

delete collection of ClusterRole

*/
type DeleteRbacAuthorizationV1CollectionClusterRole struct {
	Context *middleware.Context
	Handler DeleteRbacAuthorizationV1CollectionClusterRoleHandler
}

func (o *DeleteRbacAuthorizationV1CollectionClusterRole) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteRbacAuthorizationV1CollectionClusterRoleParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
