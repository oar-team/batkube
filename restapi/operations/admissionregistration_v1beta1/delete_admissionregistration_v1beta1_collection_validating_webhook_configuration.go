// Code generated by go-swagger; DO NOT EDIT.

package admissionregistration_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteAdmissionregistrationV1beta1CollectionValidatingWebhookConfigurationHandlerFunc turns a function with the right signature into a delete admissionregistration v1beta1 collection validating webhook configuration handler
type DeleteAdmissionregistrationV1beta1CollectionValidatingWebhookConfigurationHandlerFunc func(DeleteAdmissionregistrationV1beta1CollectionValidatingWebhookConfigurationParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteAdmissionregistrationV1beta1CollectionValidatingWebhookConfigurationHandlerFunc) Handle(params DeleteAdmissionregistrationV1beta1CollectionValidatingWebhookConfigurationParams) middleware.Responder {
	return fn(params)
}

// DeleteAdmissionregistrationV1beta1CollectionValidatingWebhookConfigurationHandler interface for that can handle valid delete admissionregistration v1beta1 collection validating webhook configuration params
type DeleteAdmissionregistrationV1beta1CollectionValidatingWebhookConfigurationHandler interface {
	Handle(DeleteAdmissionregistrationV1beta1CollectionValidatingWebhookConfigurationParams) middleware.Responder
}

// NewDeleteAdmissionregistrationV1beta1CollectionValidatingWebhookConfiguration creates a new http.Handler for the delete admissionregistration v1beta1 collection validating webhook configuration operation
func NewDeleteAdmissionregistrationV1beta1CollectionValidatingWebhookConfiguration(ctx *middleware.Context, handler DeleteAdmissionregistrationV1beta1CollectionValidatingWebhookConfigurationHandler) *DeleteAdmissionregistrationV1beta1CollectionValidatingWebhookConfiguration {
	return &DeleteAdmissionregistrationV1beta1CollectionValidatingWebhookConfiguration{Context: ctx, Handler: handler}
}

/*DeleteAdmissionregistrationV1beta1CollectionValidatingWebhookConfiguration swagger:route DELETE /apis/admissionregistration.k8s.io/v1beta1/validatingwebhookconfigurations admissionregistration_v1beta1 deleteAdmissionregistrationV1beta1CollectionValidatingWebhookConfiguration

delete collection of ValidatingWebhookConfiguration

*/
type DeleteAdmissionregistrationV1beta1CollectionValidatingWebhookConfiguration struct {
	Context *middleware.Context
	Handler DeleteAdmissionregistrationV1beta1CollectionValidatingWebhookConfigurationHandler
}

func (o *DeleteAdmissionregistrationV1beta1CollectionValidatingWebhookConfiguration) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteAdmissionregistrationV1beta1CollectionValidatingWebhookConfigurationParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
