// Code generated by go-swagger; DO NOT EDIT.

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PatchCoreV1NamespacedReplicationControllerHandlerFunc turns a function with the right signature into a patch core v1 namespaced replication controller handler
type PatchCoreV1NamespacedReplicationControllerHandlerFunc func(PatchCoreV1NamespacedReplicationControllerParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PatchCoreV1NamespacedReplicationControllerHandlerFunc) Handle(params PatchCoreV1NamespacedReplicationControllerParams) middleware.Responder {
	return fn(params)
}

// PatchCoreV1NamespacedReplicationControllerHandler interface for that can handle valid patch core v1 namespaced replication controller params
type PatchCoreV1NamespacedReplicationControllerHandler interface {
	Handle(PatchCoreV1NamespacedReplicationControllerParams) middleware.Responder
}

// NewPatchCoreV1NamespacedReplicationController creates a new http.Handler for the patch core v1 namespaced replication controller operation
func NewPatchCoreV1NamespacedReplicationController(ctx *middleware.Context, handler PatchCoreV1NamespacedReplicationControllerHandler) *PatchCoreV1NamespacedReplicationController {
	return &PatchCoreV1NamespacedReplicationController{Context: ctx, Handler: handler}
}

/*PatchCoreV1NamespacedReplicationController swagger:route PATCH /api/v1/namespaces/{namespace}/replicationcontrollers/{name} core_v1 patchCoreV1NamespacedReplicationController

partially update the specified ReplicationController

*/
type PatchCoreV1NamespacedReplicationController struct {
	Context *middleware.Context
	Handler PatchCoreV1NamespacedReplicationControllerHandler
}

func (o *PatchCoreV1NamespacedReplicationController) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPatchCoreV1NamespacedReplicationControllerParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
