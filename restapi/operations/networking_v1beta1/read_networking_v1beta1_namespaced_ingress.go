// Code generated by go-swagger; DO NOT EDIT.

package networking_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ReadNetworkingV1beta1NamespacedIngressHandlerFunc turns a function with the right signature into a read networking v1beta1 namespaced ingress handler
type ReadNetworkingV1beta1NamespacedIngressHandlerFunc func(ReadNetworkingV1beta1NamespacedIngressParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ReadNetworkingV1beta1NamespacedIngressHandlerFunc) Handle(params ReadNetworkingV1beta1NamespacedIngressParams) middleware.Responder {
	return fn(params)
}

// ReadNetworkingV1beta1NamespacedIngressHandler interface for that can handle valid read networking v1beta1 namespaced ingress params
type ReadNetworkingV1beta1NamespacedIngressHandler interface {
	Handle(ReadNetworkingV1beta1NamespacedIngressParams) middleware.Responder
}

// NewReadNetworkingV1beta1NamespacedIngress creates a new http.Handler for the read networking v1beta1 namespaced ingress operation
func NewReadNetworkingV1beta1NamespacedIngress(ctx *middleware.Context, handler ReadNetworkingV1beta1NamespacedIngressHandler) *ReadNetworkingV1beta1NamespacedIngress {
	return &ReadNetworkingV1beta1NamespacedIngress{Context: ctx, Handler: handler}
}

/*ReadNetworkingV1beta1NamespacedIngress swagger:route GET /apis/networking.k8s.io/v1beta1/namespaces/{namespace}/ingresses/{name} networking_v1beta1 readNetworkingV1beta1NamespacedIngress

read the specified Ingress

*/
type ReadNetworkingV1beta1NamespacedIngress struct {
	Context *middleware.Context
	Handler ReadNetworkingV1beta1NamespacedIngressHandler
}

func (o *ReadNetworkingV1beta1NamespacedIngress) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewReadNetworkingV1beta1NamespacedIngressParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
