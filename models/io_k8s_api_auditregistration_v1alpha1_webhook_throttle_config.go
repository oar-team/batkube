// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// IoK8sAPIAuditregistrationV1alpha1WebhookThrottleConfig WebhookThrottleConfig holds the configuration for throttling events
//
// swagger:model io.k8s.api.auditregistration.v1alpha1.WebhookThrottleConfig
type IoK8sAPIAuditregistrationV1alpha1WebhookThrottleConfig struct {

	// ThrottleBurst is the maximum number of events sent at the same moment default 15 QPS
	Burst int64 `json:"burst,omitempty"`

	// ThrottleQPS maximum number of batches per second default 10 QPS
	QPS int64 `json:"qps,omitempty"`
}

// Validate validates this io k8s api auditregistration v1alpha1 webhook throttle config
func (m *IoK8sAPIAuditregistrationV1alpha1WebhookThrottleConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPIAuditregistrationV1alpha1WebhookThrottleConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPIAuditregistrationV1alpha1WebhookThrottleConfig) UnmarshalBinary(b []byte) error {
	var res IoK8sAPIAuditregistrationV1alpha1WebhookThrottleConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
