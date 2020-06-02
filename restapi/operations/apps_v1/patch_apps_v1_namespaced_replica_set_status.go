// Code generated by go-swagger; DO NOT EDIT.

package apps_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PatchAppsV1NamespacedReplicaSetStatusHandlerFunc turns a function with the right signature into a patch apps v1 namespaced replica set status handler
type PatchAppsV1NamespacedReplicaSetStatusHandlerFunc func(PatchAppsV1NamespacedReplicaSetStatusParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn PatchAppsV1NamespacedReplicaSetStatusHandlerFunc) Handle(params PatchAppsV1NamespacedReplicaSetStatusParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// PatchAppsV1NamespacedReplicaSetStatusHandler interface for that can handle valid patch apps v1 namespaced replica set status params
type PatchAppsV1NamespacedReplicaSetStatusHandler interface {
	Handle(PatchAppsV1NamespacedReplicaSetStatusParams, interface{}) middleware.Responder
}

// NewPatchAppsV1NamespacedReplicaSetStatus creates a new http.Handler for the patch apps v1 namespaced replica set status operation
func NewPatchAppsV1NamespacedReplicaSetStatus(ctx *middleware.Context, handler PatchAppsV1NamespacedReplicaSetStatusHandler) *PatchAppsV1NamespacedReplicaSetStatus {
	return &PatchAppsV1NamespacedReplicaSetStatus{Context: ctx, Handler: handler}
}

/*PatchAppsV1NamespacedReplicaSetStatus swagger:route PATCH /apis/apps/v1/namespaces/{namespace}/replicasets/{name}/status apps_v1 patchAppsV1NamespacedReplicaSetStatus

partially update status of the specified ReplicaSet

*/
type PatchAppsV1NamespacedReplicaSetStatus struct {
	Context *middleware.Context
	Handler PatchAppsV1NamespacedReplicaSetStatusHandler
}

func (o *PatchAppsV1NamespacedReplicaSetStatus) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPatchAppsV1NamespacedReplicaSetStatusParams()

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
