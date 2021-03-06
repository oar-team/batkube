// Code generated by go-swagger; DO NOT EDIT.

package autoscaling_v2beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteAutoscalingV2beta1CollectionNamespacedHorizontalPodAutoscalerHandlerFunc turns a function with the right signature into a delete autoscaling v2beta1 collection namespaced horizontal pod autoscaler handler
type DeleteAutoscalingV2beta1CollectionNamespacedHorizontalPodAutoscalerHandlerFunc func(DeleteAutoscalingV2beta1CollectionNamespacedHorizontalPodAutoscalerParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteAutoscalingV2beta1CollectionNamespacedHorizontalPodAutoscalerHandlerFunc) Handle(params DeleteAutoscalingV2beta1CollectionNamespacedHorizontalPodAutoscalerParams) middleware.Responder {
	return fn(params)
}

// DeleteAutoscalingV2beta1CollectionNamespacedHorizontalPodAutoscalerHandler interface for that can handle valid delete autoscaling v2beta1 collection namespaced horizontal pod autoscaler params
type DeleteAutoscalingV2beta1CollectionNamespacedHorizontalPodAutoscalerHandler interface {
	Handle(DeleteAutoscalingV2beta1CollectionNamespacedHorizontalPodAutoscalerParams) middleware.Responder
}

// NewDeleteAutoscalingV2beta1CollectionNamespacedHorizontalPodAutoscaler creates a new http.Handler for the delete autoscaling v2beta1 collection namespaced horizontal pod autoscaler operation
func NewDeleteAutoscalingV2beta1CollectionNamespacedHorizontalPodAutoscaler(ctx *middleware.Context, handler DeleteAutoscalingV2beta1CollectionNamespacedHorizontalPodAutoscalerHandler) *DeleteAutoscalingV2beta1CollectionNamespacedHorizontalPodAutoscaler {
	return &DeleteAutoscalingV2beta1CollectionNamespacedHorizontalPodAutoscaler{Context: ctx, Handler: handler}
}

/*DeleteAutoscalingV2beta1CollectionNamespacedHorizontalPodAutoscaler swagger:route DELETE /apis/autoscaling/v2beta1/namespaces/{namespace}/horizontalpodautoscalers autoscaling_v2beta1 deleteAutoscalingV2beta1CollectionNamespacedHorizontalPodAutoscaler

delete collection of HorizontalPodAutoscaler

*/
type DeleteAutoscalingV2beta1CollectionNamespacedHorizontalPodAutoscaler struct {
	Context *middleware.Context
	Handler DeleteAutoscalingV2beta1CollectionNamespacedHorizontalPodAutoscalerHandler
}

func (o *DeleteAutoscalingV2beta1CollectionNamespacedHorizontalPodAutoscaler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteAutoscalingV2beta1CollectionNamespacedHorizontalPodAutoscalerParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
