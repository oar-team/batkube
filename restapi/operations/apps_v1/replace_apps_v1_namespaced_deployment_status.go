// Code generated by go-swagger; DO NOT EDIT.

package apps_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ReplaceAppsV1NamespacedDeploymentStatusHandlerFunc turns a function with the right signature into a replace apps v1 namespaced deployment status handler
type ReplaceAppsV1NamespacedDeploymentStatusHandlerFunc func(ReplaceAppsV1NamespacedDeploymentStatusParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn ReplaceAppsV1NamespacedDeploymentStatusHandlerFunc) Handle(params ReplaceAppsV1NamespacedDeploymentStatusParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// ReplaceAppsV1NamespacedDeploymentStatusHandler interface for that can handle valid replace apps v1 namespaced deployment status params
type ReplaceAppsV1NamespacedDeploymentStatusHandler interface {
	Handle(ReplaceAppsV1NamespacedDeploymentStatusParams, interface{}) middleware.Responder
}

// NewReplaceAppsV1NamespacedDeploymentStatus creates a new http.Handler for the replace apps v1 namespaced deployment status operation
func NewReplaceAppsV1NamespacedDeploymentStatus(ctx *middleware.Context, handler ReplaceAppsV1NamespacedDeploymentStatusHandler) *ReplaceAppsV1NamespacedDeploymentStatus {
	return &ReplaceAppsV1NamespacedDeploymentStatus{Context: ctx, Handler: handler}
}

/*ReplaceAppsV1NamespacedDeploymentStatus swagger:route PUT /apis/apps/v1/namespaces/{namespace}/deployments/{name}/status apps_v1 replaceAppsV1NamespacedDeploymentStatus

replace status of the specified Deployment

*/
type ReplaceAppsV1NamespacedDeploymentStatus struct {
	Context *middleware.Context
	Handler ReplaceAppsV1NamespacedDeploymentStatusHandler
}

func (o *ReplaceAppsV1NamespacedDeploymentStatus) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewReplaceAppsV1NamespacedDeploymentStatusParams()

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
