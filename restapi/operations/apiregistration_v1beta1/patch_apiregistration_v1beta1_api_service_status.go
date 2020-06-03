// Code generated by go-swagger; DO NOT EDIT.

package apiregistration_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PatchApiregistrationV1beta1APIServiceStatusHandlerFunc turns a function with the right signature into a patch apiregistration v1beta1 API service status handler
type PatchApiregistrationV1beta1APIServiceStatusHandlerFunc func(PatchApiregistrationV1beta1APIServiceStatusParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PatchApiregistrationV1beta1APIServiceStatusHandlerFunc) Handle(params PatchApiregistrationV1beta1APIServiceStatusParams) middleware.Responder {
	return fn(params)
}

// PatchApiregistrationV1beta1APIServiceStatusHandler interface for that can handle valid patch apiregistration v1beta1 API service status params
type PatchApiregistrationV1beta1APIServiceStatusHandler interface {
	Handle(PatchApiregistrationV1beta1APIServiceStatusParams) middleware.Responder
}

// NewPatchApiregistrationV1beta1APIServiceStatus creates a new http.Handler for the patch apiregistration v1beta1 API service status operation
func NewPatchApiregistrationV1beta1APIServiceStatus(ctx *middleware.Context, handler PatchApiregistrationV1beta1APIServiceStatusHandler) *PatchApiregistrationV1beta1APIServiceStatus {
	return &PatchApiregistrationV1beta1APIServiceStatus{Context: ctx, Handler: handler}
}

/*PatchApiregistrationV1beta1APIServiceStatus swagger:route PATCH /apis/apiregistration.k8s.io/v1beta1/apiservices/{name}/status apiregistration_v1beta1 patchApiregistrationV1beta1ApiServiceStatus

partially update status of the specified APIService

*/
type PatchApiregistrationV1beta1APIServiceStatus struct {
	Context *middleware.Context
	Handler PatchApiregistrationV1beta1APIServiceStatusHandler
}

func (o *PatchApiregistrationV1beta1APIServiceStatus) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPatchApiregistrationV1beta1APIServiceStatusParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
