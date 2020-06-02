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

// IoK8sAPICoreV1SecretKeySelector SecretKeySelector selects a key of a Secret.
//
// swagger:model io.k8s.api.core.v1.SecretKeySelector
type IoK8sAPICoreV1SecretKeySelector struct {

	// The key of the secret to select from.  Must be a valid secret key.
	// Required: true
	Key *string `json:"key"`

	// Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	Name string `json:"name,omitempty"`

	// Specify whether the Secret or its key must be defined
	Optional bool `json:"optional,omitempty"`
}

// Validate validates this io k8s api core v1 secret key selector
func (m *IoK8sAPICoreV1SecretKeySelector) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateKey(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPICoreV1SecretKeySelector) validateKey(formats strfmt.Registry) error {

	if err := validate.Required("key", "body", m.Key); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPICoreV1SecretKeySelector) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPICoreV1SecretKeySelector) UnmarshalBinary(b []byte) error {
	var res IoK8sAPICoreV1SecretKeySelector
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
