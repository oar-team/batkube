package broker

import (
	"github.com/davecgh/go-spew/spew"
	log "github.com/sirupsen/logrus"
	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/models"
	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/pkg/translate"
)

func handleBatMessage(msg translate.BatMessage) {
	for _, event := range msg.Events {
		switch event.Type {
		case "SIMULATION_BEGINS":
			log.Debugln("[broker:bathandler] Deserializing SIMULATION_BEGINS event")
			if err := translate.DeserializeSimulationBegins(event.Data, &SimData); err != nil {
				log.Panic("[broker:bathandler] Error deserializing SIMULATION_BEGINS event: ", err)
			}
			log.Tracef("[broker:bathandler] Simulation data :\n%s", spew.Sdump(SimData))

			// Translate to nodes objects
			log.Debugln("[broker:bathandler] Translating compute resources to nodes")
			err, nodesSlice := translate.ComputeResourcesToNodes(SimData.ComputeResources)
			NodeList.Items = nodesSlice
			if err != nil {
				log.Panic("[broker:bathandler] error translating compute resources to nodes: ", err)
			}

			// Add events to event list
			for _, node := range NodeList.Items {
				AddEvent(&models.IoK8sApimachineryPkgApisMetaV1WatchEvent{
					Type:   &translate.Added,
					Object: node,
				})
			}

			var nodeList []string
			for _, node := range NodeList.Items {
				nodeList = append(nodeList, node.Metadata.Name)
			}
			log.Infof("[broker:bathandler] Available nodes : %s", nodeList)

		case "JOB_SUBMITTED":
			log.Debugln("[broker:bathandler] Deserializing JOB_SUBMITTED event")
			var job translate.Job
			if err := translate.DeserializeJobSubmitted(event.Data, &job); err != nil {
				log.Panic("[broker:bathandler] Error deserializing JOB_SUBMITTED event: ", err)
			}
			log.Tracef("[broker:bathandler] Job data :\n%s", spew.Sdump(job))

			// Translate
			log.Debugln("[broker:bathandler] Translating a job to a pod")
			err, pod := translate.JobToPod(job, SimData)
			if err != nil {
				log.Panic("[broker:bathandler] error translating a job to a pod: ", err)
			}
			PodList.Items = append(PodList.Items, &pod)
			log.Tracef("pods : %s", spew.Sdump(PodList))
			log.Infof("[broker:bathandler] pod %s is ready to be scheduled", pod.Metadata.Name)

			// Add to event list
			AddEvent(&models.IoK8sApimachineryPkgApisMetaV1WatchEvent{
				Type:   &translate.Added,
				Object: pod,
			})

		case "NOTIFY":
			log.Infoln("[broker:bathandler] Got NOTIFY")

		case "JOB_COMPLETED":
			log.Debugln("[broker:bathandler] Deserializing JOB_COMPLETED event")
			var jobCompleted translate.JobCompletedData
			if err := translate.DeserializeJobCompleted(event.Data, &jobCompleted); err != nil {
				log.Panic("[broker:bathandler] Error deserializing JOB_COMPLETED event: ", err)
			}
			log.Tracef("[broker: bathandler] Job data :\n%s", spew.Sdump(jobCompleted))

			switch jobCompleted.JobState {
			case "COMPLETED_SUCCESSFULLY":
				pod, _ := GetPod(translate.GetPodNameFromJobId(jobCompleted.JobId))
				pod.Status.Phase = "Succeeded"
				IncrementPodResourceVersion(pod)
				// TODO : unbind the pod?
				AddEvent(&models.IoK8sApimachineryPkgApisMetaV1WatchEvent{
					Type:   &translate.Modified,
					Object: pod,
				})
			default:
				log.Errorf("[broker:bathandler] I don't know about this job state: %s", jobCompleted.JobState)

			}

		case "REQUESTED_CALL":

		case "SIMULATION_ENDS":
			log.Infoln("[broker:bathandler] Bye bye")
			// TODO : gracefully shutdown the server

		default:
			log.Errorf("[broker:bathandler] I don't know about this event type: %s", event.Type)
		}
	}
}
