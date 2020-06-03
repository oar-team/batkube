// Code generated by go-swagger; DO NOT EDIT.

package rbac_authorization_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ListRbacAuthorizationV1NamespacedRoleBindingHandlerFunc turns a function with the right signature into a list rbac authorization v1 namespaced role binding handler
type ListRbacAuthorizationV1NamespacedRoleBindingHandlerFunc func(ListRbacAuthorizationV1NamespacedRoleBindingParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ListRbacAuthorizationV1NamespacedRoleBindingHandlerFunc) Handle(params ListRbacAuthorizationV1NamespacedRoleBindingParams) middleware.Responder {
	return fn(params)
}

// ListRbacAuthorizationV1NamespacedRoleBindingHandler interface for that can handle valid list rbac authorization v1 namespaced role binding params
type ListRbacAuthorizationV1NamespacedRoleBindingHandler interface {
	Handle(ListRbacAuthorizationV1NamespacedRoleBindingParams) middleware.Responder
}

// NewListRbacAuthorizationV1NamespacedRoleBinding creates a new http.Handler for the list rbac authorization v1 namespaced role binding operation
func NewListRbacAuthorizationV1NamespacedRoleBinding(ctx *middleware.Context, handler ListRbacAuthorizationV1NamespacedRoleBindingHandler) *ListRbacAuthorizationV1NamespacedRoleBinding {
	return &ListRbacAuthorizationV1NamespacedRoleBinding{Context: ctx, Handler: handler}
}

/*ListRbacAuthorizationV1NamespacedRoleBinding swagger:route GET /apis/rbac.authorization.k8s.io/v1/namespaces/{namespace}/rolebindings rbacAuthorization_v1 listRbacAuthorizationV1NamespacedRoleBinding

list or watch objects of kind RoleBinding

*/
type ListRbacAuthorizationV1NamespacedRoleBinding struct {
	Context *middleware.Context
	Handler ListRbacAuthorizationV1NamespacedRoleBindingHandler
}

func (o *ListRbacAuthorizationV1NamespacedRoleBinding) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewListRbacAuthorizationV1NamespacedRoleBindingParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}