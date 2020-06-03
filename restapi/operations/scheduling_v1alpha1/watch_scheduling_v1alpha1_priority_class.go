// Code generated by go-swagger; DO NOT EDIT.

package scheduling_v1alpha1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// WatchSchedulingV1alpha1PriorityClassHandlerFunc turns a function with the right signature into a watch scheduling v1alpha1 priority class handler
type WatchSchedulingV1alpha1PriorityClassHandlerFunc func(WatchSchedulingV1alpha1PriorityClassParams) middleware.Responder

// Handle executing the request and returning a response
func (fn WatchSchedulingV1alpha1PriorityClassHandlerFunc) Handle(params WatchSchedulingV1alpha1PriorityClassParams) middleware.Responder {
	return fn(params)
}

// WatchSchedulingV1alpha1PriorityClassHandler interface for that can handle valid watch scheduling v1alpha1 priority class params
type WatchSchedulingV1alpha1PriorityClassHandler interface {
	Handle(WatchSchedulingV1alpha1PriorityClassParams) middleware.Responder
}

// NewWatchSchedulingV1alpha1PriorityClass creates a new http.Handler for the watch scheduling v1alpha1 priority class operation
func NewWatchSchedulingV1alpha1PriorityClass(ctx *middleware.Context, handler WatchSchedulingV1alpha1PriorityClassHandler) *WatchSchedulingV1alpha1PriorityClass {
	return &WatchSchedulingV1alpha1PriorityClass{Context: ctx, Handler: handler}
}

/*WatchSchedulingV1alpha1PriorityClass swagger:route GET /apis/scheduling.k8s.io/v1alpha1/watch/priorityclasses/{name} scheduling_v1alpha1 watchSchedulingV1alpha1PriorityClass

watch changes to an object of kind PriorityClass. deprecated: use the 'watch' parameter with a list operation instead, filtered to a single item with the 'fieldSelector' parameter.

*/
type WatchSchedulingV1alpha1PriorityClass struct {
	Context *middleware.Context
	Handler WatchSchedulingV1alpha1PriorityClassHandler
}

func (o *WatchSchedulingV1alpha1PriorityClass) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewWatchSchedulingV1alpha1PriorityClassParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
