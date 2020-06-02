// Code generated by go-swagger; DO NOT EDIT.

package certificates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
)

// NewGetCertificatesAPIGroupParams creates a new GetCertificatesAPIGroupParams object
// no default values defined in spec.
func NewGetCertificatesAPIGroupParams() GetCertificatesAPIGroupParams {

	return GetCertificatesAPIGroupParams{}
}

// GetCertificatesAPIGroupParams contains all the bound params for the get certificates API group operation
// typically these are obtained from a http.Request
//
// swagger:parameters getCertificatesAPIGroup
type GetCertificatesAPIGroupParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetCertificatesAPIGroupParams() beforehand.
func (o *GetCertificatesAPIGroupParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
