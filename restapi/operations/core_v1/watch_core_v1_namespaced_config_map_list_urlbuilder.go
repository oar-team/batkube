// Code generated by go-swagger; DO NOT EDIT.

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"

	"github.com/go-openapi/swag"
)

// WatchCoreV1NamespacedConfigMapListURL generates an URL for the watch core v1 namespaced config map list operation
type WatchCoreV1NamespacedConfigMapListURL struct {
	Namespace string

	AllowWatchBookmarks *bool
	Continue            *string
	FieldSelector       *string
	LabelSelector       *string
	Limit               *int64
	Pretty              *string
	ResourceVersion     *string
	TimeoutSeconds      *int64
	Watch               *bool

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *WatchCoreV1NamespacedConfigMapListURL) WithBasePath(bp string) *WatchCoreV1NamespacedConfigMapListURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *WatchCoreV1NamespacedConfigMapListURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *WatchCoreV1NamespacedConfigMapListURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/api/v1/watch/namespaces/{namespace}/configmaps"

	namespace := o.Namespace
	if namespace != "" {
		_path = strings.Replace(_path, "{namespace}", namespace, -1)
	} else {
		return nil, errors.New("namespace is required on WatchCoreV1NamespacedConfigMapListURL")
	}

	_basePath := o._basePath
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	var allowWatchBookmarksQ string
	if o.AllowWatchBookmarks != nil {
		allowWatchBookmarksQ = swag.FormatBool(*o.AllowWatchBookmarks)
	}
	if allowWatchBookmarksQ != "" {
		qs.Set("allowWatchBookmarks", allowWatchBookmarksQ)
	}

	var continueVarQ string
	if o.Continue != nil {
		continueVarQ = *o.Continue
	}
	if continueVarQ != "" {
		qs.Set("continue", continueVarQ)
	}

	var fieldSelectorQ string
	if o.FieldSelector != nil {
		fieldSelectorQ = *o.FieldSelector
	}
	if fieldSelectorQ != "" {
		qs.Set("fieldSelector", fieldSelectorQ)
	}

	var labelSelectorQ string
	if o.LabelSelector != nil {
		labelSelectorQ = *o.LabelSelector
	}
	if labelSelectorQ != "" {
		qs.Set("labelSelector", labelSelectorQ)
	}

	var limitQ string
	if o.Limit != nil {
		limitQ = swag.FormatInt64(*o.Limit)
	}
	if limitQ != "" {
		qs.Set("limit", limitQ)
	}

	var prettyQ string
	if o.Pretty != nil {
		prettyQ = *o.Pretty
	}
	if prettyQ != "" {
		qs.Set("pretty", prettyQ)
	}

	var resourceVersionQ string
	if o.ResourceVersion != nil {
		resourceVersionQ = *o.ResourceVersion
	}
	if resourceVersionQ != "" {
		qs.Set("resourceVersion", resourceVersionQ)
	}

	var timeoutSecondsQ string
	if o.TimeoutSeconds != nil {
		timeoutSecondsQ = swag.FormatInt64(*o.TimeoutSeconds)
	}
	if timeoutSecondsQ != "" {
		qs.Set("timeoutSeconds", timeoutSecondsQ)
	}

	var watchQ string
	if o.Watch != nil {
		watchQ = swag.FormatBool(*o.Watch)
	}
	if watchQ != "" {
		qs.Set("watch", watchQ)
	}

	_result.RawQuery = qs.Encode()

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *WatchCoreV1NamespacedConfigMapListURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *WatchCoreV1NamespacedConfigMapListURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *WatchCoreV1NamespacedConfigMapListURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on WatchCoreV1NamespacedConfigMapListURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on WatchCoreV1NamespacedConfigMapListURL")
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
func (o *WatchCoreV1NamespacedConfigMapListURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
