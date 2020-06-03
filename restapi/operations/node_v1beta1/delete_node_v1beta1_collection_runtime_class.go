// Code generated by go-swagger; DO NOT EDIT.

package node_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteNodeV1beta1CollectionRuntimeClassHandlerFunc turns a function with the right signature into a delete node v1beta1 collection runtime class handler
type DeleteNodeV1beta1CollectionRuntimeClassHandlerFunc func(DeleteNodeV1beta1CollectionRuntimeClassParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteNodeV1beta1CollectionRuntimeClassHandlerFunc) Handle(params DeleteNodeV1beta1CollectionRuntimeClassParams) middleware.Responder {
	return fn(params)
}

// DeleteNodeV1beta1CollectionRuntimeClassHandler interface for that can handle valid delete node v1beta1 collection runtime class params
type DeleteNodeV1beta1CollectionRuntimeClassHandler interface {
	Handle(DeleteNodeV1beta1CollectionRuntimeClassParams) middleware.Responder
}

// NewDeleteNodeV1beta1CollectionRuntimeClass creates a new http.Handler for the delete node v1beta1 collection runtime class operation
func NewDeleteNodeV1beta1CollectionRuntimeClass(ctx *middleware.Context, handler DeleteNodeV1beta1CollectionRuntimeClassHandler) *DeleteNodeV1beta1CollectionRuntimeClass {
	return &DeleteNodeV1beta1CollectionRuntimeClass{Context: ctx, Handler: handler}
}

/*DeleteNodeV1beta1CollectionRuntimeClass swagger:route DELETE /apis/node.k8s.io/v1beta1/runtimeclasses node_v1beta1 deleteNodeV1beta1CollectionRuntimeClass

delete collection of RuntimeClass

*/
type DeleteNodeV1beta1CollectionRuntimeClass struct {
	Context *middleware.Context
	Handler DeleteNodeV1beta1CollectionRuntimeClassHandler
}

func (o *DeleteNodeV1beta1CollectionRuntimeClass) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteNodeV1beta1CollectionRuntimeClassParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
