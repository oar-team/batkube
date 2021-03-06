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

// IoK8sAPIAutoscalingV2beta1ExternalMetricSource ExternalMetricSource indicates how to scale on a metric not associated with any Kubernetes object (for example length of queue in cloud messaging service, or QPS from loadbalancer running outside of cluster). Exactly one "target" type should be set.
//
// swagger:model io.k8s.api.autoscaling.v2beta1.ExternalMetricSource
type IoK8sAPIAutoscalingV2beta1ExternalMetricSource struct {

	// metricName is the name of the metric in question.
	// Required: true
	MetricName *string `json:"metricName"`

	// metricSelector is used to identify a specific time series within a given metric.
	MetricSelector *IoK8sApimachineryPkgApisMetaV1LabelSelector `json:"metricSelector,omitempty"`

	// targetAverageValue is the target per-pod value of global metric (as a quantity). Mutually exclusive with TargetValue.
	TargetAverageValue IoK8sApimachineryPkgAPIResourceQuantity `json:"targetAverageValue,omitempty"`

	// targetValue is the target value of the metric (as a quantity). Mutually exclusive with TargetAverageValue.
	TargetValue IoK8sApimachineryPkgAPIResourceQuantity `json:"targetValue,omitempty"`
}

// Validate validates this io k8s api autoscaling v2beta1 external metric source
func (m *IoK8sAPIAutoscalingV2beta1ExternalMetricSource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMetricName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetricSelector(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTargetAverageValue(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTargetValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPIAutoscalingV2beta1ExternalMetricSource) validateMetricName(formats strfmt.Registry) error {

	if err := validate.Required("metricName", "body", m.MetricName); err != nil {
		return err
	}

	return nil
}

func (m *IoK8sAPIAutoscalingV2beta1ExternalMetricSource) validateMetricSelector(formats strfmt.Registry) error {

	if swag.IsZero(m.MetricSelector) { // not required
		return nil
	}

	if m.MetricSelector != nil {
		if err := m.MetricSelector.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metricSelector")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPIAutoscalingV2beta1ExternalMetricSource) validateTargetAverageValue(formats strfmt.Registry) error {

	if swag.IsZero(m.TargetAverageValue) { // not required
		return nil
	}

	if err := m.TargetAverageValue.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("targetAverageValue")
		}
		return err
	}

	return nil
}

func (m *IoK8sAPIAutoscalingV2beta1ExternalMetricSource) validateTargetValue(formats strfmt.Registry) error {

	if swag.IsZero(m.TargetValue) { // not required
		return nil
	}

	if err := m.TargetValue.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("targetValue")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPIAutoscalingV2beta1ExternalMetricSource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPIAutoscalingV2beta1ExternalMetricSource) UnmarshalBinary(b []byte) error {
	var res IoK8sAPIAutoscalingV2beta1ExternalMetricSource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
