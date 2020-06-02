// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// IoK8sAPIAutoscalingV2beta2HorizontalPodAutoscalerStatus HorizontalPodAutoscalerStatus describes the current status of a horizontal pod autoscaler.
//
// swagger:model io.k8s.api.autoscaling.v2beta2.HorizontalPodAutoscalerStatus
type IoK8sAPIAutoscalingV2beta2HorizontalPodAutoscalerStatus struct {

	// conditions is the set of conditions required for this autoscaler to scale its target, and indicates whether or not those conditions are met.
	// Required: true
	Conditions []*IoK8sAPIAutoscalingV2beta2HorizontalPodAutoscalerCondition `json:"conditions"`

	// currentMetrics is the last read state of the metrics used by this autoscaler.
	CurrentMetrics []*IoK8sAPIAutoscalingV2beta2MetricStatus `json:"currentMetrics"`

	// currentReplicas is current number of replicas of pods managed by this autoscaler, as last seen by the autoscaler.
	// Required: true
	CurrentReplicas *int32 `json:"currentReplicas"`

	// desiredReplicas is the desired number of replicas of pods managed by this autoscaler, as last calculated by the autoscaler.
	// Required: true
	DesiredReplicas *int32 `json:"desiredReplicas"`

	// lastScaleTime is the last time the HorizontalPodAutoscaler scaled the number of pods, used by the autoscaler to control how often the number of pods is changed.
	// Format: date-time
	LastScaleTime IoK8sApimachineryPkgApisMetaV1Time `json:"lastScaleTime,omitempty"`

	// observedGeneration is the most recent generation observed by this autoscaler.
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
}

// Validate validates this io k8s api autoscaling v2beta2 horizontal pod autoscaler status
func (m *IoK8sAPIAutoscalingV2beta2HorizontalPodAutoscalerStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConditions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCurrentMetrics(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCurrentReplicas(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDesiredReplicas(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastScaleTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPIAutoscalingV2beta2HorizontalPodAutoscalerStatus) validateConditions(formats strfmt.Registry) error {

	if err := validate.Required("conditions", "body", m.Conditions); err != nil {
		return err
	}

	for i := 0; i < len(m.Conditions); i++ {
		if swag.IsZero(m.Conditions[i]) { // not required
			continue
		}

		if m.Conditions[i] != nil {
			if err := m.Conditions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("conditions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IoK8sAPIAutoscalingV2beta2HorizontalPodAutoscalerStatus) validateCurrentMetrics(formats strfmt.Registry) error {

	if swag.IsZero(m.CurrentMetrics) { // not required
		return nil
	}

	for i := 0; i < len(m.CurrentMetrics); i++ {
		if swag.IsZero(m.CurrentMetrics[i]) { // not required
			continue
		}

		if m.CurrentMetrics[i] != nil {
			if err := m.CurrentMetrics[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("currentMetrics" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IoK8sAPIAutoscalingV2beta2HorizontalPodAutoscalerStatus) validateCurrentReplicas(formats strfmt.Registry) error {

	if err := validate.Required("currentReplicas", "body", m.CurrentReplicas); err != nil {
		return err
	}

	return nil
}

func (m *IoK8sAPIAutoscalingV2beta2HorizontalPodAutoscalerStatus) validateDesiredReplicas(formats strfmt.Registry) error {

	if err := validate.Required("desiredReplicas", "body", m.DesiredReplicas); err != nil {
		return err
	}

	return nil
}

func (m *IoK8sAPIAutoscalingV2beta2HorizontalPodAutoscalerStatus) validateLastScaleTime(formats strfmt.Registry) error {

	if swag.IsZero(m.LastScaleTime) { // not required
		return nil
	}

	if err := m.LastScaleTime.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("lastScaleTime")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPIAutoscalingV2beta2HorizontalPodAutoscalerStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPIAutoscalingV2beta2HorizontalPodAutoscalerStatus) UnmarshalBinary(b []byte) error {
	var res IoK8sAPIAutoscalingV2beta2HorizontalPodAutoscalerStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
