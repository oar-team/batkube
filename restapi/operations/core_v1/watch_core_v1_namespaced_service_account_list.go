// Code generated by go-swagger; DO NOT EDIT.

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// WatchCoreV1NamespacedServiceAccountListHandlerFunc turns a function with the right signature into a watch core v1 namespaced service account list handler
type WatchCoreV1NamespacedServiceAccountListHandlerFunc func(WatchCoreV1NamespacedServiceAccountListParams) middleware.Responder

// Handle executing the request and returning a response
func (fn WatchCoreV1NamespacedServiceAccountListHandlerFunc) Handle(params WatchCoreV1NamespacedServiceAccountListParams) middleware.Responder {
	return fn(params)
}

// WatchCoreV1NamespacedServiceAccountListHandler interface for that can handle valid watch core v1 namespaced service account list params
type WatchCoreV1NamespacedServiceAccountListHandler interface {
	Handle(WatchCoreV1NamespacedServiceAccountListParams) middleware.Responder
}

// NewWatchCoreV1NamespacedServiceAccountList creates a new http.Handler for the watch core v1 namespaced service account list operation
func NewWatchCoreV1NamespacedServiceAccountList(ctx *middleware.Context, handler WatchCoreV1NamespacedServiceAccountListHandler) *WatchCoreV1NamespacedServiceAccountList {
	return &WatchCoreV1NamespacedServiceAccountList{Context: ctx, Handler: handler}
}

/*WatchCoreV1NamespacedServiceAccountList swagger:route GET /api/v1/watch/namespaces/{namespace}/serviceaccounts core_v1 watchCoreV1NamespacedServiceAccountList

watch individual changes to a list of ServiceAccount. deprecated: use the 'watch' parameter with a list operation instead.

*/
type WatchCoreV1NamespacedServiceAccountList struct {
	Context *middleware.Context
	Handler WatchCoreV1NamespacedServiceAccountListHandler
}

func (o *WatchCoreV1NamespacedServiceAccountList) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewWatchCoreV1NamespacedServiceAccountListParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
