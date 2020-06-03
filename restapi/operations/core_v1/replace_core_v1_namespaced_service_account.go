// Code generated by go-swagger; DO NOT EDIT.

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ReplaceCoreV1NamespacedServiceAccountHandlerFunc turns a function with the right signature into a replace core v1 namespaced service account handler
type ReplaceCoreV1NamespacedServiceAccountHandlerFunc func(ReplaceCoreV1NamespacedServiceAccountParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ReplaceCoreV1NamespacedServiceAccountHandlerFunc) Handle(params ReplaceCoreV1NamespacedServiceAccountParams) middleware.Responder {
	return fn(params)
}

// ReplaceCoreV1NamespacedServiceAccountHandler interface for that can handle valid replace core v1 namespaced service account params
type ReplaceCoreV1NamespacedServiceAccountHandler interface {
	Handle(ReplaceCoreV1NamespacedServiceAccountParams) middleware.Responder
}

// NewReplaceCoreV1NamespacedServiceAccount creates a new http.Handler for the replace core v1 namespaced service account operation
func NewReplaceCoreV1NamespacedServiceAccount(ctx *middleware.Context, handler ReplaceCoreV1NamespacedServiceAccountHandler) *ReplaceCoreV1NamespacedServiceAccount {
	return &ReplaceCoreV1NamespacedServiceAccount{Context: ctx, Handler: handler}
}

/*ReplaceCoreV1NamespacedServiceAccount swagger:route PUT /api/v1/namespaces/{namespace}/serviceaccounts/{name} core_v1 replaceCoreV1NamespacedServiceAccount

replace the specified ServiceAccount

*/
type ReplaceCoreV1NamespacedServiceAccount struct {
	Context *middleware.Context
	Handler ReplaceCoreV1NamespacedServiceAccountHandler
}

func (o *ReplaceCoreV1NamespacedServiceAccount) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewReplaceCoreV1NamespacedServiceAccountParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
