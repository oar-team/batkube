// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// IoK8sAPICoreV1NodeConfigStatus NodeConfigStatus describes the status of the config assigned by Node.Spec.ConfigSource.
//
// swagger:model io.k8s.api.core.v1.NodeConfigStatus
type IoK8sAPICoreV1NodeConfigStatus struct {

	// Active reports the checkpointed config the node is actively using. Active will represent either the current version of the Assigned config, or the current LastKnownGood config, depending on whether attempting to use the Assigned config results in an error.
	Active *IoK8sAPICoreV1NodeConfigSource `json:"active,omitempty"`

	// Assigned reports the checkpointed config the node will try to use. When Node.Spec.ConfigSource is updated, the node checkpoints the associated config payload to local disk, along with a record indicating intended config. The node refers to this record to choose its config checkpoint, and reports this record in Assigned. Assigned only updates in the status after the record has been checkpointed to disk. When the Kubelet is restarted, it tries to make the Assigned config the Active config by loading and validating the checkpointed payload identified by Assigned.
	Assigned *IoK8sAPICoreV1NodeConfigSource `json:"assigned,omitempty"`

	// Error describes any problems reconciling the Spec.ConfigSource to the Active config. Errors may occur, for example, attempting to checkpoint Spec.ConfigSource to the local Assigned record, attempting to checkpoint the payload associated with Spec.ConfigSource, attempting to load or validate the Assigned config, etc. Errors may occur at different points while syncing config. Earlier errors (e.g. download or checkpointing errors) will not result in a rollback to LastKnownGood, and may resolve across Kubelet retries. Later errors (e.g. loading or validating a checkpointed config) will result in a rollback to LastKnownGood. In the latter case, it is usually possible to resolve the error by fixing the config assigned in Spec.ConfigSource. You can find additional information for debugging by searching the error message in the Kubelet log. Error is a human-readable description of the error state; machines can check whether or not Error is empty, but should not rely on the stability of the Error text across Kubelet versions.
	Error string `json:"error,omitempty"`

	// LastKnownGood reports the checkpointed config the node will fall back to when it encounters an error attempting to use the Assigned config. The Assigned config becomes the LastKnownGood config when the node determines that the Assigned config is stable and correct. This is currently implemented as a 10-minute soak period starting when the local record of Assigned config is updated. If the Assigned config is Active at the end of this period, it becomes the LastKnownGood. Note that if Spec.ConfigSource is reset to nil (use local defaults), the LastKnownGood is also immediately reset to nil, because the local default config is always assumed good. You should not make assumptions about the node's method of determining config stability and correctness, as this may change or become configurable in the future.
	LastKnownGood *IoK8sAPICoreV1NodeConfigSource `json:"lastKnownGood,omitempty"`
}

// Validate validates this io k8s api core v1 node config status
func (m *IoK8sAPICoreV1NodeConfigStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActive(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAssigned(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastKnownGood(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPICoreV1NodeConfigStatus) validateActive(formats strfmt.Registry) error {

	if swag.IsZero(m.Active) { // not required
		return nil
	}

	if m.Active != nil {
		if err := m.Active.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("active")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPICoreV1NodeConfigStatus) validateAssigned(formats strfmt.Registry) error {

	if swag.IsZero(m.Assigned) { // not required
		return nil
	}

	if m.Assigned != nil {
		if err := m.Assigned.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("assigned")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPICoreV1NodeConfigStatus) validateLastKnownGood(formats strfmt.Registry) error {

	if swag.IsZero(m.LastKnownGood) { // not required
		return nil
	}

	if m.LastKnownGood != nil {
		if err := m.LastKnownGood.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lastKnownGood")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPICoreV1NodeConfigStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPICoreV1NodeConfigStatus) UnmarshalBinary(b []byte) error {
	var res IoK8sAPICoreV1NodeConfigStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
