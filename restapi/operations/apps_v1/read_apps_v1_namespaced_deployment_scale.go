// Code generated by go-swagger; DO NOT EDIT.

package apps_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ReadAppsV1NamespacedDeploymentScaleHandlerFunc turns a function with the right signature into a read apps v1 namespaced deployment scale handler
type ReadAppsV1NamespacedDeploymentScaleHandlerFunc func(ReadAppsV1NamespacedDeploymentScaleParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ReadAppsV1NamespacedDeploymentScaleHandlerFunc) Handle(params ReadAppsV1NamespacedDeploymentScaleParams) middleware.Responder {
	return fn(params)
}

// ReadAppsV1NamespacedDeploymentScaleHandler interface for that can handle valid read apps v1 namespaced deployment scale params
type ReadAppsV1NamespacedDeploymentScaleHandler interface {
	Handle(ReadAppsV1NamespacedDeploymentScaleParams) middleware.Responder
}

// NewReadAppsV1NamespacedDeploymentScale creates a new http.Handler for the read apps v1 namespaced deployment scale operation
func NewReadAppsV1NamespacedDeploymentScale(ctx *middleware.Context, handler ReadAppsV1NamespacedDeploymentScaleHandler) *ReadAppsV1NamespacedDeploymentScale {
	return &ReadAppsV1NamespacedDeploymentScale{Context: ctx, Handler: handler}
}

/*ReadAppsV1NamespacedDeploymentScale swagger:route GET /apis/apps/v1/namespaces/{namespace}/deployments/{name}/scale apps_v1 readAppsV1NamespacedDeploymentScale

read scale of the specified Deployment

*/
type ReadAppsV1NamespacedDeploymentScale struct {
	Context *middleware.Context
	Handler ReadAppsV1NamespacedDeploymentScaleHandler
}

func (o *ReadAppsV1NamespacedDeploymentScale) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewReadAppsV1NamespacedDeploymentScaleParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
