// Code generated by go-swagger; DO NOT EDIT.

package node_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/models"
)

// NewDeleteNodeV1beta1CollectionRuntimeClassParams creates a new DeleteNodeV1beta1CollectionRuntimeClassParams object
// no default values defined in spec.
func NewDeleteNodeV1beta1CollectionRuntimeClassParams() DeleteNodeV1beta1CollectionRuntimeClassParams {

	return DeleteNodeV1beta1CollectionRuntimeClassParams{}
}

// DeleteNodeV1beta1CollectionRuntimeClassParams contains all the bound params for the delete node v1beta1 collection runtime class operation
// typically these are obtained from a http.Request
//
// swagger:parameters deleteNodeV1beta1CollectionRuntimeClass
type DeleteNodeV1beta1CollectionRuntimeClassParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*allowWatchBookmarks requests watch events with type "BOOKMARK". Servers that do not implement bookmarks may ignore this flag and bookmarks are sent at the server's discretion. Clients should not assume bookmarks are returned at any specific interval, nor may they assume the server will send any BOOKMARK event during a session. If this is not a watch, this field is ignored. If the feature gate WatchBookmarks is not enabled in apiserver, this field is ignored.
	  Unique: true
	  In: query
	*/
	AllowWatchBookmarks *bool
	/*
	  In: body
	*/
	Body *models.IoK8sApimachineryPkgApisMetaV1DeleteOptions
	/*The continue option should be set when retrieving more results from the server. Since this value is server defined, clients may only use the continue value from a previous query result with identical query parameters (except for the value of continue) and the server may reject a continue value it does not recognize. If the specified continue value is no longer valid whether due to expiration (generally five to fifteen minutes) or a configuration change on the server, the server will respond with a 410 ResourceExpired error together with a continue token. If the client needs a consistent list, it must restart their list without the continue field. Otherwise, the client may send another list request with the token received with the 410 error, the server will respond with a list starting from the next key, but from the latest snapshot, which is inconsistent from the previous list results - objects that are created, modified, or deleted after the first list request will be included in the response, as long as their keys are after the "next key".

	This field is not supported when watch is true. Clients may start a watch from the last resourceVersion value returned by the server and not miss any modifications.
	  Unique: true
	  In: query
	*/
	Continue *string
	/*When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed
	  Unique: true
	  In: query
	*/
	DryRun *string
	/*A selector to restrict the list of returned objects by their fields. Defaults to everything.
	  Unique: true
	  In: query
	*/
	FieldSelector *string
	/*The duration in seconds before the object should be deleted. Value must be non-negative integer. The value zero indicates delete immediately. If this value is nil, the default grace period for the specified type will be used. Defaults to a per object value if not specified. zero means delete immediately.
	  Unique: true
	  In: query
	*/
	GracePeriodSeconds *int64
	/*A selector to restrict the list of returned objects by their labels. Defaults to everything.
	  Unique: true
	  In: query
	*/
	LabelSelector *string
	/*limit is a maximum number of responses to return for a list call. If more items exist, the server will set the `continue` field on the list metadata to a value that can be used with the same initial query to retrieve the next set of results. Setting a limit may return fewer than the requested amount of items (up to zero items) in the event all requested objects are filtered out and clients should only use the presence of the continue field to determine whether more results are available. Servers may choose not to support the limit argument and will return all of the available results. If limit is specified and the continue field is empty, clients may assume that no more results are available. This field is not supported if watch is true.

	The server guarantees that the objects returned when using continue will be identical to issuing a single list call without a limit - that is, no objects created, modified, or deleted after the first request is issued will be included in any subsequent continued requests. This is sometimes referred to as a consistent snapshot, and ensures that a client that is using limit to receive smaller chunks of a very large result can ensure they see all possible objects. If objects are updated during a chunked list the version of the object that was present at the time the first list result was calculated is returned.
	  Unique: true
	  In: query
	*/
	Limit *int64
	/*Deprecated: please use the PropagationPolicy, this field will be deprecated in 1.7. Should the dependent objects be orphaned. If true/false, the "orphan" finalizer will be added to/removed from the object's finalizers list. Either this field or PropagationPolicy may be set, but not both.
	  Unique: true
	  In: query
	*/
	OrphanDependents *bool
	/*If 'true', then the output is pretty printed.
	  Unique: true
	  In: query
	*/
	Pretty *string
	/*Whether and how garbage collection will be performed. Either this field or OrphanDependents may be set, but not both. The default policy is decided by the existing finalizer set in the metadata.finalizers and the resource-specific default policy. Acceptable values are: 'Orphan' - orphan the dependents; 'Background' - allow the garbage collector to delete the dependents in the background; 'Foreground' - a cascading policy that deletes all dependents in the foreground.
	  Unique: true
	  In: query
	*/
	PropagationPolicy *string
	/*When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history. When specified for list: - if unset, then the result is returned from remote storage based on quorum-read flag; - if it's 0, then we simply return what we currently have in cache, no guarantee; - if set to non zero, then the result is at least as fresh as given rv.
	  Unique: true
	  In: query
	*/
	ResourceVersion *string
	/*Timeout for the list/watch call. This limits the duration of the call, regardless of any activity or inactivity.
	  Unique: true
	  In: query
	*/
	TimeoutSeconds *int64
	/*Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.
	  Unique: true
	  In: query
	*/
	Watch *bool
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewDeleteNodeV1beta1CollectionRuntimeClassParams() beforehand.
func (o *DeleteNodeV1beta1CollectionRuntimeClassParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qAllowWatchBookmarks, qhkAllowWatchBookmarks, _ := qs.GetOK("allowWatchBookmarks")
	if err := o.bindAllowWatchBookmarks(qAllowWatchBookmarks, qhkAllowWatchBookmarks, route.Formats); err != nil {
		res = append(res, err)
	}

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.IoK8sApimachineryPkgApisMetaV1DeleteOptions
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			res = append(res, errors.NewParseError("body", "body", "", err))
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Body = &body
			}
		}
	}
	qContinue, qhkContinue, _ := qs.GetOK("continue")
	if err := o.bindContinue(qContinue, qhkContinue, route.Formats); err != nil {
		res = append(res, err)
	}

	qDryRun, qhkDryRun, _ := qs.GetOK("dryRun")
	if err := o.bindDryRun(qDryRun, qhkDryRun, route.Formats); err != nil {
		res = append(res, err)
	}

	qFieldSelector, qhkFieldSelector, _ := qs.GetOK("fieldSelector")
	if err := o.bindFieldSelector(qFieldSelector, qhkFieldSelector, route.Formats); err != nil {
		res = append(res, err)
	}

	qGracePeriodSeconds, qhkGracePeriodSeconds, _ := qs.GetOK("gracePeriodSeconds")
	if err := o.bindGracePeriodSeconds(qGracePeriodSeconds, qhkGracePeriodSeconds, route.Formats); err != nil {
		res = append(res, err)
	}

	qLabelSelector, qhkLabelSelector, _ := qs.GetOK("labelSelector")
	if err := o.bindLabelSelector(qLabelSelector, qhkLabelSelector, route.Formats); err != nil {
		res = append(res, err)
	}

	qLimit, qhkLimit, _ := qs.GetOK("limit")
	if err := o.bindLimit(qLimit, qhkLimit, route.Formats); err != nil {
		res = append(res, err)
	}

	qOrphanDependents, qhkOrphanDependents, _ := qs.GetOK("orphanDependents")
	if err := o.bindOrphanDependents(qOrphanDependents, qhkOrphanDependents, route.Formats); err != nil {
		res = append(res, err)
	}

	qPretty, qhkPretty, _ := qs.GetOK("pretty")
	if err := o.bindPretty(qPretty, qhkPretty, route.Formats); err != nil {
		res = append(res, err)
	}

	qPropagationPolicy, qhkPropagationPolicy, _ := qs.GetOK("propagationPolicy")
	if err := o.bindPropagationPolicy(qPropagationPolicy, qhkPropagationPolicy, route.Formats); err != nil {
		res = append(res, err)
	}

	qResourceVersion, qhkResourceVersion, _ := qs.GetOK("resourceVersion")
	if err := o.bindResourceVersion(qResourceVersion, qhkResourceVersion, route.Formats); err != nil {
		res = append(res, err)
	}

	qTimeoutSeconds, qhkTimeoutSeconds, _ := qs.GetOK("timeoutSeconds")
	if err := o.bindTimeoutSeconds(qTimeoutSeconds, qhkTimeoutSeconds, route.Formats); err != nil {
		res = append(res, err)
	}

	qWatch, qhkWatch, _ := qs.GetOK("watch")
	if err := o.bindWatch(qWatch, qhkWatch, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindAllowWatchBookmarks binds and validates parameter AllowWatchBookmarks from query.
func (o *DeleteNodeV1beta1CollectionRuntimeClassParams) bindAllowWatchBookmarks(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertBool(raw)
	if err != nil {
		return errors.InvalidType("allowWatchBookmarks", "query", "bool", raw)
	}
	o.AllowWatchBookmarks = &value

	if err := o.validateAllowWatchBookmarks(formats); err != nil {
		return err
	}

	return nil
}

// validateAllowWatchBookmarks carries on validations for parameter AllowWatchBookmarks
func (o *DeleteNodeV1beta1CollectionRuntimeClassParams) validateAllowWatchBookmarks(formats strfmt.Registry) error {

	return nil
}

// bindContinue binds and validates parameter Continue from query.
func (o *DeleteNodeV1beta1CollectionRuntimeClassParams) bindContinue(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Continue = &raw

	if err := o.validateContinue(formats); err != nil {
		return err
	}

	return nil
}

// validateContinue carries on validations for parameter Continue
func (o *DeleteNodeV1beta1CollectionRuntimeClassParams) validateContinue(formats strfmt.Registry) error {

	return nil
}

// bindDryRun binds and validates parameter DryRun from query.
func (o *DeleteNodeV1beta1CollectionRuntimeClassParams) bindDryRun(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.DryRun = &raw

	if err := o.validateDryRun(formats); err != nil {
		return err
	}

	return nil
}

// validateDryRun carries on validations for parameter DryRun
func (o *DeleteNodeV1beta1CollectionRuntimeClassParams) validateDryRun(formats strfmt.Registry) error {

	return nil
}

// bindFieldSelector binds and validates parameter FieldSelector from query.
func (o *DeleteNodeV1beta1CollectionRuntimeClassParams) bindFieldSelector(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.FieldSelector = &raw

	if err := o.validateFieldSelector(formats); err != nil {
		return err
	}

	return nil
}

// validateFieldSelector carries on validations for parameter FieldSelector
func (o *DeleteNodeV1beta1CollectionRuntimeClassParams) validateFieldSelector(formats strfmt.Registry) error {

	return nil
}

// bindGracePeriodSeconds binds and validates parameter GracePeriodSeconds from query.
func (o *DeleteNodeV1beta1CollectionRuntimeClassParams) bindGracePeriodSeconds(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("gracePeriodSeconds", "query", "int64", raw)
	}
	o.GracePeriodSeconds = &value

	if err := o.validateGracePeriodSeconds(formats); err != nil {
		return err
	}

	return nil
}

// validateGracePeriodSeconds carries on validations for parameter GracePeriodSeconds
func (o *DeleteNodeV1beta1CollectionRuntimeClassParams) validateGracePeriodSeconds(formats strfmt.Registry) error {

	return nil
}

// bindLabelSelector binds and validates parameter LabelSelector from query.
func (o *DeleteNodeV1beta1CollectionRuntimeClassParams) bindLabelSelector(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.LabelSelector = &raw

	if err := o.validateLabelSelector(formats); err != nil {
		return err
	}

	return nil
}

// validateLabelSelector carries on validations for parameter LabelSelector
func (o *DeleteNodeV1beta1CollectionRuntimeClassParams) validateLabelSelector(formats strfmt.Registry) error {

	return nil
}

// bindLimit binds and validates parameter Limit from query.
func (o *DeleteNodeV1beta1CollectionRuntimeClassParams) bindLimit(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("limit", "query", "int64", raw)
	}
	o.Limit = &value

	if err := o.validateLimit(formats); err != nil {
		return err
	}

	return nil
}

// validateLimit carries on validations for parameter Limit
func (o *DeleteNodeV1beta1CollectionRuntimeClassParams) validateLimit(formats strfmt.Registry) error {

	return nil
}

// bindOrphanDependents binds and validates parameter OrphanDependents from query.
func (o *DeleteNodeV1beta1CollectionRuntimeClassParams) bindOrphanDependents(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertBool(raw)
	if err != nil {
		return errors.InvalidType("orphanDependents", "query", "bool", raw)
	}
	o.OrphanDependents = &value

	if err := o.validateOrphanDependents(formats); err != nil {
		return err
	}

	return nil
}

// validateOrphanDependents carries on validations for parameter OrphanDependents
func (o *DeleteNodeV1beta1CollectionRuntimeClassParams) validateOrphanDependents(formats strfmt.Registry) error {

	return nil
}

// bindPretty binds and validates parameter Pretty from query.
func (o *DeleteNodeV1beta1CollectionRuntimeClassParams) bindPretty(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Pretty = &raw

	if err := o.validatePretty(formats); err != nil {
		return err
	}

	return nil
}

// validatePretty carries on validations for parameter Pretty
func (o *DeleteNodeV1beta1CollectionRuntimeClassParams) validatePretty(formats strfmt.Registry) error {

	return nil
}

// bindPropagationPolicy binds and validates parameter PropagationPolicy from query.
func (o *DeleteNodeV1beta1CollectionRuntimeClassParams) bindPropagationPolicy(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.PropagationPolicy = &raw

	if err := o.validatePropagationPolicy(formats); err != nil {
		return err
	}

	return nil
}

// validatePropagationPolicy carries on validations for parameter PropagationPolicy
func (o *DeleteNodeV1beta1CollectionRuntimeClassParams) validatePropagationPolicy(formats strfmt.Registry) error {

	return nil
}

// bindResourceVersion binds and validates parameter ResourceVersion from query.
func (o *DeleteNodeV1beta1CollectionRuntimeClassParams) bindResourceVersion(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.ResourceVersion = &raw

	if err := o.validateResourceVersion(formats); err != nil {
		return err
	}

	return nil
}

// validateResourceVersion carries on validations for parameter ResourceVersion
func (o *DeleteNodeV1beta1CollectionRuntimeClassParams) validateResourceVersion(formats strfmt.Registry) error {

	return nil
}

// bindTimeoutSeconds binds and validates parameter TimeoutSeconds from query.
func (o *DeleteNodeV1beta1CollectionRuntimeClassParams) bindTimeoutSeconds(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("timeoutSeconds", "query", "int64", raw)
	}
	o.TimeoutSeconds = &value

	if err := o.validateTimeoutSeconds(formats); err != nil {
		return err
	}

	return nil
}

// validateTimeoutSeconds carries on validations for parameter TimeoutSeconds
func (o *DeleteNodeV1beta1CollectionRuntimeClassParams) validateTimeoutSeconds(formats strfmt.Registry) error {

	return nil
}

// bindWatch binds and validates parameter Watch from query.
func (o *DeleteNodeV1beta1CollectionRuntimeClassParams) bindWatch(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertBool(raw)
	if err != nil {
		return errors.InvalidType("watch", "query", "bool", raw)
	}
	o.Watch = &value

	if err := o.validateWatch(formats); err != nil {
		return err
	}

	return nil
}

// validateWatch carries on validations for parameter Watch
func (o *DeleteNodeV1beta1CollectionRuntimeClassParams) validateWatch(formats strfmt.Registry) error {

	return nil
}
