// Code generated by go-swagger; DO NOT EDIT.

package storage_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteStorageV1CollectionVolumeAttachmentHandlerFunc turns a function with the right signature into a delete storage v1 collection volume attachment handler
type DeleteStorageV1CollectionVolumeAttachmentHandlerFunc func(DeleteStorageV1CollectionVolumeAttachmentParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteStorageV1CollectionVolumeAttachmentHandlerFunc) Handle(params DeleteStorageV1CollectionVolumeAttachmentParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// DeleteStorageV1CollectionVolumeAttachmentHandler interface for that can handle valid delete storage v1 collection volume attachment params
type DeleteStorageV1CollectionVolumeAttachmentHandler interface {
	Handle(DeleteStorageV1CollectionVolumeAttachmentParams, interface{}) middleware.Responder
}

// NewDeleteStorageV1CollectionVolumeAttachment creates a new http.Handler for the delete storage v1 collection volume attachment operation
func NewDeleteStorageV1CollectionVolumeAttachment(ctx *middleware.Context, handler DeleteStorageV1CollectionVolumeAttachmentHandler) *DeleteStorageV1CollectionVolumeAttachment {
	return &DeleteStorageV1CollectionVolumeAttachment{Context: ctx, Handler: handler}
}

/*DeleteStorageV1CollectionVolumeAttachment swagger:route DELETE /apis/storage.k8s.io/v1/volumeattachments storage_v1 deleteStorageV1CollectionVolumeAttachment

delete collection of VolumeAttachment

*/
type DeleteStorageV1CollectionVolumeAttachment struct {
	Context *middleware.Context
	Handler DeleteStorageV1CollectionVolumeAttachmentHandler
}

func (o *DeleteStorageV1CollectionVolumeAttachment) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteStorageV1CollectionVolumeAttachmentParams()

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
