// Code generated by go-swagger; DO NOT EDIT.

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// CreateCoreV1NamespacedServiceHandlerFunc turns a function with the right signature into a create core v1 namespaced service handler
type CreateCoreV1NamespacedServiceHandlerFunc func(CreateCoreV1NamespacedServiceParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateCoreV1NamespacedServiceHandlerFunc) Handle(params CreateCoreV1NamespacedServiceParams) middleware.Responder {
	return fn(params)
}

// CreateCoreV1NamespacedServiceHandler interface for that can handle valid create core v1 namespaced service params
type CreateCoreV1NamespacedServiceHandler interface {
	Handle(CreateCoreV1NamespacedServiceParams) middleware.Responder
}

// NewCreateCoreV1NamespacedService creates a new http.Handler for the create core v1 namespaced service operation
func NewCreateCoreV1NamespacedService(ctx *middleware.Context, handler CreateCoreV1NamespacedServiceHandler) *CreateCoreV1NamespacedService {
	return &CreateCoreV1NamespacedService{Context: ctx, Handler: handler}
}

/*CreateCoreV1NamespacedService swagger:route POST /api/v1/namespaces/{namespace}/services core_v1 createCoreV1NamespacedService

create a Service

*/
type CreateCoreV1NamespacedService struct {
	Context *middleware.Context
	Handler CreateCoreV1NamespacedServiceHandler
}

func (o *CreateCoreV1NamespacedService) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCreateCoreV1NamespacedServiceParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
