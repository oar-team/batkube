// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// IoK8sAPINetworkingV1NetworkPolicyPeer NetworkPolicyPeer describes a peer to allow traffic from. Only certain combinations of fields are allowed
//
// swagger:model io.k8s.api.networking.v1.NetworkPolicyPeer
type IoK8sAPINetworkingV1NetworkPolicyPeer struct {

	// IPBlock defines policy on a particular IPBlock. If this field is set then neither of the other fields can be.
	IPBlock *IoK8sAPINetworkingV1IPBlock `json:"ipBlock,omitempty"`

	// Selects Namespaces using cluster-scoped labels. This field follows standard label selector semantics; if present but empty, it selects all namespaces.
	//
	// If PodSelector is also set, then the NetworkPolicyPeer as a whole selects the Pods matching PodSelector in the Namespaces selected by NamespaceSelector. Otherwise it selects all Pods in the Namespaces selected by NamespaceSelector.
	NamespaceSelector *IoK8sApimachineryPkgApisMetaV1LabelSelector `json:"namespaceSelector,omitempty"`

	// This is a label selector which selects Pods. This field follows standard label selector semantics; if present but empty, it selects all pods.
	//
	// If NamespaceSelector is also set, then the NetworkPolicyPeer as a whole selects the Pods matching PodSelector in the Namespaces selected by NamespaceSelector. Otherwise it selects the Pods matching PodSelector in the policy's own Namespace.
	PodSelector *IoK8sApimachineryPkgApisMetaV1LabelSelector `json:"podSelector,omitempty"`
}

// Validate validates this io k8s api networking v1 network policy peer
func (m *IoK8sAPINetworkingV1NetworkPolicyPeer) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIPBlock(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNamespaceSelector(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePodSelector(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPINetworkingV1NetworkPolicyPeer) validateIPBlock(formats strfmt.Registry) error {

	if swag.IsZero(m.IPBlock) { // not required
		return nil
	}

	if m.IPBlock != nil {
		if err := m.IPBlock.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ipBlock")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPINetworkingV1NetworkPolicyPeer) validateNamespaceSelector(formats strfmt.Registry) error {

	if swag.IsZero(m.NamespaceSelector) { // not required
		return nil
	}

	if m.NamespaceSelector != nil {
		if err := m.NamespaceSelector.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("namespaceSelector")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPINetworkingV1NetworkPolicyPeer) validatePodSelector(formats strfmt.Registry) error {

	if swag.IsZero(m.PodSelector) { // not required
		return nil
	}

	if m.PodSelector != nil {
		if err := m.PodSelector.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("podSelector")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPINetworkingV1NetworkPolicyPeer) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPINetworkingV1NetworkPolicyPeer) UnmarshalBinary(b []byte) error {
	var res IoK8sAPINetworkingV1NetworkPolicyPeer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
