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

// IoK8sAPIPolicyV1beta1AllowedFlexVolume AllowedFlexVolume represents a single Flexvolume that is allowed to be used.
//
// swagger:model io.k8s.api.policy.v1beta1.AllowedFlexVolume
type IoK8sAPIPolicyV1beta1AllowedFlexVolume struct {

	// driver is the name of the Flexvolume driver.
	// Required: true
	Driver *string `json:"driver"`
}

// Validate validates this io k8s api policy v1beta1 allowed flex volume
func (m *IoK8sAPIPolicyV1beta1AllowedFlexVolume) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDriver(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPIPolicyV1beta1AllowedFlexVolume) validateDriver(formats strfmt.Registry) error {

	if err := validate.Required("driver", "body", m.Driver); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPIPolicyV1beta1AllowedFlexVolume) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPIPolicyV1beta1AllowedFlexVolume) UnmarshalBinary(b []byte) error {
	var res IoK8sAPIPolicyV1beta1AllowedFlexVolume
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}