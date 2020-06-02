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

// IoK8sAPIAutoscalingV1ScaleStatus ScaleStatus represents the current status of a scale subresource.
//
// swagger:model io.k8s.api.autoscaling.v1.ScaleStatus
type IoK8sAPIAutoscalingV1ScaleStatus struct {

	// actual number of observed instances of the scaled object.
	// Required: true
	Replicas *int32 `json:"replicas"`

	// label query over pods that should match the replicas count. This is same as the label selector but in the string format to avoid introspection by clients. The string will be in the same format as the query-param syntax. More info about label selectors: http://kubernetes.io/docs/user-guide/labels#label-selectors
	Selector string `json:"selector,omitempty"`
}

// Validate validates this io k8s api autoscaling v1 scale status
func (m *IoK8sAPIAutoscalingV1ScaleStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateReplicas(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPIAutoscalingV1ScaleStatus) validateReplicas(formats strfmt.Registry) error {

	if err := validate.Required("replicas", "body", m.Replicas); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPIAutoscalingV1ScaleStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPIAutoscalingV1ScaleStatus) UnmarshalBinary(b []byte) error {
	var res IoK8sAPIAutoscalingV1ScaleStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
