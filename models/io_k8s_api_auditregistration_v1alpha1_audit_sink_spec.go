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

// IoK8sAPIAuditregistrationV1alpha1AuditSinkSpec AuditSinkSpec holds the spec for the audit sink
//
// swagger:model io.k8s.api.auditregistration.v1alpha1.AuditSinkSpec
type IoK8sAPIAuditregistrationV1alpha1AuditSinkSpec struct {

	// Policy defines the policy for selecting which events should be sent to the webhook required
	// Required: true
	Policy *IoK8sAPIAuditregistrationV1alpha1Policy `json:"policy"`

	// Webhook to send events required
	// Required: true
	Webhook *IoK8sAPIAuditregistrationV1alpha1Webhook `json:"webhook"`
}

// Validate validates this io k8s api auditregistration v1alpha1 audit sink spec
func (m *IoK8sAPIAuditregistrationV1alpha1AuditSinkSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWebhook(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPIAuditregistrationV1alpha1AuditSinkSpec) validatePolicy(formats strfmt.Registry) error {

	if err := validate.Required("policy", "body", m.Policy); err != nil {
		return err
	}

	if m.Policy != nil {
		if err := m.Policy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("policy")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPIAuditregistrationV1alpha1AuditSinkSpec) validateWebhook(formats strfmt.Registry) error {

	if err := validate.Required("webhook", "body", m.Webhook); err != nil {
		return err
	}

	if m.Webhook != nil {
		if err := m.Webhook.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("webhook")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPIAuditregistrationV1alpha1AuditSinkSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPIAuditregistrationV1alpha1AuditSinkSpec) UnmarshalBinary(b []byte) error {
	var res IoK8sAPIAuditregistrationV1alpha1AuditSinkSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
