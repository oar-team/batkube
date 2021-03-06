// Code generated by go-swagger; DO NOT EDIT.

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ReadCoreV1NamespacedPodTemplateHandlerFunc turns a function with the right signature into a read core v1 namespaced pod template handler
type ReadCoreV1NamespacedPodTemplateHandlerFunc func(ReadCoreV1NamespacedPodTemplateParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ReadCoreV1NamespacedPodTemplateHandlerFunc) Handle(params ReadCoreV1NamespacedPodTemplateParams) middleware.Responder {
	return fn(params)
}

// ReadCoreV1NamespacedPodTemplateHandler interface for that can handle valid read core v1 namespaced pod template params
type ReadCoreV1NamespacedPodTemplateHandler interface {
	Handle(ReadCoreV1NamespacedPodTemplateParams) middleware.Responder
}

// NewReadCoreV1NamespacedPodTemplate creates a new http.Handler for the read core v1 namespaced pod template operation
func NewReadCoreV1NamespacedPodTemplate(ctx *middleware.Context, handler ReadCoreV1NamespacedPodTemplateHandler) *ReadCoreV1NamespacedPodTemplate {
	return &ReadCoreV1NamespacedPodTemplate{Context: ctx, Handler: handler}
}

/*ReadCoreV1NamespacedPodTemplate swagger:route GET /api/v1/namespaces/{namespace}/podtemplates/{name} core_v1 readCoreV1NamespacedPodTemplate

read the specified PodTemplate

*/
type ReadCoreV1NamespacedPodTemplate struct {
	Context *middleware.Context
	Handler ReadCoreV1NamespacedPodTemplateHandler
}

func (o *ReadCoreV1NamespacedPodTemplate) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewReadCoreV1NamespacedPodTemplateParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
