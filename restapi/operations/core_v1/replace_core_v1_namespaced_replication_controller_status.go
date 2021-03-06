// Code generated by go-swagger; DO NOT EDIT.

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ReplaceCoreV1NamespacedReplicationControllerStatusHandlerFunc turns a function with the right signature into a replace core v1 namespaced replication controller status handler
type ReplaceCoreV1NamespacedReplicationControllerStatusHandlerFunc func(ReplaceCoreV1NamespacedReplicationControllerStatusParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ReplaceCoreV1NamespacedReplicationControllerStatusHandlerFunc) Handle(params ReplaceCoreV1NamespacedReplicationControllerStatusParams) middleware.Responder {
	return fn(params)
}

// ReplaceCoreV1NamespacedReplicationControllerStatusHandler interface for that can handle valid replace core v1 namespaced replication controller status params
type ReplaceCoreV1NamespacedReplicationControllerStatusHandler interface {
	Handle(ReplaceCoreV1NamespacedReplicationControllerStatusParams) middleware.Responder
}

// NewReplaceCoreV1NamespacedReplicationControllerStatus creates a new http.Handler for the replace core v1 namespaced replication controller status operation
func NewReplaceCoreV1NamespacedReplicationControllerStatus(ctx *middleware.Context, handler ReplaceCoreV1NamespacedReplicationControllerStatusHandler) *ReplaceCoreV1NamespacedReplicationControllerStatus {
	return &ReplaceCoreV1NamespacedReplicationControllerStatus{Context: ctx, Handler: handler}
}

/*ReplaceCoreV1NamespacedReplicationControllerStatus swagger:route PUT /api/v1/namespaces/{namespace}/replicationcontrollers/{name}/status core_v1 replaceCoreV1NamespacedReplicationControllerStatus

replace status of the specified ReplicationController

*/
type ReplaceCoreV1NamespacedReplicationControllerStatus struct {
	Context *middleware.Context
	Handler ReplaceCoreV1NamespacedReplicationControllerStatusHandler
}

func (o *ReplaceCoreV1NamespacedReplicationControllerStatus) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewReplaceCoreV1NamespacedReplicationControllerStatusParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
