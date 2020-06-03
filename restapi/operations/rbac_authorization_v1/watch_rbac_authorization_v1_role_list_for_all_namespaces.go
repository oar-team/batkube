// Code generated by go-swagger; DO NOT EDIT.

package rbac_authorization_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// WatchRbacAuthorizationV1RoleListForAllNamespacesHandlerFunc turns a function with the right signature into a watch rbac authorization v1 role list for all namespaces handler
type WatchRbacAuthorizationV1RoleListForAllNamespacesHandlerFunc func(WatchRbacAuthorizationV1RoleListForAllNamespacesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn WatchRbacAuthorizationV1RoleListForAllNamespacesHandlerFunc) Handle(params WatchRbacAuthorizationV1RoleListForAllNamespacesParams) middleware.Responder {
	return fn(params)
}

// WatchRbacAuthorizationV1RoleListForAllNamespacesHandler interface for that can handle valid watch rbac authorization v1 role list for all namespaces params
type WatchRbacAuthorizationV1RoleListForAllNamespacesHandler interface {
	Handle(WatchRbacAuthorizationV1RoleListForAllNamespacesParams) middleware.Responder
}

// NewWatchRbacAuthorizationV1RoleListForAllNamespaces creates a new http.Handler for the watch rbac authorization v1 role list for all namespaces operation
func NewWatchRbacAuthorizationV1RoleListForAllNamespaces(ctx *middleware.Context, handler WatchRbacAuthorizationV1RoleListForAllNamespacesHandler) *WatchRbacAuthorizationV1RoleListForAllNamespaces {
	return &WatchRbacAuthorizationV1RoleListForAllNamespaces{Context: ctx, Handler: handler}
}

/*WatchRbacAuthorizationV1RoleListForAllNamespaces swagger:route GET /apis/rbac.authorization.k8s.io/v1/watch/roles rbacAuthorization_v1 watchRbacAuthorizationV1RoleListForAllNamespaces

watch individual changes to a list of Role. deprecated: use the 'watch' parameter with a list operation instead.

*/
type WatchRbacAuthorizationV1RoleListForAllNamespaces struct {
	Context *middleware.Context
	Handler WatchRbacAuthorizationV1RoleListForAllNamespacesHandler
}

func (o *WatchRbacAuthorizationV1RoleListForAllNamespaces) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewWatchRbacAuthorizationV1RoleListForAllNamespacesParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
