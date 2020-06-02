// Code generated by go-swagger; DO NOT EDIT.

package apis

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetAPIVersionsHandlerFunc turns a function with the right signature into a get API versions handler
type GetAPIVersionsHandlerFunc func(GetAPIVersionsParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetAPIVersionsHandlerFunc) Handle(params GetAPIVersionsParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetAPIVersionsHandler interface for that can handle valid get API versions params
type GetAPIVersionsHandler interface {
	Handle(GetAPIVersionsParams, interface{}) middleware.Responder
}

// NewGetAPIVersions creates a new http.Handler for the get API versions operation
func NewGetAPIVersions(ctx *middleware.Context, handler GetAPIVersionsHandler) *GetAPIVersions {
	return &GetAPIVersions{Context: ctx, Handler: handler}
}

/*GetAPIVersions swagger:route GET /apis/ apis getApiVersions

get available API versions

*/
type GetAPIVersions struct {
	Context *middleware.Context
	Handler GetAPIVersionsHandler
}

func (o *GetAPIVersions) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetAPIVersionsParams()

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
