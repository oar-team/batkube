// Code generated by go-swagger; DO NOT EDIT.

package apps_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PatchAppsV1NamespacedReplicaSetHandlerFunc turns a function with the right signature into a patch apps v1 namespaced replica set handler
type PatchAppsV1NamespacedReplicaSetHandlerFunc func(PatchAppsV1NamespacedReplicaSetParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn PatchAppsV1NamespacedReplicaSetHandlerFunc) Handle(params PatchAppsV1NamespacedReplicaSetParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// PatchAppsV1NamespacedReplicaSetHandler interface for that can handle valid patch apps v1 namespaced replica set params
type PatchAppsV1NamespacedReplicaSetHandler interface {
	Handle(PatchAppsV1NamespacedReplicaSetParams, interface{}) middleware.Responder
}

// NewPatchAppsV1NamespacedReplicaSet creates a new http.Handler for the patch apps v1 namespaced replica set operation
func NewPatchAppsV1NamespacedReplicaSet(ctx *middleware.Context, handler PatchAppsV1NamespacedReplicaSetHandler) *PatchAppsV1NamespacedReplicaSet {
	return &PatchAppsV1NamespacedReplicaSet{Context: ctx, Handler: handler}
}

/*PatchAppsV1NamespacedReplicaSet swagger:route PATCH /apis/apps/v1/namespaces/{namespace}/replicasets/{name} apps_v1 patchAppsV1NamespacedReplicaSet

partially update the specified ReplicaSet

*/
type PatchAppsV1NamespacedReplicaSet struct {
	Context *middleware.Context
	Handler PatchAppsV1NamespacedReplicaSetHandler
}

func (o *PatchAppsV1NamespacedReplicaSet) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPatchAppsV1NamespacedReplicaSetParams()

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
