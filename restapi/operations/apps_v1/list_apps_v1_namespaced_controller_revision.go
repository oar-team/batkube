// Code generated by go-swagger; DO NOT EDIT.

package apps_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ListAppsV1NamespacedControllerRevisionHandlerFunc turns a function with the right signature into a list apps v1 namespaced controller revision handler
type ListAppsV1NamespacedControllerRevisionHandlerFunc func(ListAppsV1NamespacedControllerRevisionParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ListAppsV1NamespacedControllerRevisionHandlerFunc) Handle(params ListAppsV1NamespacedControllerRevisionParams) middleware.Responder {
	return fn(params)
}

// ListAppsV1NamespacedControllerRevisionHandler interface for that can handle valid list apps v1 namespaced controller revision params
type ListAppsV1NamespacedControllerRevisionHandler interface {
	Handle(ListAppsV1NamespacedControllerRevisionParams) middleware.Responder
}

// NewListAppsV1NamespacedControllerRevision creates a new http.Handler for the list apps v1 namespaced controller revision operation
func NewListAppsV1NamespacedControllerRevision(ctx *middleware.Context, handler ListAppsV1NamespacedControllerRevisionHandler) *ListAppsV1NamespacedControllerRevision {
	return &ListAppsV1NamespacedControllerRevision{Context: ctx, Handler: handler}
}

/*ListAppsV1NamespacedControllerRevision swagger:route GET /apis/apps/v1/namespaces/{namespace}/controllerrevisions apps_v1 listAppsV1NamespacedControllerRevision

list or watch objects of kind ControllerRevision

*/
type ListAppsV1NamespacedControllerRevision struct {
	Context *middleware.Context
	Handler ListAppsV1NamespacedControllerRevisionHandler
}

func (o *ListAppsV1NamespacedControllerRevision) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewListAppsV1NamespacedControllerRevisionParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
