// Code generated by go-swagger; DO NOT EDIT.

package autoscaling_v2beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PatchAutoscalingV2beta1NamespacedHorizontalPodAutoscalerHandlerFunc turns a function with the right signature into a patch autoscaling v2beta1 namespaced horizontal pod autoscaler handler
type PatchAutoscalingV2beta1NamespacedHorizontalPodAutoscalerHandlerFunc func(PatchAutoscalingV2beta1NamespacedHorizontalPodAutoscalerParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PatchAutoscalingV2beta1NamespacedHorizontalPodAutoscalerHandlerFunc) Handle(params PatchAutoscalingV2beta1NamespacedHorizontalPodAutoscalerParams) middleware.Responder {
	return fn(params)
}

// PatchAutoscalingV2beta1NamespacedHorizontalPodAutoscalerHandler interface for that can handle valid patch autoscaling v2beta1 namespaced horizontal pod autoscaler params
type PatchAutoscalingV2beta1NamespacedHorizontalPodAutoscalerHandler interface {
	Handle(PatchAutoscalingV2beta1NamespacedHorizontalPodAutoscalerParams) middleware.Responder
}

// NewPatchAutoscalingV2beta1NamespacedHorizontalPodAutoscaler creates a new http.Handler for the patch autoscaling v2beta1 namespaced horizontal pod autoscaler operation
func NewPatchAutoscalingV2beta1NamespacedHorizontalPodAutoscaler(ctx *middleware.Context, handler PatchAutoscalingV2beta1NamespacedHorizontalPodAutoscalerHandler) *PatchAutoscalingV2beta1NamespacedHorizontalPodAutoscaler {
	return &PatchAutoscalingV2beta1NamespacedHorizontalPodAutoscaler{Context: ctx, Handler: handler}
}

/*PatchAutoscalingV2beta1NamespacedHorizontalPodAutoscaler swagger:route PATCH /apis/autoscaling/v2beta1/namespaces/{namespace}/horizontalpodautoscalers/{name} autoscaling_v2beta1 patchAutoscalingV2beta1NamespacedHorizontalPodAutoscaler

partially update the specified HorizontalPodAutoscaler

*/
type PatchAutoscalingV2beta1NamespacedHorizontalPodAutoscaler struct {
	Context *middleware.Context
	Handler PatchAutoscalingV2beta1NamespacedHorizontalPodAutoscalerHandler
}

func (o *PatchAutoscalingV2beta1NamespacedHorizontalPodAutoscaler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPatchAutoscalingV2beta1NamespacedHorizontalPodAutoscalerParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}