package broker

import (
	"github.com/davecgh/go-spew/spew"
	log "github.com/sirupsen/logrus"
	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/models"
	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/pkg/translate"
)

func handleBatMessage(msg translate.BatMessage) {
	expectedEmptyResponse = false
	for _, event := range msg.Events {
		switch event.Type {
		case "SIMULATION_BEGINS":
			expectedEmptyResponse = true
			log.Debugln("[broker:bathandler] Deserializing SIMULATION_BEGINS event")
			if err := translate.DeserializeSimulationBegins(event.Data, &simData); err != nil {
				log.Panic("[broker:bathandler] Error deserializing SIMULATION_BEGINS event: ", err)
			}
			log.Tracef("[broker:bathandler] Simulation data :\n%s", spew.Sdump(simData))

			// Translate to nodes objects
			log.Debugln("[broker:bathandler] Translating compute resources to nodes")
			err, nodesSlice := translate.ComputeResourcesToNodes(simData)
			NodeList.Items = nodesSlice
			if err != nil {
				log.Panic("[broker:bathandler] error translating compute resources to nodes: ", err)
			}

			// Add events to event list
			for _, node := range NodeList.Items {
				AddEvent(&translate.Added, node)
			}

			var nodeList []string
			for _, node := range NodeList.Items {
				nodeList = append(nodeList, node.Metadata.Name)
			}
			log.Infof("[broker:bathandler] Available nodes : %s", nodeList)

		case "JOB_SUBMITTED":
			log.Debugln("[broker:bathandler] Deserializing JOB_SUBMITTED event")
			unfinishedJobs++
			var job translate.Job
			if err := translate.DeserializeJobSubmitted(event.Data, &job); err != nil {
				log.Panic("[broker:bathandler] Error deserializing JOB_SUBMITTED event: ", err)
			}
			log.Tracef("[broker:bathandler] Job data :\n%s", spew.Sdump(job))

			// Translate
			log.Debugln("[broker:bathandler] Translating a job to a pod")
			err, pod := translate.JobToPod(job, simData)
			if err != nil {
				log.Panic("[broker:bathandler] error translating a job to a pod: ", err)
			}
			PodList.Items = append(PodList.Items, &pod)
			IncrementResourceVersion(PodList.Metadata)
			AddEvent(&translate.Modified, PodList)
			log.Tracef("pods : %s", spew.Sdump(PodList))
			log.Infof("[broker:bathandler] pod %s is ready to be scheduled", pod.Metadata.Name)

			// Add to event list
			AddEvent(&translate.Added, &pod)

		case "NOTIFY":
			expectedEmptyResponse = true
			log.Debugf("[broker:bathandler] Got notify : %s", spew.Sdump(event))
			if event.Data["type"] == "no_more_static_job_to_submit" {
				noMoreJobs = true
			}

		case "JOB_COMPLETED":
			expectedEmptyResponse = true
			log.Debugln("[broker:bathandler] Deserializing JOB_COMPLETED event")
			var jobCompleted translate.JobCompletedData
			if err := translate.DeserializeJobCompleted(event.Data, &jobCompleted); err != nil {
				log.Panic("[broker:bathandler] Error deserializing JOB_COMPLETED event: ", err)
			}
			log.Tracef("[broker: bathandler] Job data :\n%s", spew.Sdump(jobCompleted))

			switch jobCompleted.JobState {
			case "COMPLETED_SUCCESSFULLY":
				unfinishedJobs--
				podName := translate.GetPodNameFromJobId(jobCompleted.JobId)
				res, i, _ := GetResource(&podName, nil, PodList)
				pod := res.(*models.IoK8sAPICoreV1Pod)

				pod.Status.Phase = "Succeeded"
				pod.Spec.NodeName = ""
				IncrementResourceVersion(pod.Metadata)
				AddEvent(&translate.Modified, pod)

				currentTime := translate.BatsimNowToMetaV1Time(msg.Now)
				pod.Metadata.DeletionTimestamp = &currentTime
				IncrementResourceVersion(pod.Metadata)
				AddEvent(&translate.Deleted, pod)

				// Remove the pod from the pod list
				n := len(PodList.Items)
				PodList.Items[n-1], PodList.Items[i] = PodList.Items[i], PodList.Items[n-1]
				PodList.Items = PodList.Items[:n-1]
				log.Infof("[broker:bathandler] pod %s completed successfully. %d left to execute", podName, unfinishedJobs)
			default:
				log.Errorf("[broker:bathandler] I don't know about this job state: %s", jobCompleted.JobState)

			}

		case "REQUESTED_CALL":
			callMeLaters--

		case "SIMULATION_ENDS":
			expectedEmptyResponse = true
			log.Infoln("[broker:bathandler] Bye bye")
			// TODO : gracefully shutdown the server

		default:
			log.Errorf("[broker:bathandler] I don't know about this event type: %s", event.Type)
		}
	}
}
