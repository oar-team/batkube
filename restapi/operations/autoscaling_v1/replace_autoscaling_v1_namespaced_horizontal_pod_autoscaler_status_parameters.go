// Code generated by go-swagger; DO NOT EDIT.

package autoscaling_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"

	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/models"
)

// NewReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerStatusParams creates a new ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerStatusParams object
// no default values defined in spec.
func NewReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerStatusParams() ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerStatusParams {

	return ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerStatusParams{}
}

// ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerStatusParams contains all the bound params for the replace autoscaling v1 namespaced horizontal pod autoscaler status operation
// typically these are obtained from a http.Request
//
// swagger:parameters replaceAutoscalingV1NamespacedHorizontalPodAutoscalerStatus
type ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerStatusParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: body
	*/
	Body *models.IoK8sAPIAutoscalingV1HorizontalPodAutoscaler
	/*When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed
	  Unique: true
	  In: query
	*/
	DryRun *string
	/*fieldManager is a name associated with the actor or entity that is making these changes. The value must be less than or 128 characters long, and only contain printable characters, as defined by https://golang.org/pkg/unicode/#IsPrint.
	  Unique: true
	  In: query
	*/
	FieldManager *string
	/*name of the HorizontalPodAutoscaler
	  Required: true
	  Unique: true
	  In: path
	*/
	Name string
	/*object name and auth scope, such as for teams and projects
	  Required: true
	  Unique: true
	  In: path
	*/
	Namespace string
	/*If 'true', then the output is pretty printed.
	  Unique: true
	  In: query
	*/
	Pretty *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerStatusParams() beforehand.
func (o *ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerStatusParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.IoK8sAPIAutoscalingV1HorizontalPodAutoscaler
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("body", "body"))
			} else {
				res = append(res, errors.NewParseError("body", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Body = &body
			}
		}
	} else {
		res = append(res, errors.Required("body", "body"))
	}
	qDryRun, qhkDryRun, _ := qs.GetOK("dryRun")
	if err := o.bindDryRun(qDryRun, qhkDryRun, route.Formats); err != nil {
		res = append(res, err)
	}

	qFieldManager, qhkFieldManager, _ := qs.GetOK("fieldManager")
	if err := o.bindFieldManager(qFieldManager, qhkFieldManager, route.Formats); err != nil {
		res = append(res, err)
	}

	rName, rhkName, _ := route.Params.GetOK("name")
	if err := o.bindName(rName, rhkName, route.Formats); err != nil {
		res = append(res, err)
	}

	rNamespace, rhkNamespace, _ := route.Params.GetOK("namespace")
	if err := o.bindNamespace(rNamespace, rhkNamespace, route.Formats); err != nil {
		res = append(res, err)
	}

	qPretty, qhkPretty, _ := qs.GetOK("pretty")
	if err := o.bindPretty(qPretty, qhkPretty, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindDryRun binds and validates parameter DryRun from query.
func (o *ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerStatusParams) bindDryRun(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
func (o *ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerStatusParams) validateDryRun(formats strfmt.Registry) error {

	return nil
}

// bindFieldManager binds and validates parameter FieldManager from query.
func (o *ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerStatusParams) bindFieldManager(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.FieldManager = &raw

	if err := o.validateFieldManager(formats); err != nil {
		return err
	}

	return nil
}

// validateFieldManager carries on validations for parameter FieldManager
func (o *ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerStatusParams) validateFieldManager(formats strfmt.Registry) error {

	return nil
}

// bindName binds and validates parameter Name from path.
func (o *ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerStatusParams) bindName(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
func (o *ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerStatusParams) validateName(formats strfmt.Registry) error {

	return nil
}

// bindNamespace binds and validates parameter Namespace from path.
func (o *ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerStatusParams) bindNamespace(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.Namespace = raw

	if err := o.validateNamespace(formats); err != nil {
		return err
	}

	return nil
}

// validateNamespace carries on validations for parameter Namespace
func (o *ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerStatusParams) validateNamespace(formats strfmt.Registry) error {

	return nil
}

// bindPretty binds and validates parameter Pretty from query.
func (o *ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerStatusParams) bindPretty(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
func (o *ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerStatusParams) validatePretty(formats strfmt.Registry) error {

	return nil
}
