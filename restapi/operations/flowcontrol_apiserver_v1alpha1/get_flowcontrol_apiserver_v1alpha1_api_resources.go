// Code generated by go-swagger; DO NOT EDIT.

package flowcontrol_apiserver_v1alpha1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetFlowcontrolApiserverV1alpha1APIResourcesHandlerFunc turns a function with the right signature into a get flowcontrol apiserver v1alpha1 API resources handler
type GetFlowcontrolApiserverV1alpha1APIResourcesHandlerFunc func(GetFlowcontrolApiserverV1alpha1APIResourcesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetFlowcontrolApiserverV1alpha1APIResourcesHandlerFunc) Handle(params GetFlowcontrolApiserverV1alpha1APIResourcesParams) middleware.Responder {
	return fn(params)
}

// GetFlowcontrolApiserverV1alpha1APIResourcesHandler interface for that can handle valid get flowcontrol apiserver v1alpha1 API resources params
type GetFlowcontrolApiserverV1alpha1APIResourcesHandler interface {
	Handle(GetFlowcontrolApiserverV1alpha1APIResourcesParams) middleware.Responder
}

// NewGetFlowcontrolApiserverV1alpha1APIResources creates a new http.Handler for the get flowcontrol apiserver v1alpha1 API resources operation
func NewGetFlowcontrolApiserverV1alpha1APIResources(ctx *middleware.Context, handler GetFlowcontrolApiserverV1alpha1APIResourcesHandler) *GetFlowcontrolApiserverV1alpha1APIResources {
	return &GetFlowcontrolApiserverV1alpha1APIResources{Context: ctx, Handler: handler}
}

/*GetFlowcontrolApiserverV1alpha1APIResources swagger:route GET /apis/flowcontrol.apiserver.k8s.io/v1alpha1/ flowcontrolApiserver_v1alpha1 getFlowcontrolApiserverV1alpha1ApiResources

get available resources

*/
type GetFlowcontrolApiserverV1alpha1APIResources struct {
	Context *middleware.Context
	Handler GetFlowcontrolApiserverV1alpha1APIResourcesHandler
}

func (o *GetFlowcontrolApiserverV1alpha1APIResources) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetFlowcontrolApiserverV1alpha1APIResourcesParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
