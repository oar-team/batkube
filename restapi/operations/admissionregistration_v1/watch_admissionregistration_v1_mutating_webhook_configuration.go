// Code generated by go-swagger; DO NOT EDIT.

package admissionregistration_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// WatchAdmissionregistrationV1MutatingWebhookConfigurationHandlerFunc turns a function with the right signature into a watch admissionregistration v1 mutating webhook configuration handler
type WatchAdmissionregistrationV1MutatingWebhookConfigurationHandlerFunc func(WatchAdmissionregistrationV1MutatingWebhookConfigurationParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn WatchAdmissionregistrationV1MutatingWebhookConfigurationHandlerFunc) Handle(params WatchAdmissionregistrationV1MutatingWebhookConfigurationParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// WatchAdmissionregistrationV1MutatingWebhookConfigurationHandler interface for that can handle valid watch admissionregistration v1 mutating webhook configuration params
type WatchAdmissionregistrationV1MutatingWebhookConfigurationHandler interface {
	Handle(WatchAdmissionregistrationV1MutatingWebhookConfigurationParams, interface{}) middleware.Responder
}

// NewWatchAdmissionregistrationV1MutatingWebhookConfiguration creates a new http.Handler for the watch admissionregistration v1 mutating webhook configuration operation
func NewWatchAdmissionregistrationV1MutatingWebhookConfiguration(ctx *middleware.Context, handler WatchAdmissionregistrationV1MutatingWebhookConfigurationHandler) *WatchAdmissionregistrationV1MutatingWebhookConfiguration {
	return &WatchAdmissionregistrationV1MutatingWebhookConfiguration{Context: ctx, Handler: handler}
}

/*WatchAdmissionregistrationV1MutatingWebhookConfiguration swagger:route GET /apis/admissionregistration.k8s.io/v1/watch/mutatingwebhookconfigurations/{name} admissionregistration_v1 watchAdmissionregistrationV1MutatingWebhookConfiguration

watch changes to an object of kind MutatingWebhookConfiguration. deprecated: use the 'watch' parameter with a list operation instead, filtered to a single item with the 'fieldSelector' parameter.

*/
type WatchAdmissionregistrationV1MutatingWebhookConfiguration struct {
	Context *middleware.Context
	Handler WatchAdmissionregistrationV1MutatingWebhookConfigurationHandler
}

func (o *WatchAdmissionregistrationV1MutatingWebhookConfiguration) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewWatchAdmissionregistrationV1MutatingWebhookConfigurationParams()

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
