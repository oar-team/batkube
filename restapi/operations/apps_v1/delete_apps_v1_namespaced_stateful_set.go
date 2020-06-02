// Code generated by go-swagger; DO NOT EDIT.

package apps_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteAppsV1NamespacedStatefulSetHandlerFunc turns a function with the right signature into a delete apps v1 namespaced stateful set handler
type DeleteAppsV1NamespacedStatefulSetHandlerFunc func(DeleteAppsV1NamespacedStatefulSetParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteAppsV1NamespacedStatefulSetHandlerFunc) Handle(params DeleteAppsV1NamespacedStatefulSetParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// DeleteAppsV1NamespacedStatefulSetHandler interface for that can handle valid delete apps v1 namespaced stateful set params
type DeleteAppsV1NamespacedStatefulSetHandler interface {
	Handle(DeleteAppsV1NamespacedStatefulSetParams, interface{}) middleware.Responder
}

// NewDeleteAppsV1NamespacedStatefulSet creates a new http.Handler for the delete apps v1 namespaced stateful set operation
func NewDeleteAppsV1NamespacedStatefulSet(ctx *middleware.Context, handler DeleteAppsV1NamespacedStatefulSetHandler) *DeleteAppsV1NamespacedStatefulSet {
	return &DeleteAppsV1NamespacedStatefulSet{Context: ctx, Handler: handler}
}

/*DeleteAppsV1NamespacedStatefulSet swagger:route DELETE /apis/apps/v1/namespaces/{namespace}/statefulsets/{name} apps_v1 deleteAppsV1NamespacedStatefulSet

delete a StatefulSet

*/
type DeleteAppsV1NamespacedStatefulSet struct {
	Context *middleware.Context
	Handler DeleteAppsV1NamespacedStatefulSetHandler
}

func (o *DeleteAppsV1NamespacedStatefulSet) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteAppsV1NamespacedStatefulSetParams()

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
