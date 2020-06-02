// Code generated by go-swagger; DO NOT EDIT.

package events_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// WatchEventsV1beta1EventListForAllNamespacesHandlerFunc turns a function with the right signature into a watch events v1beta1 event list for all namespaces handler
type WatchEventsV1beta1EventListForAllNamespacesHandlerFunc func(WatchEventsV1beta1EventListForAllNamespacesParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn WatchEventsV1beta1EventListForAllNamespacesHandlerFunc) Handle(params WatchEventsV1beta1EventListForAllNamespacesParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// WatchEventsV1beta1EventListForAllNamespacesHandler interface for that can handle valid watch events v1beta1 event list for all namespaces params
type WatchEventsV1beta1EventListForAllNamespacesHandler interface {
	Handle(WatchEventsV1beta1EventListForAllNamespacesParams, interface{}) middleware.Responder
}

// NewWatchEventsV1beta1EventListForAllNamespaces creates a new http.Handler for the watch events v1beta1 event list for all namespaces operation
func NewWatchEventsV1beta1EventListForAllNamespaces(ctx *middleware.Context, handler WatchEventsV1beta1EventListForAllNamespacesHandler) *WatchEventsV1beta1EventListForAllNamespaces {
	return &WatchEventsV1beta1EventListForAllNamespaces{Context: ctx, Handler: handler}
}

/*WatchEventsV1beta1EventListForAllNamespaces swagger:route GET /apis/events.k8s.io/v1beta1/watch/events events_v1beta1 watchEventsV1beta1EventListForAllNamespaces

watch individual changes to a list of Event. deprecated: use the 'watch' parameter with a list operation instead.

*/
type WatchEventsV1beta1EventListForAllNamespaces struct {
	Context *middleware.Context
	Handler WatchEventsV1beta1EventListForAllNamespacesHandler
}

func (o *WatchEventsV1beta1EventListForAllNamespaces) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewWatchEventsV1beta1EventListForAllNamespacesParams()

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
