// Code generated by go-swagger; DO NOT EDIT.

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// WatchCoreV1NamespacedServiceHandlerFunc turns a function with the right signature into a watch core v1 namespaced service handler
type WatchCoreV1NamespacedServiceHandlerFunc func(WatchCoreV1NamespacedServiceParams) middleware.Responder

// Handle executing the request and returning a response
func (fn WatchCoreV1NamespacedServiceHandlerFunc) Handle(params WatchCoreV1NamespacedServiceParams) middleware.Responder {
	return fn(params)
}

// WatchCoreV1NamespacedServiceHandler interface for that can handle valid watch core v1 namespaced service params
type WatchCoreV1NamespacedServiceHandler interface {
	Handle(WatchCoreV1NamespacedServiceParams) middleware.Responder
}

// NewWatchCoreV1NamespacedService creates a new http.Handler for the watch core v1 namespaced service operation
func NewWatchCoreV1NamespacedService(ctx *middleware.Context, handler WatchCoreV1NamespacedServiceHandler) *WatchCoreV1NamespacedService {
	return &WatchCoreV1NamespacedService{Context: ctx, Handler: handler}
}

/*WatchCoreV1NamespacedService swagger:route GET /api/v1/watch/namespaces/{namespace}/services/{name} core_v1 watchCoreV1NamespacedService

watch changes to an object of kind Service. deprecated: use the 'watch' parameter with a list operation instead, filtered to a single item with the 'fieldSelector' parameter.

*/
type WatchCoreV1NamespacedService struct {
	Context *middleware.Context
	Handler WatchCoreV1NamespacedServiceHandler
}

func (o *WatchCoreV1NamespacedService) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewWatchCoreV1NamespacedServiceParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
