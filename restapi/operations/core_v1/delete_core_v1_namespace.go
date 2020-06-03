// Code generated by go-swagger; DO NOT EDIT.

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteCoreV1NamespaceHandlerFunc turns a function with the right signature into a delete core v1 namespace handler
type DeleteCoreV1NamespaceHandlerFunc func(DeleteCoreV1NamespaceParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteCoreV1NamespaceHandlerFunc) Handle(params DeleteCoreV1NamespaceParams) middleware.Responder {
	return fn(params)
}

// DeleteCoreV1NamespaceHandler interface for that can handle valid delete core v1 namespace params
type DeleteCoreV1NamespaceHandler interface {
	Handle(DeleteCoreV1NamespaceParams) middleware.Responder
}

// NewDeleteCoreV1Namespace creates a new http.Handler for the delete core v1 namespace operation
func NewDeleteCoreV1Namespace(ctx *middleware.Context, handler DeleteCoreV1NamespaceHandler) *DeleteCoreV1Namespace {
	return &DeleteCoreV1Namespace{Context: ctx, Handler: handler}
}

/*DeleteCoreV1Namespace swagger:route DELETE /api/v1/namespaces/{name} core_v1 deleteCoreV1Namespace

delete a Namespace

*/
type DeleteCoreV1Namespace struct {
	Context *middleware.Context
	Handler DeleteCoreV1NamespaceHandler
}

func (o *DeleteCoreV1Namespace) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteCoreV1NamespaceParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}