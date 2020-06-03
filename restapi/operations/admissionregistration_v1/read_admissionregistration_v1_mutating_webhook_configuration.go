// Code generated by go-swagger; DO NOT EDIT.

package admissionregistration_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ReadAdmissionregistrationV1MutatingWebhookConfigurationHandlerFunc turns a function with the right signature into a read admissionregistration v1 mutating webhook configuration handler
type ReadAdmissionregistrationV1MutatingWebhookConfigurationHandlerFunc func(ReadAdmissionregistrationV1MutatingWebhookConfigurationParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ReadAdmissionregistrationV1MutatingWebhookConfigurationHandlerFunc) Handle(params ReadAdmissionregistrationV1MutatingWebhookConfigurationParams) middleware.Responder {
	return fn(params)
}

// ReadAdmissionregistrationV1MutatingWebhookConfigurationHandler interface for that can handle valid read admissionregistration v1 mutating webhook configuration params
type ReadAdmissionregistrationV1MutatingWebhookConfigurationHandler interface {
	Handle(ReadAdmissionregistrationV1MutatingWebhookConfigurationParams) middleware.Responder
}

// NewReadAdmissionregistrationV1MutatingWebhookConfiguration creates a new http.Handler for the read admissionregistration v1 mutating webhook configuration operation
func NewReadAdmissionregistrationV1MutatingWebhookConfiguration(ctx *middleware.Context, handler ReadAdmissionregistrationV1MutatingWebhookConfigurationHandler) *ReadAdmissionregistrationV1MutatingWebhookConfiguration {
	return &ReadAdmissionregistrationV1MutatingWebhookConfiguration{Context: ctx, Handler: handler}
}

/*ReadAdmissionregistrationV1MutatingWebhookConfiguration swagger:route GET /apis/admissionregistration.k8s.io/v1/mutatingwebhookconfigurations/{name} admissionregistration_v1 readAdmissionregistrationV1MutatingWebhookConfiguration

read the specified MutatingWebhookConfiguration

*/
type ReadAdmissionregistrationV1MutatingWebhookConfiguration struct {
	Context *middleware.Context
	Handler ReadAdmissionregistrationV1MutatingWebhookConfigurationHandler
}

func (o *ReadAdmissionregistrationV1MutatingWebhookConfiguration) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewReadAdmissionregistrationV1MutatingWebhookConfigurationParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
