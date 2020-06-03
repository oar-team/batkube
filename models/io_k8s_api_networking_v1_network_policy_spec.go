// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// IoK8sAPINetworkingV1NetworkPolicySpec NetworkPolicySpec provides the specification of a NetworkPolicy
//
// swagger:model io.k8s.api.networking.v1.NetworkPolicySpec
type IoK8sAPINetworkingV1NetworkPolicySpec struct {

	// List of egress rules to be applied to the selected pods. Outgoing traffic is allowed if there are no NetworkPolicies selecting the pod (and cluster policy otherwise allows the traffic), OR if the traffic matches at least one egress rule across all of the NetworkPolicy objects whose podSelector matches the pod. If this field is empty then this NetworkPolicy limits all outgoing traffic (and serves solely to ensure that the pods it selects are isolated by default). This field is beta-level in 1.8
	Egress []*IoK8sAPINetworkingV1NetworkPolicyEgressRule `json:"egress"`

	// List of ingress rules to be applied to the selected pods. Traffic is allowed to a pod if there are no NetworkPolicies selecting the pod (and cluster policy otherwise allows the traffic), OR if the traffic source is the pod's local node, OR if the traffic matches at least one ingress rule across all of the NetworkPolicy objects whose podSelector matches the pod. If this field is empty then this NetworkPolicy does not allow any traffic (and serves solely to ensure that the pods it selects are isolated by default)
	Ingress []*IoK8sAPINetworkingV1NetworkPolicyIngressRule `json:"ingress"`

	// Selects the pods to which this NetworkPolicy object applies. The array of ingress rules is applied to any pods selected by this field. Multiple network policies can select the same set of pods. In this case, the ingress rules for each are combined additively. This field is NOT optional and follows standard label selector semantics. An empty podSelector matches all pods in this namespace.
	// Required: true
	PodSelector *IoK8sApimachineryPkgApisMetaV1LabelSelector `json:"podSelector"`

	// List of rule types that the NetworkPolicy relates to. Valid options are "Ingress", "Egress", or "Ingress,Egress". If this field is not specified, it will default based on the existence of Ingress or Egress rules; policies that contain an Egress section are assumed to affect Egress, and all policies (whether or not they contain an Ingress section) are assumed to affect Ingress. If you want to write an egress-only policy, you must explicitly specify policyTypes [ "Egress" ]. Likewise, if you want to write a policy that specifies that no egress is allowed, you must specify a policyTypes value that include "Egress" (since such a policy would not include an Egress section and would otherwise default to just [ "Ingress" ]). This field is beta-level in 1.8
	PolicyTypes []string `json:"policyTypes"`
}

// Validate validates this io k8s api networking v1 network policy spec
func (m *IoK8sAPINetworkingV1NetworkPolicySpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEgress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIngress(formats); err != nil {
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

func (m *IoK8sAPINetworkingV1NetworkPolicySpec) validateEgress(formats strfmt.Registry) error {

	if swag.IsZero(m.Egress) { // not required
		return nil
	}

	for i := 0; i < len(m.Egress); i++ {
		if swag.IsZero(m.Egress[i]) { // not required
			continue
		}

		if m.Egress[i] != nil {
			if err := m.Egress[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("egress" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IoK8sAPINetworkingV1NetworkPolicySpec) validateIngress(formats strfmt.Registry) error {

	if swag.IsZero(m.Ingress) { // not required
		return nil
	}

	for i := 0; i < len(m.Ingress); i++ {
		if swag.IsZero(m.Ingress[i]) { // not required
			continue
		}

		if m.Ingress[i] != nil {
			if err := m.Ingress[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ingress" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IoK8sAPINetworkingV1NetworkPolicySpec) validatePodSelector(formats strfmt.Registry) error {

	if err := validate.Required("podSelector", "body", m.PodSelector); err != nil {
		return err
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
func (m *IoK8sAPINetworkingV1NetworkPolicySpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPINetworkingV1NetworkPolicySpec) UnmarshalBinary(b []byte) error {
	var res IoK8sAPINetworkingV1NetworkPolicySpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}