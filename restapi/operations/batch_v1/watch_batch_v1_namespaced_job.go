// Code generated by go-swagger; DO NOT EDIT.

package batch_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// WatchBatchV1NamespacedJobHandlerFunc turns a function with the right signature into a watch batch v1 namespaced job handler
type WatchBatchV1NamespacedJobHandlerFunc func(WatchBatchV1NamespacedJobParams) middleware.Responder

// Handle executing the request and returning a response
func (fn WatchBatchV1NamespacedJobHandlerFunc) Handle(params WatchBatchV1NamespacedJobParams) middleware.Responder {
	return fn(params)
}

// WatchBatchV1NamespacedJobHandler interface for that can handle valid watch batch v1 namespaced job params
type WatchBatchV1NamespacedJobHandler interface {
	Handle(WatchBatchV1NamespacedJobParams) middleware.Responder
}

// NewWatchBatchV1NamespacedJob creates a new http.Handler for the watch batch v1 namespaced job operation
func NewWatchBatchV1NamespacedJob(ctx *middleware.Context, handler WatchBatchV1NamespacedJobHandler) *WatchBatchV1NamespacedJob {
	return &WatchBatchV1NamespacedJob{Context: ctx, Handler: handler}
}

/*WatchBatchV1NamespacedJob swagger:route GET /apis/batch/v1/watch/namespaces/{namespace}/jobs/{name} batch_v1 watchBatchV1NamespacedJob

watch changes to an object of kind Job. deprecated: use the 'watch' parameter with a list operation instead, filtered to a single item with the 'fieldSelector' parameter.

*/
type WatchBatchV1NamespacedJob struct {
	Context *middleware.Context
	Handler WatchBatchV1NamespacedJobHandler
}

func (o *WatchBatchV1NamespacedJob) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewWatchBatchV1NamespacedJobParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
