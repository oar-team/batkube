// Code generated by go-swagger; DO NOT EDIT.

package rbac_authorization_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// WatchRbacAuthorizationV1beta1RoleListForAllNamespacesHandlerFunc turns a function with the right signature into a watch rbac authorization v1beta1 role list for all namespaces handler
type WatchRbacAuthorizationV1beta1RoleListForAllNamespacesHandlerFunc func(WatchRbacAuthorizationV1beta1RoleListForAllNamespacesParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn WatchRbacAuthorizationV1beta1RoleListForAllNamespacesHandlerFunc) Handle(params WatchRbacAuthorizationV1beta1RoleListForAllNamespacesParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// WatchRbacAuthorizationV1beta1RoleListForAllNamespacesHandler interface for that can handle valid watch rbac authorization v1beta1 role list for all namespaces params
type WatchRbacAuthorizationV1beta1RoleListForAllNamespacesHandler interface {
	Handle(WatchRbacAuthorizationV1beta1RoleListForAllNamespacesParams, interface{}) middleware.Responder
}

// NewWatchRbacAuthorizationV1beta1RoleListForAllNamespaces creates a new http.Handler for the watch rbac authorization v1beta1 role list for all namespaces operation
func NewWatchRbacAuthorizationV1beta1RoleListForAllNamespaces(ctx *middleware.Context, handler WatchRbacAuthorizationV1beta1RoleListForAllNamespacesHandler) *WatchRbacAuthorizationV1beta1RoleListForAllNamespaces {
	return &WatchRbacAuthorizationV1beta1RoleListForAllNamespaces{Context: ctx, Handler: handler}
}

/*WatchRbacAuthorizationV1beta1RoleListForAllNamespaces swagger:route GET /apis/rbac.authorization.k8s.io/v1beta1/watch/roles rbacAuthorization_v1beta1 watchRbacAuthorizationV1beta1RoleListForAllNamespaces

watch individual changes to a list of Role. deprecated: use the 'watch' parameter with a list operation instead.

*/
type WatchRbacAuthorizationV1beta1RoleListForAllNamespaces struct {
	Context *middleware.Context
	Handler WatchRbacAuthorizationV1beta1RoleListForAllNamespacesHandler
}

func (o *WatchRbacAuthorizationV1beta1RoleListForAllNamespaces) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewWatchRbacAuthorizationV1beta1RoleListForAllNamespacesParams()

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
