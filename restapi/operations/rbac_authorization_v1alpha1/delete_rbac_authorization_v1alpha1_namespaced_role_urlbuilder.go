// Code generated by go-swagger; DO NOT EDIT.

package rbac_authorization_v1alpha1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"

	"github.com/go-openapi/swag"
)

// DeleteRbacAuthorizationV1alpha1NamespacedRoleURL generates an URL for the delete rbac authorization v1alpha1 namespaced role operation
type DeleteRbacAuthorizationV1alpha1NamespacedRoleURL struct {
	Name      string
	Namespace string

	DryRun             *string
	GracePeriodSeconds *int64
	OrphanDependents   *bool
	Pretty             *string
	PropagationPolicy  *string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *DeleteRbacAuthorizationV1alpha1NamespacedRoleURL) WithBasePath(bp string) *DeleteRbacAuthorizationV1alpha1NamespacedRoleURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *DeleteRbacAuthorizationV1alpha1NamespacedRoleURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *DeleteRbacAuthorizationV1alpha1NamespacedRoleURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/apis/rbac.authorization.k8s.io/v1alpha1/namespaces/{namespace}/roles/{name}"

	name := o.Name
	if name != "" {
		_path = strings.Replace(_path, "{name}", name, -1)
	} else {
		return nil, errors.New("name is required on DeleteRbacAuthorizationV1alpha1NamespacedRoleURL")
	}

	namespace := o.Namespace
	if namespace != "" {
		_path = strings.Replace(_path, "{namespace}", namespace, -1)
	} else {
		return nil, errors.New("namespace is required on DeleteRbacAuthorizationV1alpha1NamespacedRoleURL")
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

	var gracePeriodSecondsQ string
	if o.GracePeriodSeconds != nil {
		gracePeriodSecondsQ = swag.FormatInt64(*o.GracePeriodSeconds)
	}
	if gracePeriodSecondsQ != "" {
		qs.Set("gracePeriodSeconds", gracePeriodSecondsQ)
	}

	var orphanDependentsQ string
	if o.OrphanDependents != nil {
		orphanDependentsQ = swag.FormatBool(*o.OrphanDependents)
	}
	if orphanDependentsQ != "" {
		qs.Set("orphanDependents", orphanDependentsQ)
	}

	var prettyQ string
	if o.Pretty != nil {
		prettyQ = *o.Pretty
	}
	if prettyQ != "" {
		qs.Set("pretty", prettyQ)
	}

	var propagationPolicyQ string
	if o.PropagationPolicy != nil {
		propagationPolicyQ = *o.PropagationPolicy
	}
	if propagationPolicyQ != "" {
		qs.Set("propagationPolicy", propagationPolicyQ)
	}

	_result.RawQuery = qs.Encode()

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *DeleteRbacAuthorizationV1alpha1NamespacedRoleURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *DeleteRbacAuthorizationV1alpha1NamespacedRoleURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *DeleteRbacAuthorizationV1alpha1NamespacedRoleURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on DeleteRbacAuthorizationV1alpha1NamespacedRoleURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on DeleteRbacAuthorizationV1alpha1NamespacedRoleURL")
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
func (o *DeleteRbacAuthorizationV1alpha1NamespacedRoleURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
