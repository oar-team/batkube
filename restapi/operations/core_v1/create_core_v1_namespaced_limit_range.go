// Code generated by go-swagger; DO NOT EDIT.

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// CreateCoreV1NamespacedLimitRangeHandlerFunc turns a function with the right signature into a create core v1 namespaced limit range handler
type CreateCoreV1NamespacedLimitRangeHandlerFunc func(CreateCoreV1NamespacedLimitRangeParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateCoreV1NamespacedLimitRangeHandlerFunc) Handle(params CreateCoreV1NamespacedLimitRangeParams) middleware.Responder {
	return fn(params)
}

// CreateCoreV1NamespacedLimitRangeHandler interface for that can handle valid create core v1 namespaced limit range params
type CreateCoreV1NamespacedLimitRangeHandler interface {
	Handle(CreateCoreV1NamespacedLimitRangeParams) middleware.Responder
}

// NewCreateCoreV1NamespacedLimitRange creates a new http.Handler for the create core v1 namespaced limit range operation
func NewCreateCoreV1NamespacedLimitRange(ctx *middleware.Context, handler CreateCoreV1NamespacedLimitRangeHandler) *CreateCoreV1NamespacedLimitRange {
	return &CreateCoreV1NamespacedLimitRange{Context: ctx, Handler: handler}
}

/*CreateCoreV1NamespacedLimitRange swagger:route POST /api/v1/namespaces/{namespace}/limitranges core_v1 createCoreV1NamespacedLimitRange

create a LimitRange

*/
type CreateCoreV1NamespacedLimitRange struct {
	Context *middleware.Context
	Handler CreateCoreV1NamespacedLimitRangeHandler
}

func (o *CreateCoreV1NamespacedLimitRange) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCreateCoreV1NamespacedLimitRangeParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
