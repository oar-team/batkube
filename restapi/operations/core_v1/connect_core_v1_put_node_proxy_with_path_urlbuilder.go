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

// ConnectCoreV1PutNodeProxyWithPathURL generates an URL for the connect core v1 put node proxy with path operation
type ConnectCoreV1PutNodeProxyWithPathURL struct {
	Name     string
	PathPath string

	QueryPath *string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *ConnectCoreV1PutNodeProxyWithPathURL) WithBasePath(bp string) *ConnectCoreV1PutNodeProxyWithPathURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *ConnectCoreV1PutNodeProxyWithPathURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *ConnectCoreV1PutNodeProxyWithPathURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/api/v1/nodes/{name}/proxy/{path}"

	name := o.Name
	if name != "" {
		_path = strings.Replace(_path, "{name}", name, -1)
	} else {
		return nil, errors.New("name is required on ConnectCoreV1PutNodeProxyWithPathURL")
	}

	pathPath := o.PathPath
	if pathPath != "" {
		_path = strings.Replace(_path, "{path}", pathPath, -1)
	} else {
		return nil, errors.New("pathPath is required on ConnectCoreV1PutNodeProxyWithPathURL")
	}

	_basePath := o._basePath
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	var queryPathQ string
	if o.QueryPath != nil {
		queryPathQ = *o.QueryPath
	}
	if queryPathQ != "" {
		qs.Set("path", queryPathQ)
	}

	_result.RawQuery = qs.Encode()

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *ConnectCoreV1PutNodeProxyWithPathURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *ConnectCoreV1PutNodeProxyWithPathURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *ConnectCoreV1PutNodeProxyWithPathURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on ConnectCoreV1PutNodeProxyWithPathURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on ConnectCoreV1PutNodeProxyWithPathURL")
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
func (o *ConnectCoreV1PutNodeProxyWithPathURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
