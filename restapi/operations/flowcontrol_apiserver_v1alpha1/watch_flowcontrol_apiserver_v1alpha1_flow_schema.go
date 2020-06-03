// Code generated by go-swagger; DO NOT EDIT.

package flowcontrol_apiserver_v1alpha1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// WatchFlowcontrolApiserverV1alpha1FlowSchemaHandlerFunc turns a function with the right signature into a watch flowcontrol apiserver v1alpha1 flow schema handler
type WatchFlowcontrolApiserverV1alpha1FlowSchemaHandlerFunc func(WatchFlowcontrolApiserverV1alpha1FlowSchemaParams) middleware.Responder

// Handle executing the request and returning a response
func (fn WatchFlowcontrolApiserverV1alpha1FlowSchemaHandlerFunc) Handle(params WatchFlowcontrolApiserverV1alpha1FlowSchemaParams) middleware.Responder {
	return fn(params)
}

// WatchFlowcontrolApiserverV1alpha1FlowSchemaHandler interface for that can handle valid watch flowcontrol apiserver v1alpha1 flow schema params
type WatchFlowcontrolApiserverV1alpha1FlowSchemaHandler interface {
	Handle(WatchFlowcontrolApiserverV1alpha1FlowSchemaParams) middleware.Responder
}

// NewWatchFlowcontrolApiserverV1alpha1FlowSchema creates a new http.Handler for the watch flowcontrol apiserver v1alpha1 flow schema operation
func NewWatchFlowcontrolApiserverV1alpha1FlowSchema(ctx *middleware.Context, handler WatchFlowcontrolApiserverV1alpha1FlowSchemaHandler) *WatchFlowcontrolApiserverV1alpha1FlowSchema {
	return &WatchFlowcontrolApiserverV1alpha1FlowSchema{Context: ctx, Handler: handler}
}

/*WatchFlowcontrolApiserverV1alpha1FlowSchema swagger:route GET /apis/flowcontrol.apiserver.k8s.io/v1alpha1/watch/flowschemas/{name} flowcontrolApiserver_v1alpha1 watchFlowcontrolApiserverV1alpha1FlowSchema

watch changes to an object of kind FlowSchema. deprecated: use the 'watch' parameter with a list operation instead, filtered to a single item with the 'fieldSelector' parameter.

*/
type WatchFlowcontrolApiserverV1alpha1FlowSchema struct {
	Context *middleware.Context
	Handler WatchFlowcontrolApiserverV1alpha1FlowSchemaHandler
}

func (o *WatchFlowcontrolApiserverV1alpha1FlowSchema) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewWatchFlowcontrolApiserverV1alpha1FlowSchemaParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
