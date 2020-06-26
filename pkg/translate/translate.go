package translate

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/models"
)

func JobToPod(job Job, simData SimulationBeginsData) (error, models.IoK8sAPICoreV1Pod) {
	prof := simData.Profiles[strings.Split(job.Id, "!")[0]][job.Profile]
	var pod models.IoK8sAPICoreV1Pod

	var scheduler string
	if prof.Specs["scheduler"] != nil {
		scheduler = prof.Specs["scheduler"].(string)
	}

	readyConditionType := "Ready"
	trueConditionStatus := "True"
	switch prof.Type {
	case "delay":
		containerName := "sleep"
		jobSubtime := BatsimNowToMetaV1Time(job.Subtime)
		pod = models.IoK8sAPICoreV1Pod{
			Kind:       "Pod",
			APIVersion: "v1",
			Metadata: &models.IoK8sApimachineryPkgApisMetaV1ObjectMeta{
				Name:              strings.ReplaceAll(job.Id, "!", "-"),
				Namespace:         "default",
				ResourceVersion:   "0",
				UID:               uuid.New().String(),
				CreationTimestamp: &jobSubtime,
			},
			Spec: &models.IoK8sAPICoreV1PodSpec{
				SchedulerName: scheduler,
				NodeName:      "",
				Containers: []*models.IoK8sAPICoreV1Container{
					{
						Name:  &containerName,
						Image: "tutum/curl",
						Command: []string{
							"/bin/sleep",
							fmt.Sprintf("%f", prof.Specs["delay"].(float64)),
						},
					},
				},
			},
			Status: &models.IoK8sAPICoreV1PodStatus{
				Phase: "Pending",
				Conditions: []*models.IoK8sAPICoreV1PodCondition{
					{
						Type:   &readyConditionType,
						Status: &trueConditionStatus,
					},
				},
			},
		}

	default:
		log.Fatalf("[translate] I don't know this profile type : %s", prof.Type)
	}
	return nil, pod
}

func ComputeResourcesToNodes(simData SimulationBeginsData) (error, []*models.IoK8sAPICoreV1Node) {
	var nodes []*models.IoK8sAPICoreV1Node
	readyConditionType := "Ready"
	trueConditionStatus := "True"
	for _, resource := range simData.ComputeResources {
		// Capacity of the pod
		e, ok := resource.Properties["memory"].(string)
		var memory models.IoK8sApimachineryPkgAPIResourceQuantity
		if ok {
			memory = models.IoK8sApimachineryPkgAPIResourceQuantity(e)
		}
		e, ok = resource.Properties["core"].(string)
		var core models.IoK8sApimachineryPkgAPIResourceQuantity
		if ok {
			core = models.IoK8sApimachineryPkgAPIResourceQuantity(e)
		}
		// Node cpu speed is not a thing in Kubernetes
		var capacity = make(map[string]models.IoK8sApimachineryPkgAPIResourceQuantity, 0)
		capacity["memory"] = memory
		capacity["cpu"] = core
		var pods = "1"
		//if simData.AllowComputeSharing {
		//	pods = "110"
		//}
		capacity["pods"] = models.IoK8sApimachineryPkgAPIResourceQuantity(pods)

		var node models.IoK8sAPICoreV1Node = models.IoK8sAPICoreV1Node{
			Kind:       "Node",
			APIVersion: "v1",
			Metadata: &models.IoK8sApimachineryPkgApisMetaV1ObjectMeta{
				Name:              strconv.Itoa(resource.Id) + "-" + resource.Name,
				ResourceVersion:   "0",
				UID:               uuid.New().String(),
				CreationTimestamp: &models.IoK8sApimachineryPkgApisMetaV1Time{},
			},
			Status: &models.IoK8sAPICoreV1NodeStatus{
				Conditions: []*models.IoK8sAPICoreV1NodeCondition{
					{
						// Is the initial state of the
						// resource always idle?
						Type:   &readyConditionType,
						Status: &trueConditionStatus,
					},
				},
				Capacity:    capacity,
				Allocatable: capacity,
			},
		}
		nodes = append(nodes, &node)
	}
	return nil, nodes
}

func PodToExecuteJobData(pod *models.IoK8sAPICoreV1Pod) ExecuteJobData {
	return ExecuteJobData{
		JobId: GetJobIdFromPodName(pod.Metadata.Name),
		Alloc: strings.Split(pod.Spec.NodeName, "-")[0],
	}
}
