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

// IoK8sAPIStorageV1VolumeAttachmentSpec VolumeAttachmentSpec is the specification of a VolumeAttachment request.
//
// swagger:model io.k8s.api.storage.v1.VolumeAttachmentSpec
type IoK8sAPIStorageV1VolumeAttachmentSpec struct {

	// Attacher indicates the name of the volume driver that MUST handle this request. This is the name returned by GetPluginName().
	// Required: true
	Attacher *string `json:"attacher"`

	// The node that the volume should be attached to.
	// Required: true
	NodeName *string `json:"nodeName"`

	// Source represents the volume that should be attached.
	// Required: true
	Source *IoK8sAPIStorageV1VolumeAttachmentSource `json:"source"`
}

// Validate validates this io k8s api storage v1 volume attachment spec
func (m *IoK8sAPIStorageV1VolumeAttachmentSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttacher(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNodeName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSource(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPIStorageV1VolumeAttachmentSpec) validateAttacher(formats strfmt.Registry) error {

	if err := validate.Required("attacher", "body", m.Attacher); err != nil {
		return err
	}

	return nil
}

func (m *IoK8sAPIStorageV1VolumeAttachmentSpec) validateNodeName(formats strfmt.Registry) error {

	if err := validate.Required("nodeName", "body", m.NodeName); err != nil {
		return err
	}

	return nil
}

func (m *IoK8sAPIStorageV1VolumeAttachmentSpec) validateSource(formats strfmt.Registry) error {

	if err := validate.Required("source", "body", m.Source); err != nil {
		return err
	}

	if m.Source != nil {
		if err := m.Source.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("source")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPIStorageV1VolumeAttachmentSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPIStorageV1VolumeAttachmentSpec) UnmarshalBinary(b []byte) error {
	var res IoK8sAPIStorageV1VolumeAttachmentSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
