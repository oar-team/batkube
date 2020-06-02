// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// IoK8sAPIPolicyV1beta1Eviction Eviction evicts a pod from its node subject to certain policies and safety constraints. This is a subresource of Pod.  A request to cause such an eviction is created by POSTing to .../pods/<pod name>/evictions.
//
// swagger:model io.k8s.api.policy.v1beta1.Eviction
type IoK8sAPIPolicyV1beta1Eviction struct {

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	APIVersion string `json:"apiVersion,omitempty"`

	// DeleteOptions may be provided
	DeleteOptions *IoK8sApimachineryPkgApisMetaV1DeleteOptions `json:"deleteOptions,omitempty"`

	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind string `json:"kind,omitempty"`

	// ObjectMeta describes the pod that is being evicted.
	Metadata *IoK8sApimachineryPkgApisMetaV1ObjectMeta `json:"metadata,omitempty"`
}

// Validate validates this io k8s api policy v1beta1 eviction
func (m *IoK8sAPIPolicyV1beta1Eviction) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDeleteOptions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPIPolicyV1beta1Eviction) validateDeleteOptions(formats strfmt.Registry) error {

	if swag.IsZero(m.DeleteOptions) { // not required
		return nil
	}

	if m.DeleteOptions != nil {
		if err := m.DeleteOptions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deleteOptions")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPIPolicyV1beta1Eviction) validateMetadata(formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *IoK8sAPIPolicyV1beta1Eviction) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPIPolicyV1beta1Eviction) UnmarshalBinary(b []byte) error {
	var res IoK8sAPIPolicyV1beta1Eviction
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
