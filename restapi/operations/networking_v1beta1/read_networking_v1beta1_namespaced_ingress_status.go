// Code generated by go-swagger; DO NOT EDIT.

package networking_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ReadNetworkingV1beta1NamespacedIngressStatusHandlerFunc turns a function with the right signature into a read networking v1beta1 namespaced ingress status handler
type ReadNetworkingV1beta1NamespacedIngressStatusHandlerFunc func(ReadNetworkingV1beta1NamespacedIngressStatusParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ReadNetworkingV1beta1NamespacedIngressStatusHandlerFunc) Handle(params ReadNetworkingV1beta1NamespacedIngressStatusParams) middleware.Responder {
	return fn(params)
}

// ReadNetworkingV1beta1NamespacedIngressStatusHandler interface for that can handle valid read networking v1beta1 namespaced ingress status params
type ReadNetworkingV1beta1NamespacedIngressStatusHandler interface {
	Handle(ReadNetworkingV1beta1NamespacedIngressStatusParams) middleware.Responder
}

// NewReadNetworkingV1beta1NamespacedIngressStatus creates a new http.Handler for the read networking v1beta1 namespaced ingress status operation
func NewReadNetworkingV1beta1NamespacedIngressStatus(ctx *middleware.Context, handler ReadNetworkingV1beta1NamespacedIngressStatusHandler) *ReadNetworkingV1beta1NamespacedIngressStatus {
	return &ReadNetworkingV1beta1NamespacedIngressStatus{Context: ctx, Handler: handler}
}

/*ReadNetworkingV1beta1NamespacedIngressStatus swagger:route GET /apis/networking.k8s.io/v1beta1/namespaces/{namespace}/ingresses/{name}/status networking_v1beta1 readNetworkingV1beta1NamespacedIngressStatus

read status of the specified Ingress

*/
type ReadNetworkingV1beta1NamespacedIngressStatus struct {
	Context *middleware.Context
	Handler ReadNetworkingV1beta1NamespacedIngressStatusHandler
}

func (o *ReadNetworkingV1beta1NamespacedIngressStatus) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewReadNetworkingV1beta1NamespacedIngressStatusParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}