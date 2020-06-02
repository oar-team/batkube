// Code generated by go-swagger; DO NOT EDIT.

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ConnectCoreV1HeadNamespacedPodProxyHandlerFunc turns a function with the right signature into a connect core v1 head namespaced pod proxy handler
type ConnectCoreV1HeadNamespacedPodProxyHandlerFunc func(ConnectCoreV1HeadNamespacedPodProxyParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn ConnectCoreV1HeadNamespacedPodProxyHandlerFunc) Handle(params ConnectCoreV1HeadNamespacedPodProxyParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// ConnectCoreV1HeadNamespacedPodProxyHandler interface for that can handle valid connect core v1 head namespaced pod proxy params
type ConnectCoreV1HeadNamespacedPodProxyHandler interface {
	Handle(ConnectCoreV1HeadNamespacedPodProxyParams, interface{}) middleware.Responder
}

// NewConnectCoreV1HeadNamespacedPodProxy creates a new http.Handler for the connect core v1 head namespaced pod proxy operation
func NewConnectCoreV1HeadNamespacedPodProxy(ctx *middleware.Context, handler ConnectCoreV1HeadNamespacedPodProxyHandler) *ConnectCoreV1HeadNamespacedPodProxy {
	return &ConnectCoreV1HeadNamespacedPodProxy{Context: ctx, Handler: handler}
}

/*ConnectCoreV1HeadNamespacedPodProxy swagger:route HEAD /api/v1/namespaces/{namespace}/pods/{name}/proxy core_v1 connectCoreV1HeadNamespacedPodProxy

connect HEAD requests to proxy of Pod

*/
type ConnectCoreV1HeadNamespacedPodProxy struct {
	Context *middleware.Context
	Handler ConnectCoreV1HeadNamespacedPodProxyHandler
}

func (o *ConnectCoreV1HeadNamespacedPodProxy) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewConnectCoreV1HeadNamespacedPodProxyParams()

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
