package broker

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/models"
)

func FilterEventListOnKind(events []models.IoK8sApimachineryPkgApisMetaV1WatchEvent, kind string) []models.IoK8sApimachineryPkgApisMetaV1WatchEvent {
	var result []models.IoK8sApimachineryPkgApisMetaV1WatchEvent
	for _, event := range events {
		if FilterObjectOnKind(&event.Object, kind) {
			result = append(result, event)
		}
	}
	return result
}

func FilterObjectOnKind(o interface{}, kind string) bool {
	if reflect.ValueOf(o).Kind() != reflect.Ptr {
		panic("Filter function requires an indirect type")
	}
	switch o.(type) {
	case *models.IoK8sAPICoreV1Pod:
		return strings.EqualFold(kind, o.(*models.IoK8sAPICoreV1Pod).Kind)
	case *models.IoK8sAPICoreV1Node:
		return strings.EqualFold(kind, o.(*models.IoK8sAPICoreV1Node).Kind)
	default:
		logrus.Warnf("[broker/util:FilterOnObjectKind] Unknown object type : %T", o)
		return false
	}
}

func FilterEventListOnResourceVersion(events []*models.IoK8sApimachineryPkgApisMetaV1WatchEvent, resourceVersion string) []*models.IoK8sApimachineryPkgApisMetaV1WatchEvent {
	var result []*models.IoK8sApimachineryPkgApisMetaV1WatchEvent
	for _, event := range events {
		if FilterObjectOnResourceVersion(event.Object, resourceVersion) {
			result = append(result, event)
		}
	}
	return result
}

func FilterEventListOnFieldSelector(events []*models.IoK8sApimachineryPkgApisMetaV1WatchEvent, fieldSelector string) ([]*models.IoK8sApimachineryPkgApisMetaV1WatchEvent, error) {
	var result []*models.IoK8sApimachineryPkgApisMetaV1WatchEvent
	var ok bool
	var err error
	for _, event := range events {
		ok, err = FilterObjectOnFieldSelector(event.Object, fieldSelector)
		if err != nil {
			return nil, err
		} else if ok {
			result = append(result, event)
		}
	}
	return result, nil
}

func FilterObjectOnResourceVersion(o interface{}, resourceVersion string) bool {
	if reflect.ValueOf(o).Kind() != reflect.Ptr {
		panic("Filter function requires an indirect type")
	}
	expected, err := strconv.Atoi(resourceVersion)
	var rv int
	if err != nil {
		panic(err)
	}
	switch o.(type) {
	case *models.IoK8sAPICoreV1Pod:
		rv, err = strconv.Atoi(o.(*models.IoK8sAPICoreV1Pod).Metadata.ResourceVersion)
		if err != nil {
			panic(err)
		}
	case *models.IoK8sAPICoreV1Node:
		rv, err = strconv.Atoi(o.(*models.IoK8sAPICoreV1Node).Metadata.ResourceVersion)
		if err != nil {
			panic(err)
		}
	default:
		logrus.Warnf("[broker/util:FilterOnObjectKind] Unknown object type : %T", o)
		return false
	}
	return rv >= expected
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
func FilterObjectOnFieldSelector(o interface{}, selectors string) (bool, error) {
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

func IncrementPodResourceVersion(pod *models.IoK8sAPICoreV1Pod) {
	resourceVersion, err := strconv.Atoi(pod.Metadata.ResourceVersion)
	if err != nil {
		log.Panic(err)
	}
	pod.Metadata.ResourceVersion = fmt.Sprintf("%d", resourceVersion+1)
}
