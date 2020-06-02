// Code generated by go-swagger; DO NOT EDIT.

package scheduling_v1alpha1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ReplaceSchedulingV1alpha1PriorityClassHandlerFunc turns a function with the right signature into a replace scheduling v1alpha1 priority class handler
type ReplaceSchedulingV1alpha1PriorityClassHandlerFunc func(ReplaceSchedulingV1alpha1PriorityClassParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn ReplaceSchedulingV1alpha1PriorityClassHandlerFunc) Handle(params ReplaceSchedulingV1alpha1PriorityClassParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// ReplaceSchedulingV1alpha1PriorityClassHandler interface for that can handle valid replace scheduling v1alpha1 priority class params
type ReplaceSchedulingV1alpha1PriorityClassHandler interface {
	Handle(ReplaceSchedulingV1alpha1PriorityClassParams, interface{}) middleware.Responder
}

// NewReplaceSchedulingV1alpha1PriorityClass creates a new http.Handler for the replace scheduling v1alpha1 priority class operation
func NewReplaceSchedulingV1alpha1PriorityClass(ctx *middleware.Context, handler ReplaceSchedulingV1alpha1PriorityClassHandler) *ReplaceSchedulingV1alpha1PriorityClass {
	return &ReplaceSchedulingV1alpha1PriorityClass{Context: ctx, Handler: handler}
}

/*ReplaceSchedulingV1alpha1PriorityClass swagger:route PUT /apis/scheduling.k8s.io/v1alpha1/priorityclasses/{name} scheduling_v1alpha1 replaceSchedulingV1alpha1PriorityClass

replace the specified PriorityClass

*/
type ReplaceSchedulingV1alpha1PriorityClass struct {
	Context *middleware.Context
	Handler ReplaceSchedulingV1alpha1PriorityClassHandler
}

func (o *ReplaceSchedulingV1alpha1PriorityClass) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewReplaceSchedulingV1alpha1PriorityClassParams()

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
