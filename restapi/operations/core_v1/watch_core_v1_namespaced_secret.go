// Code generated by go-swagger; DO NOT EDIT.

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// WatchCoreV1NamespacedSecretHandlerFunc turns a function with the right signature into a watch core v1 namespaced secret handler
type WatchCoreV1NamespacedSecretHandlerFunc func(WatchCoreV1NamespacedSecretParams) middleware.Responder

// Handle executing the request and returning a response
func (fn WatchCoreV1NamespacedSecretHandlerFunc) Handle(params WatchCoreV1NamespacedSecretParams) middleware.Responder {
	return fn(params)
}

// WatchCoreV1NamespacedSecretHandler interface for that can handle valid watch core v1 namespaced secret params
type WatchCoreV1NamespacedSecretHandler interface {
	Handle(WatchCoreV1NamespacedSecretParams) middleware.Responder
}

// NewWatchCoreV1NamespacedSecret creates a new http.Handler for the watch core v1 namespaced secret operation
func NewWatchCoreV1NamespacedSecret(ctx *middleware.Context, handler WatchCoreV1NamespacedSecretHandler) *WatchCoreV1NamespacedSecret {
	return &WatchCoreV1NamespacedSecret{Context: ctx, Handler: handler}
}

/*WatchCoreV1NamespacedSecret swagger:route GET /api/v1/watch/namespaces/{namespace}/secrets/{name} core_v1 watchCoreV1NamespacedSecret

watch changes to an object of kind Secret. deprecated: use the 'watch' parameter with a list operation instead, filtered to a single item with the 'fieldSelector' parameter.

*/
type WatchCoreV1NamespacedSecret struct {
	Context *middleware.Context
	Handler WatchCoreV1NamespacedSecretHandler
}

func (o *WatchCoreV1NamespacedSecret) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewWatchCoreV1NamespacedSecretParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}