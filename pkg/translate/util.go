package translate

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/mitchellh/mapstructure"
	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/models"
)

func GetJobIdFromPodName(name string) string {
	return strings.ReplaceAll(name, "-", "!")
}

func GetPodNameFromJobId(name string) string {
	return strings.ReplaceAll(name, "!", "-")
}

/*
Wrapper function to help with Batsim event creation
*/
func MakeEvent(timestamp float64, eventType string, data interface{}) (error, Event) {
	var eventData map[string]interface{}
	err := mapstructure.Decode(data, &eventData)
	if err != nil {
		return err, Event{}
	}
	return nil, Event{
		Timestamp: timestamp,
		Type:      eventType,
		Data:      eventData,
	}
}

func BatsimNowToTime(now float64) time.Time {
	return time.Unix(0, int64(now*1e9))
}

func BatsimNowToMetaV1Time(now float64) models.IoK8sApimachineryPkgApisMetaV1Time {
	return models.IoK8sApimachineryPkgApisMetaV1Time(BatsimNowToTime(now))
}

func UpdatePodStatusForScheduling(pod *models.IoK8sAPICoreV1Pod, now models.IoK8sApimachineryPkgApisMetaV1Time) {
	trueStr := "True"
	scheduledStr := "PodScheduled"
	readyStr := "Ready"
	containersReadyStr := "ContainersReady"
	initializedStr := "Initialized"
	pod.Status.StartTime = &now
	pod.Status.Conditions = []*models.IoK8sAPICoreV1PodCondition{
		{
			Type:               &scheduledStr,
			Status:             &trueStr,
			LastTransitionTime: &now,
		},
		{
			Type:               &readyStr,
			Status:             &trueStr,
			LastTransitionTime: &now,
		},
		{
			Type:               &containersReadyStr,
			Status:             &trueStr,
			LastTransitionTime: &now,
		},
		{
			Type:               &initializedStr,
			Status:             &trueStr,
			LastTransitionTime: &now,
		},
	}
	boolTrue := true
	pod.Status.ContainerStatuses = []*models.IoK8sAPICoreV1ContainerStatus{
		{
			Name:    pod.Spec.Containers[0].Name,
			Ready:   &boolTrue,
			Started: true,
			State: &models.IoK8sAPICoreV1ContainerState{
				Running: &models.IoK8sAPICoreV1ContainerStateRunning{
					StartedAt: &now,
				},
			},
		},
	}

	pod.Status.Phase = "Running"
	IncrementResourceVersion(pod.Metadata)
}

func IncrementResourceVersion(metadata interface{}) {
	// Piece of code for experiments on resourceVersions

	//PodList.Metadata.ResourceVersion = incrementStr(PodList.Metadata.ResourceVersion)
	//NodeList.Metadata.ResourceVersion = incrementStr(NodeList.Metadata.ResourceVersion)
	//for _, pod := range PodList.Items {
	//	pod.Metadata.ResourceVersion = incrementStr(pod.Metadata.ResourceVersion)
	//}
	//for _, node := range NodeList.Items {
	//	node.Metadata.ResourceVersion = incrementStr(node.Metadata.ResourceVersion)
	//}

	switch metadata.(type) {
	case *models.IoK8sApimachineryPkgApisMetaV1ObjectMeta:
		meta := metadata.(*models.IoK8sApimachineryPkgApisMetaV1ObjectMeta)
		if meta.ResourceVersion == "" {
			meta.ResourceVersion = "0"
		}
		resourceVersion, err := strconv.Atoi(meta.ResourceVersion)
		if err != nil {
			log.Panic(err)
		}
		meta.ResourceVersion = fmt.Sprintf("%d", resourceVersion+1)
	case *models.IoK8sApimachineryPkgApisMetaV1ListMeta:
		meta := metadata.(*models.IoK8sApimachineryPkgApisMetaV1ListMeta)
		if meta.ResourceVersion == "" {
			meta.ResourceVersion = "0"
		}
		resourceVersion, err := strconv.Atoi(meta.ResourceVersion)
		if err != nil {
			log.Panic(err)
		}
		meta.ResourceVersion = fmt.Sprintf("%d", resourceVersion+1)
	default:
		panic(fmt.Sprintf("Unknown metadata type : %T", metadata))
	}
}
