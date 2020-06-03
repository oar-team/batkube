// Code generated by go-swagger; DO NOT EDIT.

package coordination_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// CreateCoordinationV1NamespacedLeaseHandlerFunc turns a function with the right signature into a create coordination v1 namespaced lease handler
type CreateCoordinationV1NamespacedLeaseHandlerFunc func(CreateCoordinationV1NamespacedLeaseParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateCoordinationV1NamespacedLeaseHandlerFunc) Handle(params CreateCoordinationV1NamespacedLeaseParams) middleware.Responder {
	return fn(params)
}

// CreateCoordinationV1NamespacedLeaseHandler interface for that can handle valid create coordination v1 namespaced lease params
type CreateCoordinationV1NamespacedLeaseHandler interface {
	Handle(CreateCoordinationV1NamespacedLeaseParams) middleware.Responder
}

// NewCreateCoordinationV1NamespacedLease creates a new http.Handler for the create coordination v1 namespaced lease operation
func NewCreateCoordinationV1NamespacedLease(ctx *middleware.Context, handler CreateCoordinationV1NamespacedLeaseHandler) *CreateCoordinationV1NamespacedLease {
	return &CreateCoordinationV1NamespacedLease{Context: ctx, Handler: handler}
}

/*CreateCoordinationV1NamespacedLease swagger:route POST /apis/coordination.k8s.io/v1/namespaces/{namespace}/leases coordination_v1 createCoordinationV1NamespacedLease

create a Lease

*/
type CreateCoordinationV1NamespacedLease struct {
	Context *middleware.Context
	Handler CreateCoordinationV1NamespacedLeaseHandler
}

func (o *CreateCoordinationV1NamespacedLease) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCreateCoordinationV1NamespacedLeaseParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
