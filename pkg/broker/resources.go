package broker

import (
	"github.com/pkg/errors"
	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/models"
	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/pkg/translate"
)

// Resources protected by getters and setters
var events []models.IoK8sApimachineryPkgApisMetaV1WatchEvent

var Nodes models.IoK8sAPICoreV1NodeList = models.IoK8sAPICoreV1NodeList{Kind: "NodeList", APIVersion: "v1"}
var Pods models.IoK8sAPICoreV1PodList = models.IoK8sAPICoreV1PodList{Kind: "PodList", APIVersion: "v1"}
var SimData translate.SimulationBeginsData

//var ToExecute PodStack
var ToExecute = make(chan *models.IoK8sAPICoreV1Pod)

//type PodStack []*models.IoK8sAPICoreV1Pod
//
//func (s *PodStack) Push(v *models.IoK8sAPICoreV1Pod) {
//	*s = append(*s, v)
//}
//
//func (s *PodStack) Pop() *models.IoK8sAPICoreV1Pod {
//	if len(*s) == 0 {
//		return nil
//	}
//	res := (*s)[len(*s)-1]
//	*s = (*s)[:len(*s)-1]
//	return res
//}

// TODO AddEvent(type string, object interface{}) pour pouvoir faire un check
// sur l'event type (et simplifier la procédure)
func AddEvent(event models.IoK8sApimachineryPkgApisMetaV1WatchEvent) {
	events = append(events, event)
}

func GetEvents() []models.IoK8sApimachineryPkgApisMetaV1WatchEvent {
	return events
}

func GetPod(podName string) (*models.IoK8sAPICoreV1Pod, error) {
	var pod *models.IoK8sAPICoreV1Pod
	for _, pod = range Pods.Items {
		if pod.Metadata.Name == podName {
			break
		}
	}
	if pod == nil || pod.Metadata.Name != podName {
		return nil, errors.Errorf("Could not find pod %s", podName)
	}
	return pod, nil
}

func GetNode(nodeName string) (*models.IoK8sAPICoreV1Node, error) {
	var node *models.IoK8sAPICoreV1Node
	for _, node = range Nodes.Items {
		if node.Metadata.Name == nodeName {
			break
		}
	}
	if node == nil || node.Metadata.Name != nodeName {
		return nil, errors.Errorf("Could not find node %s", nodeName)
	}
	return node, nil
}
