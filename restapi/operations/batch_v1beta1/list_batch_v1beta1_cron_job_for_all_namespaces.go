// Code generated by go-swagger; DO NOT EDIT.

package batch_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ListBatchV1beta1CronJobForAllNamespacesHandlerFunc turns a function with the right signature into a list batch v1beta1 cron job for all namespaces handler
type ListBatchV1beta1CronJobForAllNamespacesHandlerFunc func(ListBatchV1beta1CronJobForAllNamespacesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ListBatchV1beta1CronJobForAllNamespacesHandlerFunc) Handle(params ListBatchV1beta1CronJobForAllNamespacesParams) middleware.Responder {
	return fn(params)
}

// ListBatchV1beta1CronJobForAllNamespacesHandler interface for that can handle valid list batch v1beta1 cron job for all namespaces params
type ListBatchV1beta1CronJobForAllNamespacesHandler interface {
	Handle(ListBatchV1beta1CronJobForAllNamespacesParams) middleware.Responder
}

// NewListBatchV1beta1CronJobForAllNamespaces creates a new http.Handler for the list batch v1beta1 cron job for all namespaces operation
func NewListBatchV1beta1CronJobForAllNamespaces(ctx *middleware.Context, handler ListBatchV1beta1CronJobForAllNamespacesHandler) *ListBatchV1beta1CronJobForAllNamespaces {
	return &ListBatchV1beta1CronJobForAllNamespaces{Context: ctx, Handler: handler}
}

/*ListBatchV1beta1CronJobForAllNamespaces swagger:route GET /apis/batch/v1beta1/cronjobs batch_v1beta1 listBatchV1beta1CronJobForAllNamespaces

list or watch objects of kind CronJob

*/
type ListBatchV1beta1CronJobForAllNamespaces struct {
	Context *middleware.Context
	Handler ListBatchV1beta1CronJobForAllNamespacesHandler
}

func (o *ListBatchV1beta1CronJobForAllNamespaces) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewListBatchV1beta1CronJobForAllNamespacesParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
