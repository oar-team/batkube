// Code generated by go-swagger; DO NOT EDIT.

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// WatchCoreV1NamespacedPodListHandlerFunc turns a function with the right signature into a watch core v1 namespaced pod list handler
type WatchCoreV1NamespacedPodListHandlerFunc func(WatchCoreV1NamespacedPodListParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn WatchCoreV1NamespacedPodListHandlerFunc) Handle(params WatchCoreV1NamespacedPodListParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// WatchCoreV1NamespacedPodListHandler interface for that can handle valid watch core v1 namespaced pod list params
type WatchCoreV1NamespacedPodListHandler interface {
	Handle(WatchCoreV1NamespacedPodListParams, interface{}) middleware.Responder
}

// NewWatchCoreV1NamespacedPodList creates a new http.Handler for the watch core v1 namespaced pod list operation
func NewWatchCoreV1NamespacedPodList(ctx *middleware.Context, handler WatchCoreV1NamespacedPodListHandler) *WatchCoreV1NamespacedPodList {
	return &WatchCoreV1NamespacedPodList{Context: ctx, Handler: handler}
}

/*WatchCoreV1NamespacedPodList swagger:route GET /api/v1/watch/namespaces/{namespace}/pods core_v1 watchCoreV1NamespacedPodList

watch individual changes to a list of Pod. deprecated: use the 'watch' parameter with a list operation instead.

*/
type WatchCoreV1NamespacedPodList struct {
	Context *middleware.Context
	Handler WatchCoreV1NamespacedPodListHandler
}

func (o *WatchCoreV1NamespacedPodList) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewWatchCoreV1NamespacedPodListParams()

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
