// Code generated by go-swagger; DO NOT EDIT.

package auditregistration_v1alpha1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// WatchAuditregistrationV1alpha1AuditSinkListHandlerFunc turns a function with the right signature into a watch auditregistration v1alpha1 audit sink list handler
type WatchAuditregistrationV1alpha1AuditSinkListHandlerFunc func(WatchAuditregistrationV1alpha1AuditSinkListParams) middleware.Responder

// Handle executing the request and returning a response
func (fn WatchAuditregistrationV1alpha1AuditSinkListHandlerFunc) Handle(params WatchAuditregistrationV1alpha1AuditSinkListParams) middleware.Responder {
	return fn(params)
}

// WatchAuditregistrationV1alpha1AuditSinkListHandler interface for that can handle valid watch auditregistration v1alpha1 audit sink list params
type WatchAuditregistrationV1alpha1AuditSinkListHandler interface {
	Handle(WatchAuditregistrationV1alpha1AuditSinkListParams) middleware.Responder
}

// NewWatchAuditregistrationV1alpha1AuditSinkList creates a new http.Handler for the watch auditregistration v1alpha1 audit sink list operation
func NewWatchAuditregistrationV1alpha1AuditSinkList(ctx *middleware.Context, handler WatchAuditregistrationV1alpha1AuditSinkListHandler) *WatchAuditregistrationV1alpha1AuditSinkList {
	return &WatchAuditregistrationV1alpha1AuditSinkList{Context: ctx, Handler: handler}
}

/*WatchAuditregistrationV1alpha1AuditSinkList swagger:route GET /apis/auditregistration.k8s.io/v1alpha1/watch/auditsinks auditregistration_v1alpha1 watchAuditregistrationV1alpha1AuditSinkList

watch individual changes to a list of AuditSink. deprecated: use the 'watch' parameter with a list operation instead.

*/
type WatchAuditregistrationV1alpha1AuditSinkList struct {
	Context *middleware.Context
	Handler WatchAuditregistrationV1alpha1AuditSinkListHandler
}

func (o *WatchAuditregistrationV1alpha1AuditSinkList) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewWatchAuditregistrationV1alpha1AuditSinkListParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
