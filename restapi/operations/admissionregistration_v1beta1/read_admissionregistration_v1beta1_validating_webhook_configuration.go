// Code generated by go-swagger; DO NOT EDIT.

package admissionregistration_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ReadAdmissionregistrationV1beta1ValidatingWebhookConfigurationHandlerFunc turns a function with the right signature into a read admissionregistration v1beta1 validating webhook configuration handler
type ReadAdmissionregistrationV1beta1ValidatingWebhookConfigurationHandlerFunc func(ReadAdmissionregistrationV1beta1ValidatingWebhookConfigurationParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ReadAdmissionregistrationV1beta1ValidatingWebhookConfigurationHandlerFunc) Handle(params ReadAdmissionregistrationV1beta1ValidatingWebhookConfigurationParams) middleware.Responder {
	return fn(params)
}

// ReadAdmissionregistrationV1beta1ValidatingWebhookConfigurationHandler interface for that can handle valid read admissionregistration v1beta1 validating webhook configuration params
type ReadAdmissionregistrationV1beta1ValidatingWebhookConfigurationHandler interface {
	Handle(ReadAdmissionregistrationV1beta1ValidatingWebhookConfigurationParams) middleware.Responder
}

// NewReadAdmissionregistrationV1beta1ValidatingWebhookConfiguration creates a new http.Handler for the read admissionregistration v1beta1 validating webhook configuration operation
func NewReadAdmissionregistrationV1beta1ValidatingWebhookConfiguration(ctx *middleware.Context, handler ReadAdmissionregistrationV1beta1ValidatingWebhookConfigurationHandler) *ReadAdmissionregistrationV1beta1ValidatingWebhookConfiguration {
	return &ReadAdmissionregistrationV1beta1ValidatingWebhookConfiguration{Context: ctx, Handler: handler}
}

/*ReadAdmissionregistrationV1beta1ValidatingWebhookConfiguration swagger:route GET /apis/admissionregistration.k8s.io/v1beta1/validatingwebhookconfigurations/{name} admissionregistration_v1beta1 readAdmissionregistrationV1beta1ValidatingWebhookConfiguration

read the specified ValidatingWebhookConfiguration

*/
type ReadAdmissionregistrationV1beta1ValidatingWebhookConfiguration struct {
	Context *middleware.Context
	Handler ReadAdmissionregistrationV1beta1ValidatingWebhookConfigurationHandler
}

func (o *ReadAdmissionregistrationV1beta1ValidatingWebhookConfiguration) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewReadAdmissionregistrationV1beta1ValidatingWebhookConfigurationParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
