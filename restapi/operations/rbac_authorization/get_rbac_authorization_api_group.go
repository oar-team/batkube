// Code generated by go-swagger; DO NOT EDIT.

package rbac_authorization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetRbacAuthorizationAPIGroupHandlerFunc turns a function with the right signature into a get rbac authorization API group handler
type GetRbacAuthorizationAPIGroupHandlerFunc func(GetRbacAuthorizationAPIGroupParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetRbacAuthorizationAPIGroupHandlerFunc) Handle(params GetRbacAuthorizationAPIGroupParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetRbacAuthorizationAPIGroupHandler interface for that can handle valid get rbac authorization API group params
type GetRbacAuthorizationAPIGroupHandler interface {
	Handle(GetRbacAuthorizationAPIGroupParams, interface{}) middleware.Responder
}

// NewGetRbacAuthorizationAPIGroup creates a new http.Handler for the get rbac authorization API group operation
func NewGetRbacAuthorizationAPIGroup(ctx *middleware.Context, handler GetRbacAuthorizationAPIGroupHandler) *GetRbacAuthorizationAPIGroup {
	return &GetRbacAuthorizationAPIGroup{Context: ctx, Handler: handler}
}

/*GetRbacAuthorizationAPIGroup swagger:route GET /apis/rbac.authorization.k8s.io/ rbacAuthorization getRbacAuthorizationApiGroup

get information of a group

*/
type GetRbacAuthorizationAPIGroup struct {
	Context *middleware.Context
	Handler GetRbacAuthorizationAPIGroupHandler
}

func (o *GetRbacAuthorizationAPIGroup) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetRbacAuthorizationAPIGroupParams()

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
