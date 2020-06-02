// Code generated by go-swagger; DO NOT EDIT.

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

// NewConnectCoreV1OptionsNamespacedPodProxyWithPathParams creates a new ConnectCoreV1OptionsNamespacedPodProxyWithPathParams object
// no default values defined in spec.
func NewConnectCoreV1OptionsNamespacedPodProxyWithPathParams() ConnectCoreV1OptionsNamespacedPodProxyWithPathParams {

	return ConnectCoreV1OptionsNamespacedPodProxyWithPathParams{}
}

// ConnectCoreV1OptionsNamespacedPodProxyWithPathParams contains all the bound params for the connect core v1 options namespaced pod proxy with path operation
// typically these are obtained from a http.Request
//
// swagger:parameters connectCoreV1OptionsNamespacedPodProxyWithPath
type ConnectCoreV1OptionsNamespacedPodProxyWithPathParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*name of the PodProxyOptions
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
	/*path to the resource
	  Required: true
	  Unique: true
	  In: path
	*/
	PathPath string
	/*Path is the URL path to use for the current proxy request to pod.
	  Unique: true
	  In: query
	*/
	QueryPath *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewConnectCoreV1OptionsNamespacedPodProxyWithPathParams() beforehand.
func (o *ConnectCoreV1OptionsNamespacedPodProxyWithPathParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	rName, rhkName, _ := route.Params.GetOK("name")
	if err := o.bindName(rName, rhkName, route.Formats); err != nil {
		res = append(res, err)
	}

	rNamespace, rhkNamespace, _ := route.Params.GetOK("namespace")
	if err := o.bindNamespace(rNamespace, rhkNamespace, route.Formats); err != nil {
		res = append(res, err)
	}

	rPath, rhkPath, _ := route.Params.GetOK("path")
	if err := o.bindPathPath(rPath, rhkPath, route.Formats); err != nil {
		res = append(res, err)
	}

	qPath, qhkPath, _ := qs.GetOK("path")
	if err := o.bindQueryPath(qPath, qhkPath, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindName binds and validates parameter Name from path.
func (o *ConnectCoreV1OptionsNamespacedPodProxyWithPathParams) bindName(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
func (o *ConnectCoreV1OptionsNamespacedPodProxyWithPathParams) validateName(formats strfmt.Registry) error {

	return nil
}

// bindNamespace binds and validates parameter Namespace from path.
func (o *ConnectCoreV1OptionsNamespacedPodProxyWithPathParams) bindNamespace(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
func (o *ConnectCoreV1OptionsNamespacedPodProxyWithPathParams) validateNamespace(formats strfmt.Registry) error {

	return nil
}

// bindPathPath binds and validates parameter PathPath from path.
func (o *ConnectCoreV1OptionsNamespacedPodProxyWithPathParams) bindPathPath(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.PathPath = raw

	if err := o.validatePathPath(formats); err != nil {
		return err
	}

	return nil
}

// validatePathPath carries on validations for parameter PathPath
func (o *ConnectCoreV1OptionsNamespacedPodProxyWithPathParams) validatePathPath(formats strfmt.Registry) error {

	return nil
}

// bindQueryPath binds and validates parameter QueryPath from query.
func (o *ConnectCoreV1OptionsNamespacedPodProxyWithPathParams) bindQueryPath(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.QueryPath = &raw

	if err := o.validateQueryPath(formats); err != nil {
		return err
	}

	return nil
}

// validateQueryPath carries on validations for parameter QueryPath
func (o *ConnectCoreV1OptionsNamespacedPodProxyWithPathParams) validateQueryPath(formats strfmt.Registry) error {

	return nil
}
