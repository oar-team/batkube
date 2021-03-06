// Code generated by go-swagger; DO NOT EDIT.

package scheduling_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteSchedulingV1CollectionPriorityClassHandlerFunc turns a function with the right signature into a delete scheduling v1 collection priority class handler
type DeleteSchedulingV1CollectionPriorityClassHandlerFunc func(DeleteSchedulingV1CollectionPriorityClassParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteSchedulingV1CollectionPriorityClassHandlerFunc) Handle(params DeleteSchedulingV1CollectionPriorityClassParams) middleware.Responder {
	return fn(params)
}

// DeleteSchedulingV1CollectionPriorityClassHandler interface for that can handle valid delete scheduling v1 collection priority class params
type DeleteSchedulingV1CollectionPriorityClassHandler interface {
	Handle(DeleteSchedulingV1CollectionPriorityClassParams) middleware.Responder
}

// NewDeleteSchedulingV1CollectionPriorityClass creates a new http.Handler for the delete scheduling v1 collection priority class operation
func NewDeleteSchedulingV1CollectionPriorityClass(ctx *middleware.Context, handler DeleteSchedulingV1CollectionPriorityClassHandler) *DeleteSchedulingV1CollectionPriorityClass {
	return &DeleteSchedulingV1CollectionPriorityClass{Context: ctx, Handler: handler}
}

/*DeleteSchedulingV1CollectionPriorityClass swagger:route DELETE /apis/scheduling.k8s.io/v1/priorityclasses scheduling_v1 deleteSchedulingV1CollectionPriorityClass

delete collection of PriorityClass

*/
type DeleteSchedulingV1CollectionPriorityClass struct {
	Context *middleware.Context
	Handler DeleteSchedulingV1CollectionPriorityClassHandler
}

func (o *DeleteSchedulingV1CollectionPriorityClass) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteSchedulingV1CollectionPriorityClassParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
