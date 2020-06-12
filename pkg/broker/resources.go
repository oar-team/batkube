package broker

import (
	"reflect"

	"github.com/pkg/errors"
	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/models"
	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/pkg/translate"
)

// Resources protected by getters and setters
var events []*models.IoK8sApimachineryPkgApisMetaV1WatchEvent

// Kubernetes resources
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
var LeaseList models.IoK8sAPICoordinationV1LeaseList = models.IoK8sAPICoordinationV1LeaseList{
	Kind:       "LeaseList",
	APIVersion: "coordination.k8s.io/v1",
	Items:      []*models.IoK8sAPICoordinationV1Lease{},
}
var EventList models.IoK8sAPICoreV1EventList = models.IoK8sAPICoreV1EventList{
	Kind:       "EventList",
	APIVersion: "v1",
	Items:      []*models.IoK8sAPICoreV1Event{},
}

var SimData translate.SimulationBeginsData

// Pods scheduled and ready to be executed are sent over this channel to be
// retrieved by te broker and sent to Batsim.
var ToExecute = make(chan *models.IoK8sAPICoreV1Pod)

func AddEvent(event *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) {
	events = append(events, event)
}

func GetEvents() []*models.IoK8sApimachineryPkgApisMetaV1WatchEvent {
	return events
}

/*
TODO : use reflection to factorise all this. At the cost of greater code complexity.
*/
func GetPod(name string) (*models.IoK8sAPICoreV1Pod, error) {
	var r *models.IoK8sAPICoreV1Pod
	for _, r = range PodList.Items {
		if r.Metadata.Name == name {
			return r, nil
		}
	}
	return nil, errors.Errorf("Could not find %s %s", reflect.TypeOf(r).String(), name)
}

func GetNode(name string) (*models.IoK8sAPICoreV1Node, error) {
	var r *models.IoK8sAPICoreV1Node
	for _, r = range NodeList.Items {
		if r.Metadata.Name == name {
			return r, nil
		}
	}
	return nil, errors.Errorf("Could not find %s %s", reflect.TypeOf(r).String(), name)
}

func GetEndpoint(name string, namespace string) (*models.IoK8sAPICoreV1Endpoints, error) {
	var r *models.IoK8sAPICoreV1Endpoints
	for _, r = range EndpointsList.Items {
		if r.Metadata.Name == name && r.Metadata.Namespace == namespace {
			return r, nil
		}
	}
	return nil, errors.Errorf("Could not find %s %s in namespace %s", reflect.TypeOf(r).String(), name, namespace)
}

func GetLease(name string, namespace string) (*models.IoK8sAPICoordinationV1Lease, error) {
	var r *models.IoK8sAPICoordinationV1Lease
	for _, r = range LeaseList.Items {
		if r.Metadata.Name == name && r.Metadata.Namespace == namespace {
			return r, nil
		}
	}
	return nil, errors.Errorf("Could not find %s %s in namespace %s", reflect.TypeOf(r).String(), name, namespace)
}
