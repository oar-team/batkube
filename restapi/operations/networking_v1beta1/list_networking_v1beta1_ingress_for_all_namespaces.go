// Code generated by go-swagger; DO NOT EDIT.

package networking_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ListNetworkingV1beta1IngressForAllNamespacesHandlerFunc turns a function with the right signature into a list networking v1beta1 ingress for all namespaces handler
type ListNetworkingV1beta1IngressForAllNamespacesHandlerFunc func(ListNetworkingV1beta1IngressForAllNamespacesParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn ListNetworkingV1beta1IngressForAllNamespacesHandlerFunc) Handle(params ListNetworkingV1beta1IngressForAllNamespacesParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// ListNetworkingV1beta1IngressForAllNamespacesHandler interface for that can handle valid list networking v1beta1 ingress for all namespaces params
type ListNetworkingV1beta1IngressForAllNamespacesHandler interface {
	Handle(ListNetworkingV1beta1IngressForAllNamespacesParams, interface{}) middleware.Responder
}

// NewListNetworkingV1beta1IngressForAllNamespaces creates a new http.Handler for the list networking v1beta1 ingress for all namespaces operation
func NewListNetworkingV1beta1IngressForAllNamespaces(ctx *middleware.Context, handler ListNetworkingV1beta1IngressForAllNamespacesHandler) *ListNetworkingV1beta1IngressForAllNamespaces {
	return &ListNetworkingV1beta1IngressForAllNamespaces{Context: ctx, Handler: handler}
}

/*ListNetworkingV1beta1IngressForAllNamespaces swagger:route GET /apis/networking.k8s.io/v1beta1/ingresses networking_v1beta1 listNetworkingV1beta1IngressForAllNamespaces

list or watch objects of kind Ingress

*/
type ListNetworkingV1beta1IngressForAllNamespaces struct {
	Context *middleware.Context
	Handler ListNetworkingV1beta1IngressForAllNamespacesHandler
}

func (o *ListNetworkingV1beta1IngressForAllNamespaces) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewListNetworkingV1beta1IngressForAllNamespacesParams()

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
