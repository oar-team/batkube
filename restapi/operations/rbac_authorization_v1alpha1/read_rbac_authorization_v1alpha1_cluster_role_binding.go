// Code generated by go-swagger; DO NOT EDIT.

package rbac_authorization_v1alpha1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ReadRbacAuthorizationV1alpha1ClusterRoleBindingHandlerFunc turns a function with the right signature into a read rbac authorization v1alpha1 cluster role binding handler
type ReadRbacAuthorizationV1alpha1ClusterRoleBindingHandlerFunc func(ReadRbacAuthorizationV1alpha1ClusterRoleBindingParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn ReadRbacAuthorizationV1alpha1ClusterRoleBindingHandlerFunc) Handle(params ReadRbacAuthorizationV1alpha1ClusterRoleBindingParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// ReadRbacAuthorizationV1alpha1ClusterRoleBindingHandler interface for that can handle valid read rbac authorization v1alpha1 cluster role binding params
type ReadRbacAuthorizationV1alpha1ClusterRoleBindingHandler interface {
	Handle(ReadRbacAuthorizationV1alpha1ClusterRoleBindingParams, interface{}) middleware.Responder
}

// NewReadRbacAuthorizationV1alpha1ClusterRoleBinding creates a new http.Handler for the read rbac authorization v1alpha1 cluster role binding operation
func NewReadRbacAuthorizationV1alpha1ClusterRoleBinding(ctx *middleware.Context, handler ReadRbacAuthorizationV1alpha1ClusterRoleBindingHandler) *ReadRbacAuthorizationV1alpha1ClusterRoleBinding {
	return &ReadRbacAuthorizationV1alpha1ClusterRoleBinding{Context: ctx, Handler: handler}
}

/*ReadRbacAuthorizationV1alpha1ClusterRoleBinding swagger:route GET /apis/rbac.authorization.k8s.io/v1alpha1/clusterrolebindings/{name} rbacAuthorization_v1alpha1 readRbacAuthorizationV1alpha1ClusterRoleBinding

read the specified ClusterRoleBinding

*/
type ReadRbacAuthorizationV1alpha1ClusterRoleBinding struct {
	Context *middleware.Context
	Handler ReadRbacAuthorizationV1alpha1ClusterRoleBindingHandler
}

func (o *ReadRbacAuthorizationV1alpha1ClusterRoleBinding) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewReadRbacAuthorizationV1alpha1ClusterRoleBindingParams()

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
