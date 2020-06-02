// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// IoK8sAPIAppsV1RollingUpdateStatefulSetStrategy RollingUpdateStatefulSetStrategy is used to communicate parameter for RollingUpdateStatefulSetStrategyType.
//
// swagger:model io.k8s.api.apps.v1.RollingUpdateStatefulSetStrategy
type IoK8sAPIAppsV1RollingUpdateStatefulSetStrategy struct {

	// Partition indicates the ordinal at which the StatefulSet should be partitioned. Default value is 0.
	Partition int32 `json:"partition,omitempty"`
}

// Validate validates this io k8s api apps v1 rolling update stateful set strategy
func (m *IoK8sAPIAppsV1RollingUpdateStatefulSetStrategy) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPIAppsV1RollingUpdateStatefulSetStrategy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPIAppsV1RollingUpdateStatefulSetStrategy) UnmarshalBinary(b []byte) error {
	var res IoK8sAPIAppsV1RollingUpdateStatefulSetStrategy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
