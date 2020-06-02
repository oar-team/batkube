// Code generated by go-swagger; DO NOT EDIT.

package settings_v1alpha1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteSettingsV1alpha1CollectionNamespacedPodPresetHandlerFunc turns a function with the right signature into a delete settings v1alpha1 collection namespaced pod preset handler
type DeleteSettingsV1alpha1CollectionNamespacedPodPresetHandlerFunc func(DeleteSettingsV1alpha1CollectionNamespacedPodPresetParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteSettingsV1alpha1CollectionNamespacedPodPresetHandlerFunc) Handle(params DeleteSettingsV1alpha1CollectionNamespacedPodPresetParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// DeleteSettingsV1alpha1CollectionNamespacedPodPresetHandler interface for that can handle valid delete settings v1alpha1 collection namespaced pod preset params
type DeleteSettingsV1alpha1CollectionNamespacedPodPresetHandler interface {
	Handle(DeleteSettingsV1alpha1CollectionNamespacedPodPresetParams, interface{}) middleware.Responder
}

// NewDeleteSettingsV1alpha1CollectionNamespacedPodPreset creates a new http.Handler for the delete settings v1alpha1 collection namespaced pod preset operation
func NewDeleteSettingsV1alpha1CollectionNamespacedPodPreset(ctx *middleware.Context, handler DeleteSettingsV1alpha1CollectionNamespacedPodPresetHandler) *DeleteSettingsV1alpha1CollectionNamespacedPodPreset {
	return &DeleteSettingsV1alpha1CollectionNamespacedPodPreset{Context: ctx, Handler: handler}
}

/*DeleteSettingsV1alpha1CollectionNamespacedPodPreset swagger:route DELETE /apis/settings.k8s.io/v1alpha1/namespaces/{namespace}/podpresets settings_v1alpha1 deleteSettingsV1alpha1CollectionNamespacedPodPreset

delete collection of PodPreset

*/
type DeleteSettingsV1alpha1CollectionNamespacedPodPreset struct {
	Context *middleware.Context
	Handler DeleteSettingsV1alpha1CollectionNamespacedPodPresetHandler
}

func (o *DeleteSettingsV1alpha1CollectionNamespacedPodPreset) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteSettingsV1alpha1CollectionNamespacedPodPresetParams()

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
