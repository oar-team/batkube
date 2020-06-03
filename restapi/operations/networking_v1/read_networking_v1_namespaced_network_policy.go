// Code generated by go-swagger; DO NOT EDIT.

package networking_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ReadNetworkingV1NamespacedNetworkPolicyHandlerFunc turns a function with the right signature into a read networking v1 namespaced network policy handler
type ReadNetworkingV1NamespacedNetworkPolicyHandlerFunc func(ReadNetworkingV1NamespacedNetworkPolicyParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ReadNetworkingV1NamespacedNetworkPolicyHandlerFunc) Handle(params ReadNetworkingV1NamespacedNetworkPolicyParams) middleware.Responder {
	return fn(params)
}

// ReadNetworkingV1NamespacedNetworkPolicyHandler interface for that can handle valid read networking v1 namespaced network policy params
type ReadNetworkingV1NamespacedNetworkPolicyHandler interface {
	Handle(ReadNetworkingV1NamespacedNetworkPolicyParams) middleware.Responder
}

// NewReadNetworkingV1NamespacedNetworkPolicy creates a new http.Handler for the read networking v1 namespaced network policy operation
func NewReadNetworkingV1NamespacedNetworkPolicy(ctx *middleware.Context, handler ReadNetworkingV1NamespacedNetworkPolicyHandler) *ReadNetworkingV1NamespacedNetworkPolicy {
	return &ReadNetworkingV1NamespacedNetworkPolicy{Context: ctx, Handler: handler}
}

/*ReadNetworkingV1NamespacedNetworkPolicy swagger:route GET /apis/networking.k8s.io/v1/namespaces/{namespace}/networkpolicies/{name} networking_v1 readNetworkingV1NamespacedNetworkPolicy

read the specified NetworkPolicy

*/
type ReadNetworkingV1NamespacedNetworkPolicy struct {
	Context *middleware.Context
	Handler ReadNetworkingV1NamespacedNetworkPolicyHandler
}

func (o *ReadNetworkingV1NamespacedNetworkPolicy) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewReadNetworkingV1NamespacedNetworkPolicyParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
