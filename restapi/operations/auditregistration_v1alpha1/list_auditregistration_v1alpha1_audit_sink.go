// Code generated by go-swagger; DO NOT EDIT.

package auditregistration_v1alpha1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ListAuditregistrationV1alpha1AuditSinkHandlerFunc turns a function with the right signature into a list auditregistration v1alpha1 audit sink handler
type ListAuditregistrationV1alpha1AuditSinkHandlerFunc func(ListAuditregistrationV1alpha1AuditSinkParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn ListAuditregistrationV1alpha1AuditSinkHandlerFunc) Handle(params ListAuditregistrationV1alpha1AuditSinkParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// ListAuditregistrationV1alpha1AuditSinkHandler interface for that can handle valid list auditregistration v1alpha1 audit sink params
type ListAuditregistrationV1alpha1AuditSinkHandler interface {
	Handle(ListAuditregistrationV1alpha1AuditSinkParams, interface{}) middleware.Responder
}

// NewListAuditregistrationV1alpha1AuditSink creates a new http.Handler for the list auditregistration v1alpha1 audit sink operation
func NewListAuditregistrationV1alpha1AuditSink(ctx *middleware.Context, handler ListAuditregistrationV1alpha1AuditSinkHandler) *ListAuditregistrationV1alpha1AuditSink {
	return &ListAuditregistrationV1alpha1AuditSink{Context: ctx, Handler: handler}
}

/*ListAuditregistrationV1alpha1AuditSink swagger:route GET /apis/auditregistration.k8s.io/v1alpha1/auditsinks auditregistration_v1alpha1 listAuditregistrationV1alpha1AuditSink

list or watch objects of kind AuditSink

*/
type ListAuditregistrationV1alpha1AuditSink struct {
	Context *middleware.Context
	Handler ListAuditregistrationV1alpha1AuditSinkHandler
}

func (o *ListAuditregistrationV1alpha1AuditSink) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewListAuditregistrationV1alpha1AuditSinkParams()

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
