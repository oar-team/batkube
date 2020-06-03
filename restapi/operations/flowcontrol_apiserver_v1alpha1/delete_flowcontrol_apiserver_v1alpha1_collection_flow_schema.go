// Code generated by go-swagger; DO NOT EDIT.

package flowcontrol_apiserver_v1alpha1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteFlowcontrolApiserverV1alpha1CollectionFlowSchemaHandlerFunc turns a function with the right signature into a delete flowcontrol apiserver v1alpha1 collection flow schema handler
type DeleteFlowcontrolApiserverV1alpha1CollectionFlowSchemaHandlerFunc func(DeleteFlowcontrolApiserverV1alpha1CollectionFlowSchemaParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteFlowcontrolApiserverV1alpha1CollectionFlowSchemaHandlerFunc) Handle(params DeleteFlowcontrolApiserverV1alpha1CollectionFlowSchemaParams) middleware.Responder {
	return fn(params)
}

// DeleteFlowcontrolApiserverV1alpha1CollectionFlowSchemaHandler interface for that can handle valid delete flowcontrol apiserver v1alpha1 collection flow schema params
type DeleteFlowcontrolApiserverV1alpha1CollectionFlowSchemaHandler interface {
	Handle(DeleteFlowcontrolApiserverV1alpha1CollectionFlowSchemaParams) middleware.Responder
}

// NewDeleteFlowcontrolApiserverV1alpha1CollectionFlowSchema creates a new http.Handler for the delete flowcontrol apiserver v1alpha1 collection flow schema operation
func NewDeleteFlowcontrolApiserverV1alpha1CollectionFlowSchema(ctx *middleware.Context, handler DeleteFlowcontrolApiserverV1alpha1CollectionFlowSchemaHandler) *DeleteFlowcontrolApiserverV1alpha1CollectionFlowSchema {
	return &DeleteFlowcontrolApiserverV1alpha1CollectionFlowSchema{Context: ctx, Handler: handler}
}

/*DeleteFlowcontrolApiserverV1alpha1CollectionFlowSchema swagger:route DELETE /apis/flowcontrol.apiserver.k8s.io/v1alpha1/flowschemas flowcontrolApiserver_v1alpha1 deleteFlowcontrolApiserverV1alpha1CollectionFlowSchema

delete collection of FlowSchema

*/
type DeleteFlowcontrolApiserverV1alpha1CollectionFlowSchema struct {
	Context *middleware.Context
	Handler DeleteFlowcontrolApiserverV1alpha1CollectionFlowSchemaHandler
}

func (o *DeleteFlowcontrolApiserverV1alpha1CollectionFlowSchema) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteFlowcontrolApiserverV1alpha1CollectionFlowSchemaParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
