// Code generated by go-swagger; DO NOT EDIT.

package apps_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// CreateAppsV1NamespacedControllerRevisionHandlerFunc turns a function with the right signature into a create apps v1 namespaced controller revision handler
type CreateAppsV1NamespacedControllerRevisionHandlerFunc func(CreateAppsV1NamespacedControllerRevisionParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateAppsV1NamespacedControllerRevisionHandlerFunc) Handle(params CreateAppsV1NamespacedControllerRevisionParams) middleware.Responder {
	return fn(params)
}

// CreateAppsV1NamespacedControllerRevisionHandler interface for that can handle valid create apps v1 namespaced controller revision params
type CreateAppsV1NamespacedControllerRevisionHandler interface {
	Handle(CreateAppsV1NamespacedControllerRevisionParams) middleware.Responder
}

// NewCreateAppsV1NamespacedControllerRevision creates a new http.Handler for the create apps v1 namespaced controller revision operation
func NewCreateAppsV1NamespacedControllerRevision(ctx *middleware.Context, handler CreateAppsV1NamespacedControllerRevisionHandler) *CreateAppsV1NamespacedControllerRevision {
	return &CreateAppsV1NamespacedControllerRevision{Context: ctx, Handler: handler}
}

/*CreateAppsV1NamespacedControllerRevision swagger:route POST /apis/apps/v1/namespaces/{namespace}/controllerrevisions apps_v1 createAppsV1NamespacedControllerRevision

create a ControllerRevision

*/
type CreateAppsV1NamespacedControllerRevision struct {
	Context *middleware.Context
	Handler CreateAppsV1NamespacedControllerRevisionHandler
}

func (o *CreateAppsV1NamespacedControllerRevision) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCreateAppsV1NamespacedControllerRevisionParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
