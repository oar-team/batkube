// Code generated by go-swagger; DO NOT EDIT.

package node_v1alpha1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
)

// CreateNodeV1alpha1RuntimeClassURL generates an URL for the create node v1alpha1 runtime class operation
type CreateNodeV1alpha1RuntimeClassURL struct {
	DryRun       *string
	FieldManager *string
	Pretty       *string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *CreateNodeV1alpha1RuntimeClassURL) WithBasePath(bp string) *CreateNodeV1alpha1RuntimeClassURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *CreateNodeV1alpha1RuntimeClassURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *CreateNodeV1alpha1RuntimeClassURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/apis/node.k8s.io/v1alpha1/runtimeclasses"

	_basePath := o._basePath
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	var dryRunQ string
	if o.DryRun != nil {
		dryRunQ = *o.DryRun
	}
	if dryRunQ != "" {
		qs.Set("dryRun", dryRunQ)
	}

	var fieldManagerQ string
	if o.FieldManager != nil {
		fieldManagerQ = *o.FieldManager
	}
	if fieldManagerQ != "" {
		qs.Set("fieldManager", fieldManagerQ)
	}

	var prettyQ string
	if o.Pretty != nil {
		prettyQ = *o.Pretty
	}
	if prettyQ != "" {
		qs.Set("pretty", prettyQ)
	}

	_result.RawQuery = qs.Encode()

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *CreateNodeV1alpha1RuntimeClassURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *CreateNodeV1alpha1RuntimeClassURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *CreateNodeV1alpha1RuntimeClassURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on CreateNodeV1alpha1RuntimeClassURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on CreateNodeV1alpha1RuntimeClassURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *CreateNodeV1alpha1RuntimeClassURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
