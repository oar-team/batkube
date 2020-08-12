package translate

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/oar-team/batkube/models"
)

func JobToPod(job Job, simData SimulationBeginsData) (error, models.IoK8sAPICoreV1Pod) {
	jobIdsplit := strings.Split(job.Id, "!")
	if len(jobIdsplit) <= 1 {
		return errors.New("Job id must be of the form worklaod!jobId, as in a JOB_SUBMITTED event"), models.IoK8sAPICoreV1Pod{}
	}
	wl := jobIdsplit[0]
	prof, ok := simData.Profiles[wl][job.Profile]
	if !ok {
		return errors.New(fmt.Sprintf("Could not find profile %s for workload %s\n", job.Profile, wl)), models.IoK8sAPICoreV1Pod{}
	}
	var pod models.IoK8sAPICoreV1Pod

	var scheduler string
	if prof.Specs["scheduler"] != nil {
		scheduler = prof.Specs["scheduler"].(string)
	}

	requests := make(map[string]models.IoK8sApimachineryPkgAPIResourceQuantity, 0)
	if cpu := prof.Specs["cpu"]; cpu != nil {
		switch cpu.(type) {
		case float64:
			cpu = fmt.Sprintf("%f", cpu)
		}
		requests["cpu"] = models.IoK8sApimachineryPkgAPIResourceQuantity(cpu.(string))
	}

	switch prof.Type {
	case "delay":
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
						Name:  &job.Profile,
						Image: "nginx",
						Command: []string{
							"/bin/sleep",
							fmt.Sprintf("%f", prof.Specs["delay"].(float64)),
						},
						Resources: &models.IoK8sAPICoreV1ResourceRequirements{
							Requests: requests,
							Limits:   requests,
						},
					},
				},
			},
			Status: &models.IoK8sAPICoreV1PodStatus{
				Phase: "Pending",
			},
		}

	default:
		return errors.New(fmt.Sprintf("[translate] I don't know this profile type : %s", prof.Type)), models.IoK8sAPICoreV1Pod{}
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
		var pods = "110" // Default value
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
						Type:               &readyConditionType,
						Status:             &trueConditionStatus,
						LastTransitionTime: &models.IoK8sApimachineryPkgApisMetaV1Time{},
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
