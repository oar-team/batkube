// Code generated by go-swagger; DO NOT EDIT.

package apps_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ReadAppsV1NamespacedDeploymentStatusHandlerFunc turns a function with the right signature into a read apps v1 namespaced deployment status handler
type ReadAppsV1NamespacedDeploymentStatusHandlerFunc func(ReadAppsV1NamespacedDeploymentStatusParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ReadAppsV1NamespacedDeploymentStatusHandlerFunc) Handle(params ReadAppsV1NamespacedDeploymentStatusParams) middleware.Responder {
	return fn(params)
}

// ReadAppsV1NamespacedDeploymentStatusHandler interface for that can handle valid read apps v1 namespaced deployment status params
type ReadAppsV1NamespacedDeploymentStatusHandler interface {
	Handle(ReadAppsV1NamespacedDeploymentStatusParams) middleware.Responder
}

// NewReadAppsV1NamespacedDeploymentStatus creates a new http.Handler for the read apps v1 namespaced deployment status operation
func NewReadAppsV1NamespacedDeploymentStatus(ctx *middleware.Context, handler ReadAppsV1NamespacedDeploymentStatusHandler) *ReadAppsV1NamespacedDeploymentStatus {
	return &ReadAppsV1NamespacedDeploymentStatus{Context: ctx, Handler: handler}
}

/*ReadAppsV1NamespacedDeploymentStatus swagger:route GET /apis/apps/v1/namespaces/{namespace}/deployments/{name}/status apps_v1 readAppsV1NamespacedDeploymentStatus

read status of the specified Deployment

*/
type ReadAppsV1NamespacedDeploymentStatus struct {
	Context *middleware.Context
	Handler ReadAppsV1NamespacedDeploymentStatusHandler
}

func (o *ReadAppsV1NamespacedDeploymentStatus) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewReadAppsV1NamespacedDeploymentStatusParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
