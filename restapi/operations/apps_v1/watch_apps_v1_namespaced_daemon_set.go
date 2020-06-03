// Code generated by go-swagger; DO NOT EDIT.

package apps_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// WatchAppsV1NamespacedDaemonSetHandlerFunc turns a function with the right signature into a watch apps v1 namespaced daemon set handler
type WatchAppsV1NamespacedDaemonSetHandlerFunc func(WatchAppsV1NamespacedDaemonSetParams) middleware.Responder

// Handle executing the request and returning a response
func (fn WatchAppsV1NamespacedDaemonSetHandlerFunc) Handle(params WatchAppsV1NamespacedDaemonSetParams) middleware.Responder {
	return fn(params)
}

// WatchAppsV1NamespacedDaemonSetHandler interface for that can handle valid watch apps v1 namespaced daemon set params
type WatchAppsV1NamespacedDaemonSetHandler interface {
	Handle(WatchAppsV1NamespacedDaemonSetParams) middleware.Responder
}

// NewWatchAppsV1NamespacedDaemonSet creates a new http.Handler for the watch apps v1 namespaced daemon set operation
func NewWatchAppsV1NamespacedDaemonSet(ctx *middleware.Context, handler WatchAppsV1NamespacedDaemonSetHandler) *WatchAppsV1NamespacedDaemonSet {
	return &WatchAppsV1NamespacedDaemonSet{Context: ctx, Handler: handler}
}

/*WatchAppsV1NamespacedDaemonSet swagger:route GET /apis/apps/v1/watch/namespaces/{namespace}/daemonsets/{name} apps_v1 watchAppsV1NamespacedDaemonSet

watch changes to an object of kind DaemonSet. deprecated: use the 'watch' parameter with a list operation instead, filtered to a single item with the 'fieldSelector' parameter.

*/
type WatchAppsV1NamespacedDaemonSet struct {
	Context *middleware.Context
	Handler WatchAppsV1NamespacedDaemonSetHandler
}

func (o *WatchAppsV1NamespacedDaemonSet) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewWatchAppsV1NamespacedDaemonSetParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
