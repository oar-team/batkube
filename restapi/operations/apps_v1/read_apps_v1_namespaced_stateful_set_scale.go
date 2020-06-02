// Code generated by go-swagger; DO NOT EDIT.

package apps_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ReadAppsV1NamespacedStatefulSetScaleHandlerFunc turns a function with the right signature into a read apps v1 namespaced stateful set scale handler
type ReadAppsV1NamespacedStatefulSetScaleHandlerFunc func(ReadAppsV1NamespacedStatefulSetScaleParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn ReadAppsV1NamespacedStatefulSetScaleHandlerFunc) Handle(params ReadAppsV1NamespacedStatefulSetScaleParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// ReadAppsV1NamespacedStatefulSetScaleHandler interface for that can handle valid read apps v1 namespaced stateful set scale params
type ReadAppsV1NamespacedStatefulSetScaleHandler interface {
	Handle(ReadAppsV1NamespacedStatefulSetScaleParams, interface{}) middleware.Responder
}

// NewReadAppsV1NamespacedStatefulSetScale creates a new http.Handler for the read apps v1 namespaced stateful set scale operation
func NewReadAppsV1NamespacedStatefulSetScale(ctx *middleware.Context, handler ReadAppsV1NamespacedStatefulSetScaleHandler) *ReadAppsV1NamespacedStatefulSetScale {
	return &ReadAppsV1NamespacedStatefulSetScale{Context: ctx, Handler: handler}
}

/*ReadAppsV1NamespacedStatefulSetScale swagger:route GET /apis/apps/v1/namespaces/{namespace}/statefulsets/{name}/scale apps_v1 readAppsV1NamespacedStatefulSetScale

read scale of the specified StatefulSet

*/
type ReadAppsV1NamespacedStatefulSetScale struct {
	Context *middleware.Context
	Handler ReadAppsV1NamespacedStatefulSetScaleHandler
}

func (o *ReadAppsV1NamespacedStatefulSetScale) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewReadAppsV1NamespacedStatefulSetScaleParams()

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
