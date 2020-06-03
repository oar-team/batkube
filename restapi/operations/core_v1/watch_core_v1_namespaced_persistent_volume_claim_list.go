// Code generated by go-swagger; DO NOT EDIT.

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// WatchCoreV1NamespacedPersistentVolumeClaimListHandlerFunc turns a function with the right signature into a watch core v1 namespaced persistent volume claim list handler
type WatchCoreV1NamespacedPersistentVolumeClaimListHandlerFunc func(WatchCoreV1NamespacedPersistentVolumeClaimListParams) middleware.Responder

// Handle executing the request and returning a response
func (fn WatchCoreV1NamespacedPersistentVolumeClaimListHandlerFunc) Handle(params WatchCoreV1NamespacedPersistentVolumeClaimListParams) middleware.Responder {
	return fn(params)
}

// WatchCoreV1NamespacedPersistentVolumeClaimListHandler interface for that can handle valid watch core v1 namespaced persistent volume claim list params
type WatchCoreV1NamespacedPersistentVolumeClaimListHandler interface {
	Handle(WatchCoreV1NamespacedPersistentVolumeClaimListParams) middleware.Responder
}

// NewWatchCoreV1NamespacedPersistentVolumeClaimList creates a new http.Handler for the watch core v1 namespaced persistent volume claim list operation
func NewWatchCoreV1NamespacedPersistentVolumeClaimList(ctx *middleware.Context, handler WatchCoreV1NamespacedPersistentVolumeClaimListHandler) *WatchCoreV1NamespacedPersistentVolumeClaimList {
	return &WatchCoreV1NamespacedPersistentVolumeClaimList{Context: ctx, Handler: handler}
}

/*WatchCoreV1NamespacedPersistentVolumeClaimList swagger:route GET /api/v1/watch/namespaces/{namespace}/persistentvolumeclaims core_v1 watchCoreV1NamespacedPersistentVolumeClaimList

watch individual changes to a list of PersistentVolumeClaim. deprecated: use the 'watch' parameter with a list operation instead.

*/
type WatchCoreV1NamespacedPersistentVolumeClaimList struct {
	Context *middleware.Context
	Handler WatchCoreV1NamespacedPersistentVolumeClaimListHandler
}

func (o *WatchCoreV1NamespacedPersistentVolumeClaimList) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewWatchCoreV1NamespacedPersistentVolumeClaimListParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
