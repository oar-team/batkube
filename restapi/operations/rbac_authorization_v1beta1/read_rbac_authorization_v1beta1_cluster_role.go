// Code generated by go-swagger; DO NOT EDIT.

package rbac_authorization_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ReadRbacAuthorizationV1beta1ClusterRoleHandlerFunc turns a function with the right signature into a read rbac authorization v1beta1 cluster role handler
type ReadRbacAuthorizationV1beta1ClusterRoleHandlerFunc func(ReadRbacAuthorizationV1beta1ClusterRoleParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ReadRbacAuthorizationV1beta1ClusterRoleHandlerFunc) Handle(params ReadRbacAuthorizationV1beta1ClusterRoleParams) middleware.Responder {
	return fn(params)
}

// ReadRbacAuthorizationV1beta1ClusterRoleHandler interface for that can handle valid read rbac authorization v1beta1 cluster role params
type ReadRbacAuthorizationV1beta1ClusterRoleHandler interface {
	Handle(ReadRbacAuthorizationV1beta1ClusterRoleParams) middleware.Responder
}

// NewReadRbacAuthorizationV1beta1ClusterRole creates a new http.Handler for the read rbac authorization v1beta1 cluster role operation
func NewReadRbacAuthorizationV1beta1ClusterRole(ctx *middleware.Context, handler ReadRbacAuthorizationV1beta1ClusterRoleHandler) *ReadRbacAuthorizationV1beta1ClusterRole {
	return &ReadRbacAuthorizationV1beta1ClusterRole{Context: ctx, Handler: handler}
}

/*ReadRbacAuthorizationV1beta1ClusterRole swagger:route GET /apis/rbac.authorization.k8s.io/v1beta1/clusterroles/{name} rbacAuthorization_v1beta1 readRbacAuthorizationV1beta1ClusterRole

read the specified ClusterRole

*/
type ReadRbacAuthorizationV1beta1ClusterRole struct {
	Context *middleware.Context
	Handler ReadRbacAuthorizationV1beta1ClusterRoleHandler
}

func (o *ReadRbacAuthorizationV1beta1ClusterRole) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewReadRbacAuthorizationV1beta1ClusterRoleParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}