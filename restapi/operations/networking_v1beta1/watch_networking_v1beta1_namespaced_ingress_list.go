// Code generated by go-swagger; DO NOT EDIT.

package networking_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// WatchNetworkingV1beta1NamespacedIngressListHandlerFunc turns a function with the right signature into a watch networking v1beta1 namespaced ingress list handler
type WatchNetworkingV1beta1NamespacedIngressListHandlerFunc func(WatchNetworkingV1beta1NamespacedIngressListParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn WatchNetworkingV1beta1NamespacedIngressListHandlerFunc) Handle(params WatchNetworkingV1beta1NamespacedIngressListParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// WatchNetworkingV1beta1NamespacedIngressListHandler interface for that can handle valid watch networking v1beta1 namespaced ingress list params
type WatchNetworkingV1beta1NamespacedIngressListHandler interface {
	Handle(WatchNetworkingV1beta1NamespacedIngressListParams, interface{}) middleware.Responder
}

// NewWatchNetworkingV1beta1NamespacedIngressList creates a new http.Handler for the watch networking v1beta1 namespaced ingress list operation
func NewWatchNetworkingV1beta1NamespacedIngressList(ctx *middleware.Context, handler WatchNetworkingV1beta1NamespacedIngressListHandler) *WatchNetworkingV1beta1NamespacedIngressList {
	return &WatchNetworkingV1beta1NamespacedIngressList{Context: ctx, Handler: handler}
}

/*WatchNetworkingV1beta1NamespacedIngressList swagger:route GET /apis/networking.k8s.io/v1beta1/watch/namespaces/{namespace}/ingresses networking_v1beta1 watchNetworkingV1beta1NamespacedIngressList

watch individual changes to a list of Ingress. deprecated: use the 'watch' parameter with a list operation instead.

*/
type WatchNetworkingV1beta1NamespacedIngressList struct {
	Context *middleware.Context
	Handler WatchNetworkingV1beta1NamespacedIngressListHandler
}

func (o *WatchNetworkingV1beta1NamespacedIngressList) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewWatchNetworkingV1beta1NamespacedIngressListParams()

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
