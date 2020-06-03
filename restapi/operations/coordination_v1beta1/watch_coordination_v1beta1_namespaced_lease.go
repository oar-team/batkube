// Code generated by go-swagger; DO NOT EDIT.

package coordination_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// WatchCoordinationV1beta1NamespacedLeaseHandlerFunc turns a function with the right signature into a watch coordination v1beta1 namespaced lease handler
type WatchCoordinationV1beta1NamespacedLeaseHandlerFunc func(WatchCoordinationV1beta1NamespacedLeaseParams) middleware.Responder

// Handle executing the request and returning a response
func (fn WatchCoordinationV1beta1NamespacedLeaseHandlerFunc) Handle(params WatchCoordinationV1beta1NamespacedLeaseParams) middleware.Responder {
	return fn(params)
}

// WatchCoordinationV1beta1NamespacedLeaseHandler interface for that can handle valid watch coordination v1beta1 namespaced lease params
type WatchCoordinationV1beta1NamespacedLeaseHandler interface {
	Handle(WatchCoordinationV1beta1NamespacedLeaseParams) middleware.Responder
}

// NewWatchCoordinationV1beta1NamespacedLease creates a new http.Handler for the watch coordination v1beta1 namespaced lease operation
func NewWatchCoordinationV1beta1NamespacedLease(ctx *middleware.Context, handler WatchCoordinationV1beta1NamespacedLeaseHandler) *WatchCoordinationV1beta1NamespacedLease {
	return &WatchCoordinationV1beta1NamespacedLease{Context: ctx, Handler: handler}
}

/*WatchCoordinationV1beta1NamespacedLease swagger:route GET /apis/coordination.k8s.io/v1beta1/watch/namespaces/{namespace}/leases/{name} coordination_v1beta1 watchCoordinationV1beta1NamespacedLease

watch changes to an object of kind Lease. deprecated: use the 'watch' parameter with a list operation instead, filtered to a single item with the 'fieldSelector' parameter.

*/
type WatchCoordinationV1beta1NamespacedLease struct {
	Context *middleware.Context
	Handler WatchCoordinationV1beta1NamespacedLeaseHandler
}

func (o *WatchCoordinationV1beta1NamespacedLease) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewWatchCoordinationV1beta1NamespacedLeaseParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
