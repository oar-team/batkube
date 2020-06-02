// Code generated by go-swagger; DO NOT EDIT.

package coordination_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetCoordinationV1APIResourcesHandlerFunc turns a function with the right signature into a get coordination v1 API resources handler
type GetCoordinationV1APIResourcesHandlerFunc func(GetCoordinationV1APIResourcesParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetCoordinationV1APIResourcesHandlerFunc) Handle(params GetCoordinationV1APIResourcesParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetCoordinationV1APIResourcesHandler interface for that can handle valid get coordination v1 API resources params
type GetCoordinationV1APIResourcesHandler interface {
	Handle(GetCoordinationV1APIResourcesParams, interface{}) middleware.Responder
}

// NewGetCoordinationV1APIResources creates a new http.Handler for the get coordination v1 API resources operation
func NewGetCoordinationV1APIResources(ctx *middleware.Context, handler GetCoordinationV1APIResourcesHandler) *GetCoordinationV1APIResources {
	return &GetCoordinationV1APIResources{Context: ctx, Handler: handler}
}

/*GetCoordinationV1APIResources swagger:route GET /apis/coordination.k8s.io/v1/ coordination_v1 getCoordinationV1ApiResources

get available resources

*/
type GetCoordinationV1APIResources struct {
	Context *middleware.Context
	Handler GetCoordinationV1APIResourcesHandler
}

func (o *GetCoordinationV1APIResources) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetCoordinationV1APIResourcesParams()

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
