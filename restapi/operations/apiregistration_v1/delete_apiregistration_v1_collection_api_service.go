// Code generated by go-swagger; DO NOT EDIT.

package apiregistration_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteApiregistrationV1CollectionAPIServiceHandlerFunc turns a function with the right signature into a delete apiregistration v1 collection API service handler
type DeleteApiregistrationV1CollectionAPIServiceHandlerFunc func(DeleteApiregistrationV1CollectionAPIServiceParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteApiregistrationV1CollectionAPIServiceHandlerFunc) Handle(params DeleteApiregistrationV1CollectionAPIServiceParams) middleware.Responder {
	return fn(params)
}

// DeleteApiregistrationV1CollectionAPIServiceHandler interface for that can handle valid delete apiregistration v1 collection API service params
type DeleteApiregistrationV1CollectionAPIServiceHandler interface {
	Handle(DeleteApiregistrationV1CollectionAPIServiceParams) middleware.Responder
}

// NewDeleteApiregistrationV1CollectionAPIService creates a new http.Handler for the delete apiregistration v1 collection API service operation
func NewDeleteApiregistrationV1CollectionAPIService(ctx *middleware.Context, handler DeleteApiregistrationV1CollectionAPIServiceHandler) *DeleteApiregistrationV1CollectionAPIService {
	return &DeleteApiregistrationV1CollectionAPIService{Context: ctx, Handler: handler}
}

/*DeleteApiregistrationV1CollectionAPIService swagger:route DELETE /apis/apiregistration.k8s.io/v1/apiservices apiregistration_v1 deleteApiregistrationV1CollectionApiService

delete collection of APIService

*/
type DeleteApiregistrationV1CollectionAPIService struct {
	Context *middleware.Context
	Handler DeleteApiregistrationV1CollectionAPIServiceHandler
}

func (o *DeleteApiregistrationV1CollectionAPIService) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteApiregistrationV1CollectionAPIServiceParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
