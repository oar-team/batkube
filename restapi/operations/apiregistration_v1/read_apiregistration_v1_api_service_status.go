// Code generated by go-swagger; DO NOT EDIT.

package apiregistration_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ReadApiregistrationV1APIServiceStatusHandlerFunc turns a function with the right signature into a read apiregistration v1 API service status handler
type ReadApiregistrationV1APIServiceStatusHandlerFunc func(ReadApiregistrationV1APIServiceStatusParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn ReadApiregistrationV1APIServiceStatusHandlerFunc) Handle(params ReadApiregistrationV1APIServiceStatusParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// ReadApiregistrationV1APIServiceStatusHandler interface for that can handle valid read apiregistration v1 API service status params
type ReadApiregistrationV1APIServiceStatusHandler interface {
	Handle(ReadApiregistrationV1APIServiceStatusParams, interface{}) middleware.Responder
}

// NewReadApiregistrationV1APIServiceStatus creates a new http.Handler for the read apiregistration v1 API service status operation
func NewReadApiregistrationV1APIServiceStatus(ctx *middleware.Context, handler ReadApiregistrationV1APIServiceStatusHandler) *ReadApiregistrationV1APIServiceStatus {
	return &ReadApiregistrationV1APIServiceStatus{Context: ctx, Handler: handler}
}

/*ReadApiregistrationV1APIServiceStatus swagger:route GET /apis/apiregistration.k8s.io/v1/apiservices/{name}/status apiregistration_v1 readApiregistrationV1ApiServiceStatus

read status of the specified APIService

*/
type ReadApiregistrationV1APIServiceStatus struct {
	Context *middleware.Context
	Handler ReadApiregistrationV1APIServiceStatusHandler
}

func (o *ReadApiregistrationV1APIServiceStatus) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewReadApiregistrationV1APIServiceStatusParams()

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
