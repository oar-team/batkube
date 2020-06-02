// Code generated by go-swagger; DO NOT EDIT.

package storage_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewReadStorageV1beta1CSINodeParams creates a new ReadStorageV1beta1CSINodeParams object
// no default values defined in spec.
func NewReadStorageV1beta1CSINodeParams() ReadStorageV1beta1CSINodeParams {

	return ReadStorageV1beta1CSINodeParams{}
}

// ReadStorageV1beta1CSINodeParams contains all the bound params for the read storage v1beta1 c s i node operation
// typically these are obtained from a http.Request
//
// swagger:parameters readStorageV1beta1CSINode
type ReadStorageV1beta1CSINodeParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Should the export be exact.  Exact export maintains cluster-specific fields like 'Namespace'. Deprecated. Planned for removal in 1.18.
	  Unique: true
	  In: query
	*/
	Exact *bool
	/*Should this value be exported.  Export strips fields that a user can not specify. Deprecated. Planned for removal in 1.18.
	  Unique: true
	  In: query
	*/
	Export *bool
	/*name of the CSINode
	  Required: true
	  Unique: true
	  In: path
	*/
	Name string
	/*If 'true', then the output is pretty printed.
	  Unique: true
	  In: query
	*/
	Pretty *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewReadStorageV1beta1CSINodeParams() beforehand.
func (o *ReadStorageV1beta1CSINodeParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qExact, qhkExact, _ := qs.GetOK("exact")
	if err := o.bindExact(qExact, qhkExact, route.Formats); err != nil {
		res = append(res, err)
	}

	qExport, qhkExport, _ := qs.GetOK("export")
	if err := o.bindExport(qExport, qhkExport, route.Formats); err != nil {
		res = append(res, err)
	}

	rName, rhkName, _ := route.Params.GetOK("name")
	if err := o.bindName(rName, rhkName, route.Formats); err != nil {
		res = append(res, err)
	}

	qPretty, qhkPretty, _ := qs.GetOK("pretty")
	if err := o.bindPretty(qPretty, qhkPretty, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindExact binds and validates parameter Exact from query.
func (o *ReadStorageV1beta1CSINodeParams) bindExact(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertBool(raw)
	if err != nil {
		return errors.InvalidType("exact", "query", "bool", raw)
	}
	o.Exact = &value

	if err := o.validateExact(formats); err != nil {
		return err
	}

	return nil
}

// validateExact carries on validations for parameter Exact
func (o *ReadStorageV1beta1CSINodeParams) validateExact(formats strfmt.Registry) error {

	return nil
}

// bindExport binds and validates parameter Export from query.
func (o *ReadStorageV1beta1CSINodeParams) bindExport(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertBool(raw)
	if err != nil {
		return errors.InvalidType("export", "query", "bool", raw)
	}
	o.Export = &value

	if err := o.validateExport(formats); err != nil {
		return err
	}

	return nil
}

// validateExport carries on validations for parameter Export
func (o *ReadStorageV1beta1CSINodeParams) validateExport(formats strfmt.Registry) error {

	return nil
}

// bindName binds and validates parameter Name from path.
func (o *ReadStorageV1beta1CSINodeParams) bindName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.Name = raw

	if err := o.validateName(formats); err != nil {
		return err
	}

	return nil
}

// validateName carries on validations for parameter Name
func (o *ReadStorageV1beta1CSINodeParams) validateName(formats strfmt.Registry) error {

	return nil
}

// bindPretty binds and validates parameter Pretty from query.
func (o *ReadStorageV1beta1CSINodeParams) bindPretty(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Pretty = &raw

	if err := o.validatePretty(formats); err != nil {
		return err
	}

	return nil
}

// validatePretty carries on validations for parameter Pretty
func (o *ReadStorageV1beta1CSINodeParams) validatePretty(formats strfmt.Registry) error {

	return nil
}
