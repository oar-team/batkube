// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// IoK8sApiextensionsApiserverPkgApisApiextensionsV1beta1CustomResourceColumnDefinition CustomResourceColumnDefinition specifies a column for server side printing.
//
// swagger:model io.k8s.apiextensions-apiserver.pkg.apis.apiextensions.v1beta1.CustomResourceColumnDefinition
type IoK8sApiextensionsApiserverPkgApisApiextensionsV1beta1CustomResourceColumnDefinition struct {

	// JSONPath is a simple JSON path (i.e. with array notation) which is evaluated against each custom resource to produce the value for this column.
	// Required: true
	JSONPath *string `json:"JSONPath"`

	// description is a human readable description of this column.
	Description string `json:"description,omitempty"`

	// format is an optional OpenAPI type definition for this column. The 'name' format is applied to the primary identifier column to assist in clients identifying column is the resource name. See https://github.com/OAI/OpenAPI-Specification/blob/master/versions/2.0.md#data-types for details.
	Format string `json:"format,omitempty"`

	// name is a human readable name for the column.
	// Required: true
	Name *string `json:"name"`

	// priority is an integer defining the relative importance of this column compared to others. Lower numbers are considered higher priority. Columns that may be omitted in limited space scenarios should be given a priority greater than 0.
	Priority int32 `json:"priority,omitempty"`

	// type is an OpenAPI type definition for this column. See https://github.com/OAI/OpenAPI-Specification/blob/master/versions/2.0.md#data-types for details.
	// Required: true
	Type *string `json:"type"`
}

// Validate validates this io k8s apiextensions apiserver pkg apis apiextensions v1beta1 custom resource column definition
func (m *IoK8sApiextensionsApiserverPkgApisApiextensionsV1beta1CustomResourceColumnDefinition) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateJSONPath(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sApiextensionsApiserverPkgApisApiextensionsV1beta1CustomResourceColumnDefinition) validateJSONPath(formats strfmt.Registry) error {

	if err := validate.Required("JSONPath", "body", m.JSONPath); err != nil {
		return err
	}

	return nil
}

func (m *IoK8sApiextensionsApiserverPkgApisApiextensionsV1beta1CustomResourceColumnDefinition) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *IoK8sApiextensionsApiserverPkgApisApiextensionsV1beta1CustomResourceColumnDefinition) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sApiextensionsApiserverPkgApisApiextensionsV1beta1CustomResourceColumnDefinition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sApiextensionsApiserverPkgApisApiextensionsV1beta1CustomResourceColumnDefinition) UnmarshalBinary(b []byte) error {
	var res IoK8sApiextensionsApiserverPkgApisApiextensionsV1beta1CustomResourceColumnDefinition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}