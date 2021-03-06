// Code generated by go-swagger; DO NOT EDIT.

package storage_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ListStorageV1beta1StorageClassHandlerFunc turns a function with the right signature into a list storage v1beta1 storage class handler
type ListStorageV1beta1StorageClassHandlerFunc func(ListStorageV1beta1StorageClassParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ListStorageV1beta1StorageClassHandlerFunc) Handle(params ListStorageV1beta1StorageClassParams) middleware.Responder {
	return fn(params)
}

// ListStorageV1beta1StorageClassHandler interface for that can handle valid list storage v1beta1 storage class params
type ListStorageV1beta1StorageClassHandler interface {
	Handle(ListStorageV1beta1StorageClassParams) middleware.Responder
}

// NewListStorageV1beta1StorageClass creates a new http.Handler for the list storage v1beta1 storage class operation
func NewListStorageV1beta1StorageClass(ctx *middleware.Context, handler ListStorageV1beta1StorageClassHandler) *ListStorageV1beta1StorageClass {
	return &ListStorageV1beta1StorageClass{Context: ctx, Handler: handler}
}

/*ListStorageV1beta1StorageClass swagger:route GET /apis/storage.k8s.io/v1beta1/storageclasses storage_v1beta1 listStorageV1beta1StorageClass

list or watch objects of kind StorageClass

*/
type ListStorageV1beta1StorageClass struct {
	Context *middleware.Context
	Handler ListStorageV1beta1StorageClassHandler
}

func (o *ListStorageV1beta1StorageClass) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewListStorageV1beta1StorageClassParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
