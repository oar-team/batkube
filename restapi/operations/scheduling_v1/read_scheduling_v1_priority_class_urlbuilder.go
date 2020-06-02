// Code generated by go-swagger; DO NOT EDIT.

package scheduling_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"

	"github.com/go-openapi/swag"
)

// ReadSchedulingV1PriorityClassURL generates an URL for the read scheduling v1 priority class operation
type ReadSchedulingV1PriorityClassURL struct {
	Name string

	Exact  *bool
	Export *bool
	Pretty *string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *ReadSchedulingV1PriorityClassURL) WithBasePath(bp string) *ReadSchedulingV1PriorityClassURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *ReadSchedulingV1PriorityClassURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *ReadSchedulingV1PriorityClassURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/apis/scheduling.k8s.io/v1/priorityclasses/{name}"

	name := o.Name
	if name != "" {
		_path = strings.Replace(_path, "{name}", name, -1)
	} else {
		return nil, errors.New("name is required on ReadSchedulingV1PriorityClassURL")
	}

	_basePath := o._basePath
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	var exactQ string
	if o.Exact != nil {
		exactQ = swag.FormatBool(*o.Exact)
	}
	if exactQ != "" {
		qs.Set("exact", exactQ)
	}

	var exportQ string
	if o.Export != nil {
		exportQ = swag.FormatBool(*o.Export)
	}
	if exportQ != "" {
		qs.Set("export", exportQ)
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
func (o *ReadSchedulingV1PriorityClassURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *ReadSchedulingV1PriorityClassURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *ReadSchedulingV1PriorityClassURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on ReadSchedulingV1PriorityClassURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on ReadSchedulingV1PriorityClassURL")
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
func (o *ReadSchedulingV1PriorityClassURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
