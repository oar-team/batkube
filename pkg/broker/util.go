package broker

import (
	"reflect"
	"strings"

	"github.com/sirupsen/logrus"
	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/models"
)

func FilterEventListOnKind(events []models.IoK8sApimachineryPkgApisMetaV1WatchEvent, kind string) []models.IoK8sApimachineryPkgApisMetaV1WatchEvent {
	var result []models.IoK8sApimachineryPkgApisMetaV1WatchEvent
	for _, event := range events {
		if FilterObjectOnKind(event, kind) {
			result = append(result, event)
		}
	}
	return result
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

/*
Returns the underlying value from a struct corresponding to the given tag.
Recursively explores the struct

Parameters :
o : object from which to read the value
tag : string representing the tagged field you want to extract. Ex : status.phase
*/
func GetValueFromTag(o interface{}, tag string) interface{} {
	v := indirect(reflect.ValueOf(o))

	tagsliced := strings.Split(tag, ".") // slice representation of the tag

	var i int
	t := reflect.TypeOf(v)
	for ; i < t.NumField() && !strings.Contains(t.Field(i).Tag.Get("json"), tagsliced[0]); i++ {
	}

	if i < t.NumField() {
		if len(tagsliced) == 1 {
			return v.Field(i).Interface()
		}
		return GetValueFromTag(v.Field(i).Interface(), strings.Join(tagsliced[1:], "."))
	}
	return nil
}

func FilterOnTag(items interface{}, tag string, value interface{}) []interface{} {
	filtered := make([]interface{}, 0)
	v := indirect(reflect.ValueOf(items))
	if v.Kind() == reflect.Slice {
		for i := 0; i < v.Len(); i++ {
			if GetValueFromTag(v.Index(i), tag) == value {
				filtered = append(filtered, v.Index(i))
			}
		}
	} else {
		if GetValueFromTag(v, tag) == value {
			filtered = append(filtered, items) // Return the original type, not the redirected one
		}
	}

	return filtered
}

func indirect(v reflect.Value) reflect.Value {
	for v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	return v
}
