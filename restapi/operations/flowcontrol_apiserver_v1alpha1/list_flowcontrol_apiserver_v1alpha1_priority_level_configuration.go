// Code generated by go-swagger; DO NOT EDIT.

package flowcontrol_apiserver_v1alpha1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ListFlowcontrolApiserverV1alpha1PriorityLevelConfigurationHandlerFunc turns a function with the right signature into a list flowcontrol apiserver v1alpha1 priority level configuration handler
type ListFlowcontrolApiserverV1alpha1PriorityLevelConfigurationHandlerFunc func(ListFlowcontrolApiserverV1alpha1PriorityLevelConfigurationParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ListFlowcontrolApiserverV1alpha1PriorityLevelConfigurationHandlerFunc) Handle(params ListFlowcontrolApiserverV1alpha1PriorityLevelConfigurationParams) middleware.Responder {
	return fn(params)
}

// ListFlowcontrolApiserverV1alpha1PriorityLevelConfigurationHandler interface for that can handle valid list flowcontrol apiserver v1alpha1 priority level configuration params
type ListFlowcontrolApiserverV1alpha1PriorityLevelConfigurationHandler interface {
	Handle(ListFlowcontrolApiserverV1alpha1PriorityLevelConfigurationParams) middleware.Responder
}

// NewListFlowcontrolApiserverV1alpha1PriorityLevelConfiguration creates a new http.Handler for the list flowcontrol apiserver v1alpha1 priority level configuration operation
func NewListFlowcontrolApiserverV1alpha1PriorityLevelConfiguration(ctx *middleware.Context, handler ListFlowcontrolApiserverV1alpha1PriorityLevelConfigurationHandler) *ListFlowcontrolApiserverV1alpha1PriorityLevelConfiguration {
	return &ListFlowcontrolApiserverV1alpha1PriorityLevelConfiguration{Context: ctx, Handler: handler}
}

/*ListFlowcontrolApiserverV1alpha1PriorityLevelConfiguration swagger:route GET /apis/flowcontrol.apiserver.k8s.io/v1alpha1/prioritylevelconfigurations flowcontrolApiserver_v1alpha1 listFlowcontrolApiserverV1alpha1PriorityLevelConfiguration

list or watch objects of kind PriorityLevelConfiguration

*/
type ListFlowcontrolApiserverV1alpha1PriorityLevelConfiguration struct {
	Context *middleware.Context
	Handler ListFlowcontrolApiserverV1alpha1PriorityLevelConfigurationHandler
}

func (o *ListFlowcontrolApiserverV1alpha1PriorityLevelConfiguration) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewListFlowcontrolApiserverV1alpha1PriorityLevelConfigurationParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
