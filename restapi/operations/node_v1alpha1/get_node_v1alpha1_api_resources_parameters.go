// Code generated by go-swagger; DO NOT EDIT.

package node_v1alpha1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
)

// NewGetNodeV1alpha1APIResourcesParams creates a new GetNodeV1alpha1APIResourcesParams object
// no default values defined in spec.
func NewGetNodeV1alpha1APIResourcesParams() GetNodeV1alpha1APIResourcesParams {

	return GetNodeV1alpha1APIResourcesParams{}
}

// GetNodeV1alpha1APIResourcesParams contains all the bound params for the get node v1alpha1 API resources operation
// typically these are obtained from a http.Request
//
// swagger:parameters getNodeV1alpha1APIResources
type GetNodeV1alpha1APIResourcesParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetNodeV1alpha1APIResourcesParams() beforehand.
func (o *GetNodeV1alpha1APIResourcesParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
