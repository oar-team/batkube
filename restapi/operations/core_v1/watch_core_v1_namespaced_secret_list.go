// Code generated by go-swagger; DO NOT EDIT.

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// WatchCoreV1NamespacedSecretListHandlerFunc turns a function with the right signature into a watch core v1 namespaced secret list handler
type WatchCoreV1NamespacedSecretListHandlerFunc func(WatchCoreV1NamespacedSecretListParams) middleware.Responder

// Handle executing the request and returning a response
func (fn WatchCoreV1NamespacedSecretListHandlerFunc) Handle(params WatchCoreV1NamespacedSecretListParams) middleware.Responder {
	return fn(params)
}

// WatchCoreV1NamespacedSecretListHandler interface for that can handle valid watch core v1 namespaced secret list params
type WatchCoreV1NamespacedSecretListHandler interface {
	Handle(WatchCoreV1NamespacedSecretListParams) middleware.Responder
}

// NewWatchCoreV1NamespacedSecretList creates a new http.Handler for the watch core v1 namespaced secret list operation
func NewWatchCoreV1NamespacedSecretList(ctx *middleware.Context, handler WatchCoreV1NamespacedSecretListHandler) *WatchCoreV1NamespacedSecretList {
	return &WatchCoreV1NamespacedSecretList{Context: ctx, Handler: handler}
}

/*WatchCoreV1NamespacedSecretList swagger:route GET /api/v1/watch/namespaces/{namespace}/secrets core_v1 watchCoreV1NamespacedSecretList

watch individual changes to a list of Secret. deprecated: use the 'watch' parameter with a list operation instead.

*/
type WatchCoreV1NamespacedSecretList struct {
	Context *middleware.Context
	Handler WatchCoreV1NamespacedSecretListHandler
}

func (o *WatchCoreV1NamespacedSecretList) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewWatchCoreV1NamespacedSecretListParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
