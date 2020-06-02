// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// IoK8sKubeAggregatorPkgApisApiregistrationV1beta1APIService APIService represents a server for a particular GroupVersion. Name must be "version.group".
//
// swagger:model io.k8s.kube-aggregator.pkg.apis.apiregistration.v1beta1.APIService
type IoK8sKubeAggregatorPkgApisApiregistrationV1beta1APIService struct {

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	APIVersion string `json:"apiVersion,omitempty"`

	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind string `json:"kind,omitempty"`

	// metadata
	Metadata *IoK8sApimachineryPkgApisMetaV1ObjectMeta `json:"metadata,omitempty"`

	// Spec contains information for locating and communicating with a server
	Spec *IoK8sKubeAggregatorPkgApisApiregistrationV1beta1APIServiceSpec `json:"spec,omitempty"`

	// Status contains derived information about an API server
	Status *IoK8sKubeAggregatorPkgApisApiregistrationV1beta1APIServiceStatus `json:"status,omitempty"`
}

// Validate validates this io k8s kube aggregator pkg apis apiregistration v1beta1 API service
func (m *IoK8sKubeAggregatorPkgApisApiregistrationV1beta1APIService) Validate(formats strfmt.Registry) error {
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

func (m *IoK8sKubeAggregatorPkgApisApiregistrationV1beta1APIService) validateMetadata(formats strfmt.Registry) error {

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

func (m *IoK8sKubeAggregatorPkgApisApiregistrationV1beta1APIService) validateSpec(formats strfmt.Registry) error {

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

func (m *IoK8sKubeAggregatorPkgApisApiregistrationV1beta1APIService) validateStatus(formats strfmt.Registry) error {

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
func (m *IoK8sKubeAggregatorPkgApisApiregistrationV1beta1APIService) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sKubeAggregatorPkgApisApiregistrationV1beta1APIService) UnmarshalBinary(b []byte) error {
	var res IoK8sKubeAggregatorPkgApisApiregistrationV1beta1APIService
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
