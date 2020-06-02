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

// IoK8sAPIFlowcontrolV1alpha1LimitResponse LimitResponse defines how to handle requests that can not be executed right now.
//
// swagger:model io.k8s.api.flowcontrol.v1alpha1.LimitResponse
type IoK8sAPIFlowcontrolV1alpha1LimitResponse struct {

	// `queuing` holds the configuration parameters for queuing. This field may be non-empty only if `type` is `"Queue"`.
	Queuing *IoK8sAPIFlowcontrolV1alpha1QueuingConfiguration `json:"queuing,omitempty"`

	// `type` is "Queue" or "Reject". "Queue" means that requests that can not be executed upon arrival are held in a queue until they can be executed or a queuing limit is reached. "Reject" means that requests that can not be executed upon arrival are rejected. Required.
	// Required: true
	Type *string `json:"type"`
}

// Validate validates this io k8s api flowcontrol v1alpha1 limit response
func (m *IoK8sAPIFlowcontrolV1alpha1LimitResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateQueuing(formats); err != nil {
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

func (m *IoK8sAPIFlowcontrolV1alpha1LimitResponse) validateQueuing(formats strfmt.Registry) error {

	if swag.IsZero(m.Queuing) { // not required
		return nil
	}

	if m.Queuing != nil {
		if err := m.Queuing.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("queuing")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPIFlowcontrolV1alpha1LimitResponse) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPIFlowcontrolV1alpha1LimitResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPIFlowcontrolV1alpha1LimitResponse) UnmarshalBinary(b []byte) error {
	var res IoK8sAPIFlowcontrolV1alpha1LimitResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
