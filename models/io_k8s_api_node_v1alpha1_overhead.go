// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// IoK8sAPINodeV1alpha1Overhead Overhead structure represents the resource overhead associated with running a pod.
//
// swagger:model io.k8s.api.node.v1alpha1.Overhead
type IoK8sAPINodeV1alpha1Overhead struct {

	// PodFixed represents the fixed resource overhead associated with running a pod.
	PodFixed map[string]IoK8sApimachineryPkgAPIResourceQuantity `json:"podFixed,omitempty"`
}

// Validate validates this io k8s api node v1alpha1 overhead
func (m *IoK8sAPINodeV1alpha1Overhead) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePodFixed(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPINodeV1alpha1Overhead) validatePodFixed(formats strfmt.Registry) error {

	if swag.IsZero(m.PodFixed) { // not required
		return nil
	}

	for k := range m.PodFixed {

		if val, ok := m.PodFixed[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPINodeV1alpha1Overhead) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPINodeV1alpha1Overhead) UnmarshalBinary(b []byte) error {
	var res IoK8sAPINodeV1alpha1Overhead
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
