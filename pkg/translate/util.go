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
