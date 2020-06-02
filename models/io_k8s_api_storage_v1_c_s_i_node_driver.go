// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// IoK8sAPIStorageV1CSINodeDriver CSINodeDriver holds information about the specification of one CSI driver installed on a node
//
// swagger:model io.k8s.api.storage.v1.CSINodeDriver
type IoK8sAPIStorageV1CSINodeDriver struct {

	// allocatable represents the volume resources of a node that are available for scheduling. This field is beta.
	Allocatable *IoK8sAPIStorageV1VolumeNodeResources `json:"allocatable,omitempty"`

	// This is the name of the CSI driver that this object refers to. This MUST be the same name returned by the CSI GetPluginName() call for that driver.
	// Required: true
	Name *string `json:"name"`

	// nodeID of the node from the driver point of view. This field enables Kubernetes to communicate with storage systems that do not share the same nomenclature for nodes. For example, Kubernetes may refer to a given node as "node1", but the storage system may refer to the same node as "nodeA". When Kubernetes issues a command to the storage system to attach a volume to a specific node, it can use this field to refer to the node name using the ID that the storage system will understand, e.g. "nodeA" instead of "node1". This field is required.
	// Required: true
	NodeID *string `json:"nodeID"`

	// topologyKeys is the list of keys supported by the driver. When a driver is initialized on a cluster, it provides a set of topology keys that it understands (e.g. "company.com/zone", "company.com/region"). When a driver is initialized on a node, it provides the same topology keys along with values. Kubelet will expose these topology keys as labels on its own node object. When Kubernetes does topology aware provisioning, it can use this list to determine which labels it should retrieve from the node object and pass back to the driver. It is possible for different nodes to use different topology keys. This can be empty if driver does not support topology.
	TopologyKeys []string `json:"topologyKeys"`
}

// Validate validates this io k8s api storage v1 c s i node driver
func (m *IoK8sAPIStorageV1CSINodeDriver) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAllocatable(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNodeID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPIStorageV1CSINodeDriver) validateAllocatable(formats strfmt.Registry) error {

	if swag.IsZero(m.Allocatable) { // not required
		return nil
	}

	if m.Allocatable != nil {
		if err := m.Allocatable.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("allocatable")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPIStorageV1CSINodeDriver) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *IoK8sAPIStorageV1CSINodeDriver) validateNodeID(formats strfmt.Registry) error {

	if err := validate.Required("nodeID", "body", m.NodeID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPIStorageV1CSINodeDriver) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPIStorageV1CSINodeDriver) UnmarshalBinary(b []byte) error {
	var res IoK8sAPIStorageV1CSINodeDriver
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
