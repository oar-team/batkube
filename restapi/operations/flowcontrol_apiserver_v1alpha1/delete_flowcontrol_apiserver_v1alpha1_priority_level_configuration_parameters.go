// Code generated by go-swagger; DO NOT EDIT.

package flowcontrol_apiserver_v1alpha1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/oar-team/batkube/models"
)

// NewDeleteFlowcontrolApiserverV1alpha1PriorityLevelConfigurationParams creates a new DeleteFlowcontrolApiserverV1alpha1PriorityLevelConfigurationParams object
// no default values defined in spec.
func NewDeleteFlowcontrolApiserverV1alpha1PriorityLevelConfigurationParams() DeleteFlowcontrolApiserverV1alpha1PriorityLevelConfigurationParams {

	return DeleteFlowcontrolApiserverV1alpha1PriorityLevelConfigurationParams{}
}

// DeleteFlowcontrolApiserverV1alpha1PriorityLevelConfigurationParams contains all the bound params for the delete flowcontrol apiserver v1alpha1 priority level configuration operation
// typically these are obtained from a http.Request
//
// swagger:parameters deleteFlowcontrolApiserverV1alpha1PriorityLevelConfiguration
type DeleteFlowcontrolApiserverV1alpha1PriorityLevelConfigurationParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  In: body
	*/
	Body *models.IoK8sApimachineryPkgApisMetaV1DeleteOptions
	/*When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed
	  Unique: true
	  In: query
	*/
	DryRun *string
	/*The duration in seconds before the object should be deleted. Value must be non-negative integer. The value zero indicates delete immediately. If this value is nil, the default grace period for the specified type will be used. Defaults to a per object value if not specified. zero means delete immediately.
	  Unique: true
	  In: query
	*/
	GracePeriodSeconds *int64
	/*name of the PriorityLevelConfiguration
	  Required: true
	  Unique: true
	  In: path
	*/
	Name string
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
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewDeleteFlowcontrolApiserverV1alpha1PriorityLevelConfigurationParams() beforehand.
func (o *DeleteFlowcontrolApiserverV1alpha1PriorityLevelConfigurationParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

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
	qDryRun, qhkDryRun, _ := qs.GetOK("dryRun")
	if err := o.bindDryRun(qDryRun, qhkDryRun, route.Formats); err != nil {
		res = append(res, err)
	}

	qGracePeriodSeconds, qhkGracePeriodSeconds, _ := qs.GetOK("gracePeriodSeconds")
	if err := o.bindGracePeriodSeconds(qGracePeriodSeconds, qhkGracePeriodSeconds, route.Formats); err != nil {
		res = append(res, err)
	}

	rName, rhkName, _ := route.Params.GetOK("name")
	if err := o.bindName(rName, rhkName, route.Formats); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindDryRun binds and validates parameter DryRun from query.
func (o *DeleteFlowcontrolApiserverV1alpha1PriorityLevelConfigurationParams) bindDryRun(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
func (o *DeleteFlowcontrolApiserverV1alpha1PriorityLevelConfigurationParams) validateDryRun(formats strfmt.Registry) error {

	return nil
}

// bindGracePeriodSeconds binds and validates parameter GracePeriodSeconds from query.
func (o *DeleteFlowcontrolApiserverV1alpha1PriorityLevelConfigurationParams) bindGracePeriodSeconds(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
func (o *DeleteFlowcontrolApiserverV1alpha1PriorityLevelConfigurationParams) validateGracePeriodSeconds(formats strfmt.Registry) error {

	return nil
}

// bindName binds and validates parameter Name from path.
func (o *DeleteFlowcontrolApiserverV1alpha1PriorityLevelConfigurationParams) bindName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.Name = raw

	if err := o.validateName(formats); err != nil {
		return err
	}

	return nil
}

// validateName carries on validations for parameter Name
func (o *DeleteFlowcontrolApiserverV1alpha1PriorityLevelConfigurationParams) validateName(formats strfmt.Registry) error {

	return nil
}

// bindOrphanDependents binds and validates parameter OrphanDependents from query.
func (o *DeleteFlowcontrolApiserverV1alpha1PriorityLevelConfigurationParams) bindOrphanDependents(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
func (o *DeleteFlowcontrolApiserverV1alpha1PriorityLevelConfigurationParams) validateOrphanDependents(formats strfmt.Registry) error {

	return nil
}

// bindPretty binds and validates parameter Pretty from query.
func (o *DeleteFlowcontrolApiserverV1alpha1PriorityLevelConfigurationParams) bindPretty(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
func (o *DeleteFlowcontrolApiserverV1alpha1PriorityLevelConfigurationParams) validatePretty(formats strfmt.Registry) error {

	return nil
}

// bindPropagationPolicy binds and validates parameter PropagationPolicy from query.
func (o *DeleteFlowcontrolApiserverV1alpha1PriorityLevelConfigurationParams) bindPropagationPolicy(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
func (o *DeleteFlowcontrolApiserverV1alpha1PriorityLevelConfigurationParams) validatePropagationPolicy(formats strfmt.Registry) error {

	return nil
}
