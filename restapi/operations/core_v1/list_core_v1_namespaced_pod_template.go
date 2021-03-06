// Code generated by go-swagger; DO NOT EDIT.

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ListCoreV1NamespacedPodTemplateHandlerFunc turns a function with the right signature into a list core v1 namespaced pod template handler
type ListCoreV1NamespacedPodTemplateHandlerFunc func(ListCoreV1NamespacedPodTemplateParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ListCoreV1NamespacedPodTemplateHandlerFunc) Handle(params ListCoreV1NamespacedPodTemplateParams) middleware.Responder {
	return fn(params)
}

// ListCoreV1NamespacedPodTemplateHandler interface for that can handle valid list core v1 namespaced pod template params
type ListCoreV1NamespacedPodTemplateHandler interface {
	Handle(ListCoreV1NamespacedPodTemplateParams) middleware.Responder
}

// NewListCoreV1NamespacedPodTemplate creates a new http.Handler for the list core v1 namespaced pod template operation
func NewListCoreV1NamespacedPodTemplate(ctx *middleware.Context, handler ListCoreV1NamespacedPodTemplateHandler) *ListCoreV1NamespacedPodTemplate {
	return &ListCoreV1NamespacedPodTemplate{Context: ctx, Handler: handler}
}

/*ListCoreV1NamespacedPodTemplate swagger:route GET /api/v1/namespaces/{namespace}/podtemplates core_v1 listCoreV1NamespacedPodTemplate

list or watch objects of kind PodTemplate

*/
type ListCoreV1NamespacedPodTemplate struct {
	Context *middleware.Context
	Handler ListCoreV1NamespacedPodTemplateHandler
}

func (o *ListCoreV1NamespacedPodTemplate) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewListCoreV1NamespacedPodTemplateParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
