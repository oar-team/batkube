// Code generated by go-swagger; DO NOT EDIT.

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ReplaceCoreV1NodeStatusHandlerFunc turns a function with the right signature into a replace core v1 node status handler
type ReplaceCoreV1NodeStatusHandlerFunc func(ReplaceCoreV1NodeStatusParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn ReplaceCoreV1NodeStatusHandlerFunc) Handle(params ReplaceCoreV1NodeStatusParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// ReplaceCoreV1NodeStatusHandler interface for that can handle valid replace core v1 node status params
type ReplaceCoreV1NodeStatusHandler interface {
	Handle(ReplaceCoreV1NodeStatusParams, interface{}) middleware.Responder
}

// NewReplaceCoreV1NodeStatus creates a new http.Handler for the replace core v1 node status operation
func NewReplaceCoreV1NodeStatus(ctx *middleware.Context, handler ReplaceCoreV1NodeStatusHandler) *ReplaceCoreV1NodeStatus {
	return &ReplaceCoreV1NodeStatus{Context: ctx, Handler: handler}
}

/*ReplaceCoreV1NodeStatus swagger:route PUT /api/v1/nodes/{name}/status core_v1 replaceCoreV1NodeStatus

replace status of the specified Node

*/
type ReplaceCoreV1NodeStatus struct {
	Context *middleware.Context
	Handler ReplaceCoreV1NodeStatusHandler
}

func (o *ReplaceCoreV1NodeStatus) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewReplaceCoreV1NodeStatusParams()

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
