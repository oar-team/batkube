// Code generated by go-swagger; DO NOT EDIT.

package rbac_authorization_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ReplaceRbacAuthorizationV1beta1ClusterRoleHandlerFunc turns a function with the right signature into a replace rbac authorization v1beta1 cluster role handler
type ReplaceRbacAuthorizationV1beta1ClusterRoleHandlerFunc func(ReplaceRbacAuthorizationV1beta1ClusterRoleParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ReplaceRbacAuthorizationV1beta1ClusterRoleHandlerFunc) Handle(params ReplaceRbacAuthorizationV1beta1ClusterRoleParams) middleware.Responder {
	return fn(params)
}

// ReplaceRbacAuthorizationV1beta1ClusterRoleHandler interface for that can handle valid replace rbac authorization v1beta1 cluster role params
type ReplaceRbacAuthorizationV1beta1ClusterRoleHandler interface {
	Handle(ReplaceRbacAuthorizationV1beta1ClusterRoleParams) middleware.Responder
}

// NewReplaceRbacAuthorizationV1beta1ClusterRole creates a new http.Handler for the replace rbac authorization v1beta1 cluster role operation
func NewReplaceRbacAuthorizationV1beta1ClusterRole(ctx *middleware.Context, handler ReplaceRbacAuthorizationV1beta1ClusterRoleHandler) *ReplaceRbacAuthorizationV1beta1ClusterRole {
	return &ReplaceRbacAuthorizationV1beta1ClusterRole{Context: ctx, Handler: handler}
}

/*ReplaceRbacAuthorizationV1beta1ClusterRole swagger:route PUT /apis/rbac.authorization.k8s.io/v1beta1/clusterroles/{name} rbacAuthorization_v1beta1 replaceRbacAuthorizationV1beta1ClusterRole

replace the specified ClusterRole

*/
type ReplaceRbacAuthorizationV1beta1ClusterRole struct {
	Context *middleware.Context
	Handler ReplaceRbacAuthorizationV1beta1ClusterRoleHandler
}

func (o *ReplaceRbacAuthorizationV1beta1ClusterRole) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewReplaceRbacAuthorizationV1beta1ClusterRoleParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}