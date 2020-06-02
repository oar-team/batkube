// Code generated by go-swagger; DO NOT EDIT.

package auditregistration_v1alpha1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// CreateAuditregistrationV1alpha1AuditSinkHandlerFunc turns a function with the right signature into a create auditregistration v1alpha1 audit sink handler
type CreateAuditregistrationV1alpha1AuditSinkHandlerFunc func(CreateAuditregistrationV1alpha1AuditSinkParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateAuditregistrationV1alpha1AuditSinkHandlerFunc) Handle(params CreateAuditregistrationV1alpha1AuditSinkParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// CreateAuditregistrationV1alpha1AuditSinkHandler interface for that can handle valid create auditregistration v1alpha1 audit sink params
type CreateAuditregistrationV1alpha1AuditSinkHandler interface {
	Handle(CreateAuditregistrationV1alpha1AuditSinkParams, interface{}) middleware.Responder
}

// NewCreateAuditregistrationV1alpha1AuditSink creates a new http.Handler for the create auditregistration v1alpha1 audit sink operation
func NewCreateAuditregistrationV1alpha1AuditSink(ctx *middleware.Context, handler CreateAuditregistrationV1alpha1AuditSinkHandler) *CreateAuditregistrationV1alpha1AuditSink {
	return &CreateAuditregistrationV1alpha1AuditSink{Context: ctx, Handler: handler}
}

/*CreateAuditregistrationV1alpha1AuditSink swagger:route POST /apis/auditregistration.k8s.io/v1alpha1/auditsinks auditregistration_v1alpha1 createAuditregistrationV1alpha1AuditSink

create an AuditSink

*/
type CreateAuditregistrationV1alpha1AuditSink struct {
	Context *middleware.Context
	Handler CreateAuditregistrationV1alpha1AuditSinkHandler
}

func (o *CreateAuditregistrationV1alpha1AuditSink) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCreateAuditregistrationV1alpha1AuditSinkParams()

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
