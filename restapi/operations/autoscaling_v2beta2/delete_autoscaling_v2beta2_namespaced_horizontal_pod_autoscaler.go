// Code generated by go-swagger; DO NOT EDIT.

package autoscaling_v2beta2

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteAutoscalingV2beta2NamespacedHorizontalPodAutoscalerHandlerFunc turns a function with the right signature into a delete autoscaling v2beta2 namespaced horizontal pod autoscaler handler
type DeleteAutoscalingV2beta2NamespacedHorizontalPodAutoscalerHandlerFunc func(DeleteAutoscalingV2beta2NamespacedHorizontalPodAutoscalerParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteAutoscalingV2beta2NamespacedHorizontalPodAutoscalerHandlerFunc) Handle(params DeleteAutoscalingV2beta2NamespacedHorizontalPodAutoscalerParams) middleware.Responder {
	return fn(params)
}

// DeleteAutoscalingV2beta2NamespacedHorizontalPodAutoscalerHandler interface for that can handle valid delete autoscaling v2beta2 namespaced horizontal pod autoscaler params
type DeleteAutoscalingV2beta2NamespacedHorizontalPodAutoscalerHandler interface {
	Handle(DeleteAutoscalingV2beta2NamespacedHorizontalPodAutoscalerParams) middleware.Responder
}

// NewDeleteAutoscalingV2beta2NamespacedHorizontalPodAutoscaler creates a new http.Handler for the delete autoscaling v2beta2 namespaced horizontal pod autoscaler operation
func NewDeleteAutoscalingV2beta2NamespacedHorizontalPodAutoscaler(ctx *middleware.Context, handler DeleteAutoscalingV2beta2NamespacedHorizontalPodAutoscalerHandler) *DeleteAutoscalingV2beta2NamespacedHorizontalPodAutoscaler {
	return &DeleteAutoscalingV2beta2NamespacedHorizontalPodAutoscaler{Context: ctx, Handler: handler}
}

/*DeleteAutoscalingV2beta2NamespacedHorizontalPodAutoscaler swagger:route DELETE /apis/autoscaling/v2beta2/namespaces/{namespace}/horizontalpodautoscalers/{name} autoscaling_v2beta2 deleteAutoscalingV2beta2NamespacedHorizontalPodAutoscaler

delete a HorizontalPodAutoscaler

*/
type DeleteAutoscalingV2beta2NamespacedHorizontalPodAutoscaler struct {
	Context *middleware.Context
	Handler DeleteAutoscalingV2beta2NamespacedHorizontalPodAutoscalerHandler
}

func (o *DeleteAutoscalingV2beta2NamespacedHorizontalPodAutoscaler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteAutoscalingV2beta2NamespacedHorizontalPodAutoscalerParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
