// Code generated by go-swagger; DO NOT EDIT.

package storage_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ReplaceStorageV1StorageClassHandlerFunc turns a function with the right signature into a replace storage v1 storage class handler
type ReplaceStorageV1StorageClassHandlerFunc func(ReplaceStorageV1StorageClassParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn ReplaceStorageV1StorageClassHandlerFunc) Handle(params ReplaceStorageV1StorageClassParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// ReplaceStorageV1StorageClassHandler interface for that can handle valid replace storage v1 storage class params
type ReplaceStorageV1StorageClassHandler interface {
	Handle(ReplaceStorageV1StorageClassParams, interface{}) middleware.Responder
}

// NewReplaceStorageV1StorageClass creates a new http.Handler for the replace storage v1 storage class operation
func NewReplaceStorageV1StorageClass(ctx *middleware.Context, handler ReplaceStorageV1StorageClassHandler) *ReplaceStorageV1StorageClass {
	return &ReplaceStorageV1StorageClass{Context: ctx, Handler: handler}
}

/*ReplaceStorageV1StorageClass swagger:route PUT /apis/storage.k8s.io/v1/storageclasses/{name} storage_v1 replaceStorageV1StorageClass

replace the specified StorageClass

*/
type ReplaceStorageV1StorageClass struct {
	Context *middleware.Context
	Handler ReplaceStorageV1StorageClassHandler
}

func (o *ReplaceStorageV1StorageClass) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewReplaceStorageV1StorageClassParams()

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
