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

// IoK8sAPIPolicyV1beta1SELinuxStrategyOptions SELinuxStrategyOptions defines the strategy type and any options used to create the strategy.
//
// swagger:model io.k8s.api.policy.v1beta1.SELinuxStrategyOptions
type IoK8sAPIPolicyV1beta1SELinuxStrategyOptions struct {

	// rule is the strategy that will dictate the allowable labels that may be set.
	// Required: true
	Rule *string `json:"rule"`

	// seLinuxOptions required to run as; required for MustRunAs More info: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/
	SeLinuxOptions *IoK8sAPICoreV1SELinuxOptions `json:"seLinuxOptions,omitempty"`
}

// Validate validates this io k8s api policy v1beta1 s e linux strategy options
func (m *IoK8sAPIPolicyV1beta1SELinuxStrategyOptions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRule(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSeLinuxOptions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPIPolicyV1beta1SELinuxStrategyOptions) validateRule(formats strfmt.Registry) error {

	if err := validate.Required("rule", "body", m.Rule); err != nil {
		return err
	}

	return nil
}

func (m *IoK8sAPIPolicyV1beta1SELinuxStrategyOptions) validateSeLinuxOptions(formats strfmt.Registry) error {

	if swag.IsZero(m.SeLinuxOptions) { // not required
		return nil
	}

	if m.SeLinuxOptions != nil {
		if err := m.SeLinuxOptions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("seLinuxOptions")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPIPolicyV1beta1SELinuxStrategyOptions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPIPolicyV1beta1SELinuxStrategyOptions) UnmarshalBinary(b []byte) error {
	var res IoK8sAPIPolicyV1beta1SELinuxStrategyOptions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
