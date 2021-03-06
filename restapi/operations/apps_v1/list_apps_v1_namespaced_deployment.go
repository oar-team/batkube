// Code generated by go-swagger; DO NOT EDIT.

package apps_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ListAppsV1NamespacedDeploymentHandlerFunc turns a function with the right signature into a list apps v1 namespaced deployment handler
type ListAppsV1NamespacedDeploymentHandlerFunc func(ListAppsV1NamespacedDeploymentParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ListAppsV1NamespacedDeploymentHandlerFunc) Handle(params ListAppsV1NamespacedDeploymentParams) middleware.Responder {
	return fn(params)
}

// ListAppsV1NamespacedDeploymentHandler interface for that can handle valid list apps v1 namespaced deployment params
type ListAppsV1NamespacedDeploymentHandler interface {
	Handle(ListAppsV1NamespacedDeploymentParams) middleware.Responder
}

// NewListAppsV1NamespacedDeployment creates a new http.Handler for the list apps v1 namespaced deployment operation
func NewListAppsV1NamespacedDeployment(ctx *middleware.Context, handler ListAppsV1NamespacedDeploymentHandler) *ListAppsV1NamespacedDeployment {
	return &ListAppsV1NamespacedDeployment{Context: ctx, Handler: handler}
}

/*ListAppsV1NamespacedDeployment swagger:route GET /apis/apps/v1/namespaces/{namespace}/deployments apps_v1 listAppsV1NamespacedDeployment

list or watch objects of kind Deployment

*/
type ListAppsV1NamespacedDeployment struct {
	Context *middleware.Context
	Handler ListAppsV1NamespacedDeploymentHandler
}

func (o *ListAppsV1NamespacedDeployment) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewListAppsV1NamespacedDeploymentParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
