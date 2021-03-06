// Code generated by go-swagger; DO NOT EDIT.

package storage_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// WatchStorageV1CSINodeListHandlerFunc turns a function with the right signature into a watch storage v1 c s i node list handler
type WatchStorageV1CSINodeListHandlerFunc func(WatchStorageV1CSINodeListParams) middleware.Responder

// Handle executing the request and returning a response
func (fn WatchStorageV1CSINodeListHandlerFunc) Handle(params WatchStorageV1CSINodeListParams) middleware.Responder {
	return fn(params)
}

// WatchStorageV1CSINodeListHandler interface for that can handle valid watch storage v1 c s i node list params
type WatchStorageV1CSINodeListHandler interface {
	Handle(WatchStorageV1CSINodeListParams) middleware.Responder
}

// NewWatchStorageV1CSINodeList creates a new http.Handler for the watch storage v1 c s i node list operation
func NewWatchStorageV1CSINodeList(ctx *middleware.Context, handler WatchStorageV1CSINodeListHandler) *WatchStorageV1CSINodeList {
	return &WatchStorageV1CSINodeList{Context: ctx, Handler: handler}
}

/*WatchStorageV1CSINodeList swagger:route GET /apis/storage.k8s.io/v1/watch/csinodes storage_v1 watchStorageV1CSINodeList

watch individual changes to a list of CSINode. deprecated: use the 'watch' parameter with a list operation instead.

*/
type WatchStorageV1CSINodeList struct {
	Context *middleware.Context
	Handler WatchStorageV1CSINodeListHandler
}

func (o *WatchStorageV1CSINodeList) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewWatchStorageV1CSINodeListParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
