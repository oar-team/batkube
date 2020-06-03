// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// IoK8sAPICoreV1EmptyDirVolumeSource Represents an empty directory for a pod. Empty directory volumes support ownership management and SELinux relabeling.
//
// swagger:model io.k8s.api.core.v1.EmptyDirVolumeSource
type IoK8sAPICoreV1EmptyDirVolumeSource struct {

	// What type of storage medium should back this directory. The default is "" which means to use the node's default medium. Must be an empty string (default) or Memory. More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir
	Medium string `json:"medium,omitempty"`

	// Total amount of local storage required for this EmptyDir volume. The size limit is also applicable for memory medium. The maximum usage on memory medium EmptyDir would be the minimum value between the SizeLimit specified here and the sum of memory limits of all containers in a pod. The default is nil which means that the limit is undefined. More info: http://kubernetes.io/docs/user-guide/volumes#emptydir
	SizeLimit IoK8sApimachineryPkgAPIResourceQuantity `json:"sizeLimit,omitempty"`
}

// Validate validates this io k8s api core v1 empty dir volume source
func (m *IoK8sAPICoreV1EmptyDirVolumeSource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSizeLimit(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPICoreV1EmptyDirVolumeSource) validateSizeLimit(formats strfmt.Registry) error {

	if swag.IsZero(m.SizeLimit) { // not required
		return nil
	}

	if err := m.SizeLimit.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("sizeLimit")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPICoreV1EmptyDirVolumeSource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPICoreV1EmptyDirVolumeSource) UnmarshalBinary(b []byte) error {
	var res IoK8sAPICoreV1EmptyDirVolumeSource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}