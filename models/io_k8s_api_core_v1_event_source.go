// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// IoK8sAPICoreV1EventSource EventSource contains information for an event.
//
// swagger:model io.k8s.api.core.v1.EventSource
type IoK8sAPICoreV1EventSource struct {

	// Component from which the event is generated.
	Component string `json:"component,omitempty"`

	// Node name on which the event is generated.
	Host string `json:"host,omitempty"`
}

// Validate validates this io k8s api core v1 event source
func (m *IoK8sAPICoreV1EventSource) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPICoreV1EventSource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPICoreV1EventSource) UnmarshalBinary(b []byte) error {
	var res IoK8sAPICoreV1EventSource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
