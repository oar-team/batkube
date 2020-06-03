// Code generated by go-swagger; DO NOT EDIT.

package apiextensions_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ReplaceApiextensionsV1CustomResourceDefinitionStatusHandlerFunc turns a function with the right signature into a replace apiextensions v1 custom resource definition status handler
type ReplaceApiextensionsV1CustomResourceDefinitionStatusHandlerFunc func(ReplaceApiextensionsV1CustomResourceDefinitionStatusParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ReplaceApiextensionsV1CustomResourceDefinitionStatusHandlerFunc) Handle(params ReplaceApiextensionsV1CustomResourceDefinitionStatusParams) middleware.Responder {
	return fn(params)
}

// ReplaceApiextensionsV1CustomResourceDefinitionStatusHandler interface for that can handle valid replace apiextensions v1 custom resource definition status params
type ReplaceApiextensionsV1CustomResourceDefinitionStatusHandler interface {
	Handle(ReplaceApiextensionsV1CustomResourceDefinitionStatusParams) middleware.Responder
}

// NewReplaceApiextensionsV1CustomResourceDefinitionStatus creates a new http.Handler for the replace apiextensions v1 custom resource definition status operation
func NewReplaceApiextensionsV1CustomResourceDefinitionStatus(ctx *middleware.Context, handler ReplaceApiextensionsV1CustomResourceDefinitionStatusHandler) *ReplaceApiextensionsV1CustomResourceDefinitionStatus {
	return &ReplaceApiextensionsV1CustomResourceDefinitionStatus{Context: ctx, Handler: handler}
}

/*ReplaceApiextensionsV1CustomResourceDefinitionStatus swagger:route PUT /apis/apiextensions.k8s.io/v1/customresourcedefinitions/{name}/status apiextensions_v1 replaceApiextensionsV1CustomResourceDefinitionStatus

replace status of the specified CustomResourceDefinition

*/
type ReplaceApiextensionsV1CustomResourceDefinitionStatus struct {
	Context *middleware.Context
	Handler ReplaceApiextensionsV1CustomResourceDefinitionStatusHandler
}

func (o *ReplaceApiextensionsV1CustomResourceDefinitionStatus) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewReplaceApiextensionsV1CustomResourceDefinitionStatusParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}