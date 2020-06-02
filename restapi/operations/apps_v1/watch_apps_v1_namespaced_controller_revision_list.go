// Code generated by go-swagger; DO NOT EDIT.

package apps_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// WatchAppsV1NamespacedControllerRevisionListHandlerFunc turns a function with the right signature into a watch apps v1 namespaced controller revision list handler
type WatchAppsV1NamespacedControllerRevisionListHandlerFunc func(WatchAppsV1NamespacedControllerRevisionListParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn WatchAppsV1NamespacedControllerRevisionListHandlerFunc) Handle(params WatchAppsV1NamespacedControllerRevisionListParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// WatchAppsV1NamespacedControllerRevisionListHandler interface for that can handle valid watch apps v1 namespaced controller revision list params
type WatchAppsV1NamespacedControllerRevisionListHandler interface {
	Handle(WatchAppsV1NamespacedControllerRevisionListParams, interface{}) middleware.Responder
}

// NewWatchAppsV1NamespacedControllerRevisionList creates a new http.Handler for the watch apps v1 namespaced controller revision list operation
func NewWatchAppsV1NamespacedControllerRevisionList(ctx *middleware.Context, handler WatchAppsV1NamespacedControllerRevisionListHandler) *WatchAppsV1NamespacedControllerRevisionList {
	return &WatchAppsV1NamespacedControllerRevisionList{Context: ctx, Handler: handler}
}

/*WatchAppsV1NamespacedControllerRevisionList swagger:route GET /apis/apps/v1/watch/namespaces/{namespace}/controllerrevisions apps_v1 watchAppsV1NamespacedControllerRevisionList

watch individual changes to a list of ControllerRevision. deprecated: use the 'watch' parameter with a list operation instead.

*/
type WatchAppsV1NamespacedControllerRevisionList struct {
	Context *middleware.Context
	Handler WatchAppsV1NamespacedControllerRevisionListHandler
}

func (o *WatchAppsV1NamespacedControllerRevisionList) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewWatchAppsV1NamespacedControllerRevisionListParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
