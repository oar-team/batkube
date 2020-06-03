// Code generated by go-swagger; DO NOT EDIT.

package discovery_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// WatchDiscoveryV1beta1EndpointSliceListForAllNamespacesHandlerFunc turns a function with the right signature into a watch discovery v1beta1 endpoint slice list for all namespaces handler
type WatchDiscoveryV1beta1EndpointSliceListForAllNamespacesHandlerFunc func(WatchDiscoveryV1beta1EndpointSliceListForAllNamespacesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn WatchDiscoveryV1beta1EndpointSliceListForAllNamespacesHandlerFunc) Handle(params WatchDiscoveryV1beta1EndpointSliceListForAllNamespacesParams) middleware.Responder {
	return fn(params)
}

// WatchDiscoveryV1beta1EndpointSliceListForAllNamespacesHandler interface for that can handle valid watch discovery v1beta1 endpoint slice list for all namespaces params
type WatchDiscoveryV1beta1EndpointSliceListForAllNamespacesHandler interface {
	Handle(WatchDiscoveryV1beta1EndpointSliceListForAllNamespacesParams) middleware.Responder
}

// NewWatchDiscoveryV1beta1EndpointSliceListForAllNamespaces creates a new http.Handler for the watch discovery v1beta1 endpoint slice list for all namespaces operation
func NewWatchDiscoveryV1beta1EndpointSliceListForAllNamespaces(ctx *middleware.Context, handler WatchDiscoveryV1beta1EndpointSliceListForAllNamespacesHandler) *WatchDiscoveryV1beta1EndpointSliceListForAllNamespaces {
	return &WatchDiscoveryV1beta1EndpointSliceListForAllNamespaces{Context: ctx, Handler: handler}
}

/*WatchDiscoveryV1beta1EndpointSliceListForAllNamespaces swagger:route GET /apis/discovery.k8s.io/v1beta1/watch/endpointslices discovery_v1beta1 watchDiscoveryV1beta1EndpointSliceListForAllNamespaces

watch individual changes to a list of EndpointSlice. deprecated: use the 'watch' parameter with a list operation instead.

*/
type WatchDiscoveryV1beta1EndpointSliceListForAllNamespaces struct {
	Context *middleware.Context
	Handler WatchDiscoveryV1beta1EndpointSliceListForAllNamespacesHandler
}

func (o *WatchDiscoveryV1beta1EndpointSliceListForAllNamespaces) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewWatchDiscoveryV1beta1EndpointSliceListForAllNamespacesParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}