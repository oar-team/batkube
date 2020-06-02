// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// IoK8sAPIFlowcontrolV1alpha1LimitedPriorityLevelConfiguration LimitedPriorityLevelConfiguration specifies how to handle requests that are subject to limits. It addresses two issues:
//  * How are requests for this priority level limited?
//  * What should be done with requests that exceed the limit?
//
// swagger:model io.k8s.api.flowcontrol.v1alpha1.LimitedPriorityLevelConfiguration
type IoK8sAPIFlowcontrolV1alpha1LimitedPriorityLevelConfiguration struct {

	// `assuredConcurrencyShares` (ACS) configures the execution limit, which is a limit on the number of requests of this priority level that may be exeucting at a given time.  ACS must be a positive number. The server's concurrency limit (SCL) is divided among the concurrency-controlled priority levels in proportion to their assured concurrency shares. This produces the assured concurrency value (ACV) --- the number of requests that may be executing at a time --- for each such priority level:
	//
	//             ACV(l) = ceil( SCL * ACS(l) / ( sum[priority levels k] ACS(k) ) )
	//
	// bigger numbers of ACS mean more reserved concurrent requests (at the expense of every other PL). This field has a default value of 30.
	AssuredConcurrencyShares int32 `json:"assuredConcurrencyShares,omitempty"`

	// `limitResponse` indicates what to do with requests that can not be executed right now
	LimitResponse *IoK8sAPIFlowcontrolV1alpha1LimitResponse `json:"limitResponse,omitempty"`
}

// Validate validates this io k8s api flowcontrol v1alpha1 limited priority level configuration
func (m *IoK8sAPIFlowcontrolV1alpha1LimitedPriorityLevelConfiguration) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLimitResponse(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPIFlowcontrolV1alpha1LimitedPriorityLevelConfiguration) validateLimitResponse(formats strfmt.Registry) error {

	if swag.IsZero(m.LimitResponse) { // not required
		return nil
	}

	if m.LimitResponse != nil {
		if err := m.LimitResponse.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("limitResponse")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPIFlowcontrolV1alpha1LimitedPriorityLevelConfiguration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPIFlowcontrolV1alpha1LimitedPriorityLevelConfiguration) UnmarshalBinary(b []byte) error {
	var res IoK8sAPIFlowcontrolV1alpha1LimitedPriorityLevelConfiguration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
