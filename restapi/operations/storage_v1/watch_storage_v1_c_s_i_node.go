// Code generated by go-swagger; DO NOT EDIT.

package storage_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// WatchStorageV1CSINodeHandlerFunc turns a function with the right signature into a watch storage v1 c s i node handler
type WatchStorageV1CSINodeHandlerFunc func(WatchStorageV1CSINodeParams) middleware.Responder

// Handle executing the request and returning a response
func (fn WatchStorageV1CSINodeHandlerFunc) Handle(params WatchStorageV1CSINodeParams) middleware.Responder {
	return fn(params)
}

// WatchStorageV1CSINodeHandler interface for that can handle valid watch storage v1 c s i node params
type WatchStorageV1CSINodeHandler interface {
	Handle(WatchStorageV1CSINodeParams) middleware.Responder
}

// NewWatchStorageV1CSINode creates a new http.Handler for the watch storage v1 c s i node operation
func NewWatchStorageV1CSINode(ctx *middleware.Context, handler WatchStorageV1CSINodeHandler) *WatchStorageV1CSINode {
	return &WatchStorageV1CSINode{Context: ctx, Handler: handler}
}

/*WatchStorageV1CSINode swagger:route GET /apis/storage.k8s.io/v1/watch/csinodes/{name} storage_v1 watchStorageV1CSINode

watch changes to an object of kind CSINode. deprecated: use the 'watch' parameter with a list operation instead, filtered to a single item with the 'fieldSelector' parameter.

*/
type WatchStorageV1CSINode struct {
	Context *middleware.Context
	Handler WatchStorageV1CSINodeHandler
}

func (o *WatchStorageV1CSINode) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewWatchStorageV1CSINodeParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
