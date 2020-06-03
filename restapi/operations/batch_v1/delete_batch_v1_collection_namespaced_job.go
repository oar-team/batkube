// Code generated by go-swagger; DO NOT EDIT.

package batch_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteBatchV1CollectionNamespacedJobHandlerFunc turns a function with the right signature into a delete batch v1 collection namespaced job handler
type DeleteBatchV1CollectionNamespacedJobHandlerFunc func(DeleteBatchV1CollectionNamespacedJobParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteBatchV1CollectionNamespacedJobHandlerFunc) Handle(params DeleteBatchV1CollectionNamespacedJobParams) middleware.Responder {
	return fn(params)
}

// DeleteBatchV1CollectionNamespacedJobHandler interface for that can handle valid delete batch v1 collection namespaced job params
type DeleteBatchV1CollectionNamespacedJobHandler interface {
	Handle(DeleteBatchV1CollectionNamespacedJobParams) middleware.Responder
}

// NewDeleteBatchV1CollectionNamespacedJob creates a new http.Handler for the delete batch v1 collection namespaced job operation
func NewDeleteBatchV1CollectionNamespacedJob(ctx *middleware.Context, handler DeleteBatchV1CollectionNamespacedJobHandler) *DeleteBatchV1CollectionNamespacedJob {
	return &DeleteBatchV1CollectionNamespacedJob{Context: ctx, Handler: handler}
}

/*DeleteBatchV1CollectionNamespacedJob swagger:route DELETE /apis/batch/v1/namespaces/{namespace}/jobs batch_v1 deleteBatchV1CollectionNamespacedJob

delete collection of Job

*/
type DeleteBatchV1CollectionNamespacedJob struct {
	Context *middleware.Context
	Handler DeleteBatchV1CollectionNamespacedJobHandler
}

func (o *DeleteBatchV1CollectionNamespacedJob) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteBatchV1CollectionNamespacedJobParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
