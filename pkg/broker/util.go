package broker

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/models"
	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/pkg/translate"
)

/*
Returns the resource corresponding to the given fields.
Returned type is a pointer to the resource.
Also returns the index of the resource in the list (returns -1 if the resource could not be found)
*/
func GetResource(name *string, namespace *string, resourceList interface{}) (interface{}, int, error) {
	v := indirect(reflect.ValueOf(resourceList))
	itemsValue, err := getFieldByName(v, "Items")
	if err != nil {
		return nil, -1, err
	}
	//itemsValue = itemsValue.Elem()
	if itemsValue.Kind() != reflect.Slice {
		return nil, -1, errors.Errorf("Expected a slice as Items value, got %s (type %s)", itemsValue.Kind(), itemsValue.Type())
	}

	n := itemsValue.Len()
	for j := 0; j < n; j++ {
		item := indirect(itemsValue.Index(j))
		metadata, err := getFieldByName(item, "Metadata")
		if err != nil {
			return nil, -1, err
		}
		metadata = indirect(metadata)

		nameValue, err := getFieldByName(metadata, "Name")
		if err != nil {
			return nil, -1, err
		}
		nameValue = indirect(nameValue)
		if name != nil && nameValue.IsValid() && nameValue.String() != *name {
			continue
		}

		namespaceValue, err := getFieldByName(metadata, "Namespace")
		if err != nil {
			return nil, -1, err
		}
		namespaceValue = indirect(namespaceValue)
		if namespace != nil && namespaceValue.IsValid() && namespaceValue.String() != *namespace {
			continue
		}

		return item.Addr().Interface(), j, nil
	}
	return nil, -1, nil
}

func getFieldByName(resource reflect.Value, fieldName string) (reflect.Value, error) {
	if resource.Kind() == reflect.Ptr {
		resource = resource.Elem()
	}
	fieldValue := resource.FieldByName(fieldName)
	if !fieldValue.IsValid() {
		return reflect.Value{}, errors.Errorf("Type %s has no %s field", resource.Type(), fieldName)
	}
	return fieldValue, nil
}

func GetAPIResourceList(groupVersion string) *models.IoK8sApimachineryPkgApisMetaV1APIResourceList {
	resources, _ := apiResources[groupVersion]

	return &models.IoK8sApimachineryPkgApisMetaV1APIResourceList{
		Kind:         "APIResourceList",
		APIVersion:   "v1",
		GroupVersion: &groupVersion,
		Resources:    resources,
	}
}

func createAPIResource(groupVersion, name, kind string, namespaced bool, verbs []string) {
	apiResources[groupVersion] = append(apiResources[groupVersion], &models.IoK8sApimachineryPkgApisMetaV1APIResource{
		Name:       &name,
		Kind:       &kind,
		Namespaced: &namespaced,
		Verbs:      verbs,
	})
}

/*
Filters out items slice in filteredItems slice based on the given filter.

filteredItems must be a slice of same type as items with capacity superior or
equal to items lenght
*/
func filterItems(items, filteredItems interface{}, filterCondition string, filter func(interface{}, string) (bool, error)) error {
	itemsValue := indirect(reflect.ValueOf(items))
	filteredItemsValue := reflect.ValueOf(filteredItems)
	if reflect.ValueOf(filteredItems).Kind() != reflect.Ptr {
		return errors.Errorf("filteredItems (type %T) should be a pointer to a slice", filteredItems)
	}
	filteredItemsValue = filteredItemsValue.Elem()

	if itemsValue.Kind() != reflect.Slice {
		return errors.Errorf("items (type %T) should be a slice or pointer to a slice", itemsValue.Type)
	}
	if filteredItemsValue.Kind() != reflect.Slice {
		return errors.Errorf("filteredItems (type %T) should be a pointer to a slice", filteredItemsValue.Type)
	}

	n := itemsValue.Len()
	if filteredItemsValue.Cap() < n {
		return errors.Errorf("filteredItems capacity (%d) is less than items length (%d)", filteredItemsValue.Cap(), n)
	}
	i := 0 // current index in filteredItems
	var ok bool
	var err error
	for j := 0; j < n; j++ {
		ok, err = filter(itemsValue.Index(j).Interface(), filterCondition)
		if err != nil {
			return err
		} else if ok {
			filteredItemsValue.Index(i).Set(itemsValue.Index(j))
			i++ // At most, i = n
		}
	}
	filteredItemsValue.Set(filteredItemsValue.Slice(0, i))
	return nil
}

/*
Filters the given resourcelist object items based on the given filter fucntion.

resourceList must be an indirect type.
*/
func FilterResourceList(resourceList interface{}, filterCondition string, filter func(interface{}, string) (bool, error)) (interface{}, error) {
	if reflect.ValueOf(resourceList).Kind() != reflect.Ptr {
		return nil, errors.Errorf("ResourceList must be an indirect type. Given type : %T", resourceList)
	}

	// Could not find a better way. To my knowledge, objects cannot be initialized and manipulated without a concrete type.
	var err error
	switch resourceList.(type) {
	case *models.IoK8sAPICoreV1PodList:
		concreteResourceList := resourceList.(*models.IoK8sAPICoreV1PodList)
		resourceListShallowCopy := &models.IoK8sAPICoreV1PodList{
			APIVersion: concreteResourceList.APIVersion,
			Kind:       concreteResourceList.Kind,
			Metadata:   concreteResourceList.Metadata,
		}
		filteredItems := make([]*models.IoK8sAPICoreV1Pod, len(concreteResourceList.Items))
		if err = filterItems(&concreteResourceList.Items, &filteredItems, filterCondition, filter); err != nil {
			return nil, err
		}
		resourceListShallowCopy.Items = filteredItems
		return resourceListShallowCopy, nil
	case *models.IoK8sAPICoreV1NodeList:
		concreteResourceList := resourceList.(*models.IoK8sAPICoreV1NodeList)
		resourceListShallowCopy := &models.IoK8sAPICoreV1NodeList{
			APIVersion: concreteResourceList.APIVersion,
			Kind:       concreteResourceList.Kind,
			Metadata:   concreteResourceList.Metadata,
		}
		filteredItems := make([]*models.IoK8sAPICoreV1Node, len(concreteResourceList.Items))
		if err = filterItems(&concreteResourceList.Items, &filteredItems, filterCondition, filter); err != nil {
			return nil, err
		}
		resourceListShallowCopy.Items = filteredItems
		return resourceListShallowCopy, nil
	case *models.IoK8sAPIPolicyV1beta1PodDisruptionBudgetList:
		concreteResourceList := resourceList.(*models.IoK8sAPIPolicyV1beta1PodDisruptionBudgetList)
		resourceListShallowCopy := &models.IoK8sAPIPolicyV1beta1PodDisruptionBudgetList{
			APIVersion: concreteResourceList.APIVersion,
			Kind:       concreteResourceList.Kind,
			Metadata:   concreteResourceList.Metadata,
		}
		filteredItems := make([]*models.IoK8sAPIPolicyV1beta1PodDisruptionBudget, len(concreteResourceList.Items))
		if err = filterItems(&concreteResourceList.Items, &filteredItems, filterCondition, filter); err != nil {
			return nil, err
		}
		resourceListShallowCopy.Items = filteredItems
		return resourceListShallowCopy, nil
	case *models.IoK8sAPIStorageV1StorageClassList:
		concreteResourceList := resourceList.(*models.IoK8sAPIStorageV1StorageClassList)
		resourceListShallowCopy := &models.IoK8sAPIStorageV1StorageClassList{
			APIVersion: concreteResourceList.APIVersion,
			Kind:       concreteResourceList.Kind,
			Metadata:   concreteResourceList.Metadata,
		}
		filteredItems := make([]*models.IoK8sAPIStorageV1StorageClass, len(concreteResourceList.Items))
		if err = filterItems(&concreteResourceList.Items, &filteredItems, filterCondition, filter); err != nil {
			return nil, err
		}
		resourceListShallowCopy.Items = filteredItems
		return resourceListShallowCopy, nil
	case *models.IoK8sAPIStorageV1CSINodeList:
		concreteResourceList := resourceList.(*models.IoK8sAPIStorageV1CSINodeList)
		resourceListShallowCopy := &models.IoK8sAPIStorageV1CSINodeList{
			APIVersion: concreteResourceList.APIVersion,
			Kind:       concreteResourceList.Kind,
			Metadata:   concreteResourceList.Metadata,
		}
		filteredItems := make([]*models.IoK8sAPIStorageV1CSINode, len(concreteResourceList.Items))
		if err = filterItems(&concreteResourceList.Items, &filteredItems, filterCondition, filter); err != nil {
			return nil, err
		}
		resourceListShallowCopy.Items = filteredItems
		return resourceListShallowCopy, nil
	case *models.IoK8sAPICoreV1PersistentVolumeClaimList:
		concreteResourceList := resourceList.(*models.IoK8sAPICoreV1PersistentVolumeClaimList)
		resourceListShallowCopy := &models.IoK8sAPICoreV1PersistentVolumeClaimList{
			APIVersion: concreteResourceList.APIVersion,
			Kind:       concreteResourceList.Kind,
			Metadata:   concreteResourceList.Metadata,
		}
		filteredItems := make([]*models.IoK8sAPICoreV1PersistentVolumeClaim, len(concreteResourceList.Items))
		if err = filterItems(&concreteResourceList.Items, &filteredItems, filterCondition, filter); err != nil {
			return nil, err
		}
		resourceListShallowCopy.Items = filteredItems
		return resourceListShallowCopy, nil
	case *models.IoK8sAPICoreV1PersistentVolumeList:
		concreteResourceList := resourceList.(*models.IoK8sAPICoreV1PersistentVolumeList)
		resourceListShallowCopy := &models.IoK8sAPICoreV1PersistentVolumeList{
			APIVersion: concreteResourceList.APIVersion,
			Kind:       concreteResourceList.Kind,
			Metadata:   concreteResourceList.Metadata,
		}
		filteredItems := make([]*models.IoK8sAPICoreV1PersistentVolume, len(concreteResourceList.Items))
		if err = filterItems(&concreteResourceList.Items, &filteredItems, filterCondition, filter); err != nil {
			return nil, err
		}
		resourceListShallowCopy.Items = filteredItems
		return resourceListShallowCopy, nil
	case *models.IoK8sAPICoreV1ServiceList:
		concreteResourceList := resourceList.(*models.IoK8sAPICoreV1ServiceList)
		resourceListShallowCopy := &models.IoK8sAPICoreV1ServiceList{
			APIVersion: concreteResourceList.APIVersion,
			Kind:       concreteResourceList.Kind,
			Metadata:   concreteResourceList.Metadata,
		}
		filteredItems := make([]*models.IoK8sAPICoreV1Service, len(concreteResourceList.Items))
		if err = filterItems(&concreteResourceList.Items, &filteredItems, filterCondition, filter); err != nil {
			return nil, err
		}
		resourceListShallowCopy.Items = filteredItems
		return resourceListShallowCopy, nil
	case *models.IoK8sApimachineryPkgApisMetaV1APIResourceList:
		return nil, errors.Errorf("internal error : use GetAPIResourceList to list API resources")
	case *models.IoK8sAPICoreV1EndpointsList:
		concreteResourceList := resourceList.(*models.IoK8sAPICoreV1EndpointsList)
		resourceListShallowCopy := &models.IoK8sAPICoreV1EndpointsList{
			APIVersion: concreteResourceList.APIVersion,
			Kind:       concreteResourceList.Kind,
			Metadata:   concreteResourceList.Metadata,
		}
		filteredItems := make([]*models.IoK8sAPICoreV1Endpoints, len(concreteResourceList.Items))
		if err = filterItems(&concreteResourceList.Items, &filteredItems, filterCondition, filter); err != nil {
			return nil, err
		}
		resourceListShallowCopy.Items = filteredItems
		return resourceListShallowCopy, nil
	case *models.IoK8sAPICoordinationV1LeaseList:
		concreteResourceList := resourceList.(*models.IoK8sAPICoordinationV1LeaseList)
		resourceListShallowCopy := &models.IoK8sAPICoordinationV1LeaseList{
			APIVersion: concreteResourceList.APIVersion,
			Kind:       concreteResourceList.Kind,
			Metadata:   concreteResourceList.Metadata,
		}
		filteredItems := make([]*models.IoK8sAPICoordinationV1Lease, len(concreteResourceList.Items))
		if err = filterItems(&concreteResourceList.Items, &filteredItems, filterCondition, filter); err != nil {
			return nil, err
		}
		resourceListShallowCopy.Items = filteredItems
		return resourceListShallowCopy, nil
	case *models.IoK8sAPICoreV1EventList:
		concreteResourceList := resourceList.(*models.IoK8sAPICoreV1EventList)
		resourceListShallowCopy := &models.IoK8sAPICoreV1EventList{
			APIVersion: concreteResourceList.APIVersion,
			Kind:       concreteResourceList.Kind,
			Metadata:   concreteResourceList.Metadata,
		}
		filteredItems := make([]*models.IoK8sAPICoreV1Event, len(concreteResourceList.Items))
		if err = filterItems(&concreteResourceList.Items, &filteredItems, filterCondition, filter); err != nil {
			return nil, err
		}
		resourceListShallowCopy.Items = filteredItems
		return resourceListShallowCopy, nil
	case *models.IoK8sAPIEventsV1beta1EventList:
		concreteResourceList := resourceList.(*models.IoK8sAPIEventsV1beta1EventList)
		resourceListShallowCopy := &models.IoK8sAPIEventsV1beta1EventList{
			APIVersion: concreteResourceList.APIVersion,
			Kind:       concreteResourceList.Kind,
			Metadata:   concreteResourceList.Metadata,
		}
		filteredItems := make([]*models.IoK8sAPIEventsV1beta1Event, len(concreteResourceList.Items))
		if err = filterItems(&concreteResourceList.Items, &filteredItems, filterCondition, filter); err != nil {
			return nil, err
		}
		resourceListShallowCopy.Items = filteredItems
		return resourceListShallowCopy, nil
	default:
		return nil, errors.Errorf("I don't know this resource type : %T", resourceList)
	}
}

// TODO : these next two filters can be handled as fieldSelectors.
// This will reduce the amount of filters to one only, based on json tags.
func FilterObjectOnKind(o interface{}, kind string) (bool, error) {
	if kind == "" || kind == "*" {
		return true, nil
	}
	v := indirect(reflect.ValueOf(o))
	if v.Kind() == reflect.Interface {
		v = indirect(v.Elem())
	}
	if v.Kind() != reflect.Struct {
		return false, errors.Errorf("%T is neither a struct nor an interface", o)
	}
	fieldName := "Kind"
	fieldValue := v.FieldByName(fieldName)
	if !fieldValue.IsValid() {
		return false, errors.Errorf("Could not find %s in %T fields", fieldName, o)
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
	if !fieldValue.IsValid() {
		return false, errors.Errorf("Could not find %s in %T fields", fieldName, o)
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
	if v.Kind() != reflect.Struct {
		return "", errors.Errorf("Expected a struct, got a %s\n", v.Kind().String())
	}
	t := v.Type()

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
	return "", errors.Errorf("Type %s does not contain any field tagged %s", t.String(), tag)
}

/*
Indirects the given value, recursively
*/
// TODO : get rid of this, as we never encounter nested pointers. If we do, it is an error.
func indirect(v reflect.Value) reflect.Value {
	if v.Kind() == reflect.Ptr {
		return indirect(v.Elem())
	}
	return v
}

func incrementStr(str string) string {
	if str == "" {
		return "1"
	}
	n, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%d", n+1)
}

/*
Returns if int(str1) > int(str2)
*/
func compareStr(str1, str2 string) bool {
	//if str1 == "" {
	//	str1 = "0"
	//}
	//if str2 == "" {
	//	str2 = "0"
	//}
	n1, err := strconv.Atoi(str1)
	if err != nil {
		panic(err)
	}
	n2, err := strconv.Atoi(str2)
	if err != nil {
		panic(err)
	}
	return n1 > n2
}

/*
Naive deepcopy function, before getting real ones
*/
func DeepCopy(input, output interface{}) error {
	b, err := json.Marshal(input)
	if err != nil {
		return errors.New(fmt.Sprintf("error while marshaling %T: %s", input, err))
	}
	err = json.Unmarshal(b, output)
	if err != nil {
		return errors.New(fmt.Sprintf("error while unmarshaling into a %T: %s", output, err))
	}
	return nil
}

/*
If the given object's resource version is lower than given resource version,
set its resourceVersion to given value
Returns if the value was incremented
*/
func IncrementResourceVersionTo(metadata interface{}, value string) bool {
	rv, err := getFieldByName(reflect.ValueOf(metadata), "ResourceVersion")
	if err != nil {
		panic(err)
	}
	if compareStr(value, rv.String()) {
		rv.SetString(value)
		return true
	}
	return false
}

/*
Does IncrementResourceVersionTo and all pods and nodes, and adds a modified event for each of the increased nodes
*/
func IncrementAllResourceVersionsTo(value string) {
	for _, pod := range PodList.Items {
		if IncrementResourceVersionTo(pod.Metadata, value) {
			AddEvent(&translate.Modified, pod)
		}
	}
	for _, node := range NodeList.Items {
		if IncrementResourceVersionTo(node.Metadata, value) {
			AddEvent(&translate.Modified, node)
		}
	}
}

/*
Increments the resource version of the given metadata by 1
*/
func IncrementResourceVersion(metadata interface{}) {
	rv, err := getFieldByName(reflect.ValueOf(metadata), "ResourceVersion")
	if err != nil {
		panic(err)
	}
	rv.SetString(incrementStr(rv.String()))
}

func DeleteAllEventsPrior(resourceVersion string) {
	// TODO
	// Remove events that are too old to free memory
}
