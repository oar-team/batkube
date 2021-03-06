// Code generated by go-swagger; DO NOT EDIT.

package storage_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ReplaceStorageV1beta1CSINodeHandlerFunc turns a function with the right signature into a replace storage v1beta1 c s i node handler
type ReplaceStorageV1beta1CSINodeHandlerFunc func(ReplaceStorageV1beta1CSINodeParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ReplaceStorageV1beta1CSINodeHandlerFunc) Handle(params ReplaceStorageV1beta1CSINodeParams) middleware.Responder {
	return fn(params)
}

// ReplaceStorageV1beta1CSINodeHandler interface for that can handle valid replace storage v1beta1 c s i node params
type ReplaceStorageV1beta1CSINodeHandler interface {
	Handle(ReplaceStorageV1beta1CSINodeParams) middleware.Responder
}

// NewReplaceStorageV1beta1CSINode creates a new http.Handler for the replace storage v1beta1 c s i node operation
func NewReplaceStorageV1beta1CSINode(ctx *middleware.Context, handler ReplaceStorageV1beta1CSINodeHandler) *ReplaceStorageV1beta1CSINode {
	return &ReplaceStorageV1beta1CSINode{Context: ctx, Handler: handler}
}

/*ReplaceStorageV1beta1CSINode swagger:route PUT /apis/storage.k8s.io/v1beta1/csinodes/{name} storage_v1beta1 replaceStorageV1beta1CSINode

replace the specified CSINode

*/
type ReplaceStorageV1beta1CSINode struct {
	Context *middleware.Context
	Handler ReplaceStorageV1beta1CSINodeHandler
}

func (o *ReplaceStorageV1beta1CSINode) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewReplaceStorageV1beta1CSINodeParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
