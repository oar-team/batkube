// Code generated by go-swagger; DO NOT EDIT.

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// WatchCoreV1NamespacedResourceQuotaHandlerFunc turns a function with the right signature into a watch core v1 namespaced resource quota handler
type WatchCoreV1NamespacedResourceQuotaHandlerFunc func(WatchCoreV1NamespacedResourceQuotaParams) middleware.Responder

// Handle executing the request and returning a response
func (fn WatchCoreV1NamespacedResourceQuotaHandlerFunc) Handle(params WatchCoreV1NamespacedResourceQuotaParams) middleware.Responder {
	return fn(params)
}

// WatchCoreV1NamespacedResourceQuotaHandler interface for that can handle valid watch core v1 namespaced resource quota params
type WatchCoreV1NamespacedResourceQuotaHandler interface {
	Handle(WatchCoreV1NamespacedResourceQuotaParams) middleware.Responder
}

// NewWatchCoreV1NamespacedResourceQuota creates a new http.Handler for the watch core v1 namespaced resource quota operation
func NewWatchCoreV1NamespacedResourceQuota(ctx *middleware.Context, handler WatchCoreV1NamespacedResourceQuotaHandler) *WatchCoreV1NamespacedResourceQuota {
	return &WatchCoreV1NamespacedResourceQuota{Context: ctx, Handler: handler}
}

/*WatchCoreV1NamespacedResourceQuota swagger:route GET /api/v1/watch/namespaces/{namespace}/resourcequotas/{name} core_v1 watchCoreV1NamespacedResourceQuota

watch changes to an object of kind ResourceQuota. deprecated: use the 'watch' parameter with a list operation instead, filtered to a single item with the 'fieldSelector' parameter.

*/
type WatchCoreV1NamespacedResourceQuota struct {
	Context *middleware.Context
	Handler WatchCoreV1NamespacedResourceQuotaHandler
}

func (o *WatchCoreV1NamespacedResourceQuota) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewWatchCoreV1NamespacedResourceQuotaParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}