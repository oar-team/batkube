// Code generated by go-swagger; DO NOT EDIT.

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ConnectCoreV1PostNamespacedPodExecHandlerFunc turns a function with the right signature into a connect core v1 post namespaced pod exec handler
type ConnectCoreV1PostNamespacedPodExecHandlerFunc func(ConnectCoreV1PostNamespacedPodExecParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ConnectCoreV1PostNamespacedPodExecHandlerFunc) Handle(params ConnectCoreV1PostNamespacedPodExecParams) middleware.Responder {
	return fn(params)
}

// ConnectCoreV1PostNamespacedPodExecHandler interface for that can handle valid connect core v1 post namespaced pod exec params
type ConnectCoreV1PostNamespacedPodExecHandler interface {
	Handle(ConnectCoreV1PostNamespacedPodExecParams) middleware.Responder
}

// NewConnectCoreV1PostNamespacedPodExec creates a new http.Handler for the connect core v1 post namespaced pod exec operation
func NewConnectCoreV1PostNamespacedPodExec(ctx *middleware.Context, handler ConnectCoreV1PostNamespacedPodExecHandler) *ConnectCoreV1PostNamespacedPodExec {
	return &ConnectCoreV1PostNamespacedPodExec{Context: ctx, Handler: handler}
}

/*ConnectCoreV1PostNamespacedPodExec swagger:route POST /api/v1/namespaces/{namespace}/pods/{name}/exec core_v1 connectCoreV1PostNamespacedPodExec

connect POST requests to exec of Pod

*/
type ConnectCoreV1PostNamespacedPodExec struct {
	Context *middleware.Context
	Handler ConnectCoreV1PostNamespacedPodExecHandler
}

func (o *ConnectCoreV1PostNamespacedPodExec) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewConnectCoreV1PostNamespacedPodExecParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
