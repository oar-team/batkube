// Code generated by go-swagger; DO NOT EDIT.

package batch_v2alpha1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ListBatchV2alpha1CronJobForAllNamespacesHandlerFunc turns a function with the right signature into a list batch v2alpha1 cron job for all namespaces handler
type ListBatchV2alpha1CronJobForAllNamespacesHandlerFunc func(ListBatchV2alpha1CronJobForAllNamespacesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ListBatchV2alpha1CronJobForAllNamespacesHandlerFunc) Handle(params ListBatchV2alpha1CronJobForAllNamespacesParams) middleware.Responder {
	return fn(params)
}

// ListBatchV2alpha1CronJobForAllNamespacesHandler interface for that can handle valid list batch v2alpha1 cron job for all namespaces params
type ListBatchV2alpha1CronJobForAllNamespacesHandler interface {
	Handle(ListBatchV2alpha1CronJobForAllNamespacesParams) middleware.Responder
}

// NewListBatchV2alpha1CronJobForAllNamespaces creates a new http.Handler for the list batch v2alpha1 cron job for all namespaces operation
func NewListBatchV2alpha1CronJobForAllNamespaces(ctx *middleware.Context, handler ListBatchV2alpha1CronJobForAllNamespacesHandler) *ListBatchV2alpha1CronJobForAllNamespaces {
	return &ListBatchV2alpha1CronJobForAllNamespaces{Context: ctx, Handler: handler}
}

/*ListBatchV2alpha1CronJobForAllNamespaces swagger:route GET /apis/batch/v2alpha1/cronjobs batch_v2alpha1 listBatchV2alpha1CronJobForAllNamespaces

list or watch objects of kind CronJob

*/
type ListBatchV2alpha1CronJobForAllNamespaces struct {
	Context *middleware.Context
	Handler ListBatchV2alpha1CronJobForAllNamespacesHandler
}

func (o *ListBatchV2alpha1CronJobForAllNamespaces) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewListBatchV2alpha1CronJobForAllNamespacesParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
