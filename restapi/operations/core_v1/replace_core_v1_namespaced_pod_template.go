// Code generated by go-swagger; DO NOT EDIT.

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ReplaceCoreV1NamespacedPodTemplateHandlerFunc turns a function with the right signature into a replace core v1 namespaced pod template handler
type ReplaceCoreV1NamespacedPodTemplateHandlerFunc func(ReplaceCoreV1NamespacedPodTemplateParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn ReplaceCoreV1NamespacedPodTemplateHandlerFunc) Handle(params ReplaceCoreV1NamespacedPodTemplateParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// ReplaceCoreV1NamespacedPodTemplateHandler interface for that can handle valid replace core v1 namespaced pod template params
type ReplaceCoreV1NamespacedPodTemplateHandler interface {
	Handle(ReplaceCoreV1NamespacedPodTemplateParams, interface{}) middleware.Responder
}

// NewReplaceCoreV1NamespacedPodTemplate creates a new http.Handler for the replace core v1 namespaced pod template operation
func NewReplaceCoreV1NamespacedPodTemplate(ctx *middleware.Context, handler ReplaceCoreV1NamespacedPodTemplateHandler) *ReplaceCoreV1NamespacedPodTemplate {
	return &ReplaceCoreV1NamespacedPodTemplate{Context: ctx, Handler: handler}
}

/*ReplaceCoreV1NamespacedPodTemplate swagger:route PUT /api/v1/namespaces/{namespace}/podtemplates/{name} core_v1 replaceCoreV1NamespacedPodTemplate

replace the specified PodTemplate

*/
type ReplaceCoreV1NamespacedPodTemplate struct {
	Context *middleware.Context
	Handler ReplaceCoreV1NamespacedPodTemplateHandler
}

func (o *ReplaceCoreV1NamespacedPodTemplate) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewReplaceCoreV1NamespacedPodTemplateParams()

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
