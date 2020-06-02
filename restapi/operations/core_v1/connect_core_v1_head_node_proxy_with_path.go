// Code generated by go-swagger; DO NOT EDIT.

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ConnectCoreV1HeadNodeProxyWithPathHandlerFunc turns a function with the right signature into a connect core v1 head node proxy with path handler
type ConnectCoreV1HeadNodeProxyWithPathHandlerFunc func(ConnectCoreV1HeadNodeProxyWithPathParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn ConnectCoreV1HeadNodeProxyWithPathHandlerFunc) Handle(params ConnectCoreV1HeadNodeProxyWithPathParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// ConnectCoreV1HeadNodeProxyWithPathHandler interface for that can handle valid connect core v1 head node proxy with path params
type ConnectCoreV1HeadNodeProxyWithPathHandler interface {
	Handle(ConnectCoreV1HeadNodeProxyWithPathParams, interface{}) middleware.Responder
}

// NewConnectCoreV1HeadNodeProxyWithPath creates a new http.Handler for the connect core v1 head node proxy with path operation
func NewConnectCoreV1HeadNodeProxyWithPath(ctx *middleware.Context, handler ConnectCoreV1HeadNodeProxyWithPathHandler) *ConnectCoreV1HeadNodeProxyWithPath {
	return &ConnectCoreV1HeadNodeProxyWithPath{Context: ctx, Handler: handler}
}

/*ConnectCoreV1HeadNodeProxyWithPath swagger:route HEAD /api/v1/nodes/{name}/proxy/{path} core_v1 connectCoreV1HeadNodeProxyWithPath

connect HEAD requests to proxy of Node

*/
type ConnectCoreV1HeadNodeProxyWithPath struct {
	Context *middleware.Context
	Handler ConnectCoreV1HeadNodeProxyWithPathHandler
}

func (o *ConnectCoreV1HeadNodeProxyWithPath) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewConnectCoreV1HeadNodeProxyWithPathParams()

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
