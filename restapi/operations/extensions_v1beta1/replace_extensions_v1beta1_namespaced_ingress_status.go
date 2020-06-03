// Code generated by go-swagger; DO NOT EDIT.

package extensions_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ReplaceExtensionsV1beta1NamespacedIngressStatusHandlerFunc turns a function with the right signature into a replace extensions v1beta1 namespaced ingress status handler
type ReplaceExtensionsV1beta1NamespacedIngressStatusHandlerFunc func(ReplaceExtensionsV1beta1NamespacedIngressStatusParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ReplaceExtensionsV1beta1NamespacedIngressStatusHandlerFunc) Handle(params ReplaceExtensionsV1beta1NamespacedIngressStatusParams) middleware.Responder {
	return fn(params)
}

// ReplaceExtensionsV1beta1NamespacedIngressStatusHandler interface for that can handle valid replace extensions v1beta1 namespaced ingress status params
type ReplaceExtensionsV1beta1NamespacedIngressStatusHandler interface {
	Handle(ReplaceExtensionsV1beta1NamespacedIngressStatusParams) middleware.Responder
}

// NewReplaceExtensionsV1beta1NamespacedIngressStatus creates a new http.Handler for the replace extensions v1beta1 namespaced ingress status operation
func NewReplaceExtensionsV1beta1NamespacedIngressStatus(ctx *middleware.Context, handler ReplaceExtensionsV1beta1NamespacedIngressStatusHandler) *ReplaceExtensionsV1beta1NamespacedIngressStatus {
	return &ReplaceExtensionsV1beta1NamespacedIngressStatus{Context: ctx, Handler: handler}
}

/*ReplaceExtensionsV1beta1NamespacedIngressStatus swagger:route PUT /apis/extensions/v1beta1/namespaces/{namespace}/ingresses/{name}/status extensions_v1beta1 replaceExtensionsV1beta1NamespacedIngressStatus

replace status of the specified Ingress

*/
type ReplaceExtensionsV1beta1NamespacedIngressStatus struct {
	Context *middleware.Context
	Handler ReplaceExtensionsV1beta1NamespacedIngressStatusHandler
}

func (o *ReplaceExtensionsV1beta1NamespacedIngressStatus) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewReplaceExtensionsV1beta1NamespacedIngressStatusParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
