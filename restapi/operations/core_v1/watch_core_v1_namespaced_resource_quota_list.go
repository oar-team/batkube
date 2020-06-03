// Code generated by go-swagger; DO NOT EDIT.

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// WatchCoreV1NamespacedResourceQuotaListHandlerFunc turns a function with the right signature into a watch core v1 namespaced resource quota list handler
type WatchCoreV1NamespacedResourceQuotaListHandlerFunc func(WatchCoreV1NamespacedResourceQuotaListParams) middleware.Responder

// Handle executing the request and returning a response
func (fn WatchCoreV1NamespacedResourceQuotaListHandlerFunc) Handle(params WatchCoreV1NamespacedResourceQuotaListParams) middleware.Responder {
	return fn(params)
}

// WatchCoreV1NamespacedResourceQuotaListHandler interface for that can handle valid watch core v1 namespaced resource quota list params
type WatchCoreV1NamespacedResourceQuotaListHandler interface {
	Handle(WatchCoreV1NamespacedResourceQuotaListParams) middleware.Responder
}

// NewWatchCoreV1NamespacedResourceQuotaList creates a new http.Handler for the watch core v1 namespaced resource quota list operation
func NewWatchCoreV1NamespacedResourceQuotaList(ctx *middleware.Context, handler WatchCoreV1NamespacedResourceQuotaListHandler) *WatchCoreV1NamespacedResourceQuotaList {
	return &WatchCoreV1NamespacedResourceQuotaList{Context: ctx, Handler: handler}
}

/*WatchCoreV1NamespacedResourceQuotaList swagger:route GET /api/v1/watch/namespaces/{namespace}/resourcequotas core_v1 watchCoreV1NamespacedResourceQuotaList

watch individual changes to a list of ResourceQuota. deprecated: use the 'watch' parameter with a list operation instead.

*/
type WatchCoreV1NamespacedResourceQuotaList struct {
	Context *middleware.Context
	Handler WatchCoreV1NamespacedResourceQuotaListHandler
}

func (o *WatchCoreV1NamespacedResourceQuotaList) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewWatchCoreV1NamespacedResourceQuotaListParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
