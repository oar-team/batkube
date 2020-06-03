// Code generated by go-swagger; DO NOT EDIT.

package rbac_authorization_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ListRbacAuthorizationV1beta1NamespacedRoleHandlerFunc turns a function with the right signature into a list rbac authorization v1beta1 namespaced role handler
type ListRbacAuthorizationV1beta1NamespacedRoleHandlerFunc func(ListRbacAuthorizationV1beta1NamespacedRoleParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ListRbacAuthorizationV1beta1NamespacedRoleHandlerFunc) Handle(params ListRbacAuthorizationV1beta1NamespacedRoleParams) middleware.Responder {
	return fn(params)
}

// ListRbacAuthorizationV1beta1NamespacedRoleHandler interface for that can handle valid list rbac authorization v1beta1 namespaced role params
type ListRbacAuthorizationV1beta1NamespacedRoleHandler interface {
	Handle(ListRbacAuthorizationV1beta1NamespacedRoleParams) middleware.Responder
}

// NewListRbacAuthorizationV1beta1NamespacedRole creates a new http.Handler for the list rbac authorization v1beta1 namespaced role operation
func NewListRbacAuthorizationV1beta1NamespacedRole(ctx *middleware.Context, handler ListRbacAuthorizationV1beta1NamespacedRoleHandler) *ListRbacAuthorizationV1beta1NamespacedRole {
	return &ListRbacAuthorizationV1beta1NamespacedRole{Context: ctx, Handler: handler}
}

/*ListRbacAuthorizationV1beta1NamespacedRole swagger:route GET /apis/rbac.authorization.k8s.io/v1beta1/namespaces/{namespace}/roles rbacAuthorization_v1beta1 listRbacAuthorizationV1beta1NamespacedRole

list or watch objects of kind Role

*/
type ListRbacAuthorizationV1beta1NamespacedRole struct {
	Context *middleware.Context
	Handler ListRbacAuthorizationV1beta1NamespacedRoleHandler
}

func (o *ListRbacAuthorizationV1beta1NamespacedRole) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewListRbacAuthorizationV1beta1NamespacedRoleParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
