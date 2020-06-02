// Code generated by go-swagger; DO NOT EDIT.

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ConnectCoreV1PutNamespacedServiceProxyHandlerFunc turns a function with the right signature into a connect core v1 put namespaced service proxy handler
type ConnectCoreV1PutNamespacedServiceProxyHandlerFunc func(ConnectCoreV1PutNamespacedServiceProxyParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn ConnectCoreV1PutNamespacedServiceProxyHandlerFunc) Handle(params ConnectCoreV1PutNamespacedServiceProxyParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// ConnectCoreV1PutNamespacedServiceProxyHandler interface for that can handle valid connect core v1 put namespaced service proxy params
type ConnectCoreV1PutNamespacedServiceProxyHandler interface {
	Handle(ConnectCoreV1PutNamespacedServiceProxyParams, interface{}) middleware.Responder
}

// NewConnectCoreV1PutNamespacedServiceProxy creates a new http.Handler for the connect core v1 put namespaced service proxy operation
func NewConnectCoreV1PutNamespacedServiceProxy(ctx *middleware.Context, handler ConnectCoreV1PutNamespacedServiceProxyHandler) *ConnectCoreV1PutNamespacedServiceProxy {
	return &ConnectCoreV1PutNamespacedServiceProxy{Context: ctx, Handler: handler}
}

/*ConnectCoreV1PutNamespacedServiceProxy swagger:route PUT /api/v1/namespaces/{namespace}/services/{name}/proxy core_v1 connectCoreV1PutNamespacedServiceProxy

connect PUT requests to proxy of Service

*/
type ConnectCoreV1PutNamespacedServiceProxy struct {
	Context *middleware.Context
	Handler ConnectCoreV1PutNamespacedServiceProxyHandler
}

func (o *ConnectCoreV1PutNamespacedServiceProxy) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewConnectCoreV1PutNamespacedServiceProxyParams()

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
