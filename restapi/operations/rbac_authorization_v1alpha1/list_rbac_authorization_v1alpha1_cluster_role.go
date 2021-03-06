// Code generated by go-swagger; DO NOT EDIT.

package rbac_authorization_v1alpha1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ListRbacAuthorizationV1alpha1ClusterRoleHandlerFunc turns a function with the right signature into a list rbac authorization v1alpha1 cluster role handler
type ListRbacAuthorizationV1alpha1ClusterRoleHandlerFunc func(ListRbacAuthorizationV1alpha1ClusterRoleParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ListRbacAuthorizationV1alpha1ClusterRoleHandlerFunc) Handle(params ListRbacAuthorizationV1alpha1ClusterRoleParams) middleware.Responder {
	return fn(params)
}

// ListRbacAuthorizationV1alpha1ClusterRoleHandler interface for that can handle valid list rbac authorization v1alpha1 cluster role params
type ListRbacAuthorizationV1alpha1ClusterRoleHandler interface {
	Handle(ListRbacAuthorizationV1alpha1ClusterRoleParams) middleware.Responder
}

// NewListRbacAuthorizationV1alpha1ClusterRole creates a new http.Handler for the list rbac authorization v1alpha1 cluster role operation
func NewListRbacAuthorizationV1alpha1ClusterRole(ctx *middleware.Context, handler ListRbacAuthorizationV1alpha1ClusterRoleHandler) *ListRbacAuthorizationV1alpha1ClusterRole {
	return &ListRbacAuthorizationV1alpha1ClusterRole{Context: ctx, Handler: handler}
}

/*ListRbacAuthorizationV1alpha1ClusterRole swagger:route GET /apis/rbac.authorization.k8s.io/v1alpha1/clusterroles rbacAuthorization_v1alpha1 listRbacAuthorizationV1alpha1ClusterRole

list or watch objects of kind ClusterRole

*/
type ListRbacAuthorizationV1alpha1ClusterRole struct {
	Context *middleware.Context
	Handler ListRbacAuthorizationV1alpha1ClusterRoleHandler
}

func (o *ListRbacAuthorizationV1alpha1ClusterRole) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewListRbacAuthorizationV1alpha1ClusterRoleParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
