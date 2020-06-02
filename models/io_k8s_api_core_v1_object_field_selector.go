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

// IoK8sAPICoreV1ObjectFieldSelector ObjectFieldSelector selects an APIVersioned field of an object.
//
// swagger:model io.k8s.api.core.v1.ObjectFieldSelector
type IoK8sAPICoreV1ObjectFieldSelector struct {

	// Version of the schema the FieldPath is written in terms of, defaults to "v1".
	APIVersion string `json:"apiVersion,omitempty"`

	// Path of the field to select in the specified API version.
	// Required: true
	FieldPath *string `json:"fieldPath"`
}

// Validate validates this io k8s api core v1 object field selector
func (m *IoK8sAPICoreV1ObjectFieldSelector) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFieldPath(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPICoreV1ObjectFieldSelector) validateFieldPath(formats strfmt.Registry) error {

	if err := validate.Required("fieldPath", "body", m.FieldPath); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPICoreV1ObjectFieldSelector) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPICoreV1ObjectFieldSelector) UnmarshalBinary(b []byte) error {
	var res IoK8sAPICoreV1ObjectFieldSelector
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
