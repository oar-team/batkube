// Code generated by go-swagger; DO NOT EDIT.

package storage_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteStorageV1beta1CSINodeHandlerFunc turns a function with the right signature into a delete storage v1beta1 c s i node handler
type DeleteStorageV1beta1CSINodeHandlerFunc func(DeleteStorageV1beta1CSINodeParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteStorageV1beta1CSINodeHandlerFunc) Handle(params DeleteStorageV1beta1CSINodeParams) middleware.Responder {
	return fn(params)
}

// DeleteStorageV1beta1CSINodeHandler interface for that can handle valid delete storage v1beta1 c s i node params
type DeleteStorageV1beta1CSINodeHandler interface {
	Handle(DeleteStorageV1beta1CSINodeParams) middleware.Responder
}

// NewDeleteStorageV1beta1CSINode creates a new http.Handler for the delete storage v1beta1 c s i node operation
func NewDeleteStorageV1beta1CSINode(ctx *middleware.Context, handler DeleteStorageV1beta1CSINodeHandler) *DeleteStorageV1beta1CSINode {
	return &DeleteStorageV1beta1CSINode{Context: ctx, Handler: handler}
}

/*DeleteStorageV1beta1CSINode swagger:route DELETE /apis/storage.k8s.io/v1beta1/csinodes/{name} storage_v1beta1 deleteStorageV1beta1CSINode

delete a CSINode

*/
type DeleteStorageV1beta1CSINode struct {
	Context *middleware.Context
	Handler DeleteStorageV1beta1CSINodeHandler
}

func (o *DeleteStorageV1beta1CSINode) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteStorageV1beta1CSINodeParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
