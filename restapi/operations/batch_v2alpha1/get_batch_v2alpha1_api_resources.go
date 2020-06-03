// Code generated by go-swagger; DO NOT EDIT.

package batch_v2alpha1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetBatchV2alpha1APIResourcesHandlerFunc turns a function with the right signature into a get batch v2alpha1 API resources handler
type GetBatchV2alpha1APIResourcesHandlerFunc func(GetBatchV2alpha1APIResourcesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetBatchV2alpha1APIResourcesHandlerFunc) Handle(params GetBatchV2alpha1APIResourcesParams) middleware.Responder {
	return fn(params)
}

// GetBatchV2alpha1APIResourcesHandler interface for that can handle valid get batch v2alpha1 API resources params
type GetBatchV2alpha1APIResourcesHandler interface {
	Handle(GetBatchV2alpha1APIResourcesParams) middleware.Responder
}

// NewGetBatchV2alpha1APIResources creates a new http.Handler for the get batch v2alpha1 API resources operation
func NewGetBatchV2alpha1APIResources(ctx *middleware.Context, handler GetBatchV2alpha1APIResourcesHandler) *GetBatchV2alpha1APIResources {
	return &GetBatchV2alpha1APIResources{Context: ctx, Handler: handler}
}

/*GetBatchV2alpha1APIResources swagger:route GET /apis/batch/v2alpha1/ batch_v2alpha1 getBatchV2alpha1ApiResources

get available resources

*/
type GetBatchV2alpha1APIResources struct {
	Context *middleware.Context
	Handler GetBatchV2alpha1APIResourcesHandler
}

func (o *GetBatchV2alpha1APIResources) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetBatchV2alpha1APIResourcesParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
