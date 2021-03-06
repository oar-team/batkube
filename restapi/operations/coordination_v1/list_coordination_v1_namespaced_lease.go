// Code generated by go-swagger; DO NOT EDIT.

package coordination_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ListCoordinationV1NamespacedLeaseHandlerFunc turns a function with the right signature into a list coordination v1 namespaced lease handler
type ListCoordinationV1NamespacedLeaseHandlerFunc func(ListCoordinationV1NamespacedLeaseParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ListCoordinationV1NamespacedLeaseHandlerFunc) Handle(params ListCoordinationV1NamespacedLeaseParams) middleware.Responder {
	return fn(params)
}

// ListCoordinationV1NamespacedLeaseHandler interface for that can handle valid list coordination v1 namespaced lease params
type ListCoordinationV1NamespacedLeaseHandler interface {
	Handle(ListCoordinationV1NamespacedLeaseParams) middleware.Responder
}

// NewListCoordinationV1NamespacedLease creates a new http.Handler for the list coordination v1 namespaced lease operation
func NewListCoordinationV1NamespacedLease(ctx *middleware.Context, handler ListCoordinationV1NamespacedLeaseHandler) *ListCoordinationV1NamespacedLease {
	return &ListCoordinationV1NamespacedLease{Context: ctx, Handler: handler}
}

/*ListCoordinationV1NamespacedLease swagger:route GET /apis/coordination.k8s.io/v1/namespaces/{namespace}/leases coordination_v1 listCoordinationV1NamespacedLease

list or watch objects of kind Lease

*/
type ListCoordinationV1NamespacedLease struct {
	Context *middleware.Context
	Handler ListCoordinationV1NamespacedLeaseHandler
}

func (o *ListCoordinationV1NamespacedLease) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewListCoordinationV1NamespacedLeaseParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
