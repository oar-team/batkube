// Code generated by go-swagger; DO NOT EDIT.

package rbac_authorization_v1alpha1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteRbacAuthorizationV1alpha1ClusterRoleBindingHandlerFunc turns a function with the right signature into a delete rbac authorization v1alpha1 cluster role binding handler
type DeleteRbacAuthorizationV1alpha1ClusterRoleBindingHandlerFunc func(DeleteRbacAuthorizationV1alpha1ClusterRoleBindingParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteRbacAuthorizationV1alpha1ClusterRoleBindingHandlerFunc) Handle(params DeleteRbacAuthorizationV1alpha1ClusterRoleBindingParams) middleware.Responder {
	return fn(params)
}

// DeleteRbacAuthorizationV1alpha1ClusterRoleBindingHandler interface for that can handle valid delete rbac authorization v1alpha1 cluster role binding params
type DeleteRbacAuthorizationV1alpha1ClusterRoleBindingHandler interface {
	Handle(DeleteRbacAuthorizationV1alpha1ClusterRoleBindingParams) middleware.Responder
}

// NewDeleteRbacAuthorizationV1alpha1ClusterRoleBinding creates a new http.Handler for the delete rbac authorization v1alpha1 cluster role binding operation
func NewDeleteRbacAuthorizationV1alpha1ClusterRoleBinding(ctx *middleware.Context, handler DeleteRbacAuthorizationV1alpha1ClusterRoleBindingHandler) *DeleteRbacAuthorizationV1alpha1ClusterRoleBinding {
	return &DeleteRbacAuthorizationV1alpha1ClusterRoleBinding{Context: ctx, Handler: handler}
}

/*DeleteRbacAuthorizationV1alpha1ClusterRoleBinding swagger:route DELETE /apis/rbac.authorization.k8s.io/v1alpha1/clusterrolebindings/{name} rbacAuthorization_v1alpha1 deleteRbacAuthorizationV1alpha1ClusterRoleBinding

delete a ClusterRoleBinding

*/
type DeleteRbacAuthorizationV1alpha1ClusterRoleBinding struct {
	Context *middleware.Context
	Handler DeleteRbacAuthorizationV1alpha1ClusterRoleBindingHandler
}

func (o *DeleteRbacAuthorizationV1alpha1ClusterRoleBinding) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteRbacAuthorizationV1alpha1ClusterRoleBindingParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
