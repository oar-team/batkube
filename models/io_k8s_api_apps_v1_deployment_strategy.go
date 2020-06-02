// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// IoK8sAPIAppsV1DeploymentStrategy DeploymentStrategy describes how to replace existing pods with new ones.
//
// swagger:model io.k8s.api.apps.v1.DeploymentStrategy
type IoK8sAPIAppsV1DeploymentStrategy struct {

	// Rolling update config params. Present only if DeploymentStrategyType = RollingUpdate.
	RollingUpdate *IoK8sAPIAppsV1RollingUpdateDeployment `json:"rollingUpdate,omitempty"`

	// Type of deployment. Can be "Recreate" or "RollingUpdate". Default is RollingUpdate.
	Type string `json:"type,omitempty"`
}

// Validate validates this io k8s api apps v1 deployment strategy
func (m *IoK8sAPIAppsV1DeploymentStrategy) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRollingUpdate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPIAppsV1DeploymentStrategy) validateRollingUpdate(formats strfmt.Registry) error {

	if swag.IsZero(m.RollingUpdate) { // not required
		return nil
	}

	if m.RollingUpdate != nil {
		if err := m.RollingUpdate.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rollingUpdate")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPIAppsV1DeploymentStrategy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPIAppsV1DeploymentStrategy) UnmarshalBinary(b []byte) error {
	var res IoK8sAPIAppsV1DeploymentStrategy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
