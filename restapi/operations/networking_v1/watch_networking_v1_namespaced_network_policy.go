// Code generated by go-swagger; DO NOT EDIT.

package networking_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// WatchNetworkingV1NamespacedNetworkPolicyHandlerFunc turns a function with the right signature into a watch networking v1 namespaced network policy handler
type WatchNetworkingV1NamespacedNetworkPolicyHandlerFunc func(WatchNetworkingV1NamespacedNetworkPolicyParams) middleware.Responder

// Handle executing the request and returning a response
func (fn WatchNetworkingV1NamespacedNetworkPolicyHandlerFunc) Handle(params WatchNetworkingV1NamespacedNetworkPolicyParams) middleware.Responder {
	return fn(params)
}

// WatchNetworkingV1NamespacedNetworkPolicyHandler interface for that can handle valid watch networking v1 namespaced network policy params
type WatchNetworkingV1NamespacedNetworkPolicyHandler interface {
	Handle(WatchNetworkingV1NamespacedNetworkPolicyParams) middleware.Responder
}

// NewWatchNetworkingV1NamespacedNetworkPolicy creates a new http.Handler for the watch networking v1 namespaced network policy operation
func NewWatchNetworkingV1NamespacedNetworkPolicy(ctx *middleware.Context, handler WatchNetworkingV1NamespacedNetworkPolicyHandler) *WatchNetworkingV1NamespacedNetworkPolicy {
	return &WatchNetworkingV1NamespacedNetworkPolicy{Context: ctx, Handler: handler}
}

/*WatchNetworkingV1NamespacedNetworkPolicy swagger:route GET /apis/networking.k8s.io/v1/watch/namespaces/{namespace}/networkpolicies/{name} networking_v1 watchNetworkingV1NamespacedNetworkPolicy

watch changes to an object of kind NetworkPolicy. deprecated: use the 'watch' parameter with a list operation instead, filtered to a single item with the 'fieldSelector' parameter.

*/
type WatchNetworkingV1NamespacedNetworkPolicy struct {
	Context *middleware.Context
	Handler WatchNetworkingV1NamespacedNetworkPolicyHandler
}

func (o *WatchNetworkingV1NamespacedNetworkPolicy) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewWatchNetworkingV1NamespacedNetworkPolicyParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
