// Code generated by go-swagger; DO NOT EDIT.

package autoscaling_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// WatchAutoscalingV1NamespacedHorizontalPodAutoscalerListHandlerFunc turns a function with the right signature into a watch autoscaling v1 namespaced horizontal pod autoscaler list handler
type WatchAutoscalingV1NamespacedHorizontalPodAutoscalerListHandlerFunc func(WatchAutoscalingV1NamespacedHorizontalPodAutoscalerListParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn WatchAutoscalingV1NamespacedHorizontalPodAutoscalerListHandlerFunc) Handle(params WatchAutoscalingV1NamespacedHorizontalPodAutoscalerListParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// WatchAutoscalingV1NamespacedHorizontalPodAutoscalerListHandler interface for that can handle valid watch autoscaling v1 namespaced horizontal pod autoscaler list params
type WatchAutoscalingV1NamespacedHorizontalPodAutoscalerListHandler interface {
	Handle(WatchAutoscalingV1NamespacedHorizontalPodAutoscalerListParams, interface{}) middleware.Responder
}

// NewWatchAutoscalingV1NamespacedHorizontalPodAutoscalerList creates a new http.Handler for the watch autoscaling v1 namespaced horizontal pod autoscaler list operation
func NewWatchAutoscalingV1NamespacedHorizontalPodAutoscalerList(ctx *middleware.Context, handler WatchAutoscalingV1NamespacedHorizontalPodAutoscalerListHandler) *WatchAutoscalingV1NamespacedHorizontalPodAutoscalerList {
	return &WatchAutoscalingV1NamespacedHorizontalPodAutoscalerList{Context: ctx, Handler: handler}
}

/*WatchAutoscalingV1NamespacedHorizontalPodAutoscalerList swagger:route GET /apis/autoscaling/v1/watch/namespaces/{namespace}/horizontalpodautoscalers autoscaling_v1 watchAutoscalingV1NamespacedHorizontalPodAutoscalerList

watch individual changes to a list of HorizontalPodAutoscaler. deprecated: use the 'watch' parameter with a list operation instead.

*/
type WatchAutoscalingV1NamespacedHorizontalPodAutoscalerList struct {
	Context *middleware.Context
	Handler WatchAutoscalingV1NamespacedHorizontalPodAutoscalerListHandler
}

func (o *WatchAutoscalingV1NamespacedHorizontalPodAutoscalerList) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewWatchAutoscalingV1NamespacedHorizontalPodAutoscalerListParams()

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
