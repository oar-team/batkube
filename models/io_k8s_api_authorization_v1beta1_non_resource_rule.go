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

// IoK8sAPIAuthorizationV1beta1NonResourceRule NonResourceRule holds information that describes a rule for the non-resource
//
// swagger:model io.k8s.api.authorization.v1beta1.NonResourceRule
type IoK8sAPIAuthorizationV1beta1NonResourceRule struct {

	// NonResourceURLs is a set of partial urls that a user should have access to.  *s are allowed, but only as the full, final step in the path.  "*" means all.
	NonResourceURLs []string `json:"nonResourceURLs"`

	// Verb is a list of kubernetes non-resource API verbs, like: get, post, put, delete, patch, head, options.  "*" means all.
	// Required: true
	Verbs []string `json:"verbs"`
}

// Validate validates this io k8s api authorization v1beta1 non resource rule
func (m *IoK8sAPIAuthorizationV1beta1NonResourceRule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateVerbs(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPIAuthorizationV1beta1NonResourceRule) validateVerbs(formats strfmt.Registry) error {

	if err := validate.Required("verbs", "body", m.Verbs); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPIAuthorizationV1beta1NonResourceRule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPIAuthorizationV1beta1NonResourceRule) UnmarshalBinary(b []byte) error {
	var res IoK8sAPIAuthorizationV1beta1NonResourceRule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
