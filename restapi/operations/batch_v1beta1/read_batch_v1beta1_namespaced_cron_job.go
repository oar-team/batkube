// Code generated by go-swagger; DO NOT EDIT.

package batch_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ReadBatchV1beta1NamespacedCronJobHandlerFunc turns a function with the right signature into a read batch v1beta1 namespaced cron job handler
type ReadBatchV1beta1NamespacedCronJobHandlerFunc func(ReadBatchV1beta1NamespacedCronJobParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ReadBatchV1beta1NamespacedCronJobHandlerFunc) Handle(params ReadBatchV1beta1NamespacedCronJobParams) middleware.Responder {
	return fn(params)
}

// ReadBatchV1beta1NamespacedCronJobHandler interface for that can handle valid read batch v1beta1 namespaced cron job params
type ReadBatchV1beta1NamespacedCronJobHandler interface {
	Handle(ReadBatchV1beta1NamespacedCronJobParams) middleware.Responder
}

// NewReadBatchV1beta1NamespacedCronJob creates a new http.Handler for the read batch v1beta1 namespaced cron job operation
func NewReadBatchV1beta1NamespacedCronJob(ctx *middleware.Context, handler ReadBatchV1beta1NamespacedCronJobHandler) *ReadBatchV1beta1NamespacedCronJob {
	return &ReadBatchV1beta1NamespacedCronJob{Context: ctx, Handler: handler}
}

/*ReadBatchV1beta1NamespacedCronJob swagger:route GET /apis/batch/v1beta1/namespaces/{namespace}/cronjobs/{name} batch_v1beta1 readBatchV1beta1NamespacedCronJob

read the specified CronJob

*/
type ReadBatchV1beta1NamespacedCronJob struct {
	Context *middleware.Context
	Handler ReadBatchV1beta1NamespacedCronJobHandler
}

func (o *ReadBatchV1beta1NamespacedCronJob) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewReadBatchV1beta1NamespacedCronJobParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
