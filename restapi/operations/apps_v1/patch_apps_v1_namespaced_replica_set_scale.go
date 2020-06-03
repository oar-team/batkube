// Code generated by go-swagger; DO NOT EDIT.

package apps_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PatchAppsV1NamespacedReplicaSetScaleHandlerFunc turns a function with the right signature into a patch apps v1 namespaced replica set scale handler
type PatchAppsV1NamespacedReplicaSetScaleHandlerFunc func(PatchAppsV1NamespacedReplicaSetScaleParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PatchAppsV1NamespacedReplicaSetScaleHandlerFunc) Handle(params PatchAppsV1NamespacedReplicaSetScaleParams) middleware.Responder {
	return fn(params)
}

// PatchAppsV1NamespacedReplicaSetScaleHandler interface for that can handle valid patch apps v1 namespaced replica set scale params
type PatchAppsV1NamespacedReplicaSetScaleHandler interface {
	Handle(PatchAppsV1NamespacedReplicaSetScaleParams) middleware.Responder
}

// NewPatchAppsV1NamespacedReplicaSetScale creates a new http.Handler for the patch apps v1 namespaced replica set scale operation
func NewPatchAppsV1NamespacedReplicaSetScale(ctx *middleware.Context, handler PatchAppsV1NamespacedReplicaSetScaleHandler) *PatchAppsV1NamespacedReplicaSetScale {
	return &PatchAppsV1NamespacedReplicaSetScale{Context: ctx, Handler: handler}
}

/*PatchAppsV1NamespacedReplicaSetScale swagger:route PATCH /apis/apps/v1/namespaces/{namespace}/replicasets/{name}/scale apps_v1 patchAppsV1NamespacedReplicaSetScale

partially update scale of the specified ReplicaSet

*/
type PatchAppsV1NamespacedReplicaSetScale struct {
	Context *middleware.Context
	Handler PatchAppsV1NamespacedReplicaSetScaleHandler
}

func (o *PatchAppsV1NamespacedReplicaSetScale) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPatchAppsV1NamespacedReplicaSetScaleParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
