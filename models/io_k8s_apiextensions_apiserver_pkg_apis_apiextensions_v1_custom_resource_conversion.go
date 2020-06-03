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

// IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceConversion CustomResourceConversion describes how to convert different versions of a CR.
//
// swagger:model io.k8s.apiextensions-apiserver.pkg.apis.apiextensions.v1.CustomResourceConversion
type IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceConversion struct {

	// strategy specifies how custom resources are converted between versions. Allowed values are: - `None`: The converter only change the apiVersion and would not touch any other field in the custom resource. - `Webhook`: API Server will call to an external webhook to do the conversion. Additional information
	//   is needed for this option. This requires spec.preserveUnknownFields to be false, and spec.conversion.webhook to be set.
	// Required: true
	Strategy *string `json:"strategy"`

	// webhook describes how to call the conversion webhook. Required when `strategy` is set to `Webhook`.
	Webhook *IoK8sApiextensionsApiserverPkgApisApiextensionsV1WebhookConversion `json:"webhook,omitempty"`
}

// Validate validates this io k8s apiextensions apiserver pkg apis apiextensions v1 custom resource conversion
func (m *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceConversion) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStrategy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWebhook(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceConversion) validateStrategy(formats strfmt.Registry) error {

	if err := validate.Required("strategy", "body", m.Strategy); err != nil {
		return err
	}

	return nil
}

func (m *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceConversion) validateWebhook(formats strfmt.Registry) error {

	if swag.IsZero(m.Webhook) { // not required
		return nil
	}

	if m.Webhook != nil {
		if err := m.Webhook.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("webhook")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceConversion) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceConversion) UnmarshalBinary(b []byte) error {
	var res IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceConversion
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}