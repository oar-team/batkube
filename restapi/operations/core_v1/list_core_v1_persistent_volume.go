// Code generated by go-swagger; DO NOT EDIT.

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ListCoreV1PersistentVolumeHandlerFunc turns a function with the right signature into a list core v1 persistent volume handler
type ListCoreV1PersistentVolumeHandlerFunc func(ListCoreV1PersistentVolumeParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ListCoreV1PersistentVolumeHandlerFunc) Handle(params ListCoreV1PersistentVolumeParams) middleware.Responder {
	return fn(params)
}

// ListCoreV1PersistentVolumeHandler interface for that can handle valid list core v1 persistent volume params
type ListCoreV1PersistentVolumeHandler interface {
	Handle(ListCoreV1PersistentVolumeParams) middleware.Responder
}

// NewListCoreV1PersistentVolume creates a new http.Handler for the list core v1 persistent volume operation
func NewListCoreV1PersistentVolume(ctx *middleware.Context, handler ListCoreV1PersistentVolumeHandler) *ListCoreV1PersistentVolume {
	return &ListCoreV1PersistentVolume{Context: ctx, Handler: handler}
}

/*ListCoreV1PersistentVolume swagger:route GET /api/v1/persistentvolumes core_v1 listCoreV1PersistentVolume

list or watch objects of kind PersistentVolume

*/
type ListCoreV1PersistentVolume struct {
	Context *middleware.Context
	Handler ListCoreV1PersistentVolumeHandler
}

func (o *ListCoreV1PersistentVolume) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewListCoreV1PersistentVolumeParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}