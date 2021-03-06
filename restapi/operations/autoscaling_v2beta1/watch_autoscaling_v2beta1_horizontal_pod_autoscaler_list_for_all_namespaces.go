// Code generated by go-swagger; DO NOT EDIT.

package autoscaling_v2beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// WatchAutoscalingV2beta1HorizontalPodAutoscalerListForAllNamespacesHandlerFunc turns a function with the right signature into a watch autoscaling v2beta1 horizontal pod autoscaler list for all namespaces handler
type WatchAutoscalingV2beta1HorizontalPodAutoscalerListForAllNamespacesHandlerFunc func(WatchAutoscalingV2beta1HorizontalPodAutoscalerListForAllNamespacesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn WatchAutoscalingV2beta1HorizontalPodAutoscalerListForAllNamespacesHandlerFunc) Handle(params WatchAutoscalingV2beta1HorizontalPodAutoscalerListForAllNamespacesParams) middleware.Responder {
	return fn(params)
}

// WatchAutoscalingV2beta1HorizontalPodAutoscalerListForAllNamespacesHandler interface for that can handle valid watch autoscaling v2beta1 horizontal pod autoscaler list for all namespaces params
type WatchAutoscalingV2beta1HorizontalPodAutoscalerListForAllNamespacesHandler interface {
	Handle(WatchAutoscalingV2beta1HorizontalPodAutoscalerListForAllNamespacesParams) middleware.Responder
}

// NewWatchAutoscalingV2beta1HorizontalPodAutoscalerListForAllNamespaces creates a new http.Handler for the watch autoscaling v2beta1 horizontal pod autoscaler list for all namespaces operation
func NewWatchAutoscalingV2beta1HorizontalPodAutoscalerListForAllNamespaces(ctx *middleware.Context, handler WatchAutoscalingV2beta1HorizontalPodAutoscalerListForAllNamespacesHandler) *WatchAutoscalingV2beta1HorizontalPodAutoscalerListForAllNamespaces {
	return &WatchAutoscalingV2beta1HorizontalPodAutoscalerListForAllNamespaces{Context: ctx, Handler: handler}
}

/*WatchAutoscalingV2beta1HorizontalPodAutoscalerListForAllNamespaces swagger:route GET /apis/autoscaling/v2beta1/watch/horizontalpodautoscalers autoscaling_v2beta1 watchAutoscalingV2beta1HorizontalPodAutoscalerListForAllNamespaces

watch individual changes to a list of HorizontalPodAutoscaler. deprecated: use the 'watch' parameter with a list operation instead.

*/
type WatchAutoscalingV2beta1HorizontalPodAutoscalerListForAllNamespaces struct {
	Context *middleware.Context
	Handler WatchAutoscalingV2beta1HorizontalPodAutoscalerListForAllNamespacesHandler
}

func (o *WatchAutoscalingV2beta1HorizontalPodAutoscalerListForAllNamespaces) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewWatchAutoscalingV2beta1HorizontalPodAutoscalerListForAllNamespacesParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
