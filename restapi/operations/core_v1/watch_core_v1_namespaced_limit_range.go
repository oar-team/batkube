// Code generated by go-swagger; DO NOT EDIT.

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// WatchCoreV1NamespacedLimitRangeHandlerFunc turns a function with the right signature into a watch core v1 namespaced limit range handler
type WatchCoreV1NamespacedLimitRangeHandlerFunc func(WatchCoreV1NamespacedLimitRangeParams) middleware.Responder

// Handle executing the request and returning a response
func (fn WatchCoreV1NamespacedLimitRangeHandlerFunc) Handle(params WatchCoreV1NamespacedLimitRangeParams) middleware.Responder {
	return fn(params)
}

// WatchCoreV1NamespacedLimitRangeHandler interface for that can handle valid watch core v1 namespaced limit range params
type WatchCoreV1NamespacedLimitRangeHandler interface {
	Handle(WatchCoreV1NamespacedLimitRangeParams) middleware.Responder
}

// NewWatchCoreV1NamespacedLimitRange creates a new http.Handler for the watch core v1 namespaced limit range operation
func NewWatchCoreV1NamespacedLimitRange(ctx *middleware.Context, handler WatchCoreV1NamespacedLimitRangeHandler) *WatchCoreV1NamespacedLimitRange {
	return &WatchCoreV1NamespacedLimitRange{Context: ctx, Handler: handler}
}

/*WatchCoreV1NamespacedLimitRange swagger:route GET /api/v1/watch/namespaces/{namespace}/limitranges/{name} core_v1 watchCoreV1NamespacedLimitRange

watch changes to an object of kind LimitRange. deprecated: use the 'watch' parameter with a list operation instead, filtered to a single item with the 'fieldSelector' parameter.

*/
type WatchCoreV1NamespacedLimitRange struct {
	Context *middleware.Context
	Handler WatchCoreV1NamespacedLimitRangeHandler
}

func (o *WatchCoreV1NamespacedLimitRange) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewWatchCoreV1NamespacedLimitRangeParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
