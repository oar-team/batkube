// Code generated by go-swagger; DO NOT EDIT.

package apps_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ReplaceAppsV1NamespacedDaemonSetStatusHandlerFunc turns a function with the right signature into a replace apps v1 namespaced daemon set status handler
type ReplaceAppsV1NamespacedDaemonSetStatusHandlerFunc func(ReplaceAppsV1NamespacedDaemonSetStatusParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ReplaceAppsV1NamespacedDaemonSetStatusHandlerFunc) Handle(params ReplaceAppsV1NamespacedDaemonSetStatusParams) middleware.Responder {
	return fn(params)
}

// ReplaceAppsV1NamespacedDaemonSetStatusHandler interface for that can handle valid replace apps v1 namespaced daemon set status params
type ReplaceAppsV1NamespacedDaemonSetStatusHandler interface {
	Handle(ReplaceAppsV1NamespacedDaemonSetStatusParams) middleware.Responder
}

// NewReplaceAppsV1NamespacedDaemonSetStatus creates a new http.Handler for the replace apps v1 namespaced daemon set status operation
func NewReplaceAppsV1NamespacedDaemonSetStatus(ctx *middleware.Context, handler ReplaceAppsV1NamespacedDaemonSetStatusHandler) *ReplaceAppsV1NamespacedDaemonSetStatus {
	return &ReplaceAppsV1NamespacedDaemonSetStatus{Context: ctx, Handler: handler}
}

/*ReplaceAppsV1NamespacedDaemonSetStatus swagger:route PUT /apis/apps/v1/namespaces/{namespace}/daemonsets/{name}/status apps_v1 replaceAppsV1NamespacedDaemonSetStatus

replace status of the specified DaemonSet

*/
type ReplaceAppsV1NamespacedDaemonSetStatus struct {
	Context *middleware.Context
	Handler ReplaceAppsV1NamespacedDaemonSetStatusHandler
}

func (o *ReplaceAppsV1NamespacedDaemonSetStatus) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewReplaceAppsV1NamespacedDaemonSetStatusParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
