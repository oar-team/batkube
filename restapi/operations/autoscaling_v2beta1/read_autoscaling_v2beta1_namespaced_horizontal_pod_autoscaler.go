// Code generated by go-swagger; DO NOT EDIT.

package autoscaling_v2beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ReadAutoscalingV2beta1NamespacedHorizontalPodAutoscalerHandlerFunc turns a function with the right signature into a read autoscaling v2beta1 namespaced horizontal pod autoscaler handler
type ReadAutoscalingV2beta1NamespacedHorizontalPodAutoscalerHandlerFunc func(ReadAutoscalingV2beta1NamespacedHorizontalPodAutoscalerParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ReadAutoscalingV2beta1NamespacedHorizontalPodAutoscalerHandlerFunc) Handle(params ReadAutoscalingV2beta1NamespacedHorizontalPodAutoscalerParams) middleware.Responder {
	return fn(params)
}

// ReadAutoscalingV2beta1NamespacedHorizontalPodAutoscalerHandler interface for that can handle valid read autoscaling v2beta1 namespaced horizontal pod autoscaler params
type ReadAutoscalingV2beta1NamespacedHorizontalPodAutoscalerHandler interface {
	Handle(ReadAutoscalingV2beta1NamespacedHorizontalPodAutoscalerParams) middleware.Responder
}

// NewReadAutoscalingV2beta1NamespacedHorizontalPodAutoscaler creates a new http.Handler for the read autoscaling v2beta1 namespaced horizontal pod autoscaler operation
func NewReadAutoscalingV2beta1NamespacedHorizontalPodAutoscaler(ctx *middleware.Context, handler ReadAutoscalingV2beta1NamespacedHorizontalPodAutoscalerHandler) *ReadAutoscalingV2beta1NamespacedHorizontalPodAutoscaler {
	return &ReadAutoscalingV2beta1NamespacedHorizontalPodAutoscaler{Context: ctx, Handler: handler}
}

/*ReadAutoscalingV2beta1NamespacedHorizontalPodAutoscaler swagger:route GET /apis/autoscaling/v2beta1/namespaces/{namespace}/horizontalpodautoscalers/{name} autoscaling_v2beta1 readAutoscalingV2beta1NamespacedHorizontalPodAutoscaler

read the specified HorizontalPodAutoscaler

*/
type ReadAutoscalingV2beta1NamespacedHorizontalPodAutoscaler struct {
	Context *middleware.Context
	Handler ReadAutoscalingV2beta1NamespacedHorizontalPodAutoscalerHandler
}

func (o *ReadAutoscalingV2beta1NamespacedHorizontalPodAutoscaler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewReadAutoscalingV2beta1NamespacedHorizontalPodAutoscalerParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
