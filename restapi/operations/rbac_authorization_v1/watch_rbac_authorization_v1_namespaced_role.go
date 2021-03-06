// Code generated by go-swagger; DO NOT EDIT.

package rbac_authorization_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// WatchRbacAuthorizationV1NamespacedRoleHandlerFunc turns a function with the right signature into a watch rbac authorization v1 namespaced role handler
type WatchRbacAuthorizationV1NamespacedRoleHandlerFunc func(WatchRbacAuthorizationV1NamespacedRoleParams) middleware.Responder

// Handle executing the request and returning a response
func (fn WatchRbacAuthorizationV1NamespacedRoleHandlerFunc) Handle(params WatchRbacAuthorizationV1NamespacedRoleParams) middleware.Responder {
	return fn(params)
}

// WatchRbacAuthorizationV1NamespacedRoleHandler interface for that can handle valid watch rbac authorization v1 namespaced role params
type WatchRbacAuthorizationV1NamespacedRoleHandler interface {
	Handle(WatchRbacAuthorizationV1NamespacedRoleParams) middleware.Responder
}

// NewWatchRbacAuthorizationV1NamespacedRole creates a new http.Handler for the watch rbac authorization v1 namespaced role operation
func NewWatchRbacAuthorizationV1NamespacedRole(ctx *middleware.Context, handler WatchRbacAuthorizationV1NamespacedRoleHandler) *WatchRbacAuthorizationV1NamespacedRole {
	return &WatchRbacAuthorizationV1NamespacedRole{Context: ctx, Handler: handler}
}

/*WatchRbacAuthorizationV1NamespacedRole swagger:route GET /apis/rbac.authorization.k8s.io/v1/watch/namespaces/{namespace}/roles/{name} rbacAuthorization_v1 watchRbacAuthorizationV1NamespacedRole

watch changes to an object of kind Role. deprecated: use the 'watch' parameter with a list operation instead, filtered to a single item with the 'fieldSelector' parameter.

*/
type WatchRbacAuthorizationV1NamespacedRole struct {
	Context *middleware.Context
	Handler WatchRbacAuthorizationV1NamespacedRoleHandler
}

func (o *WatchRbacAuthorizationV1NamespacedRole) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewWatchRbacAuthorizationV1NamespacedRoleParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
