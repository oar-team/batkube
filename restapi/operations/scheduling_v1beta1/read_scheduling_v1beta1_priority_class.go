// Code generated by go-swagger; DO NOT EDIT.

package scheduling_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ReadSchedulingV1beta1PriorityClassHandlerFunc turns a function with the right signature into a read scheduling v1beta1 priority class handler
type ReadSchedulingV1beta1PriorityClassHandlerFunc func(ReadSchedulingV1beta1PriorityClassParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ReadSchedulingV1beta1PriorityClassHandlerFunc) Handle(params ReadSchedulingV1beta1PriorityClassParams) middleware.Responder {
	return fn(params)
}

// ReadSchedulingV1beta1PriorityClassHandler interface for that can handle valid read scheduling v1beta1 priority class params
type ReadSchedulingV1beta1PriorityClassHandler interface {
	Handle(ReadSchedulingV1beta1PriorityClassParams) middleware.Responder
}

// NewReadSchedulingV1beta1PriorityClass creates a new http.Handler for the read scheduling v1beta1 priority class operation
func NewReadSchedulingV1beta1PriorityClass(ctx *middleware.Context, handler ReadSchedulingV1beta1PriorityClassHandler) *ReadSchedulingV1beta1PriorityClass {
	return &ReadSchedulingV1beta1PriorityClass{Context: ctx, Handler: handler}
}

/*ReadSchedulingV1beta1PriorityClass swagger:route GET /apis/scheduling.k8s.io/v1beta1/priorityclasses/{name} scheduling_v1beta1 readSchedulingV1beta1PriorityClass

read the specified PriorityClass

*/
type ReadSchedulingV1beta1PriorityClass struct {
	Context *middleware.Context
	Handler ReadSchedulingV1beta1PriorityClassHandler
}

func (o *ReadSchedulingV1beta1PriorityClass) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewReadSchedulingV1beta1PriorityClassParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}