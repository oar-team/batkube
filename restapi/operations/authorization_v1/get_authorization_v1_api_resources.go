// Code generated by go-swagger; DO NOT EDIT.

package authorization_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetAuthorizationV1APIResourcesHandlerFunc turns a function with the right signature into a get authorization v1 API resources handler
type GetAuthorizationV1APIResourcesHandlerFunc func(GetAuthorizationV1APIResourcesParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetAuthorizationV1APIResourcesHandlerFunc) Handle(params GetAuthorizationV1APIResourcesParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetAuthorizationV1APIResourcesHandler interface for that can handle valid get authorization v1 API resources params
type GetAuthorizationV1APIResourcesHandler interface {
	Handle(GetAuthorizationV1APIResourcesParams, interface{}) middleware.Responder
}

// NewGetAuthorizationV1APIResources creates a new http.Handler for the get authorization v1 API resources operation
func NewGetAuthorizationV1APIResources(ctx *middleware.Context, handler GetAuthorizationV1APIResourcesHandler) *GetAuthorizationV1APIResources {
	return &GetAuthorizationV1APIResources{Context: ctx, Handler: handler}
}

/*GetAuthorizationV1APIResources swagger:route GET /apis/authorization.k8s.io/v1/ authorization_v1 getAuthorizationV1ApiResources

get available resources

*/
type GetAuthorizationV1APIResources struct {
	Context *middleware.Context
	Handler GetAuthorizationV1APIResourcesHandler
}

func (o *GetAuthorizationV1APIResources) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetAuthorizationV1APIResourcesParams()

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
