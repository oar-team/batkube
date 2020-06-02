// Code generated by go-swagger; DO NOT EDIT.

package auditregistration_v1alpha1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteAuditregistrationV1alpha1CollectionAuditSinkHandlerFunc turns a function with the right signature into a delete auditregistration v1alpha1 collection audit sink handler
type DeleteAuditregistrationV1alpha1CollectionAuditSinkHandlerFunc func(DeleteAuditregistrationV1alpha1CollectionAuditSinkParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteAuditregistrationV1alpha1CollectionAuditSinkHandlerFunc) Handle(params DeleteAuditregistrationV1alpha1CollectionAuditSinkParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// DeleteAuditregistrationV1alpha1CollectionAuditSinkHandler interface for that can handle valid delete auditregistration v1alpha1 collection audit sink params
type DeleteAuditregistrationV1alpha1CollectionAuditSinkHandler interface {
	Handle(DeleteAuditregistrationV1alpha1CollectionAuditSinkParams, interface{}) middleware.Responder
}

// NewDeleteAuditregistrationV1alpha1CollectionAuditSink creates a new http.Handler for the delete auditregistration v1alpha1 collection audit sink operation
func NewDeleteAuditregistrationV1alpha1CollectionAuditSink(ctx *middleware.Context, handler DeleteAuditregistrationV1alpha1CollectionAuditSinkHandler) *DeleteAuditregistrationV1alpha1CollectionAuditSink {
	return &DeleteAuditregistrationV1alpha1CollectionAuditSink{Context: ctx, Handler: handler}
}

/*DeleteAuditregistrationV1alpha1CollectionAuditSink swagger:route DELETE /apis/auditregistration.k8s.io/v1alpha1/auditsinks auditregistration_v1alpha1 deleteAuditregistrationV1alpha1CollectionAuditSink

delete collection of AuditSink

*/
type DeleteAuditregistrationV1alpha1CollectionAuditSink struct {
	Context *middleware.Context
	Handler DeleteAuditregistrationV1alpha1CollectionAuditSinkHandler
}

func (o *DeleteAuditregistrationV1alpha1CollectionAuditSink) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteAuditregistrationV1alpha1CollectionAuditSinkParams()

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
