// Code generated by go-swagger; DO NOT EDIT.

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ConnectCoreV1PutNodeProxyWithPathHandlerFunc turns a function with the right signature into a connect core v1 put node proxy with path handler
type ConnectCoreV1PutNodeProxyWithPathHandlerFunc func(ConnectCoreV1PutNodeProxyWithPathParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ConnectCoreV1PutNodeProxyWithPathHandlerFunc) Handle(params ConnectCoreV1PutNodeProxyWithPathParams) middleware.Responder {
	return fn(params)
}

// ConnectCoreV1PutNodeProxyWithPathHandler interface for that can handle valid connect core v1 put node proxy with path params
type ConnectCoreV1PutNodeProxyWithPathHandler interface {
	Handle(ConnectCoreV1PutNodeProxyWithPathParams) middleware.Responder
}

// NewConnectCoreV1PutNodeProxyWithPath creates a new http.Handler for the connect core v1 put node proxy with path operation
func NewConnectCoreV1PutNodeProxyWithPath(ctx *middleware.Context, handler ConnectCoreV1PutNodeProxyWithPathHandler) *ConnectCoreV1PutNodeProxyWithPath {
	return &ConnectCoreV1PutNodeProxyWithPath{Context: ctx, Handler: handler}
}

/*ConnectCoreV1PutNodeProxyWithPath swagger:route PUT /api/v1/nodes/{name}/proxy/{path} core_v1 connectCoreV1PutNodeProxyWithPath

connect PUT requests to proxy of Node

*/
type ConnectCoreV1PutNodeProxyWithPath struct {
	Context *middleware.Context
	Handler ConnectCoreV1PutNodeProxyWithPathHandler
}

func (o *ConnectCoreV1PutNodeProxyWithPath) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewConnectCoreV1PutNodeProxyWithPathParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}