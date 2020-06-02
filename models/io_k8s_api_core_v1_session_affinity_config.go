// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// IoK8sAPICoreV1SessionAffinityConfig SessionAffinityConfig represents the configurations of session affinity.
//
// swagger:model io.k8s.api.core.v1.SessionAffinityConfig
type IoK8sAPICoreV1SessionAffinityConfig struct {

	// clientIP contains the configurations of Client IP based session affinity.
	ClientIP *IoK8sAPICoreV1ClientIPConfig `json:"clientIP,omitempty"`
}

// Validate validates this io k8s api core v1 session affinity config
func (m *IoK8sAPICoreV1SessionAffinityConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClientIP(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPICoreV1SessionAffinityConfig) validateClientIP(formats strfmt.Registry) error {

	if swag.IsZero(m.ClientIP) { // not required
		return nil
	}

	if m.ClientIP != nil {
		if err := m.ClientIP.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("clientIP")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPICoreV1SessionAffinityConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPICoreV1SessionAffinityConfig) UnmarshalBinary(b []byte) error {
	var res IoK8sAPICoreV1SessionAffinityConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
