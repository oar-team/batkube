// Code generated by go-swagger; DO NOT EDIT.

package settings_v1alpha1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PatchSettingsV1alpha1NamespacedPodPresetHandlerFunc turns a function with the right signature into a patch settings v1alpha1 namespaced pod preset handler
type PatchSettingsV1alpha1NamespacedPodPresetHandlerFunc func(PatchSettingsV1alpha1NamespacedPodPresetParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn PatchSettingsV1alpha1NamespacedPodPresetHandlerFunc) Handle(params PatchSettingsV1alpha1NamespacedPodPresetParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// PatchSettingsV1alpha1NamespacedPodPresetHandler interface for that can handle valid patch settings v1alpha1 namespaced pod preset params
type PatchSettingsV1alpha1NamespacedPodPresetHandler interface {
	Handle(PatchSettingsV1alpha1NamespacedPodPresetParams, interface{}) middleware.Responder
}

// NewPatchSettingsV1alpha1NamespacedPodPreset creates a new http.Handler for the patch settings v1alpha1 namespaced pod preset operation
func NewPatchSettingsV1alpha1NamespacedPodPreset(ctx *middleware.Context, handler PatchSettingsV1alpha1NamespacedPodPresetHandler) *PatchSettingsV1alpha1NamespacedPodPreset {
	return &PatchSettingsV1alpha1NamespacedPodPreset{Context: ctx, Handler: handler}
}

/*PatchSettingsV1alpha1NamespacedPodPreset swagger:route PATCH /apis/settings.k8s.io/v1alpha1/namespaces/{namespace}/podpresets/{name} settings_v1alpha1 patchSettingsV1alpha1NamespacedPodPreset

partially update the specified PodPreset

*/
type PatchSettingsV1alpha1NamespacedPodPreset struct {
	Context *middleware.Context
	Handler PatchSettingsV1alpha1NamespacedPodPresetHandler
}

func (o *PatchSettingsV1alpha1NamespacedPodPreset) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPatchSettingsV1alpha1NamespacedPodPresetParams()

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
