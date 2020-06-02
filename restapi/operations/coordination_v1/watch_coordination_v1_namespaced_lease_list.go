// Code generated by go-swagger; DO NOT EDIT.

package coordination_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// WatchCoordinationV1NamespacedLeaseListHandlerFunc turns a function with the right signature into a watch coordination v1 namespaced lease list handler
type WatchCoordinationV1NamespacedLeaseListHandlerFunc func(WatchCoordinationV1NamespacedLeaseListParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn WatchCoordinationV1NamespacedLeaseListHandlerFunc) Handle(params WatchCoordinationV1NamespacedLeaseListParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// WatchCoordinationV1NamespacedLeaseListHandler interface for that can handle valid watch coordination v1 namespaced lease list params
type WatchCoordinationV1NamespacedLeaseListHandler interface {
	Handle(WatchCoordinationV1NamespacedLeaseListParams, interface{}) middleware.Responder
}

// NewWatchCoordinationV1NamespacedLeaseList creates a new http.Handler for the watch coordination v1 namespaced lease list operation
func NewWatchCoordinationV1NamespacedLeaseList(ctx *middleware.Context, handler WatchCoordinationV1NamespacedLeaseListHandler) *WatchCoordinationV1NamespacedLeaseList {
	return &WatchCoordinationV1NamespacedLeaseList{Context: ctx, Handler: handler}
}

/*WatchCoordinationV1NamespacedLeaseList swagger:route GET /apis/coordination.k8s.io/v1/watch/namespaces/{namespace}/leases coordination_v1 watchCoordinationV1NamespacedLeaseList

watch individual changes to a list of Lease. deprecated: use the 'watch' parameter with a list operation instead.

*/
type WatchCoordinationV1NamespacedLeaseList struct {
	Context *middleware.Context
	Handler WatchCoordinationV1NamespacedLeaseListHandler
}

func (o *WatchCoordinationV1NamespacedLeaseList) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewWatchCoordinationV1NamespacedLeaseListParams()

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
