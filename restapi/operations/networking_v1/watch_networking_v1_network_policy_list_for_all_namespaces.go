// Code generated by go-swagger; DO NOT EDIT.

package networking_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// WatchNetworkingV1NetworkPolicyListForAllNamespacesHandlerFunc turns a function with the right signature into a watch networking v1 network policy list for all namespaces handler
type WatchNetworkingV1NetworkPolicyListForAllNamespacesHandlerFunc func(WatchNetworkingV1NetworkPolicyListForAllNamespacesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn WatchNetworkingV1NetworkPolicyListForAllNamespacesHandlerFunc) Handle(params WatchNetworkingV1NetworkPolicyListForAllNamespacesParams) middleware.Responder {
	return fn(params)
}

// WatchNetworkingV1NetworkPolicyListForAllNamespacesHandler interface for that can handle valid watch networking v1 network policy list for all namespaces params
type WatchNetworkingV1NetworkPolicyListForAllNamespacesHandler interface {
	Handle(WatchNetworkingV1NetworkPolicyListForAllNamespacesParams) middleware.Responder
}

// NewWatchNetworkingV1NetworkPolicyListForAllNamespaces creates a new http.Handler for the watch networking v1 network policy list for all namespaces operation
func NewWatchNetworkingV1NetworkPolicyListForAllNamespaces(ctx *middleware.Context, handler WatchNetworkingV1NetworkPolicyListForAllNamespacesHandler) *WatchNetworkingV1NetworkPolicyListForAllNamespaces {
	return &WatchNetworkingV1NetworkPolicyListForAllNamespaces{Context: ctx, Handler: handler}
}

/*WatchNetworkingV1NetworkPolicyListForAllNamespaces swagger:route GET /apis/networking.k8s.io/v1/watch/networkpolicies networking_v1 watchNetworkingV1NetworkPolicyListForAllNamespaces

watch individual changes to a list of NetworkPolicy. deprecated: use the 'watch' parameter with a list operation instead.

*/
type WatchNetworkingV1NetworkPolicyListForAllNamespaces struct {
	Context *middleware.Context
	Handler WatchNetworkingV1NetworkPolicyListForAllNamespacesHandler
}

func (o *WatchNetworkingV1NetworkPolicyListForAllNamespaces) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewWatchNetworkingV1NetworkPolicyListForAllNamespacesParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}