// Code generated by go-swagger; DO NOT EDIT.

package autoscaling_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"
)

// CreateAutoscalingV1NamespacedHorizontalPodAutoscalerURL generates an URL for the create autoscaling v1 namespaced horizontal pod autoscaler operation
type CreateAutoscalingV1NamespacedHorizontalPodAutoscalerURL struct {
	Namespace string

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
func (o *CreateAutoscalingV1NamespacedHorizontalPodAutoscalerURL) WithBasePath(bp string) *CreateAutoscalingV1NamespacedHorizontalPodAutoscalerURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *CreateAutoscalingV1NamespacedHorizontalPodAutoscalerURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *CreateAutoscalingV1NamespacedHorizontalPodAutoscalerURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/apis/autoscaling/v1/namespaces/{namespace}/horizontalpodautoscalers"

	namespace := o.Namespace
	if namespace != "" {
		_path = strings.Replace(_path, "{namespace}", namespace, -1)
	} else {
		return nil, errors.New("namespace is required on CreateAutoscalingV1NamespacedHorizontalPodAutoscalerURL")
	}

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
func (o *CreateAutoscalingV1NamespacedHorizontalPodAutoscalerURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *CreateAutoscalingV1NamespacedHorizontalPodAutoscalerURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *CreateAutoscalingV1NamespacedHorizontalPodAutoscalerURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on CreateAutoscalingV1NamespacedHorizontalPodAutoscalerURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on CreateAutoscalingV1NamespacedHorizontalPodAutoscalerURL")
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
func (o *CreateAutoscalingV1NamespacedHorizontalPodAutoscalerURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
