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

// IoK8sAPICoreV1VolumeDevice volumeDevice describes a mapping of a raw block device within a container.
//
// swagger:model io.k8s.api.core.v1.VolumeDevice
type IoK8sAPICoreV1VolumeDevice struct {

	// devicePath is the path inside of the container that the device will be mapped to.
	// Required: true
	DevicePath *string `json:"devicePath"`

	// name must match the name of a persistentVolumeClaim in the pod
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this io k8s api core v1 volume device
func (m *IoK8sAPICoreV1VolumeDevice) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDevicePath(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPICoreV1VolumeDevice) validateDevicePath(formats strfmt.Registry) error {

	if err := validate.Required("devicePath", "body", m.DevicePath); err != nil {
		return err
	}

	return nil
}

func (m *IoK8sAPICoreV1VolumeDevice) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPICoreV1VolumeDevice) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPICoreV1VolumeDevice) UnmarshalBinary(b []byte) error {
	var res IoK8sAPICoreV1VolumeDevice
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}