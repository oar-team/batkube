// Code generated by go-swagger; DO NOT EDIT.

package flowcontrol_apiserver_v1alpha1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ReplaceFlowcontrolApiserverV1alpha1PriorityLevelConfigurationStatusHandlerFunc turns a function with the right signature into a replace flowcontrol apiserver v1alpha1 priority level configuration status handler
type ReplaceFlowcontrolApiserverV1alpha1PriorityLevelConfigurationStatusHandlerFunc func(ReplaceFlowcontrolApiserverV1alpha1PriorityLevelConfigurationStatusParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ReplaceFlowcontrolApiserverV1alpha1PriorityLevelConfigurationStatusHandlerFunc) Handle(params ReplaceFlowcontrolApiserverV1alpha1PriorityLevelConfigurationStatusParams) middleware.Responder {
	return fn(params)
}

// ReplaceFlowcontrolApiserverV1alpha1PriorityLevelConfigurationStatusHandler interface for that can handle valid replace flowcontrol apiserver v1alpha1 priority level configuration status params
type ReplaceFlowcontrolApiserverV1alpha1PriorityLevelConfigurationStatusHandler interface {
	Handle(ReplaceFlowcontrolApiserverV1alpha1PriorityLevelConfigurationStatusParams) middleware.Responder
}

// NewReplaceFlowcontrolApiserverV1alpha1PriorityLevelConfigurationStatus creates a new http.Handler for the replace flowcontrol apiserver v1alpha1 priority level configuration status operation
func NewReplaceFlowcontrolApiserverV1alpha1PriorityLevelConfigurationStatus(ctx *middleware.Context, handler ReplaceFlowcontrolApiserverV1alpha1PriorityLevelConfigurationStatusHandler) *ReplaceFlowcontrolApiserverV1alpha1PriorityLevelConfigurationStatus {
	return &ReplaceFlowcontrolApiserverV1alpha1PriorityLevelConfigurationStatus{Context: ctx, Handler: handler}
}

/*ReplaceFlowcontrolApiserverV1alpha1PriorityLevelConfigurationStatus swagger:route PUT /apis/flowcontrol.apiserver.k8s.io/v1alpha1/prioritylevelconfigurations/{name}/status flowcontrolApiserver_v1alpha1 replaceFlowcontrolApiserverV1alpha1PriorityLevelConfigurationStatus

replace status of the specified PriorityLevelConfiguration

*/
type ReplaceFlowcontrolApiserverV1alpha1PriorityLevelConfigurationStatus struct {
	Context *middleware.Context
	Handler ReplaceFlowcontrolApiserverV1alpha1PriorityLevelConfigurationStatusHandler
}

func (o *ReplaceFlowcontrolApiserverV1alpha1PriorityLevelConfigurationStatus) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewReplaceFlowcontrolApiserverV1alpha1PriorityLevelConfigurationStatusParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
