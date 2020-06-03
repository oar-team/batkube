// Code generated by go-swagger; DO NOT EDIT.

package storage_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteStorageV1beta1CSIDriverHandlerFunc turns a function with the right signature into a delete storage v1beta1 c s i driver handler
type DeleteStorageV1beta1CSIDriverHandlerFunc func(DeleteStorageV1beta1CSIDriverParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteStorageV1beta1CSIDriverHandlerFunc) Handle(params DeleteStorageV1beta1CSIDriverParams) middleware.Responder {
	return fn(params)
}

// DeleteStorageV1beta1CSIDriverHandler interface for that can handle valid delete storage v1beta1 c s i driver params
type DeleteStorageV1beta1CSIDriverHandler interface {
	Handle(DeleteStorageV1beta1CSIDriverParams) middleware.Responder
}

// NewDeleteStorageV1beta1CSIDriver creates a new http.Handler for the delete storage v1beta1 c s i driver operation
func NewDeleteStorageV1beta1CSIDriver(ctx *middleware.Context, handler DeleteStorageV1beta1CSIDriverHandler) *DeleteStorageV1beta1CSIDriver {
	return &DeleteStorageV1beta1CSIDriver{Context: ctx, Handler: handler}
}

/*DeleteStorageV1beta1CSIDriver swagger:route DELETE /apis/storage.k8s.io/v1beta1/csidrivers/{name} storage_v1beta1 deleteStorageV1beta1CSIDriver

delete a CSIDriver

*/
type DeleteStorageV1beta1CSIDriver struct {
	Context *middleware.Context
	Handler DeleteStorageV1beta1CSIDriverHandler
}

func (o *DeleteStorageV1beta1CSIDriver) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteStorageV1beta1CSIDriverParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}