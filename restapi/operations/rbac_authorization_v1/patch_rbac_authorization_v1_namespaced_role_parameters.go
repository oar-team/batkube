// Code generated by go-swagger; DO NOT EDIT.

package rbac_authorization_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"gitlab.com/ryax-tech/internships/2020/scheduling_simulation/batkube/models"
)

// NewPatchRbacAuthorizationV1NamespacedRoleParams creates a new PatchRbacAuthorizationV1NamespacedRoleParams object
// no default values defined in spec.
func NewPatchRbacAuthorizationV1NamespacedRoleParams() PatchRbacAuthorizationV1NamespacedRoleParams {

	return PatchRbacAuthorizationV1NamespacedRoleParams{}
}

// PatchRbacAuthorizationV1NamespacedRoleParams contains all the bound params for the patch rbac authorization v1 namespaced role operation
// typically these are obtained from a http.Request
//
// swagger:parameters patchRbacAuthorizationV1NamespacedRole
type PatchRbacAuthorizationV1NamespacedRoleParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: body
	*/
	Body models.IoK8sApimachineryPkgApisMetaV1Patch
	/*When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed
	  Unique: true
	  In: query
	*/
	DryRun *string
	/*fieldManager is a name associated with the actor or entity that is making these changes. The value must be less than or 128 characters long, and only contain printable characters, as defined by https://golang.org/pkg/unicode/#IsPrint. This field is required for apply requests (application/apply-patch) but optional for non-apply patch types (JsonPatch, MergePatch, StrategicMergePatch).
	  Unique: true
	  In: query
	*/
	FieldManager *string
	/*Force is going to "force" Apply requests. It means user will re-acquire conflicting fields owned by other people. Force flag must be unset for non-apply patch requests.
	  Unique: true
	  In: query
	*/
	Force *bool
	/*name of the Role
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
// To ensure default values, the struct must have been initialized with NewPatchRbacAuthorizationV1NamespacedRoleParams() beforehand.
func (o *PatchRbacAuthorizationV1NamespacedRoleParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.IoK8sApimachineryPkgApisMetaV1Patch
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("body", "body"))
			} else {
				res = append(res, errors.NewParseError("body", "body", "", err))
			}
		} else {
			// no validation on generic interface
			o.Body = body
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

	qForce, qhkForce, _ := qs.GetOK("force")
	if err := o.bindForce(qForce, qhkForce, route.Formats); err != nil {
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
func (o *PatchRbacAuthorizationV1NamespacedRoleParams) bindDryRun(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
func (o *PatchRbacAuthorizationV1NamespacedRoleParams) validateDryRun(formats strfmt.Registry) error {

	return nil
}

// bindFieldManager binds and validates parameter FieldManager from query.
func (o *PatchRbacAuthorizationV1NamespacedRoleParams) bindFieldManager(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
func (o *PatchRbacAuthorizationV1NamespacedRoleParams) validateFieldManager(formats strfmt.Registry) error {

	return nil
}

// bindForce binds and validates parameter Force from query.
func (o *PatchRbacAuthorizationV1NamespacedRoleParams) bindForce(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
		return errors.InvalidType("force", "query", "bool", raw)
	}
	o.Force = &value

	if err := o.validateForce(formats); err != nil {
		return err
	}

	return nil
}

// validateForce carries on validations for parameter Force
func (o *PatchRbacAuthorizationV1NamespacedRoleParams) validateForce(formats strfmt.Registry) error {

	return nil
}

// bindName binds and validates parameter Name from path.
func (o *PatchRbacAuthorizationV1NamespacedRoleParams) bindName(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
func (o *PatchRbacAuthorizationV1NamespacedRoleParams) validateName(formats strfmt.Registry) error {

	return nil
}

// bindNamespace binds and validates parameter Namespace from path.
func (o *PatchRbacAuthorizationV1NamespacedRoleParams) bindNamespace(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
func (o *PatchRbacAuthorizationV1NamespacedRoleParams) validateNamespace(formats strfmt.Registry) error {

	return nil
}

// bindPretty binds and validates parameter Pretty from query.
func (o *PatchRbacAuthorizationV1NamespacedRoleParams) bindPretty(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
func (o *PatchRbacAuthorizationV1NamespacedRoleParams) validatePretty(formats strfmt.Registry) error {

	return nil
}
