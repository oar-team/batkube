// Code generated by go-swagger; DO NOT EDIT.

package admissionregistration_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ReplaceAdmissionregistrationV1MutatingWebhookConfigurationHandlerFunc turns a function with the right signature into a replace admissionregistration v1 mutating webhook configuration handler
type ReplaceAdmissionregistrationV1MutatingWebhookConfigurationHandlerFunc func(ReplaceAdmissionregistrationV1MutatingWebhookConfigurationParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ReplaceAdmissionregistrationV1MutatingWebhookConfigurationHandlerFunc) Handle(params ReplaceAdmissionregistrationV1MutatingWebhookConfigurationParams) middleware.Responder {
	return fn(params)
}

// ReplaceAdmissionregistrationV1MutatingWebhookConfigurationHandler interface for that can handle valid replace admissionregistration v1 mutating webhook configuration params
type ReplaceAdmissionregistrationV1MutatingWebhookConfigurationHandler interface {
	Handle(ReplaceAdmissionregistrationV1MutatingWebhookConfigurationParams) middleware.Responder
}

// NewReplaceAdmissionregistrationV1MutatingWebhookConfiguration creates a new http.Handler for the replace admissionregistration v1 mutating webhook configuration operation
func NewReplaceAdmissionregistrationV1MutatingWebhookConfiguration(ctx *middleware.Context, handler ReplaceAdmissionregistrationV1MutatingWebhookConfigurationHandler) *ReplaceAdmissionregistrationV1MutatingWebhookConfiguration {
	return &ReplaceAdmissionregistrationV1MutatingWebhookConfiguration{Context: ctx, Handler: handler}
}

/*ReplaceAdmissionregistrationV1MutatingWebhookConfiguration swagger:route PUT /apis/admissionregistration.k8s.io/v1/mutatingwebhookconfigurations/{name} admissionregistration_v1 replaceAdmissionregistrationV1MutatingWebhookConfiguration

replace the specified MutatingWebhookConfiguration

*/
type ReplaceAdmissionregistrationV1MutatingWebhookConfiguration struct {
	Context *middleware.Context
	Handler ReplaceAdmissionregistrationV1MutatingWebhookConfigurationHandler
}

func (o *ReplaceAdmissionregistrationV1MutatingWebhookConfiguration) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewReplaceAdmissionregistrationV1MutatingWebhookConfigurationParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
