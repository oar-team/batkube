// Code generated by go-swagger; DO NOT EDIT.

package rbac_authorization_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// WatchRbacAuthorizationV1NamespacedRoleBindingListHandlerFunc turns a function with the right signature into a watch rbac authorization v1 namespaced role binding list handler
type WatchRbacAuthorizationV1NamespacedRoleBindingListHandlerFunc func(WatchRbacAuthorizationV1NamespacedRoleBindingListParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn WatchRbacAuthorizationV1NamespacedRoleBindingListHandlerFunc) Handle(params WatchRbacAuthorizationV1NamespacedRoleBindingListParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// WatchRbacAuthorizationV1NamespacedRoleBindingListHandler interface for that can handle valid watch rbac authorization v1 namespaced role binding list params
type WatchRbacAuthorizationV1NamespacedRoleBindingListHandler interface {
	Handle(WatchRbacAuthorizationV1NamespacedRoleBindingListParams, interface{}) middleware.Responder
}

// NewWatchRbacAuthorizationV1NamespacedRoleBindingList creates a new http.Handler for the watch rbac authorization v1 namespaced role binding list operation
func NewWatchRbacAuthorizationV1NamespacedRoleBindingList(ctx *middleware.Context, handler WatchRbacAuthorizationV1NamespacedRoleBindingListHandler) *WatchRbacAuthorizationV1NamespacedRoleBindingList {
	return &WatchRbacAuthorizationV1NamespacedRoleBindingList{Context: ctx, Handler: handler}
}

/*WatchRbacAuthorizationV1NamespacedRoleBindingList swagger:route GET /apis/rbac.authorization.k8s.io/v1/watch/namespaces/{namespace}/rolebindings rbacAuthorization_v1 watchRbacAuthorizationV1NamespacedRoleBindingList

watch individual changes to a list of RoleBinding. deprecated: use the 'watch' parameter with a list operation instead.

*/
type WatchRbacAuthorizationV1NamespacedRoleBindingList struct {
	Context *middleware.Context
	Handler WatchRbacAuthorizationV1NamespacedRoleBindingListHandler
}

func (o *WatchRbacAuthorizationV1NamespacedRoleBindingList) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewWatchRbacAuthorizationV1NamespacedRoleBindingListParams()

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
