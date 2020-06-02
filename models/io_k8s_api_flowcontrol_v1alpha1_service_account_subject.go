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

// IoK8sAPIFlowcontrolV1alpha1ServiceAccountSubject ServiceAccountSubject holds detailed information for service-account-kind subject.
//
// swagger:model io.k8s.api.flowcontrol.v1alpha1.ServiceAccountSubject
type IoK8sAPIFlowcontrolV1alpha1ServiceAccountSubject struct {

	// `name` is the name of matching ServiceAccount objects, or "*" to match regardless of name. Required.
	// Required: true
	Name *string `json:"name"`

	// `namespace` is the namespace of matching ServiceAccount objects. Required.
	// Required: true
	Namespace *string `json:"namespace"`
}

// Validate validates this io k8s api flowcontrol v1alpha1 service account subject
func (m *IoK8sAPIFlowcontrolV1alpha1ServiceAccountSubject) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNamespace(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPIFlowcontrolV1alpha1ServiceAccountSubject) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *IoK8sAPIFlowcontrolV1alpha1ServiceAccountSubject) validateNamespace(formats strfmt.Registry) error {

	if err := validate.Required("namespace", "body", m.Namespace); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPIFlowcontrolV1alpha1ServiceAccountSubject) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPIFlowcontrolV1alpha1ServiceAccountSubject) UnmarshalBinary(b []byte) error {
	var res IoK8sAPIFlowcontrolV1alpha1ServiceAccountSubject
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
