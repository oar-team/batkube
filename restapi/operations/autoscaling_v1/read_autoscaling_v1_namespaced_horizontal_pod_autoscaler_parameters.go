// Code generated by go-swagger; DO NOT EDIT.

package autoscaling_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewReadAutoscalingV1NamespacedHorizontalPodAutoscalerParams creates a new ReadAutoscalingV1NamespacedHorizontalPodAutoscalerParams object
// no default values defined in spec.
func NewReadAutoscalingV1NamespacedHorizontalPodAutoscalerParams() ReadAutoscalingV1NamespacedHorizontalPodAutoscalerParams {

	return ReadAutoscalingV1NamespacedHorizontalPodAutoscalerParams{}
}

// ReadAutoscalingV1NamespacedHorizontalPodAutoscalerParams contains all the bound params for the read autoscaling v1 namespaced horizontal pod autoscaler operation
// typically these are obtained from a http.Request
//
// swagger:parameters readAutoscalingV1NamespacedHorizontalPodAutoscaler
type ReadAutoscalingV1NamespacedHorizontalPodAutoscalerParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Should the export be exact.  Exact export maintains cluster-specific fields like 'Namespace'. Deprecated. Planned for removal in 1.18.
	  Unique: true
	  In: query
	*/
	Exact *bool
	/*Should this value be exported.  Export strips fields that a user can not specify. Deprecated. Planned for removal in 1.18.
	  Unique: true
	  In: query
	*/
	Export *bool
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
// To ensure default values, the struct must have been initialized with NewReadAutoscalingV1NamespacedHorizontalPodAutoscalerParams() beforehand.
func (o *ReadAutoscalingV1NamespacedHorizontalPodAutoscalerParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qExact, qhkExact, _ := qs.GetOK("exact")
	if err := o.bindExact(qExact, qhkExact, route.Formats); err != nil {
		res = append(res, err)
	}

	qExport, qhkExport, _ := qs.GetOK("export")
	if err := o.bindExport(qExport, qhkExport, route.Formats); err != nil {
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

// bindExact binds and validates parameter Exact from query.
func (o *ReadAutoscalingV1NamespacedHorizontalPodAutoscalerParams) bindExact(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
		return errors.InvalidType("exact", "query", "bool", raw)
	}
	o.Exact = &value

	if err := o.validateExact(formats); err != nil {
		return err
	}

	return nil
}

// validateExact carries on validations for parameter Exact
func (o *ReadAutoscalingV1NamespacedHorizontalPodAutoscalerParams) validateExact(formats strfmt.Registry) error {

	return nil
}

// bindExport binds and validates parameter Export from query.
func (o *ReadAutoscalingV1NamespacedHorizontalPodAutoscalerParams) bindExport(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
		return errors.InvalidType("export", "query", "bool", raw)
	}
	o.Export = &value

	if err := o.validateExport(formats); err != nil {
		return err
	}

	return nil
}

// validateExport carries on validations for parameter Export
func (o *ReadAutoscalingV1NamespacedHorizontalPodAutoscalerParams) validateExport(formats strfmt.Registry) error {

	return nil
}

// bindName binds and validates parameter Name from path.
func (o *ReadAutoscalingV1NamespacedHorizontalPodAutoscalerParams) bindName(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
func (o *ReadAutoscalingV1NamespacedHorizontalPodAutoscalerParams) validateName(formats strfmt.Registry) error {

	return nil
}

// bindNamespace binds and validates parameter Namespace from path.
func (o *ReadAutoscalingV1NamespacedHorizontalPodAutoscalerParams) bindNamespace(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
func (o *ReadAutoscalingV1NamespacedHorizontalPodAutoscalerParams) validateNamespace(formats strfmt.Registry) error {

	return nil
}

// bindPretty binds and validates parameter Pretty from query.
func (o *ReadAutoscalingV1NamespacedHorizontalPodAutoscalerParams) bindPretty(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
func (o *ReadAutoscalingV1NamespacedHorizontalPodAutoscalerParams) validatePretty(formats strfmt.Registry) error {

	return nil
}
