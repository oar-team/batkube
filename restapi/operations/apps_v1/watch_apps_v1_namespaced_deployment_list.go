// Code generated by go-swagger; DO NOT EDIT.

package apps_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// WatchAppsV1NamespacedDeploymentListHandlerFunc turns a function with the right signature into a watch apps v1 namespaced deployment list handler
type WatchAppsV1NamespacedDeploymentListHandlerFunc func(WatchAppsV1NamespacedDeploymentListParams) middleware.Responder

// Handle executing the request and returning a response
func (fn WatchAppsV1NamespacedDeploymentListHandlerFunc) Handle(params WatchAppsV1NamespacedDeploymentListParams) middleware.Responder {
	return fn(params)
}

// WatchAppsV1NamespacedDeploymentListHandler interface for that can handle valid watch apps v1 namespaced deployment list params
type WatchAppsV1NamespacedDeploymentListHandler interface {
	Handle(WatchAppsV1NamespacedDeploymentListParams) middleware.Responder
}

// NewWatchAppsV1NamespacedDeploymentList creates a new http.Handler for the watch apps v1 namespaced deployment list operation
func NewWatchAppsV1NamespacedDeploymentList(ctx *middleware.Context, handler WatchAppsV1NamespacedDeploymentListHandler) *WatchAppsV1NamespacedDeploymentList {
	return &WatchAppsV1NamespacedDeploymentList{Context: ctx, Handler: handler}
}

/*WatchAppsV1NamespacedDeploymentList swagger:route GET /apis/apps/v1/watch/namespaces/{namespace}/deployments apps_v1 watchAppsV1NamespacedDeploymentList

watch individual changes to a list of Deployment. deprecated: use the 'watch' parameter with a list operation instead.

*/
type WatchAppsV1NamespacedDeploymentList struct {
	Context *middleware.Context
	Handler WatchAppsV1NamespacedDeploymentListHandler
}

func (o *WatchAppsV1NamespacedDeploymentList) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewWatchAppsV1NamespacedDeploymentListParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
