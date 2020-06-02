// Code generated by go-swagger; DO NOT EDIT.

package rbac_authorization_v1alpha1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteRbacAuthorizationV1alpha1CollectionClusterRoleBindingHandlerFunc turns a function with the right signature into a delete rbac authorization v1alpha1 collection cluster role binding handler
type DeleteRbacAuthorizationV1alpha1CollectionClusterRoleBindingHandlerFunc func(DeleteRbacAuthorizationV1alpha1CollectionClusterRoleBindingParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteRbacAuthorizationV1alpha1CollectionClusterRoleBindingHandlerFunc) Handle(params DeleteRbacAuthorizationV1alpha1CollectionClusterRoleBindingParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// DeleteRbacAuthorizationV1alpha1CollectionClusterRoleBindingHandler interface for that can handle valid delete rbac authorization v1alpha1 collection cluster role binding params
type DeleteRbacAuthorizationV1alpha1CollectionClusterRoleBindingHandler interface {
	Handle(DeleteRbacAuthorizationV1alpha1CollectionClusterRoleBindingParams, interface{}) middleware.Responder
}

// NewDeleteRbacAuthorizationV1alpha1CollectionClusterRoleBinding creates a new http.Handler for the delete rbac authorization v1alpha1 collection cluster role binding operation
func NewDeleteRbacAuthorizationV1alpha1CollectionClusterRoleBinding(ctx *middleware.Context, handler DeleteRbacAuthorizationV1alpha1CollectionClusterRoleBindingHandler) *DeleteRbacAuthorizationV1alpha1CollectionClusterRoleBinding {
	return &DeleteRbacAuthorizationV1alpha1CollectionClusterRoleBinding{Context: ctx, Handler: handler}
}

/*DeleteRbacAuthorizationV1alpha1CollectionClusterRoleBinding swagger:route DELETE /apis/rbac.authorization.k8s.io/v1alpha1/clusterrolebindings rbacAuthorization_v1alpha1 deleteRbacAuthorizationV1alpha1CollectionClusterRoleBinding

delete collection of ClusterRoleBinding

*/
type DeleteRbacAuthorizationV1alpha1CollectionClusterRoleBinding struct {
	Context *middleware.Context
	Handler DeleteRbacAuthorizationV1alpha1CollectionClusterRoleBindingHandler
}

func (o *DeleteRbacAuthorizationV1alpha1CollectionClusterRoleBinding) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteRbacAuthorizationV1alpha1CollectionClusterRoleBindingParams()

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
