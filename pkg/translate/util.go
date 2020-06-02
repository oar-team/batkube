package translate

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"

	"github.com/mitchellh/mapstructure"
	"github.com/sirupsen/logrus"
	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/models"
)

/*
Underlying fields must be concrete types, not pointers

Returns the underlying value from a struct corresponding to the given tag.
Recursively explores the struct
*/
func GetValueFromTags(o interface{}, tags []string) interface{} {
	var i int
	t := reflect.TypeOf(o)
	v := reflect.ValueOf(o)

	for ; i < t.NumField() && !strings.Contains(t.Field(i).Tag.Get("json"), tags[0]); i++ {
	}

	if i < t.NumField() {
		if len(tags) == 1 {
			return v.Field(i).Interface()
		}
		return GetValueFromTags(v.Field(i).Interface(), tags[1:])
	}
	var nothing interface{}
	return nothing
}

func FilterEventListOnKind(events []models.IoK8sApimachineryPkgApisMetaV1WatchEvent, kind string) []models.IoK8sApimachineryPkgApisMetaV1WatchEvent {
	var result []models.IoK8sApimachineryPkgApisMetaV1WatchEvent
	for _, event := range events {
		if FilterEventOnObjectKind(event, kind) {
			result = append(result, event)
		}
	}
	return result
}

func FilterEventOnObjectKind(event models.IoK8sApimachineryPkgApisMetaV1WatchEvent, kind string) bool {
	return FilterObjectOnKind(event.Object, kind)
}

func FilterObjectOnKind(o interface{}, kind string) bool {
	switch o.(type) {
	case models.IoK8sAPICoreV1Pod, *models.IoK8sAPICoreV1Pod:
		return strings.EqualFold(kind, "pod")
	case models.IoK8sAPICoreV1Node, *models.IoK8sAPICoreV1Node:
		return strings.EqualFold(kind, "node")
	default:
		logrus.Warnf("[util:FilterOnObjectKind] Unknown object type : %T", o)
		return false
	}
}

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
