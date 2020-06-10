package broker

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/models"
)

/*
Filters the given resourcelist object items based on the given filter fucntion.

resourceList must be an indirect type.
*/
func FilterResourceList(resourceList interface{}, filterCondition string, filter func(interface{}, string) (bool, error)) (interface{}, error) {
	if reflect.ValueOf(resourceList).Kind() != reflect.Ptr {
		return nil, errors.Errorf("An indirect type is required as input")
	}

	// Could not find a better way, as we can't define interface methods with swagger.
	switch resourceList.(type) {
	case *models.IoK8sAPICoreV1PodList:
		concreteResourceList := resourceList.(*models.IoK8sAPICoreV1PodList)
		resourceListShallowCopy := &models.IoK8sAPICoreV1PodList{}
		resourceListShallowCopy.APIVersion = concreteResourceList.APIVersion
		resourceListShallowCopy.Kind = concreteResourceList.Kind
		resourceListShallowCopy.Metadata = concreteResourceList.Metadata

		selectedItems := make([]*models.IoK8sAPICoreV1Pod, 0)
		var ok bool
		var err error
		for _, item := range concreteResourceList.Items {
			ok, err = filter(item, filterCondition)
			if err != nil {
				return nil, err
			} else if ok {
				selectedItems = append(selectedItems, item)
			}
		}

		resourceListShallowCopy.Items = selectedItems
		return resourceListShallowCopy, nil

	case *models.IoK8sAPICoreV1NodeList:
		concreteResourceList := resourceList.(*models.IoK8sAPICoreV1NodeList)
		resourceListShallowCopy := &models.IoK8sAPICoreV1NodeList{}
		resourceListShallowCopy.APIVersion = concreteResourceList.APIVersion
		resourceListShallowCopy.Kind = concreteResourceList.Kind
		resourceListShallowCopy.Metadata = concreteResourceList.Metadata

		selectedItems := make([]*models.IoK8sAPICoreV1Node, 0)
		var ok bool
		var err error
		for _, item := range concreteResourceList.Items {
			ok, err = filter(item, filterCondition)
			if err != nil {
				return nil, err
			} else if ok {
				selectedItems = append(selectedItems, item)
			}
		}

		resourceListShallowCopy.Items = selectedItems
		return resourceListShallowCopy, nil

	default:
		return nil, errors.Errorf("I don't know this resource type : %T", resourceList)
	}
}

func FilterObjectOnKind(o interface{}, kind string) (bool, error) {
	v := indirect(reflect.ValueOf(o))
	if v.Kind() == reflect.Interface {
		v = indirect(v.Elem())
	}
	if v.Kind() != reflect.Struct {
		return false, errors.Errorf("%T is neither a struct nor an interface", o)
	}
	fieldName := "Kind"
	fieldValue := v.FieldByName(fieldName)
	if fieldValue.IsZero() {
		return false, errors.Errorf("Could not find %s in %T fields", o, fieldName)
	}
	return strings.EqualFold(kind, fieldValue.String()), nil
}

func FilterObjectOnResourceVersion(o interface{}, resourceVersion string) (bool, error) {
	expected, err := strconv.Atoi(resourceVersion)
	if err != nil {
		panic(err)
	}

	v := indirect(reflect.ValueOf(o))
	if v.Kind() == reflect.Interface {
		v = indirect(v.Elem())
	}
	if v.Kind() != reflect.Struct {
		return false, errors.Errorf("%T is neither a struct nor an interface", o)
	}
	fieldName := "Metadata"
	fieldValue := v.FieldByName(fieldName)
	if fieldValue.IsZero() {
		return false, errors.Errorf("Could not find %s in %T fields", o, fieldName)
	}

	metadata, ok := fieldValue.Interface().(*models.IoK8sApimachineryPkgApisMetaV1ObjectMeta)
	if !ok {
		return false, errors.Errorf("Field %s of %T is %s (was excpecting *models.IoK8sApimachineryPkgApisMetaV1ObjectMeta)", fieldName, o, fieldValue.Type().String())
	}
	rv, err := strconv.Atoi(metadata.ResourceVersion)
	return rv >= expected, nil
}

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

func FilterEventListOnKind(events []*models.IoK8sApimachineryPkgApisMetaV1WatchEvent, kind string) []*models.IoK8sApimachineryPkgApisMetaV1WatchEvent {
	var result []*models.IoK8sApimachineryPkgApisMetaV1WatchEvent
	var ok bool
	for _, event := range events {
		ok, _ = FilterObjectOnKind(&event.Object, kind)
		if ok {
			result = append(result, event)
		}
	}
	return result
}

func FilterEventListOnResourceVersion(events []*models.IoK8sApimachineryPkgApisMetaV1WatchEvent, resourceVersion string) []*models.IoK8sApimachineryPkgApisMetaV1WatchEvent {
	var result []*models.IoK8sApimachineryPkgApisMetaV1WatchEvent
	var ok bool
	for _, event := range events {
		ok, _ = FilterObjectOnResourceVersion(event.Object, resourceVersion)
		if ok {
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
