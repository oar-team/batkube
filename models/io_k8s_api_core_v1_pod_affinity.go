// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// IoK8sAPICoreV1PodAffinity Pod affinity is a group of inter pod affinity scheduling rules.
//
// swagger:model io.k8s.api.core.v1.PodAffinity
type IoK8sAPICoreV1PodAffinity struct {

	// The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding "weight" to the sum if the node has pods which matches the corresponding podAffinityTerm; the node(s) with the highest sum are the most preferred.
	PreferredDuringSchedulingIgnoredDuringExecution []*IoK8sAPICoreV1WeightedPodAffinityTerm `json:"preferredDuringSchedulingIgnoredDuringExecution"`

	// If the affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to a pod label update), the system may or may not try to eventually evict the pod from its node. When there are multiple elements, the lists of nodes corresponding to each podAffinityTerm are intersected, i.e. all terms must be satisfied.
	RequiredDuringSchedulingIgnoredDuringExecution []*IoK8sAPICoreV1PodAffinityTerm `json:"requiredDuringSchedulingIgnoredDuringExecution"`
}

// Validate validates this io k8s api core v1 pod affinity
func (m *IoK8sAPICoreV1PodAffinity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePreferredDuringSchedulingIgnoredDuringExecution(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequiredDuringSchedulingIgnoredDuringExecution(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPICoreV1PodAffinity) validatePreferredDuringSchedulingIgnoredDuringExecution(formats strfmt.Registry) error {

	if swag.IsZero(m.PreferredDuringSchedulingIgnoredDuringExecution) { // not required
		return nil
	}

	for i := 0; i < len(m.PreferredDuringSchedulingIgnoredDuringExecution); i++ {
		if swag.IsZero(m.PreferredDuringSchedulingIgnoredDuringExecution[i]) { // not required
			continue
		}

		if m.PreferredDuringSchedulingIgnoredDuringExecution[i] != nil {
			if err := m.PreferredDuringSchedulingIgnoredDuringExecution[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("preferredDuringSchedulingIgnoredDuringExecution" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IoK8sAPICoreV1PodAffinity) validateRequiredDuringSchedulingIgnoredDuringExecution(formats strfmt.Registry) error {

	if swag.IsZero(m.RequiredDuringSchedulingIgnoredDuringExecution) { // not required
		return nil
	}

	for i := 0; i < len(m.RequiredDuringSchedulingIgnoredDuringExecution); i++ {
		if swag.IsZero(m.RequiredDuringSchedulingIgnoredDuringExecution[i]) { // not required
			continue
		}

		if m.RequiredDuringSchedulingIgnoredDuringExecution[i] != nil {
			if err := m.RequiredDuringSchedulingIgnoredDuringExecution[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("requiredDuringSchedulingIgnoredDuringExecution" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPICoreV1PodAffinity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPICoreV1PodAffinity) UnmarshalBinary(b []byte) error {
	var res IoK8sAPICoreV1PodAffinity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
