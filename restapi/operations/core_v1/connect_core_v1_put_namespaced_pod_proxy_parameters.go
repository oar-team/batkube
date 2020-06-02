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

// NewConnectCoreV1PutNamespacedPodProxyParams creates a new ConnectCoreV1PutNamespacedPodProxyParams object
// no default values defined in spec.
func NewConnectCoreV1PutNamespacedPodProxyParams() ConnectCoreV1PutNamespacedPodProxyParams {

	return ConnectCoreV1PutNamespacedPodProxyParams{}
}

// ConnectCoreV1PutNamespacedPodProxyParams contains all the bound params for the connect core v1 put namespaced pod proxy operation
// typically these are obtained from a http.Request
//
// swagger:parameters connectCoreV1PutNamespacedPodProxy
type ConnectCoreV1PutNamespacedPodProxyParams struct {

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
	/*Path is the URL path to use for the current proxy request to pod.
	  Unique: true
	  In: query
	*/
	Path *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewConnectCoreV1PutNamespacedPodProxyParams() beforehand.
func (o *ConnectCoreV1PutNamespacedPodProxyParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
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

	qPath, qhkPath, _ := qs.GetOK("path")
	if err := o.bindPath(qPath, qhkPath, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindName binds and validates parameter Name from path.
func (o *ConnectCoreV1PutNamespacedPodProxyParams) bindName(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
func (o *ConnectCoreV1PutNamespacedPodProxyParams) validateName(formats strfmt.Registry) error {

	return nil
}

// bindNamespace binds and validates parameter Namespace from path.
func (o *ConnectCoreV1PutNamespacedPodProxyParams) bindNamespace(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
func (o *ConnectCoreV1PutNamespacedPodProxyParams) validateNamespace(formats strfmt.Registry) error {

	return nil
}

// bindPath binds and validates parameter Path from query.
func (o *ConnectCoreV1PutNamespacedPodProxyParams) bindPath(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Path = &raw

	if err := o.validatePath(formats); err != nil {
		return err
	}

	return nil
}

// validatePath carries on validations for parameter Path
func (o *ConnectCoreV1PutNamespacedPodProxyParams) validatePath(formats strfmt.Registry) error {

	return nil
}
