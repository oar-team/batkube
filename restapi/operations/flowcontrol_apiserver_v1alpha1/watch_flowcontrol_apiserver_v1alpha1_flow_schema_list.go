// Code generated by go-swagger; DO NOT EDIT.

package flowcontrol_apiserver_v1alpha1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// WatchFlowcontrolApiserverV1alpha1FlowSchemaListHandlerFunc turns a function with the right signature into a watch flowcontrol apiserver v1alpha1 flow schema list handler
type WatchFlowcontrolApiserverV1alpha1FlowSchemaListHandlerFunc func(WatchFlowcontrolApiserverV1alpha1FlowSchemaListParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn WatchFlowcontrolApiserverV1alpha1FlowSchemaListHandlerFunc) Handle(params WatchFlowcontrolApiserverV1alpha1FlowSchemaListParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// WatchFlowcontrolApiserverV1alpha1FlowSchemaListHandler interface for that can handle valid watch flowcontrol apiserver v1alpha1 flow schema list params
type WatchFlowcontrolApiserverV1alpha1FlowSchemaListHandler interface {
	Handle(WatchFlowcontrolApiserverV1alpha1FlowSchemaListParams, interface{}) middleware.Responder
}

// NewWatchFlowcontrolApiserverV1alpha1FlowSchemaList creates a new http.Handler for the watch flowcontrol apiserver v1alpha1 flow schema list operation
func NewWatchFlowcontrolApiserverV1alpha1FlowSchemaList(ctx *middleware.Context, handler WatchFlowcontrolApiserverV1alpha1FlowSchemaListHandler) *WatchFlowcontrolApiserverV1alpha1FlowSchemaList {
	return &WatchFlowcontrolApiserverV1alpha1FlowSchemaList{Context: ctx, Handler: handler}
}

/*WatchFlowcontrolApiserverV1alpha1FlowSchemaList swagger:route GET /apis/flowcontrol.apiserver.k8s.io/v1alpha1/watch/flowschemas flowcontrolApiserver_v1alpha1 watchFlowcontrolApiserverV1alpha1FlowSchemaList

watch individual changes to a list of FlowSchema. deprecated: use the 'watch' parameter with a list operation instead.

*/
type WatchFlowcontrolApiserverV1alpha1FlowSchemaList struct {
	Context *middleware.Context
	Handler WatchFlowcontrolApiserverV1alpha1FlowSchemaListHandler
}

func (o *WatchFlowcontrolApiserverV1alpha1FlowSchemaList) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewWatchFlowcontrolApiserverV1alpha1FlowSchemaListParams()

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
