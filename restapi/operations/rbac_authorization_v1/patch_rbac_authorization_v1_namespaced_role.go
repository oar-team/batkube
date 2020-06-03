// Code generated by go-swagger; DO NOT EDIT.

package rbac_authorization_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PatchRbacAuthorizationV1NamespacedRoleHandlerFunc turns a function with the right signature into a patch rbac authorization v1 namespaced role handler
type PatchRbacAuthorizationV1NamespacedRoleHandlerFunc func(PatchRbacAuthorizationV1NamespacedRoleParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PatchRbacAuthorizationV1NamespacedRoleHandlerFunc) Handle(params PatchRbacAuthorizationV1NamespacedRoleParams) middleware.Responder {
	return fn(params)
}

// PatchRbacAuthorizationV1NamespacedRoleHandler interface for that can handle valid patch rbac authorization v1 namespaced role params
type PatchRbacAuthorizationV1NamespacedRoleHandler interface {
	Handle(PatchRbacAuthorizationV1NamespacedRoleParams) middleware.Responder
}

// NewPatchRbacAuthorizationV1NamespacedRole creates a new http.Handler for the patch rbac authorization v1 namespaced role operation
func NewPatchRbacAuthorizationV1NamespacedRole(ctx *middleware.Context, handler PatchRbacAuthorizationV1NamespacedRoleHandler) *PatchRbacAuthorizationV1NamespacedRole {
	return &PatchRbacAuthorizationV1NamespacedRole{Context: ctx, Handler: handler}
}

/*PatchRbacAuthorizationV1NamespacedRole swagger:route PATCH /apis/rbac.authorization.k8s.io/v1/namespaces/{namespace}/roles/{name} rbacAuthorization_v1 patchRbacAuthorizationV1NamespacedRole

partially update the specified Role

*/
type PatchRbacAuthorizationV1NamespacedRole struct {
	Context *middleware.Context
	Handler PatchRbacAuthorizationV1NamespacedRoleHandler
}

func (o *PatchRbacAuthorizationV1NamespacedRole) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPatchRbacAuthorizationV1NamespacedRoleParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}