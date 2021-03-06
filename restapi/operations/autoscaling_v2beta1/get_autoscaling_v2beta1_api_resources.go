// Code generated by go-swagger; DO NOT EDIT.

package autoscaling_v2beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetAutoscalingV2beta1APIResourcesHandlerFunc turns a function with the right signature into a get autoscaling v2beta1 API resources handler
type GetAutoscalingV2beta1APIResourcesHandlerFunc func(GetAutoscalingV2beta1APIResourcesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetAutoscalingV2beta1APIResourcesHandlerFunc) Handle(params GetAutoscalingV2beta1APIResourcesParams) middleware.Responder {
	return fn(params)
}

// GetAutoscalingV2beta1APIResourcesHandler interface for that can handle valid get autoscaling v2beta1 API resources params
type GetAutoscalingV2beta1APIResourcesHandler interface {
	Handle(GetAutoscalingV2beta1APIResourcesParams) middleware.Responder
}

// NewGetAutoscalingV2beta1APIResources creates a new http.Handler for the get autoscaling v2beta1 API resources operation
func NewGetAutoscalingV2beta1APIResources(ctx *middleware.Context, handler GetAutoscalingV2beta1APIResourcesHandler) *GetAutoscalingV2beta1APIResources {
	return &GetAutoscalingV2beta1APIResources{Context: ctx, Handler: handler}
}

/*GetAutoscalingV2beta1APIResources swagger:route GET /apis/autoscaling/v2beta1/ autoscaling_v2beta1 getAutoscalingV2beta1ApiResources

get available resources

*/
type GetAutoscalingV2beta1APIResources struct {
	Context *middleware.Context
	Handler GetAutoscalingV2beta1APIResourcesHandler
}

func (o *GetAutoscalingV2beta1APIResources) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetAutoscalingV2beta1APIResourcesParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
