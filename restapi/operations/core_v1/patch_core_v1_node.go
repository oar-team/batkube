// Code generated by go-swagger; DO NOT EDIT.

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PatchCoreV1NodeHandlerFunc turns a function with the right signature into a patch core v1 node handler
type PatchCoreV1NodeHandlerFunc func(PatchCoreV1NodeParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn PatchCoreV1NodeHandlerFunc) Handle(params PatchCoreV1NodeParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// PatchCoreV1NodeHandler interface for that can handle valid patch core v1 node params
type PatchCoreV1NodeHandler interface {
	Handle(PatchCoreV1NodeParams, interface{}) middleware.Responder
}

// NewPatchCoreV1Node creates a new http.Handler for the patch core v1 node operation
func NewPatchCoreV1Node(ctx *middleware.Context, handler PatchCoreV1NodeHandler) *PatchCoreV1Node {
	return &PatchCoreV1Node{Context: ctx, Handler: handler}
}

/*PatchCoreV1Node swagger:route PATCH /api/v1/nodes/{name} core_v1 patchCoreV1Node

partially update the specified Node

*/
type PatchCoreV1Node struct {
	Context *middleware.Context
	Handler PatchCoreV1NodeHandler
}

func (o *PatchCoreV1Node) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPatchCoreV1NodeParams()

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
