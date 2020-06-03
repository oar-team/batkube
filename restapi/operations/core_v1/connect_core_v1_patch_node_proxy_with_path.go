// Code generated by go-swagger; DO NOT EDIT.

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ConnectCoreV1PatchNodeProxyWithPathHandlerFunc turns a function with the right signature into a connect core v1 patch node proxy with path handler
type ConnectCoreV1PatchNodeProxyWithPathHandlerFunc func(ConnectCoreV1PatchNodeProxyWithPathParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ConnectCoreV1PatchNodeProxyWithPathHandlerFunc) Handle(params ConnectCoreV1PatchNodeProxyWithPathParams) middleware.Responder {
	return fn(params)
}

// ConnectCoreV1PatchNodeProxyWithPathHandler interface for that can handle valid connect core v1 patch node proxy with path params
type ConnectCoreV1PatchNodeProxyWithPathHandler interface {
	Handle(ConnectCoreV1PatchNodeProxyWithPathParams) middleware.Responder
}

// NewConnectCoreV1PatchNodeProxyWithPath creates a new http.Handler for the connect core v1 patch node proxy with path operation
func NewConnectCoreV1PatchNodeProxyWithPath(ctx *middleware.Context, handler ConnectCoreV1PatchNodeProxyWithPathHandler) *ConnectCoreV1PatchNodeProxyWithPath {
	return &ConnectCoreV1PatchNodeProxyWithPath{Context: ctx, Handler: handler}
}

/*ConnectCoreV1PatchNodeProxyWithPath swagger:route PATCH /api/v1/nodes/{name}/proxy/{path} core_v1 connectCoreV1PatchNodeProxyWithPath

connect PATCH requests to proxy of Node

*/
type ConnectCoreV1PatchNodeProxyWithPath struct {
	Context *middleware.Context
	Handler ConnectCoreV1PatchNodeProxyWithPathHandler
}

func (o *ConnectCoreV1PatchNodeProxyWithPath) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewConnectCoreV1PatchNodeProxyWithPathParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
