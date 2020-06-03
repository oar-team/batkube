// Code generated by go-swagger; DO NOT EDIT.

package events_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteEventsV1beta1NamespacedEventHandlerFunc turns a function with the right signature into a delete events v1beta1 namespaced event handler
type DeleteEventsV1beta1NamespacedEventHandlerFunc func(DeleteEventsV1beta1NamespacedEventParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteEventsV1beta1NamespacedEventHandlerFunc) Handle(params DeleteEventsV1beta1NamespacedEventParams) middleware.Responder {
	return fn(params)
}

// DeleteEventsV1beta1NamespacedEventHandler interface for that can handle valid delete events v1beta1 namespaced event params
type DeleteEventsV1beta1NamespacedEventHandler interface {
	Handle(DeleteEventsV1beta1NamespacedEventParams) middleware.Responder
}

// NewDeleteEventsV1beta1NamespacedEvent creates a new http.Handler for the delete events v1beta1 namespaced event operation
func NewDeleteEventsV1beta1NamespacedEvent(ctx *middleware.Context, handler DeleteEventsV1beta1NamespacedEventHandler) *DeleteEventsV1beta1NamespacedEvent {
	return &DeleteEventsV1beta1NamespacedEvent{Context: ctx, Handler: handler}
}

/*DeleteEventsV1beta1NamespacedEvent swagger:route DELETE /apis/events.k8s.io/v1beta1/namespaces/{namespace}/events/{name} events_v1beta1 deleteEventsV1beta1NamespacedEvent

delete an Event

*/
type DeleteEventsV1beta1NamespacedEvent struct {
	Context *middleware.Context
	Handler DeleteEventsV1beta1NamespacedEventHandler
}

func (o *DeleteEventsV1beta1NamespacedEvent) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteEventsV1beta1NamespacedEventParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
