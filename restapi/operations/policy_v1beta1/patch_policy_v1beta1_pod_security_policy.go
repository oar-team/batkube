// Code generated by go-swagger; DO NOT EDIT.

package policy_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PatchPolicyV1beta1PodSecurityPolicyHandlerFunc turns a function with the right signature into a patch policy v1beta1 pod security policy handler
type PatchPolicyV1beta1PodSecurityPolicyHandlerFunc func(PatchPolicyV1beta1PodSecurityPolicyParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PatchPolicyV1beta1PodSecurityPolicyHandlerFunc) Handle(params PatchPolicyV1beta1PodSecurityPolicyParams) middleware.Responder {
	return fn(params)
}

// PatchPolicyV1beta1PodSecurityPolicyHandler interface for that can handle valid patch policy v1beta1 pod security policy params
type PatchPolicyV1beta1PodSecurityPolicyHandler interface {
	Handle(PatchPolicyV1beta1PodSecurityPolicyParams) middleware.Responder
}

// NewPatchPolicyV1beta1PodSecurityPolicy creates a new http.Handler for the patch policy v1beta1 pod security policy operation
func NewPatchPolicyV1beta1PodSecurityPolicy(ctx *middleware.Context, handler PatchPolicyV1beta1PodSecurityPolicyHandler) *PatchPolicyV1beta1PodSecurityPolicy {
	return &PatchPolicyV1beta1PodSecurityPolicy{Context: ctx, Handler: handler}
}

/*PatchPolicyV1beta1PodSecurityPolicy swagger:route PATCH /apis/policy/v1beta1/podsecuritypolicies/{name} policy_v1beta1 patchPolicyV1beta1PodSecurityPolicy

partially update the specified PodSecurityPolicy

*/
type PatchPolicyV1beta1PodSecurityPolicy struct {
	Context *middleware.Context
	Handler PatchPolicyV1beta1PodSecurityPolicyHandler
}

func (o *PatchPolicyV1beta1PodSecurityPolicy) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPatchPolicyV1beta1PodSecurityPolicyParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}