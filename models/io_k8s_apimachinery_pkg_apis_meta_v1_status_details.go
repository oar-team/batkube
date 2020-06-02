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

// IoK8sApimachineryPkgApisMetaV1StatusDetails StatusDetails is a set of additional properties that MAY be set by the server to provide additional information about a response. The Reason field of a Status object defines what attributes will be set. Clients must ignore fields that do not match the defined type of each attribute, and should assume that any attribute may be empty, invalid, or under defined.
//
// swagger:model io.k8s.apimachinery.pkg.apis.meta.v1.StatusDetails
type IoK8sApimachineryPkgApisMetaV1StatusDetails struct {

	// The Causes array includes more details associated with the StatusReason failure. Not all StatusReasons may provide detailed causes.
	Causes []*IoK8sApimachineryPkgApisMetaV1StatusCause `json:"causes"`

	// The group attribute of the resource associated with the status StatusReason.
	Group string `json:"group,omitempty"`

	// The kind attribute of the resource associated with the status StatusReason. On some operations may differ from the requested resource Kind. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind string `json:"kind,omitempty"`

	// The name attribute of the resource associated with the status StatusReason (when there is a single name which can be described).
	Name string `json:"name,omitempty"`

	// If specified, the time in seconds before the operation should be retried. Some errors may indicate the client must take an alternate action - for those errors this field may indicate how long to wait before taking the alternate action.
	RetryAfterSeconds int32 `json:"retryAfterSeconds,omitempty"`

	// UID of the resource. (when there is a single resource which can be described). More info: http://kubernetes.io/docs/user-guide/identifiers#uids
	UID string `json:"uid,omitempty"`
}

// Validate validates this io k8s apimachinery pkg apis meta v1 status details
func (m *IoK8sApimachineryPkgApisMetaV1StatusDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCauses(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sApimachineryPkgApisMetaV1StatusDetails) validateCauses(formats strfmt.Registry) error {

	if swag.IsZero(m.Causes) { // not required
		return nil
	}

	for i := 0; i < len(m.Causes); i++ {
		if swag.IsZero(m.Causes[i]) { // not required
			continue
		}

		if m.Causes[i] != nil {
			if err := m.Causes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("causes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sApimachineryPkgApisMetaV1StatusDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sApimachineryPkgApisMetaV1StatusDetails) UnmarshalBinary(b []byte) error {
	var res IoK8sApimachineryPkgApisMetaV1StatusDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
