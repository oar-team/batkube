// Code generated by go-swagger; DO NOT EDIT.

package apps_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PatchAppsV1NamespacedStatefulSetStatusHandlerFunc turns a function with the right signature into a patch apps v1 namespaced stateful set status handler
type PatchAppsV1NamespacedStatefulSetStatusHandlerFunc func(PatchAppsV1NamespacedStatefulSetStatusParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PatchAppsV1NamespacedStatefulSetStatusHandlerFunc) Handle(params PatchAppsV1NamespacedStatefulSetStatusParams) middleware.Responder {
	return fn(params)
}

// PatchAppsV1NamespacedStatefulSetStatusHandler interface for that can handle valid patch apps v1 namespaced stateful set status params
type PatchAppsV1NamespacedStatefulSetStatusHandler interface {
	Handle(PatchAppsV1NamespacedStatefulSetStatusParams) middleware.Responder
}

// NewPatchAppsV1NamespacedStatefulSetStatus creates a new http.Handler for the patch apps v1 namespaced stateful set status operation
func NewPatchAppsV1NamespacedStatefulSetStatus(ctx *middleware.Context, handler PatchAppsV1NamespacedStatefulSetStatusHandler) *PatchAppsV1NamespacedStatefulSetStatus {
	return &PatchAppsV1NamespacedStatefulSetStatus{Context: ctx, Handler: handler}
}

/*PatchAppsV1NamespacedStatefulSetStatus swagger:route PATCH /apis/apps/v1/namespaces/{namespace}/statefulsets/{name}/status apps_v1 patchAppsV1NamespacedStatefulSetStatus

partially update status of the specified StatefulSet

*/
type PatchAppsV1NamespacedStatefulSetStatus struct {
	Context *middleware.Context
	Handler PatchAppsV1NamespacedStatefulSetStatusHandler
}

func (o *PatchAppsV1NamespacedStatefulSetStatus) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPatchAppsV1NamespacedStatefulSetStatusParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
