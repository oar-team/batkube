// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// IoK8sAPICoreV1PersistentVolume PersistentVolume (PV) is a storage resource provisioned by an administrator. It is analogous to a node. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes
//
// swagger:model io.k8s.api.core.v1.PersistentVolume
type IoK8sAPICoreV1PersistentVolume struct {

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	APIVersion string `json:"apiVersion,omitempty"`

	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind string `json:"kind,omitempty"`

	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *IoK8sApimachineryPkgApisMetaV1ObjectMeta `json:"metadata,omitempty"`

	// Spec defines a specification of a persistent volume owned by the cluster. Provisioned by an administrator. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistent-volumes
	Spec *IoK8sAPICoreV1PersistentVolumeSpec `json:"spec,omitempty"`

	// Status represents the current information/status for the persistent volume. Populated by the system. Read-only. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistent-volumes
	Status *IoK8sAPICoreV1PersistentVolumeStatus `json:"status,omitempty"`
}

// Validate validates this io k8s api core v1 persistent volume
func (m *IoK8sAPICoreV1PersistentVolume) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSpec(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPICoreV1PersistentVolume) validateMetadata(formats strfmt.Registry) error {

	if swag.IsZero(m.Metadata) { // not required
		return nil
	}

	if m.Metadata != nil {
		if err := m.Metadata.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metadata")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPICoreV1PersistentVolume) validateSpec(formats strfmt.Registry) error {

	if swag.IsZero(m.Spec) { // not required
		return nil
	}

	if m.Spec != nil {
		if err := m.Spec.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("spec")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPICoreV1PersistentVolume) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if m.Status != nil {
		if err := m.Status.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPICoreV1PersistentVolume) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPICoreV1PersistentVolume) UnmarshalBinary(b []byte) error {
	var res IoK8sAPICoreV1PersistentVolume
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
