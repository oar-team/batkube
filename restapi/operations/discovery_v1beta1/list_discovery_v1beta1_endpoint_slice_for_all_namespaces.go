// Code generated by go-swagger; DO NOT EDIT.

package discovery_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ListDiscoveryV1beta1EndpointSliceForAllNamespacesHandlerFunc turns a function with the right signature into a list discovery v1beta1 endpoint slice for all namespaces handler
type ListDiscoveryV1beta1EndpointSliceForAllNamespacesHandlerFunc func(ListDiscoveryV1beta1EndpointSliceForAllNamespacesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ListDiscoveryV1beta1EndpointSliceForAllNamespacesHandlerFunc) Handle(params ListDiscoveryV1beta1EndpointSliceForAllNamespacesParams) middleware.Responder {
	return fn(params)
}

// ListDiscoveryV1beta1EndpointSliceForAllNamespacesHandler interface for that can handle valid list discovery v1beta1 endpoint slice for all namespaces params
type ListDiscoveryV1beta1EndpointSliceForAllNamespacesHandler interface {
	Handle(ListDiscoveryV1beta1EndpointSliceForAllNamespacesParams) middleware.Responder
}

// NewListDiscoveryV1beta1EndpointSliceForAllNamespaces creates a new http.Handler for the list discovery v1beta1 endpoint slice for all namespaces operation
func NewListDiscoveryV1beta1EndpointSliceForAllNamespaces(ctx *middleware.Context, handler ListDiscoveryV1beta1EndpointSliceForAllNamespacesHandler) *ListDiscoveryV1beta1EndpointSliceForAllNamespaces {
	return &ListDiscoveryV1beta1EndpointSliceForAllNamespaces{Context: ctx, Handler: handler}
}

/*ListDiscoveryV1beta1EndpointSliceForAllNamespaces swagger:route GET /apis/discovery.k8s.io/v1beta1/endpointslices discovery_v1beta1 listDiscoveryV1beta1EndpointSliceForAllNamespaces

list or watch objects of kind EndpointSlice

*/
type ListDiscoveryV1beta1EndpointSliceForAllNamespaces struct {
	Context *middleware.Context
	Handler ListDiscoveryV1beta1EndpointSliceForAllNamespacesHandler
}

func (o *ListDiscoveryV1beta1EndpointSliceForAllNamespaces) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewListDiscoveryV1beta1EndpointSliceForAllNamespacesParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
