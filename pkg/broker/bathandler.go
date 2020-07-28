package broker

import (
	"github.com/davecgh/go-spew/spew"
	log "github.com/sirupsen/logrus"
	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/models"
	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/pkg/translate"
)

func handleBatMessage(b *broker, msg translate.BatMessage) {
	for _, event := range msg.Events {
		switch event.Type {
		case "SIMULATION_BEGINS":
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
			b.unfinishedJobs++
			b.currentSimulationTimestep = b.baseSimulationTimestep // reset backoff

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
			log.Debugf("[broker:bathandler] Got notify : %s", spew.Sdump(event))
			if event.Data["type"] == "no_more_static_job_to_submit" {
				b.noMoreJobs = true
			}

		case "JOB_COMPLETED":
			log.Debugln("[broker:bathandler] Deserializing JOB_COMPLETED event")
			b.unfinishedJobs--
			b.runningJobs--
			b.currentSimulationTimestep = b.baseSimulationTimestep // reset backoff

			var jobCompleted translate.JobCompletedData
			if err := translate.DeserializeJobCompleted(event.Data, &jobCompleted); err != nil {
				log.Panic("[broker:bathandler] Error deserializing JOB_COMPLETED event: ", err)
			}
			log.Tracef("[broker: bathandler] Job data :\n%s", spew.Sdump(jobCompleted))

			switch jobCompleted.JobState {
			case "COMPLETED_SUCCESSFULLY":
				podName := translate.GetPodNameFromJobId(jobCompleted.JobId)
				res, i, _ := GetResource(&podName, nil, PodList)
				//res, _, _ := GetResource(&podName, nil, PodList)
				pod := res.(*models.IoK8sAPICoreV1Pod)

				pod.Status.Phase = "Succeeded"

				// update conditions and conditions
				var falseBool bool
				falseStr := "False"
				var exitCode0 int32
				currMetaV1Time := translate.BatsimNowToMetaV1Time(msg.Now)
				pod.Status.ContainerStatuses = []*models.IoK8sAPICoreV1ContainerStatus{
					{
						Image: &pod.Spec.Containers[0].Image,
						Name:  pod.Spec.Containers[0].Name,
						Ready: &falseBool,
						LastState: &models.IoK8sAPICoreV1ContainerState{
							Terminated: &models.IoK8sAPICoreV1ContainerStateTerminated{
								ExitCode:   &exitCode0,
								FinishedAt: &currMetaV1Time,
								Reason:     "Completed",
							},
						},
					},
				}
				for _, condition := range pod.Status.Conditions {
					if *condition.Type == "Ready" || *condition.Type == "ContainersReady" {
						condition.Status = &falseStr
						condition.LastTransitionTime = &currMetaV1Time
					}
				}
				IncrementResourceVersion(pod.Metadata)
				AddEvent(&translate.Modified, pod)

				// Remove the pod from the pod list
				currentTime := translate.BatsimNowToMetaV1Time(msg.Now)
				pod.Metadata.DeletionTimestamp = &currentTime
				IncrementResourceVersion(pod.Metadata)
				AddEvent(&translate.Deleted, pod)
				n := len(PodList.Items)
				PodList.Items[n-1], PodList.Items[i] = PodList.Items[i], PodList.Items[n-1]
				PodList.Items = PodList.Items[:n-1]

				log.Infof("[broker:bathandler] pod %s completed successfully. %d left to execute (%d running, %d pending)", podName, b.unfinishedJobs, b.runningJobs, b.unfinishedJobs-b.runningJobs)
			default:
				log.Errorf("[broker:bathandler] I don't know about this job state: %s", jobCompleted.JobState)

			}

		case "REQUESTED_CALL":
			b.requestedCalls = b.requestedCalls[1:]

		case "SIMULATION_ENDS":
			b.receivedSimulationEnded = true
			log.Infoln("[broker:bathandler] Bye bye")
			// TODO : gracefully shutdown the server

		default:
			log.Errorf("[broker:bathandler] I don't know about this event type: %s", event.Type)
		}
	}
}
