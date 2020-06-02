// Code generated by go-swagger; DO NOT EDIT.

package storage_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteStorageV1CollectionCSIDriverHandlerFunc turns a function with the right signature into a delete storage v1 collection c s i driver handler
type DeleteStorageV1CollectionCSIDriverHandlerFunc func(DeleteStorageV1CollectionCSIDriverParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteStorageV1CollectionCSIDriverHandlerFunc) Handle(params DeleteStorageV1CollectionCSIDriverParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// DeleteStorageV1CollectionCSIDriverHandler interface for that can handle valid delete storage v1 collection c s i driver params
type DeleteStorageV1CollectionCSIDriverHandler interface {
	Handle(DeleteStorageV1CollectionCSIDriverParams, interface{}) middleware.Responder
}

// NewDeleteStorageV1CollectionCSIDriver creates a new http.Handler for the delete storage v1 collection c s i driver operation
func NewDeleteStorageV1CollectionCSIDriver(ctx *middleware.Context, handler DeleteStorageV1CollectionCSIDriverHandler) *DeleteStorageV1CollectionCSIDriver {
	return &DeleteStorageV1CollectionCSIDriver{Context: ctx, Handler: handler}
}

/*DeleteStorageV1CollectionCSIDriver swagger:route DELETE /apis/storage.k8s.io/v1/csidrivers storage_v1 deleteStorageV1CollectionCSIDriver

delete collection of CSIDriver

*/
type DeleteStorageV1CollectionCSIDriver struct {
	Context *middleware.Context
	Handler DeleteStorageV1CollectionCSIDriverHandler
}

func (o *DeleteStorageV1CollectionCSIDriver) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteStorageV1CollectionCSIDriverParams()

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
