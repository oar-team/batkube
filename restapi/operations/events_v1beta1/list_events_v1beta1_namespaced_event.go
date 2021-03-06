// Code generated by go-swagger; DO NOT EDIT.

package events_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ListEventsV1beta1NamespacedEventHandlerFunc turns a function with the right signature into a list events v1beta1 namespaced event handler
type ListEventsV1beta1NamespacedEventHandlerFunc func(ListEventsV1beta1NamespacedEventParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ListEventsV1beta1NamespacedEventHandlerFunc) Handle(params ListEventsV1beta1NamespacedEventParams) middleware.Responder {
	return fn(params)
}

// ListEventsV1beta1NamespacedEventHandler interface for that can handle valid list events v1beta1 namespaced event params
type ListEventsV1beta1NamespacedEventHandler interface {
	Handle(ListEventsV1beta1NamespacedEventParams) middleware.Responder
}

// NewListEventsV1beta1NamespacedEvent creates a new http.Handler for the list events v1beta1 namespaced event operation
func NewListEventsV1beta1NamespacedEvent(ctx *middleware.Context, handler ListEventsV1beta1NamespacedEventHandler) *ListEventsV1beta1NamespacedEvent {
	return &ListEventsV1beta1NamespacedEvent{Context: ctx, Handler: handler}
}

/*ListEventsV1beta1NamespacedEvent swagger:route GET /apis/events.k8s.io/v1beta1/namespaces/{namespace}/events events_v1beta1 listEventsV1beta1NamespacedEvent

list or watch objects of kind Event

*/
type ListEventsV1beta1NamespacedEvent struct {
	Context *middleware.Context
	Handler ListEventsV1beta1NamespacedEventHandler
}

func (o *ListEventsV1beta1NamespacedEvent) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewListEventsV1beta1NamespacedEventParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
