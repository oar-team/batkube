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

// IoK8sAPICoreV1NodeSystemInfo NodeSystemInfo is a set of ids/uuids to uniquely identify the node.
//
// swagger:model io.k8s.api.core.v1.NodeSystemInfo
type IoK8sAPICoreV1NodeSystemInfo struct {

	// The Architecture reported by the node
	// Required: true
	Architecture *string `json:"architecture"`

	// Boot ID reported by the node.
	// Required: true
	BootID *string `json:"bootID"`

	// ContainerRuntime Version reported by the node through runtime remote API (e.g. docker://1.5.0).
	// Required: true
	ContainerRuntimeVersion *string `json:"containerRuntimeVersion"`

	// Kernel Version reported by the node from 'uname -r' (e.g. 3.16.0-0.bpo.4-amd64).
	// Required: true
	KernelVersion *string `json:"kernelVersion"`

	// KubeProxy Version reported by the node.
	// Required: true
	KubeProxyVersion *string `json:"kubeProxyVersion"`

	// Kubelet Version reported by the node.
	// Required: true
	KubeletVersion *string `json:"kubeletVersion"`

	// MachineID reported by the node. For unique machine identification in the cluster this field is preferred. Learn more from man(5) machine-id: http://man7.org/linux/man-pages/man5/machine-id.5.html
	// Required: true
	MachineID *string `json:"machineID"`

	// The Operating System reported by the node
	// Required: true
	OperatingSystem *string `json:"operatingSystem"`

	// OS Image reported by the node from /etc/os-release (e.g. Debian GNU/Linux 7 (wheezy)).
	// Required: true
	OsImage *string `json:"osImage"`

	// SystemUUID reported by the node. For unique machine identification MachineID is preferred. This field is specific to Red Hat hosts https://access.redhat.com/documentation/en-US/Red_Hat_Subscription_Management/1/html/RHSM/getting-system-uuid.html
	// Required: true
	SystemUUID *string `json:"systemUUID"`
}

// Validate validates this io k8s api core v1 node system info
func (m *IoK8sAPICoreV1NodeSystemInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateArchitecture(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBootID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContainerRuntimeVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKernelVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKubeProxyVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKubeletVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMachineID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperatingSystem(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOsImage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSystemUUID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPICoreV1NodeSystemInfo) validateArchitecture(formats strfmt.Registry) error {

	if err := validate.Required("architecture", "body", m.Architecture); err != nil {
		return err
	}

	return nil
}

func (m *IoK8sAPICoreV1NodeSystemInfo) validateBootID(formats strfmt.Registry) error {

	if err := validate.Required("bootID", "body", m.BootID); err != nil {
		return err
	}

	return nil
}

func (m *IoK8sAPICoreV1NodeSystemInfo) validateContainerRuntimeVersion(formats strfmt.Registry) error {

	if err := validate.Required("containerRuntimeVersion", "body", m.ContainerRuntimeVersion); err != nil {
		return err
	}

	return nil
}

func (m *IoK8sAPICoreV1NodeSystemInfo) validateKernelVersion(formats strfmt.Registry) error {

	if err := validate.Required("kernelVersion", "body", m.KernelVersion); err != nil {
		return err
	}

	return nil
}

func (m *IoK8sAPICoreV1NodeSystemInfo) validateKubeProxyVersion(formats strfmt.Registry) error {

	if err := validate.Required("kubeProxyVersion", "body", m.KubeProxyVersion); err != nil {
		return err
	}

	return nil
}

func (m *IoK8sAPICoreV1NodeSystemInfo) validateKubeletVersion(formats strfmt.Registry) error {

	if err := validate.Required("kubeletVersion", "body", m.KubeletVersion); err != nil {
		return err
	}

	return nil
}

func (m *IoK8sAPICoreV1NodeSystemInfo) validateMachineID(formats strfmt.Registry) error {

	if err := validate.Required("machineID", "body", m.MachineID); err != nil {
		return err
	}

	return nil
}

func (m *IoK8sAPICoreV1NodeSystemInfo) validateOperatingSystem(formats strfmt.Registry) error {

	if err := validate.Required("operatingSystem", "body", m.OperatingSystem); err != nil {
		return err
	}

	return nil
}

func (m *IoK8sAPICoreV1NodeSystemInfo) validateOsImage(formats strfmt.Registry) error {

	if err := validate.Required("osImage", "body", m.OsImage); err != nil {
		return err
	}

	return nil
}

func (m *IoK8sAPICoreV1NodeSystemInfo) validateSystemUUID(formats strfmt.Registry) error {

	if err := validate.Required("systemUUID", "body", m.SystemUUID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPICoreV1NodeSystemInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPICoreV1NodeSystemInfo) UnmarshalBinary(b []byte) error {
	var res IoK8sAPICoreV1NodeSystemInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
