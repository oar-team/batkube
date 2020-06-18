package broker

import (
	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/models"
	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/pkg/translate"
)

// Resources protected by getters and setters
var events []*models.IoK8sApimachineryPkgApisMetaV1WatchEvent

// Kubernetes resources
var APIGroupList models.IoK8sApimachineryPkgApisMetaV1APIGroupList
var NodeList models.IoK8sAPICoreV1NodeList
var PodList models.IoK8sAPICoreV1PodList
var PodDisruptionBudgetList models.IoK8sAPIPolicyV1beta1PodDisruptionBudgetList
var StorageClassList models.IoK8sAPIStorageV1StorageClassList
var CSINodeList models.IoK8sAPIStorageV1CSINodeList
var PersistentVolumeClaimList models.IoK8sAPICoreV1PersistentVolumeClaimList
var PersistentVolumeList models.IoK8sAPICoreV1PersistentVolumeList
var ServiceList models.IoK8sAPICoreV1ServiceList
var EndpointsList models.IoK8sAPICoreV1EndpointsList
var LeaseList models.IoK8sAPICoordinationV1LeaseList
var CoreV1EventList models.IoK8sAPICoreV1EventList // Deprecated, here for backwards compatibility
var EventV1beta1EventList models.IoK8sAPIEventsV1beta1EventList

// APIResources, sorted by groupVersion
var apiResources = make(map[string][]*models.IoK8sApimachineryPkgApisMetaV1APIResource)

// List of verbs
const create = "create"
const deleteVerb = "delete"
const deletecollection = "deletecollection"
const get = "get"
const list = "list"
const patch = "patch"
const update = "update"
const watch = "watch"

var SimData translate.SimulationBeginsData

// Pods scheduled and ready to be executed are sent over this channel to be
// retrieved by te broker and sent to Batsim.
var ToExecute = make(chan *models.IoK8sAPICoreV1Pod)

var APIGroupsMap = make(map[string][]string)

func InitResources() {
	NodeList = models.IoK8sAPICoreV1NodeList{
		Kind:       "NodeList",
		APIVersion: "v1",
		Items:      []*models.IoK8sAPICoreV1Node{},
	}
	createAPIResource("v1", "nodes", "Node", false, []string{create, deleteVerb, deletecollection, get, list, patch, update, watch})
	PodList = models.IoK8sAPICoreV1PodList{
		Kind:       "PodList",
		APIVersion: "v1",
		Items:      []*models.IoK8sAPICoreV1Pod{},
	}
	createAPIResource("v1", "pods", "Pod", true, []string{create, deleteVerb, deletecollection, get, list, patch, update, watch})
	PodDisruptionBudgetList = models.IoK8sAPIPolicyV1beta1PodDisruptionBudgetList{
		Kind:       "PodDisruptionBudgetList",
		APIVersion: "policy/v1beta1",
		Items:      []*models.IoK8sAPIPolicyV1beta1PodDisruptionBudget{},
	}
	createAPIResource("policy/v1beta1", "poddisruptionbudgets", "PodDisruptionBudget", true, []string{create, deleteVerb, deletecollection, get, list, patch, update, watch})
	StorageClassList = models.IoK8sAPIStorageV1StorageClassList{
		Kind:       "StorageClassList",
		APIVersion: "storage.k8s.io/v1",
		Items:      []*models.IoK8sAPIStorageV1StorageClass{},
	}
	createAPIResource("storage.k8s.io/v1", "storageclasses", "StorageClass", false, []string{create, deleteVerb, deletecollection, get, list, patch, update, watch})
	CSINodeList = models.IoK8sAPIStorageV1CSINodeList{
		Kind:       "CSINodeList",
		APIVersion: "storage.k8s.io/v1",
		Items:      []*models.IoK8sAPIStorageV1CSINode{},
	}
	createAPIResource("storage.k8s.io/v1", "csinodes", "CSINode", false, []string{create, deleteVerb, deletecollection, get, list, patch, update, watch})
	PersistentVolumeClaimList = models.IoK8sAPICoreV1PersistentVolumeClaimList{
		Kind:       "PersistentVolumeClaimList",
		APIVersion: "v1",
		Items:      []*models.IoK8sAPICoreV1PersistentVolumeClaim{},
	}
	createAPIResource("v1", "persistentvolumeclaims", "PersistentVolumeClaim", true, []string{create, deleteVerb, deletecollection, get, list, patch, update, watch})
	PersistentVolumeList = models.IoK8sAPICoreV1PersistentVolumeList{
		Kind:       "PersistentVolumeList",
		APIVersion: "v1",
		Items:      []*models.IoK8sAPICoreV1PersistentVolume{},
	}
	createAPIResource("v1", "persistentvolume", "PersistentVolume", false, []string{create, deleteVerb, deletecollection, get, list, patch, update, watch})
	ServiceList = models.IoK8sAPICoreV1ServiceList{
		Kind:       "ServiceList",
		APIVersion: "v1",
		Items:      []*models.IoK8sAPICoreV1Service{},
	}
	createAPIResource("v1", "services", "Service", true, []string{create, deleteVerb, get, list, patch, update, watch})
	EndpointsList = models.IoK8sAPICoreV1EndpointsList{
		Kind:       "EndpointsList",
		APIVersion: "v1",
		Items:      []*models.IoK8sAPICoreV1Endpoints{},
	}
	createAPIResource("v1", "endpoints", "Endpoints", true, []string{create, deleteVerb, deletecollection, get, list, patch, update, watch})
	LeaseList = models.IoK8sAPICoordinationV1LeaseList{
		Kind:       "LeaseList",
		APIVersion: "coordination.k8s.io/v1",
		Items:      []*models.IoK8sAPICoordinationV1Lease{},
	}
	createAPIResource("coordination.k8s.io/v1", "leases", "Lease", true, []string{create, deleteVerb, deletecollection, get, list, patch, update, watch})
	CoreV1EventList = models.IoK8sAPICoreV1EventList{
		Kind:       "EventList",
		APIVersion: "v1",
		Items:      []*models.IoK8sAPICoreV1Event{},
	}
	createAPIResource("v1", "events", "Event", true, []string{create, deleteVerb, deletecollection, get, list, patch, update, watch})
	EventV1beta1EventList = models.IoK8sAPIEventsV1beta1EventList{
		Kind:       "EventList",
		APIVersion: "events.k8s.io/v1beta1",
		Items:      []*models.IoK8sAPIEventsV1beta1Event{},
	}
	createAPIResource("events.k8s.io/v1beta1", "events", "Event", true, []string{create, deleteVerb, deletecollection, get, list, patch, update, watch})

	// Made this loop to factorize a bit. It is still hard coded, though
	APIGroupsMap["policy"] = []string{"v1beta1"}
	APIGroupsMap["storage.k8s.io"] = []string{"v1"}
	APIGroupsMap["coordination.k8s.io"] = []string{"v1"}
	APIGroupsMap["events.k8s.io"] = []string{"v1beta1"}
	APIGroupList = models.IoK8sApimachineryPkgApisMetaV1APIGroupList{
		Kind:       "APIGroupList",
		APIVersion: "v1",
		Groups:     []*models.IoK8sApimachineryPkgApisMetaV1APIGroup{},
	}
	for groupName, versions := range APIGroupsMap {
		// groupName location in memory remains the same. Setting its
		// pointer as a value directly results in every goup having the
		// same name - the last key of APIGroupsMap
		groupNameStr := groupName
		group := models.IoK8sApimachineryPkgApisMetaV1APIGroup{Name: &groupNameStr}
		for _, version := range versions {
			groupVersion := groupName + "/" + version
			group.Versions = append(group.Versions, &models.IoK8sApimachineryPkgApisMetaV1GroupVersionForDiscovery{
				GroupVersion: &groupVersion,
				Version:      &version,
			})
		}
		group.PreferredVersion = group.Versions[0] // For now there's only one version per group
		APIGroupList.Groups = append(APIGroupList.Groups, &group)
	}
}

func AddEvent(event *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) {
	events = append(events, event)
}

func GetEvents() []*models.IoK8sApimachineryPkgApisMetaV1WatchEvent {
	return events
}
