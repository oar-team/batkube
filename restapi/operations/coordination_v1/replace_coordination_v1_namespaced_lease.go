// Code generated by go-swagger; DO NOT EDIT.

package coordination_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ReplaceCoordinationV1NamespacedLeaseHandlerFunc turns a function with the right signature into a replace coordination v1 namespaced lease handler
type ReplaceCoordinationV1NamespacedLeaseHandlerFunc func(ReplaceCoordinationV1NamespacedLeaseParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn ReplaceCoordinationV1NamespacedLeaseHandlerFunc) Handle(params ReplaceCoordinationV1NamespacedLeaseParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// ReplaceCoordinationV1NamespacedLeaseHandler interface for that can handle valid replace coordination v1 namespaced lease params
type ReplaceCoordinationV1NamespacedLeaseHandler interface {
	Handle(ReplaceCoordinationV1NamespacedLeaseParams, interface{}) middleware.Responder
}

// NewReplaceCoordinationV1NamespacedLease creates a new http.Handler for the replace coordination v1 namespaced lease operation
func NewReplaceCoordinationV1NamespacedLease(ctx *middleware.Context, handler ReplaceCoordinationV1NamespacedLeaseHandler) *ReplaceCoordinationV1NamespacedLease {
	return &ReplaceCoordinationV1NamespacedLease{Context: ctx, Handler: handler}
}

/*ReplaceCoordinationV1NamespacedLease swagger:route PUT /apis/coordination.k8s.io/v1/namespaces/{namespace}/leases/{name} coordination_v1 replaceCoordinationV1NamespacedLease

replace the specified Lease

*/
type ReplaceCoordinationV1NamespacedLease struct {
	Context *middleware.Context
	Handler ReplaceCoordinationV1NamespacedLeaseHandler
}

func (o *ReplaceCoordinationV1NamespacedLease) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewReplaceCoordinationV1NamespacedLeaseParams()

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
