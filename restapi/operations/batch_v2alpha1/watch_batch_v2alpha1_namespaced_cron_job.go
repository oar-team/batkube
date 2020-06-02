// Code generated by go-swagger; DO NOT EDIT.

package batch_v2alpha1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// WatchBatchV2alpha1NamespacedCronJobHandlerFunc turns a function with the right signature into a watch batch v2alpha1 namespaced cron job handler
type WatchBatchV2alpha1NamespacedCronJobHandlerFunc func(WatchBatchV2alpha1NamespacedCronJobParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn WatchBatchV2alpha1NamespacedCronJobHandlerFunc) Handle(params WatchBatchV2alpha1NamespacedCronJobParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// WatchBatchV2alpha1NamespacedCronJobHandler interface for that can handle valid watch batch v2alpha1 namespaced cron job params
type WatchBatchV2alpha1NamespacedCronJobHandler interface {
	Handle(WatchBatchV2alpha1NamespacedCronJobParams, interface{}) middleware.Responder
}

// NewWatchBatchV2alpha1NamespacedCronJob creates a new http.Handler for the watch batch v2alpha1 namespaced cron job operation
func NewWatchBatchV2alpha1NamespacedCronJob(ctx *middleware.Context, handler WatchBatchV2alpha1NamespacedCronJobHandler) *WatchBatchV2alpha1NamespacedCronJob {
	return &WatchBatchV2alpha1NamespacedCronJob{Context: ctx, Handler: handler}
}

/*WatchBatchV2alpha1NamespacedCronJob swagger:route GET /apis/batch/v2alpha1/watch/namespaces/{namespace}/cronjobs/{name} batch_v2alpha1 watchBatchV2alpha1NamespacedCronJob

watch changes to an object of kind CronJob. deprecated: use the 'watch' parameter with a list operation instead, filtered to a single item with the 'fieldSelector' parameter.

*/
type WatchBatchV2alpha1NamespacedCronJob struct {
	Context *middleware.Context
	Handler WatchBatchV2alpha1NamespacedCronJobHandler
}

func (o *WatchBatchV2alpha1NamespacedCronJob) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewWatchBatchV2alpha1NamespacedCronJobParams()

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
