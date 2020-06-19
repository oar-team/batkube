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

// IoK8sAPIAuthenticationV1TokenRequestStatus TokenRequestStatus is the result of a token request.
//
// swagger:model io.k8s.api.authentication.v1.TokenRequestStatus
type IoK8sAPIAuthenticationV1TokenRequestStatus struct {

	// ExpirationTimestamp is the time of expiration of the returned token.
	// Required: true
	// Format: date-time
	ExpirationTimestamp *IoK8sApimachineryPkgApisMetaV1Time `json:"expirationTimestamp"`

	// Token is the opaque bearer token.
	// Required: true
	Token *string `json:"token"`
}

// Validate validates this io k8s api authentication v1 token request status
func (m *IoK8sAPIAuthenticationV1TokenRequestStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExpirationTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateToken(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPIAuthenticationV1TokenRequestStatus) validateExpirationTimestamp(formats strfmt.Registry) error {

	if err := m.ExpirationTimestamp.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("expirationTimestamp")
		}
		return err
	}

	return nil
}

func (m *IoK8sAPIAuthenticationV1TokenRequestStatus) validateToken(formats strfmt.Registry) error {

	if err := validate.Required("token", "body", m.Token); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPIAuthenticationV1TokenRequestStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPIAuthenticationV1TokenRequestStatus) UnmarshalBinary(b []byte) error {
	var res IoK8sAPIAuthenticationV1TokenRequestStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
