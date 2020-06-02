// Code generated by go-swagger; DO NOT EDIT.

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ReadCoreV1NamespacedReplicationControllerHandlerFunc turns a function with the right signature into a read core v1 namespaced replication controller handler
type ReadCoreV1NamespacedReplicationControllerHandlerFunc func(ReadCoreV1NamespacedReplicationControllerParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn ReadCoreV1NamespacedReplicationControllerHandlerFunc) Handle(params ReadCoreV1NamespacedReplicationControllerParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// ReadCoreV1NamespacedReplicationControllerHandler interface for that can handle valid read core v1 namespaced replication controller params
type ReadCoreV1NamespacedReplicationControllerHandler interface {
	Handle(ReadCoreV1NamespacedReplicationControllerParams, interface{}) middleware.Responder
}

// NewReadCoreV1NamespacedReplicationController creates a new http.Handler for the read core v1 namespaced replication controller operation
func NewReadCoreV1NamespacedReplicationController(ctx *middleware.Context, handler ReadCoreV1NamespacedReplicationControllerHandler) *ReadCoreV1NamespacedReplicationController {
	return &ReadCoreV1NamespacedReplicationController{Context: ctx, Handler: handler}
}

/*ReadCoreV1NamespacedReplicationController swagger:route GET /api/v1/namespaces/{namespace}/replicationcontrollers/{name} core_v1 readCoreV1NamespacedReplicationController

read the specified ReplicationController

*/
type ReadCoreV1NamespacedReplicationController struct {
	Context *middleware.Context
	Handler ReadCoreV1NamespacedReplicationControllerHandler
}

func (o *ReadCoreV1NamespacedReplicationController) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewReadCoreV1NamespacedReplicationControllerParams()

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
