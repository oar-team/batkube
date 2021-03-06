// Code generated by go-swagger; DO NOT EDIT.

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ListCoreV1EventForAllNamespacesHandlerFunc turns a function with the right signature into a list core v1 event for all namespaces handler
type ListCoreV1EventForAllNamespacesHandlerFunc func(ListCoreV1EventForAllNamespacesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ListCoreV1EventForAllNamespacesHandlerFunc) Handle(params ListCoreV1EventForAllNamespacesParams) middleware.Responder {
	return fn(params)
}

// ListCoreV1EventForAllNamespacesHandler interface for that can handle valid list core v1 event for all namespaces params
type ListCoreV1EventForAllNamespacesHandler interface {
	Handle(ListCoreV1EventForAllNamespacesParams) middleware.Responder
}

// NewListCoreV1EventForAllNamespaces creates a new http.Handler for the list core v1 event for all namespaces operation
func NewListCoreV1EventForAllNamespaces(ctx *middleware.Context, handler ListCoreV1EventForAllNamespacesHandler) *ListCoreV1EventForAllNamespaces {
	return &ListCoreV1EventForAllNamespaces{Context: ctx, Handler: handler}
}

/*ListCoreV1EventForAllNamespaces swagger:route GET /api/v1/events core_v1 listCoreV1EventForAllNamespaces

list or watch objects of kind Event

*/
type ListCoreV1EventForAllNamespaces struct {
	Context *middleware.Context
	Handler ListCoreV1EventForAllNamespacesHandler
}

func (o *ListCoreV1EventForAllNamespaces) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewListCoreV1EventForAllNamespacesParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
