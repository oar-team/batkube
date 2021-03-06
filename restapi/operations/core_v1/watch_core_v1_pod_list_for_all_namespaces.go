// Code generated by go-swagger; DO NOT EDIT.

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// WatchCoreV1PodListForAllNamespacesHandlerFunc turns a function with the right signature into a watch core v1 pod list for all namespaces handler
type WatchCoreV1PodListForAllNamespacesHandlerFunc func(WatchCoreV1PodListForAllNamespacesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn WatchCoreV1PodListForAllNamespacesHandlerFunc) Handle(params WatchCoreV1PodListForAllNamespacesParams) middleware.Responder {
	return fn(params)
}

// WatchCoreV1PodListForAllNamespacesHandler interface for that can handle valid watch core v1 pod list for all namespaces params
type WatchCoreV1PodListForAllNamespacesHandler interface {
	Handle(WatchCoreV1PodListForAllNamespacesParams) middleware.Responder
}

// NewWatchCoreV1PodListForAllNamespaces creates a new http.Handler for the watch core v1 pod list for all namespaces operation
func NewWatchCoreV1PodListForAllNamespaces(ctx *middleware.Context, handler WatchCoreV1PodListForAllNamespacesHandler) *WatchCoreV1PodListForAllNamespaces {
	return &WatchCoreV1PodListForAllNamespaces{Context: ctx, Handler: handler}
}

/*WatchCoreV1PodListForAllNamespaces swagger:route GET /api/v1/watch/pods core_v1 watchCoreV1PodListForAllNamespaces

watch individual changes to a list of Pod. deprecated: use the 'watch' parameter with a list operation instead.

*/
type WatchCoreV1PodListForAllNamespaces struct {
	Context *middleware.Context
	Handler WatchCoreV1PodListForAllNamespacesHandler
}

func (o *WatchCoreV1PodListForAllNamespaces) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewWatchCoreV1PodListForAllNamespacesParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
