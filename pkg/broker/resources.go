package broker

import (
	"github.com/pkg/errors"
	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/models"
	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/pkg/translate"
)

// Resources protected by getters and setters
var events []models.IoK8sApimachineryPkgApisMetaV1WatchEvent

var NodeList models.IoK8sAPICoreV1NodeList = models.IoK8sAPICoreV1NodeList{
	Kind:       "NodeList",
	APIVersion: "v1",
	Items:      []*models.IoK8sAPICoreV1Node{},
}
var PodList models.IoK8sAPICoreV1PodList = models.IoK8sAPICoreV1PodList{
	Kind:       "PodList",
	APIVersion: "v1",
	Items:      []*models.IoK8sAPICoreV1Pod{},
}
var PodDisruptionBudgetList models.IoK8sAPIPolicyV1beta1PodDisruptionBudgetList = models.IoK8sAPIPolicyV1beta1PodDisruptionBudgetList{
	Kind:       "PodDisruptionBudgetList",
	APIVersion: "policy/v1beta1",
	Items:      []*models.IoK8sAPIPolicyV1beta1PodDisruptionBudget{},
}
var StorageClassList models.IoK8sAPIStorageV1StorageClassList = models.IoK8sAPIStorageV1StorageClassList{
	Kind:       "StorageClassList",
	APIVersion: "storage.k8s.io/v1",
	Items:      []*models.IoK8sAPIStorageV1StorageClass{},
}
var CSINodeList models.IoK8sAPIStorageV1CSINodeList = models.IoK8sAPIStorageV1CSINodeList{
	Kind:       "CSINodeList",
	APIVersion: "storage.k8s.io/v1",
	Items:      []*models.IoK8sAPIStorageV1CSINode{},
}
var PersistentVolumeClaimList models.IoK8sAPICoreV1PersistentVolumeClaimList = models.IoK8sAPICoreV1PersistentVolumeClaimList{
	Kind:       "PersistentVolumeClaimList",
	APIVersion: "v1",
	Items:      []*models.IoK8sAPICoreV1PersistentVolumeClaim{},
}
var PersistentVolumeList models.IoK8sAPICoreV1PersistentVolumeList = models.IoK8sAPICoreV1PersistentVolumeList{
	Kind:       "PersistentVolumeList",
	APIVersion: "v1",
	Items:      []*models.IoK8sAPICoreV1PersistentVolume{},
}
var ServiceList models.IoK8sAPICoreV1ServiceList = models.IoK8sAPICoreV1ServiceList{
	Kind:       "ServiceList",
	APIVersion: "v1",
	Items:      []*models.IoK8sAPICoreV1Service{},
}
var APIResourceList models.IoK8sApimachineryPkgApisMetaV1APIResourceList = models.IoK8sApimachineryPkgApisMetaV1APIResourceList{
	Kind:       "APIResourceList",
	APIVersion: "v1",
	Resources:  []*models.IoK8sApimachineryPkgApisMetaV1APIResource{},
}
var EndpointsList models.IoK8sAPICoreV1EndpointsList = models.IoK8sAPICoreV1EndpointsList{
	Kind:       "EndpointsList",
	APIVersion: "v1",
	Items:      []*models.IoK8sAPICoreV1Endpoints{},
}

var SimData translate.SimulationBeginsData

var ToExecute = make(chan *models.IoK8sAPICoreV1Pod)

// TODO AddEvent(type string, object interface{}) pour pouvoir faire un check
// sur l'event type (et simplifier la procédure)
func AddEvent(event models.IoK8sApimachineryPkgApisMetaV1WatchEvent) {
	events = append(events, event)
}

func GetEvents() []models.IoK8sApimachineryPkgApisMetaV1WatchEvent {
	return events
}

func GetPod(name string) (*models.IoK8sAPICoreV1Pod, error) {
	var pod *models.IoK8sAPICoreV1Pod
	for _, pod = range PodList.Items {
		if pod.Metadata.Name == name {
			break
		}
	}
	if pod == nil || pod.Metadata.Name != name {
		return nil, errors.Errorf("Could not find pod %s", name)
	}
	return pod, nil
}

func GetNode(name string) (*models.IoK8sAPICoreV1Node, error) {
	var node *models.IoK8sAPICoreV1Node
	for _, node = range NodeList.Items {
		if node.Metadata.Name == name {
			break
		}
	}
	if node == nil || node.Metadata.Name != name {
		return nil, errors.Errorf("Could not find node %s", name)
	}
	return node, nil
}

func GetEndpoint(name string, namespace string) (*models.IoK8sAPICoreV1Endpoints, error) {
	if EndpointsList.Items == nil {
		return nil, errors.Errorf("No endpoints yet listed on the API server")
	}
	var r *models.IoK8sAPICoreV1Endpoints
	for _, r = range EndpointsList.Items {
		if r.Metadata.Name == name && r.Metadata.Namespace == namespace {
			break
		}
	}
	if r == nil || r.Metadata.Name != name || r.Metadata.Namespace != namespace {
		return nil, errors.Errorf("Could not find enpoints %s", name)
	}
	return r, nil
}
