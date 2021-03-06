// Code generated by go-swagger; DO NOT EDIT.

package auditregistration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetAuditregistrationAPIGroupHandlerFunc turns a function with the right signature into a get auditregistration API group handler
type GetAuditregistrationAPIGroupHandlerFunc func(GetAuditregistrationAPIGroupParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetAuditregistrationAPIGroupHandlerFunc) Handle(params GetAuditregistrationAPIGroupParams) middleware.Responder {
	return fn(params)
}

// GetAuditregistrationAPIGroupHandler interface for that can handle valid get auditregistration API group params
type GetAuditregistrationAPIGroupHandler interface {
	Handle(GetAuditregistrationAPIGroupParams) middleware.Responder
}

// NewGetAuditregistrationAPIGroup creates a new http.Handler for the get auditregistration API group operation
func NewGetAuditregistrationAPIGroup(ctx *middleware.Context, handler GetAuditregistrationAPIGroupHandler) *GetAuditregistrationAPIGroup {
	return &GetAuditregistrationAPIGroup{Context: ctx, Handler: handler}
}

/*GetAuditregistrationAPIGroup swagger:route GET /apis/auditregistration.k8s.io/ auditregistration getAuditregistrationApiGroup

get information of a group

*/
type GetAuditregistrationAPIGroup struct {
	Context *middleware.Context
	Handler GetAuditregistrationAPIGroupHandler
}

func (o *GetAuditregistrationAPIGroup) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetAuditregistrationAPIGroupParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
