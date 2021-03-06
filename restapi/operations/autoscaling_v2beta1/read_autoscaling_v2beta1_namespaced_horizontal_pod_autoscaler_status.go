// Code generated by go-swagger; DO NOT EDIT.

package autoscaling_v2beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ReadAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusHandlerFunc turns a function with the right signature into a read autoscaling v2beta1 namespaced horizontal pod autoscaler status handler
type ReadAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusHandlerFunc func(ReadAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ReadAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusHandlerFunc) Handle(params ReadAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusParams) middleware.Responder {
	return fn(params)
}

// ReadAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusHandler interface for that can handle valid read autoscaling v2beta1 namespaced horizontal pod autoscaler status params
type ReadAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusHandler interface {
	Handle(ReadAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusParams) middleware.Responder
}

// NewReadAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatus creates a new http.Handler for the read autoscaling v2beta1 namespaced horizontal pod autoscaler status operation
func NewReadAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatus(ctx *middleware.Context, handler ReadAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusHandler) *ReadAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatus {
	return &ReadAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatus{Context: ctx, Handler: handler}
}

/*ReadAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatus swagger:route GET /apis/autoscaling/v2beta1/namespaces/{namespace}/horizontalpodautoscalers/{name}/status autoscaling_v2beta1 readAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatus

read status of the specified HorizontalPodAutoscaler

*/
type ReadAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatus struct {
	Context *middleware.Context
	Handler ReadAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusHandler
}

func (o *ReadAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatus) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewReadAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
