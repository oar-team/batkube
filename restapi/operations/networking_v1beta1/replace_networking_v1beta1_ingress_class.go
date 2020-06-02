// Code generated by go-swagger; DO NOT EDIT.

package networking_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ReplaceNetworkingV1beta1IngressClassHandlerFunc turns a function with the right signature into a replace networking v1beta1 ingress class handler
type ReplaceNetworkingV1beta1IngressClassHandlerFunc func(ReplaceNetworkingV1beta1IngressClassParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn ReplaceNetworkingV1beta1IngressClassHandlerFunc) Handle(params ReplaceNetworkingV1beta1IngressClassParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// ReplaceNetworkingV1beta1IngressClassHandler interface for that can handle valid replace networking v1beta1 ingress class params
type ReplaceNetworkingV1beta1IngressClassHandler interface {
	Handle(ReplaceNetworkingV1beta1IngressClassParams, interface{}) middleware.Responder
}

// NewReplaceNetworkingV1beta1IngressClass creates a new http.Handler for the replace networking v1beta1 ingress class operation
func NewReplaceNetworkingV1beta1IngressClass(ctx *middleware.Context, handler ReplaceNetworkingV1beta1IngressClassHandler) *ReplaceNetworkingV1beta1IngressClass {
	return &ReplaceNetworkingV1beta1IngressClass{Context: ctx, Handler: handler}
}

/*ReplaceNetworkingV1beta1IngressClass swagger:route PUT /apis/networking.k8s.io/v1beta1/ingressclasses/{name} networking_v1beta1 replaceNetworkingV1beta1IngressClass

replace the specified IngressClass

*/
type ReplaceNetworkingV1beta1IngressClass struct {
	Context *middleware.Context
	Handler ReplaceNetworkingV1beta1IngressClassHandler
}

func (o *ReplaceNetworkingV1beta1IngressClass) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewReplaceNetworkingV1beta1IngressClassParams()

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
