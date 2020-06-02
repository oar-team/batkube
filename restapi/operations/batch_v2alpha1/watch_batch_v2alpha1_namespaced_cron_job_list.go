// Code generated by go-swagger; DO NOT EDIT.

package batch_v2alpha1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// WatchBatchV2alpha1NamespacedCronJobListHandlerFunc turns a function with the right signature into a watch batch v2alpha1 namespaced cron job list handler
type WatchBatchV2alpha1NamespacedCronJobListHandlerFunc func(WatchBatchV2alpha1NamespacedCronJobListParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn WatchBatchV2alpha1NamespacedCronJobListHandlerFunc) Handle(params WatchBatchV2alpha1NamespacedCronJobListParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// WatchBatchV2alpha1NamespacedCronJobListHandler interface for that can handle valid watch batch v2alpha1 namespaced cron job list params
type WatchBatchV2alpha1NamespacedCronJobListHandler interface {
	Handle(WatchBatchV2alpha1NamespacedCronJobListParams, interface{}) middleware.Responder
}

// NewWatchBatchV2alpha1NamespacedCronJobList creates a new http.Handler for the watch batch v2alpha1 namespaced cron job list operation
func NewWatchBatchV2alpha1NamespacedCronJobList(ctx *middleware.Context, handler WatchBatchV2alpha1NamespacedCronJobListHandler) *WatchBatchV2alpha1NamespacedCronJobList {
	return &WatchBatchV2alpha1NamespacedCronJobList{Context: ctx, Handler: handler}
}

/*WatchBatchV2alpha1NamespacedCronJobList swagger:route GET /apis/batch/v2alpha1/watch/namespaces/{namespace}/cronjobs batch_v2alpha1 watchBatchV2alpha1NamespacedCronJobList

watch individual changes to a list of CronJob. deprecated: use the 'watch' parameter with a list operation instead.

*/
type WatchBatchV2alpha1NamespacedCronJobList struct {
	Context *middleware.Context
	Handler WatchBatchV2alpha1NamespacedCronJobListHandler
}

func (o *WatchBatchV2alpha1NamespacedCronJobList) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewWatchBatchV2alpha1NamespacedCronJobListParams()

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
