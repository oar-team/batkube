// Code generated by go-swagger; DO NOT EDIT.

package flowcontrol_apiserver_v1alpha1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
)

// NewGetFlowcontrolApiserverV1alpha1APIResourcesParams creates a new GetFlowcontrolApiserverV1alpha1APIResourcesParams object
// no default values defined in spec.
func NewGetFlowcontrolApiserverV1alpha1APIResourcesParams() GetFlowcontrolApiserverV1alpha1APIResourcesParams {

	return GetFlowcontrolApiserverV1alpha1APIResourcesParams{}
}

// GetFlowcontrolApiserverV1alpha1APIResourcesParams contains all the bound params for the get flowcontrol apiserver v1alpha1 API resources operation
// typically these are obtained from a http.Request
//
// swagger:parameters getFlowcontrolApiserverV1alpha1APIResources
type GetFlowcontrolApiserverV1alpha1APIResourcesParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetFlowcontrolApiserverV1alpha1APIResourcesParams() beforehand.
func (o *GetFlowcontrolApiserverV1alpha1APIResourcesParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
