// Code generated by go-swagger; DO NOT EDIT.

package rbac_authorization_v1alpha1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PatchRbacAuthorizationV1alpha1NamespacedRoleHandlerFunc turns a function with the right signature into a patch rbac authorization v1alpha1 namespaced role handler
type PatchRbacAuthorizationV1alpha1NamespacedRoleHandlerFunc func(PatchRbacAuthorizationV1alpha1NamespacedRoleParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn PatchRbacAuthorizationV1alpha1NamespacedRoleHandlerFunc) Handle(params PatchRbacAuthorizationV1alpha1NamespacedRoleParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// PatchRbacAuthorizationV1alpha1NamespacedRoleHandler interface for that can handle valid patch rbac authorization v1alpha1 namespaced role params
type PatchRbacAuthorizationV1alpha1NamespacedRoleHandler interface {
	Handle(PatchRbacAuthorizationV1alpha1NamespacedRoleParams, interface{}) middleware.Responder
}

// NewPatchRbacAuthorizationV1alpha1NamespacedRole creates a new http.Handler for the patch rbac authorization v1alpha1 namespaced role operation
func NewPatchRbacAuthorizationV1alpha1NamespacedRole(ctx *middleware.Context, handler PatchRbacAuthorizationV1alpha1NamespacedRoleHandler) *PatchRbacAuthorizationV1alpha1NamespacedRole {
	return &PatchRbacAuthorizationV1alpha1NamespacedRole{Context: ctx, Handler: handler}
}

/*PatchRbacAuthorizationV1alpha1NamespacedRole swagger:route PATCH /apis/rbac.authorization.k8s.io/v1alpha1/namespaces/{namespace}/roles/{name} rbacAuthorization_v1alpha1 patchRbacAuthorizationV1alpha1NamespacedRole

partially update the specified Role

*/
type PatchRbacAuthorizationV1alpha1NamespacedRole struct {
	Context *middleware.Context
	Handler PatchRbacAuthorizationV1alpha1NamespacedRoleHandler
}

func (o *PatchRbacAuthorizationV1alpha1NamespacedRole) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPatchRbacAuthorizationV1alpha1NamespacedRoleParams()

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
