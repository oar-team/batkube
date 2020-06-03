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

// IoK8sAPIAutoscalingV2beta2HorizontalPodAutoscalerCondition HorizontalPodAutoscalerCondition describes the state of a HorizontalPodAutoscaler at a certain point.
//
// swagger:model io.k8s.api.autoscaling.v2beta2.HorizontalPodAutoscalerCondition
type IoK8sAPIAutoscalingV2beta2HorizontalPodAutoscalerCondition struct {

	// lastTransitionTime is the last time the condition transitioned from one status to another
	// Format: date-time
	LastTransitionTime IoK8sApimachineryPkgApisMetaV1Time `json:"lastTransitionTime,omitempty"`

	// message is a human-readable explanation containing details about the transition
	Message string `json:"message,omitempty"`

	// reason is the reason for the condition's last transition.
	Reason string `json:"reason,omitempty"`

	// status is the status of the condition (True, False, Unknown)
	// Required: true
	Status *string `json:"status"`

	// type describes the current condition
	// Required: true
	Type *string `json:"type"`
}

// Validate validates this io k8s api autoscaling v2beta2 horizontal pod autoscaler condition
func (m *IoK8sAPIAutoscalingV2beta2HorizontalPodAutoscalerCondition) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLastTransitionTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPIAutoscalingV2beta2HorizontalPodAutoscalerCondition) validateLastTransitionTime(formats strfmt.Registry) error {

	if swag.IsZero(m.LastTransitionTime) { // not required
		return nil
	}

	if err := m.LastTransitionTime.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("lastTransitionTime")
		}
		return err
	}

	return nil
}

func (m *IoK8sAPIAutoscalingV2beta2HorizontalPodAutoscalerCondition) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *IoK8sAPIAutoscalingV2beta2HorizontalPodAutoscalerCondition) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPIAutoscalingV2beta2HorizontalPodAutoscalerCondition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPIAutoscalingV2beta2HorizontalPodAutoscalerCondition) UnmarshalBinary(b []byte) error {
	var res IoK8sAPIAutoscalingV2beta2HorizontalPodAutoscalerCondition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}