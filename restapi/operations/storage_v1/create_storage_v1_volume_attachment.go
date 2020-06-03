// Code generated by go-swagger; DO NOT EDIT.

package storage_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// CreateStorageV1VolumeAttachmentHandlerFunc turns a function with the right signature into a create storage v1 volume attachment handler
type CreateStorageV1VolumeAttachmentHandlerFunc func(CreateStorageV1VolumeAttachmentParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateStorageV1VolumeAttachmentHandlerFunc) Handle(params CreateStorageV1VolumeAttachmentParams) middleware.Responder {
	return fn(params)
}

// CreateStorageV1VolumeAttachmentHandler interface for that can handle valid create storage v1 volume attachment params
type CreateStorageV1VolumeAttachmentHandler interface {
	Handle(CreateStorageV1VolumeAttachmentParams) middleware.Responder
}

// NewCreateStorageV1VolumeAttachment creates a new http.Handler for the create storage v1 volume attachment operation
func NewCreateStorageV1VolumeAttachment(ctx *middleware.Context, handler CreateStorageV1VolumeAttachmentHandler) *CreateStorageV1VolumeAttachment {
	return &CreateStorageV1VolumeAttachment{Context: ctx, Handler: handler}
}

/*CreateStorageV1VolumeAttachment swagger:route POST /apis/storage.k8s.io/v1/volumeattachments storage_v1 createStorageV1VolumeAttachment

create a VolumeAttachment

*/
type CreateStorageV1VolumeAttachment struct {
	Context *middleware.Context
	Handler CreateStorageV1VolumeAttachmentHandler
}

func (o *CreateStorageV1VolumeAttachment) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCreateStorageV1VolumeAttachmentParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
