// Code generated by go-swagger; DO NOT EDIT.

package policy_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PatchPolicyV1beta1NamespacedPodDisruptionBudgetStatusHandlerFunc turns a function with the right signature into a patch policy v1beta1 namespaced pod disruption budget status handler
type PatchPolicyV1beta1NamespacedPodDisruptionBudgetStatusHandlerFunc func(PatchPolicyV1beta1NamespacedPodDisruptionBudgetStatusParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PatchPolicyV1beta1NamespacedPodDisruptionBudgetStatusHandlerFunc) Handle(params PatchPolicyV1beta1NamespacedPodDisruptionBudgetStatusParams) middleware.Responder {
	return fn(params)
}

// PatchPolicyV1beta1NamespacedPodDisruptionBudgetStatusHandler interface for that can handle valid patch policy v1beta1 namespaced pod disruption budget status params
type PatchPolicyV1beta1NamespacedPodDisruptionBudgetStatusHandler interface {
	Handle(PatchPolicyV1beta1NamespacedPodDisruptionBudgetStatusParams) middleware.Responder
}

// NewPatchPolicyV1beta1NamespacedPodDisruptionBudgetStatus creates a new http.Handler for the patch policy v1beta1 namespaced pod disruption budget status operation
func NewPatchPolicyV1beta1NamespacedPodDisruptionBudgetStatus(ctx *middleware.Context, handler PatchPolicyV1beta1NamespacedPodDisruptionBudgetStatusHandler) *PatchPolicyV1beta1NamespacedPodDisruptionBudgetStatus {
	return &PatchPolicyV1beta1NamespacedPodDisruptionBudgetStatus{Context: ctx, Handler: handler}
}

/*PatchPolicyV1beta1NamespacedPodDisruptionBudgetStatus swagger:route PATCH /apis/policy/v1beta1/namespaces/{namespace}/poddisruptionbudgets/{name}/status policy_v1beta1 patchPolicyV1beta1NamespacedPodDisruptionBudgetStatus

partially update status of the specified PodDisruptionBudget

*/
type PatchPolicyV1beta1NamespacedPodDisruptionBudgetStatus struct {
	Context *middleware.Context
	Handler PatchPolicyV1beta1NamespacedPodDisruptionBudgetStatusHandler
}

func (o *PatchPolicyV1beta1NamespacedPodDisruptionBudgetStatus) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPatchPolicyV1beta1NamespacedPodDisruptionBudgetStatusParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}