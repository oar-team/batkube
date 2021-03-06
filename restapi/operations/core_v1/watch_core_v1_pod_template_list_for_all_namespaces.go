// Code generated by go-swagger; DO NOT EDIT.

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// WatchCoreV1PodTemplateListForAllNamespacesHandlerFunc turns a function with the right signature into a watch core v1 pod template list for all namespaces handler
type WatchCoreV1PodTemplateListForAllNamespacesHandlerFunc func(WatchCoreV1PodTemplateListForAllNamespacesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn WatchCoreV1PodTemplateListForAllNamespacesHandlerFunc) Handle(params WatchCoreV1PodTemplateListForAllNamespacesParams) middleware.Responder {
	return fn(params)
}

// WatchCoreV1PodTemplateListForAllNamespacesHandler interface for that can handle valid watch core v1 pod template list for all namespaces params
type WatchCoreV1PodTemplateListForAllNamespacesHandler interface {
	Handle(WatchCoreV1PodTemplateListForAllNamespacesParams) middleware.Responder
}

// NewWatchCoreV1PodTemplateListForAllNamespaces creates a new http.Handler for the watch core v1 pod template list for all namespaces operation
func NewWatchCoreV1PodTemplateListForAllNamespaces(ctx *middleware.Context, handler WatchCoreV1PodTemplateListForAllNamespacesHandler) *WatchCoreV1PodTemplateListForAllNamespaces {
	return &WatchCoreV1PodTemplateListForAllNamespaces{Context: ctx, Handler: handler}
}

/*WatchCoreV1PodTemplateListForAllNamespaces swagger:route GET /api/v1/watch/podtemplates core_v1 watchCoreV1PodTemplateListForAllNamespaces

watch individual changes to a list of PodTemplate. deprecated: use the 'watch' parameter with a list operation instead.

*/
type WatchCoreV1PodTemplateListForAllNamespaces struct {
	Context *middleware.Context
	Handler WatchCoreV1PodTemplateListForAllNamespacesHandler
}

func (o *WatchCoreV1PodTemplateListForAllNamespaces) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewWatchCoreV1PodTemplateListForAllNamespacesParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
