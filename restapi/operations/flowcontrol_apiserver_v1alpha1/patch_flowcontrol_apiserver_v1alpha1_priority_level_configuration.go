// Code generated by go-swagger; DO NOT EDIT.

package flowcontrol_apiserver_v1alpha1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationHandlerFunc turns a function with the right signature into a patch flowcontrol apiserver v1alpha1 priority level configuration handler
type PatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationHandlerFunc func(PatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationHandlerFunc) Handle(params PatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationParams) middleware.Responder {
	return fn(params)
}

// PatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationHandler interface for that can handle valid patch flowcontrol apiserver v1alpha1 priority level configuration params
type PatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationHandler interface {
	Handle(PatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationParams) middleware.Responder
}

// NewPatchFlowcontrolApiserverV1alpha1PriorityLevelConfiguration creates a new http.Handler for the patch flowcontrol apiserver v1alpha1 priority level configuration operation
func NewPatchFlowcontrolApiserverV1alpha1PriorityLevelConfiguration(ctx *middleware.Context, handler PatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationHandler) *PatchFlowcontrolApiserverV1alpha1PriorityLevelConfiguration {
	return &PatchFlowcontrolApiserverV1alpha1PriorityLevelConfiguration{Context: ctx, Handler: handler}
}

/*PatchFlowcontrolApiserverV1alpha1PriorityLevelConfiguration swagger:route PATCH /apis/flowcontrol.apiserver.k8s.io/v1alpha1/prioritylevelconfigurations/{name} flowcontrolApiserver_v1alpha1 patchFlowcontrolApiserverV1alpha1PriorityLevelConfiguration

partially update the specified PriorityLevelConfiguration

*/
type PatchFlowcontrolApiserverV1alpha1PriorityLevelConfiguration struct {
	Context *middleware.Context
	Handler PatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationHandler
}

func (o *PatchFlowcontrolApiserverV1alpha1PriorityLevelConfiguration) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
