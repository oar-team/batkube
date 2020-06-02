// Code generated by go-swagger; DO NOT EDIT.

package storage_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
)

// NewGetStorageV1beta1APIResourcesParams creates a new GetStorageV1beta1APIResourcesParams object
// no default values defined in spec.
func NewGetStorageV1beta1APIResourcesParams() GetStorageV1beta1APIResourcesParams {

	return GetStorageV1beta1APIResourcesParams{}
}

// GetStorageV1beta1APIResourcesParams contains all the bound params for the get storage v1beta1 API resources operation
// typically these are obtained from a http.Request
//
// swagger:parameters getStorageV1beta1APIResources
type GetStorageV1beta1APIResourcesParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetStorageV1beta1APIResourcesParams() beforehand.
func (o *GetStorageV1beta1APIResourcesParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
