// Code generated by go-swagger; DO NOT EDIT.

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ListCoreV1SecretForAllNamespacesHandlerFunc turns a function with the right signature into a list core v1 secret for all namespaces handler
type ListCoreV1SecretForAllNamespacesHandlerFunc func(ListCoreV1SecretForAllNamespacesParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn ListCoreV1SecretForAllNamespacesHandlerFunc) Handle(params ListCoreV1SecretForAllNamespacesParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// ListCoreV1SecretForAllNamespacesHandler interface for that can handle valid list core v1 secret for all namespaces params
type ListCoreV1SecretForAllNamespacesHandler interface {
	Handle(ListCoreV1SecretForAllNamespacesParams, interface{}) middleware.Responder
}

// NewListCoreV1SecretForAllNamespaces creates a new http.Handler for the list core v1 secret for all namespaces operation
func NewListCoreV1SecretForAllNamespaces(ctx *middleware.Context, handler ListCoreV1SecretForAllNamespacesHandler) *ListCoreV1SecretForAllNamespaces {
	return &ListCoreV1SecretForAllNamespaces{Context: ctx, Handler: handler}
}

/*ListCoreV1SecretForAllNamespaces swagger:route GET /api/v1/secrets core_v1 listCoreV1SecretForAllNamespaces

list or watch objects of kind Secret

*/
type ListCoreV1SecretForAllNamespaces struct {
	Context *middleware.Context
	Handler ListCoreV1SecretForAllNamespacesHandler
}

func (o *ListCoreV1SecretForAllNamespaces) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewListCoreV1SecretForAllNamespacesParams()

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
