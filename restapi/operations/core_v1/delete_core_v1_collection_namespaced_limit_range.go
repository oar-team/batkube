// Code generated by go-swagger; DO NOT EDIT.

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteCoreV1CollectionNamespacedLimitRangeHandlerFunc turns a function with the right signature into a delete core v1 collection namespaced limit range handler
type DeleteCoreV1CollectionNamespacedLimitRangeHandlerFunc func(DeleteCoreV1CollectionNamespacedLimitRangeParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteCoreV1CollectionNamespacedLimitRangeHandlerFunc) Handle(params DeleteCoreV1CollectionNamespacedLimitRangeParams) middleware.Responder {
	return fn(params)
}

// DeleteCoreV1CollectionNamespacedLimitRangeHandler interface for that can handle valid delete core v1 collection namespaced limit range params
type DeleteCoreV1CollectionNamespacedLimitRangeHandler interface {
	Handle(DeleteCoreV1CollectionNamespacedLimitRangeParams) middleware.Responder
}

// NewDeleteCoreV1CollectionNamespacedLimitRange creates a new http.Handler for the delete core v1 collection namespaced limit range operation
func NewDeleteCoreV1CollectionNamespacedLimitRange(ctx *middleware.Context, handler DeleteCoreV1CollectionNamespacedLimitRangeHandler) *DeleteCoreV1CollectionNamespacedLimitRange {
	return &DeleteCoreV1CollectionNamespacedLimitRange{Context: ctx, Handler: handler}
}

/*DeleteCoreV1CollectionNamespacedLimitRange swagger:route DELETE /api/v1/namespaces/{namespace}/limitranges core_v1 deleteCoreV1CollectionNamespacedLimitRange

delete collection of LimitRange

*/
type DeleteCoreV1CollectionNamespacedLimitRange struct {
	Context *middleware.Context
	Handler DeleteCoreV1CollectionNamespacedLimitRangeHandler
}

func (o *DeleteCoreV1CollectionNamespacedLimitRange) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteCoreV1CollectionNamespacedLimitRangeParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
