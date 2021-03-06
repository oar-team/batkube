// Code generated by go-swagger; DO NOT EDIT.

package core

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetCoreAPIVersionsHandlerFunc turns a function with the right signature into a get core API versions handler
type GetCoreAPIVersionsHandlerFunc func(GetCoreAPIVersionsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetCoreAPIVersionsHandlerFunc) Handle(params GetCoreAPIVersionsParams) middleware.Responder {
	return fn(params)
}

// GetCoreAPIVersionsHandler interface for that can handle valid get core API versions params
type GetCoreAPIVersionsHandler interface {
	Handle(GetCoreAPIVersionsParams) middleware.Responder
}

// NewGetCoreAPIVersions creates a new http.Handler for the get core API versions operation
func NewGetCoreAPIVersions(ctx *middleware.Context, handler GetCoreAPIVersionsHandler) *GetCoreAPIVersions {
	return &GetCoreAPIVersions{Context: ctx, Handler: handler}
}

/*GetCoreAPIVersions swagger:route GET /api/ core getCoreApiVersions

get available API versions

*/
type GetCoreAPIVersions struct {
	Context *middleware.Context
	Handler GetCoreAPIVersionsHandler
}

func (o *GetCoreAPIVersions) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetCoreAPIVersionsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
