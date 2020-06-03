// Code generated by go-swagger; DO NOT EDIT.

package apps_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ReplaceAppsV1NamespacedReplicaSetStatusHandlerFunc turns a function with the right signature into a replace apps v1 namespaced replica set status handler
type ReplaceAppsV1NamespacedReplicaSetStatusHandlerFunc func(ReplaceAppsV1NamespacedReplicaSetStatusParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ReplaceAppsV1NamespacedReplicaSetStatusHandlerFunc) Handle(params ReplaceAppsV1NamespacedReplicaSetStatusParams) middleware.Responder {
	return fn(params)
}

// ReplaceAppsV1NamespacedReplicaSetStatusHandler interface for that can handle valid replace apps v1 namespaced replica set status params
type ReplaceAppsV1NamespacedReplicaSetStatusHandler interface {
	Handle(ReplaceAppsV1NamespacedReplicaSetStatusParams) middleware.Responder
}

// NewReplaceAppsV1NamespacedReplicaSetStatus creates a new http.Handler for the replace apps v1 namespaced replica set status operation
func NewReplaceAppsV1NamespacedReplicaSetStatus(ctx *middleware.Context, handler ReplaceAppsV1NamespacedReplicaSetStatusHandler) *ReplaceAppsV1NamespacedReplicaSetStatus {
	return &ReplaceAppsV1NamespacedReplicaSetStatus{Context: ctx, Handler: handler}
}

/*ReplaceAppsV1NamespacedReplicaSetStatus swagger:route PUT /apis/apps/v1/namespaces/{namespace}/replicasets/{name}/status apps_v1 replaceAppsV1NamespacedReplicaSetStatus

replace status of the specified ReplicaSet

*/
type ReplaceAppsV1NamespacedReplicaSetStatus struct {
	Context *middleware.Context
	Handler ReplaceAppsV1NamespacedReplicaSetStatusHandler
}

func (o *ReplaceAppsV1NamespacedReplicaSetStatus) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewReplaceAppsV1NamespacedReplicaSetStatusParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}