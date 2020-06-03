// Code generated by go-swagger; DO NOT EDIT.

package extensions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetExtensionsAPIGroupHandlerFunc turns a function with the right signature into a get extensions API group handler
type GetExtensionsAPIGroupHandlerFunc func(GetExtensionsAPIGroupParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetExtensionsAPIGroupHandlerFunc) Handle(params GetExtensionsAPIGroupParams) middleware.Responder {
	return fn(params)
}

// GetExtensionsAPIGroupHandler interface for that can handle valid get extensions API group params
type GetExtensionsAPIGroupHandler interface {
	Handle(GetExtensionsAPIGroupParams) middleware.Responder
}

// NewGetExtensionsAPIGroup creates a new http.Handler for the get extensions API group operation
func NewGetExtensionsAPIGroup(ctx *middleware.Context, handler GetExtensionsAPIGroupHandler) *GetExtensionsAPIGroup {
	return &GetExtensionsAPIGroup{Context: ctx, Handler: handler}
}

/*GetExtensionsAPIGroup swagger:route GET /apis/extensions/ extensions getExtensionsApiGroup

get information of a group

*/
type GetExtensionsAPIGroup struct {
	Context *middleware.Context
	Handler GetExtensionsAPIGroupHandler
}

func (o *GetExtensionsAPIGroup) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetExtensionsAPIGroupParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
