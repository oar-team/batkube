// Code generated by go-swagger; DO NOT EDIT.

package events_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// WatchEventsV1beta1NamespacedEventListHandlerFunc turns a function with the right signature into a watch events v1beta1 namespaced event list handler
type WatchEventsV1beta1NamespacedEventListHandlerFunc func(WatchEventsV1beta1NamespacedEventListParams) middleware.Responder

// Handle executing the request and returning a response
func (fn WatchEventsV1beta1NamespacedEventListHandlerFunc) Handle(params WatchEventsV1beta1NamespacedEventListParams) middleware.Responder {
	return fn(params)
}

// WatchEventsV1beta1NamespacedEventListHandler interface for that can handle valid watch events v1beta1 namespaced event list params
type WatchEventsV1beta1NamespacedEventListHandler interface {
	Handle(WatchEventsV1beta1NamespacedEventListParams) middleware.Responder
}

// NewWatchEventsV1beta1NamespacedEventList creates a new http.Handler for the watch events v1beta1 namespaced event list operation
func NewWatchEventsV1beta1NamespacedEventList(ctx *middleware.Context, handler WatchEventsV1beta1NamespacedEventListHandler) *WatchEventsV1beta1NamespacedEventList {
	return &WatchEventsV1beta1NamespacedEventList{Context: ctx, Handler: handler}
}

/*WatchEventsV1beta1NamespacedEventList swagger:route GET /apis/events.k8s.io/v1beta1/watch/namespaces/{namespace}/events events_v1beta1 watchEventsV1beta1NamespacedEventList

watch individual changes to a list of Event. deprecated: use the 'watch' parameter with a list operation instead.

*/
type WatchEventsV1beta1NamespacedEventList struct {
	Context *middleware.Context
	Handler WatchEventsV1beta1NamespacedEventListHandler
}

func (o *WatchEventsV1beta1NamespacedEventList) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewWatchEventsV1beta1NamespacedEventListParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
