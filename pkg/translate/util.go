package translate

import (
	"strings"

	"github.com/mitchellh/mapstructure"
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
