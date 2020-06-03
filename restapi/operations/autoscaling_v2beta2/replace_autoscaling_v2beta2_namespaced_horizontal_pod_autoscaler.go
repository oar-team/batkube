// Code generated by go-swagger; DO NOT EDIT.

package autoscaling_v2beta2

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ReplaceAutoscalingV2beta2NamespacedHorizontalPodAutoscalerHandlerFunc turns a function with the right signature into a replace autoscaling v2beta2 namespaced horizontal pod autoscaler handler
type ReplaceAutoscalingV2beta2NamespacedHorizontalPodAutoscalerHandlerFunc func(ReplaceAutoscalingV2beta2NamespacedHorizontalPodAutoscalerParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ReplaceAutoscalingV2beta2NamespacedHorizontalPodAutoscalerHandlerFunc) Handle(params ReplaceAutoscalingV2beta2NamespacedHorizontalPodAutoscalerParams) middleware.Responder {
	return fn(params)
}

// ReplaceAutoscalingV2beta2NamespacedHorizontalPodAutoscalerHandler interface for that can handle valid replace autoscaling v2beta2 namespaced horizontal pod autoscaler params
type ReplaceAutoscalingV2beta2NamespacedHorizontalPodAutoscalerHandler interface {
	Handle(ReplaceAutoscalingV2beta2NamespacedHorizontalPodAutoscalerParams) middleware.Responder
}

// NewReplaceAutoscalingV2beta2NamespacedHorizontalPodAutoscaler creates a new http.Handler for the replace autoscaling v2beta2 namespaced horizontal pod autoscaler operation
func NewReplaceAutoscalingV2beta2NamespacedHorizontalPodAutoscaler(ctx *middleware.Context, handler ReplaceAutoscalingV2beta2NamespacedHorizontalPodAutoscalerHandler) *ReplaceAutoscalingV2beta2NamespacedHorizontalPodAutoscaler {
	return &ReplaceAutoscalingV2beta2NamespacedHorizontalPodAutoscaler{Context: ctx, Handler: handler}
}

/*ReplaceAutoscalingV2beta2NamespacedHorizontalPodAutoscaler swagger:route PUT /apis/autoscaling/v2beta2/namespaces/{namespace}/horizontalpodautoscalers/{name} autoscaling_v2beta2 replaceAutoscalingV2beta2NamespacedHorizontalPodAutoscaler

replace the specified HorizontalPodAutoscaler

*/
type ReplaceAutoscalingV2beta2NamespacedHorizontalPodAutoscaler struct {
	Context *middleware.Context
	Handler ReplaceAutoscalingV2beta2NamespacedHorizontalPodAutoscalerHandler
}

func (o *ReplaceAutoscalingV2beta2NamespacedHorizontalPodAutoscaler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewReplaceAutoscalingV2beta2NamespacedHorizontalPodAutoscalerParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}