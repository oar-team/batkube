// Code generated by go-swagger; DO NOT EDIT.

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// WatchCoreV1NamespacedPersistentVolumeClaimHandlerFunc turns a function with the right signature into a watch core v1 namespaced persistent volume claim handler
type WatchCoreV1NamespacedPersistentVolumeClaimHandlerFunc func(WatchCoreV1NamespacedPersistentVolumeClaimParams) middleware.Responder

// Handle executing the request and returning a response
func (fn WatchCoreV1NamespacedPersistentVolumeClaimHandlerFunc) Handle(params WatchCoreV1NamespacedPersistentVolumeClaimParams) middleware.Responder {
	return fn(params)
}

// WatchCoreV1NamespacedPersistentVolumeClaimHandler interface for that can handle valid watch core v1 namespaced persistent volume claim params
type WatchCoreV1NamespacedPersistentVolumeClaimHandler interface {
	Handle(WatchCoreV1NamespacedPersistentVolumeClaimParams) middleware.Responder
}

// NewWatchCoreV1NamespacedPersistentVolumeClaim creates a new http.Handler for the watch core v1 namespaced persistent volume claim operation
func NewWatchCoreV1NamespacedPersistentVolumeClaim(ctx *middleware.Context, handler WatchCoreV1NamespacedPersistentVolumeClaimHandler) *WatchCoreV1NamespacedPersistentVolumeClaim {
	return &WatchCoreV1NamespacedPersistentVolumeClaim{Context: ctx, Handler: handler}
}

/*WatchCoreV1NamespacedPersistentVolumeClaim swagger:route GET /api/v1/watch/namespaces/{namespace}/persistentvolumeclaims/{name} core_v1 watchCoreV1NamespacedPersistentVolumeClaim

watch changes to an object of kind PersistentVolumeClaim. deprecated: use the 'watch' parameter with a list operation instead, filtered to a single item with the 'fieldSelector' parameter.

*/
type WatchCoreV1NamespacedPersistentVolumeClaim struct {
	Context *middleware.Context
	Handler WatchCoreV1NamespacedPersistentVolumeClaimHandler
}

func (o *WatchCoreV1NamespacedPersistentVolumeClaim) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewWatchCoreV1NamespacedPersistentVolumeClaimParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
