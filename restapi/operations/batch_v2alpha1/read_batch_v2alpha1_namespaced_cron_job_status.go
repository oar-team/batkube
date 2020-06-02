// Code generated by go-swagger; DO NOT EDIT.

package batch_v2alpha1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ReadBatchV2alpha1NamespacedCronJobStatusHandlerFunc turns a function with the right signature into a read batch v2alpha1 namespaced cron job status handler
type ReadBatchV2alpha1NamespacedCronJobStatusHandlerFunc func(ReadBatchV2alpha1NamespacedCronJobStatusParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn ReadBatchV2alpha1NamespacedCronJobStatusHandlerFunc) Handle(params ReadBatchV2alpha1NamespacedCronJobStatusParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// ReadBatchV2alpha1NamespacedCronJobStatusHandler interface for that can handle valid read batch v2alpha1 namespaced cron job status params
type ReadBatchV2alpha1NamespacedCronJobStatusHandler interface {
	Handle(ReadBatchV2alpha1NamespacedCronJobStatusParams, interface{}) middleware.Responder
}

// NewReadBatchV2alpha1NamespacedCronJobStatus creates a new http.Handler for the read batch v2alpha1 namespaced cron job status operation
func NewReadBatchV2alpha1NamespacedCronJobStatus(ctx *middleware.Context, handler ReadBatchV2alpha1NamespacedCronJobStatusHandler) *ReadBatchV2alpha1NamespacedCronJobStatus {
	return &ReadBatchV2alpha1NamespacedCronJobStatus{Context: ctx, Handler: handler}
}

/*ReadBatchV2alpha1NamespacedCronJobStatus swagger:route GET /apis/batch/v2alpha1/namespaces/{namespace}/cronjobs/{name}/status batch_v2alpha1 readBatchV2alpha1NamespacedCronJobStatus

read status of the specified CronJob

*/
type ReadBatchV2alpha1NamespacedCronJobStatus struct {
	Context *middleware.Context
	Handler ReadBatchV2alpha1NamespacedCronJobStatusHandler
}

func (o *ReadBatchV2alpha1NamespacedCronJobStatus) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewReadBatchV2alpha1NamespacedCronJobStatusParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
