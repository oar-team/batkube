// Code generated by go-swagger; DO NOT EDIT.

package flowcontrol_apiserver_v1alpha1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// CreateFlowcontrolApiserverV1alpha1FlowSchemaHandlerFunc turns a function with the right signature into a create flowcontrol apiserver v1alpha1 flow schema handler
type CreateFlowcontrolApiserverV1alpha1FlowSchemaHandlerFunc func(CreateFlowcontrolApiserverV1alpha1FlowSchemaParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateFlowcontrolApiserverV1alpha1FlowSchemaHandlerFunc) Handle(params CreateFlowcontrolApiserverV1alpha1FlowSchemaParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// CreateFlowcontrolApiserverV1alpha1FlowSchemaHandler interface for that can handle valid create flowcontrol apiserver v1alpha1 flow schema params
type CreateFlowcontrolApiserverV1alpha1FlowSchemaHandler interface {
	Handle(CreateFlowcontrolApiserverV1alpha1FlowSchemaParams, interface{}) middleware.Responder
}

// NewCreateFlowcontrolApiserverV1alpha1FlowSchema creates a new http.Handler for the create flowcontrol apiserver v1alpha1 flow schema operation
func NewCreateFlowcontrolApiserverV1alpha1FlowSchema(ctx *middleware.Context, handler CreateFlowcontrolApiserverV1alpha1FlowSchemaHandler) *CreateFlowcontrolApiserverV1alpha1FlowSchema {
	return &CreateFlowcontrolApiserverV1alpha1FlowSchema{Context: ctx, Handler: handler}
}

/*CreateFlowcontrolApiserverV1alpha1FlowSchema swagger:route POST /apis/flowcontrol.apiserver.k8s.io/v1alpha1/flowschemas flowcontrolApiserver_v1alpha1 createFlowcontrolApiserverV1alpha1FlowSchema

create a FlowSchema

*/
type CreateFlowcontrolApiserverV1alpha1FlowSchema struct {
	Context *middleware.Context
	Handler CreateFlowcontrolApiserverV1alpha1FlowSchemaHandler
}

func (o *CreateFlowcontrolApiserverV1alpha1FlowSchema) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCreateFlowcontrolApiserverV1alpha1FlowSchemaParams()

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
