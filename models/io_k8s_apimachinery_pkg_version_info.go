// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// IoK8sApimachineryPkgVersionInfo Info contains versioning information. how we'll want to distribute that information.
//
// swagger:model io.k8s.apimachinery.pkg.version.Info
type IoK8sApimachineryPkgVersionInfo struct {

	// build date
	// Required: true
	BuildDate *string `json:"buildDate"`

	// compiler
	// Required: true
	Compiler *string `json:"compiler"`

	// git commit
	// Required: true
	GitCommit *string `json:"gitCommit"`

	// git tree state
	// Required: true
	GitTreeState *string `json:"gitTreeState"`

	// git version
	// Required: true
	GitVersion *string `json:"gitVersion"`

	// go version
	// Required: true
	GoVersion *string `json:"goVersion"`

	// major
	// Required: true
	Major *string `json:"major"`

	// minor
	// Required: true
	Minor *string `json:"minor"`

	// platform
	// Required: true
	Platform *string `json:"platform"`
}

// Validate validates this io k8s apimachinery pkg version info
func (m *IoK8sApimachineryPkgVersionInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBuildDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCompiler(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGitCommit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGitTreeState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGitVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGoVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMajor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMinor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlatform(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sApimachineryPkgVersionInfo) validateBuildDate(formats strfmt.Registry) error {

	if err := validate.Required("buildDate", "body", m.BuildDate); err != nil {
		return err
	}

	return nil
}

func (m *IoK8sApimachineryPkgVersionInfo) validateCompiler(formats strfmt.Registry) error {

	if err := validate.Required("compiler", "body", m.Compiler); err != nil {
		return err
	}

	return nil
}

func (m *IoK8sApimachineryPkgVersionInfo) validateGitCommit(formats strfmt.Registry) error {

	if err := validate.Required("gitCommit", "body", m.GitCommit); err != nil {
		return err
	}

	return nil
}

func (m *IoK8sApimachineryPkgVersionInfo) validateGitTreeState(formats strfmt.Registry) error {

	if err := validate.Required("gitTreeState", "body", m.GitTreeState); err != nil {
		return err
	}

	return nil
}

func (m *IoK8sApimachineryPkgVersionInfo) validateGitVersion(formats strfmt.Registry) error {

	if err := validate.Required("gitVersion", "body", m.GitVersion); err != nil {
		return err
	}

	return nil
}

func (m *IoK8sApimachineryPkgVersionInfo) validateGoVersion(formats strfmt.Registry) error {

	if err := validate.Required("goVersion", "body", m.GoVersion); err != nil {
		return err
	}

	return nil
}

func (m *IoK8sApimachineryPkgVersionInfo) validateMajor(formats strfmt.Registry) error {

	if err := validate.Required("major", "body", m.Major); err != nil {
		return err
	}

	return nil
}

func (m *IoK8sApimachineryPkgVersionInfo) validateMinor(formats strfmt.Registry) error {

	if err := validate.Required("minor", "body", m.Minor); err != nil {
		return err
	}

	return nil
}

func (m *IoK8sApimachineryPkgVersionInfo) validatePlatform(formats strfmt.Registry) error {

	if err := validate.Required("platform", "body", m.Platform); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sApimachineryPkgVersionInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sApimachineryPkgVersionInfo) UnmarshalBinary(b []byte) error {
	var res IoK8sApimachineryPkgVersionInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
