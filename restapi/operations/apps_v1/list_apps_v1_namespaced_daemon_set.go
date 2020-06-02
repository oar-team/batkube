// Code generated by go-swagger; DO NOT EDIT.

package apps_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ListAppsV1NamespacedDaemonSetHandlerFunc turns a function with the right signature into a list apps v1 namespaced daemon set handler
type ListAppsV1NamespacedDaemonSetHandlerFunc func(ListAppsV1NamespacedDaemonSetParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn ListAppsV1NamespacedDaemonSetHandlerFunc) Handle(params ListAppsV1NamespacedDaemonSetParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// ListAppsV1NamespacedDaemonSetHandler interface for that can handle valid list apps v1 namespaced daemon set params
type ListAppsV1NamespacedDaemonSetHandler interface {
	Handle(ListAppsV1NamespacedDaemonSetParams, interface{}) middleware.Responder
}

// NewListAppsV1NamespacedDaemonSet creates a new http.Handler for the list apps v1 namespaced daemon set operation
func NewListAppsV1NamespacedDaemonSet(ctx *middleware.Context, handler ListAppsV1NamespacedDaemonSetHandler) *ListAppsV1NamespacedDaemonSet {
	return &ListAppsV1NamespacedDaemonSet{Context: ctx, Handler: handler}
}

/*ListAppsV1NamespacedDaemonSet swagger:route GET /apis/apps/v1/namespaces/{namespace}/daemonsets apps_v1 listAppsV1NamespacedDaemonSet

list or watch objects of kind DaemonSet

*/
type ListAppsV1NamespacedDaemonSet struct {
	Context *middleware.Context
	Handler ListAppsV1NamespacedDaemonSetHandler
}

func (o *ListAppsV1NamespacedDaemonSet) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewListAppsV1NamespacedDaemonSetParams()

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
