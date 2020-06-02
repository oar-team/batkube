// Code generated by go-swagger; DO NOT EDIT.

package auditregistration_v1alpha1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ReadAuditregistrationV1alpha1AuditSinkHandlerFunc turns a function with the right signature into a read auditregistration v1alpha1 audit sink handler
type ReadAuditregistrationV1alpha1AuditSinkHandlerFunc func(ReadAuditregistrationV1alpha1AuditSinkParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn ReadAuditregistrationV1alpha1AuditSinkHandlerFunc) Handle(params ReadAuditregistrationV1alpha1AuditSinkParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// ReadAuditregistrationV1alpha1AuditSinkHandler interface for that can handle valid read auditregistration v1alpha1 audit sink params
type ReadAuditregistrationV1alpha1AuditSinkHandler interface {
	Handle(ReadAuditregistrationV1alpha1AuditSinkParams, interface{}) middleware.Responder
}

// NewReadAuditregistrationV1alpha1AuditSink creates a new http.Handler for the read auditregistration v1alpha1 audit sink operation
func NewReadAuditregistrationV1alpha1AuditSink(ctx *middleware.Context, handler ReadAuditregistrationV1alpha1AuditSinkHandler) *ReadAuditregistrationV1alpha1AuditSink {
	return &ReadAuditregistrationV1alpha1AuditSink{Context: ctx, Handler: handler}
}

/*ReadAuditregistrationV1alpha1AuditSink swagger:route GET /apis/auditregistration.k8s.io/v1alpha1/auditsinks/{name} auditregistration_v1alpha1 readAuditregistrationV1alpha1AuditSink

read the specified AuditSink

*/
type ReadAuditregistrationV1alpha1AuditSink struct {
	Context *middleware.Context
	Handler ReadAuditregistrationV1alpha1AuditSinkHandler
}

func (o *ReadAuditregistrationV1alpha1AuditSink) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewReadAuditregistrationV1alpha1AuditSinkParams()

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
