// Code generated by go-swagger; DO NOT EDIT.

package scheduling_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// CreateSchedulingV1beta1PriorityClassHandlerFunc turns a function with the right signature into a create scheduling v1beta1 priority class handler
type CreateSchedulingV1beta1PriorityClassHandlerFunc func(CreateSchedulingV1beta1PriorityClassParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateSchedulingV1beta1PriorityClassHandlerFunc) Handle(params CreateSchedulingV1beta1PriorityClassParams) middleware.Responder {
	return fn(params)
}

// CreateSchedulingV1beta1PriorityClassHandler interface for that can handle valid create scheduling v1beta1 priority class params
type CreateSchedulingV1beta1PriorityClassHandler interface {
	Handle(CreateSchedulingV1beta1PriorityClassParams) middleware.Responder
}

// NewCreateSchedulingV1beta1PriorityClass creates a new http.Handler for the create scheduling v1beta1 priority class operation
func NewCreateSchedulingV1beta1PriorityClass(ctx *middleware.Context, handler CreateSchedulingV1beta1PriorityClassHandler) *CreateSchedulingV1beta1PriorityClass {
	return &CreateSchedulingV1beta1PriorityClass{Context: ctx, Handler: handler}
}

/*CreateSchedulingV1beta1PriorityClass swagger:route POST /apis/scheduling.k8s.io/v1beta1/priorityclasses scheduling_v1beta1 createSchedulingV1beta1PriorityClass

create a PriorityClass

*/
type CreateSchedulingV1beta1PriorityClass struct {
	Context *middleware.Context
	Handler CreateSchedulingV1beta1PriorityClassHandler
}

func (o *CreateSchedulingV1beta1PriorityClass) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCreateSchedulingV1beta1PriorityClassParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
