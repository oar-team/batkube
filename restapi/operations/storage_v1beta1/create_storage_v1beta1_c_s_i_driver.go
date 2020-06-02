// Code generated by go-swagger; DO NOT EDIT.

package storage_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// CreateStorageV1beta1CSIDriverHandlerFunc turns a function with the right signature into a create storage v1beta1 c s i driver handler
type CreateStorageV1beta1CSIDriverHandlerFunc func(CreateStorageV1beta1CSIDriverParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateStorageV1beta1CSIDriverHandlerFunc) Handle(params CreateStorageV1beta1CSIDriverParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// CreateStorageV1beta1CSIDriverHandler interface for that can handle valid create storage v1beta1 c s i driver params
type CreateStorageV1beta1CSIDriverHandler interface {
	Handle(CreateStorageV1beta1CSIDriverParams, interface{}) middleware.Responder
}

// NewCreateStorageV1beta1CSIDriver creates a new http.Handler for the create storage v1beta1 c s i driver operation
func NewCreateStorageV1beta1CSIDriver(ctx *middleware.Context, handler CreateStorageV1beta1CSIDriverHandler) *CreateStorageV1beta1CSIDriver {
	return &CreateStorageV1beta1CSIDriver{Context: ctx, Handler: handler}
}

/*CreateStorageV1beta1CSIDriver swagger:route POST /apis/storage.k8s.io/v1beta1/csidrivers storage_v1beta1 createStorageV1beta1CSIDriver

create a CSIDriver

*/
type CreateStorageV1beta1CSIDriver struct {
	Context *middleware.Context
	Handler CreateStorageV1beta1CSIDriverHandler
}

func (o *CreateStorageV1beta1CSIDriver) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCreateStorageV1beta1CSIDriverParams()

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
