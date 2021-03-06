// Code generated by go-swagger; DO NOT EDIT.

package rbac_authorization_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// CreateRbacAuthorizationV1beta1ClusterRoleHandlerFunc turns a function with the right signature into a create rbac authorization v1beta1 cluster role handler
type CreateRbacAuthorizationV1beta1ClusterRoleHandlerFunc func(CreateRbacAuthorizationV1beta1ClusterRoleParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateRbacAuthorizationV1beta1ClusterRoleHandlerFunc) Handle(params CreateRbacAuthorizationV1beta1ClusterRoleParams) middleware.Responder {
	return fn(params)
}

// CreateRbacAuthorizationV1beta1ClusterRoleHandler interface for that can handle valid create rbac authorization v1beta1 cluster role params
type CreateRbacAuthorizationV1beta1ClusterRoleHandler interface {
	Handle(CreateRbacAuthorizationV1beta1ClusterRoleParams) middleware.Responder
}

// NewCreateRbacAuthorizationV1beta1ClusterRole creates a new http.Handler for the create rbac authorization v1beta1 cluster role operation
func NewCreateRbacAuthorizationV1beta1ClusterRole(ctx *middleware.Context, handler CreateRbacAuthorizationV1beta1ClusterRoleHandler) *CreateRbacAuthorizationV1beta1ClusterRole {
	return &CreateRbacAuthorizationV1beta1ClusterRole{Context: ctx, Handler: handler}
}

/*CreateRbacAuthorizationV1beta1ClusterRole swagger:route POST /apis/rbac.authorization.k8s.io/v1beta1/clusterroles rbacAuthorization_v1beta1 createRbacAuthorizationV1beta1ClusterRole

create a ClusterRole

*/
type CreateRbacAuthorizationV1beta1ClusterRole struct {
	Context *middleware.Context
	Handler CreateRbacAuthorizationV1beta1ClusterRoleHandler
}

func (o *CreateRbacAuthorizationV1beta1ClusterRole) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCreateRbacAuthorizationV1beta1ClusterRoleParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
