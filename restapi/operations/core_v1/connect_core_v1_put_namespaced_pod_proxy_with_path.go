// Code generated by go-swagger; DO NOT EDIT.

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ConnectCoreV1PutNamespacedPodProxyWithPathHandlerFunc turns a function with the right signature into a connect core v1 put namespaced pod proxy with path handler
type ConnectCoreV1PutNamespacedPodProxyWithPathHandlerFunc func(ConnectCoreV1PutNamespacedPodProxyWithPathParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ConnectCoreV1PutNamespacedPodProxyWithPathHandlerFunc) Handle(params ConnectCoreV1PutNamespacedPodProxyWithPathParams) middleware.Responder {
	return fn(params)
}

// ConnectCoreV1PutNamespacedPodProxyWithPathHandler interface for that can handle valid connect core v1 put namespaced pod proxy with path params
type ConnectCoreV1PutNamespacedPodProxyWithPathHandler interface {
	Handle(ConnectCoreV1PutNamespacedPodProxyWithPathParams) middleware.Responder
}

// NewConnectCoreV1PutNamespacedPodProxyWithPath creates a new http.Handler for the connect core v1 put namespaced pod proxy with path operation
func NewConnectCoreV1PutNamespacedPodProxyWithPath(ctx *middleware.Context, handler ConnectCoreV1PutNamespacedPodProxyWithPathHandler) *ConnectCoreV1PutNamespacedPodProxyWithPath {
	return &ConnectCoreV1PutNamespacedPodProxyWithPath{Context: ctx, Handler: handler}
}

/*ConnectCoreV1PutNamespacedPodProxyWithPath swagger:route PUT /api/v1/namespaces/{namespace}/pods/{name}/proxy/{path} core_v1 connectCoreV1PutNamespacedPodProxyWithPath

connect PUT requests to proxy of Pod

*/
type ConnectCoreV1PutNamespacedPodProxyWithPath struct {
	Context *middleware.Context
	Handler ConnectCoreV1PutNamespacedPodProxyWithPathHandler
}

func (o *ConnectCoreV1PutNamespacedPodProxyWithPath) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewConnectCoreV1PutNamespacedPodProxyWithPathParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
