// Code generated by go-swagger; DO NOT EDIT.

package node_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PatchNodeV1beta1RuntimeClassHandlerFunc turns a function with the right signature into a patch node v1beta1 runtime class handler
type PatchNodeV1beta1RuntimeClassHandlerFunc func(PatchNodeV1beta1RuntimeClassParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn PatchNodeV1beta1RuntimeClassHandlerFunc) Handle(params PatchNodeV1beta1RuntimeClassParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// PatchNodeV1beta1RuntimeClassHandler interface for that can handle valid patch node v1beta1 runtime class params
type PatchNodeV1beta1RuntimeClassHandler interface {
	Handle(PatchNodeV1beta1RuntimeClassParams, interface{}) middleware.Responder
}

// NewPatchNodeV1beta1RuntimeClass creates a new http.Handler for the patch node v1beta1 runtime class operation
func NewPatchNodeV1beta1RuntimeClass(ctx *middleware.Context, handler PatchNodeV1beta1RuntimeClassHandler) *PatchNodeV1beta1RuntimeClass {
	return &PatchNodeV1beta1RuntimeClass{Context: ctx, Handler: handler}
}

/*PatchNodeV1beta1RuntimeClass swagger:route PATCH /apis/node.k8s.io/v1beta1/runtimeclasses/{name} node_v1beta1 patchNodeV1beta1RuntimeClass

partially update the specified RuntimeClass

*/
type PatchNodeV1beta1RuntimeClass struct {
	Context *middleware.Context
	Handler PatchNodeV1beta1RuntimeClassHandler
}

func (o *PatchNodeV1beta1RuntimeClass) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPatchNodeV1beta1RuntimeClassParams()

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
