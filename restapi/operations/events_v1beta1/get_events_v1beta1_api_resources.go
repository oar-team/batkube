// Code generated by go-swagger; DO NOT EDIT.

package events_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetEventsV1beta1APIResourcesHandlerFunc turns a function with the right signature into a get events v1beta1 API resources handler
type GetEventsV1beta1APIResourcesHandlerFunc func(GetEventsV1beta1APIResourcesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetEventsV1beta1APIResourcesHandlerFunc) Handle(params GetEventsV1beta1APIResourcesParams) middleware.Responder {
	return fn(params)
}

// GetEventsV1beta1APIResourcesHandler interface for that can handle valid get events v1beta1 API resources params
type GetEventsV1beta1APIResourcesHandler interface {
	Handle(GetEventsV1beta1APIResourcesParams) middleware.Responder
}

// NewGetEventsV1beta1APIResources creates a new http.Handler for the get events v1beta1 API resources operation
func NewGetEventsV1beta1APIResources(ctx *middleware.Context, handler GetEventsV1beta1APIResourcesHandler) *GetEventsV1beta1APIResources {
	return &GetEventsV1beta1APIResources{Context: ctx, Handler: handler}
}

/*GetEventsV1beta1APIResources swagger:route GET /apis/events.k8s.io/v1beta1/ events_v1beta1 getEventsV1beta1ApiResources

get available resources

*/
type GetEventsV1beta1APIResources struct {
	Context *middleware.Context
	Handler GetEventsV1beta1APIResourcesHandler
}

func (o *GetEventsV1beta1APIResources) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetEventsV1beta1APIResourcesParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
