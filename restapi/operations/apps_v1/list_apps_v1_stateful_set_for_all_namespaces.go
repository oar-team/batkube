// Code generated by go-swagger; DO NOT EDIT.

package apps_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ListAppsV1StatefulSetForAllNamespacesHandlerFunc turns a function with the right signature into a list apps v1 stateful set for all namespaces handler
type ListAppsV1StatefulSetForAllNamespacesHandlerFunc func(ListAppsV1StatefulSetForAllNamespacesParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn ListAppsV1StatefulSetForAllNamespacesHandlerFunc) Handle(params ListAppsV1StatefulSetForAllNamespacesParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// ListAppsV1StatefulSetForAllNamespacesHandler interface for that can handle valid list apps v1 stateful set for all namespaces params
type ListAppsV1StatefulSetForAllNamespacesHandler interface {
	Handle(ListAppsV1StatefulSetForAllNamespacesParams, interface{}) middleware.Responder
}

// NewListAppsV1StatefulSetForAllNamespaces creates a new http.Handler for the list apps v1 stateful set for all namespaces operation
func NewListAppsV1StatefulSetForAllNamespaces(ctx *middleware.Context, handler ListAppsV1StatefulSetForAllNamespacesHandler) *ListAppsV1StatefulSetForAllNamespaces {
	return &ListAppsV1StatefulSetForAllNamespaces{Context: ctx, Handler: handler}
}

/*ListAppsV1StatefulSetForAllNamespaces swagger:route GET /apis/apps/v1/statefulsets apps_v1 listAppsV1StatefulSetForAllNamespaces

list or watch objects of kind StatefulSet

*/
type ListAppsV1StatefulSetForAllNamespaces struct {
	Context *middleware.Context
	Handler ListAppsV1StatefulSetForAllNamespacesHandler
}

func (o *ListAppsV1StatefulSetForAllNamespaces) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewListAppsV1StatefulSetForAllNamespacesParams()

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
