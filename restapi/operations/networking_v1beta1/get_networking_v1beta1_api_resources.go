// Code generated by go-swagger; DO NOT EDIT.

package networking_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetNetworkingV1beta1APIResourcesHandlerFunc turns a function with the right signature into a get networking v1beta1 API resources handler
type GetNetworkingV1beta1APIResourcesHandlerFunc func(GetNetworkingV1beta1APIResourcesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetNetworkingV1beta1APIResourcesHandlerFunc) Handle(params GetNetworkingV1beta1APIResourcesParams) middleware.Responder {
	return fn(params)
}

// GetNetworkingV1beta1APIResourcesHandler interface for that can handle valid get networking v1beta1 API resources params
type GetNetworkingV1beta1APIResourcesHandler interface {
	Handle(GetNetworkingV1beta1APIResourcesParams) middleware.Responder
}

// NewGetNetworkingV1beta1APIResources creates a new http.Handler for the get networking v1beta1 API resources operation
func NewGetNetworkingV1beta1APIResources(ctx *middleware.Context, handler GetNetworkingV1beta1APIResourcesHandler) *GetNetworkingV1beta1APIResources {
	return &GetNetworkingV1beta1APIResources{Context: ctx, Handler: handler}
}

/*GetNetworkingV1beta1APIResources swagger:route GET /apis/networking.k8s.io/v1beta1/ networking_v1beta1 getNetworkingV1beta1ApiResources

get available resources

*/
type GetNetworkingV1beta1APIResources struct {
	Context *middleware.Context
	Handler GetNetworkingV1beta1APIResourcesHandler
}

func (o *GetNetworkingV1beta1APIResources) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetNetworkingV1beta1APIResourcesParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
