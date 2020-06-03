// Code generated by go-swagger; DO NOT EDIT.

package rbac_authorization_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// WatchRbacAuthorizationV1beta1NamespacedRoleBindingListHandlerFunc turns a function with the right signature into a watch rbac authorization v1beta1 namespaced role binding list handler
type WatchRbacAuthorizationV1beta1NamespacedRoleBindingListHandlerFunc func(WatchRbacAuthorizationV1beta1NamespacedRoleBindingListParams) middleware.Responder

// Handle executing the request and returning a response
func (fn WatchRbacAuthorizationV1beta1NamespacedRoleBindingListHandlerFunc) Handle(params WatchRbacAuthorizationV1beta1NamespacedRoleBindingListParams) middleware.Responder {
	return fn(params)
}

// WatchRbacAuthorizationV1beta1NamespacedRoleBindingListHandler interface for that can handle valid watch rbac authorization v1beta1 namespaced role binding list params
type WatchRbacAuthorizationV1beta1NamespacedRoleBindingListHandler interface {
	Handle(WatchRbacAuthorizationV1beta1NamespacedRoleBindingListParams) middleware.Responder
}

// NewWatchRbacAuthorizationV1beta1NamespacedRoleBindingList creates a new http.Handler for the watch rbac authorization v1beta1 namespaced role binding list operation
func NewWatchRbacAuthorizationV1beta1NamespacedRoleBindingList(ctx *middleware.Context, handler WatchRbacAuthorizationV1beta1NamespacedRoleBindingListHandler) *WatchRbacAuthorizationV1beta1NamespacedRoleBindingList {
	return &WatchRbacAuthorizationV1beta1NamespacedRoleBindingList{Context: ctx, Handler: handler}
}

/*WatchRbacAuthorizationV1beta1NamespacedRoleBindingList swagger:route GET /apis/rbac.authorization.k8s.io/v1beta1/watch/namespaces/{namespace}/rolebindings rbacAuthorization_v1beta1 watchRbacAuthorizationV1beta1NamespacedRoleBindingList

watch individual changes to a list of RoleBinding. deprecated: use the 'watch' parameter with a list operation instead.

*/
type WatchRbacAuthorizationV1beta1NamespacedRoleBindingList struct {
	Context *middleware.Context
	Handler WatchRbacAuthorizationV1beta1NamespacedRoleBindingListHandler
}

func (o *WatchRbacAuthorizationV1beta1NamespacedRoleBindingList) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewWatchRbacAuthorizationV1beta1NamespacedRoleBindingListParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
