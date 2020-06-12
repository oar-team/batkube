package translate

import (
	"fmt"
	"strings"

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

	switch prof.Type {
	case "delay":
		containerName := "sleep"
		pod = models.IoK8sAPICoreV1Pod{
			Kind:       "Pod",
			APIVersion: "v1",
			Metadata: &models.IoK8sApimachineryPkgApisMetaV1ObjectMeta{
				Name:            strings.ReplaceAll(job.Id, "!", "-"),
				Namespace:       "default",
				ResourceVersion: "0",
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
			},
		}

	default:
		log.Fatalf("[translate] I don't know this profile type : %s", prof.Type)
	}
	return nil, pod
}

func ComputeResourcesToNodes(resources []ComputeResource) (error, []*models.IoK8sAPICoreV1Node) {
	var nodes []*models.IoK8sAPICoreV1Node
	var node models.IoK8sAPICoreV1Node
	nodeConditionType := "Ready"
	nodeConditionStatus := "True"
	for _, resource := range resources {
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
		//e, ok = resource.Properties["speed"].(string)
		//var speed models.IoK8sApimachineryPkgAPIResourceQuantity
		//if ok {
		//	speed = models.IoK8sApimachineryPkgAPIResourceQuantity(e)
		//}

		node = models.IoK8sAPICoreV1Node{
			Kind:       "Node",
			APIVersion: "v1",
			Metadata: &models.IoK8sApimachineryPkgApisMetaV1ObjectMeta{
				Name:            fmt.Sprintf("%v", resource.Id) + "-" + resource.Name,
				ResourceVersion: "0",
			},
			Status: &models.IoK8sAPICoreV1NodeStatus{
				Capacity: map[string]models.IoK8sApimachineryPkgAPIResourceQuantity{
					"memory": memory,
					"cpu":    core,
				},
				Conditions: []*models.IoK8sAPICoreV1NodeCondition{
					{
						// Is the initial state of the
						// resource always idle? -> To
						// be discussed with batsim devs
						Type:   &nodeConditionType,
						Status: &nodeConditionStatus,
					},
				},
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
