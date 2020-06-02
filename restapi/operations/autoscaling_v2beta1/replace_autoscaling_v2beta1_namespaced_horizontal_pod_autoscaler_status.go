// Code generated by go-swagger; DO NOT EDIT.

package autoscaling_v2beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ReplaceAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusHandlerFunc turns a function with the right signature into a replace autoscaling v2beta1 namespaced horizontal pod autoscaler status handler
type ReplaceAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusHandlerFunc func(ReplaceAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn ReplaceAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusHandlerFunc) Handle(params ReplaceAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// ReplaceAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusHandler interface for that can handle valid replace autoscaling v2beta1 namespaced horizontal pod autoscaler status params
type ReplaceAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusHandler interface {
	Handle(ReplaceAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusParams, interface{}) middleware.Responder
}

// NewReplaceAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatus creates a new http.Handler for the replace autoscaling v2beta1 namespaced horizontal pod autoscaler status operation
func NewReplaceAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatus(ctx *middleware.Context, handler ReplaceAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusHandler) *ReplaceAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatus {
	return &ReplaceAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatus{Context: ctx, Handler: handler}
}

/*ReplaceAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatus swagger:route PUT /apis/autoscaling/v2beta1/namespaces/{namespace}/horizontalpodautoscalers/{name}/status autoscaling_v2beta1 replaceAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatus

replace status of the specified HorizontalPodAutoscaler

*/
type ReplaceAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatus struct {
	Context *middleware.Context
	Handler ReplaceAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusHandler
}

func (o *ReplaceAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatus) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewReplaceAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusParams()

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
