// Code generated by go-swagger; DO NOT EDIT.

package rbac_authorization_v1alpha1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// WatchRbacAuthorizationV1alpha1NamespacedRoleBindingHandlerFunc turns a function with the right signature into a watch rbac authorization v1alpha1 namespaced role binding handler
type WatchRbacAuthorizationV1alpha1NamespacedRoleBindingHandlerFunc func(WatchRbacAuthorizationV1alpha1NamespacedRoleBindingParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn WatchRbacAuthorizationV1alpha1NamespacedRoleBindingHandlerFunc) Handle(params WatchRbacAuthorizationV1alpha1NamespacedRoleBindingParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// WatchRbacAuthorizationV1alpha1NamespacedRoleBindingHandler interface for that can handle valid watch rbac authorization v1alpha1 namespaced role binding params
type WatchRbacAuthorizationV1alpha1NamespacedRoleBindingHandler interface {
	Handle(WatchRbacAuthorizationV1alpha1NamespacedRoleBindingParams, interface{}) middleware.Responder
}

// NewWatchRbacAuthorizationV1alpha1NamespacedRoleBinding creates a new http.Handler for the watch rbac authorization v1alpha1 namespaced role binding operation
func NewWatchRbacAuthorizationV1alpha1NamespacedRoleBinding(ctx *middleware.Context, handler WatchRbacAuthorizationV1alpha1NamespacedRoleBindingHandler) *WatchRbacAuthorizationV1alpha1NamespacedRoleBinding {
	return &WatchRbacAuthorizationV1alpha1NamespacedRoleBinding{Context: ctx, Handler: handler}
}

/*WatchRbacAuthorizationV1alpha1NamespacedRoleBinding swagger:route GET /apis/rbac.authorization.k8s.io/v1alpha1/watch/namespaces/{namespace}/rolebindings/{name} rbacAuthorization_v1alpha1 watchRbacAuthorizationV1alpha1NamespacedRoleBinding

watch changes to an object of kind RoleBinding. deprecated: use the 'watch' parameter with a list operation instead, filtered to a single item with the 'fieldSelector' parameter.

*/
type WatchRbacAuthorizationV1alpha1NamespacedRoleBinding struct {
	Context *middleware.Context
	Handler WatchRbacAuthorizationV1alpha1NamespacedRoleBindingHandler
}

func (o *WatchRbacAuthorizationV1alpha1NamespacedRoleBinding) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewWatchRbacAuthorizationV1alpha1NamespacedRoleBindingParams()

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
