// Code generated by go-swagger; DO NOT EDIT.

package storage_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

// NewReadStorageV1VolumeAttachmentStatusParams creates a new ReadStorageV1VolumeAttachmentStatusParams object
// no default values defined in spec.
func NewReadStorageV1VolumeAttachmentStatusParams() ReadStorageV1VolumeAttachmentStatusParams {

	return ReadStorageV1VolumeAttachmentStatusParams{}
}

// ReadStorageV1VolumeAttachmentStatusParams contains all the bound params for the read storage v1 volume attachment status operation
// typically these are obtained from a http.Request
//
// swagger:parameters readStorageV1VolumeAttachmentStatus
type ReadStorageV1VolumeAttachmentStatusParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*name of the VolumeAttachment
	  Required: true
	  Unique: true
	  In: path
	*/
	Name string
	/*If 'true', then the output is pretty printed.
	  Unique: true
	  In: query
	*/
	Pretty *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewReadStorageV1VolumeAttachmentStatusParams() beforehand.
func (o *ReadStorageV1VolumeAttachmentStatusParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	rName, rhkName, _ := route.Params.GetOK("name")
	if err := o.bindName(rName, rhkName, route.Formats); err != nil {
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

// bindName binds and validates parameter Name from path.
func (o *ReadStorageV1VolumeAttachmentStatusParams) bindName(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
func (o *ReadStorageV1VolumeAttachmentStatusParams) validateName(formats strfmt.Registry) error {

	return nil
}

// bindPretty binds and validates parameter Pretty from query.
func (o *ReadStorageV1VolumeAttachmentStatusParams) bindPretty(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
func (o *ReadStorageV1VolumeAttachmentStatusParams) validatePretty(formats strfmt.Registry) error {

	return nil
}
