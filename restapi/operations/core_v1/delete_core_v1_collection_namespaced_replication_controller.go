// Code generated by go-swagger; DO NOT EDIT.

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteCoreV1CollectionNamespacedReplicationControllerHandlerFunc turns a function with the right signature into a delete core v1 collection namespaced replication controller handler
type DeleteCoreV1CollectionNamespacedReplicationControllerHandlerFunc func(DeleteCoreV1CollectionNamespacedReplicationControllerParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteCoreV1CollectionNamespacedReplicationControllerHandlerFunc) Handle(params DeleteCoreV1CollectionNamespacedReplicationControllerParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// DeleteCoreV1CollectionNamespacedReplicationControllerHandler interface for that can handle valid delete core v1 collection namespaced replication controller params
type DeleteCoreV1CollectionNamespacedReplicationControllerHandler interface {
	Handle(DeleteCoreV1CollectionNamespacedReplicationControllerParams, interface{}) middleware.Responder
}

// NewDeleteCoreV1CollectionNamespacedReplicationController creates a new http.Handler for the delete core v1 collection namespaced replication controller operation
func NewDeleteCoreV1CollectionNamespacedReplicationController(ctx *middleware.Context, handler DeleteCoreV1CollectionNamespacedReplicationControllerHandler) *DeleteCoreV1CollectionNamespacedReplicationController {
	return &DeleteCoreV1CollectionNamespacedReplicationController{Context: ctx, Handler: handler}
}

/*DeleteCoreV1CollectionNamespacedReplicationController swagger:route DELETE /api/v1/namespaces/{namespace}/replicationcontrollers core_v1 deleteCoreV1CollectionNamespacedReplicationController

delete collection of ReplicationController

*/
type DeleteCoreV1CollectionNamespacedReplicationController struct {
	Context *middleware.Context
	Handler DeleteCoreV1CollectionNamespacedReplicationControllerHandler
}

func (o *DeleteCoreV1CollectionNamespacedReplicationController) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteCoreV1CollectionNamespacedReplicationControllerParams()

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
