// Code generated by go-swagger; DO NOT EDIT.

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ReadCoreV1NamespacedEventHandlerFunc turns a function with the right signature into a read core v1 namespaced event handler
type ReadCoreV1NamespacedEventHandlerFunc func(ReadCoreV1NamespacedEventParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn ReadCoreV1NamespacedEventHandlerFunc) Handle(params ReadCoreV1NamespacedEventParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// ReadCoreV1NamespacedEventHandler interface for that can handle valid read core v1 namespaced event params
type ReadCoreV1NamespacedEventHandler interface {
	Handle(ReadCoreV1NamespacedEventParams, interface{}) middleware.Responder
}

// NewReadCoreV1NamespacedEvent creates a new http.Handler for the read core v1 namespaced event operation
func NewReadCoreV1NamespacedEvent(ctx *middleware.Context, handler ReadCoreV1NamespacedEventHandler) *ReadCoreV1NamespacedEvent {
	return &ReadCoreV1NamespacedEvent{Context: ctx, Handler: handler}
}

/*ReadCoreV1NamespacedEvent swagger:route GET /api/v1/namespaces/{namespace}/events/{name} core_v1 readCoreV1NamespacedEvent

read the specified Event

*/
type ReadCoreV1NamespacedEvent struct {
	Context *middleware.Context
	Handler ReadCoreV1NamespacedEventHandler
}

func (o *ReadCoreV1NamespacedEvent) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewReadCoreV1NamespacedEventParams()

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
