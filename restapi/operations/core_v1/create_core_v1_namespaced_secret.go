// Code generated by go-swagger; DO NOT EDIT.

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// CreateCoreV1NamespacedSecretHandlerFunc turns a function with the right signature into a create core v1 namespaced secret handler
type CreateCoreV1NamespacedSecretHandlerFunc func(CreateCoreV1NamespacedSecretParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateCoreV1NamespacedSecretHandlerFunc) Handle(params CreateCoreV1NamespacedSecretParams) middleware.Responder {
	return fn(params)
}

// CreateCoreV1NamespacedSecretHandler interface for that can handle valid create core v1 namespaced secret params
type CreateCoreV1NamespacedSecretHandler interface {
	Handle(CreateCoreV1NamespacedSecretParams) middleware.Responder
}

// NewCreateCoreV1NamespacedSecret creates a new http.Handler for the create core v1 namespaced secret operation
func NewCreateCoreV1NamespacedSecret(ctx *middleware.Context, handler CreateCoreV1NamespacedSecretHandler) *CreateCoreV1NamespacedSecret {
	return &CreateCoreV1NamespacedSecret{Context: ctx, Handler: handler}
}

/*CreateCoreV1NamespacedSecret swagger:route POST /api/v1/namespaces/{namespace}/secrets core_v1 createCoreV1NamespacedSecret

create a Secret

*/
type CreateCoreV1NamespacedSecret struct {
	Context *middleware.Context
	Handler CreateCoreV1NamespacedSecretHandler
}

func (o *CreateCoreV1NamespacedSecret) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCreateCoreV1NamespacedSecretParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
