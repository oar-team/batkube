// Code generated by go-swagger; DO NOT EDIT.

package storage_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ReadStorageV1VolumeAttachmentStatusHandlerFunc turns a function with the right signature into a read storage v1 volume attachment status handler
type ReadStorageV1VolumeAttachmentStatusHandlerFunc func(ReadStorageV1VolumeAttachmentStatusParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ReadStorageV1VolumeAttachmentStatusHandlerFunc) Handle(params ReadStorageV1VolumeAttachmentStatusParams) middleware.Responder {
	return fn(params)
}

// ReadStorageV1VolumeAttachmentStatusHandler interface for that can handle valid read storage v1 volume attachment status params
type ReadStorageV1VolumeAttachmentStatusHandler interface {
	Handle(ReadStorageV1VolumeAttachmentStatusParams) middleware.Responder
}

// NewReadStorageV1VolumeAttachmentStatus creates a new http.Handler for the read storage v1 volume attachment status operation
func NewReadStorageV1VolumeAttachmentStatus(ctx *middleware.Context, handler ReadStorageV1VolumeAttachmentStatusHandler) *ReadStorageV1VolumeAttachmentStatus {
	return &ReadStorageV1VolumeAttachmentStatus{Context: ctx, Handler: handler}
}

/*ReadStorageV1VolumeAttachmentStatus swagger:route GET /apis/storage.k8s.io/v1/volumeattachments/{name}/status storage_v1 readStorageV1VolumeAttachmentStatus

read status of the specified VolumeAttachment

*/
type ReadStorageV1VolumeAttachmentStatus struct {
	Context *middleware.Context
	Handler ReadStorageV1VolumeAttachmentStatusHandler
}

func (o *ReadStorageV1VolumeAttachmentStatus) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewReadStorageV1VolumeAttachmentStatusParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
