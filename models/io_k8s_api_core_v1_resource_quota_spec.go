// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// IoK8sAPICoreV1ResourceQuotaSpec ResourceQuotaSpec defines the desired hard limits to enforce for Quota.
//
// swagger:model io.k8s.api.core.v1.ResourceQuotaSpec
type IoK8sAPICoreV1ResourceQuotaSpec struct {

	// hard is the set of desired hard limits for each named resource. More info: https://kubernetes.io/docs/concepts/policy/resource-quotas/
	Hard map[string]IoK8sApimachineryPkgAPIResourceQuantity `json:"hard,omitempty"`

	// scopeSelector is also a collection of filters like scopes that must match each object tracked by a quota but expressed using ScopeSelectorOperator in combination with possible values. For a resource to match, both scopes AND scopeSelector (if specified in spec), must be matched.
	ScopeSelector *IoK8sAPICoreV1ScopeSelector `json:"scopeSelector,omitempty"`

	// A collection of filters that must match each object tracked by a quota. If not specified, the quota matches all objects.
	Scopes []string `json:"scopes"`
}

// Validate validates this io k8s api core v1 resource quota spec
func (m *IoK8sAPICoreV1ResourceQuotaSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHard(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScopeSelector(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPICoreV1ResourceQuotaSpec) validateHard(formats strfmt.Registry) error {

	if swag.IsZero(m.Hard) { // not required
		return nil
	}

	for k := range m.Hard {

		if val, ok := m.Hard[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *IoK8sAPICoreV1ResourceQuotaSpec) validateScopeSelector(formats strfmt.Registry) error {

	if swag.IsZero(m.ScopeSelector) { // not required
		return nil
	}

	if m.ScopeSelector != nil {
		if err := m.ScopeSelector.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("scopeSelector")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPICoreV1ResourceQuotaSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPICoreV1ResourceQuotaSpec) UnmarshalBinary(b []byte) error {
	var res IoK8sAPICoreV1ResourceQuotaSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
