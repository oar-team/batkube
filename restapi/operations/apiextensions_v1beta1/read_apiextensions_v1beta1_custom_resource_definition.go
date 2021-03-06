// Code generated by go-swagger; DO NOT EDIT.

package apiextensions_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ReadApiextensionsV1beta1CustomResourceDefinitionHandlerFunc turns a function with the right signature into a read apiextensions v1beta1 custom resource definition handler
type ReadApiextensionsV1beta1CustomResourceDefinitionHandlerFunc func(ReadApiextensionsV1beta1CustomResourceDefinitionParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ReadApiextensionsV1beta1CustomResourceDefinitionHandlerFunc) Handle(params ReadApiextensionsV1beta1CustomResourceDefinitionParams) middleware.Responder {
	return fn(params)
}

// ReadApiextensionsV1beta1CustomResourceDefinitionHandler interface for that can handle valid read apiextensions v1beta1 custom resource definition params
type ReadApiextensionsV1beta1CustomResourceDefinitionHandler interface {
	Handle(ReadApiextensionsV1beta1CustomResourceDefinitionParams) middleware.Responder
}

// NewReadApiextensionsV1beta1CustomResourceDefinition creates a new http.Handler for the read apiextensions v1beta1 custom resource definition operation
func NewReadApiextensionsV1beta1CustomResourceDefinition(ctx *middleware.Context, handler ReadApiextensionsV1beta1CustomResourceDefinitionHandler) *ReadApiextensionsV1beta1CustomResourceDefinition {
	return &ReadApiextensionsV1beta1CustomResourceDefinition{Context: ctx, Handler: handler}
}

/*ReadApiextensionsV1beta1CustomResourceDefinition swagger:route GET /apis/apiextensions.k8s.io/v1beta1/customresourcedefinitions/{name} apiextensions_v1beta1 readApiextensionsV1beta1CustomResourceDefinition

read the specified CustomResourceDefinition

*/
type ReadApiextensionsV1beta1CustomResourceDefinition struct {
	Context *middleware.Context
	Handler ReadApiextensionsV1beta1CustomResourceDefinitionHandler
}

func (o *ReadApiextensionsV1beta1CustomResourceDefinition) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewReadApiextensionsV1beta1CustomResourceDefinitionParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
