// Code generated by go-swagger; DO NOT EDIT.

package apps_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// CreateAppsV1NamespacedStatefulSetHandlerFunc turns a function with the right signature into a create apps v1 namespaced stateful set handler
type CreateAppsV1NamespacedStatefulSetHandlerFunc func(CreateAppsV1NamespacedStatefulSetParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateAppsV1NamespacedStatefulSetHandlerFunc) Handle(params CreateAppsV1NamespacedStatefulSetParams) middleware.Responder {
	return fn(params)
}

// CreateAppsV1NamespacedStatefulSetHandler interface for that can handle valid create apps v1 namespaced stateful set params
type CreateAppsV1NamespacedStatefulSetHandler interface {
	Handle(CreateAppsV1NamespacedStatefulSetParams) middleware.Responder
}

// NewCreateAppsV1NamespacedStatefulSet creates a new http.Handler for the create apps v1 namespaced stateful set operation
func NewCreateAppsV1NamespacedStatefulSet(ctx *middleware.Context, handler CreateAppsV1NamespacedStatefulSetHandler) *CreateAppsV1NamespacedStatefulSet {
	return &CreateAppsV1NamespacedStatefulSet{Context: ctx, Handler: handler}
}

/*CreateAppsV1NamespacedStatefulSet swagger:route POST /apis/apps/v1/namespaces/{namespace}/statefulsets apps_v1 createAppsV1NamespacedStatefulSet

create a StatefulSet

*/
type CreateAppsV1NamespacedStatefulSet struct {
	Context *middleware.Context
	Handler CreateAppsV1NamespacedStatefulSetHandler
}

func (o *CreateAppsV1NamespacedStatefulSet) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCreateAppsV1NamespacedStatefulSetParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
