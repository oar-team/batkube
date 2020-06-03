// Code generated by go-swagger; DO NOT EDIT.

package apps_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ReplaceAppsV1NamespacedDeploymentHandlerFunc turns a function with the right signature into a replace apps v1 namespaced deployment handler
type ReplaceAppsV1NamespacedDeploymentHandlerFunc func(ReplaceAppsV1NamespacedDeploymentParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ReplaceAppsV1NamespacedDeploymentHandlerFunc) Handle(params ReplaceAppsV1NamespacedDeploymentParams) middleware.Responder {
	return fn(params)
}

// ReplaceAppsV1NamespacedDeploymentHandler interface for that can handle valid replace apps v1 namespaced deployment params
type ReplaceAppsV1NamespacedDeploymentHandler interface {
	Handle(ReplaceAppsV1NamespacedDeploymentParams) middleware.Responder
}

// NewReplaceAppsV1NamespacedDeployment creates a new http.Handler for the replace apps v1 namespaced deployment operation
func NewReplaceAppsV1NamespacedDeployment(ctx *middleware.Context, handler ReplaceAppsV1NamespacedDeploymentHandler) *ReplaceAppsV1NamespacedDeployment {
	return &ReplaceAppsV1NamespacedDeployment{Context: ctx, Handler: handler}
}

/*ReplaceAppsV1NamespacedDeployment swagger:route PUT /apis/apps/v1/namespaces/{namespace}/deployments/{name} apps_v1 replaceAppsV1NamespacedDeployment

replace the specified Deployment

*/
type ReplaceAppsV1NamespacedDeployment struct {
	Context *middleware.Context
	Handler ReplaceAppsV1NamespacedDeploymentHandler
}

func (o *ReplaceAppsV1NamespacedDeployment) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewReplaceAppsV1NamespacedDeploymentParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}