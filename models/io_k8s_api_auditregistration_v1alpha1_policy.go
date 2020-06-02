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

// IoK8sAPIAuditregistrationV1alpha1Policy Policy defines the configuration of how audit events are logged
//
// swagger:model io.k8s.api.auditregistration.v1alpha1.Policy
type IoK8sAPIAuditregistrationV1alpha1Policy struct {

	// The Level that all requests are recorded at. available options: None, Metadata, Request, RequestResponse required
	// Required: true
	Level *string `json:"level"`

	// Stages is a list of stages for which events are created.
	Stages []string `json:"stages"`
}

// Validate validates this io k8s api auditregistration v1alpha1 policy
func (m *IoK8sAPIAuditregistrationV1alpha1Policy) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLevel(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPIAuditregistrationV1alpha1Policy) validateLevel(formats strfmt.Registry) error {

	if err := validate.Required("level", "body", m.Level); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPIAuditregistrationV1alpha1Policy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPIAuditregistrationV1alpha1Policy) UnmarshalBinary(b []byte) error {
	var res IoK8sAPIAuditregistrationV1alpha1Policy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
