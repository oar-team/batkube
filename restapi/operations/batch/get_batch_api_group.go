// Code generated by go-swagger; DO NOT EDIT.

package batch

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetBatchAPIGroupHandlerFunc turns a function with the right signature into a get batch API group handler
type GetBatchAPIGroupHandlerFunc func(GetBatchAPIGroupParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetBatchAPIGroupHandlerFunc) Handle(params GetBatchAPIGroupParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetBatchAPIGroupHandler interface for that can handle valid get batch API group params
type GetBatchAPIGroupHandler interface {
	Handle(GetBatchAPIGroupParams, interface{}) middleware.Responder
}

// NewGetBatchAPIGroup creates a new http.Handler for the get batch API group operation
func NewGetBatchAPIGroup(ctx *middleware.Context, handler GetBatchAPIGroupHandler) *GetBatchAPIGroup {
	return &GetBatchAPIGroup{Context: ctx, Handler: handler}
}

/*GetBatchAPIGroup swagger:route GET /apis/batch/ batch getBatchApiGroup

get information of a group

*/
type GetBatchAPIGroup struct {
	Context *middleware.Context
	Handler GetBatchAPIGroupHandler
}

func (o *GetBatchAPIGroup) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetBatchAPIGroupParams()

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
