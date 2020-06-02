// Code generated by go-swagger; DO NOT EDIT.

package coordination_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// WatchCoordinationV1NamespacedLeaseHandlerFunc turns a function with the right signature into a watch coordination v1 namespaced lease handler
type WatchCoordinationV1NamespacedLeaseHandlerFunc func(WatchCoordinationV1NamespacedLeaseParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn WatchCoordinationV1NamespacedLeaseHandlerFunc) Handle(params WatchCoordinationV1NamespacedLeaseParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// WatchCoordinationV1NamespacedLeaseHandler interface for that can handle valid watch coordination v1 namespaced lease params
type WatchCoordinationV1NamespacedLeaseHandler interface {
	Handle(WatchCoordinationV1NamespacedLeaseParams, interface{}) middleware.Responder
}

// NewWatchCoordinationV1NamespacedLease creates a new http.Handler for the watch coordination v1 namespaced lease operation
func NewWatchCoordinationV1NamespacedLease(ctx *middleware.Context, handler WatchCoordinationV1NamespacedLeaseHandler) *WatchCoordinationV1NamespacedLease {
	return &WatchCoordinationV1NamespacedLease{Context: ctx, Handler: handler}
}

/*WatchCoordinationV1NamespacedLease swagger:route GET /apis/coordination.k8s.io/v1/watch/namespaces/{namespace}/leases/{name} coordination_v1 watchCoordinationV1NamespacedLease

watch changes to an object of kind Lease. deprecated: use the 'watch' parameter with a list operation instead, filtered to a single item with the 'fieldSelector' parameter.

*/
type WatchCoordinationV1NamespacedLease struct {
	Context *middleware.Context
	Handler WatchCoordinationV1NamespacedLeaseHandler
}

func (o *WatchCoordinationV1NamespacedLease) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewWatchCoordinationV1NamespacedLeaseParams()

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
