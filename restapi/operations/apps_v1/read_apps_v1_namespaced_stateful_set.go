// Code generated by go-swagger; DO NOT EDIT.

package apps_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ReadAppsV1NamespacedStatefulSetHandlerFunc turns a function with the right signature into a read apps v1 namespaced stateful set handler
type ReadAppsV1NamespacedStatefulSetHandlerFunc func(ReadAppsV1NamespacedStatefulSetParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn ReadAppsV1NamespacedStatefulSetHandlerFunc) Handle(params ReadAppsV1NamespacedStatefulSetParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// ReadAppsV1NamespacedStatefulSetHandler interface for that can handle valid read apps v1 namespaced stateful set params
type ReadAppsV1NamespacedStatefulSetHandler interface {
	Handle(ReadAppsV1NamespacedStatefulSetParams, interface{}) middleware.Responder
}

// NewReadAppsV1NamespacedStatefulSet creates a new http.Handler for the read apps v1 namespaced stateful set operation
func NewReadAppsV1NamespacedStatefulSet(ctx *middleware.Context, handler ReadAppsV1NamespacedStatefulSetHandler) *ReadAppsV1NamespacedStatefulSet {
	return &ReadAppsV1NamespacedStatefulSet{Context: ctx, Handler: handler}
}

/*ReadAppsV1NamespacedStatefulSet swagger:route GET /apis/apps/v1/namespaces/{namespace}/statefulsets/{name} apps_v1 readAppsV1NamespacedStatefulSet

read the specified StatefulSet

*/
type ReadAppsV1NamespacedStatefulSet struct {
	Context *middleware.Context
	Handler ReadAppsV1NamespacedStatefulSetHandler
}

func (o *ReadAppsV1NamespacedStatefulSet) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewReadAppsV1NamespacedStatefulSetParams()

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
