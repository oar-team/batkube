// Code generated by go-swagger; DO NOT EDIT.

package rbac_authorization_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"

	"github.com/go-openapi/swag"
)

// PatchRbacAuthorizationV1beta1ClusterRoleURL generates an URL for the patch rbac authorization v1beta1 cluster role operation
type PatchRbacAuthorizationV1beta1ClusterRoleURL struct {
	Name string

	DryRun       *string
	FieldManager *string
	Force        *bool
	Pretty       *string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *PatchRbacAuthorizationV1beta1ClusterRoleURL) WithBasePath(bp string) *PatchRbacAuthorizationV1beta1ClusterRoleURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *PatchRbacAuthorizationV1beta1ClusterRoleURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *PatchRbacAuthorizationV1beta1ClusterRoleURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/apis/rbac.authorization.k8s.io/v1beta1/clusterroles/{name}"

	name := o.Name
	if name != "" {
		_path = strings.Replace(_path, "{name}", name, -1)
	} else {
		return nil, errors.New("name is required on PatchRbacAuthorizationV1beta1ClusterRoleURL")
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

	var forceQ string
	if o.Force != nil {
		forceQ = swag.FormatBool(*o.Force)
	}
	if forceQ != "" {
		qs.Set("force", forceQ)
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
func (o *PatchRbacAuthorizationV1beta1ClusterRoleURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *PatchRbacAuthorizationV1beta1ClusterRoleURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *PatchRbacAuthorizationV1beta1ClusterRoleURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on PatchRbacAuthorizationV1beta1ClusterRoleURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on PatchRbacAuthorizationV1beta1ClusterRoleURL")
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
func (o *PatchRbacAuthorizationV1beta1ClusterRoleURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
