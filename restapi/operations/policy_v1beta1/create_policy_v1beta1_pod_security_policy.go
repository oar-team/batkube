// Code generated by go-swagger; DO NOT EDIT.

package policy_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// CreatePolicyV1beta1PodSecurityPolicyHandlerFunc turns a function with the right signature into a create policy v1beta1 pod security policy handler
type CreatePolicyV1beta1PodSecurityPolicyHandlerFunc func(CreatePolicyV1beta1PodSecurityPolicyParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn CreatePolicyV1beta1PodSecurityPolicyHandlerFunc) Handle(params CreatePolicyV1beta1PodSecurityPolicyParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// CreatePolicyV1beta1PodSecurityPolicyHandler interface for that can handle valid create policy v1beta1 pod security policy params
type CreatePolicyV1beta1PodSecurityPolicyHandler interface {
	Handle(CreatePolicyV1beta1PodSecurityPolicyParams, interface{}) middleware.Responder
}

// NewCreatePolicyV1beta1PodSecurityPolicy creates a new http.Handler for the create policy v1beta1 pod security policy operation
func NewCreatePolicyV1beta1PodSecurityPolicy(ctx *middleware.Context, handler CreatePolicyV1beta1PodSecurityPolicyHandler) *CreatePolicyV1beta1PodSecurityPolicy {
	return &CreatePolicyV1beta1PodSecurityPolicy{Context: ctx, Handler: handler}
}

/*CreatePolicyV1beta1PodSecurityPolicy swagger:route POST /apis/policy/v1beta1/podsecuritypolicies policy_v1beta1 createPolicyV1beta1PodSecurityPolicy

create a PodSecurityPolicy

*/
type CreatePolicyV1beta1PodSecurityPolicy struct {
	Context *middleware.Context
	Handler CreatePolicyV1beta1PodSecurityPolicyHandler
}

func (o *CreatePolicyV1beta1PodSecurityPolicy) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCreatePolicyV1beta1PodSecurityPolicyParams()

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
