// Code generated by go-swagger; DO NOT EDIT.

package autoscaling_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ListAutoscalingV1NamespacedHorizontalPodAutoscalerHandlerFunc turns a function with the right signature into a list autoscaling v1 namespaced horizontal pod autoscaler handler
type ListAutoscalingV1NamespacedHorizontalPodAutoscalerHandlerFunc func(ListAutoscalingV1NamespacedHorizontalPodAutoscalerParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ListAutoscalingV1NamespacedHorizontalPodAutoscalerHandlerFunc) Handle(params ListAutoscalingV1NamespacedHorizontalPodAutoscalerParams) middleware.Responder {
	return fn(params)
}

// ListAutoscalingV1NamespacedHorizontalPodAutoscalerHandler interface for that can handle valid list autoscaling v1 namespaced horizontal pod autoscaler params
type ListAutoscalingV1NamespacedHorizontalPodAutoscalerHandler interface {
	Handle(ListAutoscalingV1NamespacedHorizontalPodAutoscalerParams) middleware.Responder
}

// NewListAutoscalingV1NamespacedHorizontalPodAutoscaler creates a new http.Handler for the list autoscaling v1 namespaced horizontal pod autoscaler operation
func NewListAutoscalingV1NamespacedHorizontalPodAutoscaler(ctx *middleware.Context, handler ListAutoscalingV1NamespacedHorizontalPodAutoscalerHandler) *ListAutoscalingV1NamespacedHorizontalPodAutoscaler {
	return &ListAutoscalingV1NamespacedHorizontalPodAutoscaler{Context: ctx, Handler: handler}
}

/*ListAutoscalingV1NamespacedHorizontalPodAutoscaler swagger:route GET /apis/autoscaling/v1/namespaces/{namespace}/horizontalpodautoscalers autoscaling_v1 listAutoscalingV1NamespacedHorizontalPodAutoscaler

list or watch objects of kind HorizontalPodAutoscaler

*/
type ListAutoscalingV1NamespacedHorizontalPodAutoscaler struct {
	Context *middleware.Context
	Handler ListAutoscalingV1NamespacedHorizontalPodAutoscalerHandler
}

func (o *ListAutoscalingV1NamespacedHorizontalPodAutoscaler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewListAutoscalingV1NamespacedHorizontalPodAutoscalerParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
