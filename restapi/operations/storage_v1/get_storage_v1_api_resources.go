// Code generated by go-swagger; DO NOT EDIT.

package storage_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetStorageV1APIResourcesHandlerFunc turns a function with the right signature into a get storage v1 API resources handler
type GetStorageV1APIResourcesHandlerFunc func(GetStorageV1APIResourcesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetStorageV1APIResourcesHandlerFunc) Handle(params GetStorageV1APIResourcesParams) middleware.Responder {
	return fn(params)
}

// GetStorageV1APIResourcesHandler interface for that can handle valid get storage v1 API resources params
type GetStorageV1APIResourcesHandler interface {
	Handle(GetStorageV1APIResourcesParams) middleware.Responder
}

// NewGetStorageV1APIResources creates a new http.Handler for the get storage v1 API resources operation
func NewGetStorageV1APIResources(ctx *middleware.Context, handler GetStorageV1APIResourcesHandler) *GetStorageV1APIResources {
	return &GetStorageV1APIResources{Context: ctx, Handler: handler}
}

/*GetStorageV1APIResources swagger:route GET /apis/storage.k8s.io/v1/ storage_v1 getStorageV1ApiResources

get available resources

*/
type GetStorageV1APIResources struct {
	Context *middleware.Context
	Handler GetStorageV1APIResourcesHandler
}

func (o *GetStorageV1APIResources) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetStorageV1APIResourcesParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
