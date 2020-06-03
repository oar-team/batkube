// Code generated by go-swagger; DO NOT EDIT.

package policy_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PatchPolicyV1beta1NamespacedPodDisruptionBudgetHandlerFunc turns a function with the right signature into a patch policy v1beta1 namespaced pod disruption budget handler
type PatchPolicyV1beta1NamespacedPodDisruptionBudgetHandlerFunc func(PatchPolicyV1beta1NamespacedPodDisruptionBudgetParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PatchPolicyV1beta1NamespacedPodDisruptionBudgetHandlerFunc) Handle(params PatchPolicyV1beta1NamespacedPodDisruptionBudgetParams) middleware.Responder {
	return fn(params)
}

// PatchPolicyV1beta1NamespacedPodDisruptionBudgetHandler interface for that can handle valid patch policy v1beta1 namespaced pod disruption budget params
type PatchPolicyV1beta1NamespacedPodDisruptionBudgetHandler interface {
	Handle(PatchPolicyV1beta1NamespacedPodDisruptionBudgetParams) middleware.Responder
}

// NewPatchPolicyV1beta1NamespacedPodDisruptionBudget creates a new http.Handler for the patch policy v1beta1 namespaced pod disruption budget operation
func NewPatchPolicyV1beta1NamespacedPodDisruptionBudget(ctx *middleware.Context, handler PatchPolicyV1beta1NamespacedPodDisruptionBudgetHandler) *PatchPolicyV1beta1NamespacedPodDisruptionBudget {
	return &PatchPolicyV1beta1NamespacedPodDisruptionBudget{Context: ctx, Handler: handler}
}

/*PatchPolicyV1beta1NamespacedPodDisruptionBudget swagger:route PATCH /apis/policy/v1beta1/namespaces/{namespace}/poddisruptionbudgets/{name} policy_v1beta1 patchPolicyV1beta1NamespacedPodDisruptionBudget

partially update the specified PodDisruptionBudget

*/
type PatchPolicyV1beta1NamespacedPodDisruptionBudget struct {
	Context *middleware.Context
	Handler PatchPolicyV1beta1NamespacedPodDisruptionBudgetHandler
}

func (o *PatchPolicyV1beta1NamespacedPodDisruptionBudget) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPatchPolicyV1beta1NamespacedPodDisruptionBudgetParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
