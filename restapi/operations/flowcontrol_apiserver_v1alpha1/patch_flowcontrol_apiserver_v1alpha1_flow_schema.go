// Code generated by go-swagger; DO NOT EDIT.

package flowcontrol_apiserver_v1alpha1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PatchFlowcontrolApiserverV1alpha1FlowSchemaHandlerFunc turns a function with the right signature into a patch flowcontrol apiserver v1alpha1 flow schema handler
type PatchFlowcontrolApiserverV1alpha1FlowSchemaHandlerFunc func(PatchFlowcontrolApiserverV1alpha1FlowSchemaParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PatchFlowcontrolApiserverV1alpha1FlowSchemaHandlerFunc) Handle(params PatchFlowcontrolApiserverV1alpha1FlowSchemaParams) middleware.Responder {
	return fn(params)
}

// PatchFlowcontrolApiserverV1alpha1FlowSchemaHandler interface for that can handle valid patch flowcontrol apiserver v1alpha1 flow schema params
type PatchFlowcontrolApiserverV1alpha1FlowSchemaHandler interface {
	Handle(PatchFlowcontrolApiserverV1alpha1FlowSchemaParams) middleware.Responder
}

// NewPatchFlowcontrolApiserverV1alpha1FlowSchema creates a new http.Handler for the patch flowcontrol apiserver v1alpha1 flow schema operation
func NewPatchFlowcontrolApiserverV1alpha1FlowSchema(ctx *middleware.Context, handler PatchFlowcontrolApiserverV1alpha1FlowSchemaHandler) *PatchFlowcontrolApiserverV1alpha1FlowSchema {
	return &PatchFlowcontrolApiserverV1alpha1FlowSchema{Context: ctx, Handler: handler}
}

/*PatchFlowcontrolApiserverV1alpha1FlowSchema swagger:route PATCH /apis/flowcontrol.apiserver.k8s.io/v1alpha1/flowschemas/{name} flowcontrolApiserver_v1alpha1 patchFlowcontrolApiserverV1alpha1FlowSchema

partially update the specified FlowSchema

*/
type PatchFlowcontrolApiserverV1alpha1FlowSchema struct {
	Context *middleware.Context
	Handler PatchFlowcontrolApiserverV1alpha1FlowSchemaHandler
}

func (o *PatchFlowcontrolApiserverV1alpha1FlowSchema) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPatchFlowcontrolApiserverV1alpha1FlowSchemaParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
