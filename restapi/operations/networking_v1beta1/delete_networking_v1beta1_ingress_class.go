// Code generated by go-swagger; DO NOT EDIT.

package networking_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteNetworkingV1beta1IngressClassHandlerFunc turns a function with the right signature into a delete networking v1beta1 ingress class handler
type DeleteNetworkingV1beta1IngressClassHandlerFunc func(DeleteNetworkingV1beta1IngressClassParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteNetworkingV1beta1IngressClassHandlerFunc) Handle(params DeleteNetworkingV1beta1IngressClassParams) middleware.Responder {
	return fn(params)
}

// DeleteNetworkingV1beta1IngressClassHandler interface for that can handle valid delete networking v1beta1 ingress class params
type DeleteNetworkingV1beta1IngressClassHandler interface {
	Handle(DeleteNetworkingV1beta1IngressClassParams) middleware.Responder
}

// NewDeleteNetworkingV1beta1IngressClass creates a new http.Handler for the delete networking v1beta1 ingress class operation
func NewDeleteNetworkingV1beta1IngressClass(ctx *middleware.Context, handler DeleteNetworkingV1beta1IngressClassHandler) *DeleteNetworkingV1beta1IngressClass {
	return &DeleteNetworkingV1beta1IngressClass{Context: ctx, Handler: handler}
}

/*DeleteNetworkingV1beta1IngressClass swagger:route DELETE /apis/networking.k8s.io/v1beta1/ingressclasses/{name} networking_v1beta1 deleteNetworkingV1beta1IngressClass

delete an IngressClass

*/
type DeleteNetworkingV1beta1IngressClass struct {
	Context *middleware.Context
	Handler DeleteNetworkingV1beta1IngressClassHandler
}

func (o *DeleteNetworkingV1beta1IngressClass) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteNetworkingV1beta1IngressClassParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
