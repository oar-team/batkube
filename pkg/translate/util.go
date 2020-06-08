package translate

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/mitchellh/mapstructure"
	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/models"
)

func GetJobIdFromPodName(name string) string {
	return strings.ReplaceAll(name, "-", "!")
}

func GetPodNameFromJobId(name string) string {
	return strings.ReplaceAll(name, "!", "-")
}

func IncrementPodResourceVersion(pod *models.IoK8sAPICoreV1Pod) {
	resourceVersion, err := strconv.Atoi(pod.Metadata.ResourceVersion)
	if err != nil {
		log.Panic(err)
	}
	pod.Metadata.ResourceVersion = fmt.Sprintf("%d", resourceVersion+1)
}

/*
Wrapper function top help with Batsim event creation
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
