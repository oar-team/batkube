// Code generated by go-swagger; DO NOT EDIT.

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// CreateCoreV1NamespacedPodTemplateHandlerFunc turns a function with the right signature into a create core v1 namespaced pod template handler
type CreateCoreV1NamespacedPodTemplateHandlerFunc func(CreateCoreV1NamespacedPodTemplateParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateCoreV1NamespacedPodTemplateHandlerFunc) Handle(params CreateCoreV1NamespacedPodTemplateParams) middleware.Responder {
	return fn(params)
}

// CreateCoreV1NamespacedPodTemplateHandler interface for that can handle valid create core v1 namespaced pod template params
type CreateCoreV1NamespacedPodTemplateHandler interface {
	Handle(CreateCoreV1NamespacedPodTemplateParams) middleware.Responder
}

// NewCreateCoreV1NamespacedPodTemplate creates a new http.Handler for the create core v1 namespaced pod template operation
func NewCreateCoreV1NamespacedPodTemplate(ctx *middleware.Context, handler CreateCoreV1NamespacedPodTemplateHandler) *CreateCoreV1NamespacedPodTemplate {
	return &CreateCoreV1NamespacedPodTemplate{Context: ctx, Handler: handler}
}

/*CreateCoreV1NamespacedPodTemplate swagger:route POST /api/v1/namespaces/{namespace}/podtemplates core_v1 createCoreV1NamespacedPodTemplate

create a PodTemplate

*/
type CreateCoreV1NamespacedPodTemplate struct {
	Context *middleware.Context
	Handler CreateCoreV1NamespacedPodTemplateHandler
}

func (o *CreateCoreV1NamespacedPodTemplate) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCreateCoreV1NamespacedPodTemplateParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
