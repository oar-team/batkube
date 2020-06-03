// Code generated by go-swagger; DO NOT EDIT.

package rbac_authorization_v1alpha1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// WatchRbacAuthorizationV1alpha1NamespacedRoleListHandlerFunc turns a function with the right signature into a watch rbac authorization v1alpha1 namespaced role list handler
type WatchRbacAuthorizationV1alpha1NamespacedRoleListHandlerFunc func(WatchRbacAuthorizationV1alpha1NamespacedRoleListParams) middleware.Responder

// Handle executing the request and returning a response
func (fn WatchRbacAuthorizationV1alpha1NamespacedRoleListHandlerFunc) Handle(params WatchRbacAuthorizationV1alpha1NamespacedRoleListParams) middleware.Responder {
	return fn(params)
}

// WatchRbacAuthorizationV1alpha1NamespacedRoleListHandler interface for that can handle valid watch rbac authorization v1alpha1 namespaced role list params
type WatchRbacAuthorizationV1alpha1NamespacedRoleListHandler interface {
	Handle(WatchRbacAuthorizationV1alpha1NamespacedRoleListParams) middleware.Responder
}

// NewWatchRbacAuthorizationV1alpha1NamespacedRoleList creates a new http.Handler for the watch rbac authorization v1alpha1 namespaced role list operation
func NewWatchRbacAuthorizationV1alpha1NamespacedRoleList(ctx *middleware.Context, handler WatchRbacAuthorizationV1alpha1NamespacedRoleListHandler) *WatchRbacAuthorizationV1alpha1NamespacedRoleList {
	return &WatchRbacAuthorizationV1alpha1NamespacedRoleList{Context: ctx, Handler: handler}
}

/*WatchRbacAuthorizationV1alpha1NamespacedRoleList swagger:route GET /apis/rbac.authorization.k8s.io/v1alpha1/watch/namespaces/{namespace}/roles rbacAuthorization_v1alpha1 watchRbacAuthorizationV1alpha1NamespacedRoleList

watch individual changes to a list of Role. deprecated: use the 'watch' parameter with a list operation instead.

*/
type WatchRbacAuthorizationV1alpha1NamespacedRoleList struct {
	Context *middleware.Context
	Handler WatchRbacAuthorizationV1alpha1NamespacedRoleListHandler
}

func (o *WatchRbacAuthorizationV1alpha1NamespacedRoleList) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewWatchRbacAuthorizationV1alpha1NamespacedRoleListParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
