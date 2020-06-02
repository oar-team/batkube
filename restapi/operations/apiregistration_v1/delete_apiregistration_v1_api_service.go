// Code generated by go-swagger; DO NOT EDIT.

package apiregistration_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteApiregistrationV1APIServiceHandlerFunc turns a function with the right signature into a delete apiregistration v1 API service handler
type DeleteApiregistrationV1APIServiceHandlerFunc func(DeleteApiregistrationV1APIServiceParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteApiregistrationV1APIServiceHandlerFunc) Handle(params DeleteApiregistrationV1APIServiceParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// DeleteApiregistrationV1APIServiceHandler interface for that can handle valid delete apiregistration v1 API service params
type DeleteApiregistrationV1APIServiceHandler interface {
	Handle(DeleteApiregistrationV1APIServiceParams, interface{}) middleware.Responder
}

// NewDeleteApiregistrationV1APIService creates a new http.Handler for the delete apiregistration v1 API service operation
func NewDeleteApiregistrationV1APIService(ctx *middleware.Context, handler DeleteApiregistrationV1APIServiceHandler) *DeleteApiregistrationV1APIService {
	return &DeleteApiregistrationV1APIService{Context: ctx, Handler: handler}
}

/*DeleteApiregistrationV1APIService swagger:route DELETE /apis/apiregistration.k8s.io/v1/apiservices/{name} apiregistration_v1 deleteApiregistrationV1ApiService

delete an APIService

*/
type DeleteApiregistrationV1APIService struct {
	Context *middleware.Context
	Handler DeleteApiregistrationV1APIServiceHandler
}

func (o *DeleteApiregistrationV1APIService) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteApiregistrationV1APIServiceParams()

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
