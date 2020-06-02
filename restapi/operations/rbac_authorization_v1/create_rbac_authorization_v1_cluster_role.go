// Code generated by go-swagger; DO NOT EDIT.

package rbac_authorization_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// CreateRbacAuthorizationV1ClusterRoleHandlerFunc turns a function with the right signature into a create rbac authorization v1 cluster role handler
type CreateRbacAuthorizationV1ClusterRoleHandlerFunc func(CreateRbacAuthorizationV1ClusterRoleParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateRbacAuthorizationV1ClusterRoleHandlerFunc) Handle(params CreateRbacAuthorizationV1ClusterRoleParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// CreateRbacAuthorizationV1ClusterRoleHandler interface for that can handle valid create rbac authorization v1 cluster role params
type CreateRbacAuthorizationV1ClusterRoleHandler interface {
	Handle(CreateRbacAuthorizationV1ClusterRoleParams, interface{}) middleware.Responder
}

// NewCreateRbacAuthorizationV1ClusterRole creates a new http.Handler for the create rbac authorization v1 cluster role operation
func NewCreateRbacAuthorizationV1ClusterRole(ctx *middleware.Context, handler CreateRbacAuthorizationV1ClusterRoleHandler) *CreateRbacAuthorizationV1ClusterRole {
	return &CreateRbacAuthorizationV1ClusterRole{Context: ctx, Handler: handler}
}

/*CreateRbacAuthorizationV1ClusterRole swagger:route POST /apis/rbac.authorization.k8s.io/v1/clusterroles rbacAuthorization_v1 createRbacAuthorizationV1ClusterRole

create a ClusterRole

*/
type CreateRbacAuthorizationV1ClusterRole struct {
	Context *middleware.Context
	Handler CreateRbacAuthorizationV1ClusterRoleHandler
}

func (o *CreateRbacAuthorizationV1ClusterRole) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCreateRbacAuthorizationV1ClusterRoleParams()

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
