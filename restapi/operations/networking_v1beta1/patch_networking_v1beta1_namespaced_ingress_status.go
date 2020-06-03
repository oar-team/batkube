// Code generated by go-swagger; DO NOT EDIT.

package networking_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PatchNetworkingV1beta1NamespacedIngressStatusHandlerFunc turns a function with the right signature into a patch networking v1beta1 namespaced ingress status handler
type PatchNetworkingV1beta1NamespacedIngressStatusHandlerFunc func(PatchNetworkingV1beta1NamespacedIngressStatusParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PatchNetworkingV1beta1NamespacedIngressStatusHandlerFunc) Handle(params PatchNetworkingV1beta1NamespacedIngressStatusParams) middleware.Responder {
	return fn(params)
}

// PatchNetworkingV1beta1NamespacedIngressStatusHandler interface for that can handle valid patch networking v1beta1 namespaced ingress status params
type PatchNetworkingV1beta1NamespacedIngressStatusHandler interface {
	Handle(PatchNetworkingV1beta1NamespacedIngressStatusParams) middleware.Responder
}

// NewPatchNetworkingV1beta1NamespacedIngressStatus creates a new http.Handler for the patch networking v1beta1 namespaced ingress status operation
func NewPatchNetworkingV1beta1NamespacedIngressStatus(ctx *middleware.Context, handler PatchNetworkingV1beta1NamespacedIngressStatusHandler) *PatchNetworkingV1beta1NamespacedIngressStatus {
	return &PatchNetworkingV1beta1NamespacedIngressStatus{Context: ctx, Handler: handler}
}

/*PatchNetworkingV1beta1NamespacedIngressStatus swagger:route PATCH /apis/networking.k8s.io/v1beta1/namespaces/{namespace}/ingresses/{name}/status networking_v1beta1 patchNetworkingV1beta1NamespacedIngressStatus

partially update status of the specified Ingress

*/
type PatchNetworkingV1beta1NamespacedIngressStatus struct {
	Context *middleware.Context
	Handler PatchNetworkingV1beta1NamespacedIngressStatusHandler
}

func (o *PatchNetworkingV1beta1NamespacedIngressStatus) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPatchNetworkingV1beta1NamespacedIngressStatusParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
