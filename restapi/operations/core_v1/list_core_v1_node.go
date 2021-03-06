// Code generated by go-swagger; DO NOT EDIT.

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ListCoreV1NodeHandlerFunc turns a function with the right signature into a list core v1 node handler
type ListCoreV1NodeHandlerFunc func(ListCoreV1NodeParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ListCoreV1NodeHandlerFunc) Handle(params ListCoreV1NodeParams) middleware.Responder {
	return fn(params)
}

// ListCoreV1NodeHandler interface for that can handle valid list core v1 node params
type ListCoreV1NodeHandler interface {
	Handle(ListCoreV1NodeParams) middleware.Responder
}

// NewListCoreV1Node creates a new http.Handler for the list core v1 node operation
func NewListCoreV1Node(ctx *middleware.Context, handler ListCoreV1NodeHandler) *ListCoreV1Node {
	return &ListCoreV1Node{Context: ctx, Handler: handler}
}

/*ListCoreV1Node swagger:route GET /api/v1/nodes core_v1 listCoreV1Node

list or watch objects of kind Node

*/
type ListCoreV1Node struct {
	Context *middleware.Context
	Handler ListCoreV1NodeHandler
}

func (o *ListCoreV1Node) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewListCoreV1NodeParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
