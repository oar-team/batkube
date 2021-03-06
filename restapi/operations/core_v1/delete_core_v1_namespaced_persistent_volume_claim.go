// Code generated by go-swagger; DO NOT EDIT.

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteCoreV1NamespacedPersistentVolumeClaimHandlerFunc turns a function with the right signature into a delete core v1 namespaced persistent volume claim handler
type DeleteCoreV1NamespacedPersistentVolumeClaimHandlerFunc func(DeleteCoreV1NamespacedPersistentVolumeClaimParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteCoreV1NamespacedPersistentVolumeClaimHandlerFunc) Handle(params DeleteCoreV1NamespacedPersistentVolumeClaimParams) middleware.Responder {
	return fn(params)
}

// DeleteCoreV1NamespacedPersistentVolumeClaimHandler interface for that can handle valid delete core v1 namespaced persistent volume claim params
type DeleteCoreV1NamespacedPersistentVolumeClaimHandler interface {
	Handle(DeleteCoreV1NamespacedPersistentVolumeClaimParams) middleware.Responder
}

// NewDeleteCoreV1NamespacedPersistentVolumeClaim creates a new http.Handler for the delete core v1 namespaced persistent volume claim operation
func NewDeleteCoreV1NamespacedPersistentVolumeClaim(ctx *middleware.Context, handler DeleteCoreV1NamespacedPersistentVolumeClaimHandler) *DeleteCoreV1NamespacedPersistentVolumeClaim {
	return &DeleteCoreV1NamespacedPersistentVolumeClaim{Context: ctx, Handler: handler}
}

/*DeleteCoreV1NamespacedPersistentVolumeClaim swagger:route DELETE /api/v1/namespaces/{namespace}/persistentvolumeclaims/{name} core_v1 deleteCoreV1NamespacedPersistentVolumeClaim

delete a PersistentVolumeClaim

*/
type DeleteCoreV1NamespacedPersistentVolumeClaim struct {
	Context *middleware.Context
	Handler DeleteCoreV1NamespacedPersistentVolumeClaimHandler
}

func (o *DeleteCoreV1NamespacedPersistentVolumeClaim) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteCoreV1NamespacedPersistentVolumeClaimParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
