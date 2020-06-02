// Code generated by go-swagger; DO NOT EDIT.

package discovery_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// WatchDiscoveryV1beta1NamespacedEndpointSliceHandlerFunc turns a function with the right signature into a watch discovery v1beta1 namespaced endpoint slice handler
type WatchDiscoveryV1beta1NamespacedEndpointSliceHandlerFunc func(WatchDiscoveryV1beta1NamespacedEndpointSliceParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn WatchDiscoveryV1beta1NamespacedEndpointSliceHandlerFunc) Handle(params WatchDiscoveryV1beta1NamespacedEndpointSliceParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// WatchDiscoveryV1beta1NamespacedEndpointSliceHandler interface for that can handle valid watch discovery v1beta1 namespaced endpoint slice params
type WatchDiscoveryV1beta1NamespacedEndpointSliceHandler interface {
	Handle(WatchDiscoveryV1beta1NamespacedEndpointSliceParams, interface{}) middleware.Responder
}

// NewWatchDiscoveryV1beta1NamespacedEndpointSlice creates a new http.Handler for the watch discovery v1beta1 namespaced endpoint slice operation
func NewWatchDiscoveryV1beta1NamespacedEndpointSlice(ctx *middleware.Context, handler WatchDiscoveryV1beta1NamespacedEndpointSliceHandler) *WatchDiscoveryV1beta1NamespacedEndpointSlice {
	return &WatchDiscoveryV1beta1NamespacedEndpointSlice{Context: ctx, Handler: handler}
}

/*WatchDiscoveryV1beta1NamespacedEndpointSlice swagger:route GET /apis/discovery.k8s.io/v1beta1/watch/namespaces/{namespace}/endpointslices/{name} discovery_v1beta1 watchDiscoveryV1beta1NamespacedEndpointSlice

watch changes to an object of kind EndpointSlice. deprecated: use the 'watch' parameter with a list operation instead, filtered to a single item with the 'fieldSelector' parameter.

*/
type WatchDiscoveryV1beta1NamespacedEndpointSlice struct {
	Context *middleware.Context
	Handler WatchDiscoveryV1beta1NamespacedEndpointSliceHandler
}

func (o *WatchDiscoveryV1beta1NamespacedEndpointSlice) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewWatchDiscoveryV1beta1NamespacedEndpointSliceParams()

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
