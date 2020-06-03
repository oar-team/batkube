// Code generated by go-swagger; DO NOT EDIT.

package apps_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// WatchAppsV1StatefulSetListForAllNamespacesHandlerFunc turns a function with the right signature into a watch apps v1 stateful set list for all namespaces handler
type WatchAppsV1StatefulSetListForAllNamespacesHandlerFunc func(WatchAppsV1StatefulSetListForAllNamespacesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn WatchAppsV1StatefulSetListForAllNamespacesHandlerFunc) Handle(params WatchAppsV1StatefulSetListForAllNamespacesParams) middleware.Responder {
	return fn(params)
}

// WatchAppsV1StatefulSetListForAllNamespacesHandler interface for that can handle valid watch apps v1 stateful set list for all namespaces params
type WatchAppsV1StatefulSetListForAllNamespacesHandler interface {
	Handle(WatchAppsV1StatefulSetListForAllNamespacesParams) middleware.Responder
}

// NewWatchAppsV1StatefulSetListForAllNamespaces creates a new http.Handler for the watch apps v1 stateful set list for all namespaces operation
func NewWatchAppsV1StatefulSetListForAllNamespaces(ctx *middleware.Context, handler WatchAppsV1StatefulSetListForAllNamespacesHandler) *WatchAppsV1StatefulSetListForAllNamespaces {
	return &WatchAppsV1StatefulSetListForAllNamespaces{Context: ctx, Handler: handler}
}

/*WatchAppsV1StatefulSetListForAllNamespaces swagger:route GET /apis/apps/v1/watch/statefulsets apps_v1 watchAppsV1StatefulSetListForAllNamespaces

watch individual changes to a list of StatefulSet. deprecated: use the 'watch' parameter with a list operation instead.

*/
type WatchAppsV1StatefulSetListForAllNamespaces struct {
	Context *middleware.Context
	Handler WatchAppsV1StatefulSetListForAllNamespacesHandler
}

func (o *WatchAppsV1StatefulSetListForAllNamespaces) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewWatchAppsV1StatefulSetListForAllNamespacesParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}