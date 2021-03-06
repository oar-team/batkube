// Code generated by go-swagger; DO NOT EDIT.

package scheduling_v1alpha1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ListSchedulingV1alpha1PriorityClassHandlerFunc turns a function with the right signature into a list scheduling v1alpha1 priority class handler
type ListSchedulingV1alpha1PriorityClassHandlerFunc func(ListSchedulingV1alpha1PriorityClassParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ListSchedulingV1alpha1PriorityClassHandlerFunc) Handle(params ListSchedulingV1alpha1PriorityClassParams) middleware.Responder {
	return fn(params)
}

// ListSchedulingV1alpha1PriorityClassHandler interface for that can handle valid list scheduling v1alpha1 priority class params
type ListSchedulingV1alpha1PriorityClassHandler interface {
	Handle(ListSchedulingV1alpha1PriorityClassParams) middleware.Responder
}

// NewListSchedulingV1alpha1PriorityClass creates a new http.Handler for the list scheduling v1alpha1 priority class operation
func NewListSchedulingV1alpha1PriorityClass(ctx *middleware.Context, handler ListSchedulingV1alpha1PriorityClassHandler) *ListSchedulingV1alpha1PriorityClass {
	return &ListSchedulingV1alpha1PriorityClass{Context: ctx, Handler: handler}
}

/*ListSchedulingV1alpha1PriorityClass swagger:route GET /apis/scheduling.k8s.io/v1alpha1/priorityclasses scheduling_v1alpha1 listSchedulingV1alpha1PriorityClass

list or watch objects of kind PriorityClass

*/
type ListSchedulingV1alpha1PriorityClass struct {
	Context *middleware.Context
	Handler ListSchedulingV1alpha1PriorityClassHandler
}

func (o *ListSchedulingV1alpha1PriorityClass) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewListSchedulingV1alpha1PriorityClassParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
