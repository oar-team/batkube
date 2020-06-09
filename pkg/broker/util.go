package broker

import (
	"reflect"
	"strings"

	"github.com/pkg/errors"
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
func getValueFromTag(o interface{}, tag string) (string, error) {
	v := indirect(reflect.ValueOf(o))
	t := v.Type()
	if v.Kind() != reflect.Struct {
		return "", errors.Errorf("Expected a struct, got a %s\n", v.Kind().String())
	}

	tagsliced := strings.Split(tag, ".") // slice representation of the tag

	var i int
	for ; i < t.NumField(); i++ {
		tagContent, ok := t.Field(i).Tag.Lookup("json")
		if ok && strings.Contains(tagContent, tagsliced[0]) {
			break
		}
	}

	if i < t.NumField() {
		if len(tagsliced) == 1 {
			if v.Field(i).Kind() != reflect.String {
				return "", errors.Errorf("Only fields containing strings are supported")
			}
			return v.Field(i).String(), nil
		}
		value, err := getValueFromTag(v.Field(i).Interface(), strings.Join(tagsliced[1:], "."))
		if err != nil {
			return "", errors.Errorf("Error looking for %s in %s : %s", tag, t.String(), err)
		}
		return value, nil
	}
	return "", errors.Errorf("Type %s does not contain any field %s", t.String(), tag)
}

/*
Returns whether the given struct complies with the given fieldSelector
*/
func FilterOnFieldSelector(o interface{}, selectors string) (bool, error) {
	selectorsSlice := strings.Split(selectors, ",")
	for _, selector := range selectorsSlice {
		if strings.Contains(selector, "!=") {
			selectorSlice := strings.Split(selector, "!=")
			value, err := getValueFromTag(o, selectorSlice[0])
			if err != nil {
				return false, err
			} else if value == selectorSlice[1] {
				return false, nil
			}
		} else if strings.Contains(selector, "=") {
			selectorSlice := strings.Split(selector, "=")
			value, err := getValueFromTag(o, selectorSlice[0])
			if err != nil {
				return false, err
			} else if value != selectorSlice[1] {
				return false, nil
			}
		} else {
			return false, errors.Errorf("Wrong fieldSelector : %s", selectors)
		}
	}
	return true, nil
}

func indirect(v reflect.Value) reflect.Value {
	for v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	return v
}
