// Code generated by go-swagger; DO NOT EDIT.

package rbac_authorization_v1alpha1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ListRbacAuthorizationV1alpha1RoleBindingForAllNamespacesHandlerFunc turns a function with the right signature into a list rbac authorization v1alpha1 role binding for all namespaces handler
type ListRbacAuthorizationV1alpha1RoleBindingForAllNamespacesHandlerFunc func(ListRbacAuthorizationV1alpha1RoleBindingForAllNamespacesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ListRbacAuthorizationV1alpha1RoleBindingForAllNamespacesHandlerFunc) Handle(params ListRbacAuthorizationV1alpha1RoleBindingForAllNamespacesParams) middleware.Responder {
	return fn(params)
}

// ListRbacAuthorizationV1alpha1RoleBindingForAllNamespacesHandler interface for that can handle valid list rbac authorization v1alpha1 role binding for all namespaces params
type ListRbacAuthorizationV1alpha1RoleBindingForAllNamespacesHandler interface {
	Handle(ListRbacAuthorizationV1alpha1RoleBindingForAllNamespacesParams) middleware.Responder
}

// NewListRbacAuthorizationV1alpha1RoleBindingForAllNamespaces creates a new http.Handler for the list rbac authorization v1alpha1 role binding for all namespaces operation
func NewListRbacAuthorizationV1alpha1RoleBindingForAllNamespaces(ctx *middleware.Context, handler ListRbacAuthorizationV1alpha1RoleBindingForAllNamespacesHandler) *ListRbacAuthorizationV1alpha1RoleBindingForAllNamespaces {
	return &ListRbacAuthorizationV1alpha1RoleBindingForAllNamespaces{Context: ctx, Handler: handler}
}

/*ListRbacAuthorizationV1alpha1RoleBindingForAllNamespaces swagger:route GET /apis/rbac.authorization.k8s.io/v1alpha1/rolebindings rbacAuthorization_v1alpha1 listRbacAuthorizationV1alpha1RoleBindingForAllNamespaces

list or watch objects of kind RoleBinding

*/
type ListRbacAuthorizationV1alpha1RoleBindingForAllNamespaces struct {
	Context *middleware.Context
	Handler ListRbacAuthorizationV1alpha1RoleBindingForAllNamespacesHandler
}

func (o *ListRbacAuthorizationV1alpha1RoleBindingForAllNamespaces) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewListRbacAuthorizationV1alpha1RoleBindingForAllNamespacesParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
