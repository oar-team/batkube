// Code generated by go-swagger; DO NOT EDIT.

package storage_v1alpha1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ReplaceStorageV1alpha1VolumeAttachmentHandlerFunc turns a function with the right signature into a replace storage v1alpha1 volume attachment handler
type ReplaceStorageV1alpha1VolumeAttachmentHandlerFunc func(ReplaceStorageV1alpha1VolumeAttachmentParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ReplaceStorageV1alpha1VolumeAttachmentHandlerFunc) Handle(params ReplaceStorageV1alpha1VolumeAttachmentParams) middleware.Responder {
	return fn(params)
}

// ReplaceStorageV1alpha1VolumeAttachmentHandler interface for that can handle valid replace storage v1alpha1 volume attachment params
type ReplaceStorageV1alpha1VolumeAttachmentHandler interface {
	Handle(ReplaceStorageV1alpha1VolumeAttachmentParams) middleware.Responder
}

// NewReplaceStorageV1alpha1VolumeAttachment creates a new http.Handler for the replace storage v1alpha1 volume attachment operation
func NewReplaceStorageV1alpha1VolumeAttachment(ctx *middleware.Context, handler ReplaceStorageV1alpha1VolumeAttachmentHandler) *ReplaceStorageV1alpha1VolumeAttachment {
	return &ReplaceStorageV1alpha1VolumeAttachment{Context: ctx, Handler: handler}
}

/*ReplaceStorageV1alpha1VolumeAttachment swagger:route PUT /apis/storage.k8s.io/v1alpha1/volumeattachments/{name} storage_v1alpha1 replaceStorageV1alpha1VolumeAttachment

replace the specified VolumeAttachment

*/
type ReplaceStorageV1alpha1VolumeAttachment struct {
	Context *middleware.Context
	Handler ReplaceStorageV1alpha1VolumeAttachmentHandler
}

func (o *ReplaceStorageV1alpha1VolumeAttachment) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewReplaceStorageV1alpha1VolumeAttachmentParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
