// Code generated by go-swagger; DO NOT EDIT.

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ConnectCoreV1HeadNamespacedServiceProxyWithPathHandlerFunc turns a function with the right signature into a connect core v1 head namespaced service proxy with path handler
type ConnectCoreV1HeadNamespacedServiceProxyWithPathHandlerFunc func(ConnectCoreV1HeadNamespacedServiceProxyWithPathParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ConnectCoreV1HeadNamespacedServiceProxyWithPathHandlerFunc) Handle(params ConnectCoreV1HeadNamespacedServiceProxyWithPathParams) middleware.Responder {
	return fn(params)
}

// ConnectCoreV1HeadNamespacedServiceProxyWithPathHandler interface for that can handle valid connect core v1 head namespaced service proxy with path params
type ConnectCoreV1HeadNamespacedServiceProxyWithPathHandler interface {
	Handle(ConnectCoreV1HeadNamespacedServiceProxyWithPathParams) middleware.Responder
}

// NewConnectCoreV1HeadNamespacedServiceProxyWithPath creates a new http.Handler for the connect core v1 head namespaced service proxy with path operation
func NewConnectCoreV1HeadNamespacedServiceProxyWithPath(ctx *middleware.Context, handler ConnectCoreV1HeadNamespacedServiceProxyWithPathHandler) *ConnectCoreV1HeadNamespacedServiceProxyWithPath {
	return &ConnectCoreV1HeadNamespacedServiceProxyWithPath{Context: ctx, Handler: handler}
}

/*ConnectCoreV1HeadNamespacedServiceProxyWithPath swagger:route HEAD /api/v1/namespaces/{namespace}/services/{name}/proxy/{path} core_v1 connectCoreV1HeadNamespacedServiceProxyWithPath

connect HEAD requests to proxy of Service

*/
type ConnectCoreV1HeadNamespacedServiceProxyWithPath struct {
	Context *middleware.Context
	Handler ConnectCoreV1HeadNamespacedServiceProxyWithPathHandler
}

func (o *ConnectCoreV1HeadNamespacedServiceProxyWithPath) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewConnectCoreV1HeadNamespacedServiceProxyWithPathParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}