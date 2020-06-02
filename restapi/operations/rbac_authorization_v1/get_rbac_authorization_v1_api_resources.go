// Code generated by go-swagger; DO NOT EDIT.

package rbac_authorization_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetRbacAuthorizationV1APIResourcesHandlerFunc turns a function with the right signature into a get rbac authorization v1 API resources handler
type GetRbacAuthorizationV1APIResourcesHandlerFunc func(GetRbacAuthorizationV1APIResourcesParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetRbacAuthorizationV1APIResourcesHandlerFunc) Handle(params GetRbacAuthorizationV1APIResourcesParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetRbacAuthorizationV1APIResourcesHandler interface for that can handle valid get rbac authorization v1 API resources params
type GetRbacAuthorizationV1APIResourcesHandler interface {
	Handle(GetRbacAuthorizationV1APIResourcesParams, interface{}) middleware.Responder
}

// NewGetRbacAuthorizationV1APIResources creates a new http.Handler for the get rbac authorization v1 API resources operation
func NewGetRbacAuthorizationV1APIResources(ctx *middleware.Context, handler GetRbacAuthorizationV1APIResourcesHandler) *GetRbacAuthorizationV1APIResources {
	return &GetRbacAuthorizationV1APIResources{Context: ctx, Handler: handler}
}

/*GetRbacAuthorizationV1APIResources swagger:route GET /apis/rbac.authorization.k8s.io/v1/ rbacAuthorization_v1 getRbacAuthorizationV1ApiResources

get available resources

*/
type GetRbacAuthorizationV1APIResources struct {
	Context *middleware.Context
	Handler GetRbacAuthorizationV1APIResourcesHandler
}

func (o *GetRbacAuthorizationV1APIResources) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetRbacAuthorizationV1APIResourcesParams()

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
