// Code generated by go-swagger; DO NOT EDIT.

package batch_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ReadBatchV1NamespacedJobHandlerFunc turns a function with the right signature into a read batch v1 namespaced job handler
type ReadBatchV1NamespacedJobHandlerFunc func(ReadBatchV1NamespacedJobParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ReadBatchV1NamespacedJobHandlerFunc) Handle(params ReadBatchV1NamespacedJobParams) middleware.Responder {
	return fn(params)
}

// ReadBatchV1NamespacedJobHandler interface for that can handle valid read batch v1 namespaced job params
type ReadBatchV1NamespacedJobHandler interface {
	Handle(ReadBatchV1NamespacedJobParams) middleware.Responder
}

// NewReadBatchV1NamespacedJob creates a new http.Handler for the read batch v1 namespaced job operation
func NewReadBatchV1NamespacedJob(ctx *middleware.Context, handler ReadBatchV1NamespacedJobHandler) *ReadBatchV1NamespacedJob {
	return &ReadBatchV1NamespacedJob{Context: ctx, Handler: handler}
}

/*ReadBatchV1NamespacedJob swagger:route GET /apis/batch/v1/namespaces/{namespace}/jobs/{name} batch_v1 readBatchV1NamespacedJob

read the specified Job

*/
type ReadBatchV1NamespacedJob struct {
	Context *middleware.Context
	Handler ReadBatchV1NamespacedJobHandler
}

func (o *ReadBatchV1NamespacedJob) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewReadBatchV1NamespacedJobParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}