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

// IoK8sAPIAutoscalingV2beta2ExternalMetricStatus ExternalMetricStatus indicates the current value of a global metric not associated with any Kubernetes object.
//
// swagger:model io.k8s.api.autoscaling.v2beta2.ExternalMetricStatus
type IoK8sAPIAutoscalingV2beta2ExternalMetricStatus struct {

	// current contains the current value for the given metric
	// Required: true
	Current *IoK8sAPIAutoscalingV2beta2MetricValueStatus `json:"current"`

	// metric identifies the target metric by name and selector
	// Required: true
	Metric *IoK8sAPIAutoscalingV2beta2MetricIdentifier `json:"metric"`
}

// Validate validates this io k8s api autoscaling v2beta2 external metric status
func (m *IoK8sAPIAutoscalingV2beta2ExternalMetricStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCurrent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetric(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPIAutoscalingV2beta2ExternalMetricStatus) validateCurrent(formats strfmt.Registry) error {

	if err := validate.Required("current", "body", m.Current); err != nil {
		return err
	}

	if m.Current != nil {
		if err := m.Current.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("current")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPIAutoscalingV2beta2ExternalMetricStatus) validateMetric(formats strfmt.Registry) error {

	if err := validate.Required("metric", "body", m.Metric); err != nil {
		return err
	}

	if m.Metric != nil {
		if err := m.Metric.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metric")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPIAutoscalingV2beta2ExternalMetricStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPIAutoscalingV2beta2ExternalMetricStatus) UnmarshalBinary(b []byte) error {
	var res IoK8sAPIAutoscalingV2beta2ExternalMetricStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
