// Code generated by go-swagger; DO NOT EDIT.

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PatchCoreV1NamespacedResourceQuotaStatusHandlerFunc turns a function with the right signature into a patch core v1 namespaced resource quota status handler
type PatchCoreV1NamespacedResourceQuotaStatusHandlerFunc func(PatchCoreV1NamespacedResourceQuotaStatusParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PatchCoreV1NamespacedResourceQuotaStatusHandlerFunc) Handle(params PatchCoreV1NamespacedResourceQuotaStatusParams) middleware.Responder {
	return fn(params)
}

// PatchCoreV1NamespacedResourceQuotaStatusHandler interface for that can handle valid patch core v1 namespaced resource quota status params
type PatchCoreV1NamespacedResourceQuotaStatusHandler interface {
	Handle(PatchCoreV1NamespacedResourceQuotaStatusParams) middleware.Responder
}

// NewPatchCoreV1NamespacedResourceQuotaStatus creates a new http.Handler for the patch core v1 namespaced resource quota status operation
func NewPatchCoreV1NamespacedResourceQuotaStatus(ctx *middleware.Context, handler PatchCoreV1NamespacedResourceQuotaStatusHandler) *PatchCoreV1NamespacedResourceQuotaStatus {
	return &PatchCoreV1NamespacedResourceQuotaStatus{Context: ctx, Handler: handler}
}

/*PatchCoreV1NamespacedResourceQuotaStatus swagger:route PATCH /api/v1/namespaces/{namespace}/resourcequotas/{name}/status core_v1 patchCoreV1NamespacedResourceQuotaStatus

partially update status of the specified ResourceQuota

*/
type PatchCoreV1NamespacedResourceQuotaStatus struct {
	Context *middleware.Context
	Handler PatchCoreV1NamespacedResourceQuotaStatusHandler
}

func (o *PatchCoreV1NamespacedResourceQuotaStatus) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPatchCoreV1NamespacedResourceQuotaStatusParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
