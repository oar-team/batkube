// Code generated by go-swagger; DO NOT EDIT.

package storage_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ListStorageV1CSIDriverHandlerFunc turns a function with the right signature into a list storage v1 c s i driver handler
type ListStorageV1CSIDriverHandlerFunc func(ListStorageV1CSIDriverParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ListStorageV1CSIDriverHandlerFunc) Handle(params ListStorageV1CSIDriverParams) middleware.Responder {
	return fn(params)
}

// ListStorageV1CSIDriverHandler interface for that can handle valid list storage v1 c s i driver params
type ListStorageV1CSIDriverHandler interface {
	Handle(ListStorageV1CSIDriverParams) middleware.Responder
}

// NewListStorageV1CSIDriver creates a new http.Handler for the list storage v1 c s i driver operation
func NewListStorageV1CSIDriver(ctx *middleware.Context, handler ListStorageV1CSIDriverHandler) *ListStorageV1CSIDriver {
	return &ListStorageV1CSIDriver{Context: ctx, Handler: handler}
}

/*ListStorageV1CSIDriver swagger:route GET /apis/storage.k8s.io/v1/csidrivers storage_v1 listStorageV1CSIDriver

list or watch objects of kind CSIDriver

*/
type ListStorageV1CSIDriver struct {
	Context *middleware.Context
	Handler ListStorageV1CSIDriverHandler
}

func (o *ListStorageV1CSIDriver) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewListStorageV1CSIDriverParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
