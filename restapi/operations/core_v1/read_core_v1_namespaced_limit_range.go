// Code generated by go-swagger; DO NOT EDIT.

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ReadCoreV1NamespacedLimitRangeHandlerFunc turns a function with the right signature into a read core v1 namespaced limit range handler
type ReadCoreV1NamespacedLimitRangeHandlerFunc func(ReadCoreV1NamespacedLimitRangeParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ReadCoreV1NamespacedLimitRangeHandlerFunc) Handle(params ReadCoreV1NamespacedLimitRangeParams) middleware.Responder {
	return fn(params)
}

// ReadCoreV1NamespacedLimitRangeHandler interface for that can handle valid read core v1 namespaced limit range params
type ReadCoreV1NamespacedLimitRangeHandler interface {
	Handle(ReadCoreV1NamespacedLimitRangeParams) middleware.Responder
}

// NewReadCoreV1NamespacedLimitRange creates a new http.Handler for the read core v1 namespaced limit range operation
func NewReadCoreV1NamespacedLimitRange(ctx *middleware.Context, handler ReadCoreV1NamespacedLimitRangeHandler) *ReadCoreV1NamespacedLimitRange {
	return &ReadCoreV1NamespacedLimitRange{Context: ctx, Handler: handler}
}

/*ReadCoreV1NamespacedLimitRange swagger:route GET /api/v1/namespaces/{namespace}/limitranges/{name} core_v1 readCoreV1NamespacedLimitRange

read the specified LimitRange

*/
type ReadCoreV1NamespacedLimitRange struct {
	Context *middleware.Context
	Handler ReadCoreV1NamespacedLimitRangeHandler
}

func (o *ReadCoreV1NamespacedLimitRange) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewReadCoreV1NamespacedLimitRangeParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
