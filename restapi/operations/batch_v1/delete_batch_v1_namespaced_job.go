// Code generated by go-swagger; DO NOT EDIT.

package batch_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteBatchV1NamespacedJobHandlerFunc turns a function with the right signature into a delete batch v1 namespaced job handler
type DeleteBatchV1NamespacedJobHandlerFunc func(DeleteBatchV1NamespacedJobParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteBatchV1NamespacedJobHandlerFunc) Handle(params DeleteBatchV1NamespacedJobParams) middleware.Responder {
	return fn(params)
}

// DeleteBatchV1NamespacedJobHandler interface for that can handle valid delete batch v1 namespaced job params
type DeleteBatchV1NamespacedJobHandler interface {
	Handle(DeleteBatchV1NamespacedJobParams) middleware.Responder
}

// NewDeleteBatchV1NamespacedJob creates a new http.Handler for the delete batch v1 namespaced job operation
func NewDeleteBatchV1NamespacedJob(ctx *middleware.Context, handler DeleteBatchV1NamespacedJobHandler) *DeleteBatchV1NamespacedJob {
	return &DeleteBatchV1NamespacedJob{Context: ctx, Handler: handler}
}

/*DeleteBatchV1NamespacedJob swagger:route DELETE /apis/batch/v1/namespaces/{namespace}/jobs/{name} batch_v1 deleteBatchV1NamespacedJob

delete a Job

*/
type DeleteBatchV1NamespacedJob struct {
	Context *middleware.Context
	Handler DeleteBatchV1NamespacedJobHandler
}

func (o *DeleteBatchV1NamespacedJob) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteBatchV1NamespacedJobParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
