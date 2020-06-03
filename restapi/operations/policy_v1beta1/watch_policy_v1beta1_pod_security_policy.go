// Code generated by go-swagger; DO NOT EDIT.

package policy_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// WatchPolicyV1beta1PodSecurityPolicyHandlerFunc turns a function with the right signature into a watch policy v1beta1 pod security policy handler
type WatchPolicyV1beta1PodSecurityPolicyHandlerFunc func(WatchPolicyV1beta1PodSecurityPolicyParams) middleware.Responder

// Handle executing the request and returning a response
func (fn WatchPolicyV1beta1PodSecurityPolicyHandlerFunc) Handle(params WatchPolicyV1beta1PodSecurityPolicyParams) middleware.Responder {
	return fn(params)
}

// WatchPolicyV1beta1PodSecurityPolicyHandler interface for that can handle valid watch policy v1beta1 pod security policy params
type WatchPolicyV1beta1PodSecurityPolicyHandler interface {
	Handle(WatchPolicyV1beta1PodSecurityPolicyParams) middleware.Responder
}

// NewWatchPolicyV1beta1PodSecurityPolicy creates a new http.Handler for the watch policy v1beta1 pod security policy operation
func NewWatchPolicyV1beta1PodSecurityPolicy(ctx *middleware.Context, handler WatchPolicyV1beta1PodSecurityPolicyHandler) *WatchPolicyV1beta1PodSecurityPolicy {
	return &WatchPolicyV1beta1PodSecurityPolicy{Context: ctx, Handler: handler}
}

/*WatchPolicyV1beta1PodSecurityPolicy swagger:route GET /apis/policy/v1beta1/watch/podsecuritypolicies/{name} policy_v1beta1 watchPolicyV1beta1PodSecurityPolicy

watch changes to an object of kind PodSecurityPolicy. deprecated: use the 'watch' parameter with a list operation instead, filtered to a single item with the 'fieldSelector' parameter.

*/
type WatchPolicyV1beta1PodSecurityPolicy struct {
	Context *middleware.Context
	Handler WatchPolicyV1beta1PodSecurityPolicyHandler
}

func (o *WatchPolicyV1beta1PodSecurityPolicy) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewWatchPolicyV1beta1PodSecurityPolicyParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}