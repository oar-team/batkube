// Code generated by go-swagger; DO NOT EDIT.

package apiregistration_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ReplaceApiregistrationV1beta1APIServiceHandlerFunc turns a function with the right signature into a replace apiregistration v1beta1 API service handler
type ReplaceApiregistrationV1beta1APIServiceHandlerFunc func(ReplaceApiregistrationV1beta1APIServiceParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ReplaceApiregistrationV1beta1APIServiceHandlerFunc) Handle(params ReplaceApiregistrationV1beta1APIServiceParams) middleware.Responder {
	return fn(params)
}

// ReplaceApiregistrationV1beta1APIServiceHandler interface for that can handle valid replace apiregistration v1beta1 API service params
type ReplaceApiregistrationV1beta1APIServiceHandler interface {
	Handle(ReplaceApiregistrationV1beta1APIServiceParams) middleware.Responder
}

// NewReplaceApiregistrationV1beta1APIService creates a new http.Handler for the replace apiregistration v1beta1 API service operation
func NewReplaceApiregistrationV1beta1APIService(ctx *middleware.Context, handler ReplaceApiregistrationV1beta1APIServiceHandler) *ReplaceApiregistrationV1beta1APIService {
	return &ReplaceApiregistrationV1beta1APIService{Context: ctx, Handler: handler}
}

/*ReplaceApiregistrationV1beta1APIService swagger:route PUT /apis/apiregistration.k8s.io/v1beta1/apiservices/{name} apiregistration_v1beta1 replaceApiregistrationV1beta1ApiService

replace the specified APIService

*/
type ReplaceApiregistrationV1beta1APIService struct {
	Context *middleware.Context
	Handler ReplaceApiregistrationV1beta1APIServiceHandler
}

func (o *ReplaceApiregistrationV1beta1APIService) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewReplaceApiregistrationV1beta1APIServiceParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
