// Code generated by go-swagger; DO NOT EDIT.

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"
)

// ConnectCoreV1OptionsNamespacedServiceProxyURL generates an URL for the connect core v1 options namespaced service proxy operation
type ConnectCoreV1OptionsNamespacedServiceProxyURL struct {
	Name      string
	Namespace string

	Path *string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *ConnectCoreV1OptionsNamespacedServiceProxyURL) WithBasePath(bp string) *ConnectCoreV1OptionsNamespacedServiceProxyURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *ConnectCoreV1OptionsNamespacedServiceProxyURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *ConnectCoreV1OptionsNamespacedServiceProxyURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/api/v1/namespaces/{namespace}/services/{name}/proxy"

	name := o.Name
	if name != "" {
		_path = strings.Replace(_path, "{name}", name, -1)
	} else {
		return nil, errors.New("name is required on ConnectCoreV1OptionsNamespacedServiceProxyURL")
	}

	namespace := o.Namespace
	if namespace != "" {
		_path = strings.Replace(_path, "{namespace}", namespace, -1)
	} else {
		return nil, errors.New("namespace is required on ConnectCoreV1OptionsNamespacedServiceProxyURL")
	}

	_basePath := o._basePath
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	var pathQ string
	if o.Path != nil {
		pathQ = *o.Path
	}
	if pathQ != "" {
		qs.Set("path", pathQ)
	}

	_result.RawQuery = qs.Encode()

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *ConnectCoreV1OptionsNamespacedServiceProxyURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *ConnectCoreV1OptionsNamespacedServiceProxyURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *ConnectCoreV1OptionsNamespacedServiceProxyURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on ConnectCoreV1OptionsNamespacedServiceProxyURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on ConnectCoreV1OptionsNamespacedServiceProxyURL")
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
func (o *ConnectCoreV1OptionsNamespacedServiceProxyURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
