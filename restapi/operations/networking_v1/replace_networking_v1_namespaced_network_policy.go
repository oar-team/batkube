// Code generated by go-swagger; DO NOT EDIT.

package networking_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ReplaceNetworkingV1NamespacedNetworkPolicyHandlerFunc turns a function with the right signature into a replace networking v1 namespaced network policy handler
type ReplaceNetworkingV1NamespacedNetworkPolicyHandlerFunc func(ReplaceNetworkingV1NamespacedNetworkPolicyParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ReplaceNetworkingV1NamespacedNetworkPolicyHandlerFunc) Handle(params ReplaceNetworkingV1NamespacedNetworkPolicyParams) middleware.Responder {
	return fn(params)
}

// ReplaceNetworkingV1NamespacedNetworkPolicyHandler interface for that can handle valid replace networking v1 namespaced network policy params
type ReplaceNetworkingV1NamespacedNetworkPolicyHandler interface {
	Handle(ReplaceNetworkingV1NamespacedNetworkPolicyParams) middleware.Responder
}

// NewReplaceNetworkingV1NamespacedNetworkPolicy creates a new http.Handler for the replace networking v1 namespaced network policy operation
func NewReplaceNetworkingV1NamespacedNetworkPolicy(ctx *middleware.Context, handler ReplaceNetworkingV1NamespacedNetworkPolicyHandler) *ReplaceNetworkingV1NamespacedNetworkPolicy {
	return &ReplaceNetworkingV1NamespacedNetworkPolicy{Context: ctx, Handler: handler}
}

/*ReplaceNetworkingV1NamespacedNetworkPolicy swagger:route PUT /apis/networking.k8s.io/v1/namespaces/{namespace}/networkpolicies/{name} networking_v1 replaceNetworkingV1NamespacedNetworkPolicy

replace the specified NetworkPolicy

*/
type ReplaceNetworkingV1NamespacedNetworkPolicy struct {
	Context *middleware.Context
	Handler ReplaceNetworkingV1NamespacedNetworkPolicyHandler
}

func (o *ReplaceNetworkingV1NamespacedNetworkPolicy) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewReplaceNetworkingV1NamespacedNetworkPolicyParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
