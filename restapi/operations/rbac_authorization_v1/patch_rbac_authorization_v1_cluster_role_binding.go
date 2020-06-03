// Code generated by go-swagger; DO NOT EDIT.

package rbac_authorization_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PatchRbacAuthorizationV1ClusterRoleBindingHandlerFunc turns a function with the right signature into a patch rbac authorization v1 cluster role binding handler
type PatchRbacAuthorizationV1ClusterRoleBindingHandlerFunc func(PatchRbacAuthorizationV1ClusterRoleBindingParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PatchRbacAuthorizationV1ClusterRoleBindingHandlerFunc) Handle(params PatchRbacAuthorizationV1ClusterRoleBindingParams) middleware.Responder {
	return fn(params)
}

// PatchRbacAuthorizationV1ClusterRoleBindingHandler interface for that can handle valid patch rbac authorization v1 cluster role binding params
type PatchRbacAuthorizationV1ClusterRoleBindingHandler interface {
	Handle(PatchRbacAuthorizationV1ClusterRoleBindingParams) middleware.Responder
}

// NewPatchRbacAuthorizationV1ClusterRoleBinding creates a new http.Handler for the patch rbac authorization v1 cluster role binding operation
func NewPatchRbacAuthorizationV1ClusterRoleBinding(ctx *middleware.Context, handler PatchRbacAuthorizationV1ClusterRoleBindingHandler) *PatchRbacAuthorizationV1ClusterRoleBinding {
	return &PatchRbacAuthorizationV1ClusterRoleBinding{Context: ctx, Handler: handler}
}

/*PatchRbacAuthorizationV1ClusterRoleBinding swagger:route PATCH /apis/rbac.authorization.k8s.io/v1/clusterrolebindings/{name} rbacAuthorization_v1 patchRbacAuthorizationV1ClusterRoleBinding

partially update the specified ClusterRoleBinding

*/
type PatchRbacAuthorizationV1ClusterRoleBinding struct {
	Context *middleware.Context
	Handler PatchRbacAuthorizationV1ClusterRoleBindingHandler
}

func (o *PatchRbacAuthorizationV1ClusterRoleBinding) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPatchRbacAuthorizationV1ClusterRoleBindingParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
