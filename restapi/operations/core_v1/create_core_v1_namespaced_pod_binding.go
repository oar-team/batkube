// Code generated by go-swagger; DO NOT EDIT.

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// CreateCoreV1NamespacedPodBindingHandlerFunc turns a function with the right signature into a create core v1 namespaced pod binding handler
type CreateCoreV1NamespacedPodBindingHandlerFunc func(CreateCoreV1NamespacedPodBindingParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateCoreV1NamespacedPodBindingHandlerFunc) Handle(params CreateCoreV1NamespacedPodBindingParams) middleware.Responder {
	return fn(params)
}

// CreateCoreV1NamespacedPodBindingHandler interface for that can handle valid create core v1 namespaced pod binding params
type CreateCoreV1NamespacedPodBindingHandler interface {
	Handle(CreateCoreV1NamespacedPodBindingParams) middleware.Responder
}

// NewCreateCoreV1NamespacedPodBinding creates a new http.Handler for the create core v1 namespaced pod binding operation
func NewCreateCoreV1NamespacedPodBinding(ctx *middleware.Context, handler CreateCoreV1NamespacedPodBindingHandler) *CreateCoreV1NamespacedPodBinding {
	return &CreateCoreV1NamespacedPodBinding{Context: ctx, Handler: handler}
}

/*CreateCoreV1NamespacedPodBinding swagger:route POST /api/v1/namespaces/{namespace}/pods/{name}/binding core_v1 createCoreV1NamespacedPodBinding

create binding of a Pod

*/
type CreateCoreV1NamespacedPodBinding struct {
	Context *middleware.Context
	Handler CreateCoreV1NamespacedPodBindingHandler
}

func (o *CreateCoreV1NamespacedPodBinding) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCreateCoreV1NamespacedPodBindingParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
