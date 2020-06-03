// Code generated by go-swagger; DO NOT EDIT.

package auditregistration_v1alpha1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteAuditregistrationV1alpha1AuditSinkHandlerFunc turns a function with the right signature into a delete auditregistration v1alpha1 audit sink handler
type DeleteAuditregistrationV1alpha1AuditSinkHandlerFunc func(DeleteAuditregistrationV1alpha1AuditSinkParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteAuditregistrationV1alpha1AuditSinkHandlerFunc) Handle(params DeleteAuditregistrationV1alpha1AuditSinkParams) middleware.Responder {
	return fn(params)
}

// DeleteAuditregistrationV1alpha1AuditSinkHandler interface for that can handle valid delete auditregistration v1alpha1 audit sink params
type DeleteAuditregistrationV1alpha1AuditSinkHandler interface {
	Handle(DeleteAuditregistrationV1alpha1AuditSinkParams) middleware.Responder
}

// NewDeleteAuditregistrationV1alpha1AuditSink creates a new http.Handler for the delete auditregistration v1alpha1 audit sink operation
func NewDeleteAuditregistrationV1alpha1AuditSink(ctx *middleware.Context, handler DeleteAuditregistrationV1alpha1AuditSinkHandler) *DeleteAuditregistrationV1alpha1AuditSink {
	return &DeleteAuditregistrationV1alpha1AuditSink{Context: ctx, Handler: handler}
}

/*DeleteAuditregistrationV1alpha1AuditSink swagger:route DELETE /apis/auditregistration.k8s.io/v1alpha1/auditsinks/{name} auditregistration_v1alpha1 deleteAuditregistrationV1alpha1AuditSink

delete an AuditSink

*/
type DeleteAuditregistrationV1alpha1AuditSink struct {
	Context *middleware.Context
	Handler DeleteAuditregistrationV1alpha1AuditSinkHandler
}

func (o *DeleteAuditregistrationV1alpha1AuditSink) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteAuditregistrationV1alpha1AuditSinkParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}