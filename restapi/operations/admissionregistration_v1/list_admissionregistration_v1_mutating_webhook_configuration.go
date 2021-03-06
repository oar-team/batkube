// Code generated by go-swagger; DO NOT EDIT.

package admissionregistration_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ListAdmissionregistrationV1MutatingWebhookConfigurationHandlerFunc turns a function with the right signature into a list admissionregistration v1 mutating webhook configuration handler
type ListAdmissionregistrationV1MutatingWebhookConfigurationHandlerFunc func(ListAdmissionregistrationV1MutatingWebhookConfigurationParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ListAdmissionregistrationV1MutatingWebhookConfigurationHandlerFunc) Handle(params ListAdmissionregistrationV1MutatingWebhookConfigurationParams) middleware.Responder {
	return fn(params)
}

// ListAdmissionregistrationV1MutatingWebhookConfigurationHandler interface for that can handle valid list admissionregistration v1 mutating webhook configuration params
type ListAdmissionregistrationV1MutatingWebhookConfigurationHandler interface {
	Handle(ListAdmissionregistrationV1MutatingWebhookConfigurationParams) middleware.Responder
}

// NewListAdmissionregistrationV1MutatingWebhookConfiguration creates a new http.Handler for the list admissionregistration v1 mutating webhook configuration operation
func NewListAdmissionregistrationV1MutatingWebhookConfiguration(ctx *middleware.Context, handler ListAdmissionregistrationV1MutatingWebhookConfigurationHandler) *ListAdmissionregistrationV1MutatingWebhookConfiguration {
	return &ListAdmissionregistrationV1MutatingWebhookConfiguration{Context: ctx, Handler: handler}
}

/*ListAdmissionregistrationV1MutatingWebhookConfiguration swagger:route GET /apis/admissionregistration.k8s.io/v1/mutatingwebhookconfigurations admissionregistration_v1 listAdmissionregistrationV1MutatingWebhookConfiguration

list or watch objects of kind MutatingWebhookConfiguration

*/
type ListAdmissionregistrationV1MutatingWebhookConfiguration struct {
	Context *middleware.Context
	Handler ListAdmissionregistrationV1MutatingWebhookConfigurationHandler
}

func (o *ListAdmissionregistrationV1MutatingWebhookConfiguration) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewListAdmissionregistrationV1MutatingWebhookConfigurationParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
