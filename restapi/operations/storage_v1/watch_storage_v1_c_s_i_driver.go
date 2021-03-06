// Code generated by go-swagger; DO NOT EDIT.

package storage_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// WatchStorageV1CSIDriverHandlerFunc turns a function with the right signature into a watch storage v1 c s i driver handler
type WatchStorageV1CSIDriverHandlerFunc func(WatchStorageV1CSIDriverParams) middleware.Responder

// Handle executing the request and returning a response
func (fn WatchStorageV1CSIDriverHandlerFunc) Handle(params WatchStorageV1CSIDriverParams) middleware.Responder {
	return fn(params)
}

// WatchStorageV1CSIDriverHandler interface for that can handle valid watch storage v1 c s i driver params
type WatchStorageV1CSIDriverHandler interface {
	Handle(WatchStorageV1CSIDriverParams) middleware.Responder
}

// NewWatchStorageV1CSIDriver creates a new http.Handler for the watch storage v1 c s i driver operation
func NewWatchStorageV1CSIDriver(ctx *middleware.Context, handler WatchStorageV1CSIDriverHandler) *WatchStorageV1CSIDriver {
	return &WatchStorageV1CSIDriver{Context: ctx, Handler: handler}
}

/*WatchStorageV1CSIDriver swagger:route GET /apis/storage.k8s.io/v1/watch/csidrivers/{name} storage_v1 watchStorageV1CSIDriver

watch changes to an object of kind CSIDriver. deprecated: use the 'watch' parameter with a list operation instead, filtered to a single item with the 'fieldSelector' parameter.

*/
type WatchStorageV1CSIDriver struct {
	Context *middleware.Context
	Handler WatchStorageV1CSIDriverHandler
}

func (o *WatchStorageV1CSIDriver) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewWatchStorageV1CSIDriverParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
