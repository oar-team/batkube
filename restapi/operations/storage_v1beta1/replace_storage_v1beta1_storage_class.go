// Code generated by go-swagger; DO NOT EDIT.

package storage_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ReplaceStorageV1beta1StorageClassHandlerFunc turns a function with the right signature into a replace storage v1beta1 storage class handler
type ReplaceStorageV1beta1StorageClassHandlerFunc func(ReplaceStorageV1beta1StorageClassParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ReplaceStorageV1beta1StorageClassHandlerFunc) Handle(params ReplaceStorageV1beta1StorageClassParams) middleware.Responder {
	return fn(params)
}

// ReplaceStorageV1beta1StorageClassHandler interface for that can handle valid replace storage v1beta1 storage class params
type ReplaceStorageV1beta1StorageClassHandler interface {
	Handle(ReplaceStorageV1beta1StorageClassParams) middleware.Responder
}

// NewReplaceStorageV1beta1StorageClass creates a new http.Handler for the replace storage v1beta1 storage class operation
func NewReplaceStorageV1beta1StorageClass(ctx *middleware.Context, handler ReplaceStorageV1beta1StorageClassHandler) *ReplaceStorageV1beta1StorageClass {
	return &ReplaceStorageV1beta1StorageClass{Context: ctx, Handler: handler}
}

/*ReplaceStorageV1beta1StorageClass swagger:route PUT /apis/storage.k8s.io/v1beta1/storageclasses/{name} storage_v1beta1 replaceStorageV1beta1StorageClass

replace the specified StorageClass

*/
type ReplaceStorageV1beta1StorageClass struct {
	Context *middleware.Context
	Handler ReplaceStorageV1beta1StorageClassHandler
}

func (o *ReplaceStorageV1beta1StorageClass) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewReplaceStorageV1beta1StorageClassParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
