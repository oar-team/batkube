// Code generated by go-swagger; DO NOT EDIT.

package admissionregistration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetAdmissionregistrationAPIGroupHandlerFunc turns a function with the right signature into a get admissionregistration API group handler
type GetAdmissionregistrationAPIGroupHandlerFunc func(GetAdmissionregistrationAPIGroupParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetAdmissionregistrationAPIGroupHandlerFunc) Handle(params GetAdmissionregistrationAPIGroupParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetAdmissionregistrationAPIGroupHandler interface for that can handle valid get admissionregistration API group params
type GetAdmissionregistrationAPIGroupHandler interface {
	Handle(GetAdmissionregistrationAPIGroupParams, interface{}) middleware.Responder
}

// NewGetAdmissionregistrationAPIGroup creates a new http.Handler for the get admissionregistration API group operation
func NewGetAdmissionregistrationAPIGroup(ctx *middleware.Context, handler GetAdmissionregistrationAPIGroupHandler) *GetAdmissionregistrationAPIGroup {
	return &GetAdmissionregistrationAPIGroup{Context: ctx, Handler: handler}
}

/*GetAdmissionregistrationAPIGroup swagger:route GET /apis/admissionregistration.k8s.io/ admissionregistration getAdmissionregistrationApiGroup

get information of a group

*/
type GetAdmissionregistrationAPIGroup struct {
	Context *middleware.Context
	Handler GetAdmissionregistrationAPIGroupHandler
}

func (o *GetAdmissionregistrationAPIGroup) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetAdmissionregistrationAPIGroupParams()

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
