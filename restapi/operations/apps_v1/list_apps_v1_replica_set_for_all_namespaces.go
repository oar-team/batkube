// Code generated by go-swagger; DO NOT EDIT.

package apps_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ListAppsV1ReplicaSetForAllNamespacesHandlerFunc turns a function with the right signature into a list apps v1 replica set for all namespaces handler
type ListAppsV1ReplicaSetForAllNamespacesHandlerFunc func(ListAppsV1ReplicaSetForAllNamespacesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ListAppsV1ReplicaSetForAllNamespacesHandlerFunc) Handle(params ListAppsV1ReplicaSetForAllNamespacesParams) middleware.Responder {
	return fn(params)
}

// ListAppsV1ReplicaSetForAllNamespacesHandler interface for that can handle valid list apps v1 replica set for all namespaces params
type ListAppsV1ReplicaSetForAllNamespacesHandler interface {
	Handle(ListAppsV1ReplicaSetForAllNamespacesParams) middleware.Responder
}

// NewListAppsV1ReplicaSetForAllNamespaces creates a new http.Handler for the list apps v1 replica set for all namespaces operation
func NewListAppsV1ReplicaSetForAllNamespaces(ctx *middleware.Context, handler ListAppsV1ReplicaSetForAllNamespacesHandler) *ListAppsV1ReplicaSetForAllNamespaces {
	return &ListAppsV1ReplicaSetForAllNamespaces{Context: ctx, Handler: handler}
}

/*ListAppsV1ReplicaSetForAllNamespaces swagger:route GET /apis/apps/v1/replicasets apps_v1 listAppsV1ReplicaSetForAllNamespaces

list or watch objects of kind ReplicaSet

*/
type ListAppsV1ReplicaSetForAllNamespaces struct {
	Context *middleware.Context
	Handler ListAppsV1ReplicaSetForAllNamespacesHandler
}

func (o *ListAppsV1ReplicaSetForAllNamespaces) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewListAppsV1ReplicaSetForAllNamespacesParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
