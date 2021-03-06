// Code generated by go-swagger; DO NOT EDIT.

package apps_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetAppsV1APIResourcesHandlerFunc turns a function with the right signature into a get apps v1 API resources handler
type GetAppsV1APIResourcesHandlerFunc func(GetAppsV1APIResourcesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetAppsV1APIResourcesHandlerFunc) Handle(params GetAppsV1APIResourcesParams) middleware.Responder {
	return fn(params)
}

// GetAppsV1APIResourcesHandler interface for that can handle valid get apps v1 API resources params
type GetAppsV1APIResourcesHandler interface {
	Handle(GetAppsV1APIResourcesParams) middleware.Responder
}

// NewGetAppsV1APIResources creates a new http.Handler for the get apps v1 API resources operation
func NewGetAppsV1APIResources(ctx *middleware.Context, handler GetAppsV1APIResourcesHandler) *GetAppsV1APIResources {
	return &GetAppsV1APIResources{Context: ctx, Handler: handler}
}

/*GetAppsV1APIResources swagger:route GET /apis/apps/v1/ apps_v1 getAppsV1ApiResources

get available resources

*/
type GetAppsV1APIResources struct {
	Context *middleware.Context
	Handler GetAppsV1APIResourcesHandler
}

func (o *GetAppsV1APIResources) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetAppsV1APIResourcesParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
