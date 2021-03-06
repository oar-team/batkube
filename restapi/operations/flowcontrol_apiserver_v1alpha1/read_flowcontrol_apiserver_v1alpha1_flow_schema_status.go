// Code generated by go-swagger; DO NOT EDIT.

package flowcontrol_apiserver_v1alpha1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ReadFlowcontrolApiserverV1alpha1FlowSchemaStatusHandlerFunc turns a function with the right signature into a read flowcontrol apiserver v1alpha1 flow schema status handler
type ReadFlowcontrolApiserverV1alpha1FlowSchemaStatusHandlerFunc func(ReadFlowcontrolApiserverV1alpha1FlowSchemaStatusParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ReadFlowcontrolApiserverV1alpha1FlowSchemaStatusHandlerFunc) Handle(params ReadFlowcontrolApiserverV1alpha1FlowSchemaStatusParams) middleware.Responder {
	return fn(params)
}

// ReadFlowcontrolApiserverV1alpha1FlowSchemaStatusHandler interface for that can handle valid read flowcontrol apiserver v1alpha1 flow schema status params
type ReadFlowcontrolApiserverV1alpha1FlowSchemaStatusHandler interface {
	Handle(ReadFlowcontrolApiserverV1alpha1FlowSchemaStatusParams) middleware.Responder
}

// NewReadFlowcontrolApiserverV1alpha1FlowSchemaStatus creates a new http.Handler for the read flowcontrol apiserver v1alpha1 flow schema status operation
func NewReadFlowcontrolApiserverV1alpha1FlowSchemaStatus(ctx *middleware.Context, handler ReadFlowcontrolApiserverV1alpha1FlowSchemaStatusHandler) *ReadFlowcontrolApiserverV1alpha1FlowSchemaStatus {
	return &ReadFlowcontrolApiserverV1alpha1FlowSchemaStatus{Context: ctx, Handler: handler}
}

/*ReadFlowcontrolApiserverV1alpha1FlowSchemaStatus swagger:route GET /apis/flowcontrol.apiserver.k8s.io/v1alpha1/flowschemas/{name}/status flowcontrolApiserver_v1alpha1 readFlowcontrolApiserverV1alpha1FlowSchemaStatus

read status of the specified FlowSchema

*/
type ReadFlowcontrolApiserverV1alpha1FlowSchemaStatus struct {
	Context *middleware.Context
	Handler ReadFlowcontrolApiserverV1alpha1FlowSchemaStatusHandler
}

func (o *ReadFlowcontrolApiserverV1alpha1FlowSchemaStatus) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewReadFlowcontrolApiserverV1alpha1FlowSchemaStatusParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
