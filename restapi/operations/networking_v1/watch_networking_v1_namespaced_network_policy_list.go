// Code generated by go-swagger; DO NOT EDIT.

package networking_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// WatchNetworkingV1NamespacedNetworkPolicyListHandlerFunc turns a function with the right signature into a watch networking v1 namespaced network policy list handler
type WatchNetworkingV1NamespacedNetworkPolicyListHandlerFunc func(WatchNetworkingV1NamespacedNetworkPolicyListParams) middleware.Responder

// Handle executing the request and returning a response
func (fn WatchNetworkingV1NamespacedNetworkPolicyListHandlerFunc) Handle(params WatchNetworkingV1NamespacedNetworkPolicyListParams) middleware.Responder {
	return fn(params)
}

// WatchNetworkingV1NamespacedNetworkPolicyListHandler interface for that can handle valid watch networking v1 namespaced network policy list params
type WatchNetworkingV1NamespacedNetworkPolicyListHandler interface {
	Handle(WatchNetworkingV1NamespacedNetworkPolicyListParams) middleware.Responder
}

// NewWatchNetworkingV1NamespacedNetworkPolicyList creates a new http.Handler for the watch networking v1 namespaced network policy list operation
func NewWatchNetworkingV1NamespacedNetworkPolicyList(ctx *middleware.Context, handler WatchNetworkingV1NamespacedNetworkPolicyListHandler) *WatchNetworkingV1NamespacedNetworkPolicyList {
	return &WatchNetworkingV1NamespacedNetworkPolicyList{Context: ctx, Handler: handler}
}

/*WatchNetworkingV1NamespacedNetworkPolicyList swagger:route GET /apis/networking.k8s.io/v1/watch/namespaces/{namespace}/networkpolicies networking_v1 watchNetworkingV1NamespacedNetworkPolicyList

watch individual changes to a list of NetworkPolicy. deprecated: use the 'watch' parameter with a list operation instead.

*/
type WatchNetworkingV1NamespacedNetworkPolicyList struct {
	Context *middleware.Context
	Handler WatchNetworkingV1NamespacedNetworkPolicyListHandler
}

func (o *WatchNetworkingV1NamespacedNetworkPolicyList) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewWatchNetworkingV1NamespacedNetworkPolicyListParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
