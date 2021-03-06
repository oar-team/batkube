// Code generated by go-swagger; DO NOT EDIT.

package admissionregistration_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetAdmissionregistrationV1APIResourcesHandlerFunc turns a function with the right signature into a get admissionregistration v1 API resources handler
type GetAdmissionregistrationV1APIResourcesHandlerFunc func(GetAdmissionregistrationV1APIResourcesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetAdmissionregistrationV1APIResourcesHandlerFunc) Handle(params GetAdmissionregistrationV1APIResourcesParams) middleware.Responder {
	return fn(params)
}

// GetAdmissionregistrationV1APIResourcesHandler interface for that can handle valid get admissionregistration v1 API resources params
type GetAdmissionregistrationV1APIResourcesHandler interface {
	Handle(GetAdmissionregistrationV1APIResourcesParams) middleware.Responder
}

// NewGetAdmissionregistrationV1APIResources creates a new http.Handler for the get admissionregistration v1 API resources operation
func NewGetAdmissionregistrationV1APIResources(ctx *middleware.Context, handler GetAdmissionregistrationV1APIResourcesHandler) *GetAdmissionregistrationV1APIResources {
	return &GetAdmissionregistrationV1APIResources{Context: ctx, Handler: handler}
}

/*GetAdmissionregistrationV1APIResources swagger:route GET /apis/admissionregistration.k8s.io/v1/ admissionregistration_v1 getAdmissionregistrationV1ApiResources

get available resources

*/
type GetAdmissionregistrationV1APIResources struct {
	Context *middleware.Context
	Handler GetAdmissionregistrationV1APIResourcesHandler
}

func (o *GetAdmissionregistrationV1APIResources) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetAdmissionregistrationV1APIResourcesParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
