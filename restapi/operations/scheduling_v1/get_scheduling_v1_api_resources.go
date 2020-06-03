// Code generated by go-swagger; DO NOT EDIT.

package scheduling_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetSchedulingV1APIResourcesHandlerFunc turns a function with the right signature into a get scheduling v1 API resources handler
type GetSchedulingV1APIResourcesHandlerFunc func(GetSchedulingV1APIResourcesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetSchedulingV1APIResourcesHandlerFunc) Handle(params GetSchedulingV1APIResourcesParams) middleware.Responder {
	return fn(params)
}

// GetSchedulingV1APIResourcesHandler interface for that can handle valid get scheduling v1 API resources params
type GetSchedulingV1APIResourcesHandler interface {
	Handle(GetSchedulingV1APIResourcesParams) middleware.Responder
}

// NewGetSchedulingV1APIResources creates a new http.Handler for the get scheduling v1 API resources operation
func NewGetSchedulingV1APIResources(ctx *middleware.Context, handler GetSchedulingV1APIResourcesHandler) *GetSchedulingV1APIResources {
	return &GetSchedulingV1APIResources{Context: ctx, Handler: handler}
}

/*GetSchedulingV1APIResources swagger:route GET /apis/scheduling.k8s.io/v1/ scheduling_v1 getSchedulingV1ApiResources

get available resources

*/
type GetSchedulingV1APIResources struct {
	Context *middleware.Context
	Handler GetSchedulingV1APIResourcesHandler
}

func (o *GetSchedulingV1APIResources) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetSchedulingV1APIResourcesParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
