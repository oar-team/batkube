// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// IoK8sAPICoreV1PersistentVolumeSpec PersistentVolumeSpec is the specification of a persistent volume.
//
// swagger:model io.k8s.api.core.v1.PersistentVolumeSpec
type IoK8sAPICoreV1PersistentVolumeSpec struct {

	// AccessModes contains all ways the volume can be mounted. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes
	AccessModes []string `json:"accessModes"`

	// AWSElasticBlockStore represents an AWS Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore
	AwsElasticBlockStore *IoK8sAPICoreV1AWSElasticBlockStoreVolumeSource `json:"awsElasticBlockStore,omitempty"`

	// AzureDisk represents an Azure Data Disk mount on the host and bind mount to the pod.
	AzureDisk *IoK8sAPICoreV1AzureDiskVolumeSource `json:"azureDisk,omitempty"`

	// AzureFile represents an Azure File Service mount on the host and bind mount to the pod.
	AzureFile *IoK8sAPICoreV1AzureFilePersistentVolumeSource `json:"azureFile,omitempty"`

	// A description of the persistent volume's resources and capacity. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#capacity
	Capacity map[string]IoK8sApimachineryPkgAPIResourceQuantity `json:"capacity,omitempty"`

	// CephFS represents a Ceph FS mount on the host that shares a pod's lifetime
	Cephfs *IoK8sAPICoreV1CephFSPersistentVolumeSource `json:"cephfs,omitempty"`

	// Cinder represents a cinder volume attached and mounted on kubelets host machine. More info: https://examples.k8s.io/mysql-cinder-pd/README.md
	Cinder *IoK8sAPICoreV1CinderPersistentVolumeSource `json:"cinder,omitempty"`

	// ClaimRef is part of a bi-directional binding between PersistentVolume and PersistentVolumeClaim. Expected to be non-nil when bound. claim.VolumeName is the authoritative bind between PV and PVC. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#binding
	ClaimRef *IoK8sAPICoreV1ObjectReference `json:"claimRef,omitempty"`

	// CSI represents storage that is handled by an external CSI driver (Beta feature).
	Csi *IoK8sAPICoreV1CSIPersistentVolumeSource `json:"csi,omitempty"`

	// FC represents a Fibre Channel resource that is attached to a kubelet's host machine and then exposed to the pod.
	Fc *IoK8sAPICoreV1FCVolumeSource `json:"fc,omitempty"`

	// FlexVolume represents a generic volume resource that is provisioned/attached using an exec based plugin.
	FlexVolume *IoK8sAPICoreV1FlexPersistentVolumeSource `json:"flexVolume,omitempty"`

	// Flocker represents a Flocker volume attached to a kubelet's host machine and exposed to the pod for its usage. This depends on the Flocker control service being running
	Flocker *IoK8sAPICoreV1FlockerVolumeSource `json:"flocker,omitempty"`

	// GCEPersistentDisk represents a GCE Disk resource that is attached to a kubelet's host machine and then exposed to the pod. Provisioned by an admin. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk
	GcePersistentDisk *IoK8sAPICoreV1GCEPersistentDiskVolumeSource `json:"gcePersistentDisk,omitempty"`

	// Glusterfs represents a Glusterfs volume that is attached to a host and exposed to the pod. Provisioned by an admin. More info: https://examples.k8s.io/volumes/glusterfs/README.md
	Glusterfs *IoK8sAPICoreV1GlusterfsPersistentVolumeSource `json:"glusterfs,omitempty"`

	// HostPath represents a directory on the host. Provisioned by a developer or tester. This is useful for single-node development and testing only! On-host storage is not supported in any way and WILL NOT WORK in a multi-node cluster. More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath
	HostPath *IoK8sAPICoreV1HostPathVolumeSource `json:"hostPath,omitempty"`

	// ISCSI represents an ISCSI Disk resource that is attached to a kubelet's host machine and then exposed to the pod. Provisioned by an admin.
	Iscsi *IoK8sAPICoreV1ISCSIPersistentVolumeSource `json:"iscsi,omitempty"`

	// Local represents directly-attached storage with node affinity
	Local *IoK8sAPICoreV1LocalVolumeSource `json:"local,omitempty"`

	// A list of mount options, e.g. ["ro", "soft"]. Not validated - mount will simply fail if one is invalid. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes/#mount-options
	MountOptions []string `json:"mountOptions"`

	// NFS represents an NFS mount on the host. Provisioned by an admin. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs
	Nfs *IoK8sAPICoreV1NFSVolumeSource `json:"nfs,omitempty"`

	// NodeAffinity defines constraints that limit what nodes this volume can be accessed from. This field influences the scheduling of pods that use this volume.
	NodeAffinity *IoK8sAPICoreV1VolumeNodeAffinity `json:"nodeAffinity,omitempty"`

	// What happens to a persistent volume when released from its claim. Valid options are Retain (default for manually created PersistentVolumes), Delete (default for dynamically provisioned PersistentVolumes), and Recycle (deprecated). Recycle must be supported by the volume plugin underlying this PersistentVolume. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#reclaiming
	PersistentVolumeReclaimPolicy string `json:"persistentVolumeReclaimPolicy,omitempty"`

	// PhotonPersistentDisk represents a PhotonController persistent disk attached and mounted on kubelets host machine
	PhotonPersistentDisk *IoK8sAPICoreV1PhotonPersistentDiskVolumeSource `json:"photonPersistentDisk,omitempty"`

	// PortworxVolume represents a portworx volume attached and mounted on kubelets host machine
	PortworxVolume *IoK8sAPICoreV1PortworxVolumeSource `json:"portworxVolume,omitempty"`

	// Quobyte represents a Quobyte mount on the host that shares a pod's lifetime
	Quobyte *IoK8sAPICoreV1QuobyteVolumeSource `json:"quobyte,omitempty"`

	// RBD represents a Rados Block Device mount on the host that shares a pod's lifetime. More info: https://examples.k8s.io/volumes/rbd/README.md
	Rbd *IoK8sAPICoreV1RBDPersistentVolumeSource `json:"rbd,omitempty"`

	// ScaleIO represents a ScaleIO persistent volume attached and mounted on Kubernetes nodes.
	ScaleIO *IoK8sAPICoreV1ScaleIOPersistentVolumeSource `json:"scaleIO,omitempty"`

	// Name of StorageClass to which this persistent volume belongs. Empty value means that this volume does not belong to any StorageClass.
	StorageClassName string `json:"storageClassName,omitempty"`

	// StorageOS represents a StorageOS volume that is attached to the kubelet's host machine and mounted into the pod More info: https://examples.k8s.io/volumes/storageos/README.md
	Storageos *IoK8sAPICoreV1StorageOSPersistentVolumeSource `json:"storageos,omitempty"`

	// volumeMode defines if a volume is intended to be used with a formatted filesystem or to remain in raw block state. Value of Filesystem is implied when not included in spec.
	VolumeMode string `json:"volumeMode,omitempty"`

	// VsphereVolume represents a vSphere volume attached and mounted on kubelets host machine
	VsphereVolume *IoK8sAPICoreV1VsphereVirtualDiskVolumeSource `json:"vsphereVolume,omitempty"`
}

// Validate validates this io k8s api core v1 persistent volume spec
func (m *IoK8sAPICoreV1PersistentVolumeSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAwsElasticBlockStore(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAzureDisk(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAzureFile(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCapacity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCephfs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCinder(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClaimRef(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCsi(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFc(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFlexVolume(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFlocker(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGcePersistentDisk(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGlusterfs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHostPath(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIscsi(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocal(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNfs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNodeAffinity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePhotonPersistentDisk(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePortworxVolume(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQuobyte(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRbd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScaleIO(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStorageos(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVsphereVolume(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPICoreV1PersistentVolumeSpec) validateAwsElasticBlockStore(formats strfmt.Registry) error {

	if swag.IsZero(m.AwsElasticBlockStore) { // not required
		return nil
	}

	if m.AwsElasticBlockStore != nil {
		if err := m.AwsElasticBlockStore.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("awsElasticBlockStore")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPICoreV1PersistentVolumeSpec) validateAzureDisk(formats strfmt.Registry) error {

	if swag.IsZero(m.AzureDisk) { // not required
		return nil
	}

	if m.AzureDisk != nil {
		if err := m.AzureDisk.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("azureDisk")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPICoreV1PersistentVolumeSpec) validateAzureFile(formats strfmt.Registry) error {

	if swag.IsZero(m.AzureFile) { // not required
		return nil
	}

	if m.AzureFile != nil {
		if err := m.AzureFile.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("azureFile")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPICoreV1PersistentVolumeSpec) validateCapacity(formats strfmt.Registry) error {

	if swag.IsZero(m.Capacity) { // not required
		return nil
	}

	for k := range m.Capacity {

		if val, ok := m.Capacity[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *IoK8sAPICoreV1PersistentVolumeSpec) validateCephfs(formats strfmt.Registry) error {

	if swag.IsZero(m.Cephfs) { // not required
		return nil
	}

	if m.Cephfs != nil {
		if err := m.Cephfs.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cephfs")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPICoreV1PersistentVolumeSpec) validateCinder(formats strfmt.Registry) error {

	if swag.IsZero(m.Cinder) { // not required
		return nil
	}

	if m.Cinder != nil {
		if err := m.Cinder.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cinder")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPICoreV1PersistentVolumeSpec) validateClaimRef(formats strfmt.Registry) error {

	if swag.IsZero(m.ClaimRef) { // not required
		return nil
	}

	if m.ClaimRef != nil {
		if err := m.ClaimRef.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("claimRef")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPICoreV1PersistentVolumeSpec) validateCsi(formats strfmt.Registry) error {

	if swag.IsZero(m.Csi) { // not required
		return nil
	}

	if m.Csi != nil {
		if err := m.Csi.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("csi")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPICoreV1PersistentVolumeSpec) validateFc(formats strfmt.Registry) error {

	if swag.IsZero(m.Fc) { // not required
		return nil
	}

	if m.Fc != nil {
		if err := m.Fc.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fc")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPICoreV1PersistentVolumeSpec) validateFlexVolume(formats strfmt.Registry) error {

	if swag.IsZero(m.FlexVolume) { // not required
		return nil
	}

	if m.FlexVolume != nil {
		if err := m.FlexVolume.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("flexVolume")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPICoreV1PersistentVolumeSpec) validateFlocker(formats strfmt.Registry) error {

	if swag.IsZero(m.Flocker) { // not required
		return nil
	}

	if m.Flocker != nil {
		if err := m.Flocker.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("flocker")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPICoreV1PersistentVolumeSpec) validateGcePersistentDisk(formats strfmt.Registry) error {

	if swag.IsZero(m.GcePersistentDisk) { // not required
		return nil
	}

	if m.GcePersistentDisk != nil {
		if err := m.GcePersistentDisk.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gcePersistentDisk")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPICoreV1PersistentVolumeSpec) validateGlusterfs(formats strfmt.Registry) error {

	if swag.IsZero(m.Glusterfs) { // not required
		return nil
	}

	if m.Glusterfs != nil {
		if err := m.Glusterfs.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("glusterfs")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPICoreV1PersistentVolumeSpec) validateHostPath(formats strfmt.Registry) error {

	if swag.IsZero(m.HostPath) { // not required
		return nil
	}

	if m.HostPath != nil {
		if err := m.HostPath.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hostPath")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPICoreV1PersistentVolumeSpec) validateIscsi(formats strfmt.Registry) error {

	if swag.IsZero(m.Iscsi) { // not required
		return nil
	}

	if m.Iscsi != nil {
		if err := m.Iscsi.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("iscsi")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPICoreV1PersistentVolumeSpec) validateLocal(formats strfmt.Registry) error {

	if swag.IsZero(m.Local) { // not required
		return nil
	}

	if m.Local != nil {
		if err := m.Local.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("local")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPICoreV1PersistentVolumeSpec) validateNfs(formats strfmt.Registry) error {

	if swag.IsZero(m.Nfs) { // not required
		return nil
	}

	if m.Nfs != nil {
		if err := m.Nfs.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nfs")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPICoreV1PersistentVolumeSpec) validateNodeAffinity(formats strfmt.Registry) error {

	if swag.IsZero(m.NodeAffinity) { // not required
		return nil
	}

	if m.NodeAffinity != nil {
		if err := m.NodeAffinity.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nodeAffinity")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPICoreV1PersistentVolumeSpec) validatePhotonPersistentDisk(formats strfmt.Registry) error {

	if swag.IsZero(m.PhotonPersistentDisk) { // not required
		return nil
	}

	if m.PhotonPersistentDisk != nil {
		if err := m.PhotonPersistentDisk.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("photonPersistentDisk")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPICoreV1PersistentVolumeSpec) validatePortworxVolume(formats strfmt.Registry) error {

	if swag.IsZero(m.PortworxVolume) { // not required
		return nil
	}

	if m.PortworxVolume != nil {
		if err := m.PortworxVolume.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("portworxVolume")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPICoreV1PersistentVolumeSpec) validateQuobyte(formats strfmt.Registry) error {

	if swag.IsZero(m.Quobyte) { // not required
		return nil
	}

	if m.Quobyte != nil {
		if err := m.Quobyte.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("quobyte")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPICoreV1PersistentVolumeSpec) validateRbd(formats strfmt.Registry) error {

	if swag.IsZero(m.Rbd) { // not required
		return nil
	}

	if m.Rbd != nil {
		if err := m.Rbd.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rbd")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPICoreV1PersistentVolumeSpec) validateScaleIO(formats strfmt.Registry) error {

	if swag.IsZero(m.ScaleIO) { // not required
		return nil
	}

	if m.ScaleIO != nil {
		if err := m.ScaleIO.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("scaleIO")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPICoreV1PersistentVolumeSpec) validateStorageos(formats strfmt.Registry) error {

	if swag.IsZero(m.Storageos) { // not required
		return nil
	}

	if m.Storageos != nil {
		if err := m.Storageos.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("storageos")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPICoreV1PersistentVolumeSpec) validateVsphereVolume(formats strfmt.Registry) error {

	if swag.IsZero(m.VsphereVolume) { // not required
		return nil
	}

	if m.VsphereVolume != nil {
		if err := m.VsphereVolume.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vsphereVolume")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPICoreV1PersistentVolumeSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPICoreV1PersistentVolumeSpec) UnmarshalBinary(b []byte) error {
	var res IoK8sAPICoreV1PersistentVolumeSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}