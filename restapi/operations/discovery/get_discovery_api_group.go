// Code generated by go-swagger; DO NOT EDIT.

package discovery

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetDiscoveryAPIGroupHandlerFunc turns a function with the right signature into a get discovery API group handler
type GetDiscoveryAPIGroupHandlerFunc func(GetDiscoveryAPIGroupParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetDiscoveryAPIGroupHandlerFunc) Handle(params GetDiscoveryAPIGroupParams) middleware.Responder {
	return fn(params)
}

// GetDiscoveryAPIGroupHandler interface for that can handle valid get discovery API group params
type GetDiscoveryAPIGroupHandler interface {
	Handle(GetDiscoveryAPIGroupParams) middleware.Responder
}

// NewGetDiscoveryAPIGroup creates a new http.Handler for the get discovery API group operation
func NewGetDiscoveryAPIGroup(ctx *middleware.Context, handler GetDiscoveryAPIGroupHandler) *GetDiscoveryAPIGroup {
	return &GetDiscoveryAPIGroup{Context: ctx, Handler: handler}
}

/*GetDiscoveryAPIGroup swagger:route GET /apis/discovery.k8s.io/ discovery getDiscoveryApiGroup

get information of a group

*/
type GetDiscoveryAPIGroup struct {
	Context *middleware.Context
	Handler GetDiscoveryAPIGroupHandler
}

func (o *GetDiscoveryAPIGroup) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetDiscoveryAPIGroupParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
