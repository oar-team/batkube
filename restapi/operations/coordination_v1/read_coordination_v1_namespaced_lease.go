// Code generated by go-swagger; DO NOT EDIT.

package coordination_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ReadCoordinationV1NamespacedLeaseHandlerFunc turns a function with the right signature into a read coordination v1 namespaced lease handler
type ReadCoordinationV1NamespacedLeaseHandlerFunc func(ReadCoordinationV1NamespacedLeaseParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ReadCoordinationV1NamespacedLeaseHandlerFunc) Handle(params ReadCoordinationV1NamespacedLeaseParams) middleware.Responder {
	return fn(params)
}

// ReadCoordinationV1NamespacedLeaseHandler interface for that can handle valid read coordination v1 namespaced lease params
type ReadCoordinationV1NamespacedLeaseHandler interface {
	Handle(ReadCoordinationV1NamespacedLeaseParams) middleware.Responder
}

// NewReadCoordinationV1NamespacedLease creates a new http.Handler for the read coordination v1 namespaced lease operation
func NewReadCoordinationV1NamespacedLease(ctx *middleware.Context, handler ReadCoordinationV1NamespacedLeaseHandler) *ReadCoordinationV1NamespacedLease {
	return &ReadCoordinationV1NamespacedLease{Context: ctx, Handler: handler}
}

/*ReadCoordinationV1NamespacedLease swagger:route GET /apis/coordination.k8s.io/v1/namespaces/{namespace}/leases/{name} coordination_v1 readCoordinationV1NamespacedLease

read the specified Lease

*/
type ReadCoordinationV1NamespacedLease struct {
	Context *middleware.Context
	Handler ReadCoordinationV1NamespacedLeaseHandler
}

func (o *ReadCoordinationV1NamespacedLease) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewReadCoordinationV1NamespacedLeaseParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
