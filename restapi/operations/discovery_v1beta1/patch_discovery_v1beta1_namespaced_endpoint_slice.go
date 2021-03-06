// Code generated by go-swagger; DO NOT EDIT.

package discovery_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PatchDiscoveryV1beta1NamespacedEndpointSliceHandlerFunc turns a function with the right signature into a patch discovery v1beta1 namespaced endpoint slice handler
type PatchDiscoveryV1beta1NamespacedEndpointSliceHandlerFunc func(PatchDiscoveryV1beta1NamespacedEndpointSliceParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PatchDiscoveryV1beta1NamespacedEndpointSliceHandlerFunc) Handle(params PatchDiscoveryV1beta1NamespacedEndpointSliceParams) middleware.Responder {
	return fn(params)
}

// PatchDiscoveryV1beta1NamespacedEndpointSliceHandler interface for that can handle valid patch discovery v1beta1 namespaced endpoint slice params
type PatchDiscoveryV1beta1NamespacedEndpointSliceHandler interface {
	Handle(PatchDiscoveryV1beta1NamespacedEndpointSliceParams) middleware.Responder
}

// NewPatchDiscoveryV1beta1NamespacedEndpointSlice creates a new http.Handler for the patch discovery v1beta1 namespaced endpoint slice operation
func NewPatchDiscoveryV1beta1NamespacedEndpointSlice(ctx *middleware.Context, handler PatchDiscoveryV1beta1NamespacedEndpointSliceHandler) *PatchDiscoveryV1beta1NamespacedEndpointSlice {
	return &PatchDiscoveryV1beta1NamespacedEndpointSlice{Context: ctx, Handler: handler}
}

/*PatchDiscoveryV1beta1NamespacedEndpointSlice swagger:route PATCH /apis/discovery.k8s.io/v1beta1/namespaces/{namespace}/endpointslices/{name} discovery_v1beta1 patchDiscoveryV1beta1NamespacedEndpointSlice

partially update the specified EndpointSlice

*/
type PatchDiscoveryV1beta1NamespacedEndpointSlice struct {
	Context *middleware.Context
	Handler PatchDiscoveryV1beta1NamespacedEndpointSliceHandler
}

func (o *PatchDiscoveryV1beta1NamespacedEndpointSlice) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPatchDiscoveryV1beta1NamespacedEndpointSliceParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
