// Code generated by go-swagger; DO NOT EDIT.

package admissionregistration_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// WatchAdmissionregistrationV1MutatingWebhookConfigurationListHandlerFunc turns a function with the right signature into a watch admissionregistration v1 mutating webhook configuration list handler
type WatchAdmissionregistrationV1MutatingWebhookConfigurationListHandlerFunc func(WatchAdmissionregistrationV1MutatingWebhookConfigurationListParams) middleware.Responder

// Handle executing the request and returning a response
func (fn WatchAdmissionregistrationV1MutatingWebhookConfigurationListHandlerFunc) Handle(params WatchAdmissionregistrationV1MutatingWebhookConfigurationListParams) middleware.Responder {
	return fn(params)
}

// WatchAdmissionregistrationV1MutatingWebhookConfigurationListHandler interface for that can handle valid watch admissionregistration v1 mutating webhook configuration list params
type WatchAdmissionregistrationV1MutatingWebhookConfigurationListHandler interface {
	Handle(WatchAdmissionregistrationV1MutatingWebhookConfigurationListParams) middleware.Responder
}

// NewWatchAdmissionregistrationV1MutatingWebhookConfigurationList creates a new http.Handler for the watch admissionregistration v1 mutating webhook configuration list operation
func NewWatchAdmissionregistrationV1MutatingWebhookConfigurationList(ctx *middleware.Context, handler WatchAdmissionregistrationV1MutatingWebhookConfigurationListHandler) *WatchAdmissionregistrationV1MutatingWebhookConfigurationList {
	return &WatchAdmissionregistrationV1MutatingWebhookConfigurationList{Context: ctx, Handler: handler}
}

/*WatchAdmissionregistrationV1MutatingWebhookConfigurationList swagger:route GET /apis/admissionregistration.k8s.io/v1/watch/mutatingwebhookconfigurations admissionregistration_v1 watchAdmissionregistrationV1MutatingWebhookConfigurationList

watch individual changes to a list of MutatingWebhookConfiguration. deprecated: use the 'watch' parameter with a list operation instead.

*/
type WatchAdmissionregistrationV1MutatingWebhookConfigurationList struct {
	Context *middleware.Context
	Handler WatchAdmissionregistrationV1MutatingWebhookConfigurationListHandler
}

func (o *WatchAdmissionregistrationV1MutatingWebhookConfigurationList) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewWatchAdmissionregistrationV1MutatingWebhookConfigurationListParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
