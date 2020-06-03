// Code generated by go-swagger; DO NOT EDIT.

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ReplaceCoreV1NamespaceStatusHandlerFunc turns a function with the right signature into a replace core v1 namespace status handler
type ReplaceCoreV1NamespaceStatusHandlerFunc func(ReplaceCoreV1NamespaceStatusParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ReplaceCoreV1NamespaceStatusHandlerFunc) Handle(params ReplaceCoreV1NamespaceStatusParams) middleware.Responder {
	return fn(params)
}

// ReplaceCoreV1NamespaceStatusHandler interface for that can handle valid replace core v1 namespace status params
type ReplaceCoreV1NamespaceStatusHandler interface {
	Handle(ReplaceCoreV1NamespaceStatusParams) middleware.Responder
}

// NewReplaceCoreV1NamespaceStatus creates a new http.Handler for the replace core v1 namespace status operation
func NewReplaceCoreV1NamespaceStatus(ctx *middleware.Context, handler ReplaceCoreV1NamespaceStatusHandler) *ReplaceCoreV1NamespaceStatus {
	return &ReplaceCoreV1NamespaceStatus{Context: ctx, Handler: handler}
}

/*ReplaceCoreV1NamespaceStatus swagger:route PUT /api/v1/namespaces/{name}/status core_v1 replaceCoreV1NamespaceStatus

replace status of the specified Namespace

*/
type ReplaceCoreV1NamespaceStatus struct {
	Context *middleware.Context
	Handler ReplaceCoreV1NamespaceStatusHandler
}

func (o *ReplaceCoreV1NamespaceStatus) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewReplaceCoreV1NamespaceStatusParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}