// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// IoK8sAPICoreV1PodIP IP address information for entries in the (plural) PodIPs field. Each entry includes:
//    IP: An IP address allocated to the pod. Routable at least within the cluster.
//
// swagger:model io.k8s.api.core.v1.PodIP
type IoK8sAPICoreV1PodIP struct {

	// ip is an IP address (IPv4 or IPv6) assigned to the pod
	IP string `json:"ip,omitempty"`
}

// Validate validates this io k8s api core v1 pod IP
func (m *IoK8sAPICoreV1PodIP) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPICoreV1PodIP) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPICoreV1PodIP) UnmarshalBinary(b []byte) error {
	var res IoK8sAPICoreV1PodIP
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
