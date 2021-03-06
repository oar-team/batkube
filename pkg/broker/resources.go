package broker

import (
	"fmt"
	"reflect"

	"github.com/oar-team/batkube/models"
	"github.com/oar-team/batkube/pkg/translate"
)

// TODO: move all this under broker struct defined in broker.go

// Resources protected by getters and setters
var events []*models.IoK8sApimachineryPkgApisMetaV1WatchEvent

// Discard events older than most recent resourceVersion - maxResourceAge to free memory.
var maxResourceAge int = 100

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
var CoreV1EventList models.IoK8sAPICoreV1EventList
var EventV1beta1EventList models.IoK8sAPIEventsV1beta1EventList
var ReplicaSetList models.IoK8sAPIAppsV1ReplicaSetList
var StatefulSetList models.IoK8sAPIAppsV1StatefulSetList
var ReplicationControllerList models.IoK8sAPICoreV1ReplicationControllerList

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

var simData translate.SimulationBeginsData = translate.SimulationBeginsData{}

// Pods scheduled and ready to be executed are sent over this channel to be
// retrieved by te broker and sent to Batsim.
var ToExecute = make(chan *models.IoK8sAPICoreV1Pod)

var APIGroupsMap = make(map[string][]string)

func InitResources() {
	NodeList = models.IoK8sAPICoreV1NodeList{
		Kind:       "NodeList",
		APIVersion: "v1",
		Items:      []*models.IoK8sAPICoreV1Node{},
		Metadata:   &models.IoK8sApimachineryPkgApisMetaV1ListMeta{},
	}
	createAPIResource("v1", "nodes", "Node", false, []string{create, deleteVerb, deletecollection, get, list, patch, update, watch})
	PodList = models.IoK8sAPICoreV1PodList{
		Kind:       "PodList",
		APIVersion: "v1",
		Items:      []*models.IoK8sAPICoreV1Pod{},
		Metadata:   &models.IoK8sApimachineryPkgApisMetaV1ListMeta{},
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
	createAPIResource("v1", "persistentvolumes", "PersistentVolume", false, []string{create, deleteVerb, deletecollection, get, list, patch, update, watch})
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
	ReplicaSetList = models.IoK8sAPIAppsV1ReplicaSetList{
		Kind:       "ReplicaSetList",
		APIVersion: "apps/v1",
		Items:      []*models.IoK8sAPIAppsV1ReplicaSet{},
	}
	createAPIResource("apps/v1", "replicasets", "ReplicaSet", true, []string{create, deleteVerb, deletecollection, get, list, patch, update, watch})
	StatefulSetList = models.IoK8sAPIAppsV1StatefulSetList{
		Kind:       "StatefulSetList",
		APIVersion: "apps/v1",
		Items:      []*models.IoK8sAPIAppsV1StatefulSet{},
	}
	createAPIResource("apps/v1", "statefulsets", "StatefulSet", true, []string{create, deleteVerb, deletecollection, get, list, patch, update, watch})
	ReplicationControllerList = models.IoK8sAPICoreV1ReplicationControllerList{
		Kind:       "ReplicationController",
		APIVersion: "v1",
		Items:      []*models.IoK8sAPICoreV1ReplicationController{},
	}
	createAPIResource("v1", "replicationcontrollers", "ReplicationController", true, []string{create, deleteVerb, deletecollection, get, list, patch, update, watch})

	// Made this loop to factorize a bit. It is still hard coded, though
	APIGroupsMap["policy"] = []string{"v1beta1"}
	APIGroupsMap["storage.k8s.io"] = []string{"v1"}
	APIGroupsMap["coordination.k8s.io"] = []string{"v1", "v1beta1"}
	APIGroupsMap["events.k8s.io"] = []string{"v1beta1"}
	APIGroupsMap["apps"] = []string{"v1"}
	APIGroupList = models.IoK8sApimachineryPkgApisMetaV1APIGroupList{
		Kind:       "APIGroupList",
		APIVersion: "v1",
		Groups:     []*models.IoK8sApimachineryPkgApisMetaV1APIGroup{},
	}
	for groupName, versions := range APIGroupsMap {
		// We need to copy the string
		groupNameStr := groupName
		group := models.IoK8sApimachineryPkgApisMetaV1APIGroup{Name: &groupNameStr}
		for _, version := range versions {
			groupVersion := groupName + "/" + version
			group.Versions = append(group.Versions, &models.IoK8sApimachineryPkgApisMetaV1GroupVersionForDiscovery{
				GroupVersion: &groupVersion,
				Version:      &version,
			})
		}
		group.PreferredVersion = group.Versions[0]
		APIGroupList.Groups = append(APIGroupList.Groups, &group)
	}
}

func AddEvent(eventType *string, object interface{}) {
	//v := reflect.ValueOf(object)
	//if v.Kind() != reflect.Ptr {
	//	object = v.Addr().Interface()
	//}

	// Hard coded, temporary solution to the deepcopy problem
	switch object.(type) {
	case *models.IoK8sAPICoreV1Pod:
		var objectCopy models.IoK8sAPICoreV1Pod
		if err := DeepCopy(object, &objectCopy); err != nil {
			panic(fmt.Sprintf("Error while adding event to EventList : %s", err))
		}
		events = append(events, &models.IoK8sApimachineryPkgApisMetaV1WatchEvent{
			Type:   eventType,
			Object: objectCopy,
		})
	case *models.IoK8sAPICoreV1Node:
		var objectCopy models.IoK8sAPICoreV1Node
		if err := DeepCopy(object, &objectCopy); err != nil {
			panic(fmt.Sprintf("Error while adding event to EventList : %s", err))
		}
		events = append(events, &models.IoK8sApimachineryPkgApisMetaV1WatchEvent{
			Type:   eventType,
			Object: objectCopy,
		})
	case *models.IoK8sAPICoreV1NodeList:
		var objectCopy models.IoK8sAPICoreV1NodeList
		if err := DeepCopy(object, &objectCopy); err != nil {
			panic(fmt.Sprintf("Error while adding event to EventList : %s", err))
		}
		events = append(events, &models.IoK8sApimachineryPkgApisMetaV1WatchEvent{
			Type:   eventType,
			Object: objectCopy,
		})
	case *models.IoK8sAPICoreV1PodList:
		var objectCopy models.IoK8sAPICoreV1PodList
		if err := DeepCopy(object, &objectCopy); err != nil {
			panic(fmt.Sprintf("Error while adding event to EventList : %s", err))
		}
		events = append(events, &models.IoK8sApimachineryPkgApisMetaV1WatchEvent{
			Type:   eventType,
			Object: objectCopy,
		})
	}
}

func GetEvents() []*models.IoK8sApimachineryPkgApisMetaV1WatchEvent {
	return events
}

/*
Updates LastProbeTime for pods containers and LastHeartbeatTime for nodes.
Attempt at resolving the resource update issues
*/
func UpdateProbeAndHeartbeatTimes(now float64) {
	currentTime := translate.BatsimNowToMetaV1Time(now)
	for _, pod := range PodList.Items {
		for _, condition := range pod.Status.Conditions {
			condition.LastProbeTime = &currentTime
		}
		IncrementResourceVersion(pod.Metadata)
	}
	for _, node := range NodeList.Items {
		for _, condition := range node.Status.Conditions {
			condition.LastHeartbeatTime = &currentTime
		}
		IncrementResourceVersion(node.Metadata)
	}
}

func ClearEmptyEvents() {
	toRemove := make([]int, 0)
	for i, event := range events {
		if event.Type == nil {
			toRemove = append(toRemove, i)
		}
	}
	deleteEvents(toRemove)
}

/*
Warning : not usable in its current state.
The fact that resources are stored in non thread safe data structures make this code really vulnerable
*/
func ClearOldEvents() {
	// First pass : get the most recent resourceVersion.
	var maxResourceVersion string = "0"
	for _, event := range events {
		meta, err := getFieldByName(reflect.ValueOf(event.Object), "Metadata")
		if err != nil {
			panic(err)
		}
		rvValue, err := getFieldByName(meta, "ResourceVersion")
		if err != nil {
			panic(err)
		}
		rv := rvValue.String()
		if compareStr(rv, maxResourceVersion) {
			maxResourceVersion = rv
		}
	}

	// Second pass : determine which events need to be cleared
	toRemove := make([]int, 0)
	for i, event := range events {
		meta, err := getFieldByName(reflect.ValueOf(event.Object), "Metadata")
		if err != nil {
			panic(err)
		}
		rvValue, err := getFieldByName(meta, "ResourceVersion")
		if err != nil {
			panic(err)
		}
		rv := rvValue.String()
		if compareStr(maxResourceVersion, incrementStr(rv, maxResourceAge)) {
			toRemove = append(toRemove, i)
		}
	}
	deleteEvents(toRemove)
}

func deleteEvents(toRemove []int) {
	last := len(toRemove) - 1
	for i := range toRemove {
		reverse_i := toRemove[last-i]
		events = append(
			events[:reverse_i],
			events[reverse_i+1:]...,
		)
	}
}
