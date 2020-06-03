// Code generated by go-swagger; DO NOT EDIT.

package node_v1alpha1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// WatchNodeV1alpha1RuntimeClassHandlerFunc turns a function with the right signature into a watch node v1alpha1 runtime class handler
type WatchNodeV1alpha1RuntimeClassHandlerFunc func(WatchNodeV1alpha1RuntimeClassParams) middleware.Responder

// Handle executing the request and returning a response
func (fn WatchNodeV1alpha1RuntimeClassHandlerFunc) Handle(params WatchNodeV1alpha1RuntimeClassParams) middleware.Responder {
	return fn(params)
}

// WatchNodeV1alpha1RuntimeClassHandler interface for that can handle valid watch node v1alpha1 runtime class params
type WatchNodeV1alpha1RuntimeClassHandler interface {
	Handle(WatchNodeV1alpha1RuntimeClassParams) middleware.Responder
}

// NewWatchNodeV1alpha1RuntimeClass creates a new http.Handler for the watch node v1alpha1 runtime class operation
func NewWatchNodeV1alpha1RuntimeClass(ctx *middleware.Context, handler WatchNodeV1alpha1RuntimeClassHandler) *WatchNodeV1alpha1RuntimeClass {
	return &WatchNodeV1alpha1RuntimeClass{Context: ctx, Handler: handler}
}

/*WatchNodeV1alpha1RuntimeClass swagger:route GET /apis/node.k8s.io/v1alpha1/watch/runtimeclasses/{name} node_v1alpha1 watchNodeV1alpha1RuntimeClass

watch changes to an object of kind RuntimeClass. deprecated: use the 'watch' parameter with a list operation instead, filtered to a single item with the 'fieldSelector' parameter.

*/
type WatchNodeV1alpha1RuntimeClass struct {
	Context *middleware.Context
	Handler WatchNodeV1alpha1RuntimeClassHandler
}

func (o *WatchNodeV1alpha1RuntimeClass) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewWatchNodeV1alpha1RuntimeClassParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
