// Code generated by go-swagger; DO NOT EDIT.

package autoscaling_v2beta2

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// WatchAutoscalingV2beta2HorizontalPodAutoscalerListForAllNamespacesHandlerFunc turns a function with the right signature into a watch autoscaling v2beta2 horizontal pod autoscaler list for all namespaces handler
type WatchAutoscalingV2beta2HorizontalPodAutoscalerListForAllNamespacesHandlerFunc func(WatchAutoscalingV2beta2HorizontalPodAutoscalerListForAllNamespacesParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn WatchAutoscalingV2beta2HorizontalPodAutoscalerListForAllNamespacesHandlerFunc) Handle(params WatchAutoscalingV2beta2HorizontalPodAutoscalerListForAllNamespacesParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// WatchAutoscalingV2beta2HorizontalPodAutoscalerListForAllNamespacesHandler interface for that can handle valid watch autoscaling v2beta2 horizontal pod autoscaler list for all namespaces params
type WatchAutoscalingV2beta2HorizontalPodAutoscalerListForAllNamespacesHandler interface {
	Handle(WatchAutoscalingV2beta2HorizontalPodAutoscalerListForAllNamespacesParams, interface{}) middleware.Responder
}

// NewWatchAutoscalingV2beta2HorizontalPodAutoscalerListForAllNamespaces creates a new http.Handler for the watch autoscaling v2beta2 horizontal pod autoscaler list for all namespaces operation
func NewWatchAutoscalingV2beta2HorizontalPodAutoscalerListForAllNamespaces(ctx *middleware.Context, handler WatchAutoscalingV2beta2HorizontalPodAutoscalerListForAllNamespacesHandler) *WatchAutoscalingV2beta2HorizontalPodAutoscalerListForAllNamespaces {
	return &WatchAutoscalingV2beta2HorizontalPodAutoscalerListForAllNamespaces{Context: ctx, Handler: handler}
}

/*WatchAutoscalingV2beta2HorizontalPodAutoscalerListForAllNamespaces swagger:route GET /apis/autoscaling/v2beta2/watch/horizontalpodautoscalers autoscaling_v2beta2 watchAutoscalingV2beta2HorizontalPodAutoscalerListForAllNamespaces

watch individual changes to a list of HorizontalPodAutoscaler. deprecated: use the 'watch' parameter with a list operation instead.

*/
type WatchAutoscalingV2beta2HorizontalPodAutoscalerListForAllNamespaces struct {
	Context *middleware.Context
	Handler WatchAutoscalingV2beta2HorizontalPodAutoscalerListForAllNamespacesHandler
}

func (o *WatchAutoscalingV2beta2HorizontalPodAutoscalerListForAllNamespaces) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewWatchAutoscalingV2beta2HorizontalPodAutoscalerListForAllNamespacesParams()

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
