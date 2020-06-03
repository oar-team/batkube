// Code generated by go-swagger; DO NOT EDIT.

package admissionregistration_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// CreateAdmissionregistrationV1beta1MutatingWebhookConfigurationHandlerFunc turns a function with the right signature into a create admissionregistration v1beta1 mutating webhook configuration handler
type CreateAdmissionregistrationV1beta1MutatingWebhookConfigurationHandlerFunc func(CreateAdmissionregistrationV1beta1MutatingWebhookConfigurationParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateAdmissionregistrationV1beta1MutatingWebhookConfigurationHandlerFunc) Handle(params CreateAdmissionregistrationV1beta1MutatingWebhookConfigurationParams) middleware.Responder {
	return fn(params)
}

// CreateAdmissionregistrationV1beta1MutatingWebhookConfigurationHandler interface for that can handle valid create admissionregistration v1beta1 mutating webhook configuration params
type CreateAdmissionregistrationV1beta1MutatingWebhookConfigurationHandler interface {
	Handle(CreateAdmissionregistrationV1beta1MutatingWebhookConfigurationParams) middleware.Responder
}

// NewCreateAdmissionregistrationV1beta1MutatingWebhookConfiguration creates a new http.Handler for the create admissionregistration v1beta1 mutating webhook configuration operation
func NewCreateAdmissionregistrationV1beta1MutatingWebhookConfiguration(ctx *middleware.Context, handler CreateAdmissionregistrationV1beta1MutatingWebhookConfigurationHandler) *CreateAdmissionregistrationV1beta1MutatingWebhookConfiguration {
	return &CreateAdmissionregistrationV1beta1MutatingWebhookConfiguration{Context: ctx, Handler: handler}
}

/*CreateAdmissionregistrationV1beta1MutatingWebhookConfiguration swagger:route POST /apis/admissionregistration.k8s.io/v1beta1/mutatingwebhookconfigurations admissionregistration_v1beta1 createAdmissionregistrationV1beta1MutatingWebhookConfiguration

create a MutatingWebhookConfiguration

*/
type CreateAdmissionregistrationV1beta1MutatingWebhookConfiguration struct {
	Context *middleware.Context
	Handler CreateAdmissionregistrationV1beta1MutatingWebhookConfigurationHandler
}

func (o *CreateAdmissionregistrationV1beta1MutatingWebhookConfiguration) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCreateAdmissionregistrationV1beta1MutatingWebhookConfigurationParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}