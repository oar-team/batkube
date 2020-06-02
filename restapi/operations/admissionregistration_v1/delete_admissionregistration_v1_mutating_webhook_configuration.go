// Code generated by go-swagger; DO NOT EDIT.

package admissionregistration_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteAdmissionregistrationV1MutatingWebhookConfigurationHandlerFunc turns a function with the right signature into a delete admissionregistration v1 mutating webhook configuration handler
type DeleteAdmissionregistrationV1MutatingWebhookConfigurationHandlerFunc func(DeleteAdmissionregistrationV1MutatingWebhookConfigurationParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteAdmissionregistrationV1MutatingWebhookConfigurationHandlerFunc) Handle(params DeleteAdmissionregistrationV1MutatingWebhookConfigurationParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// DeleteAdmissionregistrationV1MutatingWebhookConfigurationHandler interface for that can handle valid delete admissionregistration v1 mutating webhook configuration params
type DeleteAdmissionregistrationV1MutatingWebhookConfigurationHandler interface {
	Handle(DeleteAdmissionregistrationV1MutatingWebhookConfigurationParams, interface{}) middleware.Responder
}

// NewDeleteAdmissionregistrationV1MutatingWebhookConfiguration creates a new http.Handler for the delete admissionregistration v1 mutating webhook configuration operation
func NewDeleteAdmissionregistrationV1MutatingWebhookConfiguration(ctx *middleware.Context, handler DeleteAdmissionregistrationV1MutatingWebhookConfigurationHandler) *DeleteAdmissionregistrationV1MutatingWebhookConfiguration {
	return &DeleteAdmissionregistrationV1MutatingWebhookConfiguration{Context: ctx, Handler: handler}
}

/*DeleteAdmissionregistrationV1MutatingWebhookConfiguration swagger:route DELETE /apis/admissionregistration.k8s.io/v1/mutatingwebhookconfigurations/{name} admissionregistration_v1 deleteAdmissionregistrationV1MutatingWebhookConfiguration

delete a MutatingWebhookConfiguration

*/
type DeleteAdmissionregistrationV1MutatingWebhookConfiguration struct {
	Context *middleware.Context
	Handler DeleteAdmissionregistrationV1MutatingWebhookConfigurationHandler
}

func (o *DeleteAdmissionregistrationV1MutatingWebhookConfiguration) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteAdmissionregistrationV1MutatingWebhookConfigurationParams()

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
