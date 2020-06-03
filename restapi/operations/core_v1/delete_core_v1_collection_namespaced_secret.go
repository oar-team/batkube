// Code generated by go-swagger; DO NOT EDIT.

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteCoreV1CollectionNamespacedSecretHandlerFunc turns a function with the right signature into a delete core v1 collection namespaced secret handler
type DeleteCoreV1CollectionNamespacedSecretHandlerFunc func(DeleteCoreV1CollectionNamespacedSecretParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteCoreV1CollectionNamespacedSecretHandlerFunc) Handle(params DeleteCoreV1CollectionNamespacedSecretParams) middleware.Responder {
	return fn(params)
}

// DeleteCoreV1CollectionNamespacedSecretHandler interface for that can handle valid delete core v1 collection namespaced secret params
type DeleteCoreV1CollectionNamespacedSecretHandler interface {
	Handle(DeleteCoreV1CollectionNamespacedSecretParams) middleware.Responder
}

// NewDeleteCoreV1CollectionNamespacedSecret creates a new http.Handler for the delete core v1 collection namespaced secret operation
func NewDeleteCoreV1CollectionNamespacedSecret(ctx *middleware.Context, handler DeleteCoreV1CollectionNamespacedSecretHandler) *DeleteCoreV1CollectionNamespacedSecret {
	return &DeleteCoreV1CollectionNamespacedSecret{Context: ctx, Handler: handler}
}

/*DeleteCoreV1CollectionNamespacedSecret swagger:route DELETE /api/v1/namespaces/{namespace}/secrets core_v1 deleteCoreV1CollectionNamespacedSecret

delete collection of Secret

*/
type DeleteCoreV1CollectionNamespacedSecret struct {
	Context *middleware.Context
	Handler DeleteCoreV1CollectionNamespacedSecretHandler
}

func (o *DeleteCoreV1CollectionNamespacedSecret) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteCoreV1CollectionNamespacedSecretParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
