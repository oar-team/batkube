// Code generated by go-swagger; DO NOT EDIT.

package apps_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ReadAppsV1NamespacedDeploymentHandlerFunc turns a function with the right signature into a read apps v1 namespaced deployment handler
type ReadAppsV1NamespacedDeploymentHandlerFunc func(ReadAppsV1NamespacedDeploymentParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn ReadAppsV1NamespacedDeploymentHandlerFunc) Handle(params ReadAppsV1NamespacedDeploymentParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// ReadAppsV1NamespacedDeploymentHandler interface for that can handle valid read apps v1 namespaced deployment params
type ReadAppsV1NamespacedDeploymentHandler interface {
	Handle(ReadAppsV1NamespacedDeploymentParams, interface{}) middleware.Responder
}

// NewReadAppsV1NamespacedDeployment creates a new http.Handler for the read apps v1 namespaced deployment operation
func NewReadAppsV1NamespacedDeployment(ctx *middleware.Context, handler ReadAppsV1NamespacedDeploymentHandler) *ReadAppsV1NamespacedDeployment {
	return &ReadAppsV1NamespacedDeployment{Context: ctx, Handler: handler}
}

/*ReadAppsV1NamespacedDeployment swagger:route GET /apis/apps/v1/namespaces/{namespace}/deployments/{name} apps_v1 readAppsV1NamespacedDeployment

read the specified Deployment

*/
type ReadAppsV1NamespacedDeployment struct {
	Context *middleware.Context
	Handler ReadAppsV1NamespacedDeploymentHandler
}

func (o *ReadAppsV1NamespacedDeployment) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewReadAppsV1NamespacedDeploymentParams()

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
