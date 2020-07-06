package translate

import (
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
}
