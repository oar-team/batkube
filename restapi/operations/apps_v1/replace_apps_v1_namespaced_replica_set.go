// Code generated by go-swagger; DO NOT EDIT.

package apps_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ReplaceAppsV1NamespacedReplicaSetHandlerFunc turns a function with the right signature into a replace apps v1 namespaced replica set handler
type ReplaceAppsV1NamespacedReplicaSetHandlerFunc func(ReplaceAppsV1NamespacedReplicaSetParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ReplaceAppsV1NamespacedReplicaSetHandlerFunc) Handle(params ReplaceAppsV1NamespacedReplicaSetParams) middleware.Responder {
	return fn(params)
}

// ReplaceAppsV1NamespacedReplicaSetHandler interface for that can handle valid replace apps v1 namespaced replica set params
type ReplaceAppsV1NamespacedReplicaSetHandler interface {
	Handle(ReplaceAppsV1NamespacedReplicaSetParams) middleware.Responder
}

// NewReplaceAppsV1NamespacedReplicaSet creates a new http.Handler for the replace apps v1 namespaced replica set operation
func NewReplaceAppsV1NamespacedReplicaSet(ctx *middleware.Context, handler ReplaceAppsV1NamespacedReplicaSetHandler) *ReplaceAppsV1NamespacedReplicaSet {
	return &ReplaceAppsV1NamespacedReplicaSet{Context: ctx, Handler: handler}
}

/*ReplaceAppsV1NamespacedReplicaSet swagger:route PUT /apis/apps/v1/namespaces/{namespace}/replicasets/{name} apps_v1 replaceAppsV1NamespacedReplicaSet

replace the specified ReplicaSet

*/
type ReplaceAppsV1NamespacedReplicaSet struct {
	Context *middleware.Context
	Handler ReplaceAppsV1NamespacedReplicaSetHandler
}

func (o *ReplaceAppsV1NamespacedReplicaSet) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewReplaceAppsV1NamespacedReplicaSetParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
