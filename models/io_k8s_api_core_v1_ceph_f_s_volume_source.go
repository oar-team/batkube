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

// IoK8sAPICoreV1CephFSVolumeSource Represents a Ceph Filesystem mount that lasts the lifetime of a pod Cephfs volumes do not support ownership management or SELinux relabeling.
//
// swagger:model io.k8s.api.core.v1.CephFSVolumeSource
type IoK8sAPICoreV1CephFSVolumeSource struct {

	// Required: Monitors is a collection of Ceph monitors More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it
	// Required: true
	Monitors []string `json:"monitors"`

	// Optional: Used as the mounted root, rather than the full Ceph tree, default is /
	Path string `json:"path,omitempty"`

	// Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it
	ReadOnly bool `json:"readOnly,omitempty"`

	// Optional: SecretFile is the path to key ring for User, default is /etc/ceph/user.secret More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it
	SecretFile string `json:"secretFile,omitempty"`

	// Optional: SecretRef is reference to the authentication secret for User, default is empty. More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it
	SecretRef *IoK8sAPICoreV1LocalObjectReference `json:"secretRef,omitempty"`

	// Optional: User is the rados user name, default is admin More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it
	User string `json:"user,omitempty"`
}

// Validate validates this io k8s api core v1 ceph f s volume source
func (m *IoK8sAPICoreV1CephFSVolumeSource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMonitors(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecretRef(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPICoreV1CephFSVolumeSource) validateMonitors(formats strfmt.Registry) error {

	if err := validate.Required("monitors", "body", m.Monitors); err != nil {
		return err
	}

	return nil
}

func (m *IoK8sAPICoreV1CephFSVolumeSource) validateSecretRef(formats strfmt.Registry) error {

	if swag.IsZero(m.SecretRef) { // not required
		return nil
	}

	if m.SecretRef != nil {
		if err := m.SecretRef.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("secretRef")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPICoreV1CephFSVolumeSource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPICoreV1CephFSVolumeSource) UnmarshalBinary(b []byte) error {
	var res IoK8sAPICoreV1CephFSVolumeSource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
