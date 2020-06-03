// Code generated by go-swagger; DO NOT EDIT.

package settings_v1alpha1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// CreateSettingsV1alpha1NamespacedPodPresetHandlerFunc turns a function with the right signature into a create settings v1alpha1 namespaced pod preset handler
type CreateSettingsV1alpha1NamespacedPodPresetHandlerFunc func(CreateSettingsV1alpha1NamespacedPodPresetParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateSettingsV1alpha1NamespacedPodPresetHandlerFunc) Handle(params CreateSettingsV1alpha1NamespacedPodPresetParams) middleware.Responder {
	return fn(params)
}

// CreateSettingsV1alpha1NamespacedPodPresetHandler interface for that can handle valid create settings v1alpha1 namespaced pod preset params
type CreateSettingsV1alpha1NamespacedPodPresetHandler interface {
	Handle(CreateSettingsV1alpha1NamespacedPodPresetParams) middleware.Responder
}

// NewCreateSettingsV1alpha1NamespacedPodPreset creates a new http.Handler for the create settings v1alpha1 namespaced pod preset operation
func NewCreateSettingsV1alpha1NamespacedPodPreset(ctx *middleware.Context, handler CreateSettingsV1alpha1NamespacedPodPresetHandler) *CreateSettingsV1alpha1NamespacedPodPreset {
	return &CreateSettingsV1alpha1NamespacedPodPreset{Context: ctx, Handler: handler}
}

/*CreateSettingsV1alpha1NamespacedPodPreset swagger:route POST /apis/settings.k8s.io/v1alpha1/namespaces/{namespace}/podpresets settings_v1alpha1 createSettingsV1alpha1NamespacedPodPreset

create a PodPreset

*/
type CreateSettingsV1alpha1NamespacedPodPreset struct {
	Context *middleware.Context
	Handler CreateSettingsV1alpha1NamespacedPodPresetHandler
}

func (o *CreateSettingsV1alpha1NamespacedPodPreset) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCreateSettingsV1alpha1NamespacedPodPresetParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
