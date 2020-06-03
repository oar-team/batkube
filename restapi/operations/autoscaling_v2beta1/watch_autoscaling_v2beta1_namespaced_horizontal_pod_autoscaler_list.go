// Code generated by go-swagger; DO NOT EDIT.

package autoscaling_v2beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// WatchAutoscalingV2beta1NamespacedHorizontalPodAutoscalerListHandlerFunc turns a function with the right signature into a watch autoscaling v2beta1 namespaced horizontal pod autoscaler list handler
type WatchAutoscalingV2beta1NamespacedHorizontalPodAutoscalerListHandlerFunc func(WatchAutoscalingV2beta1NamespacedHorizontalPodAutoscalerListParams) middleware.Responder

// Handle executing the request and returning a response
func (fn WatchAutoscalingV2beta1NamespacedHorizontalPodAutoscalerListHandlerFunc) Handle(params WatchAutoscalingV2beta1NamespacedHorizontalPodAutoscalerListParams) middleware.Responder {
	return fn(params)
}

// WatchAutoscalingV2beta1NamespacedHorizontalPodAutoscalerListHandler interface for that can handle valid watch autoscaling v2beta1 namespaced horizontal pod autoscaler list params
type WatchAutoscalingV2beta1NamespacedHorizontalPodAutoscalerListHandler interface {
	Handle(WatchAutoscalingV2beta1NamespacedHorizontalPodAutoscalerListParams) middleware.Responder
}

// NewWatchAutoscalingV2beta1NamespacedHorizontalPodAutoscalerList creates a new http.Handler for the watch autoscaling v2beta1 namespaced horizontal pod autoscaler list operation
func NewWatchAutoscalingV2beta1NamespacedHorizontalPodAutoscalerList(ctx *middleware.Context, handler WatchAutoscalingV2beta1NamespacedHorizontalPodAutoscalerListHandler) *WatchAutoscalingV2beta1NamespacedHorizontalPodAutoscalerList {
	return &WatchAutoscalingV2beta1NamespacedHorizontalPodAutoscalerList{Context: ctx, Handler: handler}
}

/*WatchAutoscalingV2beta1NamespacedHorizontalPodAutoscalerList swagger:route GET /apis/autoscaling/v2beta1/watch/namespaces/{namespace}/horizontalpodautoscalers autoscaling_v2beta1 watchAutoscalingV2beta1NamespacedHorizontalPodAutoscalerList

watch individual changes to a list of HorizontalPodAutoscaler. deprecated: use the 'watch' parameter with a list operation instead.

*/
type WatchAutoscalingV2beta1NamespacedHorizontalPodAutoscalerList struct {
	Context *middleware.Context
	Handler WatchAutoscalingV2beta1NamespacedHorizontalPodAutoscalerListHandler
}

func (o *WatchAutoscalingV2beta1NamespacedHorizontalPodAutoscalerList) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewWatchAutoscalingV2beta1NamespacedHorizontalPodAutoscalerListParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}